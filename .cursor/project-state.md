# Project State

Last Updated: 2025-10-31

## Current Focus
- Working on: [Current task description]
- Branch: v5
- Recent changes: [Brief summary of recent modifications]

## Active Tasks

### Task 1: [Task Name]
- **Status**: In Progress / Completed / Blocked
- **Files Modified**:
  - `file1.go`
  - `file2.go`
- **Next Steps**:
  1. [Step 1]
  2. [Step 2]
- **Blockers**: [Any blockers or dependencies]

## Project Structure

This is a Telegram bot library (telebot/tb) written in Go:

```
/Users/jhin/projects/tb/
├── telegram/          # Telegram API types and structures
├── bot/              # Core bot implementation
├── params/           # Request parameters and options
├── request/          # HTTP request handling
├── sendables/        # Sendable message types
├── files/            # File handling
├── censorship/       # Content filtering
├── errors/           # Error handling
└── examples/         # Usage examples
```

### Key Directories
- **telegram/**: Telegram Bot API type definitions (messages, users, chats, etc.)
- **params/**: Parameter builders for API calls
- **bot/**: Main bot implementation and handlers
- **examples/**: Demonstration code for library usage

## Key Patterns and Conventions

### Code Style
- Go standard formatting (gofmt)
- Descriptive variable names
- Comprehensive error handling
- Interface-based design

### Error Handling
- Use custom error types from `errors/` package
- Always check and handle errors explicitly
- Provide context in error messages

### Testing Approach
- Unit tests in `_test.go` files
- Table-driven tests where appropriate
- Mock external dependencies

### Naming Conventions
- PascalCase for exported types and functions
- camelCase for internal functions
- Descriptive names over brevity

### Architecture
- Clean separation between Telegram API types and bot logic
- Interface-based abstractions for extensibility
- Builder pattern for complex parameters

## Recent Decisions

### [Date]: [Decision Title]
- **Context**: [What problem was being solved]
- **Decision**: [What was decided]
- **Rationale**: [Why this approach was chosen]
- **Alternatives Considered**: [Other options that were evaluated]
- **Consequences**: [Impact and implications]

## Known Issues

### Issue 1: [Description]
- **Status**: [Open/In Progress/Resolved]
- **Severity**: [Critical/High/Medium/Low]
- **Impact**: [What's affected]
- **Workaround**: [Temporary solution if any]
- **Plan**: [How it will be resolved]

## Dependencies and Configuration

### Key Dependencies
- Go 1.x (check go.mod for exact version)
- Standard library packages
- [List any external dependencies from go.mod]

### Environment Variables
- [List any required environment variables]
- [Configuration options]

### Configuration Files
- `go.mod`: Go module definition
- `go.sum`: Dependency checksums
- `.cursor/rules/`: Agent behavior rules

## Testing Strategy

### Unit Tests
- Location: `*_test.go` files alongside source
- Convention: `Test<FunctionName>` format
- Run: `go test ./...`

### Integration Tests
- [Describe integration test approach if any]

### Test Data
- [Location of test fixtures or data]

## Build and Run

```bash
# Build
go build ./...

# Test
go test ./...

# Run examples
cd examples/basic && go run main.go
```

## Git Workflow

- Main branch: `v5`
- Feature branches: `feature/<name>`
- Commit messages: Conventional commits format preferred

## Notes

- This file serves as persistent memory for the AI agent
- Update after significant changes or decisions
- Keep focused on information relevant to development
- Remove outdated information regularly

