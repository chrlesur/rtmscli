# Appliance Management

The RTMS CLI provides several commands for managing appliances. This document outlines the available commands and their usage.

## Available Commands

- `rtmscli appliances list`: List all appliances
- `rtmscli appliances details`: Get details of a specific appliance
- `rtmscli appliances services`: Get services of a specific appliance
- `rtmscli appliances synchronize`: Synchronize an appliance
- `rtmscli appliances configuration`: Get appliance configuration
- `rtmscli appliances healthcheck`: Get or post appliance health check

## Usage Examples

### List Appliances

To list all appliances:

```
rtmscli appliances list
```

Options:
- `--name`: Filter appliances by name
- `--cloud-temple-id`: Specify the Cloud Temple ID (required)

Example:
```
rtmscli appliances list --cloud-temple-id=your_id --name=app1
```

### Get Appliance Details

To get details of a specific appliance:

```
rtmscli appliances details [appliance-id]
```

Example:
```
rtmscli appliances details 12345
```

### Get Appliance Services

To get services of a specific appliance:

```
rtmscli appliances services [appliance-id]
```

Example:
```
rtmscli appliances services 12345
```

### Synchronize Appliance

To synchronize an appliance:

```
rtmscli appliances synchronize [appliance-id]
```

Example:
```
rtmscli appliances synchronize 12345
```

### Get Appliance Configuration

To get the configuration of an appliance:

```
rtmscli appliances configuration [appliance-id] --appliance-version=[version] --plugins-path=[path]
```

Example:
```
rtmscli appliances configuration 12345 --appliance-version=1.0.0 --plugins-path=/path/to/plugins
```

### Appliance Health Check

To get the last health check of an appliance:

```
rtmscli appliances healthcheck get [appliance-id]
```

To post a health check for an appliance:

```
rtmscli appliances healthcheck post [appliance-id] --appliance-version=[version] --nagios-operating-state=[state] --details=[details]
```

Example:
```
rtmscli appliances healthcheck post 12345 --appliance-version=1.0.0 --nagios-operating-state=OK --details="Everything is running smoothly"
```

## Common Options

All appliance commands support the following options:

- `-f, --format`: Specify output format (json, html, markdown)
- `-H, --host`: Specify the RTMS API host (default is "rtms-api.cloud-temple.com")

Example using format option:
```
rtmscli -f markdown appliances list --cloud-temple-id=your_id
```

For more detailed information on each command and its options, use the `--help` flag:

```
rtmscli appliances --help
rtmscli appliances [command] --help
```

