# D-Link DWR-932C (LTE Router) Unofficial CLI tool

Created for personal purposes

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

Available subcommands:
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