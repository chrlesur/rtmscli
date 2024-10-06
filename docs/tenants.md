
# Tenants Module Documentation

The Tenants module provides commands to manage tenants in the RTMS system. The Tenants module allows users to manage various aspects of tenants in the RTMS system, including creating and listing tenants, managing SSH keys, and configuring workflow emails. These functionalities are crucial for maintaining and configuring multi-tenant environments within the RTMS system.

This documentation covers all the verbs (subcommands) available in the Tenants module as defined in the provided code. Each command's purpose, required arguments, and available options are detailed to provide a comprehensive guide for users of the rtmscli tool.


## Base Command

```
rtmscli tenants
```

This is the base command for all tenant-related operations. It doesn't perform any action on its own but serves as a parent for the subcommands.

## Subcommands

### 1. List Tenants

```
rtmscli tenants list
```

Retrieves a list of Tenants.

Options:
- `--name`: (Optional) Filter tenants by name.
- `--responsible-team-id`: (Optional) Filter tenants by responsible team ID.
- `--sdm-id`: (Optional) Filter tenants by SDM ID.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 2. Create Tenant

```
rtmscli tenants create
```

Creates a new Tenant.

Options:
- `--name`: (Required) Tenant's name.
- `--phone`: (Required) Tenant's phone.
- `--address`: (Required) Tenant's address.
- `--postal-code`: (Required) Tenant's postal code.
- `--city`: (Required) Tenant's city.
- `--country`: (Required) Tenant's country.
- `--responsible-team`: (Required) Tenant's responsible team ID.
- `--contact`: (Required) Tenant's contact ID.
- `--watchers`: (Optional) List of default watcher email addresses.
- `--is-enabled`: (Optional) Is Tenant active? Default is true.
- `--cloud-temple-id`: (Optional) MySI Tenant's identifier.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 3. Get Tenant Details

```
rtmscli tenants details [id]
```

Retrieves detailed information about a specific tenant.

Arguments:
- `id`: (Required) The ID of the tenant.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 4. Get Tenant Contacts

```
rtmscli tenants contacts [id]
```

Retrieves contacts for a specific tenant.

Arguments:
- `id`: (Required) The ID of the tenant.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 5. Request Tenant Deletion

```
rtmscli tenants request-deletion [id]
```

Requests deletion of a specific tenant.

Arguments:
- `id`: (Required) The ID of the tenant.

Options:
- `--delete`: (Required) Enable or disable deletion request.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 6. List Tenant SSH Keys

```
rtmscli tenants ssh-keys list [tenant-id]
```

Lists all SSH keys of a tenant.

Arguments:
- `tenant-id`: (Required) The ID of the tenant.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 7. Generate Tenant SSH Key

```
rtmscli tenants ssh-keys generate [tenant-id]
```

Generates a new SSH key for a tenant.

Arguments:
- `tenant-id`: (Required) The ID of the tenant.

Options:
- `--comment`: (Optional) Free comment.
- `--is-active`: (Optional) Key state. Default is false.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 8. Delete Tenant SSH Key

```
rtmscli tenants ssh-keys delete [key-id]
```

Deletes an SSH key of a tenant.

Arguments:
- `key-id`: (Required) The ID of the SSH key to delete.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 9. Update Tenant SSH Key

```
rtmscli tenants ssh-keys update [key-id]
```

Updates an SSH key.

Arguments:
- `key-id`: (Required) The ID of the SSH key to update.

Options:
- `--is-active`: (Optional) Key state. Default is false.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 10. Get Tenant Workflow Emails

```
rtmscli tenants workflow-emails get [tenant-id]
```

Retrieves workflow emails details for a tenant.

Arguments:
- `tenant-id`: (Required) The ID of the tenant.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 11. Edit Tenant Workflow Emails Generalities

```
rtmscli tenants workflow-emails edit-generalities [tenant-id]
```

Edits workflow emails generalities for a tenant.

Arguments:
- `tenant-id`: (Required) The ID of the tenant.

Options:
- `--format`: (Optional) Format of emails sent (HTML or TEXT).
- `--from`: (Optional) Email address used to send emails.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 12. Edit Tenant Workflow Emails Create Ticket

```
rtmscli tenants workflow-emails edit-create-ticket [tenant-id]
```

Edits workflow emails for ticket creation for a tenant.

Arguments:
- `tenant-id`: (Required) The ID of the tenant.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 13. Edit Tenant Workflow Emails Update Ticket

```
rtmscli tenants workflow-emails edit-update-ticket [tenant-id]
```

Edits workflow emails for ticket updates for a tenant.

Arguments:
- `tenant-id`: (Required) The ID of the tenant.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 14. Edit Tenant Workflow Emails Validation Client Ticket

```
rtmscli tenants workflow-emails edit-validation-client-ticket [tenant-id]
```

Edits workflow emails for client ticket validation for a tenant.

Arguments:
- `tenant-id`: (Required) The ID of the tenant.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 15. Edit Tenant Workflow Emails Close Ticket

```
rtmscli tenants workflow-emails edit-close-ticket [tenant-id]
```

Edits workflow emails for ticket closure for a tenant.

Arguments:
- `tenant-id`: (Required) The ID of the tenant.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

Note: All commands support the global flags defined in the root command, such as `--debug` for enabling debug mode.