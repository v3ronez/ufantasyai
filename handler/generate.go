package handler

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/v3ronez/ufantasyai/db"
	"github.com/v3ronez/ufantasyai/types"
	"github.com/v3ronez/ufantasyai/view/generate"
)

func HandlerGenerateIndex(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	images, err := db.GetImagesFromUserId(user.ID)
	if err != nil {
		return err
	}
	viewData := generate.ViewData{Images: images}
	return RenderComponent(w, r, generate.Index(viewData))
}
func HandlerGenerateImageCreate(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	prompt := r.FormValue("prompt")
	qtdImage, _ := strconv.Atoi(r.FormValue("amount"))
	formData := generate.FormData{
		Prompt: prompt,
		Amount: qtdImage,
	}
	var formErr = generate.FormDataErr{}
	// if err != nil {
	// 	return err
	// }
	if len(prompt) < 5 {
		formErr.PromptErr = "Invalid prompt"
		return RenderComponent(w, r, generate.Form(formData, formErr))
	}
	if qtdImage <= 0 || qtdImage > 8 {
		formErr.AmountErr = "Select the amount"
		return RenderComponent(w, r, generate.Form(formData, formErr))
	}
	err := db.Bun.RunInTx(r.Context(), &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		batchID := uuid.New()
		for i := 0; i < qtdImage; i++ {
			image := types.Image{
				UserId:        user.ID,
				Prompt:        formData.Prompt,
				BatchId:       batchID,
				Status:        types.ImageStatusPending,
				ImageLocation: "https://img.freepik.com/premium-photo/cute-girl-with-pretty-face-creative-ai_634423-2810.jpg",
			}
			if err := db.CreateImage(&image); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return err
	}
	return HtmxRedirect(w, r, "/generate")
	// image := types.Image{
	// 	ID:            2,
	// 	UserId:        user.ID,
	// 	Prompt:        "beautiful womem",
	// 	BatchId:       batchID,
	// 	Status:        types.ImageStatusPending,
	// 	ImageLocation: "https://img.freepik.com/premium-photo/cute-girl-with-pretty-face-creative-ai_634423-2810.jpg",
	// }

	// if err := db.CreateImage(&image); err != nil {
	// 	return err
	// }
	// return RenderComponent(w, r, generate.GalleryImage(image))
}

func HandlerGenerateImageStatus(w http.ResponseWriter, r *http.Request) error {
	imageID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}
	//fetch from db
	image, err := db.GetImageForID(imageID)
	if err != nil {
		return err
	}
	return RenderComponent(w, r, generate.GalleryImage(image))
}
