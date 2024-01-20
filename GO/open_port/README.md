 <h2>Ce dossier continent un environnement de test permettant l'ouverture de ports</h2>
## Features

- open.go > script permettant d'ouvrir un port TCP spécifié
- web_server.go > script permettant de créer un serveur web en se basant sur la template index.html
- permet tester le bon fonctionnement de notre scanner

## Usage

```python
Usage open.go :
go run ./open.go
--> ouvrir un port spécifié en TCP. Le port par défaut est :20001.

Usage web_server.go :
go run ./web_server.go
--> permet d'ouvrir un serveur web sur un port spécifié dans le code, en utilisant
comme template "index.html". Le port par défaut est :9999

```
