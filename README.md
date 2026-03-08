# CryptoDash

Dashboard cryptomonnaie de démonstration, réalisé pour illustrer les capacités techniques mises en œuvre dans des projets web professionnels destinés aux futurs clients.

## Présentation

CryptoDash est une application web monolithique qui simule un tableau de bord de portefeuille crypto. Elle met en scène un environnement de trading fictif avec vue d'ensemble du portefeuille, suivi des actifs et passage d'ordres via une interface moderne et réactive.

Fonctionnalités principales :

- Vue d'ensemble du portefeuille (solde total, profit/perte 24h, ordres en attente)
- Tableau des marchés avec prix, variation et quantités détenues
- Modale de négociation pour acheter des actifs par quantité
- Mise à jour partielle de l'interface sans rechargement complet
- Notifications toast de confirmation après validation d'un ordre

## Stack technique

- **Backend** : Go (net/http standard)
- **Templates HTML** : templ, composants réutilisables
- **UI** : TemplUI, Tailwind CSS
- **Interactivité** : HTMX, Alpine.js
- **Conteneurisation** : Docker

L'architecture vise à montrer une approche pragmatique : rendu HTML côté serveur, mises à jour ciblées via HTMX, et composants UI modernes sans framework JavaScript lourd.

## Prérequis

- Go 1.24 ou supérieur
- Node.js (pour Tailwind CSS)
- templ CLI (`go install github.com/a-h/templ/cmd/templ@latest`)

## Installation et démarrage

```bash
# Génération des templates et du CSS
templ generate
npm install
npx tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css

# Lancement en développement (optionnel : utiliser air pour le hot-reload)
go run .
```

L'application écoute sur `http://localhost:8090`.

## Déploiement

Un Dockerfile multi-stage est fourni pour la production :

```bash
docker build -t crypto-demo .
docker run -p 8090:8090 crypto-demo
```

## Structure du projet

```
crypto-demo/
├── main.go                 # Point d'entrée et routes HTTP
├── internal/
│   ├── models/             # Modèles métier (Asset, Portfolio, Transaction)
│   └── repository/         # Couche d'accès données (mock en mémoire)
├── ui/
│   ├── components/         # Composants TemplUI réutilisables
│   ├── layouts/            # Layouts de page (base, sidebar, navbar)
│   └── pages/              # Pages et fragments (dashboard, modale trade)
└── assets/                 # CSS et assets statiques
```

## À propos

Projet de démonstration conçu comme référence pour des applications web sur mesure, mettant en avant la rapidité de développement, la maintenabilité et une expérience utilisateur soignée avec des technologies éprouvées.
