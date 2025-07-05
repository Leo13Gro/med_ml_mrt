package kt

import (
	"github.com/google/uuid"
)

type CreateKTArg struct {
	Author      uuid.UUID
	DeviceID    int
	Description *string
}
