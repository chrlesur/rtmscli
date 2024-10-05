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
	}
	getMetricHistoryCmd.Flags().String("start-date", "", "Start date timestamp or milliseconds of searched period")
	getMetricHistoryCmd.Flags().String("end-date", "", "End date timestamp or milliseconds of searched period")
	getMetricHistoryCmd.Flags().StringSlice("metric-name", nil, "List of metric names")
	getMetricHistoryCmd.Flags().String("version-order", "", "Version order: asc or desc")
	updateListCommand(getMetricHistoryCmd, "/monitoringServices/{id}/metricHistory", func() map[string]string {
		params := make(map[string]string)
		startDate, _ := getMetricHistoryCmd.Flags().GetString("start-date")
		if startDate != "" {
			params["startDate"] = startDate
		}
		endDate, _ := getMetricHistoryCmd.Flags().GetString("end-date")
		if endDate != "" {
			params["endDate"] = endDate
		}
		metricNames, _ := getMetricHistoryCmd.Flags().GetStringSlice("metric-name")
		if len(metricNames) > 0 {
			params["metricName[]"] = strings.Join(metricNames, ",")
		}
		versionOrder, _ := getMetricHistoryCmd.Flags().GetString("version-order")
		if versionOrder != "" {
			params["versionOrder"] = versionOrder
		}
		return params
	})
	monitoringServicePerformanceCmd.AddCommand(getMetricHistoryCmd)

	// Get graph configurations
	getGraphConfigurationsCmd := &cobra.Command{
		Use:   "graph-configurations [service-id]",
		Short: "Get a list of graph configurations for a given monitoring service",
		Args:  cobra.ExactArgs(1),
	}
	getGraphConfigurationsCmd.Flags().String("label", "", "Filter graph by a string contained in label field")
	updateListCommand(getGraphConfigurationsCmd, "/monitoringServices/{id}/graphs", func() map[string]string {
		params := make(map[string]string)
		label, _ := getGraphConfigurationsCmd.Flags().GetString("label")
		if label != "" {
			params["label"] = label
		}
		return params
	})
	monitoringServicePerformanceCmd.AddCommand(getGraphConfigurationsCmd)
}

func getMetricHistory(cmd *cobra.Command, args []string) error {
	serviceID := args[0]
	format, _ := cmd.Flags().GetString("format")

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

	dataChan, errChan := client.StreamData(fmt.Sprintf("/monitoringServices/%s/metricHistory", serviceID), params, batchSize)

	for item := range dataChan {
		formattedOutput, err := formatOutput(item, format)
		if err != nil {
			return err
		}
		fmt.Println(formattedOutput)
	}

	if err := <-errChan; err != nil {
		return fmt.Errorf("erreur lors de la récupération de l'historique des métriques : %w", err)
	}

	return nil
}

func getGraphConfigurations(cmd *cobra.Command, args []string) error {
	serviceID := args[0]
	format, _ := cmd.Flags().GetString("format")

	label, _ := cmd.Flags().GetString("label")

	params := make(map[string]string)
	if label != "" {
		params["label"] = label
	}

	dataChan, errChan := client.StreamData(fmt.Sprintf("/monitoringServices/%s/graphs", serviceID), params, batchSize)

	for item := range dataChan {
		formattedOutput, err := formatOutput(item, format)
		if err != nil {
			return err
		}
		fmt.Println(formattedOutput)
	}

	if err := <-errChan; err != nil {
		return fmt.Errorf("erreur lors de la récupération des configurations de graphiques : %w", err)
	}

	return nil
}