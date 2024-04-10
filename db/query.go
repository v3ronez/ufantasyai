package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/v3ronez/ufantasyai/types"
)

func GetImagesFromUserId(userId uuid.UUID) ([]types.Image, error) {
	var images []types.Image

	err := Bun.NewSelect().Model(&images).
		Where("deleted = ?", false).
		Where("user_id = ?", userId).
		Order("created_at DESC").
		Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return images, nil
}

func GetImageForID(imageID int) (types.Image, error) {
	var image types.Image

	err := Bun.NewSelect().
		Model(&image).
		Where("id = ?", imageID).
		Scan(context.Background())
	return image, err
}

func GetImagesForBatchID(batchID uuid.UUID) ([]types.Image, error) {
	var images []types.Image
	err := Bun.NewSelect().
		Model(&images).
		Where("batch_id = ?", batchID).
		Scan(context.Background())
	return images, err
}

func CreateImage(tx bun.Tx, image *types.Image) error {
	_, err := tx.NewInsert().
		Model(image).
		Exec(context.Background())
	return err
}
func UpdateImage(tx bun.Tx, image *types.Image) error {
	_, err := Bun.NewUpdate().
		Model(image).
		WherePK().
		Exec(context.Background())
	return err
}

func CreateNewAccount(account *types.Account) error {
	_, err := Bun.NewInsert().Model(account).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func GetAccountUserByID(id uuid.UUID) (types.Account, error) {
	var account types.Account
	err := Bun.NewSelect().Model(&account).Where("user_id = ?", id).Scan(context.Background())
	return account, err
}

func UpdateAccount(account *types.Account) error {
	_, err := Bun.NewUpdate().Model(account).WherePK().Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
