package databases

import (
	"fmt"
)

type (
	PostArticleRequest struct {
		UserId     int64
		Content    string
		Visibility string
	}
)

// func: add a new article
func AddArticle(r *PostArticleRequest) (int64, error) {
	result, err1 := connMysql.Exec("INSERT INTO article (user_id, content, visibility) VALUES (?, ?, ?)", r.UserId, r.Content, r.Visibility)
	if err1 != nil {
		return 0, fmt.Errorf("AddRegister: %v", err1)
	}
	id, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, fmt.Errorf("AddRegister: %v", err2)
	}

	return id, nil
}
