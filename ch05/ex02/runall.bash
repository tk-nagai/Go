go build fetch.go
go build elements.go
./fetch http://golang.org | ./elements
