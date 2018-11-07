# Protocol Buffer
## Requirement
* Golang version 1.9 or higher
* install protoc `brew install protobuf` check in [here](http://google.github.io/proto-lens/installing-protoc.html)
* `go get -u github.com/golang/protobuf/protoc-gen-go`. The compiler plugin `protoc-gen-go` will be installed in `$GOBIN`, defaulting to `$GOPATH/bin`. It must be in your `$PATH` for the protocol compiler `protoc` to find it.

## How to run
1. Create your directory to this project
2. make `.proto` file
3. Generate using `protoc` in terminal. Usually after generated you get file `name.pb.go` contain default of struct and function that can be use in your application
4. write some code in your `main.go` like
5. Run your application
