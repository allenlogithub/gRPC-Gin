package databases

import (
	"fmt"

	proto "post-get-server/proto"
)

type (
	GetPersonalArticleRequest struct {
		UserId int64
	}

	GetArticleCommentRequest struct {
		ArticleId int64
		UserId    int64
	}
)

// func: get articles of an user
func GetPersonalArticle(r *GetPersonalArticleRequest) (*proto.GetPersonalArticleReply, error) {
	q := `
	SELECT register.name, article.id, article.content, articlecomment.content
	FROM article
		LEFT JOIN articlecomment
			ON article.id=articlecomment.article_id
		LEFT JOIN register
			ON articlecomment.user_id=register.id
	WHERE article.user_id=?	
	`
	rows, err := connMysql.Query(q, r.UserId)
	if err != nil {
		return nil, fmt.Errorf("GetPersonalArticle.Query: %v", err)
	}

	rp, existedArtsId := proto.GetPersonalArticleReply{}, make(map[int64]int8)
	for rows.Next() {
		var userName *string // the user of whom leave the comment
		var articleId int64
		var articleContent string
		var articleCommentContent *string
		if err := rows.Scan(&userName, &articleId, &articleContent, &articleCommentContent); err != nil {
			return nil, fmt.Errorf("GetPersonalArticle.Scan: %v", err)
		}
		pc, isEmptyCmt := proto.PersonalCommentItem{}, true
		if userName != nil { // if suerName == nil, then comment == nil
			pc.UserName = *userName
			pc.Comment = *articleCommentContent
			isEmptyCmt = false
		}
		pai := proto.PersonalArticleItem{}
		if _, ok := existedArtsId[articleId]; !ok {
			pai.ArticleId = articleId
			pai.Content = articleContent
			if !isEmptyCmt {
				pai.Items = []*proto.PersonalCommentItem{&pc}
			}
			rp.Items = append(rp.Items, &pai)
			existedArtsId[articleId] = 1
		} else {
			newItems := rp.Items[len(rp.Items)-1].Items
			if !isEmptyCmt {
				newItems = append(newItems, &pc)
			}
			rp.Items[len(rp.Items)-1].Items = newItems
		}
	}

	return &rp, nil
}

// not handling the empty comment situation since this API is called
// after a comment post action which guarantees the comment length of
// an aticle will > 0
func GetArticleComment(r *GetArticleCommentRequest) (*proto.GetArticleCommentReply, error) {
	// NEED TO INPLEMENTED FRIEND LIST CHECKING IN THE FUTURE IN ORDER TO GET THE COMMENT FROM FRIEND'S ARTICLE
	// THE SITUATION SO FAR IS NO LIMITATION
	q := `
	SELECT register.name, articlecomment.content
	FROM articlecomment
		LEFT JOIN register
			ON articlecomment.user_id=register.id
	WHERE articlecomment.article_id=?
	`
	rows, err := connMysql.Query(q, r.ArticleId)
	if err != nil {
		return nil, fmt.Errorf("GetArticleComment.Query: %v", err)
	}
	rp := proto.GetArticleCommentReply{}
	for rows.Next() {
		var userName string
		var commentContent string
		if err := rows.Scan(&userName, &commentContent); err != nil {
			return nil, fmt.Errorf("GetArticleComment.Scan: %v", err)
		}
		ac := proto.ArticleComment{}
		ac.UserName = userName
		ac.Comment = commentContent
		rp.Items = append(rp.Items, &ac)
	}

	return &rp, nil
}
