package main

import (
    //"fmt"
    "io/ioutil"
    "io"
    "os"
    "net"
    //"reflect"
    "encoding/json"
    "fmt"
    "reflect"
    "bytes"
    "strings"

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

        _, err = conn.Write([]byte(matrices))
        if err != nil {
                println("Write matriceA failed:", err.Error())
                os.Exit(1)
        }

        // Chopper la matriceR

        /*received, err := ioutil.ReadAll(conn)
	if err != nil {
		if err != io.EOF {
			fmt.Println("read error:", err)
		}
		panic(err)
	}*/

        received := bytes.NewBuffer(nil)
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
                received.Write(tmp[:n])
                if bytes.Contains(received.Bytes(), delimiter) {
                    break
                }
        }

        fmt.Print(received.String())
        fmt.Println(" de type : ", reflect.TypeOf(received.String()))
        new := strings.Split(received.String(),"...")
        fmt.Print(" new : ", new)
        fmt.Println(" de longueur : ", len(new))

        /*received := make([]byte, 1024)
        n, err := conn.Read(received)
        if err != nil {
                println("Read matriceR failed:", err.Error())
                os.Exit(1)
        }
        /*matriceR := make([][]int, 3)
        for i := 0; i < 3; i++ {
            matriceR[i] = make([]int, 2)
        }*/
        var matriceR [][]int64         
        json.Unmarshal(received.Bytes(), &matriceR)

        fmt.Println("Type : ", reflect.TypeOf(received))
        println(received, "donne : ", (received.String()))
        fmt.Print("MatriceR : ", matriceR)
        fmt.Print(" de type : ", reflect.TypeOf(matriceR))

        //Mettre la réponse dans un fichier text

        file, err := os.Create("matriceR")
        if err != nil {
                panic(err)
        }
        defer file.Close()

        // Write the matrix contents to the file
        for i := range matriceR {
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
        }


        conn.Close()
}


// Utiliser flush à la fin de chaque envoie pour pas que tcp attente
// Utiliser ReadString (ou Readlines ?) du côté du client parce que du côté du serveur on va envoyer la matrice ligne par ligne. Recommendation du prof : ReadString("\n"), comme ca on peut remplir le fichier en même temps qu'on lit. 