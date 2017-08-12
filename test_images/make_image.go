package test_images

import (
    "image"
    "image/color"
    //"image/gif"
    //"image/jpeg"
    //"image/png"
)

// WhiteImage will return an image that is completely white.
func WhiteImage() (*image.RGBA) {
    img := image.NewRGBA(image.Rect(0, 0, 256, 256))

    for j := 0; j < img.Bounds().Max.Y; j++ {
        for i := 0; i < img.Bounds().Max.X; i++ {
            img.Set(i, j, color.White)
        }
    }
    return img
}

// BlackImage will return an image that is completely black.
func BlackImage() (*image.Uniform) {
    img := image.NewRGBA(image.Rect(0, 0, 256, 256))

    for j := 0; j < img.Bounds().Max.Y; j++ {
        for i := 0; i < img.Bounds().Max.X; i++ {
            img.Set(i, j, color.Black)
        }
    }
    return img
}

// FlatImage will return an image that has a flat histogram.
func FlatImage() {

}
