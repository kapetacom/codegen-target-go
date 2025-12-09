# Model Context Protocol (MCP) Server

This block implements a Model Context Protocol (MCP) server using the [mcp-go](https://github.com/mark3labs/mcp-go) library. MCP enables seamless integration between LLM applications and external data sources.

## Features

- MCP server with streamable HTTP transport support
- Support for registering tools, resources, and prompts
- Integration with Kapeta configuration system
- Docker-ready deployment
- Full bidirectional communication support

## Generated Files

- `main.go` - The main MCP server implementation
- `go.mod` - Go module dependencies
- `Dockerfile` - Containerization configuration
- Other supporting files

## How to Use

1. Define your MCP tools, resources, and prompts in `main.go`
2. Build and run the server
3. Connect MCP-compatible clients via HTTP transport

## Transport Options

The server supports HTTP transport only:
- **Streamable HTTP transport**: Server listens on configured port for HTTP-based connections

## Configuration

The server integrates with Kapeta's configuration system for managing settings and connections to other services.
The server port can be configured via the "mcp" server port setting.