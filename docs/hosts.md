# Hosts Module Documentation

This documentation covers all the verbs (subcommands) available in the hosts module, including host tags management, as defined in the provided code. Each command's purpose, required arguments, and available options are detailed to provide a comprehensive guide for users of the rtmscli tool.

## Base Command

```
rtmscli hosts
```

This is the base command for all host-related operations. It doesn't perform any action on its own but serves as a parent for the subcommands.

## Subcommands

### 1. List Hosts

```
rtmscli hosts list
```

Retrieves a list of Hosts.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID to filter the hosts.
- `--name`: (Optional) Filter hosts by name.
- `--status`: (Optional) Filter by hosts status (UP, DOWN, PENDING, UNREACHABLE). Can be specified multiple times.
- `--is-monitored`: (Optional) Filter by monitored hosts.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 2. Create Host

```
rtmscli hosts create
```

Creates a new Host.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID for the new host.
- `--name`: (Required) Host name.
- `--address`: (Required) Host monitoring IP address.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 3. Get Host Details

```
rtmscli hosts details [id]
```

Retrieves detailed information about a specific host.

Arguments:
- `id`: (Required) The ID of the host.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 4. Remove Host

```
rtmscli hosts remove [id]
```

Removes a specific host.

Arguments:
- `id`: (Required) The ID of the host to remove.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 5. Update Host

```
rtmscli hosts update [id]
```

Updates information for a specific host.

Arguments:
- `id`: (Required) The ID of the host to update.

Options:
- `--name`: (Optional) New host name.
- `--address`: (Optional) New host monitoring IP address.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 6. Get Host Services

```
rtmscli hosts services [id]
```

Retrieves the services associated with a specific host.

Arguments:
- `id`: (Required) The ID of the host.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 7. Update Host Tags

```
rtmscli hosts update-tags [id]
```

Updates the tags associated with a specific host.

Arguments:
- `id`: (Required) The ID of the host.

Options:
- `--tags`: (Required) List of tag IDs to associate with the host.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 8. Switch Host Monitoring

```
rtmscli hosts switch-monitoring [id]
```

Enables or disables monitoring for all or specific host's services.

Arguments:
- `id`: (Required) The ID of the host.

Options:
- `--enable`: (Required) Enable or disable monitoring.
- `--services`: (Optional) List of service IDs. If not provided, applies to all services.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 9. Switch Host Monitoring Notifications

```
rtmscli hosts switch-notifications [id]
```

Enables or disables monitoring notifications for all or specific host's services.

Arguments:
- `id`: (Required) The ID of the host.

Options:
- `--enable`: (Required) Enable or disable notifications.
- `--services`: (Optional) List of service IDs. If not provided, applies to all services.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 10. Get Hosts Stats

```
rtmscli hosts stats
```

Retrieves hosts status statistics.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID to filter the stats.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

## Host Tags Subcommands

### 11. List Host Tags

```
rtmscli hosts tags list
```

Retrieves a list of host tags.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID to filter the tags.
- `--label`: (Optional) Filter by label.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 12. Create Host Tag

```
rtmscli hosts tags create
```

Creates a new host tag.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID for the new tag.
- `--label`: (Required) Tag label.
- `--description`: (Optional) Tag description.
- `--hosts`: (Optional) List of host IDs to associate with the tag.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 13. Get Host Tag Details

```
rtmscli hosts tags details [id]
```

Retrieves details of a specific tag.

Arguments:
- `id`: (Required) The ID of the tag.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 14. Remove Host Tag

```
rtmscli hosts tags remove [id]
```

Removes a specific host tag.

Arguments:
- `id`: (Required) The ID of the tag to remove.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 15. Edit Host Tag

```
rtmscli hosts tags edit [id]
```

Edits an existing host tag.

Arguments:
- `id`: (Required) The ID of the tag to edit.

Options:
- `--label`: (Optional) New tag label.
- `--description`: (Optional) New tag description.
- `--hosts`: (Optional) New list of host IDs to associate with the tag.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 16. Get Hosts by Tag

```
rtmscli hosts tags hosts [id]
```

Retrieves hosts that match a given tag.

Arguments:
- `id`: (Required) The ID of the tag.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

Note: All commands support the global flags defined in the root command, such as `--debug` for enabling debug mode.