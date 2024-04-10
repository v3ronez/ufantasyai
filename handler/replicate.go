package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/v3ronez/ufantasyai/db"
	"github.com/v3ronez/ufantasyai/types"
)

const (
	succeded   = "succeeded"
	processing = "processing"
)

type ReplicateResponse struct {
	Input struct {
		Prompt string `json:"prompt"`
	} `json:"input"`
	Status string   `json:"status"`
	Output []string `json:"output"`
}

func HandleReplicateCallback(w http.ResponseWriter, r *http.Request) error {
	userID := chi.URLParam(r, "userID")
	_ = userID
	batchID, err := uuid.Parse(chi.URLParam(r, "batchID"))
	if err != nil {
		return err
	}
	var replicateResponse ReplicateResponse
	if err := json.NewDecoder(r.Body).Decode(&replicateResponse); err != nil {
		return err
	}
	if replicateResponse.Status == processing {
		return nil
	}
	if replicateResponse.Status != succeded {
		return fmt.Errorf("status if not succeeded: %s ", err)
	}

	images, err := db.GetImagesForBatchID(batchID)
	if err != nil {
		return fmt.Errorf("images not founds error: %s ", err)
	}

	err = db.Bun.RunInTx(r.Context(), &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		for idx, imageURL := range replicateResponse.Output {
			images[idx].Prompt = replicateResponse.Input.Prompt
			images[idx].Status = types.ImageStatusCompleted
			images[idx].ImageLocation = imageURL
			if err := db.UpdateImage(tx, &images[idx]); err != nil {
				return err
			}
		}
		return err
	})
	return err
}
