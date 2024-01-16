package main

import (
	"fmt"
	"net"
	"sync"
	"time"
	"strconv"
)

var wg sync.WaitGroup //crée une variable wg de type sync.WaitGroup, utilisée pour synchroniser l'exécution de goroutines en attendant leur achèvement.


func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Ajoutez le code pour gérer la connexion ici (lecture/écriture de données, etc.)
}


func startServer(port string, wg *sync.WaitGroup) {
	defer wg.Done() //chacune des goroutines s'exécute et appelle la méthode Done() lorsque la goroutine a terminé de s'exécuter
//le mot-clé defer s'exécutera à chaque fois qu'on quittera notre fonction même en cas de panique (plantage) de la fonction
	listener, err := net.Listen("tcp", ":"+port) //met le retour de la connection dans des variables
	//test si erreur est !vide --> si pas vide alors retoune le message associé dans la variable
	if err != nil {
		fmt.Printf("Erreur lors de l'écoute sur le port %s, le message d'erreur --> %v\n", port, err)
		return
	}

	defer listener.Close() //ferme la connexion réseau et permet aussi de libèrer les ressources associées

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

/*
deux fonctions : 
- handleConnection() : retourn un message au client ou autre --> a choisir
- startServer() -->  initialise la connexion et ouvertur de port sur le serveur

*/

func main() {
    debut := time.Now()
	ports := []string{}  // Déclaration de la tranche en type string 

/*
demande à l'user quel port de début, quel port de fin et chaque combien l'ouvrir

prot -> tcp ou udp uniquement
port de début
port de fin
pas de combien
*/



	for i := 10000; i <= 20000; i+=50 {
		ports = append(ports, strconv.Itoa(i)) //strconv.Itoa(i) pour convertir l'entier i en une chaîne de caractères avant de l'ajouter à la tranche, car la tranche est de type []string
	}

	for _, port := range ports { 
		wg.Add(1) //permet de définir le nombre de goroutines à attendre (on l'incrémente de 1 à chaque création de goroutine).
		go startServer(port, &wg)
	}
    fin := time.Now()
	fmt.Println(fin.Sub(debut))

	wg.Wait() //la méthode Wait() est utilisée pour empêcher l'exécution d'autres lignes de code jusqu'à ce que toutes les goroutines soient terminées
}

