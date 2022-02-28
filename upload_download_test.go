package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_uploadImage(t *testing.T) {
	blob, err := ioutil.ReadFile("./testing/gopher.png")
	require.NoError(t, err)

	err = uploadImage(blob)
	require.NoError(t, err)
}
