package domain

import (
	"time"

	"github.com/google/uuid"
)

type KT struct {
	Id       uuid.UUID
	CreateAt time.Time
}
