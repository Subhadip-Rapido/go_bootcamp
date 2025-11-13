package rating

import (
	"fmt"
	"time"
)

type RatingDetails struct {
	userId         string
	rating         int
	commentDetails CommentDetails
}

type CommentDetails struct {
	date    string
	comment string
}

type Rating struct {
	productId     string
	averageRating float64
	ratingDetails []RatingDetails
}

func (rd RatingDetails) String() string {
	return fmt.Sprintf("User ID: %s, Rating: %d, Comment: %s, Date: %s", rd.userId, rd.rating, rd.commentDetails.comment, rd.commentDetails.date)
}

func (c CommentDetails) String() string {
	return fmt.Sprintf("Comment: %s, Date: %s", c.comment, c.date)
}

func (r *Rating) AddRating(userId string, rating int, comment string) error {
	if rating < 1 || rating > 5 {
		return fmt.Errorf("invalid rating: %d. Rating must be between 1 and 5", rating)
	}
	r.ratingDetails = append(r.ratingDetails, RatingDetails{
		userId: userId,
		rating: rating,
		commentDetails: CommentDetails{
			date:    time.Now().Format("2006-01-02 15:04:05"),
			comment: comment,
		}})

	r.averageRating = float64(totalRating(r.ratingDetails)) / float64(len(r.ratingDetails))
	return nil
}

func totalRating(rating []RatingDetails) int {
	total := 0
	for _, ra := range rating {
		total += ra.rating
	}
	return total
}
