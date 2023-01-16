package main

import (
    //"fmt"
    "io/ioutil"
    "os"
    "net"
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

        matrices := string(matriceA) + string(matriceB)

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

        // buffer to get matriceA
        received := make([]byte, 1024)
        _, err = conn.Read(received)
        if err != nil {
                println("Read matriceA failed:", err.Error())
                os.Exit(1)
        }

        println("Received message:", string(received))

        conn.Close()

}