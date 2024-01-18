package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	a := time.Now()
	BegPort := 1
	EndPort := 65535
	portsTCP := 0
	// Boucle pour scanner les ports dans la plage donnée.
	for port := BegPort; port <= EndPort; port++ {
		// Construction de l'adresse du site à scanner en utilisant l'adresse IP locale "127.0.0.1" et le port.
		site := fmt.Sprintf("127.0.0.1:%d", port)
		// Tentative de connexion TCP avec un délai de timeout de 1 seconde.
		connTCP, errTCP := net.DialTimeout("tcp", site, 1*time.Second) // 1 secondes si firewall ou que port fermé. Plus long est overkill
		// Vérification s'il n'y a pas d'erreur lors de la connexion, indiquant que le port est ouvert.
		if errTCP == nil {
			// Fermeture de la connexion TCP après confirmation du port ouvert.
			connTCP.Close()

			// Incrémentation du compteur de ports ouverts.
			portsTCP++
			// Affichage d'un message indiquant que le port est ouvert.
			fmt.Printf("port %d ouvert\n", port)
		} /*else {
			results <- result{false, port, "TCP"} // ---> cette partie nous permet de renvoyer les ports fermés, qui nous intéressent pas
		}
		*/
	}
	b := time.Now()
	temps := b.Sub(a)
	fmt.Printf("-------------------------\n")
	fmt.Printf("temps écoulé : %s\n", temps)
	fmt.Printf("-------------------------\n")

}
