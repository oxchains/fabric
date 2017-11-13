package bftsmart

import (
    "testing"
    "os"
    "fmt"
    "time"
)

func TestNew(t *testing.T) {
    host, err := os.Hostname()
    a := time.Now().UTC().String()
    fmt.Println(a)
    fmt.Println(host)
    fmt.Println(err)
}
