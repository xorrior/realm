---
title: Tavern
tags:
 - Admin Guide
description: Information on managing a Tavern deployment.
permalink: admin-guide/tavern
---

## Overview

Tavern is a teamserver for Realm, providing a UI to control deployments and implants during an engagement. The majority of Tavern's functionality is exposed through a GraphQL API, which is used by both implants and the UI.

If you would like to help contribute to Tavern, please take a look at our [open issues](https://github.com/spellshift/realm/issues?q=is%3Aopen+is%3Aissue+label%3Atavern).

## Deployment

Tavern can be deployed using Docker Compose or run directly. This section covers both methods.

### Quick Start with Docker Compose

The simplest way to deploy Tavern is using Docker Compose, which sets up both Tavern and MySQL automatically.

#### 1. Clone the Repository

```bash
git clone https://github.com/spellshift/realm.git
cd realm
```

#### 2. Start the Services

```bash
docker-compose up -d --build
```

This will:
- Build the Tavern container from source
- Start a MySQL 8.0 database
- Start Tavern on port 8000

#### 3. Create Your Admin Account

Open `http://localhost:8000` in your browser. On first visit, you'll be redirected to the setup page to create your administrator account.

1. Enter a username (3-25 characters)
2. Enter a password (minimum 8 characters)
3. Click "Create Account"

You'll be automatically logged in and redirected to the dashboard.

#### 4. Verify the Deployment

```bash
# Check server status
curl http://localhost:8000/status

# Get the server's public key (needed for agent configuration)
./bin/getpubkey.sh
```

#### Docker Compose Commands

```bash
# Start services in background
docker-compose up -d

# View logs
docker-compose logs -f tavern

# Stop services
docker-compose down

# Reset everything (delete all data)
docker-compose down -v

# Rebuild after code changes
docker-compose up -d --build
```

### Manual Deployment

For more control over your deployment, you can run Tavern directly.

#### Prerequisites

- Go 1.24+
- MySQL 8.0 (optional, SQLite used by default)
- Node.js 18+ (for frontend development)

#### 1. Start MySQL (Optional)

```bash
docker run -d \
  --name tavern-mysql \
  -e MYSQL_ROOT_PASSWORD=changeme \
  -e MYSQL_DATABASE=tavern \
  -e MYSQL_USER=tavern \
  -e MYSQL_PASSWORD=tavern \
  -p 3306:3306 \
  mysql:8.0
```

#### 2. Build the Frontend

```bash
cd tavern/internal/www
npm install
npm run build
cd ../../..
```

#### 3. Run Tavern

```bash
# With MySQL
MYSQL_ADDR=127.0.0.1:3306 \
MYSQL_USER=tavern \
MYSQL_PASSWD=tavern \
MYSQL_DB=tavern \
SECRETS_FILE_PATH=/tmp/tavern-secrets \
go run ./tavern

# With SQLite (development only)
SECRETS_FILE_PATH=/tmp/tavern-secrets go run ./tavern
```

#### 4. Access Tavern

Open `http://localhost:8000` and create your admin account.

## Authentication

Tavern uses username/password authentication with bcrypt password hashing.

### First User Setup (Trust on First Use)

When no users exist in the database, Tavern enables a setup flow:

1. Visit the Tavern URL
2. You'll be redirected to `/setup`
3. Create the first user account
4. This user is automatically an activated administrator

### Subsequent Users

Administrators can create additional users via the GraphQL API:

```graphql
mutation {
  createUser(input: {
    name: "newuser"
    password: "securepassword123"
    isAdmin: false
    isActivated: true
  }) {
    id
    name
  }
}
```

### Changing Passwords

Users can change their own password:

```graphql
mutation {
  changePassword(
    currentPassword: "oldpassword"
    newPassword: "newpassword123"
  )
}
```

### CLI Application Authentication

Tavern supports `access_tokens` for CLI authentication. See the CLI Authentication section below for details.

## Redirectors

By default Tavern only supports gRPC connections directly to the server. To enable additional protocols or additional IPs/domain names in your callbacks, utilize Tavern redirectors which receive traffic using a specific protocol (like HTTP/1.1 or DNS) and then forward it to an upstream Tavern server over gRPC.

### Available Redirectors

Realm includes three built-in redirector implementations:

- **`grpc`** - Direct gRPC passthrough redirector
- **`http1`** - HTTP/1.1 to gRPC redirector
- **`dns`** - DNS to gRPC redirector

### Basic Usage

List available redirectors:

```bash
tavern redirector list
```

Start a redirector:

```bash
tavern redirector --transport <TRANSPORT> --listen <LISTEN_ADDR> <UPSTREAM_GRPC_ADDR>
```

### HTTP/1.1 Redirector

The HTTP/1.1 redirector accepts HTTP/1.1 traffic from agents and forwards it to an upstream gRPC server.

```bash
# Start HTTP/1.1 redirector on port 8080
tavern redirector --transport http1 --listen ":8080" localhost:8000
```

### DNS Redirector

The DNS redirector tunnels C2 traffic through DNS queries and responses, providing a covert communication channel. It supports TXT, A, and AAAA record types.

```bash
# Start DNS redirector on UDP port 53 for domain c2.example.com
tavern redirector --transport dns --listen "0.0.0.0:53?domain=c2.example.com" localhost:8000

# Support multiple domains
tavern redirector --transport dns --listen "0.0.0.0:53?domain=c2.example.com&domain=backup.example.com" localhost:8000
```

**DNS Configuration Requirements:**

1. Configure your DNS server to delegate queries for your C2 domain to the redirector IP
2. Or run the redirector as your authoritative DNS server for the domain
3. Ensure UDP port 53 is accessible

**Server Behavior:**

- **Benign responses**: Non-C2 queries to A records return `0.0.0.0` instead of NXDOMAIN to avoid breaking recursive DNS lookups (e.g., when using Cloudflare as an intermediary)
- **Conversation tracking**: The server tracks up to 200,000 concurrent conversations
- **Timeout management**: Conversations timeout after 15 minutes of inactivity (reduced to 5 minutes when at capacity)
- **Maximum data size**: 50MB per request

See the [DNS Transport Configuration](/user-guide/imix#dns-transport-configuration) section in the Imix user guide for more details on agent-side configuration.

### gRPC Redirector

The gRPC redirector provides a passthrough for gRPC traffic, useful for deploying multiple Tavern endpoints or load balancing.

```bash
# Start gRPC redirector on port 9000
tavern redirector --transport grpc --listen ":9000" localhost:8000
```

## Configuration

### Environment Variables

#### HTTP Server

| Env Var | Description | Default |
| ------- | ----------- | ------- |
| HTTP_LISTEN_ADDR | Address for the HTTP server to bind to | `0.0.0.0:8000` |
| HTTP_METRICS_LISTEN_ADDR | Address for the metrics HTTP server | `127.0.0.1:8080` |

#### MySQL Database

| Env Var | Description | Default |
| ------- | ----------- | ------- |
| MYSQL_ADDR | Address of the MySQL server (e.g. `host[:port]`) | N/A (uses SQLite) |
| MYSQL_NET | Network type (e.g. unix) | `tcp` |
| MYSQL_USER | User to authenticate with | `root` |
| MYSQL_PASSWD | Password to authenticate with | (empty) |
| MYSQL_DB | Name of the database to use | `tavern` |
| DB_MAX_IDLE_CONNS | Max idle MySQL connections | `10` |
| DB_MAX_OPEN_CONNS | Max open MySQL connections | `100` |
| DB_MAX_CONN_LIFETIME | Max connection lifetime (seconds) | `3600` |

#### Secrets

| Env Var | Description | Default |
| ------- | ----------- | ------- |
| SECRETS_FILE_PATH | Path for file-based secrets storage | (required) |

#### Features

| Env Var | Description | Default |
| ------- | ----------- | ------- |
| ENABLE_METRICS | Enable the `/metrics` Prometheus endpoint | Disabled |
| ENABLE_PPROF | Enable performance profiling (development only) | Disabled |
| ENABLE_TEST_DATA | Populate database with test data | Disabled |
| DISABLE_DEFAULT_TOMES | Skip importing default tomes | Disabled |
| ENABLE_DEBUG_LOGGING | Enable verbose debug logs | Disabled |
| ENABLE_JSON_LOGGING | Output logs in JSON format | Disabled |

### Secrets

Tavern uses file-based secrets for cryptographic key management. Set the `SECRETS_FILE_PATH` environment variable to specify where keys should be stored:

```bash
SECRETS_FILE_PATH="/data/secrets" go run ./tavern
```

### MySQL

By default, Tavern uses an in-memory SQLite database. For production deployments, use MySQL:

```bash
MYSQL_ADDR=127.0.0.1:3306 \
MYSQL_USER=tavern \
MYSQL_PASSWD=securepassword \
MYSQL_DB=tavern \
go run ./tavern
```

**MySQL Notes:**
- MySQL 8.0 is required
- The database must be created before starting Tavern
- Tavern will automatically run migrations on startup

### Metrics

Enable Prometheus metrics collection:

```bash
ENABLE_METRICS=1 go run ./tavern
```

Metrics are available at the `/metrics` endpoint on the metrics server (default `127.0.0.1:8080`). This is hosted on a separate server to allow restricting access, as the endpoint is unauthenticated.

### Test Data

Populate the database with sample data for testing:

```bash
ENABLE_TEST_DATA=1 go run ./tavern
```

## CLI Authentication

Tavern supports `access_tokens` for CLI application authentication. This allows CLI tools to authenticate using a user's existing browser session.

```golang
package main

import (
 "context"
 "fmt"
 "net/http"
 "time"

 "github.com/pkg/browser"
 "realm.pub/tavern/cli/auth"
)

func main() {
 ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
 defer cancel()

 tavernURL := "http://127.0.0.1:8000"
 browser := auth.BrowserFunc(browser.OpenURL)

 token, err := auth.Authenticate(ctx, browser, tavernURL)
 if err != nil {
  panic(err)
 }

 req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/status", tavernURL), nil)
 if err != nil {
  panic(err)
 }

 token.Authenticate(req)
 // Send the request...
}
```

## Docker Images

Tavern publishes several Docker image tags:

| Tag | Description |
| --- | ----------- |
| `vX.Y.Z` | Specific release version |
| `latest` | Latest stable release |
| `edge`, `main` | Latest from main branch (development) |
| `sha-<hash>` | Specific git commit |

### Building Custom Images

```bash
# Build
cd realm
docker build --tag tavern:custom --file ./docker/tavern.Dockerfile .

# Push to Docker Hub
docker tag tavern:custom yourusername/tavern:custom
docker push yourusername/tavern:custom
```

## GraphQL API

### Playground

Explore the GraphQL API at the `/playground` endpoint. This provides an interactive environment for testing queries and mutations.

### Example Queries

#### List Shells

```graphql
query shells {
  shells {
    edges {
      node {
        id
        beacon {
          id
          host {
            name
          }
        }
      }
    }
  }
}
```

#### List Tomes

```graphql
query listtomes {
  tomes {
    edges {
      node {
        id
        name
        tactic
        description
        paramDefs
      }
    }
  }
}
```

#### Create Asset Link

```graphql
mutation tempLink {
  createLink(input: {
    expiresAt: "2026-02-02T21:33:18Z"
    assetID: 4
  }) {
    path
  }
}
```

## CDN HTTP API

### Upload - POST /cdn/upload

Upload files to the Tavern CDN (requires authentication):

```bash
curl --cookie "auth-session=TOKEN" \
  -F "fileName=myfile" \
  -F "fileContent=@/path/to/file" \
  https://tavern.example.com/cdn/upload
```

### Download - GET /cdn/{path}

Download files via generated links (unauthenticated):

```python
# From Eldritch
f = http.get(f"https://tavern.example.com/cdn/{path}", allow_insecure=True)
```

### Create Download Link

```graphql
mutation {
  createLink(input: {
    expiresAt: "2026-02-02T21:33:18Z"
    downloadsRemaining: 10
    path: "myfile"
    assetID: 4
  }) {
    path
  }
}
```
