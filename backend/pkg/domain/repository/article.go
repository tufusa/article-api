package repository

import (
	"context"

	"github.com/tufusa/article-api/backend/pkg/domain/model"
)

type IArticleRepository interface {
	FindAllArticles(ctx context.Context) (model.Articles, error)
	FindArticleByID(ctx context.Context, id uint) (*model.Article, error)
}
