package gateway

import (
	"github.com/tomiok/course-phones-review/internal/database"
	"github.com/tomiok/course-phones-review/reviews/models"
)

type ReviewGateway interface {
	AddReview(cmd *models.CreateReviewCMD) (string, error)
}

type ReviewGtw struct {
	ReviewStorage
}

func NewReviewGateway(mongoClient *database.Mongo) ReviewGateway {
	return &ReviewGtw{&ReviewStg{mongoClient}}
}