# MCP (Model Context Protocol) Server Block

This block template generates a Go-based Model Context Protocol (MCP) server. MCP enables seamless integration between LLM applications and external data sources, tools, and resources.

## What is MCP?

The Model Context Protocol (MCP) is an open protocol that allows LLM applications to connect with external systems. It provides standardized ways to:
- Access resources (data, files, APIs, etc.)
- Execute tools and actions
- Retrieve prompts and templates

## Features

- MCP server implementation using [mcp-go](https://github.com/mark3labs/mcp-go)
- HTTP transport support only
- Full bidirectional communication support
- Integration with Kapeta configuration system
- Ready for Docker deployment
- Hot reloading during development

## Getting Started

1. The server is pre-configured with an example tool, resource, and prompt
2. Customize `main.go` and `generated/mcpHandlers.go` to add your own tools, resources, and prompts
3. Use the development scripts to run with hot reloading
4. Build and deploy using Docker
5. Connect using HTTP clients

## Transport Options

The server supports HTTP transport only:
- **Streamable HTTP transport**: Server listens on configured port for HTTP-based connections

## Development

```bash
# Run with HTTP transport
./start-dev.sh  # Linux/Mac
start-dev.cmd   # Windows
```

## Production

```bash
# Build and run with HTTP transport
./start-prod.sh  # Linux/Mac
start-prod.cmd   # Windows
```

## Docker

```bash
# Build and run in Docker with HTTP transport
docker build -t mcp-server .
docker run -it -p 8080:8080 mcp-server
```