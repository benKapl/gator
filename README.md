# Gator - L'agrégateur de flux RSS en Go 🐊

## Introduction
Gator est un agrégateur de flux RSS en ligne de commande écrit en Go. Il permet aux utilisateurs de :
- Ajouter des flux RSS depuis internet pour les collecter
- Stocker les publications collectées dans une base de données PostgreSQL
- Suivre et ne plus suivre les flux RSS ajoutés par d'autres utilisateurs
- Afficher des résumés des publications agrégées dans le terminal, avec un lien vers l'article complet

Les flux RSS sont un excellent moyen de suivre vos blogs, sites d'actualités, podcasts préférés et bien plus encore. Gator vous aide à gérer et parcourir ces flux directement depuis votre terminal.

## Prérequis
Avant d'utiliser Gator, assurez-vous d'avoir installé :
- [Go](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)

## Installation
Pour installer Gator, exécutez :
```sh
go install github.com/benKapl/gator
```
Cela installera la commande `gator` globalement pour que vous puissiez l'utiliser depuis le terminal.

## Configuration
Gator nécessite un fichier de configuration pour se connecter à la base de données et gérer l'utilisateur actuel. Ce fichier de configuration doit être placé à la racine du répertoire utilisateur :

- **`~/.gator_config.json`** (Linux/macOS)
- **`%USERPROFILE%\.gator_config.json`** (Windows)
```json
{
  "db_url": "connection_string_goes_here",
  "current_user_name": ""
}
```
- Remplacez `connection_string_goes_here` par votre chaîne de connexion PostgreSQL.
- `current_user_name` est défini via les commandes CLI.

## Utilisation
Une fois installé et configuré, vous pouvez exécuter des commandes en utilisant la syntaxe suivante :
```sh
gator <commande> <options>
```
### Commandes
- **Enregistrer un nouvel utilisateur et se connecter :**
  ```sh
  gator register <nom_utilisateur>
  ```
  Cette commande ajoute un nouvel utilisateur à la base de données et le sélectionne.

- **Se connecter à un utilisateur existant :**
  ```sh
  gator login <nom_utilisateur>
  ```
  Cette commande change l'utilisateur actuel pour un utilisateur existant.

- **Lister tous les utilisateurs enregistrés :**
  ```sh
  gator users
  ```
  Affiche tous les utilisateurs enregistrés dans la base de données.

- **Ajouter un flux RSS et le suivre :**
  ```sh
  gator addfeed <nom> <url>
  ```
  Ajoute un flux RSS à la base de données et l'associe à l'utilisateur actuel.

- **Lister tous les flux disponibles :**
  ```sh
  gator feeds
  ```
  Affiche tous les flux RSS ajoutés au système.

- **Suivre un flux existant :**
  ```sh
  gator follow <url>
  ```
  Marque un flux existant comme suivi par l'utilisateur actuel.

- **Lister les flux suivis par l'utilisateur actuel :**
  ```sh
  gator follows
  ```
  Affiche tous les flux suivis par l'utilisateur actuel.

- **Agrégation des nouveaux articles à intervalle régulier :**
  ```sh
  gator agg <intervalle_temps>
  ```
  Récupère les nouvelles publications des flux suivis et les stocke dans la base de données à intervalles réguliers. Cette commande fonctionne en boucle infinie en arrière-plan.

- **Parcourir les publications récentes :**
  ```sh
  gator browse <limite>
  ```
  Récupère les publications les plus récentes des flux suivis par l'utilisateur actuel.

## Exemple de workflow
1. Installez Gator.
2. Créez une base de données PostgreSQL et mettez à jour `~/.gator_config.json` avec la chaîne de connexion.
3. Enregistrez un nouvel utilisateur :
   ```sh
   gator register sacha
   ```
4. Ajoutez un flux :
   ```sh
   gator addfeed technews https://example.com/rss
   ```
5. Lancez l'agrégation des nouveaux articles :
   ```sh
   gator agg 10m
   ```
6. Parcourez les articles récents :
   ```sh
   gator browse 5
   ```

## Remerciements
Un grand merci à **Lane Wagner** ([@wagslane](https://github.com/wagslane)) et son équipe chez la plateforme [**Boot.dev**](https://www.boot.dev/tracks/backend) pour ce projet guidé inspirant. Leur travail est une excellente ressource pour l'apprentissage du développement backend.

---
Commencez dès aujourd'hui à agréger vos flux RSS préférés avec Gator ! 🐊

