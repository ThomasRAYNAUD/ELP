# Jarnac

Jarnac est un jeu de société français qui combine des éléments du scrabble et du jeu de cartes, où les joueurs créent des mots à partir de lettres tirées aléatoirement pour marquer des points.

## Lancement du jeu

Ce projet est codé intégralement en JavaScript, utilisant Nodejs. 
Afin de lancer le projet, il faut : 
```bash
sudo apt install nodejs    # installer nodejs si ce n'est pas encore fait
git clone https://github.com/ThomasRAYNAUD/ELP.git       # clone le dépôt
cd ELP/JS\ -Jarnac
node Jarnac.js
```
S'il y a une erreur sur l'importation des modules, il faut installer : 
```bash
npm install chalk
npm install cli-table3
```

## Format de fichier utilisé

Notre format de fichier enregistre les grilles et fait les modifications sur ce fichier appelé ***coups.txt***.
Ce dernier se créé automatiquement, s'il en existe déjà un, il sera supprimé en début de partie.
Dans le fichier on retrouvera par exemple : 
``` bash
cat coups.txt
Joueur 1 a joué : HELLO
Joueur 1 a joué : TOTO
Joueur 2 a joué : PLAGE
```
