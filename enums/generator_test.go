package enums

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateGoCodeFromJSON(t *testing.T) {
	err := GenerateGoCodeFromJSON("LINE_ADS_ENUM", "input.json", "out")
	require.NoError(t, err)
}
