package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Task struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" validate:"required,min=6,max=45"`
	Description string    `json:"description" validate:"required,min=5"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t *Task) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
