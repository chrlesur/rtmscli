# Users Module Documentation

The Users module provides commands to manage users in the RTMS system. The Users module allows administrators to manage user accounts within the RTMS system, including creating new users, updating existing user information, and retrieving user details. It also provides functionality to handle special user types like "not assigned" and "on delegation" users.

This documentation covers all the verbs (subcommands) available in the Users module as defined in the provided code. Each command's purpose, required arguments, and available options are detailed to provide a comprehensive guide for users of the rtmscli tool.


## Base Command

```
rtmscli users
```

This is the base command for all user-related operations. It doesn't perform any action on its own but serves as a parent for the subcommands.

## Subcommands

### 1. List Users

```
rtmscli users list
```

Retrieves a list of users.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID to filter the users.
- `--name`: (Optional) Filter users by name.
- `--enabled`: (Optional) Filter by enabled users. Default is true.
- `--email`: (Optional) Filter users by email address.
- `--is-contact`: (Optional) Show only contact users for the tenant. Default is false.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 2. Create User

```
rtmscli users create
```

Creates a new User.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID for the new user.
- `--firstname`: (Required) User's firstname.
- `--lastname`: (Required) User's lastname.
- `--email`: (Required) User's email.
- `--enabled`: (Optional) Is User enabled? Default is true.
- `--mobile-phone`: (Optional) User's mobile phone number.
- `--is-contact`: (Optional) Is the user a contact person for its own tenant? Default is false.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 3. Get User Details

```
rtmscli users details [id]
```

Retrieves detailed information about a specific user.

Arguments:
- `id`: (Required) The ID of the user.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 4. Update User

```
rtmscli users update [id]
```

Updates information for a specific user.

Arguments:
- `id`: (Required) The ID of the user to update.

Options:
- `--firstname`: (Optional) User's new firstname.
- `--lastname`: (Optional) User's new lastname.
- `--email`: (Optional) User's new email.
- `--enabled`: (Optional) Is User enabled?
- `--mobile-phone`: (Optional) User's new mobile phone number.
- `--is-contact`: (Optional) Is the user a contact person for its own tenant?
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 5. Get Logged In User Details

```
rtmscli users whoami
```

Retrieves details of the currently logged in user.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 6. Get Not Assigned User Details

```
rtmscli users not-assigned
```

Retrieves details from the not assigned user.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 7. Get On Delegation User Details

```
rtmscli users on-delegation
```

Retrieves details from the on delegation user.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

Note: All commands support the global flags defined in the root command, such as `--debug` for enabling debug mode.

## Detailed Command Descriptions

### 1. List Users

This command retrieves a list of users. You can filter the users by various criteria such as name, email, enabled status, and whether they are contact users.

Usage example:
```
rtmscli users list --cloud-temple-id "your-id" --name "John" --email "john@example.com" --is-contact
```

### 2. Create User

This command creates a new user in the system. You need to provide essential information such as firstname, lastname, and email.

Usage example:
```
rtmscli users create --cloud-temple-id "your-id" --firstname "John" --lastname "Doe" --email "john.doe@example.com" --mobile-phone "+1234567890" --is-contact
```

### 3. Get User Details

This command retrieves detailed information about a specific user based on their ID.

Usage example:
```
rtmscli users details user-123
```

### 4. Update User

This command allows you to update information for an existing user. You can modify various fields such as firstname, lastname, email, enabled status, mobile phone number, and contact status.

Usage example:
```
rtmscli users update user-123 --firstname "Johnny" --email "johnny.doe@example.com" --enabled false
```

### 5. Get Logged In User Details

This command retrieves details of the currently logged in user. It's useful for checking your own user information or permissions.

Usage example:
```
rtmscli users whoami
```

### 6. Get Not Assigned User Details

This command retrieves details of the "not assigned" user. This is typically a special user account used when no specific user is assigned to a task or item.

Usage example:
```
rtmscli users not-assigned
```

### 7. Get On Delegation User Details

This command retrieves details of the "on delegation" user. This is typically used when tasks or responsibilities are delegated to a temporary or substitute user.

Usage example:
```
rtmscli users on-delegation
```