package domain

import "time"

type Task struct {
	ID          string    `json:"id" bson:"_id" validate:"required"`
	Title       string    `json:"title" bson:"_title" validate:"required"`
	Description string    `json:"description" bson:"_description" validate:"required"`
	DueDate     time.Time `json:"due_date" bson:"_duedate" validate:"required"`
	Status      string    `json:"status" bson:"_status" validate:"required"`
}

type User struct {
	ID       string `json:"id" bson:"_id" validate:"required"`
	Role     string `json:"role" bson:"_role" validate:"required"`
	Password string `json:"password" bson:"_password" validate:"required"`
	Token    string `json:"token" bson:"_token,omitempty"`
}
