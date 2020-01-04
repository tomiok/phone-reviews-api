package gateway

import "github.com/tomiok/course-phones-review/reviews/models"

func (g *ReviewGtw) AddReview(cmd *models.CreateReviewCMD) (string, error) {
	return g.Add(cmd)
}
