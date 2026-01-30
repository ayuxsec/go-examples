```go
package main

import (
    "github.com/alecthomas/kong"
)

// Pattern: Define struct with tags, call kong.Parse()

// Example 1: Basic flags
type CLI struct {
    Debug bool   `help:"Enable debug mode"`
    Port  int    `default:"8080" help:"Server port"`
    Host  string `short:"h" default:"localhost"`
}

func main() {
    var cli CLI
    kong.Parse(&cli)
    // Use cli.Debug, cli.Port, cli.Host
}

// Example 2: Subcommands
type Context struct {
    Globals struct {
        Config string `type:"path" default:"config.yaml"`
    }

    Run struct {
        Name string `arg:"" required:"" help:"Service name"`
        Env  string `short:"e" default:"dev"`
    } `cmd:"" help:"Run service"`

    Stop struct {
        Force bool `short:"f"`
    } `cmd:"" help:"Stop service"`
}

func main() {
    ctx := kong.Parse(&Context{})
    switch ctx.Command() {
    case "run <name>":
        // ctx.Run.Name, ctx.Run.Env
    case "stop":
        // ctx.Stop.Force
    }
}

// Example 3: Required args, slices, maps
type Advanced struct {
    Files  []string          `arg:"" required:"" help:"Input files"`
    Tags   []string          `short:"t" help:"Tags to apply"`
    Env    map[string]string `short:"e" help:"Environment vars"`
    Retries int              `default:"3" enum:"1,3,5,10"`
}
```
