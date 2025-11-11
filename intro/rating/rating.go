package rating

import (
	"fmt"
	"time"
)

type Rating struct {
	ratingId      int
	averageRating float64
	rating        []RatingDetails
}

type RatingDetails struct {
	userId  string
	rating  int
	comment Comment
}

type Comment struct {
	date       string
	commentMsg string
}

func (RatingDetails RatingDetails) String() string {
	return fmt.Sprintf("User: %v, Rating: %v, Comment: %v", RatingDetails.userId, RatingDetails.rating, RatingDetails.comment)
}

func (Comment Comment) String() string {
	return fmt.Sprintf("Date: %v, Comment: %v", Comment.date, Comment.commentMsg)
}

func (r *Rating) AddRating(userId string, rating int, commentMsg string) error {

	if rating < 0 || rating > 5 {
		return fmt.Errorf("invalid rating")
	}

	r.rating = append(r.rating, RatingDetails{
		userId: userId,
		rating: rating,
		comment: Comment{
			date:       time.Now().Format("2006-01-02 15:04:05"),
			commentMsg: commentMsg,
		},
	})
	return nil
}

func (r *Rating) AverageRating() {
	total := 0
	for _, rd := range r.rating {
		total += rd.rating
	}
	r.averageRating = float64(total) / float64(len(r.rating))

}

func convertRatingToStars(rating int) string {

	stars := ""
	for i := 0; i < rating; i++ {
		stars += "*"
	}
	return stars
}
