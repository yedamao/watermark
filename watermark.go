package main

import (
	"strings"

	"gopkg.in/gographics/imagick.v2/imagick"
)

func addComment(blob []byte, info string) ([]byte, string, error) {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	if err := mw.ReadImageBlob(blob); err != nil {
		return nil, "", err
	}

	if err := mw.CommentImage(info); err != nil {
		return nil, "", err
	}

	format := strings.ToLower(mw.GetImageFormat())
	if format == "png" {
		mw.SetOption("png:include-chunk", "all")
	}

	return mw.GetImageBlob(), format, nil
}
