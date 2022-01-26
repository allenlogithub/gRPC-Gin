package databases

import (
	"fmt"
)

type (
	DelArticleRequest struct {
		UserId    int64
		ArticleId int64
	}
)

// func: delete an article
func DelArticle(r *DelArticleRequest) error {
	_, err := connMysql.Exec("DELETE FROM article WHERE id = ? AND user_id = ?", r.ArticleId, r.UserId)
	if err != nil {
		return fmt.Errorf("DelArticle: %v", err)
	}

	return nil
}
