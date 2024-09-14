# Appliance Management

The RTMS CLI provides several commands for managing appliances. This document outlines the available commands and their usage.

## Available Commands

- `rtmscli get-appliances`: List all appliances
- `rtmscli get-appliance-details`: Get details of a specific appliance
- `rtmscli get-appliance-services`: Get services of a specific appliance
- `rtmscli synchronize-appliance`: Synchronize an appliance
- `rtmscli get-appliance-healthcheck`: Get or post appliance health check

## Usage Examples

### List Appliances

To list all appliances:

```
rtmscli get-appliances
```

Options:
- `--cloud-temple-id`: Specify the Cloud Temple ID (**required**)


### Get Appliance Details

To get details of a specific appliance:

```
rtmscli get-appliance-details [appliance-id]
```

Example:
```
rtmscli get-appliance-details 12345
```

### Get Appliance Services

To get services of a specific appliance:

```
rtmscli  get-appliance-services [appliance-id]
```

Example:
```
rtmscli  get-appliance-services 12345
```

### Synchronize Appliance

To synchronize an appliance:

```
rtmscli synchronize-appliance [appliance-id]
```

Example:
```
rtmscli synchronize-appliance 12345
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
rtmscli get-appliance-healthcheck [appliance-id]
```

## Common Options

All appliance commands support the following options:

- `-f, --format`: Specify output format (json, html, markdown)
- `-H, --host`: Specify the RTMS API host (default is "rtms-api.cloud-temple.com")

Example using format option:
```
rtmscli -f markdown get-appliances --cloud-temple-id=your_id
```

For more detailed information on each command and its options, use the `--help` flag:
