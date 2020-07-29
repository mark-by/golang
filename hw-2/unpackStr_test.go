package unpack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWithoutEscapeExamples(t *testing.T) {
	require.Equal(t, "aaaabccddddde", Unpack("a4bc2d5e"), "check a4bc2d5e")
	require.Equal(t, "abcd", Unpack("abcd"), "check abcd")
	require.Equal(t, "", Unpack("45"), "check 45")
}

func TestWithEscapeExamples(t *testing.T) {
	require.Equal(t, `qwe45`, Unpack(`qwe\4\5`), `check qwe\4\5`)
	require.Equal(t, `qwe44444`, Unpack(`qwe\45`), `check qwe\45`)
	require.Equal(t, `qwe\\\\\`, Unpack(`qwe\\5`), `check qwe\\5`)
}
