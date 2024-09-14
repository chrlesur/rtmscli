package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getAppliancesCmd)
	rootCmd.AddCommand(getApplianceDetailsCmd)
	rootCmd.AddCommand(getApplianceServicesCmd)
	rootCmd.AddCommand(synchronizeApplianceCmd)
	rootCmd.AddCommand(getApplianceConfigurationCmd)
	rootCmd.AddCommand(getApplianceHealthCheckCmd)
	rootCmd.AddCommand(postApplianceHealthCheckCmd)
}

var getAppliancesCmd = &cobra.Command{
	Use:   "get-appliances",
	Short: "Get a list of appliances",
	RunE: func(cmd *cobra.Command, args []string) error {
		cloudTempleID, _ := cmd.Flags().GetString("cloud-temple-id")
		response, err := client.GetAppliances(cloudTempleID)
		if err != nil {
			return err
		}
		fmt.Println(string(response))
		return nil
	},
}

var getApplianceDetailsCmd = &cobra.Command{
	Use:   "get-appliance-details [id]",
	Short: "Get Appliance details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		response, err := client.GetApplianceDetails(args[0])
		if err != nil {
			return err
		}
		fmt.Println(string(response))
		return nil
	},
}

var getApplianceServicesCmd = &cobra.Command{
	Use:   "get-appliance-services [id]",
	Short: "Get Appliance services",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		response, err := client.GetApplianceServices(args[0])
		if err != nil {
			return err
		}
		fmt.Println(string(response))
		return nil
	},
}

var synchronizeApplianceCmd = &cobra.Command{
	Use:   "synchronize-appliance [id]",
	Short: "Synchronize Appliance",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		response, err := client.SynchronizeAppliance(args[0])
		if err != nil {
			return err
		}
		fmt.Println(string(response))
		return nil
	},
}

var getApplianceConfigurationCmd = &cobra.Command{
	Use:   "get-appliance-configuration [id]",
	Short: "Get appliances configuration",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		applianceVersion, _ := cmd.Flags().GetString("appliance-version")
		pluginsPath, _ := cmd.Flags().GetString("plugins-path")
		response, err := client.GetApplianceConfiguration(args[0], applianceVersion, pluginsPath)
		if err != nil {
			return err
		}
		fmt.Println(string(response))
		return nil
	},
}

func init() {
	getApplianceConfigurationCmd.Flags().String("appliance-version", "", "Appliance version")
	getApplianceConfigurationCmd.Flags().String("plugins-path", "", "Absolute path to the plugins installation directory on the appliance")
	getApplianceConfigurationCmd.MarkFlagRequired("appliance-version")
	getApplianceConfigurationCmd.MarkFlagRequired("plugins-path")
}

var getApplianceHealthCheckCmd = &cobra.Command{
	Use:   "get-appliance-healthcheck [id]",
	Short: "Get a last heartbeat of an appliance",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		response, err := client.GetApplianceHealthCheck(args[0])
		if err != nil {
			return err
		}
		fmt.Println(string(response))
		return nil
	},
}

var postApplianceHealthCheckCmd = &cobra.Command{
	Use:   "post-appliance-healthcheck [id]",
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
		fmt.Println(string(response))
		return nil
	},
}

func init() {
	postApplianceHealthCheckCmd.Flags().String("appliance-version", "", "Appliance version")
	postApplianceHealthCheckCmd.Flags().String("nagios-operating-state", "", "Nagios operating state (OK, WARNING, CRITICAL)")
	postApplianceHealthCheckCmd.Flags().String("details", "", "Any details to explain the current operating state")
	postApplianceHealthCheckCmd.MarkFlagRequired("appliance-version")
	postApplianceHealthCheckCmd.MarkFlagRequired("nagios-operating-state")
}
