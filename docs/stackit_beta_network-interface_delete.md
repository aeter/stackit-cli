## stackit beta network-interface delete

Deletes a network interface

### Synopsis

Deletes a network interface.

```
stackit beta network-interface delete [flags]
```

### Examples

```
  Delete network interface with nic id "xxx" and network ID "yyy"
  $ stackit beta network-interface delete xxx --network-id yyy
```

### Options

```
  -h, --help                Help for "stackit beta network-interface delete"
      --network-id string   Network ID
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

* [stackit beta network-interface](./stackit_beta_network-interface.md)	 - Provides functionality for Network Interface
