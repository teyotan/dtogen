package schema

import (
	"github.com/google/uuid"
)

type Company struct {
	ID        uuid.UUID
	Email     string
	Name      *string
	UserCount uint8
}
