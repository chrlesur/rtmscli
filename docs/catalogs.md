# Catalogs Module Documentation

The catalogs module allows users to interact with various aspects of the ticket classification system, including listing catalogs, retrieving default catalogs, getting items for specific catalogs, and accessing root catalogs. These commands provide flexibility in managing and viewing the catalog structure within the RTMS system.

## Base Command

```
rtmscli catalogs
```

This is the base command for all catalog-related operations. It doesn't perform any action on its own but serves as a parent for the subcommands.

## Subcommands

### 1. List Catalogs

```
rtmscli catalogs list
```

Retrieves a list of Ticket classification catalogs and items.

Options:
- `--cloud-temple-id`: (Required) The Cloud Temple ID to filter the catalogs.
- `--available-items`: (Optional) Show classification catalog with their available items. Default is false.
- `--is-root`: (Optional) If true, only classification root catalogs will be displayed. Default is false.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 2. Get Default Catalogs

```
rtmscli catalogs defaults
```

Retrieves a list of all default ticket classification catalogs and catalog items.

Options:
- `--available-items`: (Optional) Show classification catalog with their available items. Default is false.
- `--is-root`: (Optional) If true, only classification root catalogs will be displayed. Default is false.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 3. Get Catalog Items

```
rtmscli catalogs items [catalog-id]
```

Retrieves a list of items for a specific catalog.

Arguments:
- `catalog-id`: (Required) The ID of the catalog.

Options:
- `--enabled`: (Optional) Display only enabled or disabled catalog items. Default is false.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

### 4. Get Root Catalog

```
rtmscli catalogs root
```

Retrieves the root required catalog.

Options:
- `--type`: (Required) Required Catalog type (origin, perimeter, or nature).
- `--available-items`: (Optional) Display associated catalog items. Default is false.
- `--format`: (Optional) Output format (json, text, html, markdown). Default is json.

Note: All commands support the global flags defined in the root command, such as `--debug` for enabling debug mode.