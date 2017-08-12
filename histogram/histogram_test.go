package histogram

import (
    //"fmt"
    "testing"
)

func TestHistogram1D(t *testing.T) {
    var xr = []float32{0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}

    // Test NewHist.
    h, err := NewHist(256)
    if err != nil {
        t.Error(err)
    }
    if h.ranges == nil {
        t.Error("NewHist returned an invalid range array.")
    }
    if h.bins == nil {
        t.Error("NewHist returned an invalid bin array.")
    }
    if h.n != 256 {
        t.Error("NewHist returned an invalid size.")
    }   

    hr, err := AllocateRanges(10, xr)
    if err != nil {
        t.Error(err)
    }
    if hr.ranges == nil {
        t.Error("NewHist returned an invalid range array.")
    }
    if hr.bins == nil {
        t.Error("NewHist returned an invalid bin array.")
    }
    if hr.n != 10 {
        t.Error("NewHist returned an invalid size.")
    }   

    // Zero the ranges
    for i := 0; i <= 10; i++ {
        hr.ranges[i] = 0.0
    }

    // Test SetRanges
    hr.SetRanges(xr, 11)

    for i := 0; i <= 10; i++ {
        if hr.ranges[i] != xr[i] {
            t.Error("SetRanges didn't set the ranges.") 
        }
    }

    // Test Accumulate.
    for i := 0; i < 256; i++ {
        h.Accumulate(float32(i), float32(i))
    }

    for i := 0; i < 256; i++ {
        if h.bins[i] != float32(i) {
            t.Error("Accumulate didn't write to the array.")
        } 
    }

    // Test Get.
    for i := 0; i > 256; i++ {
        if h.Get(uint32(i)) != float32(i) {
            t.Error("Get didn't read from the array properly.")
        }
    }
    // Test Copy.
    h1, err := h.Copy()
    if err != nil {
        t.Error("Copy messed up %f", err)
    }
    for i := 0; i <= 256; i++ {
        if h.ranges[i] != h1.ranges[i] {
            t.Error("Copy miscopied the ranges.")
        }
    }
    for i := 0; i < 256; i++ {
        if h.bins[i] != h1.bins[i] {
            t.Error("Copy miscopied the bins.")
        }
    }
    
    // Test Reset.
    h.Reset()
    for i := 0; i < 256; i++ {
        if h.bins[1] != 0 {
            t.Error("Reset didn't zero the array.")
        }
    }

    // Test Increment.
    for i := 0; i < 256; i++ {
        h.Increment(float32(i))

        for j := 0; j <= i; j++ {
            if h.bins[i] != 1 {
                t.Error("Increment didn't add one to the bin.")
            }
        }
    }

    // Test GetRange
    for i := 0; i < 256; i++ {
        x0, x1 := h.GetRange(uint32(i))

        if x0 != float32(i) || x1 != float32(i) + 1.0 {
            t.Error("GetRange didn't return the proper range.")
        }
    }

    // Test Max.
    if h.Max() != 256 {
        t.Errorf("Max expected 1 got %f", h.Max())
    }
    
    // Test Min.
    if h.Min() != 0.0 {
        t.Error("Min returned an incorrect value.")
    }
    
    // Test Bins.
    if h.Bins() != 256 {
        t.Error("Bins returned an incorrect value.")
    }

    // Create min and max bins.
    h.bins[2] = 123456.0
    h.bins[4] = -654321

    // Test MaxVal
    max := h.MaxVal()
    if max != 123456.0 {
        t.Error("MaxVal returned an incorrect value.")
    }

    // Test MinVal.
    min := h.MinVal()
    if min != -654321 {
        t.Error("MinVal returned an incorrect value.")
    }

    // Test MaxBin.
    imax := h.MaxBin()
    if imax != 2 {
        t.Error("MaxBin returned an incorrect bin.")
    }
    
    // Test MinBin.
    imin := h.MinBin()
    if imin != 4 {
        t.Error("MinBin returned an incorrect bin.")
    }

    // Test Sum.
    for i := 0; i < 256; i++ {
        h.bins[i] = float32(i + 27)
    }
    sum := h.Sum()
    if sum != 256*27 + (255*256) / 2 {
        t.Error("Sum returned an incorrect result.")
    }

    // Test Add.
    g, err := h1.Copy()
    if err != nil {
        t.Error("Copy messed up")
    }
    h1.Add(h)
    for i := 0; i < 256; i++ {
        if h1.bins[i] != g.bins[i] + h.bins[i] {
            t.Error("Add added wrong.")
        }
    }
    // Test Subtract
    g, err = h1.Copy()
    if err != nil {
        t.Error("Copy messed up")
    }
    h1.Subtract(h)
    for i := 0; i < 256; i++ {
        if h1.bins[i] != g.bins[i] - h.bins[i] {
            t.Error("Subtract subtracted wrong.")
        }
    }

    // Test Multiply.
    g, err = h1.Copy()
    if err != nil {
        t.Error("Copy messed up")
    }
    h1.Multiply(h)
    for i := 0; i < 256; i++ {
        if h1.bins[i] != g.bins[i] * h.bins[i] {
            t.Error("Multiply multiplied wrong.")
        }
    }

    // Test Divide.
    g, err = h1.Copy()
    if err != nil {
        t.Error("Copy messed up")
    }
    h1.Divide(h)
    for i := 0; i < 256; i++ {
        if h1.bins[i] != g.bins[i] / h.bins[i] {
            t.Error("Multiply multiplied wrong.")
        }
    }

    // Test Shift.
    g, err = h1.Copy()
    if err != nil {
        t.Error("Copy messed up")
    }
    h1.Shift(0.5)
    for i := 0; i < 256; i++ {
        if h1.bins[i] != g.bins[i] + 0.5 {
            t.Error("Shift shifted wrong.")
        }
    }

    // Test Scale
    g, err = h1.Copy()
    if err != nil {
        t.Error("Copy messed up")
    }
    h1.Scale(0.5)
    for i := 0; i < 256; i++ {
        if h1.bins[i] != g.bins[i] * 0.5 {
            t.Error("Scale scaled wrong.")
        }
    }
}
