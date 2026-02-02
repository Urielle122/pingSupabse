# pingSupabse

[![CI](https://github.com/Urielle122/pingSupabse/actions/workflows/main.yml/badge.svg)](https://github.com/Urielle122/pingSupabse/actions/workflows/main.yml)
[![License](https://img.shields.io/github/license/Urielle122/pingSupabse.svg)](https://github.com/Urielle122/pingSupabse/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/Urielle122/pingSupabse.svg)](https://github.com/Urielle122/pingSupabse)
[![Go Report Card](https://goreportcard.com/badge/github.com/Urielle122/pingSupabse)](https://goreportcard.com/report/github.com/Urielle122/pingSupabse)
[![pkg.go.dev](https://pkg.go.dev/badge/github.com/Urielle122/pingSupabse.svg)](https://pkg.go.dev/github.com/Urielle122/pingSupabse)
[![Coverage Status](https://img.shields.io/codecov/c/github/Urielle122/pingSupabse.svg)](https://codecov.io/gh/Urielle122/pingSupabse)

Français / French & English version included below.

---

FRANÇAIS

Petit utilitaire en Go pour "ping" la base Supabase (Postgres) et garder l'instance éveillée. Le script peut aussi envoyer des notifications vers Slack via un webhook.

Résumé
- Langage : Go
- Fichier principal : `main.go` (ping DB, notifications Slack)
- CI : GitHub Actions (`.github/workflows/main.yml`) — workflow programmé toutes les 48h et exécutable manuellement

Prérequis
- Go 1.25+ installé (local)
- Une base Supabase (DATABASE_URL)
- (Optionnel) Un webhook Slack pour recevoir les notifications (SLACK_WEBHOOK_URL)

Installation
1. Cloner le dépôt :

```bash
git clone https://github.com/Urielle122/pingSupabse.git
cd pingSupabse
```

2. Installer les dépendances Go :

```bash
go mod tidy
```

Configuration
Créez un fichier `.env` à la racine (utile en local) avec au minimum :

```env
DATABASE_URL=postgres://user:password@host:port/dbname
# Optionnel: webhook Slack pour notifications
SLACK_WEBHOOK_URL=https://hooks.slack.com/services/XXX/YYY/ZZZ
```

Un exemple prêt à l’usage est fourni dans `.env.example`.

En production (GitHub Actions), définissez `DATABASE_URL` et `SLACK_WEBHOOK_URL` dans les Secrets du dépôt (voir la section "How to set GitHub Secrets" plus bas).

Usage
- Exécution locale :

```bash
go run main.go
```

- Comportement :
  - Le programme tente de se connecter à la base et de faire un `Ping`.
  - Il réessaie jusqu'à 3 fois (`maxRetries`) avec 10s de pause entre chaque tentative.
  - En cas de succès/échec, il envoie une notification Slack si `SLACK_WEBHOOK_URL` est défini.

GitHub Actions
Le workflow `.github/workflows/main.yml` exécute le script toutes les 48 heures et peut aussi être déclenché manuellement. Il utilise les secrets `DATABASE_URL` et `SLACK_WEBHOOK_URL`.

How to set GitHub Secrets (FR)
1. Va dans ton dépôt sur GitHub > "Settings" > "Secrets and variables" > "Actions".
2. Clique sur "New repository secret".
3. Ajoute `DATABASE_URL` comme nom et colle ta chaîne de connexion Supabase comme valeur.
4. (Optionnel) Ajoute `SLACK_WEBHOOK_URL` avec l'URL du webhook Slack.
5. Clique sur "Add secret". Les secrets sont chiffrés et disponibles seulement pour les workflows Actions.

Notes de sécurité :
- Ne commite jamais une chaîne de connexion ou un webhook directement dans le dépôt.
- Utilise les Secrets GitHub pour protéger ces valeurs dans CI.

Variables d'environnement
- `DATABASE_URL` (obligatoire) : chaîne de connexion Postgres vers Supabase  
- `SLACK_WEBHOOK_URL` (optionnel) : URL du webhook Slack pour notifications

Débogage / FAQ
- Erreur "DATABASE_URL non définie" : vérifiez que la variable d'environnement est bien définie localement ou dans les Secrets GitHub.
- Notifications Slack non envoyées : vérifiez que le webhook est correct. Les erreurs sont loggées côté action.

Licence
Ce dépôt contient un fichier `LICENSE` (par défaut proposé : MIT) — voir section Licence en bas.

Contribuer
Les contributions sont bienvenues : ouvre une issue ou une PR.

---

ENGLISH

Small Go tool to "ping" a Supabase (Postgres) database to keep the instance awake. The script can also send Slack notifications via an incoming webhook.

Summary
- Language: Go
- Main file: `main.go` (DB ping + Slack notifications)
- CI: GitHub Actions (`.github/workflows/main.yml`) — runs every 48h and is manually triggerable

Prerequisites
- Go 1.25+ (local)
- A Supabase database (DATABASE_URL)
- (Optional) Slack webhook for notifications (SLACK_WEBHOOK_URL)

Install
1. Clone the repo:

```bash
git clone https://github.com/Urielle122/pingSupabse.git
cd pingSupabse
```

2. Install Go dependencies:

```bash
go mod tidy
```

Configuration
Create a `.env` file at the project root (useful locally) with at least:

```env
DATABASE_URL=postgres://user:password@host:port/dbname
# Optional: Slack webhook for notifications
SLACK_WEBHOOK_URL=https://hooks.slack.com/services/XXX/YYY/ZZZ
```

A ready-to-use example is included as `.env.example`.

In production (GitHub Actions), set `DATABASE_URL` and `SLACK_WEBHOOK_URL` in the repository Secrets (see "How to set GitHub Secrets" below).

Usage
- Run locally:

```bash
go run main.go
```

- Behavior:
  - The program connects to the DB and performs a `Ping`.
  - It retries up to 3 times (`maxRetries`) with 10s delay between attempts.
  - On success/failure it sends a Slack notification if `SLACK_WEBHOOK_URL` is set.

How to set GitHub Secrets (EN)
1. Go to your repository on GitHub > "Settings" > "Secrets and variables" > "Actions".
2. Click "New repository secret".
3. Add `DATABASE_URL` as the name and paste your Supabase connection string as the value.
4. (Optional) Add `SLACK_WEBHOOK_URL` with the incoming Slack webhook URL.
5. Click "Add secret". Secrets are encrypted and only available to Actions workflows.

Security notes:
- Never commit connection strings or webhooks into the repo.
- Use GitHub Secrets to protect these values in CI.

Environment variables
- `DATABASE_URL` (required) : Postgres connection string to Supabase  
- `SLACK_WEBHOOK_URL` (optional) : Slack incoming webhook URL for notifications

Troubleshooting / FAQ
- "DATABASE_URL not set" error: ensure the env var is set locally or in GitHub Secrets.
- Slack notifications not sent: verify the webhook; errors are logged by the action.

License
This repository includes a `LICENSE` file (suggested: MIT). See `LICENSE` for details.

Contributing
Contributions are welcome — open an issue or a PR.

Notes
- I updated README with extra badges and a detailed section explaining how to set GitHub Secrets.