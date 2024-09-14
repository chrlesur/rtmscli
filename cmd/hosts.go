package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var hostsCmd = &cobra.Command{
	Use:   "hosts",
	Short: "Manage hosts",
	Long:  `Manage hosts, including listing, creating, updating, and deleting hosts, as well as managing their services and monitoring.`,
}

func init() {
	rootCmd.AddCommand(hostsCmd)

	// Get hosts
	getHostsCmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of Hosts",
		RunE:  getHosts,
	}
	getHostsCmd.Flags().String("name", "", "Filter hosts by name")
	getHostsCmd.Flags().StringSlice("status", nil, "Filter by hosts status (UP, DOWN, PENDING, UNREACHABLE)")
	getHostsCmd.Flags().Bool("is-monitored", false, "Filter by monitored hosts")
	hostsCmd.AddCommand(getHostsCmd)

	// Create host
	createHostCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new Host",
		RunE:  createHost,
	}
	createHostCmd.Flags().String("name", "", "Host name")
	createHostCmd.Flags().String("address", "", "Host monitoring ip address")
	createHostCmd.MarkFlagRequired("name")
	createHostCmd.MarkFlagRequired("address")
	hostsCmd.AddCommand(createHostCmd)

	// Get host details
	getHostDetailsCmd := &cobra.Command{
		Use:   "details [id]",
		Short: "Get Host details",
		Args:  cobra.ExactArgs(1),
		RunE:  getHostDetails,
	}
	hostsCmd.AddCommand(getHostDetailsCmd)

	// Remove host
	removeHostCmd := &cobra.Command{
		Use:   "remove [id]",
		Short: "Remove Host",
		Args:  cobra.ExactArgs(1),
		RunE:  removeHost,
	}
	hostsCmd.AddCommand(removeHostCmd)

	// Update host
	updateHostCmd := &cobra.Command{
		Use:   "update [id]",
		Short: "Update a Host",
		Args:  cobra.ExactArgs(1),
		RunE:  updateHost,
	}
	updateHostCmd.Flags().String("name", "", "Host name")
	updateHostCmd.Flags().String("address", "", "Host monitoring ip address")
	hostsCmd.AddCommand(updateHostCmd)

	// Get host services
	getHostServicesCmd := &cobra.Command{
		Use:   "services [id]",
		Short: "Get Host services",
		Args:  cobra.ExactArgs(1),
		RunE:  getHostServices,
	}
	hostsCmd.AddCommand(getHostServicesCmd)

	// Update host tags
	updateHostTagsCmd := &cobra.Command{
		Use:   "update-tags [id]",
		Short: "Update Host tags",
		Args:  cobra.ExactArgs(1),
		RunE:  updateHostTags,
	}
	updateHostTagsCmd.Flags().IntSlice("tags", nil, "List of tag IDs")
	updateHostTagsCmd.MarkFlagRequired("tags")
	hostsCmd.AddCommand(updateHostTagsCmd)

	// Switch host monitoring
	switchHostMonitoringCmd := &cobra.Command{
		Use:   "switch-monitoring [id]",
		Short: "Disable/enable monitoring for all or specific host's services",
		Args:  cobra.ExactArgs(1),
		RunE:  switchHostMonitoring,
	}
	switchHostMonitoringCmd.Flags().Bool("enable", false, "Enable or disable monitoring")
	switchHostMonitoringCmd.Flags().IntSlice("services", nil, "List of service IDs")
	switchHostMonitoringCmd.MarkFlagRequired("enable")
	hostsCmd.AddCommand(switchHostMonitoringCmd)

	// Switch host monitoring notifications
	switchHostMonitoringNotificationsCmd := &cobra.Command{
		Use:   "switch-notifications [id]",
		Short: "Disable/enable monitoring notifications for all or specific host's services",
		Args:  cobra.ExactArgs(1),
		RunE:  switchHostMonitoringNotifications,
	}
	switchHostMonitoringNotificationsCmd.Flags().Bool("enable", false, "Enable or disable notifications")
	switchHostMonitoringNotificationsCmd.Flags().IntSlice("services", nil, "List of service IDs")
	switchHostMonitoringNotificationsCmd.MarkFlagRequired("enable")
	hostsCmd.AddCommand(switchHostMonitoringNotificationsCmd)

	// Get hosts stats
	getHostsStatsCmd := &cobra.Command{
		Use:   "stats",
		Short: "Get hosts status stats",
		RunE:  getHostsStats,
	}
	hostsCmd.AddCommand(getHostsStatsCmd)
}

func getHosts(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	status, _ := cmd.Flags().GetStringSlice("status")
	isMonitored, _ := cmd.Flags().GetBool("is-monitored")

	params := make(map[string]string)
	if name != "" {
		params["name"] = name
	}
	if len(status) > 0 {
		params["status[]"] = fmt.Sprintf("[%s]", strconv.Quote(status[0]))
	}
	if isMonitored {
		params["isMonitored"] = "true"
	}

	response, err := client.GetHosts(cloudTempleID, params)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func createHost(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	address, _ := cmd.Flags().GetString("address")

	hostData := map[string]interface{}{
		"name":    name,
		"address": address,
	}

	response, err := client.CreateHost(cloudTempleID, hostData)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getHostDetails(cmd *cobra.Command, args []string) error {
	response, err := client.GetHostDetails(args[0])
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func removeHost(cmd *cobra.Command, args []string) error {
	response, err := client.RemoveHost(args[0])
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func updateHost(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	address, _ := cmd.Flags().GetString("address")

	hostData := make(map[string]interface{})
	if name != "" {
		hostData["name"] = name
	}
	if address != "" {
		hostData["address"] = address
	}

	response, err := client.UpdateHost(args[0], hostData)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getHostServices(cmd *cobra.Command, args []string) error {
	response, err := client.GetHostServices(args[0], nil)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func updateHostTags(cmd *cobra.Command, args []string) error {
	tags, _ := cmd.Flags().GetIntSlice("tags")
	response, err := client.UpdateHostTags(args[0], tags)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func switchHostMonitoring(cmd *cobra.Command, args []string) error {
	enable, _ := cmd.Flags().GetBool("enable")
	services, _ := cmd.Flags().GetIntSlice("services")
	response, err := client.SwitchHostMonitoring(args[0], enable, services)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func switchHostMonitoringNotifications(cmd *cobra.Command, args []string) error {
	enable, _ := cmd.Flags().GetBool("enable")
	services, _ := cmd.Flags().GetIntSlice("services")
	response, err := client.SwitchHostMonitoringNotifications(args[0], enable, services)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getHostsStats(cmd *cobra.Command, args []string) error {
	response, err := client.GetHostsStats(cloudTempleID)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}
