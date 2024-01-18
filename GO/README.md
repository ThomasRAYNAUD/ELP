 <h2>**InsightNet Scanner** est un scanneur de ports écrit en Go qui vous permet d'énumérer les ports ouverts en TCP sur un hôte distant.</h2>
## Features

- Scan de port TCP en utilisant des goroutines --> utilisation de paramètres afin d'en tirer le meilleur temps
- Choix du nombre de **workers**
- Choix de la **plage de ports** ou du **port**
- Choix du nombre de **port par plage**
- Un environnement de test est aussi disponible : serveur web (web_server.go et index.html) et ouvrir des port TCP (open.go)


## Installation

Utiliser le package [git]([https://pip.pypa.io/en/stable/](https://git-scm.com/book/fr/v2/D%C3%A9marrage-rapide-Installation-de-Git)) pour installer le scanner.

```bash
git clone https://github.com/ThomasRAYNAUD/ELP.git 
```

## Usage

```python
Usage:
go run ./scan.go [cible] [flags]

cible :
# choix de la machine à scanner
<IPv4 adresse> ou <nom de domaine>

flags :
# choix de la plage de port ou un seul port
-p [port de début(int):(int)port de fin] ou -p [port à scanner (int)]

# choix du nombre de worker
-w [nbr de workers (int) ]

# choix du nombre de ports par plage
-n [nbr de ports (int) ]

```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
