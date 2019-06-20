package goheif

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"io"
	"io/ioutil"

	"github.com/hmage/goheif/heif"
)

func ExtractExif(ra io.ReaderAt) ([]byte, error) {
	hf := heif.Open(ra)
	return hf.EXIF()
}

func DecodeConfig(r io.Reader) (image.Config, error) {
	var config image.Config

	ra, err := asReaderAt(r)
	if err != nil {
		return config, err
	}

	hf := heif.Open(ra)

	it, err := hf.PrimaryItem()
	if err != nil {
		return config, err
	}

	width, height, ok := it.SpatialExtents()
	if !ok {
		return config, fmt.Errorf("No dimension")
	}

	config = image.Config{
		ColorModel: color.YCbCrModel,
		Width:      width,
		Height:     height,
	}
	return config, nil
}

func asReaderAt(r io.Reader) (io.ReaderAt, error) {
	if ra, ok := r.(io.ReaderAt); ok {
		return ra, nil
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil
}
