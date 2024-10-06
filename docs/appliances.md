# Appliances Module Documentation

The appliances module provides commands to manage and interact with appliances in the RTMS system.

## Base Command

```
rtmscli appliances
```

This is the base command for all appliance-related operations. It doesn't perform any action on its own but serves as a parent for the subcommands.

## Subcommands

### 1. List Appliances

```
rtmscli appliances list
```

Retrieves a list of appliances.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID to filter the appliances.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 2. Get Appliance Details

```
rtmscli appliances details [id]
```

Retrieves detailed information about a specific appliance.

Arguments:
- `id`: (Required) The ID of the appliance.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 3. Get Appliance Services

```
rtmscli appliances services [id]
```

Retrieves the services associated with a specific appliance.

Arguments:
- `id`: (Required) The ID of the appliance.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 4. Synchronize Appliance

```
rtmscli appliances synchronize [id]
```

Initiates a synchronization process for a specific appliance.

Arguments:
- `id`: (Required) The ID of the appliance to synchronize.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 5. Get Appliance Configuration

```
rtmscli appliances configuration [id]
```

Retrieves the configuration for a specific appliance.

Arguments:
- `id`: (Required) The ID of the appliance.

Options:
- `--appliance-version`: (Required) The version of the appliance.
- `--plugins-path`: (Required) The absolute path to the plugins installation directory on the appliance.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 6. Get Appliance Health Check

```
rtmscli appliances healthcheck [id]
```

Retrieves the last heartbeat (health check) of a specific appliance.

Arguments:
- `id`: (Required) The ID of the appliance.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 7. Post Appliance Health Check

```
rtmscli appliances post-healthcheck [id]
```

Posts a heartbeat (health check) for a specific appliance.

Arguments:
- `id`: (Required) The ID of the appliance.

Options:
- `--appliance-version`: (Required) The version of the appliance.
- `--nagios-operating-state`: (Required) The Nagios operating state (OK, WARNING, CRITICAL).
- `--details`: (Optional) Any details to explain the current operating state.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

Note: All commands support the global flags defined in the root command, such as `--debug` for enabling debug mode.