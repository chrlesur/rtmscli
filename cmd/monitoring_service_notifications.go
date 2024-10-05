package cmd

import (
	"fmt"

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
	}
	getServiceNotificationsCmd.Flags().Bool("attach", false, "List only notifications attached to a ticket or not")
	updateListCommand(getServiceNotificationsCmd, "/monitoringServices/{id}/notifications", func() map[string]string {
		params := make(map[string]string)
		attach, _ := getServiceNotificationsCmd.Flags().GetBool("attach")
		if attach {
			params["attach"] = "true"
		}
		return params
	})
	monitoringServiceNotificationsCmd.AddCommand(getServiceNotificationsCmd)

	// Get all notifications
	getAllNotificationsCmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of all notifications",
	}
	getAllNotificationsCmd.Flags().Bool("attach", false, "List only notifications attached to a ticket or not")
	getAllNotificationsCmd.Flags().IntSlice("staffs", nil, "Filter by staff identifiers")
	getAllNotificationsCmd.Flags().IntSlice("perimeters", nil, "Filter by perimeter identifiers")
	updateListCommand(getAllNotificationsCmd, "/monitoringServices/notifications", func() map[string]string {
		params := make(map[string]string)
		params["cloudTempleId"] = cloudTempleID
		attach, _ := getAllNotificationsCmd.Flags().GetBool("attach")
		if attach {
			params["attach"] = "true"
		}
		staffs, _ := getAllNotificationsCmd.Flags().GetIntSlice("staffs")
		if len(staffs) > 0 {
			params["staffs[]"] = intSliceToString(staffs)
		}
		perimeters, _ := getAllNotificationsCmd.Flags().GetIntSlice("perimeters")
		if len(perimeters) > 0 {
			params["perimeters[]"] = intSliceToString(perimeters)
		}
		return params
	})
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
}

func createNotification(cmd *cobra.Command, args []string) error {
	serviceID, _ := cmd.Flags().GetInt("service-id")
	state, _ := cmd.Flags().GetString("state")
	content, _ := cmd.Flags().GetString("content")
	subject, _ := cmd.Flags().GetString("subject")
	format, _ := cmd.Flags().GetString("format")

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
	formattedOutput, err := formatOutput(response, format)
	if err != nil {
		return err
	}
	fmt.Println(formattedOutput)
	return nil
}

func getNotificationDetails(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	response, err := client.GetNotificationDetails(args[0])
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

func getTicketSuggestions(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	response, err := client.GetTicketSuggestions(args[0])
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

func attachNotificationToTicket(cmd *cobra.Command, args []string) error {
	ticketID, _ := cmd.Flags().GetInt("ticket-id")
	format, _ := cmd.Flags().GetString("format")
	response, err := client.AttachNotificationToTicket(args[0], ticketID)
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

func detachNotificationFromTicket(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	response, err := client.DetachNotificationFromTicket(args[0])
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