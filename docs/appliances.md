# Appliances

The `appliances` command allows you to manage and interact with appliances in the RTMS system.

## Usage

```
rtmscli appliances [command]
```

## Available Commands

### list

Get a list of appliances.

```
rtmscli appliances list
```

Flags:
- `--cloud-temple-id string`: Cloud Temple Tenant identifier (required)

### details

Get Appliance details.

```
rtmscli appliances details [id]
```

### services

Get Appliance services.

```
rtmscli appliances services [id]
```

### synchronize

Synchronize Appliance.

```
rtmscli appliances synchronize [id]
```

### configuration

Get appliances configuration.

```
rtmscli appliances configuration [id]
```

Flags:
- `--appliance-version string`: Appliance version (required)
- `--plugins-path string`: Absolute path to the plugins installation directory on the appliance (required)

### healthcheck

Get a last heartbeat of an appliance.

```
rtmscli appliances healthcheck [id]
```

### post-healthcheck

Posts an appliance heartbeat.

```
rtmscli appliances post-healthcheck [id]
```

Flags:
- `--appliance-version string`: Appliance version (required)
- `--nagios-operating-state string`: Nagios operating state (OK, WARNING, CRITICAL) (required)
- `--details string`: Any details to explain the current operating state

## Examples

1. List all appliances:
   ```
   rtmscli appliances list --cloud-temple-id your-cloud-temple-id
   ```

2. Get details of a specific appliance:
   ```
   rtmscli appliances details 12345
   ```

3. Get services of a specific appliance:
   ```
   rtmscli appliances services 12345
   ```

4. Synchronize a specific appliance:
   ```
   rtmscli appliances synchronize 12345
   ```

5. Get configuration of a specific appliance:
   ```
   rtmscli appliances configuration 12345 --appliance-version 1.0.0 --plugins-path /path/to/plugins
   ```

6. Get the last heartbeat of a specific appliance:
   ```
   rtmscli appliances healthcheck 12345
   ```

7. Post a heartbeat for a specific appliance:
   ```
   rtmscli appliances post-healthcheck 12345 --appliance-version 1.0.0 --nagios-operating-state OK --details "Everything is running smoothly"
   ```

For more information on a specific command, use `rtmscli appliances [command] --help`.
```
