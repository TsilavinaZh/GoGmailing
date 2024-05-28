package cmd

import (
    "github.com/spf13/cobra"
    "github.com/fatih/color"
)

var rootCmd = &cobra.Command{
    Use:   "email-client",
    Short: "Email Client CLI",
    Long:  `Un client email en ligne de commande pour envoyer et recevoir des emails.`,
}

func Execute() {
    color.Blue("Bienvenue dans Email Client CLI")
    if err := rootCmd.Execute(); err != nil {
        color.Red("Erreur: %v", err)
    }
}

func init() {
    cobra.OnInitialize()
    rootCmd.AddCommand(sendCmd)
}