package model

import "log"

type TagModel struct {
	Model
}

type Tag struct {
	Name string
	Id   int
}

func (t TagModel) FetchTagsByArticleId(articleId int) []Tag {
	t.InitSlave()
	stmt, err := t.Slave.Prepare("SELECT t.name,t.tag_id FROM article_tag AS at LEFT JOIN tag AS t ON at.tag_id=t.tag_id WHERE" +
		" at.article_id=?")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query(articleId)
	if err != nil {
		log.Fatal(err)
	}

	var tags []Tag
	for rows.Next() {
		tag := Tag{}
		rows.Scan(&tag.Name, &tag.Id)
		tags = append(tags, tag)
	}
	return tags
}
