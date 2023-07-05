package inmemory

import (
	"context"

	"github.com/tufusa/article-api/backend/pkg/domain/model"
	"github.com/tufusa/article-api/backend/pkg/domain/repository"
	"gorm.io/gorm"
)

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) repository.IArticleRepository {
	return &articleRepository{
		db: db,
	}
}

func (repo *articleRepository) FindAllArticles(ctx context.Context) (model.Articles, error) {
	articles := model.Articles{}
	err := repo.db.Find(&articles).Error
	return articles, err
}

func (repo *articleRepository) FindArticleByID(ctx context.Context, id uint) (*model.Article, error) {
	article := model.Article{}
	err := repo.db.Where("ID = ?", id).First(&article).Error
	return &article, err
}
