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

	// Staffs subcommand
	staffsCmd := &cobra.Command{
		Use:   "staffs",
		Short: "Manage notification staffs",
		Long:  `Manage notification staffs, including listing and getting details.`,
	}
	monitoringServiceNotificationsCmd.AddCommand(staffsCmd)

	// Get notification staffs
	getStaffsCmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of notification staffs",
		RunE:  getNotificationStaffs,
	}
	getStaffsCmd.Flags().String("name", "", "Filter by staff name")
	staffsCmd.AddCommand(getStaffsCmd)

	// Get notification staff details
	getStaffDetailsCmd := &cobra.Command{
		Use:   "details [id]",
		Short: "Get notification staff details",
		Args:  cobra.ExactArgs(1),
		RunE:  getNotificationStaffDetails,
	}
	staffsCmd.AddCommand(getStaffDetailsCmd)

	// Time Periods subcommand
	timePeriodsCmd := &cobra.Command{
		Use:   "timeperiods",
		Short: "Manage notification time periods",
		Long:  `Manage notification time periods, including listing time periods.`,
	}
	monitoringServiceNotificationsCmd.AddCommand(timePeriodsCmd)

	// Get notification time periods
	getTimePeriodsCmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of notification time periods",
		RunE:  getNotificationTimePeriods,
	}
	timePeriodsCmd.AddCommand(getTimePeriodsCmd)

	// Time Period Stops subcommand
	timePeriodStopsCmd := &cobra.Command{
		Use:   "timeperiod-stops",
		Short: "Manage notification time period stops",
		Long:  `Manage notification time period stops, including listing, creating, getting details, and removing stops.`,
	}
	monitoringServiceNotificationsCmd.AddCommand(timePeriodStopsCmd)

	// Get notification time period stops
	getTimePeriodStopsCmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of monitoring services' notifications planned time periods stops",
		RunE:  getNotificationTimePeriodStops,
	}
	getTimePeriodStopsCmd.Flags().Bool("currently-active", false, "Filter time period stops by currently active and expired date ranges")
	getTimePeriodStopsCmd.Flags().String("name", "", "Filter time period stops by name")
	timePeriodStopsCmd.AddCommand(getTimePeriodStopsCmd)

	// Create notification time period stop
	createTimePeriodStopCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a notification time period stop",
		RunE:  createNotificationTimePeriodStop,
	}
	createTimePeriodStopCmd.Flags().String("name", "", "Name of the time period stop")
	createTimePeriodStopCmd.Flags().String("reason", "", "The reason of the time period stop creation")
	createTimePeriodStopCmd.Flags().String("start", "", "Date to start the time period stop (YYYY-MM-DD)")
	createTimePeriodStopCmd.Flags().String("end", "", "Date to end the time period stop (YYYY-MM-DD)")
	createTimePeriodStopCmd.Flags().IntSlice("services", nil, "List of monitoring service IDs to apply time period stop")
	createTimePeriodStopCmd.MarkFlagRequired("name")
	createTimePeriodStopCmd.MarkFlagRequired("end")
	createTimePeriodStopCmd.MarkFlagRequired("services")
	timePeriodStopsCmd.AddCommand(createTimePeriodStopCmd)

	// Get notification time period stop details
	getTimePeriodStopDetailsCmd := &cobra.Command{
		Use:   "details [id]",
		Short: "Get a notification time period stop details",
		Args:  cobra.ExactArgs(1),
		RunE:  getNotificationTimePeriodStopDetails,
	}
	timePeriodStopsCmd.AddCommand(getTimePeriodStopDetailsCmd)

	// Remove notification time period stop
	removeTimePeriodStopCmd := &cobra.Command{
		Use:   "remove [id]",
		Short: "Remove notification time period stop",
		Args:  cobra.ExactArgs(1),
		RunE:  removeNotificationTimePeriodStop,
	}
	timePeriodStopsCmd.AddCommand(removeTimePeriodStopCmd)

	// Triggers subcommand
	triggersCmd := &cobra.Command{
		Use:   "triggers",
		Short: "Manage notification triggers",
		Long:  `Manage notification triggers, including listing and getting details.`,
	}
	monitoringServiceNotificationsCmd.AddCommand(triggersCmd)

	// Get notification triggers
	getTriggersCmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of notification triggers",
		RunE:  getNotificationTriggers,
	}
	getTriggersCmd.Flags().String("name", "", "Filter by trigger name")
	triggersCmd.AddCommand(getTriggersCmd)

	// Get notification trigger details
	getTriggerDetailsCmd := &cobra.Command{
		Use:   "details [id]",
		Short: "Get notification trigger details",
		Args:  cobra.ExactArgs(1),
		RunE:  getNotificationTriggerDetails,
	}
	triggersCmd.AddCommand(getTriggerDetailsCmd)
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
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
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
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
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
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func getNotificationDetails(cmd *cobra.Command, args []string) error {
	response, err := client.GetNotificationDetails(args[0])
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func getTicketSuggestions(cmd *cobra.Command, args []string) error {
	response, err := client.GetTicketSuggestions(args[0])
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func attachNotificationToTicket(cmd *cobra.Command, args []string) error {
	ticketID, _ := cmd.Flags().GetInt("ticket-id")
	response, err := client.AttachNotificationToTicket(args[0], ticketID)
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func detachNotificationFromTicket(cmd *cobra.Command, args []string) error {
	response, err := client.DetachNotificationFromTicket(args[0])
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
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
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func getNotificationPerimeterDetails(cmd *cobra.Command, args []string) error {
	response, err := client.GetNotificationPerimeter(args[0])
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
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
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func getNotificationStaffs(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")

	params := make(map[string]string)
	if name != "" {
		params["name"] = name
	}

	response, err := client.GetNotificationStaffs(cloudTempleID, params)
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func getNotificationStaffDetails(cmd *cobra.Command, args []string) error {
	response, err := client.GetNotificationStaff(args[0])
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func getNotificationTimePeriods(cmd *cobra.Command, args []string) error {
	response, err := client.GetNotificationTimePeriods(cloudTempleID, nil)
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func getNotificationTimePeriodStops(cmd *cobra.Command, args []string) error {
	currentlyActive, _ := cmd.Flags().GetBool("currently-active")
	name, _ := cmd.Flags().GetString("name")

	params := make(map[string]string)
	if currentlyActive {
		params["currentlyActive"] = "true"
	}
	if name != "" {
		params["name"] = name
	}

	response, err := client.GetNotificationTimePeriodStops(cloudTempleID, params)
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func createNotificationTimePeriodStop(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	reason, _ := cmd.Flags().GetString("reason")
	start, _ := cmd.Flags().GetString("start")
	end, _ := cmd.Flags().GetString("end")
	services, _ := cmd.Flags().GetIntSlice("services")

	stopData := map[string]interface{}{
		"name":               name,
		"timePeriodEnd":      end,
		"monitoringServices": services,
	}
	if reason != "" {
		stopData["reason"] = reason
	}
	if start != "" {
		stopData["timePeriodStart"] = start
	}

	response, err := client.CreateNotificationTimePeriodStop(cloudTempleID, stopData)
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func getNotificationTimePeriodStopDetails(cmd *cobra.Command, args []string) error {
	response, err := client.GetNotificationTimePeriodStop(args[0])
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func removeNotificationTimePeriodStop(cmd *cobra.Command, args []string) error {
	response, err := client.RemoveNotificationTimePeriodStop(args[0])
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func getNotificationTriggers(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")

	params := make(map[string]string)
	if name != "" {
		params["name"] = name
	}

	response, err := client.GetNotificationTriggers(cloudTempleID, params)
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}

func getNotificationTriggerDetails(cmd *cobra.Command, args []string) error {
	response, err := client.GetNotificationTriggerDetails(args[0])
	if err != nil {
		return err
	}
	// Utilisation de formatOutput pour formater la réponse
	formattedOutput, err := formatOutput(response)
	if err != nil {
		return err
	}

	// Affichage de la réponse formatée
	fmt.Println(formattedOutput)
	return nil
}
