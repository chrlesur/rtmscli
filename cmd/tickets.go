package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ticketsCmd = &cobra.Command{
	Use:   "tickets",
	Short: "Manage tickets",
	Long:  `Manage tickets, including listing, creating, editing, and viewing ticket details and statistics.`,
}

func init() {
	rootCmd.AddCommand(ticketsCmd)

	// Get tickets
	getTicketsCmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of Tickets",
		RunE:  getTickets,
	}
	getTicketsCmd.Flags().String("name", "", "Filter tickets by subject (name)")
	getTicketsCmd.Flags().IntSlice("status", nil, "Filter Tickets by one or more status (0-6)")
	getTicketsCmd.Flags().String("owner", "", "Filter tickets by owner name")
	getTicketsCmd.Flags().IntSlice("owner-ids", nil, "Filter tickets by one or more owner RTMS identifiers")
	getTicketsCmd.Flags().Bool("is-not-assigned", false, "Filter non assigned tickets")
	getTicketsCmd.Flags().Bool("is-on-delegation", false, "Filter tickets on delegation")
	ticketsCmd.AddCommand(getTicketsCmd)

	// Create ticket
	createTicketCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new Ticket",
		RunE:  createTicket,
	}
	createTicketCmd.Flags().String("name", "", "A title, the issue in short")
	createTicketCmd.Flags().String("description", "", "Detailed description of the issue")
	createTicketCmd.Flags().Int("owner", 0, "Identifier of the user in charge of solving the issue")
	createTicketCmd.Flags().IntSlice("catalog-items", nil, "Collection of classification catalog item identifiers")
	createTicketCmd.MarkFlagRequired("name")
	createTicketCmd.MarkFlagRequired("description")
	ticketsCmd.AddCommand(createTicketCmd)

	// Get tickets count
	getTicketsCountCmd := &cobra.Command{
		Use:   "count",
		Short: "Get number of Tickets",
		RunE:  getTicketsCount,
	}
	getTicketsCountCmd.Flags().Int("status", -1, "Filter Tickets by status (0-6)")
	ticketsCmd.AddCommand(getTicketsCountCmd)

	// Get ticket details
	getTicketDetailsCmd := &cobra.Command{
		Use:   "details [id]",
		Short: "Get Ticket details",
		Args:  cobra.ExactArgs(1),
		RunE:  getTicketDetails,
	}
	ticketsCmd.AddCommand(getTicketDetailsCmd)

	// Edit ticket
	editTicketCmd := &cobra.Command{
		Use:   "edit [id]",
		Short: "Edit Ticket information",
		Args:  cobra.ExactArgs(1),
		RunE:  editTicket,
	}
	editTicketCmd.Flags().String("name", "", "A new title, the issue in short")
	editTicketCmd.Flags().String("description", "", "A new detailed description of the issue")
	editTicketCmd.Flags().Int("owner", 0, "New identifier of the user in charge of solving the issue")
	editTicketCmd.Flags().IntSlice("catalog-items", nil, "New collection of classification catalog item identifiers")
	ticketsCmd.AddCommand(editTicketCmd)

	// Get ticket catalogs
	getTicketCatalogsCmd := &cobra.Command{
		Use:   "catalogs [id]",
		Short: "Get Ticket catalogs",
		Args:  cobra.ExactArgs(1),
		RunE:  getTicketCatalogs,
	}
	getTicketCatalogsCmd.Flags().Bool("selected-item", false, "Show classification catalog with selected items for this ticket")
	getTicketCatalogsCmd.Flags().Bool("available-items", false, "Show classification catalog with all available items for this ticket")
	getTicketCatalogsCmd.Flags().Bool("is-root", false, "If true, only classification root catalogs will be displayed")
	ticketsCmd.AddCommand(getTicketCatalogsCmd)

	// Get tickets stats
	getTicketsStatsCmd := &cobra.Command{
		Use:   "stats",
		Short: "Get tickets status stats",
		RunE:  getTicketsStats,
	}
	ticketsCmd.AddCommand(getTicketsStatsCmd)
}

func getTickets(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	status, _ := cmd.Flags().GetIntSlice("status")
	owner, _ := cmd.Flags().GetString("owner")
	ownerIDs, _ := cmd.Flags().GetIntSlice("owner-ids")
	isNotAssigned, _ := cmd.Flags().GetBool("is-not-assigned")
	isOnDelegation, _ := cmd.Flags().GetBool("is-on-delegation")

	params := make(map[string]string)
	if name != "" {
		params["name"] = name
	}
	if len(status) > 0 {
		params["status[]"] = intSliceToString(status)
	}
	if owner != "" {
		params["owner"] = owner
	}
	if len(ownerIDs) > 0 {
		params["ownerIds[]"] = intSliceToString(ownerIDs)
	}
	if isNotAssigned {
		params["isNotAssigned"] = "true"
	}
	if isOnDelegation {
		params["isOnDelegation"] = "true"
	}

	response, err := client.GetTickets(cloudTempleID, params)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func createTicket(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	description, _ := cmd.Flags().GetString("description")
	owner, _ := cmd.Flags().GetInt("owner")
	catalogItems, _ := cmd.Flags().GetIntSlice("catalog-items")

	ticketData := map[string]interface{}{
		"name":        name,
		"description": description,
	}
	if owner != 0 {
		ticketData["owner"] = owner
	}
	if len(catalogItems) > 0 {
		ticketData["catalogItemsCollection"] = catalogItems
	}

	response, err := client.CreateTicket(cloudTempleID, ticketData)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getTicketsCount(cmd *cobra.Command, args []string) error {
	status, _ := cmd.Flags().GetInt("status")
	response, err := client.GetTicketsCount(cloudTempleID, status)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getTicketDetails(cmd *cobra.Command, args []string) error {
	response, err := client.GetTicketDetails(args[0])
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func editTicket(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	description, _ := cmd.Flags().GetString("description")
	owner, _ := cmd.Flags().GetInt("owner")
	catalogItems, _ := cmd.Flags().GetIntSlice("catalog-items")

	ticketData := make(map[string]interface{})
	if name != "" {
		ticketData["name"] = name
	}
	if description != "" {
		ticketData["description"] = description
	}
	if owner != 0 {
		ticketData["owner"] = owner
	}
	if len(catalogItems) > 0 {
		ticketData["catalogItemsCollection"] = catalogItems
	}

	response, err := client.EditTicket(args[0], ticketData)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getTicketCatalogs(cmd *cobra.Command, args []string) error {
	selectedItem, _ := cmd.Flags().GetBool("selected-item")
	availableItems, _ := cmd.Flags().GetBool("available-items")
	isRoot, _ := cmd.Flags().GetBool("is-root")

	params := make(map[string]string)
	if selectedItem {
		params["selectedItem"] = "true"
	}
	if availableItems {
		params["availableItems"] = "true"
	}
	if isRoot {
		params["isRoot"] = "true"
	}

	response, err := client.GetTicketCatalogs(args[0], params)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getTicketsStats(cmd *cobra.Command, args []string) error {
	response, err := client.GetTicketsStats(cloudTempleID)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}
