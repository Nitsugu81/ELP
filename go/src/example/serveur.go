package main

import (
    "fmt"
    "net"
    //"time"
    "strings"
)

func main() {
    // Listen for incoming connections.
    addr := "localhost:8888"
    l, err := net.Listen("tcp", addr)
    if err != nil {
        panic(err)
    }
    defer l.Close()
    host, port, err := net.SplitHostPort(l.Addr().String())
    if err != nil {
        panic(err)
    }
    fmt.Printf("Listening on host: %s, port: %s\n", host, port)

    for {
        // Listen for an incoming connection

        conn, err := l.Accept()
        if err != nil {
            panic(err)
        }
        go func(conn net.Conn) {
            var i int = 0
            fmt.Print("\ni: ", i, "\n")
            buf := make([]byte, 1024)
            len, err := conn.Read(buf)
            fmt.Print("Nombre de bit :", len, "\n")
            if err != nil {
                fmt.Printf("Error reading: %#v\n", err)
                return
            }
            matrices := string(buf[:len])
            matrices_slices := strings.Split(matrices," ")
            fmt.Printf("Message received: \n%s\n", matrices_slices)
            for i, v := range matrices_slices {
                fmt.Println("\nIndex : ", i, "Valeur : ", v)
            }
            

            conn.Write([]byte("Message received.\n"))
            conn.Close()
            i++
        }(conn)
    }
}
