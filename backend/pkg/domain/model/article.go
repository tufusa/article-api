package model

import (
	"unicode/utf8"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	title string `gorm:"not null"`
	text  string `gorm:"type:text;not null"`
}

func (a *Article) GetTitle() string {
	return a.title
}

func (a *Article) GetText() string {
	return a.text
}

type Articles []Article

type ArticleTitleLengthError struct {
	message string
}

func (e ArticleTitleLengthError) Error() string {
	return e.message
}

type ArticleTextLengthError struct {
	message string
}

func (e ArticleTextLengthError) Error() string {
	return e.message
}

func NewArticle(title string, text string) (*Article, error) {
	if titleLen := utf8.RuneCountInString(title); titleLen > 256 || titleLen <= 0 {
		return nil, ArticleTitleLengthError{
			message: "article title length outside of range [1, 256]",
		}
	}

	if textLen := utf8.RuneCountInString(text); textLen > 1<<15 || textLen <= 0 {
		return nil, ArticleTextLengthError{
			message: "article text length outside of range [1, 32768]",
		}
	}

	return &Article{
		title: title,
		text:  text,
	}, nil
}
