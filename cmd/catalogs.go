package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var catalogsCmd = &cobra.Command{
	Use:   "catalogs",
	Short: "Manage ticket classification catalogs",
	Long:  `Commands to manage and interact with ticket classification catalogs in the RTMS system.`,
}

func init() {
	rootCmd.AddCommand(catalogsCmd)

	// Subcommands
	catalogsCmd.AddCommand(getCatalogsCmd)
	catalogsCmd.AddCommand(getDefaultCatalogsCmd)
	catalogsCmd.AddCommand(getCatalogItemsCmd)
	catalogsCmd.AddCommand(getRootCatalogCmd)
}

var getCatalogsCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of Ticket classification catalogs and items",
	RunE: func(cmd *cobra.Command, args []string) error {
		cloudTempleID, _ := cmd.Flags().GetString("cloud-temple-id")
		availableItems, _ := cmd.Flags().GetBool("available-items")
		isRoot, _ := cmd.Flags().GetBool("is-root")

		response, err := client.GetCatalogs(cloudTempleID, availableItems, isRoot)
		if err != nil {
			return err
		}
		formattedOutput, err := formatOutput(response)
		if err != nil {
			return err
		}
		fmt.Println(formattedOutput)
		return nil
	},
}

var getDefaultCatalogsCmd = &cobra.Command{
	Use:   "defaults",
	Short: "Get a list of all default ticket classification catalogs and catalog items",
	RunE: func(cmd *cobra.Command, args []string) error {
		availableItems, _ := cmd.Flags().GetBool("available-items")
		isRoot, _ := cmd.Flags().GetBool("is-root")

		response, err := client.GetDefaultCatalogs(availableItems, isRoot)
		if err != nil {
			return err
		}
		formattedOutput, err := formatOutput(response)
		if err != nil {
			return err
		}
		fmt.Println(formattedOutput)
		return nil
	},
}

var getCatalogItemsCmd = &cobra.Command{
	Use:   "items [catalog-id]",
	Short: "Get a list of items for a catalog",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		catalogID := args[0]
		enabled, _ := cmd.Flags().GetBool("enabled")
		var enabledPtr *bool
		if cmd.Flags().Changed("enabled") {
			enabledPtr = &enabled
		}

		response, err := client.GetCatalogItems(catalogID, enabledPtr)
		if err != nil {
			return err
		}
		formattedOutput, err := formatOutput(response)
		if err != nil {
			return err
		}
		fmt.Println(formattedOutput)
		return nil
	},
}

var getRootCatalogCmd = &cobra.Command{
	Use:   "root",
	Short: "Get the root required catalog",
	RunE: func(cmd *cobra.Command, args []string) error {
		catalogType, _ := cmd.Flags().GetString("type")
		availableItems, _ := cmd.Flags().GetBool("available-items")

		response, err := client.GetRootCatalog(catalogType, availableItems)
		if err != nil {
			return err
		}
		formattedOutput, err := formatOutput(response)
		if err != nil {
			return err
		}
		fmt.Println(formattedOutput)
		return nil
	},
}

func init() {
	getCatalogsCmd.Flags().Bool("available-items", false, "Show classification catalog with their available items")
	getCatalogsCmd.Flags().Bool("is-root", false, "If true, only classification root catalogs will be displayed")

	getDefaultCatalogsCmd.Flags().Bool("available-items", false, "Show classification catalog with their available items")
	getDefaultCatalogsCmd.Flags().Bool("is-root", false, "If true, only classification root catalogs will be displayed")

	getCatalogItemsCmd.Flags().Bool("enabled", false, "Display only enabled or disabled catalog items")

	getRootCatalogCmd.Flags().String("type", "", "Required Catalog type (origin, perimeter, or nature)")
	getRootCatalogCmd.Flags().Bool("available-items", false, "Display associated catalog items")
	getRootCatalogCmd.MarkFlagRequired("type")
}
