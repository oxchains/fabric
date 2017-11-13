package main

import (
    "net"
    "fmt"
    "encoding/binary"
    "time"
)

func sendUint32(length uint32, conn net.Conn) (int, error) {
    
    var buf [4]byte
    
    binary.BigEndian.PutUint32(buf[:], uint32(length))
    
    return conn.Write(buf[:])
}

func sendUint64(length uint64, conn net.Conn) (int, error) {
    
    var buf [8]byte
    
    binary.BigEndian.PutUint64(buf[:], uint64(length))
    
    return conn.Write(buf[:])
}

func main() {
    conn, err := net.Dial("unix", "/tmp/bft.sock")
    if err != nil {
        fmt.Println("Could not connect to proxy!")
        panic(err)
    }
    fmt.Println("connected to proxy!")
    
    if err != nil {
        fmt.Println("Could not connect to proxy!")
        return
    }
    
    for i := 0; i < 100; i++{
        var buf [4]byte
    
        binary.BigEndian.PutUint32(buf[:], uint32(10))
        fmt.Println(i)
    }
    
    //JCS: Sending pool size
    _, err = sendUint32(uint32(10), conn)
    
    if err != nil {
        fmt.Println("Error while sending pool size:", err)
        return
    }
    
    //JCS: Sending batch configuration
    _, err = sendUint32(uint32(10 * 1024), conn)
    
    if err != nil {
        fmt.Println("Error while sending PreferredMaxBytes:", err)
        return
    }
    
    _, err = sendUint32(uint32(10), conn)
    
    if err != nil {
        fmt.Println("Error while sending MaxMessageCount:", err)
        return
    }
    _, err = sendUint64(uint64(time.Duration.Nanoseconds(time.Second * 2)), conn)
    
    if err != nil {
        fmt.Println("Error while sending BatchTimeout:", err)
        return
    }
    
    //conn.Close()
    
    
    //tmp := 10
    //status, err := conn.Write([]byte{0,1,2,3})
    //if err != nil {
    //    fmt.Println("Could not write to proxy!")
    //    panic(err)
    //}
    //fmt.Println(status)
}
