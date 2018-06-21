package model

import "log"

type ArticleModel struct {
	Model
}

type Article struct {
	ArticleId int
	Title     string
	Content string
}

func (a ArticleModel) FetchAll() []Article {
	a.InitSlave()
	stmt, err := a.Slave.Prepare("SELECT article_id, title FROM article ORDER BY update_at DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, _ := stmt.Query()
	var articles []Article
	for rows.Next() {
		article := Article{}
		rows.Scan(&article.ArticleId, &article.Title)
		articles = append(articles, article)
	}

	return articles
}

func (a ArticleModel) FetchOneByArticleId() Article {
	return Article{}
}
