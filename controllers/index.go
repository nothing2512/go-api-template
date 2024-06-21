package controllers

import (
	"main/models"

	"gorm.io/gorm"
)

type InDb struct {
	db *gorm.DB
}

func NewInstance() *InDb {
	return &InDb{
		db: models.GetSqlConnection(),
	}
}
