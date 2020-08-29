package domain

import (
	"errors"
	"regexp"
	"time"
)

const (
	titleMinLength = 0
	bodyMinLength  = 0
	uuidRegex      = "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"
)

type Story struct {
	ID        string
	Title     string
	Body      string
	CreatedAt time.Time `db:"createdat"`
	UpdatedAt time.Time `db:"updatedat"`
}

func (s *Story) GetID() string {
	return s.ID
}

func (s *Story) GetTitle() string {
	return s.Title
}

func (s *Story) GetBody() string {
	return s.Body
}

func (s *Story) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s *Story) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

func NewVanillaStory(titleMaxLength, bodyMaxLength int, title, body string) (*Story, error) {
	if err := validate(titleMaxLength, bodyMaxLength, title, body); err != nil {
		return nil, err
	}

	return &Story{Title: title, Body: body}, nil
}

func NewStory(titleMaxLength, bodyMaxLength int, id, title, body string) (*Story, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}

	if err := validate(titleMaxLength, bodyMaxLength, title, body); err != nil {
		return nil, err
	}

	return &Story{
		ID:    id,
		Title: title,
		Body:  body,
	}, nil
}

func validateID(id string) error {
	if !regexp.MustCompile(uuidRegex).MatchString(id) {
		return errors.New("invalid id")
	}

	return nil
}

func validate(titleMaxLength, bodyMaxLength int, title, body string) error {
	if len(title) <= titleMinLength {
		return errors.New("title is empty")
	}

	if len(body) <= bodyMinLength {
		return errors.New("body is empty")
	}

	if len(title) > titleMaxLength {
		return errors.New("title length exceeded")
	}

	if len(body) > bodyMaxLength {
		return errors.New("body length exceeded")
	}

	return nil
}
