# Nagios Module Documentation

The Nagios module allows users to interact with various aspects of Nagios integration within the RTMS system, including listing commands, managing time periods, validating plugins, and updating commands. These functionalities are crucial for maintaining and configuring the Nagios-based monitoring capabilities of the RTMS system.

This documentation covers all the verbs (subcommands) available in the Nagios module as defined in the provided code. Each command's purpose, required arguments, and available options are detailed to provide a comprehensive guide for users of the rtmscli tool.

## Base Command

```
rtmscli nagios
```

This is the base command for all Nagios-related operations. It doesn't perform any action on its own but serves as a parent for the subcommands.

## Subcommands

### 1. Get Nagios Commands

```
rtmscli nagios commands
```

Retrieves a list of Nagios commands.

Options:
- `--name`: (Optional) Filter commands by name.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 2. Get Nagios Commands Time Periods

```
rtmscli nagios time-periods
```

Retrieves a list of Nagios commands execution time periods.

Options:
- `--name`: (Optional) Filter time periods by name.
- `--alias`: (Optional) Filter time periods by alias.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 3. Validate Nagios Plugin Package

```
rtmscli nagios validate-plugin
```

Validates a Nagios plugin package.

Options:
- `--package`: (Required) JSON string of the Nagios Plugin package's composer.json.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 4. Update Nagios Commands

```
rtmscli nagios update-commands
```

Updates Nagios commands.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

Note: All commands support the global flags defined in the root command, such as `--debug` for enabling debug mode.

## Detailed Command Descriptions

### 1. Get Nagios Commands

This command retrieves a list of Nagios commands. You can optionally filter the commands by name.

Usage example:
```
rtmscli nagios commands --name "check_http"
```

### 2. Get Nagios Commands Time Periods

This command retrieves a list of execution time periods for Nagios commands. You can filter the results by name or alias.

Usage example:
```
rtmscli nagios time-periods --name "24x7" --alias "24 Hours A Day, 7 Days A Week"
```

### 3. Validate Nagios Plugin Package

This command validates a Nagios plugin package by checking its composer.json file. The package data should be provided as a JSON string.

Usage example:
```
rtmscli nagios validate-plugin --package '{"name": "example-plugin", "version": "1.0.0"}'
```

### 4. Update Nagios Commands

This command triggers an update of the Nagios commands in the RTMS system. It doesn't require any additional parameters.

Usage example:
```
rtmscli nagios update-commands
```