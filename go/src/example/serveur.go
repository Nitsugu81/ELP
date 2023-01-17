package main

import (
    "fmt"
    "net"
    //"time"
    "strings"
    "sync"
    "reflect"
    //"strconv"
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
            buf := make([]byte, 1024)
            taille, err := conn.Read(buf)
            fmt.Print("Nombre de bit :", taille, "\n")
            if err != nil {
                fmt.Printf("Error reading: %#v\n", err)
                return
            }
            matrices := string(buf[:taille])
            matrices_slices := strings.Split(matrices,"/////")
            matriceA := matrices_slices[0]
            //matriceB := matrices_slices[1]
            //var matrice1 [2][2] int
            //matrice2 := [][]int

            fmt.Print(reflect.TypeOf(matriceA))
            fmt.Printf("LA MATRICEA EST EGAL A : %s\n", matriceA)
            k := 0
            j := 0
            matriceA_lignes := strings.Split(matriceA,"\n")
            fmt.Print("POOP : ", len((((matriceA_lignes[1])))))
            for i := 0; i < len(matriceA_lignes); i++ { //ATTENTION, UN STRING C EST UNE CHAINE DE BYTE ET NON PAS UNE CHAINE DE CHARACTERES (Histoire avec rune)
                //fmt.Print("\nBLAAAAAAAAAAA :" + string((matriceA[i])))
                matriceA_elems := strings.Split(matriceA_lignes[i], " ")
                for j := 0; j < len(matriceA_elems); j++{
                    
                }
                if string(matriceA_lignes[i]) == " "{  // A enlever cette ligne eventuellement
                    continue
                }
                if string(matriceA_lignes[i]) == "\n"{
                    k++
                    j=0
                }else{
                    //matrice1[k][j],_ = strconv.Atoi((string(matriceA_lignes[i])))
                    j++ 
                }
            }
            fmt.Println()

            /*for i := range matrice1 {
                for j := range matrice1[i] {
                    fmt.Printf("%d ",matrice1[i][j])
                }
                fmt.Println()
            }*/

            

            /*lignes_matriceA := strings.Split(matriceA,"\n")
            lignes_matriceB := strings.Split(matriceB,"\n")
            for i,v := range lignes_matriceA{
                fmt.Println("\nIndex : ", i, "Valeur : ", v)
            }
            for i,v := range lignes_matriceB{
                fmt.Println("\nIndex : ", i, "Valeur : ", v)
            }
            colonnes_matriceB = []int

            for i, v := range matrices_slices {
                fmt.Println("\nIndex : ", i, "Valeur : ", v)
            }*/
            
            conn.Write([]byte("Message received.\n"))
            conn.Close()
        }(conn)
    }
}

func multiply(a [][]int, b [][]int, result [][]int, row int, wg *sync.WaitGroup) {
    defer wg.Done()
    for col := 0; col < len(b[0]); col++ {
        for k := 0; k < len(a[0]); k++ {
            result[row][col] += a[row][k] * b[k][col]
        }
    }
}
