package cmd

import (
	"fmt"
	"os"

	"github.com/chrlesur/rtmscli/pkg/api"
	"github.com/spf13/cobra"
)

const (
	Version = "0.1.0 Alpha release" // Définissez ici le numéro de version de votre CLI
)

var (
	cloudTempleID string
	host          string
	client        *api.RTMSClient
)

var rootCmd = &cobra.Command{
	Use:   "rtmscli",
	Short: "RTMS CLI is a command line interface for the RTMS API",
	Long: fmt.Sprintf(`RTMS CLI (version %s) allows you to interact with the RTMS API from the command line.
It provides commands to manage appliances, hosts, tickets, and more.`, Version),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Skip client initialization for the version command
		if cmd.Use == "version" {
			return nil
		}

		// Client initialization
		apiKey := os.Getenv("RTMS_API_KEY")
		if apiKey == "" {
			return fmt.Errorf("RTMS_API_KEY environment variable is not set")
		}

		var err error
		client, err = api.NewRTMSClient(apiKey, host)
		if err != nil {
			return fmt.Errorf("error initializing RTMS client: %w", err)
		}

		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cloudTempleID, "cloud-temple-id", "c", "", "Cloud Temple ID (required for most commands)")
	rootCmd.PersistentFlags().StringVarP(&host, "host", "H", "rtms-api.cloud-temple.com", "RTMS API host")

	// Ajout de la commande version
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of RTMS CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("RTMS CLI version %s\n", Version)
	},
}
