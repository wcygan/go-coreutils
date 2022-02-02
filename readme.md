# Go-Coreutils

A Go implementation of gnu-coreutils programs (https://www.gnu.org/software/coreutils/manual/coreutils.html)

## Installation via [Go](https://go.dev/dl/)

Install on any platform using `go get`:

```
$ go get github.com/wcygan/go-coreutils
```

## How to run

Once the binary is installed, use the `go-coreutils` program like so:

```
$ go-coreutils
A Go implementation of gnu-coreutils programs - https://www.gnu.org/software/coreutils/

Usage:
  go-coreutils [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  du          Disk Usage
  echo        Print a line of text
  help        Help about any command
  ls          List directory contents
  ping        Tests the reachability of a host on a network
  pwd         Print working directory
  tree        Display a directory tree
  whoami      Print effective user name
  yes         Repeats the provided text until interrupted

Flags:
  -h, --help   help for go-coreutils

Use "go-coreutils [command] --help" for more information about a command.
```

You can execute a subcommand like `tree` in the following way:

```
$ go-coreutils tree
.
├── LICENSE
├── cmd
│   ├── common.go
│   ├── du.go
│   ├── echo.go
│   ├── ls.go
│   ├── ping.go
│   ├── pwd.go
│   ├── root.go
│   ├── tree.go
│   ├── whoami.go
│   └── yes.go
├── constants
│   └── shared_constants.go
├── du
│   └── du.go
├── echo
│   ├── echo.go
│   └── echo_test.go
├── gco
├── go.mod
├── go.sum
├── ls
│   ├── ls.go
│   └── ls_test.go
├── main.go
├── ping
│   └── ping.go
├── pwd
│   ├── pwd.go
│   └── pwd_test.go
├── readme.md
├── tree
│   ├── testdir
│   │   ├── a
│   │   │   └── apple
│   │   ├── b
│   │   │   └── banana
│   │   ├── bar
│   │   ├── baz
│   │   ├── c
│   │   │   └── orange
│   │   └── foo
│   └── tree.go
├── whoami
│   ├── whoami.go
│   └── whoami_test.go
└── yes
    └── yes.go

```

## Add a command using [Cobra](https://cobra.dev/)

Make sure you have the Cobra generator ([Install](https://cobra.dev/#install), [Docs](https://github.com/spf13/cobra/blob/master/cobra/README.md)) installed.

With the Cobra generator you can initialize the Cobra framework, add commands & subcommands, and more.

```
$ cobra init
$ cobra add <command>
$ cobra add <command> <subcommand>
```