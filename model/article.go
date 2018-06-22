package model

import (
	"log"
)

type ArticleModel struct {
	Model
}

type Article struct {
	ArticleId int
	Title     string
	Content   string
}

func (a ArticleModel) FetchAll() []Article {
	a.InitSlave()
	stmt, err := a.Slave.Prepare("SELECT article_id, title FROM article ORDER BY update_at DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	var articles []Article
	for rows.Next() {
		article := Article{}
		rows.Scan(&article.ArticleId, &article.Title)
		articles = append(articles, article)
	}

	return articles
}

func (a ArticleModel) FetchOneByArticleId(articleId int) Article {
	a.InitSlave()

	stmt, err := a.Slave.Prepare("SELECT article_id, title, content FROM article WHERE article_id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(articleId)
	if err != nil {
		log.Fatal(err)
	}

	var article Article
	for rows.Next() {
		rows.Scan(&article.ArticleId, &article.Title, &article.Content)
		break
	}

	return article
}
