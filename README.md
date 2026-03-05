# sdcd

[![CI](https://github.com/suzutan/sdcd-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/suzutan/sdcd-cli/actions/workflows/ci.yml)
[![Release](https://github.com/suzutan/sdcd-cli/actions/workflows/release.yml/badge.svg)](https://github.com/suzutan/sdcd-cli/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzutan/sdcd-cli)](https://goreportcard.com/report/github.com/suzutan/sdcd-cli)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A command-line interface for [Screwdriver.cd](https://screwdriver.cd).

Manage multiple Screwdriver.cd instances (production, staging, etc.) from a single CLI with a kubectl-style multi-context configuration.

## Features

- **Multi-context support** — switch between multiple Screwdriver.cd instances seamlessly
- **Full resource coverage** — pipelines, jobs, builds, events, and secrets
- **Multiple output formats** — table (default), JSON, and YAML
- **Colored status output** — RUNNING/SUCCESS/FAILURE/ABORTED highlighted automatically
- **Secure secret input** — value prompt with terminal masking when `--value` is omitted
- **Shell completion** — bash, zsh, fish, and PowerShell

## Installation

### Homebrew (macOS / Linux)

```sh
brew install suzutan/tap/sdcd
```

### Go install

```sh
go install github.com/suzutan/sdcd-cli@latest
```

### Download binary

Download the latest binary from the [Releases](https://github.com/suzutan/sdcd-cli/releases) page.

### Build from source

```sh
git clone https://github.com/suzutan/sdcd-cli.git
cd sdcd-cli
make build          # outputs to bin/sdcd
make install        # installs to $GOPATH/bin
```

## Configuration

The config file is stored at `$XDG_CONFIG_HOME/sdcd-cli/config.yaml` (defaults to `~/.config/sdcd-cli/config.yaml`) with permission `0600`.

```yaml
current-context: production

contexts:
  - name: production
    api-url: https://api.screwdriver.example.com
    token: xxxxxxx        # API token issued from Screwdriver UI
  - name: staging
    api-url: https://api.staging.screwdriver.example.com
    token: yyyyyyy

preferences:
  output: table           # table | json | yaml
  no-color: false
  page-size: 50
```

> **Note:** The raw API token is exchanged for a short-lived JWT at runtime via `GET /v4/auth/token`. The JWT is kept in memory only and never written to disk.

## Quick Start

```sh
# Add a context
sdcd context add production \
  --api-url https://api.screwdriver.example.com \
  --token <your-api-token>

# Set it as default
sdcd context use production

# List pipelines
sdcd pipeline list

# Start a pipeline
sdcd pipeline start 123 --job main

# Stream build logs
sdcd build logs 456 --step install
```

## Usage

### Global Flags

| Flag | Description |
|------|-------------|
| `--context <name>` | Override the active context for this invocation |
| `--output <format>` | Output format: `table` \| `json` \| `yaml` |
| `--no-color` | Disable ANSI color output |
| `--config <path>` | Path to config file |

### `sdcd context`

```
sdcd context add <name> --api-url <url> --token <token>
sdcd context remove <name>
sdcd context list
sdcd context use <name>
sdcd context current
```

### `sdcd pipeline`

```
sdcd pipeline list [--search <str>] [--page N] [--count N]
sdcd pipeline view <id>
sdcd pipeline create --checkout-url <url> [--root-dir <dir>]
sdcd pipeline delete <id> [--yes]
sdcd pipeline sync <id>
sdcd pipeline jobs <id>
sdcd pipeline events <id>
sdcd pipeline builds <id>
sdcd pipeline start <id> [--job <name>] [--sha <sha>]
```

### `sdcd job`

```
sdcd job view <id>
sdcd job enable <id>
sdcd job disable <id>
sdcd job builds <id>
sdcd job latest-build <id>
```

### `sdcd build`

```
sdcd build view <id>
sdcd build stop <id>
sdcd build logs <id> --step <name>
sdcd build steps <id>
sdcd build artifacts <id>
```

### `sdcd event`

```
sdcd event view <id>
sdcd event builds <id>
sdcd event stop <id>
sdcd event rerun <id> [--job <name>]
```

### `sdcd secret`

```
sdcd secret list --pipeline-id <id>
sdcd secret create --pipeline-id <id> --name <name> [--value <value>]
sdcd secret update <id> [--value <value>] [--allow-in-pr]
sdcd secret delete <id> [--yes]
```

### Shell Completion

```sh
# bash
echo 'eval "$(sdcd completion bash)"' >> ~/.bashrc

# zsh
echo 'eval "$(sdcd completion zsh)"' >> ~/.zshrc

# fish
sdcd completion fish | source
```

## Resource Hierarchy

```
Pipeline  (screwdriver.yaml definition)
  ├─ Job    (named stage, e.g. "main", "deploy")
  │    └─ Build  (each execution of that job)
  └─ Event  (one run of the whole pipeline)
       └─ Build  (one per job triggered in the event)
                  └─ Step  (individual shell commands)
```

### Typical workflow

```sh
# 1. Find a pipeline
sdcd pipeline list

# 2. Start it
sdcd pipeline start 123 --job main

# 3. Check resulting event and its builds
sdcd pipeline events 123
sdcd event builds 456

# 4. Inspect a build and stream its logs
sdcd build view 789
sdcd build logs 789 --step install
```

### How to get resource IDs

| Resource | Command |
|----------|---------|
| Pipeline ID | `sdcd pipeline list` |
| Job ID | `sdcd pipeline jobs <pipeline-id>` |
| Event ID | `sdcd pipeline events <pipeline-id>` |
| Build ID (by event) | `sdcd event builds <event-id>` |
| Build ID (by job) | `sdcd job builds <job-id>` |

## Examples

```sh
# Switch context for a single command
sdcd --context staging pipeline list

# Output as JSON and pipe to jq
sdcd pipeline list --output json | jq '.[].name'

# Disable a job
sdcd job disable 789
sdcd job view 789   # verify state is DISABLED

# Rerun an event from a specific job
sdcd event rerun 1234 --job deploy

# Create a secret (value prompted securely)
sdcd secret create --pipeline-id 123 --name DATABASE_URL
```

## Development

```sh
# Run tests
make test

# Build with version info
make build

# Lint (requires golangci-lint)
make lint
```

## Contributing

Contributions are welcome! Please open an issue or pull request.

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/my-feature`)
3. Commit your changes
4. Push to the branch and open a Pull Request

## License

[MIT License](LICENSE)
