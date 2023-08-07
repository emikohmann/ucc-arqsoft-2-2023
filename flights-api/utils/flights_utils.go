package utils

import (
	"github.com/google/uuid"
)

func NewUUID() string {
	id := uuid.New()
	return id.String()
}
