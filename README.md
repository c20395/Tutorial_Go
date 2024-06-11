$ go build main.go
$ ./main

Or, alternatively

$ go run main.go

To format Go source code, use:

$ gofmt -w main.go

Time https://golang.org/pkg/time/

# goimport main.go  // Removes unnecessary packages
# gofmt main.go // formats it
# go build main.go // Compile

Lints code

# go vet main.go

Detects race conditions

# go run --race p.go

