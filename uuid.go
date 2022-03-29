package u2utils

import (
	"github.com/satori/go.uuid"
)

func UUIDv4() string {
	u := uuid.NewV4()
	return u.String()
}
