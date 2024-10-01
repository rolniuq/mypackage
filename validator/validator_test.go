package validator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSample(t *testing.T) {
	errs := Sample()
	require.NotEmpty(t, errs)
}
