package cmd

import (
    "fmt"
    "gopkg.in/gomail.v2"
    "github.com/spf13/cobra"
    "github.com/fatih/color"
)

var (
    smtpServer string
    port       int
    username   string
    password   string
    to         string
    subject    string
    body       string
)

var sendCmd = &cobra.Command{
    Use:   "send",
    Short: "Envoyer un email",
    Long:  `Commande pour envoyer un email via SMTP.`,
    Run: func(cmd *cobra.Command, args []string) {
        sendEmail()
    },
}

func init() {
    sendCmd.Flags().StringVarP(&smtpServer, "smtp", "s", "", "Serveur SMTP")
    sendCmd.Flags().IntVarP(&port, "port", "p", 587, "Port SMTP")
    sendCmd.Flags().StringVarP(&username, "username", "u", "", "Nom d'utilisateur SMTP")
    sendCmd.Flags().StringVarP(&password, "password", "w", "", "Mot de passe SMTP")
    sendCmd.Flags().StringVarP(&to, "to", "t", "", "Adresse email du destinataire")
    sendCmd.Flags().StringVarP(&subject, "subject", "j", "", "Sujet de l'email")
    sendCmd.Flags().StringVarP(&body, "body", "b", "", "Corps de l'email")
}

func sendEmail() {
    color.Cyan("Envoi de l'email en cours...")

    m := gomail.NewMessage()
    m.SetHeader("From", username)
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/plain", body)

    d := gomail.NewDialer(smtpServer, port, username, password)

    if err := d.DialAndSend(m); err != nil {
        color.Red("Erreur lors de l'envoi de l'email: %v", err)
    } else {
        color.Green("Email envoyé avec succès!")
    }
}