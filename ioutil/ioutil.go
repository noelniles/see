package ioutil

import (
    //"fmt"
    "image"
    "image/gif"
    "image/jpeg"
    "image/png"
    "log"
    "os"
    "path"
)

// ImRead attempts to read an image file and return an Image struct.
func ImRead(fn string) (img *image.Image, err error) {
    infile, e := os.Open(fn)    

    if e != nil {
        return nil, e
    }
    defer infile.Close()

    src, _, e := image.Decode(infile)
    if e != nil {
        return nil, e
    }
    return &src, err
}

// ImWrite attempts to write an image to a file.
func ImWrite(fn string, img *image.Image) (e error) {
    
    // Get the file extension.
    ext := path.Ext(fn)

    writer, e := os.Create(fn)
    if e != nil {
        return e
    }
    defer writer.Close()

    // Write a GIF.
    if ext == ".gif" {
        if e := gif.Encode(writer, *img, nil); e != nil {
            writer.Close()
            log.Fatal(e)
        }
    }
    // Write a JPEG.
    if ext == ".jpg" {
        if e := jpeg.Encode(writer, *img, nil); e != nil {
            writer.Close()
            log.Fatal(e)
        }
    }
    // Write a PNG.
    if ext == ".png" {
        if e := png.Encode(writer, *img); e != nil {
            writer.Close()
            log.Fatal(e)
        }
    }
    return e
}
