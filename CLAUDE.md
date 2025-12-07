# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

3focus is a full-stack web application with:
- **Backend**: Go 1.21 with Gin web framework
- **Frontend**: Vue 3 with Vite build tool
- **Database**: PostgreSQL 16
- **Containerization**: Docker Compose for development environment
- **Tool Management**: mise for version management

## Development Environment

### Starting the Development Environment

Use mise tasks to manage the Docker Compose stack:

```bash
# Start all services (recommended)
mise run up
# or directly
docker-compose up -d

# View logs
mise run logs              # All services
mise run logs:backend      # Backend only
mise run logs:frontend     # Frontend only

# Stop services
mise run down
```

### Accessing Services

- Frontend: http://localhost:5173
- Backend API: http://localhost:8080
- PostgreSQL: localhost:5432 (user: 3focus, password: 3focus_dev, db: 3focus_db)

### Tool Management (mise)

The project uses mise for:
- Tool version management
- Task runner (replaces Makefile)

Tool versions:
- Node.js: 22.11.0
- Go: 1.21

Setup:
```bash
mise install        # Install tools
mise activate       # Activate environment
mise tasks          # List available tasks
```

## Code Architecture

### Backend (Go)

Location: `backend/`

- Entry point: `backend/main.go`
- Framework: Gin (HTTP web framework)
- Hot reload: Air (configured in `backend/.air.toml`)
- API routes: Prefixed with `/api/v1`
- CORS: Configured to allow frontend at localhost:5173

Key files:
- `backend/main.go`: Server setup, routes, and handlers
- `backend/go.mod`: Go module dependencies
- `backend/.air.toml`: Hot reload configuration

### Frontend (Vue)

Location: `frontend/`

- Entry point: `frontend/src/main.js`
- Framework: Vue 3 (Composition API available)
- Build tool: Vite
- Router: Vue Router
- HTTP client: Axios

Key directories:
- `frontend/src/views/`: Page components
- `frontend/src/router/`: Route definitions
- `frontend/src/App.vue`: Root component

### Docker Configuration

- `docker-compose.yml`: Orchestrates 3 services (backend, frontend, db)
- `backend/Dockerfile`: Go development container with Air
- `frontend/Dockerfile`: Node.js container with Vite dev server

Volumes are configured for hot reload:
- Backend: `./backend:/app` with Go modules cache
- Frontend: `./frontend:/app` with node_modules excluded

## Common Commands

All common tasks are defined in `.mise.toml` and run via mise:

```bash
mise tasks                # Show all available tasks
mise run up               # Start containers
mise run down             # Stop containers
mise run dev              # Start in foreground (with logs)
mise run build            # Rebuild containers
mise run rebuild          # Rebuild and restart
mise run restart          # Restart containers
mise run logs             # View all logs
mise run logs:backend     # Backend logs
mise run logs:frontend    # Frontend logs
mise run logs:db          # Database logs
mise run clean            # Remove containers and volumes
mise run ps               # Show running containers
```

### Backend Development

```bash
# Inside backend container (already running with Air)
# Changes auto-reload

# Run locally without Docker:
cd backend
go run main.go
```

### Frontend Development

```bash
# Inside frontend container (already running with Vite)
# Changes auto-reload with HMR

# Run locally without Docker:
cd frontend
npm install
npm run dev
```

### Testing

Currently no test framework is configured. When adding tests:
- Backend: Use Go's built-in testing (`go test`)
- Frontend: Consider Vitest (built-in with Vite)

## Development Workflow

1. Setup: `mise install && mise activate`
2. Start services: `mise run up`
3. Make code changes (hot reload is enabled)
4. Access frontend at http://localhost:5173
5. API requests go to http://localhost:8080/api/v1
6. View logs: `mise run logs`
7. Stop services: `mise run down`

## Notes

- Backend uses Air for hot reload in development
- Frontend uses Vite's HMR (Hot Module Replacement)
- Database data persists in Docker volume `postgres-data`
- Go modules cache persists in Docker volume `go-modules`
