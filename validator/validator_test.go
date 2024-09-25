package validator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSample(t *testing.T) {
	err := Sample()
	require.Error(t, err)
	require.True(t, err.Error() == "name is empty")
}
