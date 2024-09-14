package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tenantsCmd = &cobra.Command{
	Use:   "tenants",
	Short: "Manage tenants",
	Long:  `Manage tenants, including listing, creating, and managing tenant details, SSH keys, and workflow emails.`,
}

func init() {
	rootCmd.AddCommand(tenantsCmd)

	// Get tenants
	getTenantsCmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of Tenants",
		RunE:  getTenants,
	}
	getTenantsCmd.Flags().String("name", "", "Filter tenants by name")
	getTenantsCmd.Flags().Int("responsible-team-id", 0, "Filter tenants by responsible team ID")
	getTenantsCmd.Flags().Int("sdm-id", 0, "Filter tenants by SDM ID")
	tenantsCmd.AddCommand(getTenantsCmd)

	// Create tenant
	createTenantCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new Tenant",
		RunE:  createTenant,
	}
	createTenantCmd.Flags().String("name", "", "Tenant's name")
	createTenantCmd.Flags().String("phone", "", "Tenant's phone")
	createTenantCmd.Flags().String("address", "", "Tenant's address")
	createTenantCmd.Flags().String("postal-code", "", "Tenant's postal code")
	createTenantCmd.Flags().String("city", "", "Tenant's city")
	createTenantCmd.Flags().String("country", "", "Tenant's country")
	createTenantCmd.Flags().Int("responsible-team", 0, "Tenant's responsible team ID")
	createTenantCmd.Flags().Int("contact", 0, "Tenant's contact ID")
	createTenantCmd.Flags().StringSlice("watchers", nil, "List of default watcher email addresses")
	createTenantCmd.Flags().Bool("is-enabled", true, "Is Tenant active?")
	createTenantCmd.Flags().String("cloud-temple-id", "", "MySI Tenant's identifier")
	createTenantCmd.MarkFlagRequired("name")
	createTenantCmd.MarkFlagRequired("phone")
	createTenantCmd.MarkFlagRequired("address")
	createTenantCmd.MarkFlagRequired("postal-code")
	createTenantCmd.MarkFlagRequired("city")
	createTenantCmd.MarkFlagRequired("country")
	createTenantCmd.MarkFlagRequired("responsible-team")
	createTenantCmd.MarkFlagRequired("contact")
	tenantsCmd.AddCommand(createTenantCmd)

	// Get tenant details
	getTenantDetailsCmd := &cobra.Command{
		Use:   "details [id]",
		Short: "Get Tenant details",
		Args:  cobra.ExactArgs(1),
		RunE:  getTenantDetails,
	}
	tenantsCmd.AddCommand(getTenantDetailsCmd)

	// Get tenant contacts
	getTenantContactsCmd := &cobra.Command{
		Use:   "contacts [id]",
		Short: "Get Tenant contacts",
		Args:  cobra.ExactArgs(1),
		RunE:  getTenantContacts,
	}
	tenantsCmd.AddCommand(getTenantContactsCmd)

	// Request tenant deletion
	requestTenantDeletionCmd := &cobra.Command{
		Use:   "request-deletion [id]",
		Short: "Request tenant deletion",
		Args:  cobra.ExactArgs(1),
		RunE:  requestTenantDeletion,
	}
	requestTenantDeletionCmd.Flags().Bool("delete", false, "Enable or disable deletion request")
	requestTenantDeletionCmd.MarkFlagRequired("delete")
	tenantsCmd.AddCommand(requestTenantDeletionCmd)

	// SSH Keys subcommand
	sshKeysCmd := &cobra.Command{
		Use:   "ssh-keys",
		Short: "Manage tenant SSH keys",
	}
	tenantsCmd.AddCommand(sshKeysCmd)

	// List SSH keys
	listSSHKeysCmd := &cobra.Command{
		Use:   "list [tenant-id]",
		Short: "List all SSH keys of a tenant",
		Args:  cobra.ExactArgs(1),
		RunE:  listTenantSSHKeys,
	}
	sshKeysCmd.AddCommand(listSSHKeysCmd)

	// Generate SSH key
	generateSSHKeyCmd := &cobra.Command{
		Use:   "generate [tenant-id]",
		Short: "Generate a new SSH key for a tenant",
		Args:  cobra.ExactArgs(1),
		RunE:  generateTenantSSHKey,
	}
	generateSSHKeyCmd.Flags().String("comment", "", "Free comment")
	generateSSHKeyCmd.Flags().Bool("is-active", false, "Key state")
	sshKeysCmd.AddCommand(generateSSHKeyCmd)

	// Delete SSH key
	deleteSSHKeyCmd := &cobra.Command{
		Use:   "delete [key-id]",
		Short: "Delete an SSH key of a tenant",
		Args:  cobra.ExactArgs(1),
		RunE:  deleteTenantSSHKey,
	}
	sshKeysCmd.AddCommand(deleteSSHKeyCmd)

	// Update SSH key
	updateSSHKeyCmd := &cobra.Command{
		Use:   "update [key-id]",
		Short: "Update an SSH key",
		Args:  cobra.ExactArgs(1),
		RunE:  updateTenantSSHKey,
	}
	updateSSHKeyCmd.Flags().Bool("is-active", false, "Key state")
	sshKeysCmd.AddCommand(updateSSHKeyCmd)

	// Workflow Emails subcommand
	workflowEmailsCmd := &cobra.Command{
		Use:   "workflow-emails",
		Short: "Manage tenant workflow emails",
	}
	tenantsCmd.AddCommand(workflowEmailsCmd)

	// Get workflow emails details
	getWorkflowEmailsCmd := &cobra.Command{
		Use:   "get [tenant-id]",
		Short: "Get workflow emails details",
		Args:  cobra.ExactArgs(1),
		RunE:  getTenantWorkflowEmails,
	}
	workflowEmailsCmd.AddCommand(getWorkflowEmailsCmd)

	// Edit workflow emails generalities
	editWorkflowEmailsGeneralitiesCmd := &cobra.Command{
		Use:   "edit-generalities [tenant-id]",
		Short: "Edit workflow emails generalities",
		Args:  cobra.ExactArgs(1),
		RunE:  editTenantWorkflowEmailsGeneralities,
	}
	editWorkflowEmailsGeneralitiesCmd.Flags().String("format", "", "Format of emails sent (HTML or TEXT)")
	editWorkflowEmailsGeneralitiesCmd.Flags().String("from", "", "Email address used to send emails")
	workflowEmailsCmd.AddCommand(editWorkflowEmailsGeneralitiesCmd)

	// Edit workflow emails create ticket
	editWorkflowEmailsCreateTicketCmd := &cobra.Command{
		Use:   "edit-create-ticket [tenant-id]",
		Short: "Edit workflow emails create ticket",
		Args:  cobra.ExactArgs(1),
		RunE:  editTenantWorkflowEmailsCreateTicket,
	}
	workflowEmailsCmd.AddCommand(editWorkflowEmailsCreateTicketCmd)

	// Edit workflow emails update ticket
	editWorkflowEmailsUpdateTicketCmd := &cobra.Command{
		Use:   "edit-update-ticket [tenant-id]",
		Short: "Edit workflow emails update ticket",
		Args:  cobra.ExactArgs(1),
		RunE:  editTenantWorkflowEmailsUpdateTicket,
	}
	workflowEmailsCmd.AddCommand(editWorkflowEmailsUpdateTicketCmd)

	// Edit workflow emails validation client ticket
	editWorkflowEmailsValidationClientTicketCmd := &cobra.Command{
		Use:   "edit-validation-client-ticket [tenant-id]",
		Short: "Edit workflow emails validation client ticket",
		Args:  cobra.ExactArgs(1),
		RunE:  editTenantWorkflowEmailsValidationClientTicket,
	}
	workflowEmailsCmd.AddCommand(editWorkflowEmailsValidationClientTicketCmd)

	// Edit workflow emails close ticket
	editWorkflowEmailsCloseTicketCmd := &cobra.Command{
		Use:   "edit-close-ticket [tenant-id]",
		Short: "Edit workflow emails close ticket",
		Args:  cobra.ExactArgs(1),
		RunE:  editTenantWorkflowEmailsCloseTicket,
	}
	workflowEmailsCmd.AddCommand(editWorkflowEmailsCloseTicketCmd)
}

func getTenants(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	responsibleTeamID, _ := cmd.Flags().GetInt("responsible-team-id")
	sdmID, _ := cmd.Flags().GetInt("sdm-id")

	params := make(map[string]string)
	if name != "" {
		params["name"] = name
	}
	if responsibleTeamID != 0 {
		params["responsibleTeamId"] = fmt.Sprintf("%d", responsibleTeamID)
	}
	if sdmID != 0 {
		params["sdmId"] = fmt.Sprintf("%d", sdmID)
	}

	response, err := client.GetTenants(params)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func createTenant(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	phone, _ := cmd.Flags().GetString("phone")
	address, _ := cmd.Flags().GetString("address")
	postalCode, _ := cmd.Flags().GetString("postal-code")
	city, _ := cmd.Flags().GetString("city")
	country, _ := cmd.Flags().GetString("country")
	responsibleTeam, _ := cmd.Flags().GetInt("responsible-team")
	contact, _ := cmd.Flags().GetInt("contact")
	watchers, _ := cmd.Flags().GetStringSlice("watchers")
	isEnabled, _ := cmd.Flags().GetBool("is-enabled")
	cloudTempleID, _ := cmd.Flags().GetString("cloud-temple-id")

	tenantData := map[string]interface{}{
		"name":            name,
		"phone":           phone,
		"address":         address,
		"postalCode":      postalCode,
		"city":            city,
		"country":         country,
		"responsibleTeam": responsibleTeam,
		"contact":         contact,
		"watchers":        watchers,
		"isEnabled":       isEnabled,
		"cloudTempleId":   cloudTempleID,
	}

	response, err := client.CreateTenant(tenantData)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getTenantDetails(cmd *cobra.Command, args []string) error {
	response, err := client.GetTenantDetails(args[0])
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getTenantContacts(cmd *cobra.Command, args []string) error {
	response, err := client.GetTenantContacts(args[0])
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func requestTenantDeletion(cmd *cobra.Command, args []string) error {
	delete, _ := cmd.Flags().GetBool("delete")
	response, err := client.RequestTenantDeletion(args[0], delete)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func listTenantSSHKeys(cmd *cobra.Command, args []string) error {
	response, err := client.GetTenantSSHKeys(args[0])
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func generateTenantSSHKey(cmd *cobra.Command, args []string) error {
	comment, _ := cmd.Flags().GetString("comment")
	isActive, _ := cmd.Flags().GetBool("is-active")

	keyData := map[string]interface{}{
		"comment":  comment,
		"isActive": isActive,
	}

	response, err := client.GenerateTenantSSHKey(args[0], keyData)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func deleteTenantSSHKey(cmd *cobra.Command, args []string) error {
	response, err := client.DeleteTenantSSHKey(args[0])
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func updateTenantSSHKey(cmd *cobra.Command, args []string) error {
	isActive, _ := cmd.Flags().GetBool("is-active")

	keyData := map[string]interface{}{
		"isActive": isActive,
	}

	response, err := client.UpdateTenantSSHKey(args[0], keyData)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getTenantWorkflowEmails(cmd *cobra.Command, args []string) error {
	response, err := client.GetTenantWorkflowEmails(args[0])
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func editTenantWorkflowEmailsGeneralities(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	from, _ := cmd.Flags().GetString("from")

	data := map[string]interface{}{
		"format": format,
		"from":   from,
	}

	response, err := client.EditTenantWorkflowEmailsGeneralities(args[0], data)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func editTenantWorkflowEmailsCreateTicket(cmd *cobra.Command, args []string) error {
	// Add necessary flags and implement the function
	return fmt.Errorf("Not implemented")
}

func editTenantWorkflowEmailsUpdateTicket(cmd *cobra.Command, args []string) error {
	// Add necessary flags and implement the function
	return fmt.Errorf("Not implemented")
}

func editTenantWorkflowEmailsValidationClientTicket(cmd *cobra.Command, args []string) error {
	// Add necessary flags and implement the function
	return fmt.Errorf("Not implemented")
}

func editTenantWorkflowEmailsCloseTicket(cmd *cobra.Command, args []string) error {
	// Add necessary flags and implement the function
	return fmt.Errorf("Not implemented")
}
