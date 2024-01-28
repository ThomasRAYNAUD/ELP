# Jeu Elm

Ceci est une application Elm simple, l'objectif est de trouver le mot à partir d'une définition. Le mot est tiré au hasard dans une liste de plus de 800 mots.
Les définitions sont tirées depuis une API et le but est de trouver le maximum de mot en un temps donné (que l'on choisi), avant que la page freeze.

## Comment Jouer

1. Clonez le dépôt sur votre machine locale :

   ```bash
   git clone https://github.com/ThomasRAYNAUD/ELP.git
   cd ELP/ELM/project/src
   elm make Main.elm --output=elm.js     # compiler le code ELM en JavaScript. Attention à bien output en elm.js
   ```

2. Pour tester le code, il suffit de lancer un elm reactor :

Afin de tester le code dans l'environnement voulu, le code html contient une référence au code JS retourné par la compilation (il lance tout seul elm.js après le lancement du timer). Une fois le code compilé en JS, il suffit simplement de :

   ```bash
   elm reactor
   # puis aller dans le localhost qui nous est retourné
   ```

3. Une fois le localhost ouvert dans un navigateur :
- entrer dans le dossier /src.
- choisir le fichier index.html qui contient le code html, CSS et Javascript (obtenu par la compilation ELM).
- on importe aussi un fichier styles.css permettant de changer la forme des input en utilisant les class
