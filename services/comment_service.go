package services

import (
	"github.com/ikiat/work/models"
	"github.com/ikiat/work/repositories"
)

func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, nil
	}
	defer db.Close()

	newComment, err := repositories.InsertComment(db, comment)
	if  err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}