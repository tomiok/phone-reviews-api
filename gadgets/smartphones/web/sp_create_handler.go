package web

import (
	"github.com/tomiok/course-phones-review/gadgets/smartphones/gateway"
	"github.com/tomiok/course-phones-review/internal/database"
)

type CreateSmartphoneHandler struct {
	gateway.SmartphoneCreateGateway
}

func NewCreateSmartphoneHandler(client *database.MySqlClient) *CreateSmartphoneHandler {
	return &CreateSmartphoneHandler{gateway.NewSmartphoneCreateGateway(client)}
}
