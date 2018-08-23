package model

import (
	"errors"
	"log"
	"strconv"
)

type TagModel struct {
	Model
}

type Tag struct {
	Name     string
	Id       int
	Articles int // how many articles belong to this tag
}

func NewTagModel() TagModel {
	return TagModel{}
}

// FetchAll fetches all tags
func (t TagModel) FetchAll() []Tag {
	t.InitSlave()
	stmt, err := t.Slave.Prepare("SELECT tag_id, name FROM tag")
	if err != nil {
		log.Fatal("fail to fetch tags:" + err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("fail to fetch tags:" + err.Error())
	}
	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		tag := Tag{}
		rows.Scan(&tag.Id, &tag.Name)
		tags = append(tags, tag)
	}
	return tags
}

// fetch article tags
func (t TagModel) FetchTagsByArticleId(articleId int) []Tag {
	t.InitSlave()
	stmt, err := t.Slave.Prepare("SELECT t.name,t.tag_id FROM article_tag AS at LEFT JOIN tag AS t ON at.tag_id=t.tag_id WHERE" +
		" at.article_id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(articleId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		tag := Tag{}
		rows.Scan(&tag.Name, &tag.Id)
		tags = append(tags, tag)
	}
	return tags
}

// fetch tags with article amount
func (t TagModel) FetchTagsWithArticlesNum(status int) []Tag {
	t.InitSlave()
	stmt, err := t.Slave.Prepare("SELECT tag.name, tag.tag_id, count(1) AS num FROM tag JOIN article_tag AS at ON " +
		"tag.tag_id=at.tag_id JOIN article on at.article_id=article.article_id WHERE article.status=" +
		strconv.Itoa(status) + " group by at.tag_id")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		tag := Tag{}
		rows.Scan(&tag.Name, &tag.Id, &tag.Articles)
		tags = append(tags, tag)
	}

	return tags
}

func (t TagModel) AddArticleTags(articleId int, tagIds []int) error {
	if len(tagIds) == 0 {
		return errors.New("Invalid tag ids")
	}
	t.InitMaster()
	var vals []int
	sql := "INSERT INTO article_tag(`article_id`, `tag_id`) VALUES "
	for _, tagId := range tagIds {
		sql += "(?,?),"
		vals = append(vals, tagId)
	}
	stmt, err := t.Master.Prepare()
	if err != nil {
		log.Println("fail to insert into article_tag")
		return err
	}
	defer stmt.Close()

	stmt.Exec()
}
