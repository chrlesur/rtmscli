# View Management

The RTMS CLI provides a command for listing items in monitoring views. This document outlines the available command and its usage.

## Available Command

- `rtmscli views list`: List hosts, services, or templates in a monitoring view

## View Types

RTMS supports three types of views:

1. Host views
2. Service views
3. Template views

## Usage Examples

### List View Items

To list items in a monitoring view:

```
rtmscli views list [type] [id] [flags]
```

Parameters:
- `type`: The type of view (host, service, or template)
- `id`: The identifier of the view

Options:
- `--order`: Data sort order (ASC or DESC, default is DESC)
- `--order-by`: Attribute of the element on which to order (default is "id")
- `--page`: Current page (default is 1)
- `--items-per-page`: Number of items on a single page (default is 100, max 500)

### List Host View Items

To list items in a host view:

```
rtmscli views list host [view-id] [flags]
```

Example:
```
rtmscli views list host 12345 --order=ASC --order-by=name --page=2 --items-per-page=50
```

This command will list the hosts in the monitoring view with ID 12345, sorted by name in ascending order, displaying the second page with 50 items per page.

### List Service View Items

To list items in a service view:

```
rtmscli views list service [view-id] [flags]
```

Example:
```
rtmscli views list service 67890 --order=DESC --order-by=status
```

This command will list the services in the monitoring view with ID 67890, sorted by status in descending order.

### List Template View Items

To list items in a template view:

```
rtmscli views list template [view-id] [flags]
```

Example:
```
rtmscli views list template 54321 --items-per-page=200
```

This command will list the templates in the monitoring view with ID 54321, displaying 200 items per page.

## Common Options

The views command supports the following common options:

- `-f, --format`: Specify output format (json, html, markdown)
- `-H, --host`: Specify the RTMS API host (default is "rtms-api.cloud-temple.com")

Example using format option:
```
rtmscli -f markdown views list host 12345
```

For more detailed information on the command and its options, use the `--help` flag:

```
rtmscli views --help
rtmscli views list --help
```

## Notes

- The view type must be one of "host", "service", or "template". An error will be returned if an invalid type is provided.
- The view ID is required and must be a valid identifier for an existing view in the RTMS system.
- Pagination is supported through the `--page` and `--items-per-page` options, allowing you to navigate through large result sets.
- Sorting can be customized using the `--order` and `--order-by` options, allowing you to sort the results based on different attributes and in ascending or descending order.

