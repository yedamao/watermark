package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_uploadImage(t *testing.T) {
	os.Setenv("AWS_ACCESS_KEY_ID", "")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "")

	blob, err := ioutil.ReadFile("./testing/gopher.png")
	require.NoError(t, err)

	uploadFile := "gopher.png"
	err = uploadImage(uploadFile, blob)
	require.NoError(t, err)
}
