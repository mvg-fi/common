package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShortenUUID(t *testing.T) {
	r := require.New(t)

	uuid := "7e8e4e4e-1e4e-2e8e-3e4e-4e4e8e4e8e4e"
	fmt.Println(uuid)
	fmt.Println(ShortenUUID(uuid))
	fmt.Println(RestoreUUID(ShortenUUID(uuid)))

	r.Equal(uuid, RestoreUUID(ShortenUUID(uuid)))
}
