package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var appliancesCmd = &cobra.Command{
	Use:   "appliances",
	Short: "Manage appliances",
	Long:  `Commands to manage and interact with appliances in the RTMS system.`,
}

func init() {
	rootCmd.AddCommand(appliancesCmd)

	// Subcommands
	appliancesCmd.AddCommand(getAppliancesCmd)
	appliancesCmd.AddCommand(getApplianceDetailsCmd)
	appliancesCmd.AddCommand(getApplianceServicesCmd)
	appliancesCmd.AddCommand(synchronizeApplianceCmd)
	appliancesCmd.AddCommand(getApplianceConfigurationCmd)
	appliancesCmd.AddCommand(getApplianceHealthCheckCmd)
	appliancesCmd.AddCommand(postApplianceHealthCheckCmd)
}

var getAppliancesCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of appliances",
	RunE: func(cmd *cobra.Command, args []string) error {
		cloudTempleID, _ := cmd.Flags().GetString("cloud-temple-id")
		response, err := client.GetAppliances(cloudTempleID)
		if err != nil {
			return err
		}
		formattedOutput, err := formatOutput(response)
		if err != nil {
			return err
		}
		fmt.Println(formattedOutput)
		return nil
	},
}

var getApplianceDetailsCmd = &cobra.Command{
	Use:   "details [id]",
	Short: "Get Appliance details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		response, err := client.GetApplianceDetails(args[0])
		if err != nil {
			return err
		}
		formattedOutput, err := formatOutput(response)
		if err != nil {
			return err
		}
		fmt.Println(formattedOutput)
		return nil
	},
}

var getApplianceServicesCmd = &cobra.Command{
	Use:   "services [id]",
	Short: "Get Appliance services",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		response, err := client.GetApplianceServices(args[0])
		if err != nil {
			return err
		}
		formattedOutput, err := formatOutput(response)
		if err != nil {
			return err
		}
		fmt.Println(formattedOutput)
		return nil
	},
}

var synchronizeApplianceCmd = &cobra.Command{
	Use:   "synchronize [id]",
	Short: "Synchronize Appliance",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		response, err := client.SynchronizeAppliance(args[0])
		if err != nil {
			return err
		}
		formattedOutput, err := formatOutput(response)
		if err != nil {
			return err
		}
		fmt.Println(formattedOutput)
		return nil
	},
}

var getApplianceConfigurationCmd = &cobra.Command{
	Use:   "configuration [id]",
	Short: "Get appliances configuration",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		applianceVersion, _ := cmd.Flags().GetString("appliance-version")
		pluginsPath, _ := cmd.Flags().GetString("plugins-path")
		response, err := client.GetApplianceConfiguration(args[0], applianceVersion, pluginsPath)
		if err != nil {
			return err
		}
		formattedOutput, err := formatOutput(response)
		if err != nil {
			return err
		}
		fmt.Println(formattedOutput)
		return nil
	},
}

var getApplianceHealthCheckCmd = &cobra.Command{
	Use:   "healthcheck [id]",
	Short: "Get a last heartbeat of an appliance",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		response, err := client.GetApplianceHealthCheck(args[0])
		if err != nil {
			return err
		}
		formattedOutput, err := formatOutput(response)
		if err != nil {
			return err
		}
		fmt.Println(formattedOutput)
		return nil
	},
}

var postApplianceHealthCheckCmd = &cobra.Command{
	Use:   "post-healthcheck [id]",
	Short: "Posts an appliance heartbeat",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		applianceVersion, _ := cmd.Flags().GetString("appliance-version")
		nagiosOperatingState, _ := cmd.Flags().GetString("nagios-operating-state")
		details, _ := cmd.Flags().GetString("details")

		healthCheck := map[string]interface{}{
			"applianceVersion":     applianceVersion,
			"nagiosOperatingState": nagiosOperatingState,
			"details":              details,
		}

		response, err := client.PostApplianceHealthCheck(args[0], healthCheck)
		if err != nil {
			return err
		}
		formattedOutput, err := formatOutput(response)
		if err != nil {
			return err
		}
		fmt.Println(formattedOutput)
		return nil
	},
}

func init() {
	getApplianceConfigurationCmd.Flags().String("appliance-version", "", "Appliance version")
	getApplianceConfigurationCmd.Flags().String("plugins-path", "", "Absolute path to the plugins installation directory on the appliance")
	getApplianceConfigurationCmd.MarkFlagRequired("appliance-version")
	getApplianceConfigurationCmd.MarkFlagRequired("plugins-path")

	postApplianceHealthCheckCmd.Flags().String("appliance-version", "", "Appliance version")
	postApplianceHealthCheckCmd.Flags().String("nagios-operating-state", "", "Nagios operating state (OK, WARNING, CRITICAL)")
	postApplianceHealthCheckCmd.Flags().String("details", "", "Any details to explain the current operating state")
	postApplianceHealthCheckCmd.MarkFlagRequired("appliance-version")
	postApplianceHealthCheckCmd.MarkFlagRequired("nagios-operating-state")
}
