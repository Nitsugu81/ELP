package main

import (
    "io/ioutil"
    "io"
    "os"
    "net"
    "fmt"
    "strings"
    "bufio"
    "strconv"

)

const (
	HOST = "localhost"
	PORT = "8888"
	TYPE = "tcp"
)

func main() {

        //PARTIE LECTURE FICHIER//    

        //Ouverture MatriceA
        file1, err := os.Open("MatriceA")
        if err != nil {
                panic(err)
        }
        defer file1.Close()

        //Ouverture MatriceB
        file2, err := os.Open("MatriceB")
        if err != nil {
                panic(err)
        }
        defer file2.Close()

        // Lecture MatriceA
        matriceA, err := ioutil.ReadAll(file1)
        if err != nil {
                panic(err)
        }
        matriceA_lignes := strings.Split(string(matriceA),"\n")
        nb_lignes_matrice1 := (len(matriceA_lignes))

        // Lecture MatriceB
        matriceB, err := ioutil.ReadAll(file2)
        if err != nil {
                panic(err)
        }
        matriceB_lignes := strings.Split(string(matriceB),"\n")
        nb_colonnes_matrice2  := (len(strings.Split(matriceB_lignes[0], " ")))

        matrices := string(matriceA) + "/////" + string(matriceB)

        //PARTIE ENVOI AU SERVEUR//

        tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
        if err != nil {
                println("ResolveTCPAddr failed:", err.Error())
                os.Exit(1)
        }

        conn, err := net.DialTCP(TYPE, nil, tcpServer) //nil represente l'adresse locale
        if err != nil {
                println("Dial failed:", err.Error())
                os.Exit(1)
        }

        writer := bufio.NewWriter(conn) // Utilisation de bufio pour pouvoir utiliser flush
        _, err = writer.Write([]byte(matrices))
        if err != nil {
                println("Write matrices failed:", err.Error())
                os.Exit(1)
        }
        writer.Flush() //Permet d'envoyer les données directement sans avoir à attendre que le buffer se remplisse ou qu'il s'envoie après un timeout. 

        //PARTIE RECEP DES DONNEES ENVOYEES PAR LE SERVEUR

        reader := bufio.NewReader(conn)
        matriceR := make([][]int, nb_lignes_matrice1)
            for i := 0; i < nb_lignes_matrice1; i++ {
                matriceR[i] = make([]int, nb_colonnes_matrice2)
            }
        compteur := 0
        if err != nil {
                panic(err)
        }
        file, err := os.Create("matriceR")
        if err != nil {
                panic(err)
        }
        defer file.Close()
        for {
                line, err := reader.ReadString('\n')
                if strings.Contains(line,"Il faut autant de colonnes pour la matriceA que de lignes pour la matriceB\n"){
                        fmt.Println(line)
                        break
                }
                ligne_matrice := strings.Split(line," ")
                if strings.Contains(line, "..."){
                        break
                }
                for i,v := range ligne_matrice{ //Utilse si jamais on veut avoir la matrice R dans une variable et pas que dans un fichier text. (si jamais le client veut la remanipuler sans passer par lecture de fichier)
                        if i != len(ligne_matrice) -1{
                                matriceR[compteur][i],_ = strconv.Atoi(v)
                        }else{
                                matriceR[compteur][i],_ = strconv.Atoi(strings.Split(v,"\n")[0])
                        }                      
                }
                compteur ++ 
                if err != nil {
                        if err == io.EOF {
                                break
                        }
                        panic(err)
                }
                _, err = file.WriteString(line)
                if err != nil {
                        panic(err)
                }
        }

        fmt.Println("Matrice R : ")
        for i := range matriceR {
                for j := range matriceR[i] {
                    fmt.Printf("%d ",matriceR[i][j])
                }
                fmt.Println()
            }
 
}