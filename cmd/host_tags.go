package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hostTagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Manage host tags",
	Long:  `Manage host tags, including listing, creating, updating, and deleting tags, as well as listing hosts for a specific tag.`,
}

func init() {
	hostsCmd.AddCommand(hostTagsCmd)

	// Get host tags
	getHostTagsCmd := &cobra.Command{
		Use:   "list",
		Short: "Get a list of host tags",
		RunE:  getHostTags,
	}
	getHostTagsCmd.Flags().String("label", "", "Filter by label")
	hostTagsCmd.AddCommand(getHostTagsCmd)

	// Create host tag
	createHostTagCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a host tag",
		RunE:  createHostTag,
	}
	createHostTagCmd.Flags().String("label", "", "Tag label")
	createHostTagCmd.Flags().String("description", "", "Tag description")
	createHostTagCmd.Flags().IntSlice("hosts", nil, "List of host IDs to associate with the tag")
	createHostTagCmd.MarkFlagRequired("label")
	hostTagsCmd.AddCommand(createHostTagCmd)

	// Get host tag details
	getHostTagDetailsCmd := &cobra.Command{
		Use:   "details [id]",
		Short: "Get details of a tag",
		Args:  cobra.ExactArgs(1),
		RunE:  getHostTagDetails,
	}
	hostTagsCmd.AddCommand(getHostTagDetailsCmd)

	// Remove host tag
	removeHostTagCmd := &cobra.Command{
		Use:   "remove [id]",
		Short: "Remove a host tag",
		Args:  cobra.ExactArgs(1),
		RunE:  removeHostTag,
	}
	hostTagsCmd.AddCommand(removeHostTagCmd)

	// Edit host tag
	editHostTagCmd := &cobra.Command{
		Use:   "edit [id]",
		Short: "Edit host tag",
		Args:  cobra.ExactArgs(1),
		RunE:  editHostTag,
	}
	editHostTagCmd.Flags().String("label", "", "Tag label")
	editHostTagCmd.Flags().String("description", "", "Tag description")
	editHostTagCmd.Flags().IntSlice("hosts", nil, "List of host IDs to associate with the tag")
	hostTagsCmd.AddCommand(editHostTagCmd)

	// Get hosts by tag
	getHostsByTagCmd := &cobra.Command{
		Use:   "hosts [id]",
		Short: "Gets hosts that match a given tag",
		Args:  cobra.ExactArgs(1),
		RunE:  getHostsByTag,
	}
	hostTagsCmd.AddCommand(getHostsByTagCmd)
}

func getHostTags(cmd *cobra.Command, args []string) error {
	label, _ := cmd.Flags().GetString("label")
	params := make(map[string]string)
	if label != "" {
		params["label"] = label
	}

	response, err := client.GetHostTags(cloudTempleID, params)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func createHostTag(cmd *cobra.Command, args []string) error {
	label, _ := cmd.Flags().GetString("label")
	description, _ := cmd.Flags().GetString("description")
	hosts, _ := cmd.Flags().GetIntSlice("hosts")

	tagData := map[string]interface{}{
		"label": label,
	}
	if description != "" {
		tagData["description"] = description
	}
	if len(hosts) > 0 {
		tagData["hosts"] = hosts
	}

	response, err := client.CreateHostTag(cloudTempleID, tagData)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getHostTagDetails(cmd *cobra.Command, args []string) error {
	response, err := client.GetHostTagDetails(args[0])
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func removeHostTag(cmd *cobra.Command, args []string) error {
	response, err := client.RemoveHostTag(args[0])
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func editHostTag(cmd *cobra.Command, args []string) error {
	label, _ := cmd.Flags().GetString("label")
	description, _ := cmd.Flags().GetString("description")
	hosts, _ := cmd.Flags().GetIntSlice("hosts")

	tagData := make(map[string]interface{})
	if label != "" {
		tagData["label"] = label
	}
	if description != "" {
		tagData["description"] = description
	}
	if len(hosts) > 0 {
		tagData["hosts"] = hosts
	}

	response, err := client.EditHostTag(args[0], tagData)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

func getHostsByTag(cmd *cobra.Command, args []string) error {
	response, err := client.GetHostsByTag(args[0], nil)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}
