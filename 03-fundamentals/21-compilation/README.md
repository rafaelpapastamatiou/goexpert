####Compiles to current platform:

go build main.go
./main


####Compiles to a specified platform:

GOOS=windows go build main.go
./main.exe

GOOS=linux go build main.go
./main

GOOS=darwin go build main.go
./main


####Compiles to specificed platform/architecture

windows/amd64
GOOS=windows GOARC=amd64 go build main.go
./main.exe

windows/arm64
GOOS=windows GOARC=arm64 go build main.go
./main.exe

linux/amd64
GOOS=linux GOARC=amd64 go build main.go
./main.exe

linux/arm64
GOOS=linux GOARC=arm64 go build main.go
./main.exe

darwin/amd64
GOOS=darwin GOARC=amd64 go build main.go
./main.exe

darwin/arm64
GOOS=darwin GOARC=arm64 go build main.go
./main.exe

https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures

####Compiles with custom name

go build main.go -o my-program
./my-program