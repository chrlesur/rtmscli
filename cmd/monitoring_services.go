package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var monitoringServicesCmd = &cobra.Command{
	Use:   "monitoring-services",
	Short: "Manage monitoring services",
	Long:  `Manage monitoring services, including listing, creating, updating, and deleting services, as well as managing templates and viewing statistics.`,
}

func init() {
	rootCmd.AddCommand(monitoringServicesCmd)

	// Get monitoring services
	getMonitoringServicesCmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of monitoring services",
		RunE:  getMonitoringServices,
	}
	getMonitoringServicesCmd.Flags().String("name", "", "Filter services by name")
	getMonitoringServicesCmd.Flags().StringSlice("status", nil, "Filter services by status")
	getMonitoringServicesCmd.Flags().StringSlice("impact", nil, "Filter services by impact")
	monitoringServicesCmd.AddCommand(getMonitoringServicesCmd)

	// Create monitoring service
	createMonitoringServiceCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a monitoring service",
		RunE:  createMonitoringService,
	}
	createMonitoringServiceCmd.Flags().String("name", "", "Monitoring service name")
	createMonitoringServiceCmd.Flags().Int("appliance", 0, "Appliance ID")
	createMonitoringServiceCmd.Flags().Int("host", 0, "Host ID")
	createMonitoringServiceCmd.Flags().Int("template", 0, "Template ID")
	createMonitoringServiceCmd.MarkFlagRequired("name")
	createMonitoringServiceCmd.MarkFlagRequired("appliance")
	createMonitoringServiceCmd.MarkFlagRequired("host")
	createMonitoringServiceCmd.MarkFlagRequired("template")
	monitoringServicesCmd.AddCommand(createMonitoringServiceCmd)

	// Get monitoring service details
	getMonitoringServiceDetailsCmd := &cobra.Command{
		Use:   "details [id]",
		Short: "Get monitoring service details",
		Args:  cobra.ExactArgs(1),
		RunE:  getMonitoringServiceDetails,
	}
	monitoringServicesCmd.AddCommand(getMonitoringServiceDetailsCmd)

	// Remove monitoring service
	removeMonitoringServiceCmd := &cobra.Command{
		Use:   "remove [id]",
		Short: "Remove monitoring service",
		Args:  cobra.ExactArgs(1),
		RunE:  removeMonitoringService,
	}
	monitoringServicesCmd.AddCommand(removeMonitoringServiceCmd)

	// Update monitoring service
	updateMonitoringServiceCmd := &cobra.Command{
		Use:   "update [id]",
		Short: "Update a monitoring service",
		Args:  cobra.ExactArgs(1),
		RunE:  updateMonitoringService,
	}
	updateMonitoringServiceCmd.Flags().String("name", "", "Monitoring service name")
	updateMonitoringServiceCmd.Flags().Int("appliance", 0, "Appliance ID")
	updateMonitoringServiceCmd.Flags().Int("host", 0, "Host ID")
	updateMonitoringServiceCmd.Flags().Int("template", 0, "Template ID")
	monitoringServicesCmd.AddCommand(updateMonitoringServiceCmd)

	// Get monitoring service templates
	getMonitoringServiceTemplatesCmd := &cobra.Command{
		Use:   "templates",
		Short: "Get monitoring services templates list",
		RunE:  getMonitoringServiceTemplates,
	}
	getMonitoringServiceTemplatesCmd.Flags().String("name", "", "Filter template by name")
	getMonitoringServiceTemplatesCmd.Flags().StringSlice("impact", nil, "Filter templates by impact")
	monitoringServicesCmd.AddCommand(getMonitoringServiceTemplatesCmd)

	// Get monitoring services stats
	getMonitoringServicesStatsCmd := &cobra.Command{
		Use:   "stats",
		Short: "Get monitoring services status and impact stats",
		RunE:  getMonitoringServicesStats,
	}
	getMonitoringServicesStatsCmd.Flags().Int("host-id", 0, "Show stats of filtered monitoring services by host")
	getMonitoringServicesStatsCmd.Flags().Int("appliance-id", 0, "Show stats of filtered monitoring services by appliance")
	monitoringServicesCmd.AddCommand(getMonitoringServicesStatsCmd)
}

func getMonitoringServices(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	status, _ := cmd.Flags().GetStringSlice("status")
	impact, _ := cmd.Flags().GetStringSlice("impact")

	params := make(map[string]string)
	if name != "" {
		params["name"] = name
	}
	if len(status) > 0 {
		params["status[]"] = fmt.Sprintf("[%s]", strings.Join(status, ","))
	}
	if len(impact) > 0 {
		params["impact[]"] = fmt.Sprintf("[%s]", strings.Join(impact, ","))
	}

	response, err := client.GetMonitoringServices(cloudTempleID, params)
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

func createMonitoringService(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	appliance, _ := cmd.Flags().GetInt("appliance")
	host, _ := cmd.Flags().GetInt("host")
	template, _ := cmd.Flags().GetInt("template")

	serviceData := map[string]interface{}{
		"name":      name,
		"appliance": appliance,
		"host":      host,
		"template":  template,
	}

	response, err := client.CreateMonitoringService(cloudTempleID, serviceData)
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

func getMonitoringServiceDetails(cmd *cobra.Command, args []string) error {
	response, err := client.GetMonitoringServiceDetails(args[0])
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

func removeMonitoringService(cmd *cobra.Command, args []string) error {
	response, err := client.RemoveMonitoringService(args[0])
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

func updateMonitoringService(cmd *cobra.Command, args []string) error {
	serviceData := make(map[string]interface{})

	if name, _ := cmd.Flags().GetString("name"); name != "" {
		serviceData["name"] = name
	}
	if appliance, _ := cmd.Flags().GetInt("appliance"); appliance != 0 {
		serviceData["appliance"] = appliance
	}
	if host, _ := cmd.Flags().GetInt("host"); host != 0 {
		serviceData["host"] = host
	}
	if template, _ := cmd.Flags().GetInt("template"); template != 0 {
		serviceData["template"] = template
	}

	response, err := client.UpdateMonitoringService(args[0], serviceData)
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

func getMonitoringServiceTemplates(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	impact, _ := cmd.Flags().GetStringSlice("impact")

	params := make(map[string]string)
	if name != "" {
		params["name"] = name
	}
	if len(impact) > 0 {
		params["impact"] = fmt.Sprintf("[%s]", strings.Join(impact, ","))
	}

	response, err := client.GetMonitoringServiceTemplates(params)
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

func getMonitoringServicesStats(cmd *cobra.Command, args []string) error {
	hostID, _ := cmd.Flags().GetInt("host-id")
	applianceID, _ := cmd.Flags().GetInt("appliance-id")

	params := make(map[string]string)
	if hostID != 0 {
		params["hostId"] = fmt.Sprintf("%d", hostID)
	}
	if applianceID != 0 {
		params["applianceId"] = fmt.Sprintf("%d", applianceID)
	}

	response, err := client.GetMonitoringServicesStats(cloudTempleID, params)
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
