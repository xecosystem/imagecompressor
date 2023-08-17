package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"

	"github.com/nfnt/resize"
)

func fetchImage(url string) ([]byte, error) {
	// Fetch the image from the provided URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch image: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response contains a valid image
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("invalid image format: %w", err)
	}

	// Compress the image (without resizing)
	compressedImg := resize.Resize(uint(img.Bounds().Dx()), uint(img.Bounds().Dy()), img, resize.Lanczos3)

	// Encode the compressed image to JPEG format
	var compressedImageData bytes.Buffer
	err = jpeg.Encode(&compressedImageData, compressedImg, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to compress image: %w", err)
	}

	return compressedImageData.Bytes(), nil
}
