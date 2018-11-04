package util

import (
	"github.com/satori/go.uuid"
	"strings"
)

func UUID() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", 4)
}
