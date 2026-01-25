package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
)

func TestMain(t *testing.T) {
	require.NoError(t, fx.ValidateApp(appOpts...))
}
