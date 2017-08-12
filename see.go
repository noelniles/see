package main

import (
    "flag"
    "fmt"
    "os"
    //"image/color"
    . "github.com/see/histogram"
    . "github.com/see/ioutil"
    . "github.com/see/test_images"
)

func main() {
    ifn := flag.String("in", "foo", "input filename")
    ofn := flag.String("out", "bar", "output filename")
    flag.Parse()

    if *ifn == "foo" || *ofn == "bar" {
        flag.Usage()
        os.Exit(2)
    }
    img, e := ImRead(*ifn)
    if e != nil {
        fmt.Errorf("Problem reading file: %s\n", *ifn)
        os.Exit(2)
    }

    rh, gh, bh, e := ImageHistogram(img)

    fmt.Println("rh max: ", rh.Max())
    fmt.Println("gh max: ", gh.Max())
    fmt.Println("bh max: ", bh.Max())
    fmt.Println("e: ", e)

    fmt.Println("rh min: ", rh.Min())
    fmt.Println("gh min: ", gh.Min())
    fmt.Println("bh min: ", bh.Min())

    fmt.Println("rh maxval: ", rh.MaxVal())
    fmt.Println("gh maxval: ", gh.MaxVal())
    fmt.Println("bh maxval: ", bh.MaxVal())

    fmt.Println("rh maxbin: ", rh.MaxBin())
    fmt.Println("gh maxbin: ", gh.MaxBin())
    fmt.Println("bh maxbin: ", bh.MaxBin())

    fmt.Println("rh minbin: ", rh.MinBin())
    fmt.Println("gh minbin: ", gh.MinBin())
    fmt.Println("bh minbin: ", bh.MinBin())

    fmt.Println("rh mean: ", rh.Mean())
    fmt.Println("gh mean: ", gh.Mean())
    fmt.Println("bh mean: ", bh.Mean())

    fmt.Println("rh sigma: ", rh.Sigma())
    fmt.Println("gh sigma: ", gh.Sigma())
    fmt.Println("bh sigma: ", bh.Sigma())

    fmt.Println("rh sum: ", rh.Sum())
    fmt.Println("gh sum: ", gh.Sum())
    fmt.Println("bh sum: ", bh.Sum())

    // Try out pdf histogram
    p := NewHistPdf(rh)
    fmt.Println((*p).Sum()) 
    fmt.Println((*p).SumAt(133))
    fmt.Println((*p).Sample(0.0051912568))
    fmt.Println((*p).Sample(0.003353204))
    fmt.Println((*p).Sample(0.0))

    wimg := WhiteImage()

    for j := 0; j < wimg.Bounds().Max.Y; j++ {
        for i := 0; i < wimg.Bounds().Max.X; i++ {
            r, g, b, a := wimg.At(i, j).RGBA()
            r = r>>8
            g = g>>8
            b = b>>8
            a = a>>8

            if r != 255 || g != 255 || b != 255 || a != 255 {
                fmt.Errorf("Wrong")
            }
        } 
    }
}
