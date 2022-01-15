# Go-Coreutils

A Go implementation of gnu-coreutils programs (https://www.gnu.org/software/coreutils/)

## Build and Run

In the root directory, build the binary, invoke the program, and provide a subcommand to execute:

```
$ go build -o gco
$ ./gco <subcommand> <flags> <args>
```

For example, I can execute `echo` like so:

```
$ ./gco echo hello world
```

Where `echo` is the subcommand and `hello world` are the arguments


## Run the tests

In the root directory you can execute `go test ./...` to run the tests:

```
$ go test ./...
?       github.com/wcygan/go-coreutils  [no test files]
?       github.com/wcygan/go-coreutils/cmd      [no test files]
ok      github.com/wcygan/go-coreutils/echo     (cached)
```