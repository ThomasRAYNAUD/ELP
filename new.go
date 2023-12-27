package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	// À compléter : ajoutez le code pour gérer la connexion (lecture/écriture de données, etc.)
	//fmt.Fprintf(conn, "Bonjour, client!\n")

	// Fermez la connexion lorsque vous avez terminé
	defer conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":8777") //listen permet de créer un serveur
	if err != nil { //si err n'est pas vide 
		fmt.Println("Erreur lors de l'écoute sur le port 8777, le message d'erreur --> ", err)
		return
	}
	defer listener.Close()

	fmt.Println("Serveur en attente de connexions sur le port 8777...")

	// Boucle infinie pour accepter les connexions entrantes
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erreur lors de l'acceptation de la connexion:", err)
			continue
		}
		fmt.Println("Nouvelle connexion établie.")
		// Gérez la connexion ici (par exemple, lisez/écrivez des données)
		go handleConnection(conn)
	}
}