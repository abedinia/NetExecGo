# NetExecGo

## Overview
NetExecGo is a Go-based command-line tool designed for executing commands through a specified socks proxy server. This tool is particularly useful for scenarios where network commands need to be routed through a proxy for security, privacy, or access reasons. It provides real-time feedback and statistics on the execution of the command, including CPU and memory usage.

![Build Status](https://github.com/abedinia/NetExecGo/actions/workflows/ci.yml/badge.svg)
[![codecov](https://codecov.io/gh/abedinia/NetExecGo/branch/main/graph/badge.svg)](https://codecov.io/gh/abedinia/NetExecGo)
[![License](https://img.shields.io/github/license/abedinia/NetExecGo)](https://github.com/abedinia/NetExecGo/blob/main/LICENSE)



## Features
- Command execution through a specified proxy server.
- Real-time output display for both standard output and standard error.
- Signal handling for graceful interruption of running commands.
- Execution statistics including CPU usage and memory consumption.
- Display of panda emojis as a unique visual touch during execution.

## Prerequisites
To use NetExecGo, ensure you have the following installed:

- Go version 1.21.5 or later
- Dependencies listed in go.mod

## Installation
- Clone the NetExecGo repository.
- Navigate to the cloned directory.
- Run make build to compile the program.


## Usage
To use NetExecGo, run the following command:

```
go run main.go [proxy] [command] [command arguments]
```

Example:

```
go run main.go 127.0.0.1:2080 curl ifconfig.me
```

This example will execute curl ifconfig.me using the proxy at 127.0.0.1:2080.

## Commands
- build: Compiles the source code into an executable.
- run: Runs the compiled binary.
- clean: Cleans up the compiled binaries.

## Contact
For any questions or feedback regarding NetExecGo, please open an issue in the project's GitHub repository.


## Download and use

To install, download the latest release from the releases page. Move the executable file (netexecgo) to your bin directory and add it to your profile.

```
sudo chmod +x ./netexecgo
```

On macOS, you'll need to allow the security check.

```
sudo spctl --add --label "netexecgo" --type execute netexecgo
```

Once installed, you can run netexecgo with the following command:

```
netexecgo localhost:2080 curl ifconfig.me
```

This will execute the curl command to fetch the IP address on the specified port.