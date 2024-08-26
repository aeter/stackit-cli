## stackit beta server command template list

Lists all server command templates

### Synopsis

Lists all server command templates.

```
stackit beta server command template list [flags]
```

### Examples

```
  List all command templates
  $ stackit beta server command template list

  List all commands templates in JSON format
  $ stackit beta server command template list --output-format json
```

### Options

```
  -h, --help        Help for "stackit beta server command template list"
      --limit int   Maximum number of entries to list
```

### Options inherited from parent commands

```
  -y, --assume-yes             If set, skips all confirmation prompts
      --async                  If set, runs the command asynchronously
  -o, --output-format string   Output format, one of ["json" "pretty" "none" "yaml"]
  -p, --project-id string      Project ID
      --verbosity string       Verbosity of the CLI, one of ["debug" "info" "warning" "error"] (default "info")
```

### SEE ALSO

* [stackit beta server command template](./stackit_beta_server_command_template.md)	 - Provides functionality for Server Command Template
