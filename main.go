package main

import (
	"fmt"
	"math"
	"net"
	"sync"
	"time"
)

type travail struct {
	BegPort  int
	EndPort  int
	PortsTCP int // Ajout du champ PortsTCP pour compter le nombre de ports TCP ouverts
}
type result struct {
	State      bool
	NumberPort int
	Protocol   string
}

var wg sync.WaitGroup

// worker est une fonction qui représente le travail effectué par chaque goroutine lors du scan des ports TCP.
// Elle prend un canal pour recevoir les plages de ports à scanner (jobs) et un canal pour envoyer les résultats (results).
func worker(jobs <-chan travail, results chan<- result) {
	defer wg.Done() // Décrémentation du WaitGroup à la fin de l'exécution de la goroutine.

	// Boucle infinie pour recevoir des plages de ports à scanner depuis le canal jobs.
	for work := range jobs {
		// Initialisation du compteur de ports TCP ouverts.
		portsTCP := 0
		// Boucle pour scanner les ports dans la plage donnée.
		for port := work.BegPort; port <= work.EndPort; port++ {
			// Construction de l'adresse du site à scanner en utilisant l'adresse IP locale "127.0.0.1" et le port.
			site := fmt.Sprintf("192.168.1.169:%d", port)

			// Tentative de connexion TCP avec un délai de timeout de 1 seconde.
			connTCP, errTCP := net.DialTimeout("tcp", site, 1*time.Second)

			// Vérification s'il n'y a pas d'erreur lors de la connexion, indiquant que le port est ouvert.
			if errTCP == nil {
				// Fermeture de la connexion TCP après confirmation du port ouvert.
				connTCP.Close()

				// Incrémentation du compteur de ports ouverts.
				portsTCP++
				// Affichage d'un message indiquant que le port est ouvert.
				results <- result{true, port, "TCP"}
			} else {
				results <- result{false, port, "TCP"}
			} // Envoi des résultats au canal results.
		}
	}
}
func main() {
	a := time.Now() //début du chrono
	var work travail
	work.BegPort = 10                                                                   // port du début
	work.EndPort = 50                                                                   // port de fin
	nbrPort := 10                                                                       // nombre de port par plage
	nbrPlage := int(math.Ceil(float64(work.EndPort-work.BegPort+1) / float64(nbrPort))) // Calcul du nombre de plages
	numJobs := nbrPlage
	jobs := make(chan travail, numJobs)      //jobs est un channel de type travail
	results := make(chan result, 10*numJobs) //pareil pour results
	for w := 1; w <= numJobs; w++ {          //nombre de workers qui effectue le scan (VALEUR MAX A FAIRE VARIER POUR TROUVER LE PLUS OPTI)
		wg.Add(1)
		go worker(jobs, results) //ouverture de goroutine
	}
	for i := 0; i < numJobs; i++ { // on divise en petite plage pour transmettre aux workers
		startPort := work.BegPort + i*(nbrPort)
		endPort := startPort + nbrPort - 1
		if endPort > work.EndPort {
			endPort = work.EndPort
		}
		jobs <- travail{BegPort: startPort, EndPort: endPort} //on envoie dans le channel jobs les structures travaillent à faire par les workers
	}
	close(jobs)    //fermeture du canal
	wg.Wait()      //on attend que les goroutines finissent
	close(results) //fermeture du canal
	b := time.Now()
	temps := b.Sub(a) // fin chrono
	fmt.Printf("-------------------------\n")
	fmt.Printf("temps écoulé : %s\n", temps)
	fmt.Printf("-------------------------\n")

	// affichage du channel résultat
	for result := range results {
		state := result.State
		port := result.NumberPort
		protocol := result.Protocol
		if state {
			fmt.Printf("Le port %d est %t, protocole: %s\n", port, state, protocol)
		}
	}
}

//LE NOMBRE DE PLAGE A FAIRE = NOMBRE DE JOBS ???
/*
func ParamScan(i int,j int){
    nombrePlage := j / i
    Jobs = nombrePlage
    for a := 0; a <= nombrePlage; a++ {
        wg.Add(1)
        go AnalysePort(a*i+1, i*(a+1))
    }
}
*/
//IL FAUT QUE LA FONCTION PRENNE EN ENTREE UN NOMBRE DE WORKER ET LES JOBS A TRAITER. ELLE DOIT PRENEDRE LES JOBS ET LES DONNER AUX WORKERS, UTILISER LE TYPE TRAVAIL DEFINI

// _=var variable pas utilisé
