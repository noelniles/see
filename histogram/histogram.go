package histogram

import (
    "fmt"
    "math"
    //"image"
)

type histogram struct {
    n uint32            // number of histogram bins
    ranges []float32    // the ranges of the bins stored in n+1-array
    bins []float32      // the counts for each bin stored in an n-array
}

type histogram_pdf struct {
    n uint32
    ranges []float32
    sum []float32
}

// Sum dumps the sum array to the terminal.
func (p *histogram_pdf) Sum() []float32 {
    return p.sum
}

// SumAt returns the sum at i.
func (p* histogram_pdf) SumAt(i uint32) float32 {
    if i > p.n || i < 0 {
        fmt.Errorf("index is out of range")
    } 
    return p.sum[i]
}

// NewHist creates a new histogram struct.
func NewHist(nbins uint32) (h *histogram, err error) {
    if nbins == 0 {
        fmt.Println("Histogram length n must be positive\n")
    }

    h1 := histogram{}
    h1.ranges = make([]float32, nbins + 1)
    h1.bins = make([]float32, nbins)
    h1.n = nbins

    h1.setRangesUniform(0, float32(nbins))

    if h1.ranges == nil || h1.bins == nil {
        err = fmt.Errorf("histogram.NewHist: failed to create to histogram")
    }
    return &h1, err
}

// Copy copies src to dest.
func (src *histogram) Copy() (dest *histogram, err error) {
    dest, err = NewHist(src.n)

    var i uint32
    for i = 0; i <= dest.n; i++ {
        dest.ranges[i] = src.ranges[i]
    }
    for i = 0; i < dest.n; i++ {
        dest.bins[i] = src.bins[i]    
    }
    return dest, err
}

// MakeUniform will make a histogram range uniform on the range
// xmin to xmax.
func (h *histogram) makeUniform(ranges []float32, n uint32,
    xmin, xmax float32) {
    
    var i uint32
    var f1, f2 float32

    for i = 0; i <= n; i++ {
        f1 = float32(n-i) / float32(n)
        f2 = float32(i) / float32(n)
        ranges[i] = f1 * xmin + f2 * xmax
    }
}

// SetRangesUniform will set the ranges to a uniformly spaced
// bins between xmin and xmax.
func (h *histogram) setRangesUniform(xmin, xmax float32) {
    n := h.n

    if (xmin >= xmax) {
        fmt.Println("xmin must be greater than xmax")
    }
    // Initialize ranges.
    h.makeUniform(h.ranges, n, xmin, xmax)

    // Clear the bins
    var i uint32
    for i = 0; i < n; i ++ {
        h.bins[i] = 0
    }
}

// SetRanges will set the ranges of the histogram and also
// clear the histogram bins.
func (h *histogram) SetRanges(ranges []float32, size uint32) {
    n := h.n

    if size != n+1 {
        fmt.Println("size of range must match size of histogram")
    }
    var i uint32
    for i = 0; i <= n; i++ {
        h.ranges[i] = ranges[i]
    }
    // Clear the histograms contents
    for i = 0; i < n; i++ {
        h.bins[i] = 0
    }
    h.ranges = ranges
}

// AllocateRanges will return a new histogram with the ranges
// allocated using the given ranges array.
func AllocateRanges(n uint32, ranges []float32) (h *histogram, err error) {
    var i uint32
    i = 0

    // Check the arguments
    if n == i {
        fmt.Println("histogram must be positive length")
    }
    // Check ranges
    for i = 0; i < n; i++ {
        if ranges[i] >= ranges[i+1] {
            fmt.Println("histogram ranges must be in increasing order")
        }
    }

    H, err := NewHist(n)
    if err != nil {
        fmt.Println("Problem creating new histogram")
        fmt.Errorf("%f", err)
    }

    // Initialize the ranges.
    for i = 0; i <= n; i++ {
        H.ranges[i] = ranges[i]
    }
    // Clear the contents.
    for i = 0; i < n; i++ {
        H.bins[i] = 0
    }
    return H, err
}

// Reset sets all of the bins to zero.
func (h *histogram) Reset() {
    var i uint32
    for i = 0; i < h.n; i++ {
        h.bins[i] = 0 
    }
}

// Increment increments a bin by one.
func (h *histogram) Increment(x float32) {
    h.Accumulate(x, 1.0) 
}

// find will return true if x is found in ranges. It will also
// return the index of x or -1 if x is not found.
func find(n uint32, ranges []float32, x float32) (bool, int) {
    var (
        lower, upper, mid int
    )
    if x < ranges[0] {
        return true, -1
    }
    if x >= ranges[n] {
        return true, -1
    }
    // Perform binary search
    upper = int(n)
    lower = 0

    for ul := (upper - lower); ul > 1; ul = (upper - lower) {
        mid = (upper + lower) / 2
        if x >= ranges[mid] {
            lower = mid
        } else {
            upper = mid
        }
    }
    // Sanity check the result
    if x < ranges[lower] || x >= ranges[lower+1] {
        fmt.Printf("x not found in range\n")
    }
    return false, lower
}

// Accumulate adds an item to the proper bin
func (h *histogram) Accumulate(x, weight float32) {
    n := h.n
    
    status, index := find(n, h.ranges, x)

    if status {
        fmt.Println("Somekind of error in accumulate.")
    }
    if index >= int(n) {
        fmt.Println("index is out of range")
    } else {
        h.bins[index] += weight
    }
}

// Get returns the value of the histogram at index i.
func (h *histogram) Get(i uint32) float32 {
    if i >= h.n {
        fmt.Printf("index lies outside of valid range of 0 .. n-1")
    } 
    return h.bins[i]
}

// GetRange returns the range that index i falls into.
func (h *histogram) GetRange(i uint32) (float32, float32) {
    if i >= h.n {
        fmt.Printf("index lies outside of valid range of 0 .. n-1")
    } 
    return h.ranges[i], h.ranges[i+1]
}

// Max returns the maximum range.
func (h *histogram) Max() float32 {
    return h.ranges[h.n]
}

// Min returns the minimum range.
func (h *histogram) Min() float32 {
    return h.ranges[0]
}

// Bins returns the number of bins in the histogram.
func (h *histogram) Bins() uint32 {
    return h.n
}

// MaxVal returns the maximum value in the histogram.
func (h *histogram) MaxVal() float32 {
    max := h.bins[0]

    var i uint32
    for i = 0; i < h.n; i++ {
        if h.bins[i] > max {
            max = h.bins[i]
        }
    }
    return max
}

// MaxBin returns the index of the bin that contains the max.
func (h *histogram) MaxBin() uint32 {
    var i uint32
    var imax uint32

    max := h.bins[0]
    for i = 0; i < h.n; i++ {
        if h.bins[i] > max {
            max = h.bins[i]
            imax = i
        }
    }
    return imax
}

// MinVal returns the maximum value in the histogram.
func (h *histogram) MinVal() float32 {
    min := h.bins[0]
    var i uint32
    for i = 0; i < h.n; i++ {
        if h.bins[i] < min {
            min = h.bins[i]
        }
    }
    return min
}

// MinBin returns the index of the bin that contains the min
func (h *histogram) MinBin() uint32 {
    var i uint32
    var imin uint32

    min := h.bins[0]
    for i = 0; i < h.n; i++ {
        if h.bins[i] < min {
            min = h.bins[i]
            imin = i
        }
    }
    return imin
}

// SameBins will return true if h2 uses the same binning as h1
func (h1 *histogram) SameBins(h2 *histogram) bool {
    if h1.n != h2.n {
        return false
    }
    var i uint32
    for i = 0; i <= h1.n; i++ {
        if h1.ranges[i] != h2.ranges[i] {
            return false
        }
    }
    return true
}

// Add will add two histograms and store the result in h1.
func (h1 *histogram) Add(h2 *histogram) {
    if !h1.SameBins(h2) {
        fmt.Println("histograms have different binning")
    }
    var i uint32
    for i = 0; i < h1.n; i++ {
        h1.bins[i] += h2.bins[i]
    }
}

// Subtract will subtract two histograms and store the result in h1.
func (h1 *histogram) Subtract(h2 *histogram) {
    if !h1.SameBins(h2) {
        fmt.Println("histograms have different binning")
    }
    var i uint32
    for i = 0; i < h1.n; i++ {
        h1.bins[i] -= h2.bins[i]
    }
}

// Multiply multiplies two histogram and stores the result in h1.
func (h1 *histogram) Multiply(h2 *histogram) {
    if !h1.SameBins(h2) {
        fmt.Println("histograms have different binning")
    }
    var i uint32
    for i = 0; i < h1.n; i++ {
        h1.bins[i] *= h2.bins[i]
    }
}

// Divide divides two histograms and stores the value in h1.
func (h1 *histogram) Divide(h2 *histogram) {
    if !h1.SameBins(h2) {
        fmt.Println("histograms have different binning")
    }
    var i uint32
    for i = 0; i < h1.n; i++ {
        h1.bins[i] /= h2.bins[i]
    }
}

// Scale scales a histogram by a factor of scale.
func (h1 *histogram) Scale(scale float32) {
    var i uint32
    for i = 0; i < h1.n; i++ {
        h1.bins[i] *= scale
    }
}

// Shift shifts a histogram by shift offset
func (h *histogram) Shift(shift float32) {
    var i uint32
    for i = 0; i < h.n; i++ {
        h.bins[i] += shift
    }
}

// Mean computes the bin weighted arithmetic mean of a histogram
// using the recurrence relation:
//   M(n) = M(n-1) + (x[n] - M(n-1)) (w(n)/W(n-1) + w(n)))
//   W(n) = W(n-1) + w(n)
func (h *histogram) Mean() float64 {
    n := h.n
    var wmean, W float64
    wmean = 0
    W = 0

    var i uint32
    for i = 0; i < n; i++ {
        xi := float64((h.ranges[i+1] + h.ranges[i]) / 2)
        wi := float64(h.bins[i])

        if wi > 0 {
            W += wi
            wmean += (xi - wmean) * (wi / W)
        }
    }
    return wmean
}

// Sigma will return the standard deviation of the histogram.
func (h *histogram) Sigma() float64 {
    n := h.n
    var i uint32

    var wvariance, wmean, W float64
    wvariance, wmean, W = 0, 0, 0

    // Use a two-pass alogorithm
    // First compute the mean...
    for i = 0; i < n; i++ {
        xi := float64((h.ranges[i+1] + h.ranges[i]) /2)
        wi := float64(h.bins[i])

        if (wi > 0) {
            W += wi
            wmean += (xi - wmean) * (wi /W)
        }
    }

    // ...then compute the variance
    for i = 0; i < n; i ++ {
        xi := float64((h.ranges[i+1] + h.ranges[i]) / 2)
        wi := float64(h.bins[i])

        if wi > 0 {
            delta := xi - wmean
            W += wi
            wvariance += (delta * delta - wvariance) * (wi / W)
        }
    }
    return math.Sqrt(wvariance)
}

// Sum returns the sum of every value in the histogram.
func (h *histogram) Sum() float64 {
    var sum float64
    sum = 0
    n := h.n

    var i uint32
    for i = 0; i < n; i++ {
        sum += float64(h.bins[i])
    }
    return sum
}
