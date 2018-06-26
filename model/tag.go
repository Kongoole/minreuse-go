package model

import "log"

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
func (t TagModel) FetchTagsWithArticlesNum() []Tag {
	t.InitSlave()
	stmt, err := t.Slave.Prepare("SELECT tag.name, tag.tag_id, count(1) AS num FROM tag JOIN article_tag AS at ON " +
		"tag.tag_id=at.tag_id group by at.tag_id;")
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
