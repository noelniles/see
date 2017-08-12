package histogram

import (
    "image"
)

// ImageHistogram will take in an image and compute the
// image histogram. For color images this will produce 3
// distinct histograms. For gray images this will still
// produce 3 histograms but they will all be the same.
func ImageHistogram(img *image.Image) (*histogram, *histogram, *histogram, error) {
    width  := (*img).Bounds().Max.X
    height := (*img).Bounds().Max.Y

    rh, e := NewHist(256)
    gh, e := NewHist(256)
    bh, e := NewHist(256)

    for j := 0; j < height; j++ {
        for i := 0; i < width; i++ {
            r, g, b, _ := (*img).At(i, j).RGBA() 
            rh.Increment(float32(r>>8))
            gh.Increment(float32(g>>8))
            bh.Increment(float32(b>>8))
        }
    }
    return rh, gh, bh, e
}
