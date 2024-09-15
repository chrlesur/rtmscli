# Catalogs

The `catalogs` command allows you to manage and interact with ticket classification catalogs in the RTMS system.

## Usage

```
rtmscli catalogs [command]
```

## Available Commands

### list

Get a list of Ticket classification catalogs and items.

```
rtmscli catalogs list
```

Flags:
- `--cloud-temple-id string`: Cloud Temple Tenant identifier (required)
- `--available-items`: Show classification catalog with their available items
- `--is-root`: If true, only classification root catalogs will be displayed

### defaults

Get a list of all default ticket classification catalogs and catalog items.

```
rtmscli catalogs defaults
```

Flags:
- `--available-items`: Show classification catalog with their available items
- `--is-root`: If true, only classification root catalogs will be displayed

### items

Get a list of items for a catalog.

```
rtmscli catalogs items [catalog-id]
```

Flags:
- `--enabled`: Display only enabled or disabled catalog items

### root

Get the root required catalog.

```
rtmscli catalogs root
```

Flags:
- `--type string`: Required Catalog type (origin, perimeter, or nature) (required)
- `--available-items`: Display associated catalog items

## Examples

1. List all catalogs:
   ```
   rtmscli catalogs list --cloud-temple-id your-cloud-temple-id
   ```

2. List all catalogs with available items:
   ```
   rtmscli catalogs list --cloud-temple-id your-cloud-temple-id --available-items
   ```

3. Get default catalogs:
   ```
   rtmscli catalogs defaults
   ```

4. Get items for a specific catalog:
   ```
   rtmscli catalogs items 12345
   ```

5. Get items for a specific catalog, showing only enabled items:
   ```
   rtmscli catalogs items 12345 --enabled
   ```

6. Get the root catalog of type "origin":
   ```
   rtmscli catalogs root --type origin
   ```

7. Get the root catalog of type "perimeter" with associated items:
   ```
   rtmscli catalogs root --type perimeter --available-items
   ```

For more information on a specific command, use `rtmscli catalogs [command] --help`.
```
