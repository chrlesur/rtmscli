package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var teamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "Manage teams",
	Long:  `Manage teams, including listing, creating, updating, and deleting teams.`,
}

func init() {
	rootCmd.AddCommand(teamsCmd)

	// Get teams
	getTeamsCmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of Teams",
	}

	updateListCommand(getTeamsCmd, "/teams", func() map[string]string {
		return map[string]string{
			"cloudTempleId": cloudTempleID,
		}
	})

	teamsCmd.AddCommand(getTeamsCmd)

	// Create team
	createTeamCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new Team",
		RunE:  createTeam,
	}
	createTeamCmd.Flags().String("name", "", "Team name")
	createTeamCmd.Flags().String("information", "", "Team information")
	createTeamCmd.Flags().StringSlice("contacts", nil, "Contact email addresses")
	createTeamCmd.Flags().IntSlice("members", nil, "Team member identifiers")
	createTeamCmd.MarkFlagRequired("name")
	teamsCmd.AddCommand(createTeamCmd)

	// Get default teams
	getDefaultTeamsCmd := &cobra.Command{
		Use:   "defaults",
		Short: "Get a list of all default Teams",
		RunE:  getDefaultTeams,
	}
	teamsCmd.AddCommand(getDefaultTeamsCmd)

	// Get team details
	getTeamDetailsCmd := &cobra.Command{
		Use:   "details [id]",
		Short: "Get Team details",
		Args:  cobra.ExactArgs(1),
		RunE:  getTeamDetails,
	}
	teamsCmd.AddCommand(getTeamDetailsCmd)

	// Remove team
	removeTeamCmd := &cobra.Command{
		Use:   "remove [id]",
		Short: "Remove team",
		Args:  cobra.ExactArgs(1),
		RunE:  removeTeam,
	}
	teamsCmd.AddCommand(removeTeamCmd)

	// Edit team
	editTeamCmd := &cobra.Command{
		Use:   "edit [id]",
		Short: "Edit Team",
		Args:  cobra.ExactArgs(1),
		RunE:  editTeam,
	}
	editTeamCmd.Flags().String("name", "", "New team name")
	editTeamCmd.Flags().String("information", "", "New information about the team")
	editTeamCmd.Flags().Int("tenant", 0, "New tenant")
	editTeamCmd.Flags().StringSlice("add-contacts", nil, "List of emails to add as team contacts")
	editTeamCmd.Flags().StringSlice("remove-contacts", nil, "List of emails to delete from team contacts")
	editTeamCmd.Flags().IntSlice("add-members", nil, "List of user identifiers to add as team members")
	editTeamCmd.Flags().IntSlice("remove-members", nil, "List of user identifiers to remove from team members")
	teamsCmd.AddCommand(editTeamCmd)
}

func createTeam(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	information, _ := cmd.Flags().GetString("information")
	contacts, _ := cmd.Flags().GetStringSlice("contacts")
	members, _ := cmd.Flags().GetIntSlice("members")
	format, _ := cmd.Flags().GetString("format")

	teamData := map[string]interface{}{
		"name": name,
	}
	if information != "" {
		teamData["information"] = information
	}
	if len(contacts) > 0 {
		teamData["contacts"] = contacts
	}
	if len(members) > 0 {
		teamData["members"] = members
	}

	response, err := client.CreateTeam(cloudTempleID, teamData)
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

func getDefaultTeams(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	response, err := client.GetDefaultTeams(nil)
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

func getTeamDetails(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	response, err := client.GetTeamDetails(args[0])
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

func removeTeam(cmd *cobra.Command, args []string) error {
	format, _ := cmd.Flags().GetString("format")
	response, err := client.RemoveTeam(args[0])
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

func editTeam(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	information, _ := cmd.Flags().GetString("information")
	tenant, _ := cmd.Flags().GetInt("tenant")
	addContacts, _ := cmd.Flags().GetStringSlice("add-contacts")
	removeContacts, _ := cmd.Flags().GetStringSlice("remove-contacts")
	addMembers, _ := cmd.Flags().GetIntSlice("add-members")
	removeMembers, _ := cmd.Flags().GetIntSlice("remove-members")
	format, _ := cmd.Flags().GetString("format")

	teamData := make(map[string]interface{})
	if name != "" {
		teamData["name"] = name
	}
	if information != "" {
		teamData["information"] = information
	}
	if tenant != 0 {
		teamData["tenant"] = tenant
	}
	if len(addContacts) > 0 {
		teamData["addContacts"] = addContacts
	}
	if len(removeContacts) > 0 {
		teamData["removeContacts"] = removeContacts
	}
	if len(addMembers) > 0 {
		teamData["addMembers"] = addMembers
	}
	if len(removeMembers) > 0 {
		teamData["removeMembers"] = removeMembers
	}

	response, err := client.EditTeam(args[0], teamData)
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
