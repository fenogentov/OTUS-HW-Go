package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	dirEnv := "./testdata/env"
	errDirEnv := "./errdata/env"

	t.Run("env read", func(t *testing.T) {
		env, err := ReadDir(dirEnv)
		require.Equal(t, err, nil, "ReadDir return error: %s", err)
		require.Equal(t, env["BAR"].Value, "bar", "error read BAR")
		require.Equal(t, env["EMPTY"].Value, "", "error read EMPTY")
		require.Equal(t, env["FOO"].Value, "   foo\nwith new line", "error read FOO")
		require.Equal(t, env["HELLO"].Value, "\"hello\"", "error read HELLO")
		require.Equal(t, env["UNSET"].Value, "", "error read UNSET")
	})

	t.Run("error directory env", func(t *testing.T) {
		_, err := ReadDir(errDirEnv)
		require.EqualErrorf(t, err, "open "+errDirEnv+": no such file or directory", "actual %s:", err)
	})
}
