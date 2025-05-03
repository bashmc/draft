# Draft - Project Management System

**🚧 This project is currently under development 🚧**

## Overview
Draft is a collaborative project management system built with Go and Vue.js, designed to help teams organize and track their work efficiently.

## Tech Stack
### Backend
- Go
- PostgreSQL
- Chi router for HTTP routing
- pgx for PostgreSQL driver
- goose for database migrations

### Frontend
- Vue 3
- TypeScript
- Vite
- HTML/CSS

## Features
- User Management
  - Account creation with email verification
  - Secure password handling
  - Profile management
- Workspaces
  - Create and manage collaborative spaces
  - User memberships and permissions
- Projects (Coming soon)
  - Project creation and management
  - Task tracking
  - Team collaboration
- Tasks (Coming soon)
  - Task creation and assignment
  - Priority levels
  - Status tracking

## Development Setup

### Prerequisites
- Go 1.24 or higher
- PostgreSQL
- Node.js and npm
- Task (task runner)

### Environment Variables
Copy `.env.example` to `.env` and configure:
```sh
ENV=dev
DB_URL=postgres://user:password@localhost:5432/draft?sslmode=disable
MAIL_HOST=your-mail-host
MAIL_TOKEN=your-mail-token
SENDER_EMAIL=noreply@yourdomain.com
SENDER_NAME=Draft
```

### Run migrations
```sh

task up    # Apply migrations
task down  # Rollback migrations
task reset # Reset database

```

### Running the application
```sh
# Start backend server
task run

# Start frontend development server (from ui directory)
cd ui
npm install
npm run dev
```

### Testing
```sh
# Run database migrations for test
task test-up

# Run tests
go test ./...

```

### Project Structure
```
.
├── cmd/          # Application entrypoints
├── handlers/     # HTTP request handlers
├── mail/         # Email service
├── migrations/   # Database migrations
├── models/       # Data models
├── postgres/     # Database implementations
├── services/     # Business logic
└── ui/          # Frontend application
```
