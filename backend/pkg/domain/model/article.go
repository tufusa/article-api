package model

import (
	"unicode/utf8"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title string `json:"title" gorm:"not null"`
	Text  string `json:"text" gorm:"type:text;not null"`
}

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
		return nil, ArticleTitleLengthError{}
	}

	if textLen := utf8.RuneCountInString(text); textLen > 1<<15 || textLen <= 0 {
		return nil, ArticleTextLengthError{}
	}

	return &Article{
		Title: title,
		Text:  text,
	}, nil
}
