package repository

import (
	"gorm.io/gorm"
	"goscaf/internal/model"
)

type AccountRepository interface {
	Get(a *model.Account) ([]model.Account, error)
	GetOne(a *model.Account) (model.Account, error)

	Insert(a *model.Account) (int, error)
	Update(a *model.Account) error
	Delete(a *model.Account) error

	InsertTx(a *model.Account, tx *gorm.DB) (int, error)
	UpdateTx(a *model.Account, tx *gorm.DB) error
	DeleteTx(a *model.Account, tx *gorm.DB) error
}