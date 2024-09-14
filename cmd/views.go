package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var viewsCmd = &cobra.Command{
	Use:   "views",
	Short: "Manage monitoring views",
	Long:  `Manage monitoring views, including listing hosts, services, or templates in a monitoring view.`,
}

func init() {
	rootCmd.AddCommand(viewsCmd)

	// List view items
	listViewItemsCmd := &cobra.Command{
		Use:   "list [type] [id]",
		Short: "List hosts, services or templates in a monitoring view",
		Args:  cobra.ExactArgs(2),
		RunE:  listViewItems,
	}
	listViewItemsCmd.Flags().String("order", "DESC", "Data sort order (ASC, DESC)")
	listViewItemsCmd.Flags().String("order-by", "id", "Attribute of the element on which to order")
	listViewItemsCmd.Flags().Int("page", 1, "Current page")
	listViewItemsCmd.Flags().Int("items-per-page", 100, "Number of items on a single page (max 500)")
	viewsCmd.AddCommand(listViewItemsCmd)
}

func listViewItems(cmd *cobra.Command, args []string) error {
	viewType := args[0]
	id := args[1]

	if viewType != "host" && viewType != "service" && viewType != "template" {
		return fmt.Errorf("invalid view type. Must be 'host', 'service', or 'template'")
	}

	order, _ := cmd.Flags().GetString("order")
	orderBy, _ := cmd.Flags().GetString("order-by")
	page, _ := cmd.Flags().GetInt("page")
	itemsPerPage, _ := cmd.Flags().GetInt("items-per-page")

	params := make(map[string]string)
	params["order"] = order
	params["orderBy"] = orderBy
	params["page"] = fmt.Sprintf("%d", page)
	params["itemsPerPage"] = fmt.Sprintf("%d", itemsPerPage)

	response, err := client.GetViewItems(viewType, id, params)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}
