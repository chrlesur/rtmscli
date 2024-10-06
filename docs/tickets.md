
# Tickets Module Documentation

The Tickets module provides commands to manage tickets in the RTMS system. The Tickets module allows users to manage various aspects of tickets in the RTMS system, including creating and listing tickets, managing attachments, comments, and tags. These functionalities are crucial for maintaining and tracking issues within the RTMS system.

This documentation covers all the verbs (subcommands) available in the Tickets module as defined in the provided code. Each command's purpose, required arguments, and available options are detailed to provide a comprehensive guide for users of the rtmscli tool.

## Base Command

```
rtmscli tickets
```

This is the base command for all ticket-related operations. It doesn't perform any action on its own but serves as a parent for the subcommands.

## Subcommands

### 1. List Tickets

```
rtmscli tickets list
```

Retrieves a list of Tickets.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID to filter the tickets.
- `--name`: (Optional) Filter tickets by subject (name).
- `--status`: (Optional) Filter Tickets by one or more status (0-6). Can be specified multiple times.
- `--owner`: (Optional) Filter tickets by owner name.
- `--owner-ids`: (Optional) Filter tickets by one or more owner RTMS identifiers. Can be specified multiple times.
- `--is-not-assigned`: (Optional) Filter non assigned tickets.
- `--is-on-delegation`: (Optional) Filter tickets on delegation.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 2. Create Ticket

```
rtmscli tickets create
```

Creates a new Ticket.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID for the new ticket.
- `--name`: (Required) A title, the issue in short.
- `--description`: (Required) Detailed description of the issue.
- `--owner`: (Optional) Identifier of the user in charge of solving the issue.
- `--catalog-items`: (Optional) Collection of classification catalog item identifiers. Can be specified multiple times.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 3. Get Tickets Count

```
rtmscli tickets count
```

Retrieves the number of Tickets.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID to filter the tickets.
- `--status`: (Optional) Filter Tickets by status (0-6).
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 4. Get Ticket Details

```
rtmscli tickets details [id]
```

Retrieves detailed information about a specific ticket.

Arguments:
- `id`: (Required) The ID of the ticket.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 5. Edit Ticket

```
rtmscli tickets edit [id]
```

Edits information for a specific ticket.

Arguments:
- `id`: (Required) The ID of the ticket to edit.

Options:
- `--name`: (Optional) A new title, the issue in short.
- `--description`: (Optional) A new detailed description of the issue.
- `--owner`: (Optional) New identifier of the user in charge of solving the issue.
- `--catalog-items`: (Optional) New collection of classification catalog item identifiers. Can be specified multiple times.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 6. Get Ticket Catalogs

```
rtmscli tickets catalogs [id]
```

Retrieves catalogs for a specific ticket.

Arguments:
- `id`: (Required) The ID of the ticket.

Options:
- `--selected-item`: (Optional) Show classification catalog with selected items for this ticket.
- `--available-items`: (Optional) Show classification catalog with all available items for this ticket.
- `--is-root`: (Optional) If true, only classification root catalogs will be displayed.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 7. Get Tickets Stats

```
rtmscli tickets stats
```

Retrieves tickets status statistics.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID to filter the stats.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 8. List Ticket Attachments

```
rtmscli tickets attachments list [ticket-id]
```

Lists attachments for a specific ticket.

Arguments:
- `ticket-id`: (Required) The ID of the ticket.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 9. Upload Ticket Attachment

```
rtmscli tickets attachments upload [ticket-id] [file-path]
```

Uploads an attachment to a specific ticket.

Arguments:
- `ticket-id`: (Required) The ID of the ticket.
- `file-path`: (Required) The path to the file to be uploaded.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 10. Download Ticket Attachment

```
rtmscli tickets attachments download [attachment-id] [output-path]
```

Downloads an attachment from a ticket.

Arguments:
- `attachment-id`: (Required) The ID of the attachment.
- `output-path`: (Required) The path where the downloaded file will be saved.

### 11. Remove Ticket Attachment

```
rtmscli tickets attachments remove [attachment-id]
```

Removes an attachment from a ticket.

Arguments:
- `attachment-id`: (Required) The ID of the attachment to remove.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 12. List All Ticket Comments

```
rtmscli tickets comments list-all
```

Retrieves all Ticket comments.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID to filter the comments.
- `--ticket`: (Optional) Filter by ticket ID.
- `--user`: (Optional) Filter by user ID.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 13. List Ticket Comments

```
rtmscli tickets comments list [ticket-id]
```

Retrieves comments for a specific ticket.

Arguments:
- `ticket-id`: (Required) The ID of the ticket.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 14. Post Ticket Comment

```
rtmscli tickets comments post [ticket-id]
```

Posts a comment to a specific ticket.

Arguments:
- `ticket-id`: (Required) The ID of the ticket.

Options:
- `--content`: (Required) Comment content.
- `--private`: (Optional) Comment privacy.
- `--duration`: (Optional) Working time on the Ticket.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 15. Edit Ticket Comment

```
rtmscli tickets comments edit [comment-id]
```

Edits a specific ticket comment.

Arguments:
- `comment-id`: (Required) The ID of the comment to edit.

Options:
- `--content`: (Optional) New comment content.
- `--private`: (Optional) New comment privacy.
- `--duration`: (Optional) New working time on the Ticket.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 16. List Ticket Tags

```
rtmscli tickets tags list
```

Retrieves a list of ticket tags.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID to filter the tags.
- `--label`: (Optional) Filter by label.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 17. Create Ticket Tag

```
rtmscli tickets tags create
```

Creates a new ticket tag.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID for the new tag.
- `--label`: (Required) Tag label.
- `--description`: (Optional) Tag description.
- `--tickets`: (Optional) List of ticket IDs to associate with the tag. Can be specified multiple times.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 18. Get Ticket Tag Details

```
rtmscli tickets tags details [id]
```

Retrieves details of a specific ticket tag.

Arguments:
- `id`: (Required) The ID of the tag.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 19. Remove Ticket Tag

```
rtmscli tickets tags remove [id]
```

Removes a specific ticket tag.

Arguments:
- `id`: (Required) The ID of the tag to remove.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 20. Edit Ticket Tag

```
rtmscli tickets tags edit [id]
```

Edits a specific ticket tag.

Arguments:
- `id`: (Required) The ID of the tag to edit.

Options:
- `--label`: (Optional) New tag label.
- `--description`: (Optional) New tag description.
- `--tickets`: (Optional) New list of ticket IDs to associate with the tag. Can be specified multiple times.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 21. Get Tickets by Tag

```
rtmscli tickets tags tickets [id]
```

Retrieves tickets that match a given tag.

Arguments:
- `id`: (Required) The ID of the tag.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

Note: All commands support the global flags defined in the root command, such as `--debug` for enabling debug mode.
