const chalk = require('chalk');
const Table = require('cli-table3');
const readline = require('readline');
const { exec } = require('child_process');
const fs = require('fs');
const path = 'coups.txt';
fs.writeFile(path,"", (err) => {
  if (err) {
    console.error('Erreur lors de la création du fichier :', err);
    return;
  }
});

const lettresPossibles = 'A'.repeat(14) + 'B'.repeat(4) + 'C'.repeat(7) + 'D'.repeat(5) + 'E'.repeat(19) + 'F'.repeat(2) + 'G'.repeat(4) + 'H'.repeat(2) + 'I'.repeat(11) + 'J'.repeat(1) + 'K'.repeat(1) + 'L'.repeat(6) + 'M'.repeat(5) + 'N'.repeat(9) + 'O'.repeat(8) + 'P'.repeat(4) + 'Q'.repeat(1) + 'R'.repeat(10) + 'S'.repeat(7) + 'T'.repeat(9) + 'U'.repeat(8) + 'V'.repeat(2) + 'W'.repeat(1) + 'X'.repeat(1) + 'Y'.repeat(1) + 'Z'.repeat(2);
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

let lettresRestantes = lettresPossibles.split('');

function clearTerminal() {
  process.stdout.write('\x1Bc');
}

function afficherGrille(joueur,lettrej1, lettrej2) {
  // Inverser les lettres si le joueur est le joueur 2
  if (joueur === 2) {
    let temp = lettrej1;
    lettrej1 = lettrej2;
    lettrej2 = temp;
  }

  
  const data = fs.readFileSync('./coups.txt', 'utf8');
  const lignes = data.split('\n');
  const points = ['  ', '  ', '9 ', 16, 25, 36, 49, 64, 81];
  
  let table = new Table();
  let letters = [];

  for (let i = 0; i < lettrej1.length; i++) {
    letters.push(chalk.red(lettrej1[i]));
  }
  table.push(letters);
  console.log(chalk.green(`Lettres du joueur 1 :\n${table.toString()}`));

  let grille = new Table({
    head: [''].concat(points.map((point, i) => chalk.blue(point))),
    style: { 'padding-left': 1, 'padding-right': 1 }
  });

  let grilleData = [];
  for (let i = 0; i < 8; i++) {
    grilleData.push(Array.from({ length: 9 }, () => ' '));
  }

  let ligneIndex = 0;
  for (let ligne of lignes) {
    if (ligne.includes(`Joueur 1`)) {
      let mot = ligne.split(' : ')[1];
      for (let i = 0; i < mot.length; i++) {
        grilleData[ligneIndex][i] = mot[i];
      }
      ligneIndex++;
    }
  }
  for (let i = 0; i < 8; i++) {
    grille.push([chalk.blue(i + 1)].concat(grilleData[i]));
  }
  console.log(chalk.green(`Grille du joueur 1 :\n${grille.toString()}`));


  let tableau = new Table();
  let letterss = [];

  for (let i = 0; i < lettrej2.length; i++) {
    letterss.push(chalk.red(lettrej2[i]));
  }
  tableau.push(letterss);
  console.log(chalk.green(`Lettres du joueur 2 :\n${tableau.toString()}`));


  let grille2 = new Table({
    head: [''].concat(points.map((point, i) => chalk.blue(point))),
    style: { 'padding-left': 1, 'padding-right': 1 }
  });
  // Initialiser la grille avec des espaces
  let grilleData2 = [];
  for (let i = 0; i < 8; i++) {
    grilleData2.push(Array.from({ length: 9 }, () => ' '));
  }

  let ligneIndex2 = 0;
  for (let ligne of lignes) {
    if (ligne.includes(`Joueur 2`)) {
      let mot = ligne.split(' : ')[1];
      for (let i = 0; i < mot.length; i++) {
        grilleData2[ligneIndex2][i] = mot[i];
      }
      ligneIndex2++;
    }
  }
  // Ajouter les données à la grille
  for (let i = 0; i < 8; i++) {
    grille2.push([chalk.blue(i + 1)].concat(grilleData2[i]));
  }
  console.log(chalk.green(`Grille du joueur 2 :\n${grille2.toString()}`));

}

function estMotValide(mot, lettresTirees) {
  // Vérifier si le mot est valide (chaque lettre est dans les lettres tirées)
  for (const lettre of mot) {
    if (!lettresTirees.includes(lettre)) {
      return false;
    }
  }
  return true;
}

function enregistrerCoup(joueur, mot) {
  // Enregistrer le coup dans le fichier coups.txt (sauf si le mot est 'passe')
  if (mot.toLowerCase() !== 'passe') {
    const coup = `Joueur ${joueur} a joué : ${mot}\n`;
    fs.appendFileSync(path, coup, 'utf8');
  }
}

function selectionnerLigneJoueur(joueur, numLigne) {
  // Sélectionner la ligne correspondant au joueur et au numéro de ligne spécifié
  const data = fs.readFileSync('./coups.txt', 'utf8');
  const lignes = data.split('\n');
  let compteur = 0;

  for (let ligne of lignes) {
    console.log(ligne);
    if (ligne.includes(`Joueur ${joueur}`)) {
      compteur++;
      if (compteur === numLigne) {
        return ligne;
      }
    }
  }
  return null;
}

async function proposerOptions(joueur, lettresTirees, lettresAdversaire,mot) {
  // Proposer des options au joueur (modifier un mot, jouer un coup, passer)
  clearTerminal();
  afficherGrille(joueur,lettresTirees, lettresAdversaire);
  return new Promise(resolve => {
    rl.question(`Joueur ${joueur}, que voulez-vous faire ?\n1. Modifier un mot\n2. Jouer un coup\n3. Passer\n4. Jarnac\n`, async (choix) => {
      if (choix === '2') { // Jouer un coup
        console.log(`Voici votre main : ${lettresTirees}`);
        const result = await jouerCoup(joueur, lettresTirees, rl,lettresAdversaire);
        resolve(result);
      } else if (choix === '3') { // Passer
        resolve({ mot: 'Passe', lettresTirees, lettresAdversaire });
      } else if (choix === '1') { // Modifier un mot
        rl.question('Entrez le numéro de ligne à modifier : ', async (numLigne) => {
          const ligne = selectionnerLigneJoueur(joueur, parseInt(numLigne));
          let mot = ligne.split(' : ')[1];
          console.log(`Votre mot actuel : ${mot}`);
          rl.question('Voulez-vous modifier l\'ordre des lettres dans le mot ? (Oui/Non) : ', async (reponse) => {
            if (reponse.toLowerCase() === 'oui') {
              rl.question(`Entrez les lettres du mot dans le nouvel ordre :`, async (nouvelOrdre) => {
                if (nouvelOrdre.length > 0 && estMotValide(nouvelOrdre, mot)) {
                  ancien_mot = mot;
                  mot = nouvelOrdre;
                  console.log(`Votre nouveau mot : ${mot}`);
                } else {
                  console.log('Lettres invalides. Le joueur continuera avec le mot actuel.');
                }
                console.log(`MOT : ${mot}`);
                const result = await ajouterLettresDuDeck3(joueur, lettresTirees, mot,ancien_mot);
                resolve(result);
              });
            } else {
              const result = await ajouterLettresDuDeck(joueur, lettresTirees, mot);
              resolve(result);
            }
          });          
        });
      } else if (choix === '4') { // Jarnac : échanger des lettres avec le deck adverse
        rl.question('Entrez le numéro de ligne à modifier : ', async (numLigne) => {
          const ligne = selectionnerLigneJoueur(joueur === 1 ? 2 : 1, parseInt(numLigne));
          let mot = ligne.split(' : ')[1];
          supprimerLigne(joueur === 1 ? 2 : 1, path, '');
          console.log(`Le mot de votre adversaire : ${mot}`);
          rl.question('Voulez-vous modifier l\'ordre des lettres dans le mot ? (Oui/Non) : ', async (reponse) => {
            if (reponse.toLowerCase() === 'oui') {
              rl.question(`Entrez les lettres du mot dans le nouvel ordre :`, async (nouvelOrdre) => {
                if (nouvelOrdre.length > 0 && estMotValide(nouvelOrdre, mot)) {
                  ancien_mot = mot;
                  mot = nouvelOrdre;
                  console.log(`Votre nouveau mot : ${mot}`);
                  const result = await ajouterLettresDuDeck2(joueur, lettresAdversaire,mot,ancien_mot);
                  resolve(result);
                } else {
                  console.log('Lettres invalides. Le joueur continuera avec le mot actuel.');
                  const result = await proposerOptions(joueur, lettresTirees, lettresAdversaire,mot);
                  resolve(result);
                }
              });
            } else {
              const result = await ajouterLettresDuDeck2(joueur, lettresAdversaire,mot,mot);
              resolve(result);
            }
          });
        });
      } else {
        console.log('Choix invalide. Le joueur continuera avec le mot actuel.');
        const result = await proposerOptions(joueur, lettresTirees, lettresAdversaire,mot);
        resolve(result);
      }
    });
  });

  async function ajouterLettresDuDeck(joueur, lettresTirees, mot) {
    // Ajouter des lettres du deck au mot actuel du joueur
    return new Promise(resolve => {
      rl.question('Entrez les lettres à ajouter à votre mot (issues de votre deck) : ', (lettresAjoutees) => {
        if (lettresAjoutees.length > 0 && estMotValide(lettresAjoutees, lettresTirees)) {
          for (const lettre of lettresAjoutees) {
            const index = lettresTirees.indexOf(lettre);
            if (index !== -1) {
              lettresTirees = lettresTirees.slice(0, index) + lettresTirees.slice(index + 1);
            } else {
              console.log(`Lettre '${lettre}' non disponible dans votre deck. Le joueur continuera avec le mot actuel.`);
              const result = proposerOptions(joueur, lettresTirees, lettresAdversaire,mot);
              resolve(result);
              return;
            }
          }
          let mot_nouveau = mot+lettresAjoutees; // On ajoute les lettres au mot actuel
          supprimerLigne(joueur, path, mot_nouveau,mot);
          console.log(`Votre nouveau mot : ${mot_nouveau}`);
          lettresTirees += tirerLettre();
          const result = proposerOptions(joueur, lettresTirees, lettresAdversaire,mot_nouveau);
          resolve(result);
        } else {
          console.log('Lettres invalides ou aucune lettre ajoutée. Le joueur continuera avec le mot actuel.');
          const result = proposerOptions(joueur, lettresTirees, lettresAdversaire,mot);
          resolve(result);
        }
      });
    });
  }

  async function ajouterLettresDuDeck3(joueur, lettresTirees, mot,ancien_mot) { // Ajouter des lettres du deck au mot actuel
    return new Promise(resolve => {
      rl.question('Entrez les lettres à ajouter à votre mot (issues de ton deck) : ', (lettresAjoutees) => {
        if (lettresAjoutees.length > 0 && estMotValide(lettresAjoutees, lettresTirees)) {
          for (const lettre of lettresAjoutees) {
            const index = lettresTirees.indexOf(lettre);
            if (index !== -1) {
              lettresTirees = lettresTirees.slice(0, index) + lettresTirees.slice(index + 1); // On retire les lettres ajoutées du deck
            } else {
              console.log(`Lettre '${lettre}' non disponible dans votre deck. Le joueur continuera avec le mot actuel.`);
              const result = proposerOptions(joueur, lettresTirees, lettresAdversaire,mot);
              resolve(result);
              return;
            }
          }
          mot+=lettresAjoutees;
          supprimerLigne(joueur, path, mot,ancien_mot);
          console.log(`Votre nouveau mot : ${mot}`);
          lettresTirees += tirerLettre();
          const result = proposerOptions(joueur, lettresTirees, lettresAdversaire,mot);
          resolve(result);
        } else {
          console.log('Lettres invalides ou aucune lettre ajoutée. Le joueur continuera avec le mot actuel.');
          const result = proposerOptions(joueur, lettresTirees, lettresAdversaire,mot);
          resolve(result);
        }
      });
    });
  }


  async function ajouterLettresDuDeck2(joueur, lettresAdversaire, mot,ancien_mot) {
    // Ajouter des lettres du deck adverse au mot actuel du joueur
    return new Promise(resolve => {
      rl.question('Entrez les lettres à ajouter à votre mot (issues du deck adverse) : ', (lettresAjoutees) => {
        if (lettresAjoutees.length > 0 && estMotValide(lettresAjoutees, lettresAdversaire)) {
          for (const lettre of lettresAjoutees) {
            const index = lettresAdversaire.indexOf(lettre);
            if (index !== -1) {
              lettresAdversaire = lettresAdversaire.slice(0, index) + lettresAdversaire.slice(index + 1);
            } else {
              console.log(`Lettre '${lettre}' non disponible dans votre deck. Le joueur continuera avec le mot actuel.`);
              const result = proposerOptions(joueur, lettresTirees, lettresAdversaire,mot);
              resolve(result);
              return;
            }
          }
          mot+=lettresAjoutees;
          supprimerLigne(joueur, path, mot,ancien_mot);
          console.log(`Votre nouveau mot : ${mot}`);
          lettresTirees += tirerLettre();
          const result = proposerOptions(joueur, lettresTirees, lettresAdversaire,mot);
          resolve(result);
        } else {
          console.log('Lettres invalides ou aucune lettre ajoutée. Le joueur continuera avec le mot actuel.');
          const result = proposerOptions(joueur, lettresTirees, lettresAdversaire,mot);
          resolve(result);
        }
      });
    });
  }
}


function supprimerLigne(joueur, filePath, mot, ancien_mot) {
  // Supprimer une ligne du fichier coups.txt (pour la remplacer par une nouvelle ligne)
  const contenu = fs.readFileSync(filePath, 'utf8').split('\n');
  let nouveauContenu = '';

  for (const ligne of contenu) {
    if (ligne.includes(ancien_mot)) {
      nouveauContenu += `Joueur ${joueur} a joué : ` + mot;
    } else {
      nouveauContenu += ligne;
    }

    if (ligne !== contenu[contenu.length - 1]) {
      nouveauContenu += '\n';  // Ajoute un saut de ligne sauf pour la dernière ligne
    }
  }

  fs.writeFileSync(filePath, nouveauContenu);
}


async function jouerCoup(joueur, lettresTirees, rl,lettresAdversaire) {
  // Jouer un coup (saisir un mot)
  return new Promise(resolve => {
    rl.question(`Joueur ${joueur}, tirez vos lettres : ${lettresTirees}\nVotre mot (3 lettres minimum) : `, async (mot) => {
      if (mot.length >= 3 && estMotValide(mot, lettresTirees)) {
        for (const lettre of mot) {
          const index = lettresTirees.indexOf(lettre);
          if (index !== -1) {
            lettresTirees = lettresTirees.slice(0, index) + lettresTirees.slice(index + 1);
          }
        }
        enregistrerCoup(joueur, mot);
        lettresTirees += tirerLettre();
        const result = await proposerOptions(joueur, lettresTirees, lettresAdversaire,mot);
        resolve(result);
      } else {
        console.log("Le mot n'est pas valide. Veuillez réessayer.");
        const result = await jouerCoup(joueur, lettresTirees, rl);
        resolve(result);
      }
    });
  });
}

async function jouer() {
  let lettresJoueur1 = tirerLettres();
  let lettresJoueur2 = tirerLettres();


  let compteurJoueur1 = 0;
  let compteurJoueur2 = 0;


  let tour = 1;
  while (compteurJoueur1 < 9 && compteurJoueur2 < 9) { // Tant que les deux joueurs n'ont pas joué 9 mots

    clearTerminal();
    if (tour !== 1) {
      lettresJoueur1 += tirerLettre();
    }
    const resultJoueur1 = await proposerOptions(1, lettresJoueur1,lettresJoueur2);
    if (resultJoueur1) {
      lettresJoueur1 = resultJoueur1.lettresTirees;
      lettresJoueur2 = resultJoueur1.lettresAdversaire;
    } else {
      console.error("Erreur lors de la récupération des options du joueur 1.");
      return;
    }

    console.log(`Lettres restantes pour le joueur 1 : ${lettresJoueur1}`);
    if (tour !== 1) {
      lettresJoueur2 += tirerLettre();
    }
    clearTerminal();
    const resultJoueur2 = await proposerOptions(2, lettresJoueur2,lettresJoueur1);
    if (resultJoueur2) {
      lettresJoueur2 = resultJoueur2.lettresTirees;
      lettresJoueur1 = resultJoueur2.lettresAdversaire;
    } else {
      console.error("Erreur lors de la récupération des options du joueur 2.");
      return;
    }

    console.log(`Lettres restantes pour le joueur 2 : ${lettresJoueur2}`);
    compteurJoueur1 = 0;
    compteurJoueur2 = 0;
    let data = fs.readFileSync('./coups.txt', 'utf8');
    let lignes = data.split('\n');
    for (let ligne of lignes) {
    if (ligne.includes('Joueur 1')) {
      compteurJoueur1++;
    } else if (ligne.includes('Joueur 2')) {
      compteurJoueur2++;
    }
  }
  tour++;
  } // Fin de la boucle while
  let pts1=compterPoints(1);
  let pts2=compterPoints(2);
  if (pts1>pts2) { // Afficher le gagnant de la partie
    console.log(`Le joueur 1 a gagné avec ${pts1} points contre ${pts2} points pour le joueur 2.`);
  }
  else if (pts1<pts2) {
    console.log(`Le joueur 2 a gagné avec ${pts2} points contre ${pts1} points pour le joueur 1.`);
  }
  console.log('Fin de la partie.');
  rl.close();
}

function tirerLettre() {
  const randomIndex = Math.floor(Math.random() * lettresRestantes.length);
  return lettresRestantes.splice(randomIndex, 1)[0];
}

function tirerLettres() {
  let lettresTirees = '';
  for (let i = 0; i < 6; i++) {
    lettresTirees += tirerLettre();
  }
  return lettresTirees;
}

//fonction qui prend comme argument le numéro de joueur qui compte le nombre de lettres à chaque ligne pour un mot. Si le mot fait 3 lettres on obtient 9 points, s'il fait 4 lettres on obtient 16 points, s'il fait 5 lettres on obtient 25 points, s'il fait 6 lettres on obtient 36 points, s'il fait 7 lettres on obtient 49 points, s'il fait 8 lettres on obtient 64 points, s'il fait 9 lettres on obtient 81 points.
function compterPoints(joueur) {
  const data = fs.readFileSync('./coups.txt', 'utf8');
  const lignes = data.split('\n');
  let points = 0;
  for (let ligne of lignes) {
    if (ligne.includes(`Joueur ${joueur}`)) {
      let mot = ligne.split(' : ')[1];
      if (mot.length === 3) {
        points += 9;
      } else if (mot.length === 4) {
        points += 16;
      } else if (mot.length === 5) {
        points += 25;
      } else if (mot.length === 6) {
        points += 36;
      } else if (mot.length === 7) {
        points += 49;
      } else if (mot.length === 8) {
        points += 64;
      } else if (mot.length === 9) {
        points += 81;
      }
    }
  }
  return points;
}

// Exécuter le jeu
jouer();
