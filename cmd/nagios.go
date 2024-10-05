package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var nagiosCmd = &cobra.Command{
	Use:   "nagios",
	Short: "Manage Nagios commands and plugins",
	Long:  `Manage Nagios commands and plugins, including listing commands, time periods, and validating plugins.`,
}

func init() {
	rootCmd.AddCommand(nagiosCmd)

	// Get Nagios commands
	getNagiosCommandsCmd := &cobra.Command{
		Use:   "commands",
		Short: "Get a list of Nagios commands",
		RunE:  getNagiosCommands,
	}
	getNagiosCommandsCmd.Flags().String("name", "", "Filter by name")
	nagiosCmd.AddCommand(getNagiosCommandsCmd)

	// Get Nagios commands time periods
	getNagiosCommandsTimePeriodsCmd := &cobra.Command{
		Use:   "time-periods",
		Short: "Get Nagios commands execution time periods list",
		RunE:  getNagiosCommandsTimePeriods,
	}
	getNagiosCommandsTimePeriodsCmd.Flags().String("name", "", "Filter timeperiod by name")
	getNagiosCommandsTimePeriodsCmd.Flags().String("alias", "", "Filter timeperiods by alias")
	nagiosCmd.AddCommand(getNagiosCommandsTimePeriodsCmd)

	// Validate Nagios plugin package
	validateNagiosPluginPackageCmd := &cobra.Command{
		Use:   "validate-plugin",
		Short: "Validate a Nagios plugin package",
		RunE:  validateNagiosPluginPackage,
	}
	validateNagiosPluginPackageCmd.Flags().String("package", "", "JSON string of the Nagios Plugin package's composer.json")
	validateNagiosPluginPackageCmd.MarkFlagRequired("package")
	nagiosCmd.AddCommand(validateNagiosPluginPackageCmd)

	// Update Nagios commands
	updateNagiosCommandsCmd := &cobra.Command{
		Use:   "update-commands",
		Short: "Update Nagios commands",
		RunE:  updateNagiosCommands,
	}
	nagiosCmd.AddCommand(updateNagiosCommandsCmd)
}

func getNagiosCommands(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	format, _ := cmd.Flags().GetString("format")

	params := make(map[string]string)
	if name != "" {
		params["name"] = name
	}

	response, err := client.GetNagiosCommands(params)
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

func getNagiosCommandsTimePeriods(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	alias, _ := cmd.Flags().GetString("alias")
	format, _ := cmd.Flags().GetString("format")

	params := make(map[string]string)
	if name != "" {
		params["name"] = name
	}
	if alias != "" {
		params["alias"] = alias
	}

	response, err := client.GetNagiosCommandsTimePeriods(cloudTempleID, params)
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

func validateNagiosPluginPackage(cmd *cobra.Command, args []string) error {
	packageJSON, _ := cmd.Flags().GetString("package")
	format, _ := cmd.Flags().GetString("format")

	var packageData map[string]interface{}
	err := json.Unmarshal([]byte(packageJSON), &packageData)
	if err != nil {
		return fmt.Errorf("invalid JSON: %v", err)
	}

	response, err := client.ValidateNagiosPluginPackage(packageData)
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

func updateNagiosCommands(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	response, err := client.UpdateNagiosCommands()
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