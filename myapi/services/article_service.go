package services

import (
	_ "fmt"
	"errors"
	"database/sql"
	"github.com/yourname/reponame/apperrors"
    "github.com/yourname/reponame/models"
    "github.com/yourname/reponame/repositories"
)

// ArticleDetailHandler で使うことを想定したサービス
// 指定 ID の記事情報を返却
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {

	// 1. repositories 層の関数 SelectArticleDetail で記事の詳細を取得
	article, err1 := repositories.SelectArticleDetail(s.db, articleID)
	if err1 != nil {
		if errors.Is(err1, sql.ErrNoRows) {
			err1 = apperrors.NAData.Wrap(err1, "no data")
			return models.Article{}, err1
		}
		err1 = apperrors.GetDataFailed.Wrap(err1, "fail to get data")
			
		return models.Article{}, err1
	}
	// 2. repositories 層の関数 SelectCommentList でコメント一覧を取得
	commentList, err2 := repositories.SelectCommentList(s.db, articleID)
	if err2 != nil {
		err2 = apperrors.GetDataFailed.Wrap(err2, "fail to get data")
		return models.Article{}, err2
	}

	// 3. 2 で得たコメント一覧を、1 で得た Article 構造体に紐付ける
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article,error){

	article , err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}
	return article, nil
}

// ArticleListHandler で使うことを想定したサービス
// 指定 page の記事一覧を返却
func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {


	ArticleList , err := repositories.SelectArticleList(s.db, page)
	if err!= nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
        return nil, err
    }

	if len(ArticleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return ArticleList, nil
}


// PostNiceHandler で使うことを想定したサービス
// 指定 ID の記事のいいね数を+1 して、結果を返却
func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {

	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err!= nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice count")
			
        return models.Article{}, err
    }

	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}