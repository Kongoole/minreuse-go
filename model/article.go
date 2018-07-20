package model

import (
	"errors"
	"log"
	"strconv"
)

const PAGE_SIZE = 7

const unpublished = 1
const published = 2

type ArticleModel struct {
	Model
	StatusPublished   int
	StatusUnpublished int
}

type Article struct {
	ArticleId int
	Title     string
	Content   string
	CreateAt  string
	UpdateAt  string
}

var articleModel *ArticleModel

// ArticleModelInstance creates an ArticleModel instance
func ArticleModelInstance() *ArticleModel {
	return &ArticleModel{StatusPublished: published, StatusUnpublished: unpublished}
}

// FetchAll fetches all articles
func (a *ArticleModel) FetchAll() []Article {
	a.InitSlave()
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

// FetchAllPublished fetches all published articles
func (a *ArticleModel) FetchAllPublished() []Article {
	a.InitSlave()
	stmt, err := a.Slave.Prepare("SELECT article_id, title, create_at, update_at FROM article WHERE status=" +
		strconv.Itoa(a.StatusPublished) + " ORDER BY update_at DESC")
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

func (a *ArticleModel) FetchWithPagination(offset int, statuses ...int) []Article {
	a.InitSlave()
	sql := "SELECT article_id, title FROM article"
	if len(statuses) > 0 {
		statusStr := "("
		for _, status := range statuses {
			statusStr += strconv.Itoa(status) + ","
		}
		// remove last ","
		statusStr = statusStr[:len(statusStr)-1] + ")"
		sql += " WHERE status IN " + statusStr
	}
	sql = sql + " ORDER BY update_at DESC LIMIT " + strconv.Itoa(offset*PAGE_SIZE) + ", " + strconv.Itoa(PAGE_SIZE)
	stmt, err := a.Slave.Prepare(sql)
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
func (a *ArticleModel) FetchArticleAmount(statuses ...int) int {
	a.InitSlave()
	sql := "SELECT COUNT(article_id) FROM article"
	if len(statuses) > 0 {
		statusStr := "("
		for _, status := range statuses {
			statusStr += strconv.Itoa(status) + ","
		}
		statusStr = statusStr[:len(statusStr)-1] + ")"
		sql += " WHERE status IN " + statusStr
	}
	stmt, err := a.Slave.Prepare(sql)
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

func (a *ArticleModel) FetchOneByArticleId(articleId int) Article {
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
	defer rows.Close()

	var article Article
	for rows.Next() {
		rows.Scan(&article.ArticleId, &article.Title, &article.Content)
		break
	}

	return article
}

func (a *ArticleModel) FetchTagArticlesByTagId(tagId int) []Article {
	a.InitSlave()
	stmt, err := a.Slave.Prepare("SELECT a.article_id, a.title FROM article AS a INNER JOIN article_tag AS at" +
		" ON a.article_id=at.article_id WHERE at.tag_id=? AND a.status=" + strconv.Itoa(a.StatusPublished))
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

func (a *ArticleModel) FetchArticlesByKeyWords(keywords string) []Article {
	a.InitSlave()
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

func (a *ArticleModel) AddArticle(title, content string, author_id, status int) (int, error) {
	if title == "" {
		return 0, errors.New("title cannot be empty")
	}
	if content == "" {
		return 0, errors.New("content cannot be empty")
	}

	a.InitMaster()
	stmt, err := a.Master.Prepare("INSERT into article(`title`, `content`, `author_id`, `status`) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal("add article: failed to prepare, " + err.Error())
	}
	defer stmt.Close()
	res, err := stmt.Exec(title, content, author_id, status)
	if err != nil {
		log.Fatal("add article: failed to exec " + err.Error())
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal("add article: failed to get last insert id, " + err.Error())
	}
	return int(lastId), nil
}
