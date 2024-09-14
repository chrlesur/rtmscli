package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var monitoringServicePerformanceCmd = &cobra.Command{
	Use:   "performance",
	Short: "Manage monitoring service performance data",
	Long:  `Manage monitoring service performance data, including metric history and graph configurations.`,
}

func init() {
	monitoringServicesCmd.AddCommand(monitoringServicePerformanceCmd)

	// Get metric history
	getMetricHistoryCmd := &cobra.Command{
		Use:   "metric-history [service-id]",
		Short: "Get a list of metrics versions for a given monitoring service",
		Args:  cobra.ExactArgs(1),
		RunE:  getMetricHistory,
	}
	getMetricHistoryCmd.Flags().String("start-date", "", "Start date timestamp or milliseconds of searched period")
	getMetricHistoryCmd.Flags().String("end-date", "", "End date timestamp or milliseconds of searched period")
	getMetricHistoryCmd.Flags().StringSlice("metric-name", nil, "List of metric names")
	getMetricHistoryCmd.Flags().String("version-order", "", "Version order: asc or desc")
	monitoringServicePerformanceCmd.AddCommand(getMetricHistoryCmd)

	// Get graph configurations
	getGraphConfigurationsCmd := &cobra.Command{
		Use:   "graph-configurations [service-id]",
		Short: "Get a list of graph configurations for a given monitoring service",
		Args:  cobra.ExactArgs(1),
		RunE:  getGraphConfigurations,
	}
	getGraphConfigurationsCmd.Flags().String("label", "", "Filter graph by a string contained in label field")
	monitoringServicePerformanceCmd.AddCommand(getGraphConfigurationsCmd)
}

func getMetricHistory(cmd *cobra.Command, args []string) error {
	startDate, _ := cmd.Flags().GetString("start-date")
	endDate, _ := cmd.Flags().GetString("end-date")
	metricNames, _ := cmd.Flags().GetStringSlice("metric-name")
	versionOrder, _ := cmd.Flags().GetString("version-order")

	params := make(map[string]string)
	if startDate != "" {
		params["startDate"] = startDate
	}
	if endDate != "" {
		params["endDate"] = endDate
	}
	if len(metricNames) > 0 {
		params["metricName[]"] = strings.Join(metricNames, ",")
	}
	if versionOrder != "" {
		params["versionOrder"] = versionOrder
	}

	response, err := client.GetMetricHistory(args[0], params)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getGraphConfigurations(cmd *cobra.Command, args []string) error {
	label, _ := cmd.Flags().GetString("label")

	params := make(map[string]string)
	if label != "" {
		params["label"] = label
	}

	response, err := client.GetGraphConfigurations(args[0], params)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}
