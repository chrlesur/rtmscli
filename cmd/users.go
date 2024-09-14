package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Manage users",
	Long:  `Manage users, including listing, creating, updating, and viewing user details.`,
}

func init() {
	rootCmd.AddCommand(usersCmd)

	// Get users
	getUsersCmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of users",
		RunE:  getUsers,
	}
	getUsersCmd.Flags().String("name", "", "Filter users by name")
	getUsersCmd.Flags().Bool("enabled", true, "Filter by enabled users")
	getUsersCmd.Flags().String("email", "", "Filter users by email address")
	getUsersCmd.Flags().Bool("is-contact", false, "Show only contact users for the tenant")
	usersCmd.AddCommand(getUsersCmd)

	// Create user
	createUserCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new User",
		RunE:  createUser,
	}
	createUserCmd.Flags().String("firstname", "", "User's firstname")
	createUserCmd.Flags().String("lastname", "", "User's lastname")
	createUserCmd.Flags().String("email", "", "User's email")
	createUserCmd.Flags().Bool("enabled", true, "Is User enabled?")
	createUserCmd.Flags().String("mobile-phone", "", "User's mobile phone number")
	createUserCmd.Flags().Bool("is-contact", false, "Is the user a contact person for its own tenant?")
	createUserCmd.MarkFlagRequired("firstname")
	createUserCmd.MarkFlagRequired("lastname")
	createUserCmd.MarkFlagRequired("email")
	usersCmd.AddCommand(createUserCmd)

	// Get user details
	getUserDetailsCmd := &cobra.Command{
		Use:   "details [id]",
		Short: "Get User details",
		Args:  cobra.ExactArgs(1),
		RunE:  getUserDetails,
	}
	usersCmd.AddCommand(getUserDetailsCmd)

	// Update user
	updateUserCmd := &cobra.Command{
		Use:   "update [id]",
		Short: "Update a User",
		Args:  cobra.ExactArgs(1),
		RunE:  updateUser,
	}
	updateUserCmd.Flags().String("firstname", "", "User's firstname")
	updateUserCmd.Flags().String("lastname", "", "User's lastname")
	updateUserCmd.Flags().String("email", "", "User's email")
	updateUserCmd.Flags().Bool("enabled", true, "Is User enabled?")
	updateUserCmd.Flags().String("mobile-phone", "", "User's mobile phone number")
	updateUserCmd.Flags().Bool("is-contact", false, "Is the user a contact person for its own tenant?")
	usersCmd.AddCommand(updateUserCmd)

	// Get logged in user details
	getWhoAmICmd := &cobra.Command{
		Use:   "whoami",
		Short: "Get logged in user details",
		RunE:  getWhoAmI,
	}
	usersCmd.AddCommand(getWhoAmICmd)

	// Get not assigned user details
	getNotAssignedUserCmd := &cobra.Command{
		Use:   "not-assigned",
		Short: "Get details from the not assigned user",
		RunE:  getNotAssignedUser,
	}
	usersCmd.AddCommand(getNotAssignedUserCmd)

	// Get on delegation user details
	getOnDelegationUserCmd := &cobra.Command{
		Use:   "on-delegation",
		Short: "Get details from the on delegation user",
		RunE:  getOnDelegationUser,
	}
	usersCmd.AddCommand(getOnDelegationUserCmd)
}

func getUsers(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	enabled, _ := cmd.Flags().GetBool("enabled")
	email, _ := cmd.Flags().GetString("email")
	isContact, _ := cmd.Flags().GetBool("is-contact")

	params := make(map[string]string)
	if name != "" {
		params["name"] = name
	}
	params["enabled"] = strconv.FormatBool(enabled)
	if email != "" {
		params["email"] = email
	}
	params["isContact"] = strconv.FormatBool(isContact)

	response, err := client.GetUsers(cloudTempleID, params)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func createUser(cmd *cobra.Command, args []string) error {
	firstname, _ := cmd.Flags().GetString("firstname")
	lastname, _ := cmd.Flags().GetString("lastname")
	email, _ := cmd.Flags().GetString("email")
	enabled, _ := cmd.Flags().GetBool("enabled")
	mobilePhone, _ := cmd.Flags().GetString("mobile-phone")
	isContact, _ := cmd.Flags().GetBool("is-contact")

	userData := map[string]interface{}{
		"firstname": firstname,
		"lastname":  lastname,
		"email":     email,
		"enabled":   enabled,
	}
	if mobilePhone != "" {
		userData["mobilePhoneNumber"] = mobilePhone
	}
	userData["isContact"] = isContact

	response, err := client.CreateUser(cloudTempleID, userData)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getUserDetails(cmd *cobra.Command, args []string) error {
	response, err := client.GetUserDetails(args[0])
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func updateUser(cmd *cobra.Command, args []string) error {
	firstname, _ := cmd.Flags().GetString("firstname")
	lastname, _ := cmd.Flags().GetString("lastname")
	email, _ := cmd.Flags().GetString("email")
	enabled, _ := cmd.Flags().GetBool("enabled")
	mobilePhone, _ := cmd.Flags().GetString("mobile-phone")
	isContact, _ := cmd.Flags().GetBool("is-contact")

	userData := make(map[string]interface{})
	if firstname != "" {
		userData["firstname"] = firstname
	}
	if lastname != "" {
		userData["lastname"] = lastname
	}
	if email != "" {
		userData["email"] = email
	}
	if cmd.Flags().Changed("enabled") {
		userData["enabled"] = enabled
	}
	if mobilePhone != "" {
		userData["mobilePhoneNumber"] = mobilePhone
	}
	if cmd.Flags().Changed("is-contact") {
		userData["isContact"] = isContact
	}

	response, err := client.UpdateUser(args[0], userData)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getWhoAmI(cmd *cobra.Command, args []string) error {
	response, err := client.GetWhoAmI()
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getNotAssignedUser(cmd *cobra.Command, args []string) error {
	response, err := client.GetNotAssignedUser()
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getOnDelegationUser(cmd *cobra.Command, args []string) error {
	response, err := client.GetOnDelegationUser()
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}
