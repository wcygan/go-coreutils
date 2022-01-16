# Go-Coreutils

A Go implementation of gnu-coreutils programs (https://www.gnu.org/software/coreutils/manual/coreutils.html)

## Build and Run

In the root directory, build the binary, invoke the program, and provide a subcommand to execute:

```
$ go build -o gco
$ ./gco <subcommand> <flags> <args>
```

For example, I can execute `tree` like so:

```
$ ./gco tree ./tree
```

This yields the following:

```
./tree
├── testdir
│   ├── a
│   │   └── apple
│   ├── b
│   │   └── banana
│   ├── bar
│   ├── baz
│   ├── c
│   │   └── orange
│   └── foo
└── tree.go
```

In another example, I can execute the `ping` subcommand with the argument `google.com` and flag `-p 80` like so:

```
$ ./gco ping google.com -p 80
connected to google.com at 142.250.190.142:80
connected to google.com at 142.250.190.142:80
connected to google.com at 142.250.190.142:80
connected to google.com at 142.250.190.142:80
connected to google.com at 142.250.190.142:80
```


## Run the tests

In the root directory you can execute `go test ./...` to run the tests:

```
$ go test ./...
?       github.com/wcygan/go-coreutils  [no test files]
?       github.com/wcygan/go-coreutils/cmd      [no test files]
?       github.com/wcygan/go-coreutils/constants        [no test files]
?       github.com/wcygan/go-coreutils/du       [no test files]
ok      github.com/wcygan/go-coreutils/echo     (cached)
ok      github.com/wcygan/go-coreutils/ls       (cached)
?       github.com/wcygan/go-coreutils/ping     [no test files]
ok      github.com/wcygan/go-coreutils/pwd      0.106s
?       github.com/wcygan/go-coreutils/tree     [no test files]
ok      github.com/wcygan/go-coreutils/whoami   (cached)

```