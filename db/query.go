package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/v3ronez/ufantasyai/types"
)

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