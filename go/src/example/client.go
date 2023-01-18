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

        received, err := ioutil.ReadAll(conn)
	if err != nil {
		if err != io.EOF {
			fmt.Println("read error:", err)
		}
		panic(err)
	}

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
        json.Unmarshal(received, &matriceR)

        fmt.Println("Type : ", reflect.TypeOf(received))
        println(received, "donne : ", string(received))
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