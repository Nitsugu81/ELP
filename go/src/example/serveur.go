package main

import (
    "fmt"
    "net"
    "strings"
    "sync"
    "reflect"
    "strconv"
    "encoding/json"
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
            matriceA_lignes := strings.Split(matriceA,"\n")
            matriceB_lignes := strings.Split(matriceB,"\n")

            nb_lignes_matrice1 := (len(matriceA_lignes))
            nb_lignes_matrice2 := (len(matriceB_lignes))
            nb_colonnes_matrice1 := (len(strings.Split(matriceA_lignes[0], " ")))
            nb_colonnes_matrice2  := (len(strings.Split(matriceB_lignes[0], " ")))

            fmt.Print("nb_lignes : ", nb_lignes_matrice1)
            fmt.Println(" de type : ", reflect.TypeOf(nb_lignes_matrice1))
            //tata := 2
            matrice1 := make([][]int, nb_lignes_matrice1)
            for i := 0; i < nb_lignes_matrice1; i++ {
                matrice1[i] = make([]int, nb_colonnes_matrice1)
            } //On pouvait pas passer par des arrays parce que la taille est pas constante (depend des matrices envoyées) donc ca ne marchait pas
            matrice2 := make([][]int, nb_lignes_matrice2)
            for i := 0; i < nb_lignes_matrice2; i++ {
                matrice2[i] = make([]int, nb_colonnes_matrice2)
            }
            matriceR := make([][]int, nb_lignes_matrice1)
            for i := 0; i < nb_lignes_matrice1; i++ {
                matriceR[i] = make([]int, nb_colonnes_matrice2)
            }

            if nb_colonnes_matrice1 != nb_lignes_matrice2{
                conn.Write([]byte("Il faut autant de colonnes pour la matriceA que de lignes pour la matriceB\n"))
                conn.Close()
                return 
            }

            //fmt.Print(reflect.TypeOf(matriceA))
            //fmt.Printf("LA MATRICEA EST EGAL A : %s\n", matriceA)

            //fmt.Print("POOP : ", len((((matriceA_lignes[1])))))
            
            /*for i := 0; i < len(matriceA_lignes); i++ { //ATTENTION, UN STRING C EST UNE CHAINE DE BYTE ET NON PAS UNE CHAINE DE CHARACTERES (Histoire avec rune)
                //fmt.Print("\nBLAAAAAAAAAAA :" + string((matriceA[i])))
                matriceA_elems := strings.Split(matriceA_lignes[i], " ")
                for j := 0; j < len(matriceA_elems); j++{
                    matrice1[i][j],_ = strconv.Atoi((string(matriceA_elems[j])))
                }
                    
            }*/

            for i := 0; i < nb_lignes_matrice1; i++ { //ATTENTION, UN STRING C EST UNE CHAINE DE BYTE ET NON PAS UNE CHAINE DE CHARACTERES (Histoire avec rune)
                //fmt.Print("\nBLAAAAAAAAAAA :" + string((matriceA[i])))
                matriceA_elems := strings.Split(matriceA_lignes[i], " ")
                for j := 0; j < nb_colonnes_matrice1; j++{
                    matrice1[i][j],err = strconv.Atoi((string(matriceA_elems[j])))
                    if err != nil {
                        fmt.Print("ERREUR")
                    }
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

            for i := 0; i < nb_lignes_matrice2; i++ { //ATTENTION, UN STRING C EST UNE CHAINE DE BYTE ET NON PAS UNE CHAINE DE CHARACTERES (Histoire avec rune)
                //fmt.Print("\nBLAAAAAAAAAAA :" + string((matriceA[i])))
                matriceB_elems := strings.Split(matriceB_lignes[i], " ")
                for j := 0; j < nb_colonnes_matrice2; j++{
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

            matrice_encodee, err := json.Marshal(matriceR)
            if err != nil {
                panic(err)
            }
            fmt.Print("matrice encodee : ", matrice_encodee)
            fmt.Println(" de type : ", reflect.TypeOf(matrice_encodee))
            
            conn.Write([]byte(matrice_encodee))
            conn.Close()
        }(conn)
    }
}

func multiply(a [][]int, b [][]int, result *[][]int, row int, wg *sync.WaitGroup) {
    defer wg.Done()
    for col := 0; col < len(b[0]); col++ {
        for k := 0; k < len(a[0]); k++ {
            (*result)[row][col] += a[row][k] * b[k][col]
        }
    }
}
