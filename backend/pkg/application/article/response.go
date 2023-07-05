package article

import (
	"github.com/tufusa/article-api/backend/pkg/domain/model"
	"github.com/tufusa/article-api/backend/pkg/utils/restime"
)

type ArticleResponse struct {
	ID        uint                 `json:"id"`
	CreatedAt restime.ResourceTime `json:"created_at"`
	UpdatedAt restime.ResourceTime `json:"updated_at"`
	Title     string               `json:"title"`
	Text      string               `json:"text"`
}

func DomainToResponse(in model.Article) ArticleResponse {
	return ArticleResponse{
		ID:        in.ID,
		CreatedAt: restime.NewResourceTime(in.CreatedAt),
		UpdatedAt: restime.NewResourceTime(in.UpdatedAt),
		Title:     in.GetTitle(),
		Text:      in.GetText(),
	}
}
