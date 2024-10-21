package lor

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFunctionIf(t *testing.T) {
	t.Run("function if true", func(t *testing.T) {
		platform := "line"
		res := FunctionIf(
			platform == "line",
			func() string {
				return "groupId"
			},
			func() string {
				return ""
			},
		)

		require.True(t, res == "groupId")
	})

	t.Run("function if false", func(t *testing.T) {
		platform := "line"
		res := FunctionIf(
			platform != "line",
			func() string {
				return "groupId"
			},
			func() string {
				return ""
			},
		)

		require.True(t, res == "")
	})
}

func TestFunctionDo(t *testing.T) {
	t.Run("function do true", func(t *testing.T) {
		platform := "line"
		res := FunctionDo(
			platform == "line",
			func() string {
				return "groupId"
			},
		)

		require.True(t, res == "groupId")
	})

	t.Run("function do false", func(t *testing.T) {
		platform := "line"
		res := FunctionDo(
			platform != "line",
			func() error {
				return fmt.Errorf("error")
			},
		)

		require.True(t, res == nil)
	})
}
