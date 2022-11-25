package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnserializeURL(t *testing.T) {
	require.Equal(t, "/", unserializeURL("").String())
	require.Equal(t, "https://mytest.org/path?key=v1", unserializeURL("https://mytest.org/path?key=v1").String())
}
