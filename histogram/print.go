package histogram

import (
    "fmt"
)

// Print will print a histogram to the screen. This
// is pretty slow and useless actually.
func (h *histogram) Print() string {
    var i uint32
    var res string

    for i = 0; i < h.Bins(); i++ {
        val := h.Get(i)  
        fmt.Println("val: ", val)
        for j := 0; float32(j) < val; j++ {
            res += ("*")
        }
        fmt.Printf("%d -- %s", i, res)
        res += "\n"
    }
    return res
}
