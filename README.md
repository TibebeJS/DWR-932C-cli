# D-Link DWR-932C (LTE Router) Unofficial CLI tool

>Note: Created for personal purposes

## Setup
create a `.env` file with the following keys set. (You can copy & paste the content of `sample.env` and make any necessary changes)
```
DWR_URL="http://192.168.0.1"
DWR_USERNAME="admin"
DWR_PASSWORD="admin"
```

## Available commands
main command: `DWR-932C-cli`

```
NAME:
   DWR-932C CLI tool - Helps manage/control the D-Link DWR-932C LTE Router right from the command line

USAGE:
    [global options] command [command options] [arguments...]

COMMANDS:
   change-mode, m  change connection mode to.
   help, h         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help

```
### subcommands:
 - `DWR-932C-cli change-mode`

    ```
    NAME:
        change-mode - change connection mode to.

    USAGE:
        change-mode [command options] [arguments...]

    OPTIONS:
    --to value  desired connection mode ('auto', 'LTE', '3G' or '2G'). (default: "auto")
    ```

Built in GO!