package model

type ArticleTag struct {
	*Model
	TagId     string `json:"tag_id"`
	ArticleId string `json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
