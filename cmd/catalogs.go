package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var catalogsCmd = &cobra.Command{
	Use:   "catalogs",
	Short: "Manage ticket classification catalogs",
	Long:  `Commands to manage and interact with ticket classification catalogs in the RTMS system.`,
}

func init() {
	rootCmd.AddCommand(catalogsCmd)

	// Get catalogs
	getCatalogsCmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of Ticket classification catalogs and items",
	}
	getCatalogsCmd.Flags().Bool("available-items", false, "Show classification catalog with their available items")
	getCatalogsCmd.Flags().Bool("is-root", false, "If true, only classification root catalogs will be displayed")

	updateListCommand(getCatalogsCmd, "/catalogs", func() map[string]string {
		params := make(map[string]string)
		params["cloudTempleId"] = cloudTempleID

		availableItems, _ := getCatalogsCmd.Flags().GetBool("available-items")
		params["availableItems"] = strconv.FormatBool(availableItems)

		isRoot, _ := getCatalogsCmd.Flags().GetBool("is-root")
		params["isRoot"] = strconv.FormatBool(isRoot)

		return params
	})

	catalogsCmd.AddCommand(getCatalogsCmd)

	// Get default catalogs
	getDefaultCatalogsCmd := &cobra.Command{
		Use:   "defaults",
		Short: "Get a list of all default ticket classification catalogs and catalog items",
	}
	getDefaultCatalogsCmd.Flags().Bool("available-items", false, "Show classification catalog with their available items")
	getDefaultCatalogsCmd.Flags().Bool("is-root", false, "If true, only classification root catalogs will be displayed")

	updateListCommand(getDefaultCatalogsCmd, "/catalogs/defaults", func() map[string]string {
		params := make(map[string]string)

		availableItems, _ := getDefaultCatalogsCmd.Flags().GetBool("available-items")
		params["availableItems"] = strconv.FormatBool(availableItems)

		isRoot, _ := getDefaultCatalogsCmd.Flags().GetBool("is-root")
		params["isRoot"] = strconv.FormatBool(isRoot)

		return params
	})

	catalogsCmd.AddCommand(getDefaultCatalogsCmd)

	// Get catalog items
	getCatalogItemsCmd := &cobra.Command{
		Use:   "items [catalog-id]",
		Short: "Get a list of items for a catalog",
		Args:  cobra.ExactArgs(1),
	}
	getCatalogItemsCmd.Flags().Bool("enabled", false, "Display only enabled or disabled catalog items")

	updateListCommand(getCatalogItemsCmd, "/catalogs/{id}/items", func() map[string]string {
		params := make(map[string]string)

		enabled, _ := getCatalogItemsCmd.Flags().GetBool("enabled")
		params["enabled"] = strconv.FormatBool(enabled)

		return params
	})

	catalogsCmd.AddCommand(getCatalogItemsCmd)

	// Get root catalog
	getRootCatalogCmd := &cobra.Command{
		Use:   "root",
		Short: "Get the root required catalog",
	}
	getRootCatalogCmd.Flags().String("type", "", "Required Catalog type (origin, perimeter, or nature)")
	getRootCatalogCmd.Flags().Bool("available-items", false, "Display associated catalog items")
	getRootCatalogCmd.MarkFlagRequired("type")

	getRootCatalogCmd.RunE = func(cmd *cobra.Command, args []string) error {
		catalogType, _ := cmd.Flags().GetString("type")
		availableItems, _ := cmd.Flags().GetBool("available-items")
		format, _ := cmd.Flags().GetString("format")

		response, err := client.GetRootCatalog(catalogType, availableItems)
		if err != nil {
			return err
		}
		formattedOutput, err := formatOutput(response, format)
		if err != nil {
			return err
		}
		fmt.Println(formattedOutput)
		return nil
	}

	catalogsCmd.AddCommand(getRootCatalogCmd)
}
