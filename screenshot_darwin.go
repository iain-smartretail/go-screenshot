package screenshot

import (
	// #cgo LDFLAGS: -framework CoreGraphics
	// #cgo LDFLAGS: -framework CoreFoundation
	// #include <CoreGraphics/CoreGraphics.h>
	// #include <CoreFoundation/CoreFoundation.h>
	"C"
	"image"
)

func ScreenRect() (image.Rectangle, error) {
	displayID := C.CGMainDisplayID()
	width := int(C.CGDisplayPixelsWide(displayID))
	height := int(C.CGDisplayPixelsHigh(displayID))
	return image.Rect(0, 0, width, height), nil
}

func CaptureScreen() (*image.RGBA, error) {
	rect, err := ScreenRect()
	if err != nil {
		return nil, err
	}
	return CaptureRect(rect)
}

func CaptureRect(rect image.Rectangle) (*image.RGBA, error) {
	width := 1024
	length := 0

	imageBytes := make([]byte, length)

	img := &image.RGBA{Pix: imageBytes, Stride: 4 * width, Rect: rect}
	return img, nil
}
