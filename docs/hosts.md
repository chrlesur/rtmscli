# Host Management

RTMS CLI provides a comprehensive set of commands for managing hosts. This document outlines the available commands and their usage.

## Available Commands

- `rtmscli hosts list`: List all hosts
- `rtmscli hosts create`: Create a new host
- `rtmscli hosts details`: Get details of a specific host
- `rtmscli hosts remove`: Remove a host
- `rtmscli hosts update`: Update a host
- `rtmscli hosts services`: Get services of a specific host
- `rtmscli hosts tags`: Manage host tags
- `rtmscli hosts monitoring`: Manage host monitoring
- `rtmscli hosts stats`: Get host status statistics

## Usage Examples

### List Hosts

To list all hosts:

```
rtmscli hosts list
```

Options:
- `--name`: Filter hosts by name
- `--status`: Filter hosts by status (UP, DOWN, PENDING, UNREACHABLE)
- `--is-monitored`: Filter by monitored hosts
- `--cloud-temple-id`: Specify the Cloud Temple ID (required)

Example:
```
rtmscli hosts list --cloud-temple-id=your_id --name=server1 --status=UP
```

### Create Host

To create a new host:

```
rtmscli hosts create --name=<name> --address=<address>
```

Example:
```
rtmscli hosts create --name=newserver --address=192.168.1.100
```

### Get Host Details

To get details of a specific host:

```
rtmscli hosts details [host-id]
```

Example:
```
rtmscli hosts details 12345
```

### Remove Host

To remove a host:

```
rtmscli hosts remove [host-id]
```

Example:
```
rtmscli hosts remove 12345
```

### Update Host

To update a host:

```
rtmscli hosts update [host-id] [flags]
```

Example:
```
rtmscli hosts update 12345 --name=updatedserver --address=192.168.1.101
```

### Get Host Services

To get services of a specific host:

```
rtmscli hosts services [host-id]
```

Example:
```
rtmscli hosts services 12345
```

## Managing Host Tags

### Update Host Tags

To update tags for a host:

```
rtmscli hosts tags update [host-id] --tags=<tag1>,<tag2>
```

Example:
```
rtmscli hosts tags update 12345 --tags=production,webserver
```

## Host Monitoring

### Enable/Disable Host Monitoring

To enable or disable monitoring for a host:

```
rtmscli hosts monitoring [host-id] --enable=<true|false>
```

Example:
```
rtmscli hosts monitoring 12345 --enable=true
```

### Enable/Disable Host Monitoring Notifications

To enable or disable monitoring notifications for a host:

```
rtmscli hosts monitoring notifications [host-id] --enable=<true|false>
```

Example:
```
rtmscli hosts monitoring notifications 12345 --enable=true
```

## Host Statistics

To get host status statistics:

```
rtmscli hosts stats
```

## Common Options

All host commands support the following options:

- `-f, --format`: Specify output format (json, html, markdown)
- `-H, --host`: Specify the RTMS API host (default is "rtms-api.cloud-temple.com")

Example using format option:
```
rtmscli -f markdown hosts list --cloud-temple-id=your_id
```

For more detailed information on each command and its options, use the `--help` flag:

```
rtmscli hosts --help
rtmscli hosts [command] --help
```

## Best Practices

1. Always use the `--cloud-temple-id` flag when required to ensure you're working with the correct tenant.
2. Use tags to organize and categorize your hosts for easier management.
3. Regularly review and update host monitoring settings to ensure optimal system oversight.
4. Use the `hosts stats` command to get an overview of your host statuses and identify potential issues.

