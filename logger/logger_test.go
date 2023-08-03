package logger

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	require := require.New(t)

	out := filterOutput("hello from mvg %d", time.Now().UnixNano())
	require.Contains(out, "mvg")

	err := SetFilter("bitcoin")
	require.Nil(err)
	out = filterOutput("hello from mvg %d", time.Now().UnixNano())
	require.NotContains(out, "mvg")
	out = filterOutput("Bitcoin from mvg %d", time.Now().UnixNano())
	require.NotContains(out, "mvg")
	out = filterOutput("bitcoin from mvg %d", time.Now().UnixNano())
	require.Contains(out, "mvg")

	err = SetFilter("(?i)bitcoin")
	require.Nil(err)
	out = filterOutput("hello from mvg %d", time.Now().UnixNano())
	require.NotContains(out, "mvg")
	out = filterOutput("Bitcoin from mvg %d", time.Now().UnixNano())
	require.Contains(out, "mvg")
	out = filterOutput("bitcoin from mvg %d", time.Now().UnixNano())
	require.Contains(out, "mvg")
	out = filterOutput("ethereum from bitcoin %d", time.Now().UnixNano())
	require.NotContains(out, "mvg")

	err = SetFilter("(?i)bitcoin|MVG")
	require.Nil(err)
	out = filterOutput("hello from mvg %d", time.Now().UnixNano())
	require.Contains(out, "mvg")
	out = filterOutput("Bitcoin from mvg %d", time.Now().UnixNano())
	require.Contains(out, "mvg")
	out = filterOutput("bitcoin from mvg %d", time.Now().UnixNano())
	require.Contains(out, "mvg")
	out = filterOutput("ethereum from mvg %d", time.Now().UnixNano())
	require.Contains(out, "mvg")
	out = filterOutput("ethereum or bitcoin %d", time.Now().UnixNano())
	require.NotContains(out, "mvg")

	level = 0
	filter = nil
}
