# Views Module Documentation

The Views module provides commands to manage monitoring views in the RTMS system. The Views module allows users to interact with monitoring views in the RTMS system, providing a way to list and analyze the components (hosts, services, or templates) that make up these views. This functionality is crucial for understanding the structure of monitoring configurations and for managing large-scale monitoring setups.

This documentation covers the single verb (subcommand) available in the Views module as defined in the provided code. The command's purpose, required arguments, and available options are detailed to provide a comprehensive guide for users of the rtmscli tool.


## Base Command

```
rtmscli views
```

This is the base command for all view-related operations. It doesn't perform any action on its own but serves as a parent for the subcommands.

## Subcommand

### 1. List View Items

```
rtmscli views list [type] [id]
```

Lists hosts, services, or templates in a monitoring view.

Arguments:
- `type`: (Required) The type of items to list. Must be one of: 'host', 'service', or 'template'.
- `id`: (Required) The ID of the view.

Options:
- `--order`: (Optional) Data sort order (ASC, DESC). Default is "DESC".
- `--order-by`: (Optional) Attribute of the element on which to order. Default is "id".
- `--page`: (Optional) Current page. Default is 1.
- `--items-per-page`: (Optional) Number of items on a single page (max 500). Default is 100.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

Note: All commands support the global flags defined in the root command, such as `--debug` for enabling debug mode.

## Detailed Command Description

### List View Items

This command retrieves a list of items (hosts, services, or templates) for a specific monitoring view. It allows you to specify the type of items you want to list and provides options for sorting and pagination.

Usage example:
```
rtmscli views list host view-123 --order ASC --order-by name --page 2 --items-per-page 50
```

In this example:
- We're listing host items for the view with ID "view-123".
- The results will be sorted in ascending order by name.
- We're requesting the second page of results.
- Each page will contain 50 items.

The command will return an error if an invalid view type is provided (i.e., anything other than 'host', 'service', or 'template').

This command is useful for:
1. Reviewing the contents of a monitoring view.
2. Troubleshooting issues related to specific hosts, services, or templates within a view.
3. Generating reports on the composition of monitoring views.

The pagination options (`--page` and `--items-per-page`) are particularly useful when dealing with large views that may contain hundreds or thousands of items. By using these options, you can retrieve the data in manageable chunks.

The sorting options (`--order` and `--order-by`) allow you to customize how the results are presented, which can be helpful when looking for specific items or when you need the data in a particular order for reporting or analysis purposes.