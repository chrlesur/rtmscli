# User Management

The RTMS CLI provides several commands for managing users. This document outlines the available commands and their usage.

## Available Commands

- `rtmscli users list`: Get a list of users
- `rtmscli users create`: Create a new user
- `rtmscli users details`: Get user details
- `rtmscli users update`: Update a user
- `rtmscli users whoami`: Get logged in user details
- `rtmscli users not-assigned`: Get details of the not assigned user
- `rtmscli users on-delegation`: Get details of the on delegation user

## Usage Examples

### List Users

To list all users:

```
rtmscli users list
```

Options:
- `--name`: Filter users by name
- `--enabled`: Filter by enabled users (true/false)
- `--email`: Filter users by email address
- `--is-contact`: Show only contact users for the tenant (true/false)
- `--cloud-temple-id`: Specify the Cloud Temple ID (required)

Example:
```
rtmscli users list --cloud-temple-id=your_id --name=John --enabled=true
```

### Create User

To create a new user:

```
rtmscli users create --firstname=John --lastname=Doe --email=john.doe@example.com
```

Options:
- `--firstname`: User's first name (required)
- `--lastname`: User's last name (required)
- `--email`: User's email (required)
- `--enabled`: Is user enabled? (true/false, default is true)
- `--mobile-phone`: User's mobile phone number
- `--is-contact`: Is the user a contact person for its own tenant? (true/false)
- `--cloud-temple-id`: Specify the Cloud Temple ID (required)

Example:
```
rtmscli users create --cloud-temple-id=your_id --firstname=John --lastname=Doe --email=john.doe@example.com --mobile-phone="+1234567890" --is-contact=true
```

### Get User Details

To get details of a specific user:

```
rtmscli users details [user-id]
```

Example:
```
rtmscli users details 12345
```

### Update User

To update an existing user:

```
rtmscli users update [user-id] [flags]
```

Options:
- `--firstname`: User's new first name
- `--lastname`: User's new last name
- `--email`: User's new email
- `--enabled`: Update user enabled status (true/false)
- `--mobile-phone`: User's new mobile phone number
- `--is-contact`: Update if the user is a contact person for its own tenant (true/false)

Example:
```
rtmscli users update 12345 --firstname=John --email=new.email@example.com --enabled=false
```

### Get Logged In User Details

To get details of the currently logged in user:

```
rtmscli users whoami
```

### Get Not Assigned User Details

To get details of the not assigned user:

```
rtmscli users not-assigned
```

### Get On Delegation User Details

To get details of the on delegation user:

```
rtmscli users on-delegation
```

## Common Options

All user commands support the following options:

- `-f, --format`: Specify output format (json, html, markdown)
- `-H, --host`: Specify the RTMS API host (default is "rtms-api.cloud-temple.com")

Example using format option:
```
rtmscli -f markdown users list --cloud-temple-id=your_id
```

For more detailed information on each command and its options, use the `--help` flag:

```
rtmscli users --help
rtmscli users [command] --help
```

