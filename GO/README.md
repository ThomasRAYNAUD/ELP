 <h2>**InsightNet Scanner** est un scanneur de ports écrit en Go qui vous permet d'énumérer les ports ouverts en TCP sur un hôte distant.</h2>
## Features

- Scan de port TCP en utilisant des goroutines --> utilisation de paramètres afin d'en tirer le meilleur temps
- Choix du nombre de **workers**
- Choix de la **plage de ports** ou du **port**
- Choix du nombre de **port par plage**


## Installation

Use the package manager [pip](https://pip.pypa.io/en/stable/) to install foobar.

```bash
git clone https://github.com/ThomasRAYNAUD/ELP.git 
```

## Usage

```python
Usage:
go run ./scan.go [cible] [flags]

flags :
# choix de la plage de port
-w

# returns 'geese'
foobar.pluralize('goose')

# returns 'phenomenon'
foobar.singularize('phenomena')
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
