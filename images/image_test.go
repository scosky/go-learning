package images

import (
	"image"
	"image/png"
	"log"
	"os"
	"testing"
)

func TestBase64(t *testing.T) {
	const (
		width  = 300
		height = 200
	)

	file, err := os.Create("/Users/alice/Downloads/image/123.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := image.NewNRGBA(image.Rect(0, 0, width, height))

	err = png.Encode(file, m)
	if err != nil {
		log.Fatal(err)
	}
}
