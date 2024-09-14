# RTMS CLI
RTMS CLI is a command-line interface for interacting with the RTMS (Real-Time Monitoring System) API. It allows you to easily manage appliances, hosts, tickets, and more directly from your terminal.

**This project is in beta version: be very vigilant when using it and report any bugs you encounter for correction**

## Version

- 20240914 - 1.00 Beta Release 

## Features

- Appliance management
- Host and host tag management
- Ticket and ticket attachment management
- User management
- Monitoring view visualization
- Flexible output formatting (JSON, HTML, Markdown)

## Prerequisites

- Go 1.16 or higher
- Access to the RTMS API

## Installation

### Installing Go

Before installing RTMS CLI, you need to have Go installed on your system. Here's how to install it on different operating systems:

#### Windows

1. Download the Go installer for Windows from [the official Go website](https://golang.org/dl/).
2. Run the installer and follow the instructions.
3. Add the Go installation path (default `C:\Go\bin`) to your PATH environment variable.

#### macOS

1. Use Homebrew (recommended):
   ```
   brew install go
   ```
   Or download the macOS installer from [the official Go website](https://golang.org/dl/).

2. If using the installer, follow the provided instructions.

#### Linux

1. Use your distribution's package manager:
   
   For Ubuntu/Debian:
   ```
   sudo apt-get update
   sudo apt-get install golang
   ```
   
   For Fedora:
   ```
   sudo dnf install golang
   ```

   Or download the tar.gz archive from [the official Go website](https://golang.org/dl/) and install it manually.

2. Configure your `GOPATH` by adding these lines to your `.bashrc` or `.zshrc` file:
   ```
   export GOPATH=$HOME/go
   export PATH=$PATH:$GOPATH/bin
   ```

### Installing RTMS CLI

Once Go is installed, you can install RTMS CLI by following these steps:

1. Clone the repository:
   ```
   git clone https://github.com/chrlesur/rtmscli.git
   ```

2. Navigate to the project directory:
   ```
   cd rtmscli
   ```

3. Compile and install the CLI:
   ```
   go build
   ```

## Configuration

Before using RTMS CLI, you need to configure your RTMS API key. Set the `RTMS_API_KEY` environment variable:

#### Windows
```
setx RTMS_API_KEY "your_api_key_here"
```

#### macOS and Linux
```
echo 'export RTMS_API_KEY="your_api_key_here"' >> ~/.bashrc
source ~/.bashrc
```
## Important Note

The Cloud Temple ID (`-c` or `--cloud-temple-id`) is a required parameter for most commands. Make sure to include it in your commands, like this:

```
rtmscli [command] -c cloud_temple_id [other options]
```

This ID is specific to your Cloud Temple environment and is necessary for the CLI to interact with the correct resources in the RTMS API.

## Basic Usage

Here are some basic usage examples of RTMS CLI:

```
# Display version
rtmscli version

# List appliances
rtmscli -c cloud_temple_id get-appliances list

# Create a ticket
rtmscli -c cloud_temple_id tickets create --name="New ticket" --description="Ticket description"

# List users with HTML output format
rtmscli -c cloud_temple_id -f html users list
```

For more information on available commands, use:
```
rtmscli --help
```

## Contributing

Contributions to this project are welcome. Please follow these steps to contribute:

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request
