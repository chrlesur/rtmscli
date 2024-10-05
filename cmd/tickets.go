package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

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
	}
	getTicketsCmd.Flags().String("name", "", "Filter tickets by subject (name)")
	getTicketsCmd.Flags().IntSlice("status", nil, "Filter Tickets by one or more status (0-6)")
	getTicketsCmd.Flags().String("owner", "", "Filter tickets by owner name")
	getTicketsCmd.Flags().IntSlice("owner-ids", nil, "Filter tickets by one or more owner RTMS identifiers")
	getTicketsCmd.Flags().Bool("is-not-assigned", false, "Filter non assigned tickets")
	getTicketsCmd.Flags().Bool("is-on-delegation", false, "Filter tickets on delegation")

	updateListCommand(getTicketsCmd, "/tickets", func() map[string]string {
		params := make(map[string]string)
		params["cloudTempleId"] = cloudTempleID

		name, _ := getTicketsCmd.Flags().GetString("name")
		if name != "" {
			params["name"] = name
		}

		status, _ := getTicketsCmd.Flags().GetIntSlice("status")
		if len(status) > 0 {
			params["status[]"] = intSliceToString(status)
		}

		owner, _ := getTicketsCmd.Flags().GetString("owner")
		if owner != "" {
			params["owner"] = owner
		}

		ownerIDs, _ := getTicketsCmd.Flags().GetIntSlice("owner-ids")
		if len(ownerIDs) > 0 {
			params["ownerIds[]"] = intSliceToString(ownerIDs)
		}

		isNotAssigned, _ := getTicketsCmd.Flags().GetBool("is-not-assigned")
		if isNotAssigned {
			params["isNotAssigned"] = "true"
		}

		isOnDelegation, _ := getTicketsCmd.Flags().GetBool("is-on-delegation")
		if isOnDelegation {
			params["isOnDelegation"] = "true"
		}

		return params
	})

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

	// Attachments subcommand
	attachmentsCmd := &cobra.Command{
		Use:   "attachments",
		Short: "Manage ticket attachments",
	}
	ticketsCmd.AddCommand(attachmentsCmd)

	// List attachments
	listAttachmentsCmd := &cobra.Command{
		Use:   "list [ticket-id]",
		Short: "List ticket attachments",
		Args:  cobra.ExactArgs(1),
		RunE:  listTicketAttachments,
	}
	attachmentsCmd.AddCommand(listAttachmentsCmd)

	// Upload attachment
	uploadAttachmentCmd := &cobra.Command{
		Use:   "upload [ticket-id] [file-path]",
		Short: "Upload a ticket attachment",
		Args:  cobra.ExactArgs(2),
		RunE:  uploadTicketAttachment,
	}
	attachmentsCmd.AddCommand(uploadAttachmentCmd)

	// Download attachment
	downloadAttachmentCmd := &cobra.Command{
		Use:   "download [attachment-id] [output-path]",
		Short: "Download a ticket attachment",
		Args:  cobra.ExactArgs(2),
		RunE:  downloadTicketAttachment,
	}
	attachmentsCmd.AddCommand(downloadAttachmentCmd)

	// Remove attachment
	removeAttachmentCmd := &cobra.Command{
		Use:   "remove [attachment-id]",
		Short: "Remove a ticket attachment",
		Args:  cobra.ExactArgs(1),
		RunE:  removeTicketAttachment,
	}
	attachmentsCmd.AddCommand(removeAttachmentCmd)

	// Comments subcommand
	commentsCmd := &cobra.Command{
		Use:   "comments",
		Short: "Manage ticket comments",
	}
	ticketsCmd.AddCommand(commentsCmd)

	// List all comments
	listAllCommentsCmd := &cobra.Command{
		Use:   "list-all",
		Short: "Get all Ticket comments",
		RunE:  listAllTicketComments,
	}
	listAllCommentsCmd.Flags().Int("ticket", 0, "Filter by ticket ID")
	listAllCommentsCmd.Flags().Int("user", 0, "Filter by user ID")
	commentsCmd.AddCommand(listAllCommentsCmd)

	// List comments for a specific ticket
	listTicketCommentsCmd := &cobra.Command{
		Use:   "list [ticket-id]",
		Short: "Get Ticket comments by ticket",
		Args:  cobra.ExactArgs(1),
		RunE:  listTicketComments,
	}
	commentsCmd.AddCommand(listTicketCommentsCmd)

	// Post comment
	postCommentCmd := &cobra.Command{
		Use:   "post [ticket-id]",
		Short: "Post Ticket comment",
		Args:  cobra.ExactArgs(1),
		RunE:  postTicketComment,
	}
	postCommentCmd.Flags().String("content", "", "Comment content")
	postCommentCmd.Flags().Bool("private", false, "Comment privacy")
	postCommentCmd.Flags().Int("duration", 0, "Working time on the Ticket")
	postCommentCmd.MarkFlagRequired("content")
	commentsCmd.AddCommand(postCommentCmd)

	// Edit comment
	editCommentCmd := &cobra.Command{
		Use:   "edit [comment-id]",
		Short: "Edit Ticket comment",
		Args:  cobra.ExactArgs(1),
		RunE:  editTicketComment,
	}
	editCommentCmd.Flags().String("content", "", "Comment content")
	editCommentCmd.Flags().Bool("private", false, "Comment privacy")
	editCommentCmd.Flags().Int("duration", 0, "Working time on the Ticket")
	commentsCmd.AddCommand(editCommentCmd)

	// Tags subcommand
	tagsCmd := &cobra.Command{
		Use:   "tags",
		Short: "Manage ticket tags",
	}
	ticketsCmd.AddCommand(tagsCmd)

	// List tags
	listTagsCmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of ticket tags",
		RunE:  listTicketTags,
	}
	listTagsCmd.Flags().String("label", "", "Filter by label")
	tagsCmd.AddCommand(listTagsCmd)

	// Create tag
	createTagCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a ticket tag",
		RunE:  createTicketTag,
	}
	createTagCmd.Flags().String("label", "", "Tag label")
	createTagCmd.Flags().String("description", "", "Tag description")
	createTagCmd.Flags().IntSlice("tickets", nil, "List of ticket IDs to associate with the tag")
	createTagCmd.MarkFlagRequired("label")
	tagsCmd.AddCommand(createTagCmd)

	// Get tag details
	getTagDetailsCmd := &cobra.Command{
		Use:   "details [id]",
		Short: "Get details of a tag",
		Args:  cobra.ExactArgs(1),
		RunE:  getTicketTagDetails,
	}
	tagsCmd.AddCommand(getTagDetailsCmd)

	// Remove tag
	removeTagCmd := &cobra.Command{
		Use:   "remove [id]",
		Short: "Remove a ticket tag",
		Args:  cobra.ExactArgs(1),
		RunE:  removeTicketTag,
	}
	tagsCmd.AddCommand(removeTagCmd)

	// Edit tag
	editTagCmd := &cobra.Command{
		Use:   "edit [id]",
		Short: "Edit ticket tag",
		Args:  cobra.ExactArgs(1),
		RunE:  editTicketTag,
	}
	editTagCmd.Flags().String("label", "", "Tag label")
	editTagCmd.Flags().String("description", "", "Tag description")
	editTagCmd.Flags().IntSlice("tickets", nil, "List of ticket IDs to associate with the tag")
	tagsCmd.AddCommand(editTagCmd)

	// Get tickets by tag
	getTicketsByTagCmd := &cobra.Command{
		Use:   "tickets [id]",
		Short: "Gets tickets that match a given tag",
		Args:  cobra.ExactArgs(1),
		RunE:  getTicketsByTag,
	}
	tagsCmd.AddCommand(getTicketsByTagCmd)

}

func getTickets(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	name, _ := cmd.Flags().GetString("name")
	status, _ := cmd.Flags().GetIntSlice("status")
	owner, _ := cmd.Flags().GetString("owner")
	ownerIDs, _ := cmd.Flags().GetIntSlice("owner-ids")
	isNotAssigned, _ := cmd.Flags().GetBool("is-not-assigned")
	isOnDelegation, _ := cmd.Flags().GetBool("is-on-delegation")

	params := make(map[string]string)
	params["cloudTempleId"] = cloudTempleID
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

	dataChan, errChan := client.StreamData("/tickets", params, batchSize)

	var tickets []interface{}
	var processingError error

	for {
		select {
		case item, ok := <-dataChan:
			if !ok {
				// Le canal de données est fermé, arrêtez le traitement
				goto ProcessingComplete
			}
			tickets = append(tickets, item)
		case err, ok := <-errChan:
			if !ok {
				// Le canal d'erreurs est fermé, continuez le traitement
				continue
			}
			// Une erreur s'est produite, arrêtez le traitement
			processingError = fmt.Errorf("erreur lors de la récupération des tickets : %w", err)
			goto ProcessingComplete
		}
	}

ProcessingComplete:
	if processingError != nil {
		return processingError
	}

	// Formatage de la sortie
	output, err := formatOutput(tickets, format)
	if err != nil {
		return fmt.Errorf("erreur lors du formatage de la sortie des tickets : %w", err)
	}

	// Affichage de la sortie
	fmt.Fprintln(os.Stdout, output)

	return nil
}

func createTicket(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	description, _ := cmd.Flags().GetString("description")
	owner, _ := cmd.Flags().GetInt("owner")
	catalogItems, _ := cmd.Flags().GetIntSlice("catalog-items")
	format, _ := cmd.Flags().GetString("format")

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
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func getTicketsCount(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	status, _ := cmd.Flags().GetInt("status")
	response, err := client.GetTicketsCount(cloudTempleID, status)
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func getTicketDetails(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	response, err := client.GetTicketDetails(args[0])
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func editTicket(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	description, _ := cmd.Flags().GetString("description")
	owner, _ := cmd.Flags().GetInt("owner")
	catalogItems, _ := cmd.Flags().GetIntSlice("catalog-items")
	format, _ := cmd.Flags().GetString("format")

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
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func getTicketCatalogs(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
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
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func getTicketsStats(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	response, err := client.GetTicketsStats(cloudTempleID)
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func listTicketAttachments(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	response, err := client.ListTicketAttachments(args[0])
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func uploadTicketAttachment(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	ticketID := args[0]
	filePath := args[1]

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	filename := filepath.Base(filePath)
	response, err := client.UploadTicketAttachment(ticketID, filename, content)
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func downloadTicketAttachment(cmd *cobra.Command, args []string) error {
	attachmentID := args[0]
	outputPath := args[1]

	response, err := client.DownloadTicketAttachment(attachmentID)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outputPath, response, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	fmt.Printf("Attachment downloaded successfully to %s\n", outputPath)
	return nil
}

func removeTicketAttachment(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	response, err := client.RemoveTicketAttachment(args[0])
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func listAllTicketComments(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	ticketID, _ := cmd.Flags().GetInt("ticket")
	userID, _ := cmd.Flags().GetInt("user")

	params := make(map[string]string)
	if ticketID != 0 {
		params["ticket"] = strconv.Itoa(ticketID)
	}
	if userID != 0 {
		params["user"] = strconv.Itoa(userID)
	}

	response, err := client.GetTicketComments(cloudTempleID, params)
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func listTicketComments(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	ticketID := args[0]
	response, err := client.GetTicketCommentsByTicket(ticketID, nil)
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func postTicketComment(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	ticketID := args[0]
	content, _ := cmd.Flags().GetString("content")
	private, _ := cmd.Flags().GetBool("private")
	duration, _ := cmd.Flags().GetInt("duration")

	commentData := map[string]interface{}{
		"content": content,
		"private": private,
	}
	if duration != 0 {
		commentData["duration"] = duration
	}

	response, err := client.PostTicketComment(ticketID, commentData)
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func editTicketComment(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	commentID := args[0]
	content, _ := cmd.Flags().GetString("content")
	private, _ := cmd.Flags().GetBool("private")
	duration, _ := cmd.Flags().GetInt("duration")

	commentData := make(map[string]interface{})
	if content != "" {
		commentData["content"] = content
	}
	if cmd.Flags().Changed("private") {
		commentData["private"] = private
	}
	if duration != 0 {
		commentData["duration"] = duration
	}

	response, err := client.EditTicketComment(commentID, commentData)
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func listTicketTags(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	label, _ := cmd.Flags().GetString("label")
	params := make(map[string]string)
	if label != "" {
		params["label"] = label
	}
	response, err := client.GetTicketTags(cloudTempleID, params)
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func createTicketTag(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	label, _ := cmd.Flags().GetString("label")
	description, _ := cmd.Flags().GetString("description")
	tickets, _ := cmd.Flags().GetIntSlice("tickets")

	tagData := map[string]interface{}{
		"label": label,
	}
	if description != "" {
		tagData["description"] = description
	}
	if len(tickets) > 0 {
		tagData["tickets"] = tickets
	}

	response, err := client.CreateTicketTag(cloudTempleID, tagData)
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func getTicketTagDetails(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	response, err := client.GetTicketTagDetails(args[0])
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func removeTicketTag(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	response, err := client.RemoveTicketTag(args[0])
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func editTicketTag(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	label, _ := cmd.Flags().GetString("label")
	description, _ := cmd.Flags().GetString("description")
	tickets, _ := cmd.Flags().GetIntSlice("tickets")

	tagData := make(map[string]interface{})
	if label != "" {
		tagData["label"] = label
	}
	if description != "" {
		tagData["description"] = description
	}
	if len(tickets) > 0 {
		tagData["tickets"] = tickets
	}

	response, err := client.EditTicketTag(args[0], tagData)
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func getTicketsByTag(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	response, err := client.GetTicketsByTag(args[0], nil)
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}
