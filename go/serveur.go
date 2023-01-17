package main

import (
    "fmt"
    "net"
    "strings"
    "sync"
    //"reflect"
    "strconv"
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
            matriceB := matrices_slices[1]

            nb_lignes_matrice1 :=
            nb_lignes_matrice2 :=
            nb_colonnes_matrice1 :=
            nb_colonnes_matrice2 :=_

            var matrice1 [2][2] int
            var matrice2 [2][2] int
            var matriceR [2][2] int 

            //fmt.Print(reflect.TypeOf(matriceA))
            //fmt.Printf("LA MATRICEA EST EGAL A : %s\n", matriceA)

            matriceA_lignes := strings.Split(matriceA,"\n")
            matriceB_lignes := strings.Split(matriceB,"\n")
            //fmt.Print("POOP : ", len((((matriceA_lignes[1])))))
            for i := 0; i < len(matriceA_lignes); i++ { //ATTENTION, UN STRING C EST UNE CHAINE DE BYTE ET NON PAS UNE CHAINE DE CHARACTERES (Histoire avec rune)
                //fmt.Print("\nBLAAAAAAAAAAA :" + string((matriceA[i])))
                matriceA_elems := strings.Split(matriceA_lignes[i], " ")
                for j := 0; j < len(matriceA_elems); j++{
                    matrice1[i][j],_ = strconv.Atoi((string(matriceA_elems[j])))
                }
                    
            }

            fmt.Println("\nMatrice A : ")

            for i := range matrice1 {
                for j := range matrice1[i] {
                    fmt.Printf("%d ",matrice1[i][j])
                }
                fmt.Println()
            }
            
            fmt.Println("Matrice B : ")

            for i := 0; i < len(matriceB_lignes); i++ { //ATTENTION, UN STRING C EST UNE CHAINE DE BYTE ET NON PAS UNE CHAINE DE CHARACTERES (Histoire avec rune)
                //fmt.Print("\nBLAAAAAAAAAAA :" + string((matriceA[i])))
                matriceB_elems := strings.Split(matriceB_lignes[i], " ")
                for j := 0; j < len(matriceB_elems); j++{
                    matrice2[i][j],_ = strconv.Atoi((string(matriceB_elems[j])))
                }
                    
            }
            
             for i := range matrice2 {
                for j := range matrice2[i] {
                    fmt.Printf("%d ",matrice2[i][j])
                }
                fmt.Println()
            }

            var wg sync.WaitGroup
            for i := range matrice1 {
                wg.Add(1)
                go multiply(matrice1, matrice2, &matriceR, i, &wg)
            }
            wg.Wait()
        
            fmt.Println(matriceR)
            
            conn.Write([]byte("Message received.\n"))
            conn.Close()
        }(conn)
    }
}

func multiply(a [2][2]int, b [2][2]int, result *[2][2]int, row int, wg *sync.WaitGroup) {
    defer wg.Done()
    for col := 0; col < len(b[0]); col++ {
        for k := 0; k < len(a[0]); k++ {
            result[row][col] += a[row][k] * b[k][col]
        }
    }
}
