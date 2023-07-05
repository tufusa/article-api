package article

import "github.com/tufusa/article-api/backend/pkg/domain/model"

type ArticleRequest struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func RequestToDomain(in ArticleRequest) (*model.Article, error) {
	return model.NewArticle(in.Title, in.Text)
}
