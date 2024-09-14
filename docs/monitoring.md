# Monitoring in RTMS CLI

RTMS CLI provides comprehensive monitoring capabilities, including system health checks, monitoring services management, and notification handling. This document outlines the available commands and their usage.

## Table of Contents
1. [System Health](#system-health)
2. [Monitoring Services](#monitoring-services)
3. [Notifications](#notifications)

## System Health

RTMS CLI allows you to check the health of the RTMS services and the SLA Calculator.

### Check RTMS Services Health

To check the health of RTMS services:

```
rtmscli monitoring health [flags]
```

Options:
- `--integration-services`: List of service identifiers to test the delay of integration of monitoring results
- `--integration-delay`: Delay allowed in seconds to test the delay of integration of monitoring results

Example:
```
rtmscli monitoring health --integration-services=1,2,3 --integration-delay=60
```

### Check SLA Calculator Health

To check the health of the SLA Calculator:

```
rtmscli monitoring sla-calculator [flags]
```

Options:
- `--update-delay`: Delay allowed in seconds between the current time and the last update of a ticket's SLA

Example:
```
rtmscli monitoring sla-calculator --update-delay=3600
```

## Monitoring Services

RTMS CLI provides commands to manage monitoring services, including listing, creating, updating, and removing services.

### List Monitoring Services

To list all monitoring services:

```
rtmscli monitoring-services list [flags]
```

Options:
- `--name`: Filter services by name
- `--status`: Filter services by status (OK, WARNING, CRITICAL, UNKNOWN, PENDING)
- `--impact`: Filter services by impact (Availability, Performance, Information, Security)

Example:
```
rtmscli monitoring-services list --name=webserver --status=WARNING
```

### Create Monitoring Service

To create a new monitoring service:

```
rtmscli monitoring-services create [flags]
```

Required flags:
- `--name`: Monitoring service name
- `--appliance`: Appliance ID
- `--host`: Host ID
- `--template`: Template ID

Example:
```
rtmscli monitoring-services create --name=cpu_check --appliance=1 --host=2 --template=3
```

### Get Monitoring Service Details

To get details of a specific monitoring service:

```
rtmscli monitoring-services details [service-id]
```

Example:
```
rtmscli monitoring-services details 12345
```

### Update Monitoring Service

To update a monitoring service:

```
rtmscli monitoring-services update [service-id] [flags]
```

Example:
```
rtmscli monitoring-services update 12345 --name=new_cpu_check
```

### Remove Monitoring Service

To remove a monitoring service:

```
rtmscli monitoring-services remove [service-id]
```

Example:
```
rtmscli monitoring-services remove 12345
```

### Get Monitoring Service Templates

To get a list of monitoring service templates:

```
rtmscli monitoring-services templates [flags]
```

Options:
- `--name`: Filter template by name
- `--impact`: Filter templates by impact

Example:
```
rtmscli monitoring-services templates --name=cpu --impact=Performance
```

### Get Monitoring Services Statistics

To get statistics about monitoring services:

```
rtmscli monitoring-services stats [flags]
```

Options:
- `--host-id`: Show stats of filtered monitoring services by host
- `--appliance-id`: Show stats of filtered monitoring services by appliance

Example:
```
rtmscli monitoring-services stats --host-id=1234
```

## Notifications

RTMS CLI provides commands to manage notifications related to monitoring services.

### List Notifications

To list all notifications:

```
rtmscli monitoring-services notifications list [flags]
```

Options:
- `--attach`: List only notifications attached to a ticket or not
- `--staffs`: Filter by staff identifiers
- `--perimeters`: Filter by perimeter identifiers

Example:
```
rtmscli monitoring-services notifications list --attach --staffs=1,2,3
```

### Create Notification

To create a new notification:

```
rtmscli monitoring-services notifications create [flags]
```

Required flags:
- `--service-id`: Monitoring service ID
- `--state`: State of monitoring service (OK, WARNING, CRITICAL, UNKNOWN)
- `--content`: Content of the notification
- `--subject`: Subject that will be sent by email/sms

Example:
```
rtmscli monitoring-services notifications create --service-id=1234 --state=WARNING --content="High CPU usage" --subject="CPU Warning"
```

### Get Notification Details

To get details of a specific notification:

```
rtmscli monitoring-services notifications details [notification-id]
```

Example:
```
rtmscli monitoring-services notifications details 5678
```

### Attach Notification to Ticket

To attach a notification to a ticket:

```
rtmscli monitoring-services notifications attach [notification-id] --ticket-id=[ticket-id]
```

Example:
```
rtmscli monitoring-services notifications attach 5678 --ticket-id=9012
```

### Detach Notification from Ticket

To detach a notification from a ticket:

```
rtmscli monitoring-services notifications detach [notification-id]
```

Example:
```
rtmscli monitoring-services notifications detach 5678
```

## Common Options

All monitoring commands support the following options:

- `-f, --format`: Specify output format (json, html, markdown)
- `-H, --host`: Specify the RTMS API host (default is "rtms-api.cloud-temple.com")

For more detailed information on each command and its options, use the `--help` flag:

```
rtmscli monitoring --help
rtmscli monitoring-services --help
rtmscli monitoring-services [command] --help
```

