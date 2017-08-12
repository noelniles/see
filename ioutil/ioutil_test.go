package ioutil

import (
    "fmt"
    "testing"
)

func TestImRead(t *testing.T) {
    // Test ImRead from non-existant JPEG file.
    _, err := ImRead("non-existant-file.jpg")
    fmt.Println(err)

    // The file doesn't exist so we should throw an error.
    if err == nil {
        t.Error("ImRead didn't throw an error reading non-existant image.")
    }

    // Test ImRead with existant JPEG file.
    src, err := ImRead("../../images/png/utah_teapot.png")
    
    // This should work so make sure err is nil.
    if err != nil {
        t.Error("ImRead threw an error.")
    }

    // Check that the dimensions are correct.
    if (*src).Bounds().Max.X != 2733 || (*src).Bounds().Max.Y != 1809 {
        t.Error("ImRead returned image with wrong dimensions.")
    }
}

func TestImWrite(t *testing.T) {
    exists   := "../../images/png/utah_teapot.png"
    testpng  := "ImWrite-test-file.png"
    testjpeg := "ImWrite-test-file.jpg"
    //testgif  := "ImWrite-test-file.gif"
    
    src, e := ImRead(exists)
    if e != nil {
        t.Error("ImRead screwed up during the ImWrite test.")
    }

    // Test writing in different formats.
    // Temporarily remove the GIF test because writing GIFs is slow.
    //e = ImWrite(testgif, src)
    //if e != nil {
    //    t.Error("ImWrite had a problem writing a gif.")
    //}
    e = ImWrite(testjpeg, src)
    if e != nil {
        t.Error("ImWrite had a problem writing a jpeg.")
    }
    e = ImWrite(testpng, src)
    if e != nil {
        t.Error("ImWrite had a problem writing a png.")
    }
}
