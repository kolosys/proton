# cli API

Complete API documentation for the cli package.

**Import Path:** `github.com/kolosys/proton/internal/cli`

## Package Documentation



## Functions

### Execute
Execute adds all child commands to the root command and sets flags appropriately. It initializes the CLI application and runs the root command. Returns error if command execution fails.

```go
func Execute() error
```

**Parameters:**
None

**Returns:**
| Type | Description |
|------|-------------|
| `error` | |

**Example:**

```go
// Example usage of Execute
result := Execute(/* parameters */)
```

## External Links

- [Package Overview](../packages/cli.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/proton/internal/cli)
- [Source Code](https://github.com/kolosys/proton/tree/main/internal/cli)
