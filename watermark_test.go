package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_embed(t *testing.T) {
	// read original image

	blob, err := ioutil.ReadFile("./testing/gopher.png")
	require.NoError(t, err)

	info := "It's a mark."
	processed, err := embed(blob, info)
	require.NoError(t, err)

	err = ioutil.WriteFile("./processed.png", processed, 0644)
	require.NoError(t, err)
}
