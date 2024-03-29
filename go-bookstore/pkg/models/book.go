package models

import (
	"errors"
	"strconv"

	"github.com/deadking/go-bookstore/pkg/types"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Book struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Publication string `json:"publication,omitempty"`
	AuthorID    uint   `json:"author_id,omitempty"`
	Author      Author `json:"author" gorm:"foreignKey:AuthorID ;references:ID;omitempty"`
}

type IBookCRUD interface {
	Create(book Book) (*types.ResponseBook, error)
	Update(updateBook Book) (*types.ResponseBook, error)
	Delete(ID int) error
	Get(bookID, authorID int) []Book
}

func authorIDValidate(a uint) validation.RuleFunc {
	return func(value interface{}) error {
		if value.(uint) == 0 {
			return errors.New("enter valid author id(Numerical)")
		}
		return nil
	}
}

func bookNameValidate(a string) validation.RuleFunc {
	return func(value interface{}) error {
		name := value.(string)
		if _, err := strconv.Atoi(name); err == nil || len(name) < 1 || len(name) > 150 {
			return errors.New("enter valid book name(English) of 1 to 150 characters")
		}
		return nil
	}
}

func pubNameValidate(a string) validation.RuleFunc {
	return func(value interface{}) error {
		name := value.(string)
		if _, err := strconv.Atoi(name); err == nil || len(name) < 4 || len(name) > 150 {
			return errors.New("enter valid publication(English) of 4 to 150 characters")
		}
		return nil
	}
}

func (b Book) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Name, validation.Length(1, 150),
			validation.By(bookNameValidate(b.Name)),
			validation.Required),
		validation.Field(&b.Publication,
			validation.Length(4, 150),
			validation.By(pubNameValidate(b.Publication)),
			validation.Required),
		validation.Field(&b.AuthorID,
			validation.By(authorIDValidate(b.AuthorID)),
			validation.Required),
	)
}
