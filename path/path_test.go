package path

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRootDir(t *testing.T) {
	res := GetRootDir()
	base := filepath.Base(res)
	require.Equal(t, base, "mypackage")
}
