FROM golang:latest AS builder
WORKDIR /app
COPY . .

# ! Use the -o flag to specify the output binary name
# RUN go build -o custom_binary_name . # Build the Go application with a custom binary name

# ! Use the GOOS environment variable to specify the target operating system
# RUN GOOS=linux go build . # Build the Go application for Linux
# RUN GOOS=windows go build . # Build the Go application for Windows
# RUN GOOS=darwin go build . # Build the Go application for MacOS

# ! Use the GOARCH environment variable to specify the target architecture
# RUN GOARCH=amd64 go build . # Build the Go application for AMD64 architecture
# RUN GOARCH=arm64 go build . # Build the Go application for ARM64 architecture

# ! Use ldflags to remove the debug and symbol information from the binary
# ! -s removes the symbol table and debug information, -w removes the DWARF symbol table
# RUN go build -ldflags="-s -w" . # Build the Go application with stripped debug information

# ! Use the CGO_ENABLED environment variable to disable cgo
# ! CGO_ENABLED=0 disables cgo, which is useful for building statically linked binaries
# RUN CGO_ENABLED=0 go build . # Build the Go application with cgo disabled

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o server .

FROM scratch
COPY --from=builder /app/server .
# ! Use the EXPOSE instruction to document the port on which the container will listen for connections
EXPOSE 8080
# ! Use the CMD instruction to specify the command to run when the container starts
CMD ["./server"]
