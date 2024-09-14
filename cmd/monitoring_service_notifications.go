package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var monitoringServiceNotificationsCmd = &cobra.Command{
	Use:   "notifications",
	Short: "Manage monitoring service notifications",
	Long:  `Manage notifications for monitoring services, including listing, creating, attaching, and detaching notifications.`,
}

func init() {
	monitoringServicesCmd.AddCommand(monitoringServiceNotificationsCmd)

	// Get service notifications
	getServiceNotificationsCmd := &cobra.Command{
		Use:   "list-service [service-id]",
		Short: "Get a list of notifications of specific service",
		Args:  cobra.ExactArgs(1),
		RunE:  getServiceNotifications,
	}
	getServiceNotificationsCmd.Flags().Bool("attach", false, "List only notifications attached to a ticket or not")
	monitoringServiceNotificationsCmd.AddCommand(getServiceNotificationsCmd)

	// Get all notifications
	getAllNotificationsCmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all notifications",
		RunE:  getAllNotifications,
	}
	getAllNotificationsCmd.Flags().Bool("attach", false, "List only notifications attached to a ticket or not")
	getAllNotificationsCmd.Flags().IntSlice("staffs", nil, "Filter by staff identifiers")
	getAllNotificationsCmd.Flags().IntSlice("perimeters", nil, "Filter by perimeter identifiers")
	monitoringServiceNotificationsCmd.AddCommand(getAllNotificationsCmd)

	// Create notification
	createNotificationCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new notification",
		RunE:  createNotification,
	}
	createNotificationCmd.Flags().Int("service-id", 0, "Monitoring service ID")
	createNotificationCmd.Flags().String("state", "", "State of monitoring service (OK, WARNING, CRITICAL, UNKNOWN)")
	createNotificationCmd.Flags().String("content", "", "Content of the notification")
	createNotificationCmd.Flags().String("subject", "", "Subject that will be sent by email/sms")
	createNotificationCmd.MarkFlagRequired("service-id")
	createNotificationCmd.MarkFlagRequired("state")
	createNotificationCmd.MarkFlagRequired("content")
	createNotificationCmd.MarkFlagRequired("subject")
	monitoringServiceNotificationsCmd.AddCommand(createNotificationCmd)

	// Get notification details
	getNotificationDetailsCmd := &cobra.Command{
		Use:   "details [id]",
		Short: "Get notification details",
		Args:  cobra.ExactArgs(1),
		RunE:  getNotificationDetails,
	}
	monitoringServiceNotificationsCmd.AddCommand(getNotificationDetailsCmd)

	// Get ticket suggestions
	getTicketSuggestionsCmd := &cobra.Command{
		Use:   "suggest [id]",
		Short: "Get a ticket suggestion against a notification",
		Args:  cobra.ExactArgs(1),
		RunE:  getTicketSuggestions,
	}
	monitoringServiceNotificationsCmd.AddCommand(getTicketSuggestionsCmd)

	// Attach notification to ticket
	attachNotificationCmd := &cobra.Command{
		Use:   "attach [id]",
		Short: "Attach notification to a ticket",
		Args:  cobra.ExactArgs(1),
		RunE:  attachNotificationToTicket,
	}
	attachNotificationCmd.Flags().Int("ticket-id", 0, "Ticket ID to attach the notification to")
	attachNotificationCmd.MarkFlagRequired("ticket-id")
	monitoringServiceNotificationsCmd.AddCommand(attachNotificationCmd)

	// Detach notification from ticket
	detachNotificationCmd := &cobra.Command{
		Use:   "detach [id]",
		Short: "Detach notification from a ticket",
		Args:  cobra.ExactArgs(1),
		RunE:  detachNotificationFromTicket,
	}
	monitoringServiceNotificationsCmd.AddCommand(detachNotificationCmd)

	// Perimeters subcommand
	perimetersCmd := &cobra.Command{
		Use:   "perimeters",
		Short: "Manage notification perimeters",
		Long:  `Manage notification perimeters, including listing, getting details, and updating perimeters.`,
	}
	monitoringServiceNotificationsCmd.AddCommand(perimetersCmd)

	// Get notification perimeters
	getPerimetersCmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of notification perimeters",
		RunE:  getNotificationPerimeters,
	}
	getPerimetersCmd.Flags().String("name", "", "Filter by perimeter name")
	perimetersCmd.AddCommand(getPerimetersCmd)

	// Get notification perimeter details
	getPerimeterDetailsCmd := &cobra.Command{
		Use:   "details [id]",
		Short: "Get notification perimeter details",
		Args:  cobra.ExactArgs(1),
		RunE:  getNotificationPerimeterDetails,
	}
	perimetersCmd.AddCommand(getPerimeterDetailsCmd)

	// Update notification perimeter
	updatePerimeterCmd := &cobra.Command{
		Use:   "update [id]",
		Short: "Update a notification perimeter",
		Args:  cobra.ExactArgs(1),
		RunE:  updateNotificationPerimeter,
	}
	updatePerimeterCmd.Flags().String("name", "", "New perimeter name")
	updatePerimeterCmd.Flags().Int("tenant", 0, "Tenant scoop identifier")
	updatePerimeterCmd.Flags().IntSlice("add-services", nil, "List of monitoring service IDs to add to the perimeter")
	updatePerimeterCmd.Flags().IntSlice("remove-services", nil, "List of monitoring service IDs to remove from the perimeter")
	perimetersCmd.AddCommand(updatePerimeterCmd)
}

func getServiceNotifications(cmd *cobra.Command, args []string) error {
	serviceID := args[0]
	attach, _ := cmd.Flags().GetBool("attach")

	params := make(map[string]string)
	if attach {
		params["attach"] = "true"
	}

	response, err := client.GetServiceNotifications(serviceID, params)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getAllNotifications(cmd *cobra.Command, args []string) error {
	attach, _ := cmd.Flags().GetBool("attach")
	staffs, _ := cmd.Flags().GetIntSlice("staffs")
	perimeters, _ := cmd.Flags().GetIntSlice("perimeters")

	params := make(map[string]string)
	if attach {
		params["attach"] = "true"
	}
	if len(staffs) > 0 {
		params["staffs[]"] = intSliceToString(staffs)
	}
	if len(perimeters) > 0 {
		params["perimeters[]"] = intSliceToString(perimeters)
	}

	response, err := client.GetAllNotifications(cloudTempleID, params)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func createNotification(cmd *cobra.Command, args []string) error {
	serviceID, _ := cmd.Flags().GetInt("service-id")
	state, _ := cmd.Flags().GetString("state")
	content, _ := cmd.Flags().GetString("content")
	subject, _ := cmd.Flags().GetString("subject")

	notificationData := map[string]interface{}{
		"monitoringServiceId": serviceID,
		"state":               state,
		"content":             content,
		"subject":             subject,
	}

	response, err := client.CreateNotification(notificationData)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getNotificationDetails(cmd *cobra.Command, args []string) error {
	response, err := client.GetNotificationDetails(args[0])
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getTicketSuggestions(cmd *cobra.Command, args []string) error {
	response, err := client.GetTicketSuggestions(args[0])
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func attachNotificationToTicket(cmd *cobra.Command, args []string) error {
	ticketID, _ := cmd.Flags().GetInt("ticket-id")
	response, err := client.AttachNotificationToTicket(args[0], ticketID)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func detachNotificationFromTicket(cmd *cobra.Command, args []string) error {
	response, err := client.DetachNotificationFromTicket(args[0])
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func intSliceToString(slice []int) string {
	result := ""
	for i, v := range slice {
		if i > 0 {
			result += ","
		}
		result += strconv.Itoa(v)
	}
	return result
}

func getNotificationPerimeters(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")

	params := make(map[string]string)
	if name != "" {
		params["name"] = name
	}

	response, err := client.GetNotificationPerimeters(cloudTempleID, params)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getNotificationPerimeterDetails(cmd *cobra.Command, args []string) error {
	response, err := client.GetNotificationPerimeter(args[0])
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func updateNotificationPerimeter(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	tenant, _ := cmd.Flags().GetInt("tenant")
	addServices, _ := cmd.Flags().GetIntSlice("add-services")
	removeServices, _ := cmd.Flags().GetIntSlice("remove-services")

	perimeterData := make(map[string]interface{})
	if name != "" {
		perimeterData["name"] = name
	}
	if tenant != 0 {
		perimeterData["tenant"] = tenant
	}
	if len(addServices) > 0 {
		perimeterData["addMonitoringServices"] = addServices
	}
	if len(removeServices) > 0 {
		perimeterData["removeMonitoringServices"] = removeServices
	}

	response, err := client.UpdateNotificationPerimeter(args[0], perimeterData)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}
