package main

import (
	"fmt"
	"math"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

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

func quickSort(output []result) []result {
	if len(output) <= 1 {
		return output //si y'a qu'un terme -> tableau  déjà trié
	}

	pivotIndex := len(output) / 2 // index du milieu du tableau
	pivot := output[pivotIndex]   // valeur du milieu du tableau

	var left []result // on coup le tableau en deux
	var right []result

	for i := 0; i < len(output); i++ {
		if i == pivotIndex {
			continue
		}
		if output[i].NumberPort <= pivot.NumberPort { // trie ...
			left = append(left, output[i])
		} else {
			right = append(right, output[i])
		}
	}

	// recursivité pour trier les partitions gauche et droite jusqu'à taille de 1
	gauche := quickSort(left)
	droit := quickSort(right)

	// fusion des résultats triés
	tableau_trie := append(append(gauche, pivot), droit...)
	return tableau_trie
}

func isIPv4Address(ip string) bool {
	// Utiliser la fonction ParseIP du package "net" pour vérifier le format
	return net.ParseIP(ip) != nil
}

func isDomainName(domain string) bool {
	// Expression régulière pour vérifier le format d'un nom de domaine simple.
	// Cette expression régulière peut nécessiter des ajustements en fonction des exigences spécifiques.
	domainRegex := regexp.MustCompile(`^([a-zA-Z0-9]+\.)*[a-zA-Z0-9]+\.[a-zA-Z]{2,}$`)
	return domainRegex.MatchString(domain)
}

func estNombreNombre(chaine string) (int, int, bool) {
	// Diviser la chaîne en deux parties en fonction du délimiteur ":"
	parts := strings.Split(chaine, ":")
	if len(parts) != 2 {
		return 0, 0, false
	}

	premierNombre, err1 := strconv.Atoi(parts[0]) // si conversion en entier pas faisable car char -> une erreur dans err1
	deuxiemeNombre, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil {
		return 0, 0, false
	}
	return premierNombre, deuxiemeNombre, true
}

func estChiffre(chaine string) (int, bool) {
	p, err := strconv.Atoi(chaine)
	if err != nil {
		return p, false
	}
	return p, true
}
func loadingAnimation() {
	for {
		for _, char := range `-\|/` {
			fmt.Printf("\r Scan en cours  %c", char)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func worker(jobs <-chan travail, results chan<- result, address string) {
	defer wg.Done() // Décrémentation du WaitGroup à la fin de l'exécution de la goroutine.
	// Boucle qui tourne dans que y'a des jobs pour recevoir des plages de ports à scanner depuis le canal jobs.
	for work := range jobs {
		// Initialisation du compteur de ports TCP ouverts.
		portsTCP := 0
		// Boucle pour scanner les ports dans la plage donnée.
		for port := work.BegPort; port <= work.EndPort; port++ {
			// Construction de l'adresse du site à scanner en utilisant l'adresse IP locale "127.0.0.1" et le port.
			site := fmt.Sprintf(address+":%d", port)
			// Tentative de connexion TCP avec un délai de timeout de 1 seconde.
			connTCP, errTCP := net.DialTimeout("tcp", site, 1*time.Second) // 1 secondes si firewall ou que port fermé. Plus long est overkill
			// Vérification s'il n'y a pas d'erreur lors de la connexion, indiquant que le port est ouvert.
			if errTCP == nil {
				// Fermeture de la connexion TCP après confirmation du port ouvert.
				connTCP.Close()

				// Incrémentation du compteur de ports ouverts.
				portsTCP++
				// Affichage d'un message indiquant que le port est ouvert.
				results <- result{true, port, "TCP"}
			} /*else {
				results <- result{false, port, "TCP"} // ---> cette partie nous permet de renvoyer les ports fermés, qui nous intéressent pas
			} 
			*/ 
		}
	}
}

func main() {

	//initialisation des varibales
	var view1 bool = true
	var view2 bool = true
	var view3 bool = true
	var view3_1 bool = false
	var arg_p bool = false
	var arg_w bool = false
	var arg_n bool = false
	var BegPort int
	var BegPortbis int
	var EndPort int
	var verif1 bool
	var verif2 bool
	var numWorkers int
	var verif3 bool = false
	var nbrPort int
	var verif4 bool
	var address string
	var domain bool = false

	for _, arg := range os.Args[1:] {
		if arg == "-p" && view1 { //view1 permet de passer une seule fois dans le -p
			arg_p = true
			view1 = false
		} else if arg_p {
			BegPort, EndPort, verif1 = estNombreNombre(arg)
			BegPortbis, verif2 = estChiffre(arg)
			if verif1 ==false&& verif2==false{
				fmt.Println("Aucun argument valable de port")
				os.Exit(0)
			}
			if verif1 {
				//premier format alors réalise le traitement dessus : plage de port
				if BegPort > EndPort {
					fmt.Println("Erreur : Le port de début et supérieur au port de fin.")
					os.Exit(0)
				} else if BegPort == EndPort { //pour scanner un seul port le format "-p 80:80" n'est pas valable -> faire "-p 80"
					fmt.Println("Erreur : Les ports de début et de fin sont égaux.")
					os.Exit(0)
				} else if BegPort < 0 || EndPort > 65535 {
					fmt.Println("L'un de vos port n'est pas valide (entre 0 et 65535)")
					os.Exit(0)
				}
			} else if verif2 {
				//second format 1 seul port alors réalise le traitement dessus (met sous le format port1:port1 pour la suite du traitemet)
				BegPort = BegPortbis
				EndPort = BegPortbis
			} else {
				fmt.Println("Pas de port ou plage de ports spécifiée")
				os.Exit(0)
			}
			arg_p = false
			//nombre de workers dans le programme !
		} else if arg == "-w" && view2 {
			view2 = false
			arg_w = true
		} else if arg_w {
			numWorkers, verif3 = estChiffre(arg)
			if verif3 == false {
				fmt.Println("Le nombre de worker n'est pas compatible (entier uniquement)")
				os.Exit(0)
			}
			arg_w = false
		} else if arg == "-n" && view3 {
			view3 = false
			view3_1 = true
			arg_n = true
		} else if arg_n {
			nbrPort, verif4 = estChiffre(arg)
			if verif4 == false {
				fmt.Println("Le nombre de ports n'est pas valide")
				os.Exit(0)
			}
			arg_n = false
		} else if isIPv4Address(arg) || isDomainName(arg) {
			domain = true
			address = arg
		} else {
			fmt.Println("Le format spécifié de la commande n'est pas supporté")
			os.Exit(0)
		}
	}
	//cas par défaut si un argument n'est pas utilisé
	if (verif1 == false && verif2 == false) || view1 { //si la commande termine par -p -> une valeur par défaut par contre si le -p se trouve pas en millieu de commande et que ce qui suit n'est pas un nbr = erreur
		BegPort = 0
		EndPort = 65535
	}
	if view2 { // si pas choisi de worker
		numWorkers = 16
	}
	if view3 {
		nbrPort = 100
	}
	if view3_1 {
		if (EndPort - BegPort + 2) <= nbrPort {
			fmt.Println("Le nombre de port spécifié n'est pas possible (plus de port à scanner que de ports disponibles sur la plage)")
			os.Exit(0)
		}
	}
	if domain == false {
		fmt.Println("Pas de nom de domaine ou d'adresse IP spécifié en argument de commande.")
		os.Exit(0)
	}

	a := time.Now() //début du chrono
	var work travail
	work.BegPort = BegPort
	work.EndPort = EndPort
	//nbr de ports par plage -> nbrPort déjà def
	nbrPlage := int(math.Ceil(float64(work.EndPort-work.BegPort+1) / float64(nbrPort))) // Calcul du nombre de plages. Ceil pour arrondir à l'entier supérieur
	numJobs := nbrPlage
	jobs := make(chan travail, numJobs)      //jobs est un channel de type travail
	results := make(chan result, numWorkers) //pareil pour results
	fmt.Println("heure du scan : ", time.Now())
	fmt.Println("cible : ", address)
	if BegPort == EndPort {
		fmt.Println("Port scanné :", EndPort)
	} else {
		fmt.Sprintf("Plage de ports scannés : %d - %d\n", BegPort, EndPort)
	}
	fmt.Println("nombre de workers : ", numWorkers)
	fmt.Println("nombre de port par plage =", nbrPort)
	//go loadingAnimation() // --> annimation mais ajoute une go routine donc dans une problématique de temps à éviter d'ajouter
	for w := 1; w <= numWorkers; w++ { //nombre de workers qui effectue le scan (meilleur optimisation = 2* nombres coeurs)
		wg.Add(1)
		go worker(jobs, results, address) //ouverture de goroutine
	}
	for i := 0; i < numJobs; i++ { // on divise en petite plage pour transmettre aux workers -> création des jobs pendant que les workers travaillent dessus
		startPort := work.BegPort + i*(nbrPort) //permet de commencer à 31 si la première plage c'est fini à 30 (le +1)
		endPort := startPort + nbrPort - 1
		if endPort > work.EndPort { // si dernière plage n'a plus assez de ports on défini le max que l'on peut
			endPort = work.EndPort
		}
		jobs <- travail{BegPort: startPort, EndPort: endPort} //on envoie dans le channel jobs les structures de travail à faire par les workers
	}
	close(jobs)    //fermeture du canal quand tous les jobs ont été créés
	wg.Wait()      //on attend que les goroutines finissent
	close(results) //fermeture du canal quand les goroutines ont fini de travailler car elles ajoutent plus rien dedans 
	b := time.Now()
	temps := b.Sub(a)
	fmt.Printf("-------------------------\n")
	fmt.Printf("temps écoulé : %s\n", temps)
	fmt.Printf("-------------------------\n")
	fmt.Printf("Scan fini : tri en cours...\n")
	// affichage du channel result
	output := make([]result, 0) // creer tab output pour le quicksort
	for r := range results {
		output = append(output, r)
	}

	tableau_trie := quickSort(output)
	for _, r := range tableau_trie { //on met pas d'itérateur car variable de l'itérateur sera inutilisé
		state := r.State
		port := r.NumberPort
		protocol := r.Protocol
		if state == true {
			fmt.Printf("Port %d : ouvert -> protocole : %s\n", port, protocol)
		} else {
			fmt.Printf("Port %d : fermé -> protocole : %s\n", port, protocol)
		}
	}
}
