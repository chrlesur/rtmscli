package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var monitoringCmd = &cobra.Command{
	Use:   "monitoring",
	Short: "Check the health of RTMS services",
	Long:  `Check the health status of RTMS services and the SLA Calculator application.`,
}

func init() {
	rootCmd.AddCommand(monitoringCmd)

	// Check RTMS health
	checkRTMSHealthCmd := &cobra.Command{
		Use:   "health",
		Short: "Check if RTMS services are healthy",
		RunE:  checkRTMSHealth,
	}
	checkRTMSHealthCmd.Flags().IntSlice("integration-services", nil, "List of service identifiers used to test the delay of integration of monitoring results")
	checkRTMSHealthCmd.Flags().Int("integration-delay", 0, "Delay allowed in seconds to test the delay of integration of monitoring results")
	monitoringCmd.AddCommand(checkRTMSHealthCmd)

	// Check SLA Calculator health
	checkSLACalculatorHealthCmd := &cobra.Command{
		Use:   "sla-calculator",
		Short: "Check if the SLA Calculator app is healthy",
		RunE:  checkSLACalculatorHealth,
	}
	checkSLACalculatorHealthCmd.Flags().Int("update-delay", 0, "Delay allowed in seconds between the current time and the last update of a ticket's SLA")
	monitoringCmd.AddCommand(checkSLACalculatorHealthCmd)
}

func checkRTMSHealth(cmd *cobra.Command, args []string) error {
	integrationServices, _ := cmd.Flags().GetIntSlice("integration-services")
	integrationDelay, _ := cmd.Flags().GetInt("integration-delay")

	response, err := client.CheckRTMSHealth(integrationServices, integrationDelay)
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

func checkSLACalculatorHealth(cmd *cobra.Command, args []string) error {
	updateDelay, _ := cmd.Flags().GetInt("update-delay")

	response, err := client.CheckSLACalculatorHealth(updateDelay)
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
