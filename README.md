Set the following command to set GOBIN
- $ go env -w GOBIN={path to the project}/bin

If you need to unset a variable previously set by go env -w, use go env -u: 
- $ go env -u GOBIN

- Goto {path to the project}/morestrings and run
go build reverse.go 

- Goto {path to the project} and run
go install abhinayak.com/user/hello

- To run Goto {path to the project}/bin and run
./hello

- To run test Goto {path to the project}/gotest/{folder with test}
go test