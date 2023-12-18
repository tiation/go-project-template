# Your new golang cli application (Template golang project)

A basic template to get started with creating a command-based cli application in golang. This should be replaced with something relevant.

## Setup

> Delete or replace Setup section with something relevant for your application after you set up the project.

1. Clone this repo
2. Replace : `find -type f -exec sed -i -E -e 's/GIT_OWNER/iamowner/g' -e 's/GIT_PROJECT/myproject/g' {} \;`

## Run It 🏃

`go run main.go`

## Usage

- `project-name`

    ```properties
    Short description of the project

    All flags values can be provided via env vars starting with MYPREFIX_*
    To pass a command (e.g. 'command1') flag, use MYPREFIX_COMMAND1_FLAGNAME=somevalue

    Usage:
    project-name [flags]
    project-name [command]

    Available Commands:
    completion     Generate the autocompletion script for the specified shell
    help           Help about any command
    sample-command This is a sample command
    version        Display version and exit

    Flags:
        --config strings        Config file(s) or directories. When just dirs, file 'main' with extensions 'json, toml, yaml, yml, properties, props, prop, hcl, tfvars, dotenv, env, ini' is looked up. Can be specified multiple times (default [.,$HOME/main])
    -h, --help                  help for project-name
        --log-format string     Set log format to one of: 'console, json' (default "console")
        --log-level string      Set log level to one of: 'trace, debug, info, warn, error, fatal, panic, disabled' (default "warn")
        --project-root string   Project root directory (default "/some/current/directory")

    Use "project-name [command] --help" for more information about a command.
    ```

- `project-name sample-command -h`

    ```properties
    This is a sample command

    Usage:
    project-name sample-command [flags]

    Aliases:
    sample-command, c

    Flags:
    -c, --flag1 string       [Required] A flag called Flag1 (default "<undefined>")
    -h, --help               help for sample-command
    -p, --some-path string   Some path flag (default "./samplecommand-<undefined>")
    -t, --timeout duration   Some timeout flag (default 10m0s)

    Global Flags:
        --config strings        Config file(s) or directories. When just dirs, file 'main' with extensions 'json, toml, yaml, yml, properties, props, prop, hcl, tfvars, dotenv, env, ini' is looked up. Can be specified multiple times (default [.,$HOME/main])
        --log-format string     Set log format to one of: 'console, json' (default "console")
        --log-level string      Set log level to one of: 'trace, debug, info, warn, error, fatal, panic, disabled' (default "warn")
        --project-root string   Project root directory (default "/some/current/directory")
    ```

## Configure It ☑️

- See [sample/myconfig.yaml](./sample/myconfig.yaml) for config file
- All parameters can be set via flags or env as well: `MYPREFIX_<subcommand>_<flag>`, example: `MYPREFIX_SAMPLE_COMMAND_FLAG1=1122334455`

## Test It 🧪

Test for coverage and race conditions

`make coverage`

## Lint It 👕

`make pre-commit run --all-files --show-diff-on-failure`

## Roadmap

- [ ] ?

## Development

### Build

- Preferably: `make build` or
- `goreleaser build --snapshot --clean --single-target` or
- `scripts/dev-build.sh` (deprecated)
