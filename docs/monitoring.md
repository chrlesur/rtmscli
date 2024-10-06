# Monitoring Module Documentation

The monitoring module provides commands to manage monitoring services, notifications, and performance data in the RTMS system.
This documentation covers all the verbs (subcommands) available in the monitoring-related modules, including monitoring services, notifications, and performance, as defined in the provided code. Each command's purpose, required arguments, and available options are detailed to provide a comprehensive guide for users of the rtmscli tool.

## Base Commands

```
rtmscli monitoring-services
rtmscli monitoring
```

These are the base commands for monitoring-related operations. They don't perform any action on their own but serve as parents for the subcommands.

## Monitoring Services Subcommands

### 1. List Monitoring Services

```
rtmscli monitoring-services list
```

Retrieves a list of monitoring services.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID to filter the services.
- `--name`: (Optional) Filter services by name.
- `--status`: (Optional) Filter services by status. Can be specified multiple times.
- `--impact`: (Optional) Filter services by impact. Can be specified multiple times.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 2. Create Monitoring Service

```
rtmscli monitoring-services create
```

Creates a new monitoring service.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID for the new service.
- `--name`: (Required) Monitoring service name.
- `--appliance`: (Required) Appliance ID.
- `--host`: (Required) Host ID.
- `--template`: (Required) Template ID.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 3. Get Monitoring Service Details

```
rtmscli monitoring-services details [id]
```

Retrieves detailed information about a specific monitoring service.

Arguments:
- `id`: (Required) The ID of the monitoring service.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 4. Remove Monitoring Service

```
rtmscli monitoring-services remove [id]
```

Removes a specific monitoring service.

Arguments:
- `id`: (Required) The ID of the monitoring service to remove.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 5. Update Monitoring Service

```
rtmscli monitoring-services update [id]
```

Updates information for a specific monitoring service.

Arguments:
- `id`: (Required) The ID of the monitoring service to update.

Options:
- `--name`: (Optional) New monitoring service name.
- `--appliance`: (Optional) New appliance ID.
- `--host`: (Optional) New host ID.
- `--template`: (Optional) New template ID.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 6. Get Monitoring Service Templates

```
rtmscli monitoring-services templates
```

Retrieves a list of monitoring services templates.

Options:
- `--name`: (Optional) Filter template by name.
- `--impact`: (Optional) Filter templates by impact. Can be specified multiple times.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 7. Get Monitoring Services Stats

```
rtmscli monitoring-services stats
```

Retrieves monitoring services status and impact stats.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID to filter the stats.
- `--host-id`: (Optional) Show stats of filtered monitoring services by host.
- `--appliance-id`: (Optional) Show stats of filtered monitoring services by appliance.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

## Monitoring Service Notifications Subcommands

### 1. List Service Notifications

```
rtmscli monitoring-services notifications list-service [service-id]
```

Retrieves a list of notifications for a specific service.

Arguments:
- `service-id`: (Required) The ID of the monitoring service.

Options:
- `--attach`: (Optional) List only notifications attached to a ticket or not. Default is false.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 2. List All Notifications

```
rtmscli monitoring-services notifications list
```

Retrieves a list of all notifications.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID to filter the notifications.
- `--attach`: (Optional) List only notifications attached to a ticket or not. Default is false.
- `--staffs`: (Optional) Filter by staff identifiers. Can be specified multiple times.
- `--perimeters`: (Optional) Filter by perimeter identifiers. Can be specified multiple times.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 3. Create Notification

```
rtmscli monitoring-services notifications create
```

Creates a new notification.

Options:
- `--service-id`: (Required) Monitoring service ID.
- `--state`: (Required) State of monitoring service (OK, WARNING, CRITICAL, UNKNOWN).
- `--content`: (Required) Content of the notification.
- `--subject`: (Required) Subject that will be sent by email/sms.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 4. Get Notification Details

```
rtmscli monitoring-services notifications details [id]
```

Retrieves notification details.

Arguments:
- `id`: (Required) The ID of the notification.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 5. Get Ticket Suggestions

```
rtmscli monitoring-services notifications suggest [id]
```

Retrieves ticket suggestions for a notification.

Arguments:
- `id`: (Required) The ID of the notification.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 6. Attach Notification to Ticket

```
rtmscli monitoring-services notifications attach [id]
```

Attaches a notification to a ticket.

Arguments:
- `id`: (Required) The ID of the notification.

Options:
- `--ticket-id`: (Required) Ticket ID to attach the notification to.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 7. Detach Notification from Ticket

```
rtmscli monitoring-services notifications detach [id]
```

Detaches a notification from a ticket.

Arguments:
- `id`: (Required) The ID of the notification.

Options:
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

## Monitoring Service Performance Subcommands

### 1. Get Metric History

```
rtmscli monitoring-services performance metric-history [service-id]
```

Retrieves a list of metrics versions for a given monitoring service.

Arguments:
- `service-id`: (Required) The ID of the monitoring service.

Options:
- `--start-date`: (Optional) Start date timestamp or milliseconds of searched period.
- `--end-date`: (Optional) End date timestamp or milliseconds of searched period.
- `--metric-name`: (Optional) List of metric names. Can be specified multiple times.
- `--version-order`: (Optional) Version order: asc or desc.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 2. Get Graph Configurations

```
rtmscli monitoring-services performance graph-configurations [service-id]
```

Retrieves a list of graph configurations for a given monitoring service.

Arguments:
- `service-id`: (Required) The ID of the monitoring service.

Options:
- `--label`: (Optional) Filter graph by a string contained in label field.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

## General Monitoring Subcommands

### 1. Check RTMS Health

```
rtmscli monitoring health
```

Checks if RTMS services are healthy.

Options:
- `--integration-services`: (Optional) List of service identifiers used to test the delay of integration of monitoring results. Can be specified multiple times.
- `--integration-delay`: (Optional) Delay allowed in seconds to test the delay of integration of monitoring results.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 2. Check SLA Calculator Health

```
rtmscli monitoring sla-calculator
```

Checks if the SLA Calculator app is healthy.

Options:
- `--update-delay`: (Optional) Delay allowed in seconds between the current time and the last update of a ticket's SLA.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

Note: All commands support the global flags defined in the root command, such as `--debug` for enabling debug mode.