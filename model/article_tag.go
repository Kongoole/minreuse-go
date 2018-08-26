package model

type ArticleTagModel struct {
	Model
}

// AddArticleTags performs batch insert article tags
func (atm *ArticleTagModel) AddArticleTags(articleID int, tagIDs []int) (int64, error) {
	atm.InitMaster()
	sql := "INSERT INTO article_tag(`article_id`, `tag_id`) VALUES "
	var data []interface{}
	// make batch insert sql
	for _, tagID := range tagIDs {
		sql += "(?, ?), "
		data = append(data, articleID, tagID)
	}
	sql = sql[:len(sql)-2]

	stmt, err := atm.Master.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(data...)
	if err != nil {
		return 0, nil
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}
