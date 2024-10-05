package cmd

import (
	"fmt"
	"os"

	"github.com/chrlesur/rtmscli/pkg/api"
	"github.com/spf13/cobra"
)

const (
	Version = "1.2.0 beta release"
)

var (
	cloudTempleID string
	host          string
	client        *api.RTMSClient
	outputFormat  string
	limit         int
	batchSize     int
	filter        string
)

var rootCmd = &cobra.Command{
	Use:   "rtmscli",
	Short: "RTMS CLI est une interface en ligne de commande pour l'API RTMS",
	Long: fmt.Sprintf(`RTMS CLI (version %s) vous permet d'interagir avec l'API RTMS depuis la ligne de commande.
Il fournit des commandes pour gérer les appliances, les hôtes, les tickets et plus encore.`, Version),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Ignorer l'initialisation du client pour la commande version
		if cmd.Use == "version" {
			return nil
		}

		// Valider le format de sortie
		validFormats := map[string]bool{"json": true, "text": true, "html": true, "markdown": true}
		if !validFormats[outputFormat] {
			return fmt.Errorf("format de sortie invalide : %s. Les formats supportés sont json, text, html, et markdown", outputFormat)
		}

		// Initialisation du client
		apiKey := os.Getenv("RTMS_API_KEY")
		if apiKey == "" {
			return fmt.Errorf("la variable d'environnement RTMS_API_KEY n'est pas définie")
		}

		var err error
		client, err = api.NewRTMSClient(apiKey, host, IsBase64)
		if err != nil {
			return fmt.Errorf("erreur lors de l'initialisation du client RTMS : %w", err)
		}

		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cloudTempleID, "cloud-temple-id", "c", "", "ID Cloud Temple (requis pour la plupart des commandes)")
	rootCmd.PersistentFlags().StringVarP(&host, "host", "H", "rtms-api.cloud-temple.com", "Hôte de l'API RTMS")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "format", "f", "json", "Format de sortie (json, text, html, markdown)")

	rootCmd.PersistentFlags().IntVarP(&limit, "limit", "l", 0, "Limite le nombre de résultats retournés (par défaut : 0 pour illimité)")
	rootCmd.PersistentFlags().IntVar(&batchSize, "batch-size", 100, "Nombre d'éléments à récupérer par lot")
	rootCmd.PersistentFlags().StringVar(&filter, "filter", "", "Filtre les résultats (format dépendant de la commande)")

	// Ajout de la commande version
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Affiche le numéro de version de RTMS CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("RTMS CLI version %s\n", Version)
	},
}
