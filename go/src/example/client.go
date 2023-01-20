package main

import (
    "io/ioutil"
    "io"
    "os"
    "net"
    "fmt"
    "strings"
    "bufio"

)

const (
	HOST = "localhost"
	PORT = "8888"
	TYPE = "tcp"
)

func main() {

        //PARTIE LECTURE FICHIER//    

        //Lecture de MatriceA
        file1, err := os.Open("MatriceA")
        if err != nil {
                panic(err)
        }
        defer file1.Close()

        //Lecture de MatriceB
        file2, err := os.Open("MatriceB")
        if err != nil {
                panic(err)
        }
        defer file2.Close()

        // Read the file into a byte slice
        matriceA, err := ioutil.ReadAll(file1)
        if err != nil {
                panic(err)
        }

        // Print the contents of the file
        //fmt.Println(string(matriceA))

        // Read the file into a byte slice
        matriceB, err := ioutil.ReadAll(file2)
        if err != nil {
                panic(err)
        }

        // Print the contents of the file
        //fmt.Println(string(matriceB))

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
                if strings.Contains(line, "..."){
                        break
                }
                fmt.Print(line)
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
        //var matriceR [][]int64         
 
        //Mettre la réponse dans un fichier text

        /*file, err := os.Create("matriceR")
        if err != nil {
                panic(err)
        }
        defer file.Close()*/

        // Write the matrix contents to the file
        /*for i := range matriceR {
                for j := range matriceR[i] {
                        _, err := file.WriteString(fmt.Sprintf("%d ", matriceR[i][j]))
                        if err != nil {
                                panic(err)
                        }
                }
                _, err := file.WriteString("\n")
                if err != nil {
                panic(err)
                }
        }*/
}
// Utiliser flush à la fin de chaque envoie pour pas que tcp attente
// Creer une vraie matriceR aussi