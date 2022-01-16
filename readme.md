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
│   │   └─── apple
│   ├── b
│   │   └─── banana
│   ├── bar
│   ├── baz
│   ├── c
│   │   └─── orange
│   └─── foo
└─── tree.go
```


## Run the tests

In the root directory you can execute `go test ./...` to run the tests:

```
$ go test ./...
?   	github.com/wcygan/go-coreutils	[no test files]
?   	github.com/wcygan/go-coreutils/cmd	[no test files]
?   	github.com/wcygan/go-coreutils/constants	[no test files]
ok  	github.com/wcygan/go-coreutils/echo	(cached)
ok  	github.com/wcygan/go-coreutils/ls	(cached)
ok  	github.com/wcygan/go-coreutils/pwd	0.088s
?   	github.com/wcygan/go-coreutils/tree	[no test files]
ok  	github.com/wcygan/go-coreutils/whoami	(cached)
```