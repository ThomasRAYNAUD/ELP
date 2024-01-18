package main

import (
	"fmt"
	"net"
	"sync"
)

var wg sync.WaitGroup


func handleConnection(conn net.Conn) {
	defer conn.Close()
}

func main() {
	port:="20001"
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("Erreur lors de l'écoute sur le port %s, le message d'erreur --> %v\n", port, err)
		return
	}
	fmt.Printf("Serveur en attente de connexions sur le port %s...\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erreur lors de l'acceptation de la connexion:", err)
			continue
		}
		fmt.Println("Nouvelle connexion établie.")
		go handleConnection(conn)
	}
}

