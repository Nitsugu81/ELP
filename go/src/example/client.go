package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func main() {
	// Se connecter au serveur
	conn, err := net.Dial("tcp", "localhost:8080") //net.Dial permet d'établir la connection au serveur, ses arguments son le type de protocole et l'adresse utilisés, elle renvoie deux inforfations conn (pour les donées échangées) et err pour les erreures
	if err != nil {                                //renvoi l'erreur si il y en a une
		panic(err)
	}
	defer conn.Close() //ferme la connection

	// Créer un tableau à envoyer
	arr := []int{1, 2, 3, 4, 5}

	// Encode l'array en utlisant la méthode gob et l'envoie
	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(arr)
	if err != nil {
		panic(err)
	}

	// Recevoir le tableau inversé
	var reversedArr []int
	decoder := gob.NewDecoder(conn)
	err = decoder.Decode(&reversedArr)
	if err != nil {
		panic(err)
	}

	// Afficher le tableau inversé
	fmt.Println("Tableau inversé : ", reversedArr)
}
