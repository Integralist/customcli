The code in this repo was based off the design of [fastly/cli](https://github.com/fastly/cli).

It's incomplete, but the motivation for its existence was to validate if a pure standard library implementation, instead of a third-party dependency, would be feasible (it _is_).

The value to this work is that there is no need to workaround restrictions inherent in the design of these various third-party tools. This specifically has been an issue with fastly/cli use of the Kingpin dependency (now no longer maintained) and discussions regarding moving to a simpler tool such as https://github.com/peterbourgon/ff/tree/master/ffcli has yielded similar concerns regarding fastly/cli having more specific requirements than a third-party abstraction should be expected to handle.

It would appear that the fastly/cli would benefit from a home grown solution to allow it to simplify its design while enabling room for easier extension.

## Help Output

The current implementation of the fastly CLI isn't very flexible in its use of a help flag. For example, it supports `--help` but not `-help` or `-h`.

Here is the various output that can be expected from this program...

### `fastly`

Type the program with no arguments, and you'll get help output:

```
USAGE
  fastly [<flags>] <command> [<args> ...]

A tool to interact with the Fastly API

GLOBAL FLAGS
  -h, -help     Show context-sensitive help
  -t, -token    Fastly API token (or via FASTLY_API_TOKEN)
  -v, -verbose  Verbose logging

COMMANDS
  backend          Manipulate Fastly service version backends
  compute          Manage Compute@Edge packages
  configure        Configure the Fastly CLI
  dictionary       Manipulate Fastly edge dictionaries
  dictionaryitem   Manipulate Fastly edge dictionary items
  domain           Manipulate Fastly service version domains
  healthcheck      Manipulate Fastly service version healthchecks
  help             Show help.
  logging          Manipulate Fastly service version logging endpoints
  service          Manipulate Fastly services
  service-version  Manipulate Fastly service versions
  stats            View statistics (historical and realtime) for a Fastly service
  update           Update the CLI to the latest version
  version          Display version information for the Fastly CLI
  whoami           Get information about the currently authenticated account

ERROR: error parsing arguments: command not specified.
```

> NOTE: the `ERROR` message at the bottom of the help output, which indicates the reason the help output was shown was because no arguments were provided.

### `fastly --help`

Type the program with a single flag `--help` (or `-help`, `h`), and you'll get help output:

```
USAGE
  fastly [<flags>] <command> [<args> ...]

A tool to interact with the Fastly API

GLOBAL FLAGS
  -h, -help     Show context-sensitive help
  -t, -token    Fastly API token (or via FASTLY_API_TOKEN)
  -v, -verbose  Verbose logging

COMMANDS
  backend          Manipulate Fastly service version backends
  compute          Manage Compute@Edge packages
  configure        Configure the Fastly CLI
  dictionary       Manipulate Fastly edge dictionaries
  dictionaryitem   Manipulate Fastly edge dictionary items
  domain           Manipulate Fastly service version domains
  healthcheck      Manipulate Fastly service version healthchecks
  help             Show help.
  logging          Manipulate Fastly service version logging endpoints
  service          Manipulate Fastly services
  service-version  Manipulate Fastly service versions
  stats            View statistics (historical and realtime) for a Fastly service
  update           Update the CLI to the latest version
  version          Display version information for the Fastly CLI
  whoami           Get information about the currently authenticated account
```

> NOTE: there is no `ERROR` message at the bottom of the output anymore as you provided the relevant flag to purposefully trigger help output.

### `fastly <command>`

Type the program, followed by a recognized 'command' (but a command that itself requires a subcommand to be provided), and you'll get contextual help output for the specific command:

```
$ fastly backend

USAGE
  fastly backend <command> [<args> ...]

Manipulate Fastly service version backends

GLOBAL FLAGS
  -h, -help     Show context-sensitive help
  -t, -token    Fastly API token (or via FASTLY_API_TOKEN)
  -v, -verbose  Verbose logging

COMMANDS
  create    Create a backend on a Fastly service version
  delete    Delete a backend on a Fastly service version
  describe  Show detailed information about a backend on a Fastly service version
  list      List backends on a Fastly service version
  update    Update a backend on a Fastly service version
```

### `fastly <command> --help`

Type the program, followed by a recognized 'command', with a single flag `--help` (or `-help`, `h`), and you'll get contextual help output for the specific command:

```
$ fastly backend --help

USAGE
  fastly backend <command> [<args> ...]

Manipulate Fastly service version backends

GLOBAL FLAGS
  -h, -help     Show context-sensitive help
  -t, -token    Fastly API token (or via FASTLY_API_TOKEN)
  -v, -verbose  Verbose logging

COMMANDS
  create    Create a backend on a Fastly service version
  delete    Delete a backend on a Fastly service version
  describe  Show detailed information about a backend on a Fastly service version
  list      List backends on a Fastly service version
  update    Update a backend on a Fastly service version
```

### `fastly <command> --help <command>`

Type the program, followed by a recognized 'command', with a single flag `--help` (or `-help`, `h`), followed by another recognized 'subcommand' and you'll get contextual help output for the _parent_ command:

```
$ fastly backend --help create

USAGE
  fastly backend <command> [<args> ...]

Manipulate Fastly service version backends

GLOBAL FLAGS
  -h, -help     Show context-sensitive help
  -t, -token    Fastly API token (or via FASTLY_API_TOKEN)
  -v, -verbose  Verbose logging

COMMANDS
  create    Create a backend on a Fastly service version
  delete    Delete a backend on a Fastly service version
  describe  Show detailed information about a backend on a Fastly service version
  list      List backends on a Fastly service version
  update    Update a backend on a Fastly service version


ERROR: error parsing arguments: expected command but got --help.
```

### `fastly <command> <subcommand> --help`

Type the program, followed by a recognized 'command' and a recognized 'subcommand', with a single flag `--help` (or `-help`, `h`), and you'll get contextual help output for the subcommand:

```
$ fastly backend create --help

USAGE
  fastly backend create --address=ADDRESS --comment=COMMENT --name=NAME --service-id=SERVICE-ID --version=VERSION

Create a backend on a Fastly service version

GLOBAL FLAGS
  -h, -help     Show context-sensitive help
  -t, -token    Fastly API token (or via FASTLY_API_TOKEN)
  -v, -verbose  Verbose logging

COMMAND FLAGS
      -address     A hostname, IPv4, or IPv6 address for the backend
      -comment     A descriptive note
  -h, -help        Show context-sensitive help
  -n, -name        Backend name
  -s, -service-id  Service ID
  -s, -version     Number of service version

```
