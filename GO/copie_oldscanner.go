package main

import (
	"fmt"
	"strconv"
	"strings"
	"os"
	"regexp"
)
/*
type travail struct {
	BegPort int
	EndPort int
}

type result struct {
	numeroPort  int
	TCP bool //si 1 alors ouvert si 0 alors fermé
	UDP bool //si 1 alors ouvert si 0 alors fermé
}

func worker(address string, jobs <-chan travail, results chan<- result) {
    for j := range jobs {
        for i := j.BegPort; i <= j.EndPort; i++ {
            site := fmt.Sprintf("%s:%d", address, i)
            connTCP, errTCP := net.DialTimeout("tcp", site, 2*time.Second)
            if errTCP == nil {
                // La connexion a réussi, donc le port est ouvert -> il faut le fermer
                connTCP.Close()
                fmt.Printf("Le port %d est ouvert pour l'adresse %s\n", i, address)
				results <- result{numeroPort: i, TCP: true, UDP: false}
			} else {
            // Si errTCP n'est pas nul, la connexion a échoué, et rien n'a besoin d'être fermé explicitement ici
			results <- result{numeroPort: i, TCP: false, UDP: false}
			}
		}
    }
}

func isValidDomain(domain string) bool {
	//  vérifie s'il a une structure de nom de domaine
	domainRegex := regexp.MustCompile(`^([a-zA-Z0-9]+\.)*[a-zA-Z0-9]+\.[a-zA-Z]{2,}$`)
	return domainRegex.MatchString(domain)
}

func isValidIP(ip string) bool {
	// Utiliser la fonction ParseIP du package "net" pour vérifier le format
	return net.ParseIP(ip) != nil
}
*/

func isIPv4Address(ip string) bool {
	parts := strings.Split(ip, ".")

	if len(parts) != 4 {
		return false
	}

	for _, part := range parts {
		if len(part) == 0 {
			return false
		}
		num := 0
		for _, digit := range part {
			if digit < '0' || digit > '9' {
				return false
			}
			num = num*10 + int(digit-'0')
		}
		if num < 0 || num > 255 {
			return false
		}
	}

	return true
}

func isDomainName(domain string) bool {
	// Expression régulière pour vérifier le format d'un nom de domaine simple.
	// Cette expression régulière peut nécessiter des ajustements en fonction des exigences spécifiques.
	domainRegex := regexp.MustCompile(`^[a-zA-Z0-9.-]+$`)

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


func main() {

	/*
	
	> Récupère le nombre d'argument
		- boucle for pour tester l'arg 1..2..3..4..5..6......
			} sinon si ( -p  ) {
				l'agument suivant correspond à la plage qui est une chaine du type début:fin
			} sinon si -a {
				l'agument suivant est adresse
			} sinon si -w {
				nombre de workers
			} sinon si -n {
				nbr de port par plage -> nbr de jobs
			}
	
	> mettre des valeurs par défaut

	> adresse est le else -> si ressemble pas alors sortir de la boucle

	> 
	
	*/

	//récupérer le nombre d'arguments
	//nbArg:=len(os.Args)-1
	//fmt.Println("Nombre d'arguments :",nbArg)
	
	//initialisation des views
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
	var workerVerif int
	var verif3 bool = false
	var portVerif int
	var verif4 bool
	var address string
	var domain bool = false
	
	
	// Afficher les valeurs des arguments

	for i, arg := range os.Args[1:] {
		_=i //éviter les pb de variable pas utilisée
		//fmt.Println("i : ",i)
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
				//fmt.Println("Le port de début : ", BegPort)
				//fmt.Println("Le port de fin : ", EndPort)
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
			workerVerif,verif3 = estChiffre(arg)
			_=workerVerif
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
			portVerif,verif4 = estChiffre(arg)
			_=portVerif
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
	for j:=0;j<=5;j++ {
		if (verif1==false && verif2==false) || view1 {
			BegPort=0
			EndPort=65535
			verif1=true
			verif2=true
			view1=false
		} else if view2 { // si pas choisi de worker
			workerVerif=2
			view2=false
			// mettre une valeur par défault de worker
		} else if view3 {
			portVerif=1000
			view3=false
		} else if view3_1 {
			if portVerif > 0 && portVerif <= 65535 {
				if (EndPort-BegPort+2)<=portVerif {
					fmt.Println("Le nombre de port spécifié n'est pas possible (plus de port à scanner que de ports disponibles sur la plage)")
					os.Exit(0)
				}
			} else {
				fmt.Println("nombre de port trop grand ou trop petit")
				os.Exit(0)
			}
			view3_1=false
		} else if domain==false {
			fmt.Println("Pas de nom de domaine ou d'adresse IP spécifié en argument de commande.")
			os.Exit(0)
		}
	}
	fmt.Println(address)




			//tester si bien dans le format num_port_début:num_port_fin
			//si oui alors extraire et mettre dans les variables
			//mettre en int et vérifier si pas <0 et 65500<
//		}
		 /*else if arg == "-w" && view2 == true {
			//mettre le nbr à la varible
			// autre tests ? pas trop grand ???
		} else if arg == "-n" && view3 ==true {
			//mettre le nombre et fairele calcule -> appel de la fonction de moha
		} else view4 { //dans ce cas une adresse
			//appeler la fonction de test d'une adresse
		}*/
		//fmt.Printf("%d: %s\n", i+1, arg)
//	}




/*
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Veuillez entrer un nom de domaine ou une adresse IP à scanner : ")
	address, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erreur de lecture de l'entrée:", err)
		return
	}
	address = strings.TrimSpace(address)
	if isValidDomain(address) {
		fmt.Println("Vous avez saisi un nom de domaine valide:", address)
	} else if isValidIP(address) {
		fmt.Println("Vous avez saisi une adresse IP valide:", address)
	} else {
		fmt.Println("Entrée non valide. Veuillez saisir un nom de domaine ou une adresse IP.")
	}

	var work int
	fmt.Print("Nombre de workers : ")
	fmt.Scanln(&work)
	if work<1 {
		fmt.Println("Erreur nbr de work < 0")
	}

	fmt.Print("Nombre de workers : ")
	var work int
	fmt.Scanln(&work)

	fmt.Print("Nombre de plages/jobs : ")
	var nbrRange int
	fmt.Scanln(&nbrRange)
//faut vérfier que pas négatif, il faut entier ...
    const numJobs = 10500

    jobs := make(chan travail, numJobs)
    results := make(chan result) //ne pas laisser numJobs afin de ne pas tricher
/*
	//va créer les workers qui vont prendre les jobs parallèlement
    for w := 0; w < work; w++ {
        go worker(address, jobs, results)
    }

	//def les jobs a créer
    for j := 65355; j <= numJobs; j++ { //a définir dans une fonction --> car on veut pas taille max car trop lourd
        jobs <- j
    }



	close(jobs)
// a changer
	for w := 0; w < work; w++ {
		<-results
	}
*/
}




/**

// worker est une fonction représente le travail effectué par chaque goroutine lors du scan des ports TCP
// Elle prend en argument un canal pour recevoir les ports à scanner (jobs)
// et un canal pour envoyer les résultats (results)
func worker(address string,jobs <-chan int, results chan<- int) {
	// Initialisation du compteur de ports TCP ouverts.
	portsTCP := 0
	// Boucle infinie pour recevoir des ports à scanner depuis le canal jobs.
	for j := range jobs { //pour un job que j'ai recu, je dois prendre BegPort et EndPort afin de tous les tester
		// Construction de l'adresse du site à scanner en utilisant l'adresse IP locale "127.0.0.1" et le port j.
		site := fmt.Sprintf("127.0.0.1:%d", j)

		// Tentative de connexion TCP avec un délai de timeout de 1 seconde.
		connTCP, errTCP := net.DialTimeout("tcp", site, 1*time.Second)

		// Vérification s'il n'y a pas d'erreur lors de la connexion, indiquant que le port est ouvert.
		if errTCP == nil {
			// Fermeture de la connexion TCP après confirmation du port ouvert.
			connTCP.Close()
			// Incrémentation du compteur de ports ouverts.
			portsTCP++
			// Affichage d'un message indiquant que le port est ouvert.
			fmt.Println("Le port est ouvert :", j)
		}
	}

	// La fonction retourne implicitement à la fin de la boucle for lorsque le canal jobs est fermé.
	// À ce stade, la goroutine se termine.
}

*/