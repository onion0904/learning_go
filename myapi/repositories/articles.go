package repository

import (
    "database/sql"
    "fmt"
    "log"
	"github.com/yourname/reponame/models"
)
const (
	articleNumPerPage = 5
)

// 新規投稿をデータベースに insert する関数
// -> データベースに保存した記事内容と、発生したエラーを返り値にする
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	insert into articles (title, contents, username, nice, created_at) values
	(?, ?, ?, 0, now());
	`
	var newArticle models.Article
	newArticle.Title = article.Title
	newArticle.Contents = article.Contents
	netwArticle.UserName = article.UserName
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err!= nil {
		fmt.Println(err)
        return newArticle, err
    }
	id, _ := result.LastInsertId()
	newArticle.ID = int(id)
	return newArticle, nil
}
// 変数 page で指定されたページに表示する投稿一覧をデータベースから取得する関数
// -> 取得した記事データと、発生したエラーを返り値にする
func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, nice
		from articles
		limit ? offset ?;
	`
	rows, err := db.Query(sqlStr, articleNumPerPage, ((page - 1) * articleNumPerPage))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	articleArray := make([]models.Article,0)


// (問 2) 指定された記事データをデータベースから取得して、
// それを models.Article 構造体のスライス []models.Article に詰めて返す処理
	return articleArray, nil
}
// 投稿 ID を指定して、記事データを取得する関数
// -> 取得した記事データと、発生したエラーを返り値にする
func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
		`
// (問 3) 指定 ID の記事データをデータベースから取得して、それを models.Article 構造体の形で返す処理
	return article, nil
}
// いいねの数を update する関数
// -> 発生したエラーを返り値にする
func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
	`
	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
// (問 4) 指定された ID の記事のいいね数を+1 するようにデータベースの中身を更新する処理
	return nil
}
