package histogram

import (
    "fmt"
)

// NewHistPdf returns a new histogram_pdf given a histogram.
func NewHistPdf(h *histogram) (*histogram_pdf) {
    p := histogram_pdf{}
    p.n = h.n

    // Make sure histogram bins are non-negative.
    var i uint32
    for i = 0; i < h.n; i++ {
        if h.bins[i] < 0 {
            fmt.Errorf("histogram bins must be non-negative to compute pdf")
        }
    }

    // Init the ranges and sum arrays.
    p.ranges = make([]float32, p.n + 1)
    p.sum = make([]float32, p.n + 1)

    // Copy the ranges.
    for i = 0; i <= h.n; i++ {
        p.ranges[i] = h.ranges[i]
    }

    var mean, sum float32
    mean = 0
    sum = 0

    // Compute the mean
    for i = 0; i < h.n; i++ {
        mean += (h.bins[i] - mean) / float32(i + 1)
    }
    p.sum[0] = 0

    for i = 0; i < h.n; i++ {
        sum += (h.bins[i] / mean) / float32(h.n)
        p.sum[i+1] = sum
    }
    return &p
}

func (p *histogram_pdf) Sample(r float32) float32 {
    // Wrap the exclusive top of the bin down to the inclusive
    // bottom of the bin. Since this is a single point it 
    // should not affect the distribution.
    if r == 1.0 {
        r = 0.0
    }
    status, i := find(p.n, p.sum, r)

    var x float32
    x = 0
    if status {
        fmt.Errorf("cannot find r in the cumulative pdf")
    } else {
        delta := ((r - p.sum[i]) / (p.sum[i+1] - p.sum[i]))
        x = p.ranges[i] + delta * (p.ranges[i+1] - p.ranges[i])
    }
    return x
}
