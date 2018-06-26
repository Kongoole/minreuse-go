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
	(&a).InitSlave()
	stmt, err := a.Slave.Prepare("SELECT article_id, title FROM article ORDER BY update_at DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		article := Article{}
		rows.Scan(&article.ArticleId, &article.Title)
		articles = append(articles, article)
	}

	return articles
}

func (a ArticleModel) FetchOneByArticleId(articleId int) Article {
	(&a).InitSlave()

	stmt, err := a.Slave.Prepare("SELECT article_id, title, content FROM article WHERE article_id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(articleId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var article Article
	for rows.Next() {
		rows.Scan(&article.ArticleId, &article.Title, &article.Content)
		break
	}

	return article
}

func (a ArticleModel) FetchTagArticlesByTagId(tagId int) []Article {
	(&a).InitSlave()
	stmt, err := a.Slave.Prepare("SELECT a.article_id, a.title FROM article AS a INNER JOIN article_tag AS at" +
		" ON a.article_id=at.article_id WHERE at.tag_id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(tagId)
	defer rows.Close()
	var articles []Article
	for rows.Next() {
		article := Article{}
		rows.Scan(&article.ArticleId, &article.Title)
		articles = append(articles, article)
	}

	return articles
}

func (a ArticleModel) FetchArticlesByKeyWords(keywords string) []Article {
	(&a).InitSlave()
	stmt, err := a.Slave.Prepare("SELECT article_id, title, content FROM article WHERE title LIKE concat('%', ?, '%') OR content LIKE concat('%', ?, '%')")
	if err != nil {
		log.Fatal("failed to search article, err: " + err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(keywords, keywords)
	if err != nil {
		log.Fatal("failed to search article, err:" + err.Error())
	}
	defer rows.Close()
	var articles []Article
	for rows.Next() {
		article := Article{}
		rows.Scan(&article.ArticleId, &article.Title, &article.Content)
		articles = append(articles, article)
	}
	return articles
}
