package schema

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID
	Email string
	Name  *string
}

type UserLogin struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Provider  string
	CreatedAt time.Time
}
