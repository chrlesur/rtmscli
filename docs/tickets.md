# Ticket Management

The RTMS CLI provides comprehensive commands for managing tickets, including creating, updating, and viewing tickets, as well as managing ticket comments, attachments, and tags.

## Available Commands

### Main Ticket Commands
- `rtmscli tickets list`: List all tickets
- `rtmscli tickets create`: Create a new ticket
- `rtmscli tickets count`: Get the number of tickets
- `rtmscli tickets details`: Get details of a specific ticket
- `rtmscli tickets edit`: Edit ticket information
- `rtmscli tickets catalogs`: Get ticket catalogs
- `rtmscli tickets stats`: Get ticket status statistics

### Ticket Comments
- `rtmscli tickets comments list-all`: List all ticket comments
- `rtmscli tickets comments list`: List comments for a specific ticket
- `rtmscli tickets comments post`: Post a new comment on a ticket
- `rtmscli tickets comments edit`: Edit an existing comment

### Ticket Attachments
- `rtmscli tickets attachments list`: List attachments for a ticket
- `rtmscli tickets attachments upload`: Upload an attachment to a ticket
- `rtmscli tickets attachments download`: Download an attachment
- `rtmscli tickets attachments remove`: Remove an attachment from a ticket

### Ticket Tags
- `rtmscli tickets tags list`: List all ticket tags
- `rtmscli tickets tags create`: Create a new ticket tag
- `rtmscli tickets tags details`: Get details of a specific tag
- `rtmscli tickets tags remove`: Remove a ticket tag
- `rtmscli tickets tags edit`: Edit a ticket tag
- `rtmscli tickets tags tickets`: Get tickets associated with a tag

## Usage Examples

### List Tickets

To list all tickets:

```
rtmscli tickets list
```

Options:
- `--name`: Filter tickets by subject
- `--status`: Filter tickets by status (0-6)
- `--owner`: Filter tickets by owner name
- `--cloud-temple-id`: Specify the Cloud Temple ID (required)

Example:
```
rtmscli tickets list --cloud-temple-id=your_id --status=1,2 --owner="John Doe"
```

### Create a Ticket

To create a new ticket:

```
rtmscli tickets create --name="Ticket Subject" --description="Detailed description" --cloud-temple-id=your_id
```

Options:
- `--owner`: Specify the owner ID
- `--catalog-items`: Specify catalog item IDs

### Get Ticket Details

To get details of a specific ticket:

```
rtmscli tickets details [ticket-id]
```

### Edit a Ticket

To edit an existing ticket:

```
rtmscli tickets edit [ticket-id] --name="Updated Subject" --description="Updated description"
```

### Manage Ticket Comments

To list comments for a specific ticket:

```
rtmscli tickets comments list [ticket-id]
```

To post a new comment:

```
rtmscli tickets comments post [ticket-id] --content="Comment content" --private=false
```

### Manage Ticket Attachments

To upload an attachment:

```
rtmscli tickets attachments upload [ticket-id] [file-path]
```

To download an attachment:

```
rtmscli tickets attachments download [attachment-id] [output-path]
```

### Manage Ticket Tags

To list all ticket tags:

```
rtmscli tickets tags list
```

To create a new tag:

```
rtmscli tickets tags create --label="Urgent" --description="For urgent tickets"
```

## Common Options

All ticket commands support the following options:

- `-f, --format`: Specify output format (json, html, markdown)
- `-H, --host`: Specify the RTMS API host (default is "rtms-api.cloud-temple.com")

Example using format option:
```
rtmscli -f markdown tickets list --cloud-temple-id=your_id
```

For more detailed information on each command and its options, use the `--help` flag:

```
rtmscli tickets --help
rtmscli tickets [command] --help
```

