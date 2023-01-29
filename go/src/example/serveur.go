package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
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
			//buffer dynamique dont la taille augmente pour pouvoir lire les données en optimisant la taille du buffer
			buf := bytes.NewBuffer(nil)
			delimiter := []byte("...")
			for {
				tmp := make([]byte, 256) //De petite taille pour pas trop lire (en gros on lit 256 bytes par 256 bytes et on rajoute ca dans un buffer, comme ca ca evite de faire un buffer trop gros ou trop petit)
				n, err := conn.Read(tmp)
				if err != nil {
					if err == io.EOF {
						break
					}
					fmt.Printf("Error reading: %#v\n", err)
					return
				}
				buf.Write(tmp[:n])
				if bytes.Contains(buf.Bytes(), delimiter) {
					break
				}
			}
			//on recrée les matrices à partir des données récupérées
			matrices := buf.String()
			matrices_slices := strings.Split(matrices, "/////")
			matriceA := matrices_slices[0]
			matriceB := matrices_slices[1]
			matriceA_lignes := strings.Split(matriceA, "\n")
			matriceB_lignes := strings.Split(matriceB, "\n")
			matriceB_lignes = matriceB_lignes[:len(matriceB_lignes)-1]

			nb_lignes_matrice1 := (len(matriceA_lignes))
			nb_lignes_matrice2 := (len(matriceB_lignes))
			nb_colonnes_matrice1 := (len(strings.Split(matriceA_lignes[0], " ")))
			nb_colonnes_matrice2 := (len(strings.Split(matriceB_lignes[0], " ")))

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
			
			//vérification que la multiplication est possible
			if nb_colonnes_matrice1 != nb_lignes_matrice2 {
				conn.Write([]byte("Il faut autant de colonnes pour la matriceA que de lignes pour la matriceB\n"))
				conn.Close()
				return
			}

			for i := 0; i < nb_lignes_matrice1; i++ { //Si jamais le user fait des lignes avec des nombres de colonne differentes (pas opti car 2 boucles for mais ca fait crash le serveur sinon)
				temp := len(strings.Split(matriceA_lignes[i], " "))
				if temp != nb_colonnes_matrice1 {
					conn.Write([]byte("Il faut autant de colonnes pour la matriceA que de lignes pour la matriceB\n"))
					conn.Close()
					return
				}
			}

			for i := 0; i < nb_lignes_matrice2; i++ {
				temp := len(strings.Split(matriceB_lignes[i], " "))
				if temp != nb_colonnes_matrice2 {
					conn.Write([]byte("Il faut autant de colonnes pour la matriceA que de lignes pour la matriceB\n"))
					conn.Close()
					return
				}
			}

			for i := 0; i < nb_lignes_matrice1; i++ { //ATTENTION, UN STRING C EST UNE CHAINE DE BYTE ET NON PAS UNE CHAINE DE CHARACTERES (Histoire avec rune)
				matriceA_elems := strings.Split(matriceA_lignes[i], " ")
				for j := 0; j < nb_colonnes_matrice1; j++ {
					matrice1[i][j], err = strconv.Atoi((string(matriceA_elems[j])))
					if err != nil {
						fmt.Print("ERREUR")
					}
				}

			}
			
			//Affiche les matrices reçues
			fmt.Println("\nMatrice A : ")

			for i := range matrice1 {
				for j := range matrice1[i] {
					fmt.Printf("%d ", matrice1[i][j])
				}
				fmt.Println()
			}

			fmt.Println("Matrice B : ")

			for i := 0; i < nb_lignes_matrice2; i++ {
				matriceB_elems := strings.Split(matriceB_lignes[i], " ")
				for j := 0; j < nb_colonnes_matrice2; j++ {
					matrice2[i][j], _ = strconv.Atoi((string(matriceB_elems[j])))
				}

			}

			for i := range matrice2 {
				for j := range matrice2[i] {
					fmt.Printf("%d ", matrice2[i][j])
				}
				fmt.Println()
			}
			
			//création d'un wait groupe pour faire les calculs ligne par ligne
			var wg sync.WaitGroup
			for i := range matrice1 {
				wg.Add(1)
				go multiply(matrice1, matrice2, &matriceR, i, &wg)
			}
			wg.Wait()
			for i := range matriceR {
				for j := range matriceR[i] {
					conn.Write([]byte(strconv.Itoa(matriceR[i][j])))
					if j != len(matriceR[i])-1 {
						conn.Write([]byte(" "))
					}
				}
				conn.Write([]byte("\n"))
			}
			conn.Write([]byte("...\n"))
			conn.Close() //flush integré
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
