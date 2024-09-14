# Nagios Management

RTMS CLI provides several commands for managing Nagios-related functionalities. This document outlines the available commands and their usage.

## Available Commands

- `rtmscli nagios commands`: Get a list of Nagios commands
- `rtmscli nagios time-periods`: Get Nagios commands execution time periods list
- `rtmscli nagios validate-plugin`: Validate a Nagios plugin package
- `rtmscli nagios update-commands`: Update Nagios commands

## Usage Examples

### List Nagios Commands

To list all Nagios commands:

```
rtmscli nagios commands
```

Options:
- `--name`: Filter commands by name

Example:
```
rtmscli nagios commands --name=check_http
```

### Get Nagios Time Periods

To get the list of Nagios command execution time periods:

```
rtmscli nagios time-periods
```

Options:
- `--name`: Filter time periods by name
- `--alias`: Filter time periods by alias
- `--cloud-temple-id`: Specify the Cloud Temple ID (optional)

Example:
```
rtmscli nagios time-periods --name=24x7 --cloud-temple-id=your_id
```

### Validate Nagios Plugin Package

To validate a Nagios plugin package:

```
rtmscli nagios validate-plugin --package='{"name": "example/plugin", ...}'
```

The `--package` option should contain a valid JSON string representing the plugin's composer.json file.

Example:
```
rtmscli nagios validate-plugin --package='{"name": "external/nagios-plugins", "description": "Check CPU usage over SNMP", "version": "2.3.3", "type": "project", "require": {"cloud-temple/appliance": "^1.0"}, "bin": ["check_ping.sh"], "extra": {"commands": {"check_ping.sh": {"readme": "check_ping.md", "pluginArgs": "-H $HOSTADDRESS$ -w $ARG1$ -c $ARG2$ -p 5 -t 10"}}}}'
```

### Update Nagios Commands

To update Nagios commands:

```
rtmscli nagios update-commands
```

This command doesn't require any additional parameters.

## Common Options

All Nagios commands support the following options:

- `-f, --format`: Specify output format (json, html, markdown)
- `-H, --host`: Specify the RTMS API host (default is "rtms-api.cloud-temple.com")

Example using format option:
```
rtmscli -f markdown nagios commands
```

## Nagios Plugin Package Validation

The `validate-plugin` command checks the provided composer.json against the Cloud Temple schema for Nagios plugins. Here are some key points to remember when creating a plugin package:

1. The `name` field should follow the format "vendor/package-name".
2. The `type` field should be set to "project".
3. The `require` section should include "cloud-temple/appliance": "^1.0".
4. The `bin` section should list all executable scripts in the package.
5. The `extra.commands` section should describe each command, including its readme file and plugin arguments.

For more detailed information on each command and its options, use the `--help` flag:

```
rtmscli nagios --help
rtmscli nagios [command] --help
```

