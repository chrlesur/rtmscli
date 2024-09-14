# Tenant Management

The RTMS CLI provides comprehensive commands for managing tenants. This document outlines the available commands and their usage, including tenant details, SSH key management, and workflow email configuration.

## Available Commands

### General Tenant Commands
- `rtmscli tenants list`: List all tenants
- `rtmscli tenants create`: Create a new tenant
- `rtmscli tenants details`: Get details of a specific tenant
- `rtmscli tenants contacts`: Get tenant contacts
- `rtmscli tenants request-deletion`: Request tenant deletion

### SSH Key Management
- `rtmscli tenants ssh-keys list`: List all SSH keys of a tenant
- `rtmscli tenants ssh-keys generate`: Generate a new SSH key for a tenant
- `rtmscli tenants ssh-keys delete`: Delete an SSH key of a tenant
- `rtmscli tenants ssh-keys update`: Update an SSH key

### Workflow Emails
- `rtmscli tenants workflow-emails get`: Get workflow emails details
- `rtmscli tenants workflow-emails edit-generalities`: Edit workflow emails generalities
- `rtmscli tenants workflow-emails edit-create-ticket`: Edit workflow emails create ticket
- `rtmscli tenants workflow-emails edit-update-ticket`: Edit workflow emails update ticket
- `rtmscli tenants workflow-emails edit-validation-client-ticket`: Edit workflow emails validation client ticket
- `rtmscli tenants workflow-emails edit-close-ticket`: Edit workflow emails close ticket

## Usage Examples

### List Tenants

To list all tenants:

```
rtmscli tenants list
```

Options:
- `--name`: Filter tenants by name
- `--responsible-team-id`: Filter by responsible team ID
- `--sdm-id`: Filter by SDM ID

Example:
```
rtmscli tenants list --name="Example Corp" --responsible-team-id=123
```

### Create Tenant

To create a new tenant:

```
rtmscli tenants create --name="New Tenant" --phone="1234567890" --address="123 Main St" --postal-code="12345" --city="Example City" --country="Example Country" --responsible-team=123 --contact=456
```

### Get Tenant Details

To get details of a specific tenant:

```
rtmscli tenants details [tenant-id]
```

### Request Tenant Deletion

To request deletion of a tenant:

```
rtmscli tenants request-deletion [tenant-id] --delete=true
```

### SSH Key Management

List SSH keys:
```
rtmscli tenants ssh-keys list [tenant-id]
```

Generate a new SSH key:
```
rtmscli tenants ssh-keys generate [tenant-id] --comment="New key" --is-active=true
```

Delete an SSH key:
```
rtmscli tenants ssh-keys delete [key-id]
```

Update an SSH key:
```
rtmscli tenants ssh-keys update [key-id] --is-active=true
```

### Workflow Emails

Get workflow emails details:
```
rtmscli tenants workflow-emails get [tenant-id]
```

Edit workflow emails generalities:
```
rtmscli tenants workflow-emails edit-generalities [tenant-id] --format=HTML --from=noreply@example.com
```

Edit workflow emails for creating a ticket:
```
rtmscli tenants workflow-emails edit-create-ticket [tenant-id] --allow-requester=true --allow-contact=true
```

## Common Options

All tenant commands support the following options:

- `-f, --format`: Specify output format (json, html, markdown)
- `-H, --host`: Specify the RTMS API host (default is "rtms-api.cloud-temple.com")

Example using format option:
```
rtmscli -f markdown tenants list
```

For more detailed information on each command and its options, use the `--help` flag:

```
rtmscli tenants --help
rtmscli tenants [command] --help
```

