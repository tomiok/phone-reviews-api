package web

import (
	"encoding/json"
	"github.com/tomiok/course-phones-review/gadgets/smartphones/gateway"
	"github.com/tomiok/course-phones-review/gadgets/smartphones/models"
	"github.com/tomiok/course-phones-review/internal/database"
	"net/http"
)

func (h *CreateSmartphoneHandler) SaveSmartphoneHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cmd := parseRequest(r)
	res, err := h.Create(cmd)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "error in create smartphone"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}

type CreateSmartphoneHandler struct {
	gateway.SmartphoneCreateGateway
}

func NewCreateSmartphoneHandler(client *database.MySqlClient) *CreateSmartphoneHandler {
	return &CreateSmartphoneHandler{gateway.NewSmartphoneCreateGateway(client)}
}

func parseRequest(r *http.Request) *models.CreateSmartphoneCMD {
	body := r.Body
	defer body.Close()
	var cmd models.CreateSmartphoneCMD

	_ = json.NewDecoder(body).Decode(&cmd)

	return &cmd
}
