package web

import (
	"encoding/json"
	"github.com/tomiok/course-phones-review/internal/database"
	"github.com/tomiok/course-phones-review/reviews/gateway"
	"github.com/tomiok/course-phones-review/reviews/models"
	"net/http"
)

func (h *ReviewHandler) AddReviewHandler(w http.ResponseWriter, r *http.Request) {
	cmd := params(r)
	res, err := h.AddReview(cmd)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(res))
}

func params(r *http.Request) *models.CreateReviewCMD {
	var cmd models.CreateReviewCMD

	_ = json.NewDecoder(r.Body).Decode(&cmd)

	return &cmd
}

type ReviewHandler struct {
	gateway.ReviewGateway
}

func NewReviewHandler(mongo *database.Mongo) *ReviewHandler {
	return &ReviewHandler{gateway.NewReviewGateway(mongo)}
}