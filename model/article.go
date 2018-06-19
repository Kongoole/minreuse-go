package model

import "fmt"

type ArticleModel struct {
	Model
}

type Article struct {
	ArticleId int
	Title     string
}

func (a ArticleModel) FetchAll() {
	a.InitMaster()
	stmt, err := a.Master.Prepare("SELECT article_id, title FROM article")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	rows, _ := stmt.Query()
	var articles []Article
	for rows.Next() {
		article := Article{}
		rows.Scan(&article.ArticleId, &article.Title)
		articles = append(articles, article)
	}

	fmt.Println(articles)
}
