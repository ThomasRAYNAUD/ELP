package main

import (
	"fmt"
	"net"
)

func main() {
	port:="20001"
	listener, err := net.Listen("tcp", ":"+port)
	//créer un listener sur un port TCP spécifié
	if err != nil {
		fmt.Printf("Erreur lors de l'écoute sur le port %s, le message d'erreur --> %v\n", port, err)
		return
	}
	fmt.Printf("Serveur en attente de connexions sur le port %s...\n", port)
	for { //équivalent d'un while true qui n'existe pas en go
		_, err := listener.Accept() //accept la connection qd qlq
		if err != nil {
			fmt.Println("Erreur lors de l'acceptation de la connexion:", err)
			continue
		}
		fmt.Println("Nouvelle connexion établie.")
	}
}