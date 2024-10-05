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
	debug         bool // new debug flag
)

var rootCmd = &cobra.Command{
	Use:   "rtmscli",
	Short: "RTMS CLI is a command line interface for the RTMS API",
	Long: fmt.Sprintf(`RTMS CLI (version %s) allows you to interact with the RTMS API from the command line.
It provides commands to manage appliances, hosts, tickets, and more.`, Version),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Use == "version" {
			return nil
		}

		validFormats := map[string]bool{"json": true, "text": true, "html": true, "markdown": true}
		if !validFormats[outputFormat] {
			return fmt.Errorf("invalid output format: %s. Supported formats are json, text, html, and markdown", outputFormat)
		}

		apiKey := os.Getenv("RTMS_API_KEY")
		if apiKey == "" {
			return fmt.Errorf("environment variable RTMS_API_KEY is not set")
		}

		var err error
		client, err = api.NewRTMSClient(apiKey, host, IsBase64)
		if err != nil {
			return fmt.Errorf("error initializing RTMS client: %w", err)
		}

		client.SetDebug(debug) // Pass the debug flag to api client

		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cloudTempleID, "cloud-temple-id", "c", "", "Cloud Temple ID (required for most commands)")
	rootCmd.PersistentFlags().StringVarP(&host, "host", "H", "rtms-api.cloud-temple.com", "RTMS API host")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "format", "f", "json", "Output format (json, text, html, markdown)")
	rootCmd.PersistentFlags().IntVarP(&limit, "limit", "l", 0, "Limit the number of results returned (default: 0 for unlimited)")
	rootCmd.PersistentFlags().IntVar(&batchSize, "batch-size", 100, "Number of items to fetch per batch")
	rootCmd.PersistentFlags().StringVar(&filter, "filter", "", "Filter results (format depends on the command)")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug mode")

	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the version number of RTMS CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("RTMS CLI version %s\n", Version)
	},
}
