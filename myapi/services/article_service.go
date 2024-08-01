package services

import (
	_ "fmt"

    "github.com/yourname/reponame/models"
    "github.com/yourname/reponame/repositories"
)

// ArticleDetailHandler で使うことを想定したサービス
// 指定 ID の記事情報を返却
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {

	// 1. repositories 層の関数 SelectArticleDetail で記事の詳細を取得
	article, err1 := repositories.SelectArticleDetail(s.db, articleID)
	if err1 != nil {
		return models.Article{}, err1
	}
	// 2. repositories 層の関数 SelectCommentList でコメント一覧を取得
	commentList, err2 := repositories.SelectCommentList(s.db, articleID)
	if err2 != nil {
		return models.Article{}, err2
	}

	// 3. 2 で得たコメント一覧を、1 で得た Article 構造体に紐付ける
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article,error){

	article , err := repositories.InsertArticle(s.db, article)
	if err != nil {
		return models.Article{}, err
	}
	return article, nil
}

// ArticleListHandler で使うことを想定したサービス
// 指定 page の記事一覧を返却
func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {


	ArticleList , err := repositories.SelectArticleList(s.db, page)
	if err!= nil {
        return nil, err
    }
	
	return ArticleList, nil
}


// PostNiceHandler で使うことを想定したサービス
// 指定 ID の記事のいいね数を+1 して、結果を返却
func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {

	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err!= nil {
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