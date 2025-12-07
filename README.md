# 3focus

A full-stack application built with Go (backend) and Vue.js (frontend), orchestrated with Docker.

## Tech Stack

- **Backend**: Go 1.21 + Gin framework
- **Frontend**: Vue 3 + Vite
- **Database**: PostgreSQL 16
- **Container**: Docker + Docker Compose
- **Tool Version Manager**: mise

## Prerequisites

- Docker and Docker Compose
- mise (recommended for task management and tool versions)

## Quick Start

1. Install mise: https://mise.jdx.dev/getting-started.html

2. Install tools and activate environment:
```bash
mise install
mise activate
```

3. Start all services:
```bash
mise run up
# or directly:
docker-compose up -d
```

4. Access the application:
   - Frontend: http://localhost:5173
   - Backend API: http://localhost:8080
   - Database: localhost:5432

5. View logs:
```bash
mise run logs              # All services
mise run logs:backend      # Backend only
mise run logs:frontend     # Frontend only
```

6. Stop services:
```bash
mise run down
```

## Development

### Backend (Go)

```bash
cd backend
go run main.go
```

- Hot reload is enabled with Air in Docker
- API endpoints: http://localhost:8080/api/v1

### Frontend (Vue)

```bash
cd frontend
npm install
npm run dev
```

- Vite dev server with hot module replacement
- App: http://localhost:5173

### Database

Default PostgreSQL connection:
- Host: localhost
- Port: 5432
- User: 3focus
- Password: 3focus_dev
- Database: 3focus_db

## Available Commands

See all available mise tasks:
```bash
mise tasks
```

Common commands:
```bash
mise run up            # Start containers
mise run down          # Stop containers
mise run dev           # Start in foreground (with logs)
mise run build         # Build containers
mise run rebuild       # Rebuild and restart
mise run restart       # Restart containers
mise run logs          # View all logs
mise run logs:backend  # View backend logs
mise run logs:frontend # View frontend logs
mise run logs:db       # View database logs
mise run clean         # Remove containers and volumes
mise run ps            # Show running containers
```

## Project Structure

```
.
├── backend/           # Go backend application
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
├── frontend/          # Vue frontend application
│   ├── src/
│   ├── package.json
│   └── Dockerfile
├── docker-compose.yml
└── .mise.toml        # Tool versions and task definitions
```
