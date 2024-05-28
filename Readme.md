
# GoGmailing
<img src="GoGmailing.png />
GoGmailing est un client email en ligne de commande développé en Go. Il permet aux utilisateurs d'envoyer des emails via un serveur SMTP.

## Fonctionnalités

- Envoyer des emails via SMTP
- Personnaliser les couleurs des messages de la CLI


## Utilisation

### Commandes

- `send`: Envoyer un email

### Options

- `--smtp, -s` : Serveur SMTP (obligatoire)
- `--port, -p` : Port SMTP (par défaut : 587)
- `--username, -u` : Nom d'utilisateur SMTP (obligatoire)
- `--password, -w` : Mot de passe SMTP (obligatoire)
- `--to, -t` : Adresse email du destinataire (obligatoire)
- `--subject, -j` : Sujet de l'email (obligatoire)
- `--body, -b` : Corps de l'email (obligatoire)

### Exemple

Pour envoyer un email, exécutez la commande suivante :

```sh
go run main.go send --smtp "smtp.example.com" --port 587 --username "your_email@example.com" --password "yourpassword" --to "recipient@example.com" --subject "Hello" --body "This is a test email."
```

## Structure du projet

```plaintext
email-client/
├── cmd/
│   ├── root.go
│   └── send.go
├── go.mod
└── main.go
```

- `main.go`: Point d'entrée du programme.
- `cmd/root.go`: Définition de la commande racine.
- `cmd/send.go`: Définition de la commande `send`.
