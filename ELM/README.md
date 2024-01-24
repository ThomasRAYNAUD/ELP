# Jeu Elm

Ceci est une application Elm simple, l'objectif est de trouver le mot à partir d'une définition. Le mot est tiré au hasard dans une liste de plus de 800 mots.
Les définitions sont tirées depuis une API et le but est de trouver le maximum de mot en un temps donné (que l'on choisi), avant que la page freeze.

## Comment Jouer

1. Clonez le dépôt sur votre machine locale :

   ```bash
   git clone https://github.com/ThomasRAYNAUD/ELP.git

2. Pour tester le code, il suffit de lancer un elm reactor :

   ```bash
   elm reactor
   # puis aller dans le localhost qui nous est retourné

3. Une fois le localhost ouvert :
- entrer dans le dossier /src.
- choisir le fichier index.html.
- le fichier index.html contient le code elm sous forme de javascript que nous appelons dans un code javascript que nous avons écrit afin d'inclure la fonctionnalité du choix de timer.
- on importe aussi un fichier styles.css permettant de changer la forme des input en utilisant les class.
