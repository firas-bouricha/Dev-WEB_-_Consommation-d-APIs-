# Dev_WEB_API

Site web en Go qui affiche les titres de l'artiste JUL grâce à l'API Spotify.

## Prérequis

- Go
- Un compte développeur Spotify (Client ID et Client Secret)

## Lancement

1. Copier `.env.example` vers `.env` et remplir les valeurs.
2. Lancer le serveur : `go run .`
3. Ouvrir http://localhost:8080

## Routes

- `/` : liste des titres avec filtres
- `/tracks/{id}` : détail d'un titre
