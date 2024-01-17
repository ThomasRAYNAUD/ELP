package main

import (
	"fmt"
	"strconv"
	"strings"
	"os"
	"regexp"
	"net"
	"math"
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




//pas que ipv4 ???
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
		return 0,0,false
	}

	premierNombre, err1 := strconv.Atoi(parts[0])
	deuxiemeNombre, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil {
		return 0, 0, false
	}

	return premierNombre, deuxiemeNombre, true
}

func estChiffre(chaine string) (int,bool) {
	p, err := strconv.Atoi(chaine)
	return p,err == nil
}

func worker(jobs <-chan travail, results chan<- result, address string) {
	defer wg.Done() // Décrémentation du WaitGroup à la fin de l'exécution de la goroutine.
	// Boucle infinie pour recevoir des plages de ports à scanner depuis le canal jobs.
	for work := range jobs {
		// Initialisation du compteur de ports TCP ouverts.
		portsTCP := 0
		// Boucle pour scanner les ports dans la plage donnée.
		for port := work.BegPort; port <= work.EndPort; port++ {
			// Construction de l'adresse du site à scanner en utilisant l'adresse IP locale "127.0.0.1" et le port.
			site := fmt.Sprintf(address+":%d",port)
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
				fmt.Printf("port %d fermé\n", port)
			} // Envoi des résultats au canal results.
		}
	}
}

var wg sync.WaitGroup
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
	
	
	for i, arg := range os.Args[1:] {
		_=i //éviter les pb de variable pas utilisée
		if arg == "-p" && view1 {
			arg_p=true
			_=arg_p
			view1=false
		} else if arg_p {
			BegPort,EndPort,verif1=estNombreNombre(arg)
			BegPortbis,verif2=estChiffre(arg)
			if verif1 {
				//premier format alors réalise le traitement dessus
				if BegPort>EndPort {
					fmt.Println("Erreur : Le port de début et supérieur au port de fin.")
					os.Exit(0)
				} else if BegPort==EndPort {
					fmt.Println("Erreur : Les ports de début et de fin sont égaux.")
					os.Exit(0)
				} else if BegPort<0 || EndPort>65535 {
					fmt.Println("L'un de vos port n'est pas valide (entre 0 et 65535)")
					os.Exit(0)
				}
			} else if verif2 {
				//second format 1 seul port alors réalise le traitement dessus
				BegPort=BegPortbis
				EndPort=BegPortbis
				fmt.Println("Second format")
			}
			arg_p=false
		//nombre de workers dans le programme !
		} else if arg == "-w" && view2 {
			view2=false
			arg_w=true
			_=arg_w
		} else if arg_w {
			numWorkers,verif3 = estChiffre(arg)
			_=numWorkers
			if verif3==false {
				fmt.Println("Le nombre de worker n'est pas compatible (entier uniquement)")
				os.Exit(0)
			}
			arg_w=false
		} else if arg == "-n" && view3 {
			view3=false
			view3_1=true
			arg_n=true
			_=arg_n
		} else if arg_n {
			nbrPort,verif4 = estChiffre(arg)
			_=nbrPort
			if verif4==false {
				fmt.Println("Le nombre de ports n'est pas valide")
				os.Exit(0)
			}
			arg_n=false
		} else if isIPv4Address(arg) || isDomainName(arg) {
			domain=true
			address=arg
		} else {
			fmt.Println("Le format spécifié de la commande n'est pas supporté")
			os.Exit(0)
		}
	}
//cas où aucun port n'est définit
	if (verif1==false && verif2==false) || view1 {
		BegPort=0
		EndPort=65535
		verif1=true
		verif2=true
		view1=false
	}
	if view2 { // si pas choisi de worker
		numWorkers=2
		view2=false
		// mettre une valeur par défault de worker
	} 
	if view3 {
		nbrPort=1000
		view3=false
	} 
	if view3_1 {
		if nbrPort > 0 && nbrPort <= 65535 {
			if (EndPort-BegPort+2)<=nbrPort {
				fmt.Println("Le nombre de port spécifié n'est pas possible (plus de port à scanner que de ports disponibles sur la plage)")
				os.Exit(0)
			}
		} else {
			fmt.Println("nombre de port trop grand ou trop petit")
			os.Exit(0)
		}
		view3_1=false
	} 
	if domain==false {
		fmt.Println("Pas de nom de domaine ou d'adresse IP spécifié en argument de commande.")
		os.Exit(0)
	}
	
	a := time.Now() //début du chrono
	var work travail
	work.BegPort = BegPort
	work.EndPort = EndPort
	//nbr de ports par plage -> nbrPort déjà def
	nbrPlage := int(math.Ceil(float64(work.EndPort-work.BegPort+1) / float64(nbrPort))) // Calcul du nombre de plages
	numJobs := nbrPlage
	jobs := make(chan travail, numJobs)      //jobs est un channel de type travail
	results := make(chan result, numWorkers) //pareil pour results
	for w := 1; w <= numWorkers; w++ {       //nombre de workers qui effectue le scan (VALEUR MAX A FAIRE VARIER POUR TROUVER LE PLUS OPTI)
		wg.Add(1)
		go worker(jobs, results,address) //ouverture de goroutine
	}
	for i := 0; i < numJobs; i++ { // on divise en petite plage pour transmettre aux workers
		startPort := work.BegPort + i*(nbrPort)
		endPort := startPort + nbrPort - 1
		if endPort > work.EndPort {
			endPort = work.EndPort
		}
		jobs <- travail{BegPort: startPort, EndPort: endPort} //on envoie dans le channel jobs les structures travaillent à faire par les workers
	}
	close(jobs) //fermeture du canal
	go func() {
		wg.Wait()      //on attend que les goroutines finissent
		close(results) //fermeture du canal
		b := time.Now()
		temps := b.Sub(a) // fin chrono
		fmt.Printf("-------------------------\n")
		fmt.Printf("temps écoulé : %s\n", temps)
		fmt.Printf("-------------------------\n")

	}()

	// affichage du channel résultat
	output := make([]result, 0)
	for r := range results {
		output = append(output, r)
	}
	for _, r := range output { //on met pas d'itérateur car variable de l'itérateur sera inutilisé
		state := r.State
		port := r.NumberPort
		protocol := r.Protocol
		if state {
			fmt.Printf("Le port %d est %t, protocole: %s\n", port, state, protocol)
		}
	}
}