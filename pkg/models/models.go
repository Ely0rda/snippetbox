package models

import (
	"errors"
	"time"
)

var (ErrNoRecord = errors.New("modesl: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail = errors.New("models: duplicate email")
)
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type User struct{
	ID int
	Name string
	Email string
	HashedPassword []byte
	Created time.Time
}

