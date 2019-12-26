package gateway

import (
	"github.com/tomiok/course-phones-review/gadgets/smartphones/models"
	"github.com/tomiok/course-phones-review/internal/database"
)

type SmartphoneCreateGateway interface {
	Create(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error)
}

type SmartphoneCreateGtw struct {
	SmartphoneStorageGateway
}

func NewSmartphoneCreateGateway(client *database.MySqlClient) SmartphoneCreateGateway {
	return &SmartphoneCreateGtw{&SmartphoneStorage{client}}
}
