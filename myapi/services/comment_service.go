package services

import (
	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
)

// PostCommentHandler で使用することを想定したサービス
// 引数の情報をもとに新しいコメントを作り、結果を返却
func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err!= nil {
        return models.Comment{}, err
    }
	defer db.Close()

	comment, err2 := repositories.InsertComment(db, comment)
    if err2!= nil {
        return models.Comment{}, err2
    }	

	return comment, nil
}