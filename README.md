# pingSupabse

Petit utilitaire en Go pour "ping" la base Supabase (Postgres) et garder l'instance éveillée. Le script peut aussi envoyer des notifications vers Slack via un webhook.

## Résumé

- Langage : Go
- Fichier principal : `main.go` (ping DB, notifications Slack)
- CI : GitHub Actions (.github/workflows/main.yml) — workflow programmé toutes les 48h et exécutable manuellement

## Prérequis

- Go 1.25+ installé (local)
- Une base Supabase (DATABASE_URL)
- (Optionnel) Un webhook Slack pour recevoir les notifications (SLACK_WEBHOOK_URL)

## Installation

1. Cloner le dépôt :

```bash
git clone https://github.com/Urielle122/pingSupabse.git
cd pingSupabse
```

2. Installer les dépendances Go :

```bash
go mod tidy
```

## Configuration

Créez un fichier `.env` à la racine (utile en local) avec au minimum :

```env
DATABASE_URL=postgres://user:password@host:port/dbname
# Optionnel: webhook Slack pour notifications
SLACK_WEBHOOK_URL=https://hooks.slack.com/services/XXX/YYY/ZZZ
```

En production (GitHub Actions), définissez `DATABASE_URL` et `SLACK_WEBHOOK_URL` dans les Secrets du dépôt.

## Usage

- Exécution locale :

```bash
go run main.go
```

- Comportement :
  - Le programme tente de se connecter à la base et de faire un `Ping`.
  - Il réessaie jusqu'à 3 fois (constante `maxRetries`) avec 10s de pause entre chaque tentative.
  - En cas de succès/échec, il envoie une notification Slack si `SLACK_WEBHOOK_URL` est défini.

## GitHub Actions

Le workflow `.github/workflows/main.yml` exécute le script toutes les 48 heures et peut aussi être déclenché manuellement. Il utilise les secrets `DATABASE_URL` et `SLACK_WEBHOOK_URL`.

Voir le fichier workflow : `.github/workflows/main.yml`

## Variables d'environnement

- DATABASE_URL (obligatoire) : chaîne de connexion Postgres vers Supabase
- SLACK_WEBHOOK_URL (optionnel) : URL du webhook Slack pour notifications

## Débogage / FAQ

- Erreur "DATABASE_URL non définie" : vérifiez que la variable d'environnement est bien définie localement ou dans les Secrets GitHub.
- Notifications Slack non envoyées : vérifiez que le webhook est correct. Les erreurs sont loggées côté action.

## Contribution

Les contributions sont bienvenues : ouvre une issue ou une PR.

## Licence

Choisis une licence si nécessaire (ajoute un fichier LICENSE si tu veux publier sous une licence précise).

---

Notes de l'analyse du repo :
- J'ai parcouru les fichiers `main.go`, `go.mod` et `.github/workflows/main.yml` pour rédiger ce README.
- Les résultats de recherche de code peuvent être incomplets (limités à 10 résultats). Pour voir plus, consulte la recherche du dépôt : https://github.com/Urielle122/pingSupabse/search