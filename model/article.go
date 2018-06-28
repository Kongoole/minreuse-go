package model

import (
	"log"
	"strconv"
)

const PAGE_SIZE = 7

type ArticleModel struct {
	Model
}

type Article struct {
	ArticleId int
	Title     string
	Content   string
	CreateAt  string
	UpdateAt  string
}

func NewArticleModel() ArticleModel {
	return ArticleModel{}
}

func (a ArticleModel) FetchAll() []Article {
	(&a).InitSlave()
	stmt, err := a.Slave.Prepare("SELECT article_id, title, create_at, update_at FROM article ORDER BY update_at DESC")
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
		rows.Scan(&article.ArticleId, &article.Title, &article.CreateAt, &article.UpdateAt)
		articles = append(articles, article)
	}

	return articles
}

func (a ArticleModel) FetchWithPagination(offset int) []Article {
	(&a).InitSlave()
	stmt, err := a.Slave.Prepare("SELECT article_id, title FROM article ORDER BY update_at DESC LIMIT " +
		strconv.Itoa(offset*PAGE_SIZE) + ", " + strconv.Itoa(PAGE_SIZE))
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

// FetchArticleAmount fetch all article amount
func (a ArticleModel) FetchArticleAmount() int {
	(&a).InitSlave()
	stmt, err := a.Slave.Prepare("SELECT COUNT(article_id) FROM article")
	if err != nil {
		log.Fatal("fail to fetch total article amount")
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("fail to fetch total article amount")
	}
	defer rows.Close()

	var total int
	for rows.Next() {
		rows.Scan(&total)
	}
	return total
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
	stmt, err := a.Slave.Prepare("SELECT article_id, title, content FROM article WHERE title LIKE concat('%', ?, '%')")
	if err != nil {
		log.Fatal("failed to search article, err: " + err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(keywords)
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
