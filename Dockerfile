# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.14

# Copy the local package files to the container's workspace.
COPY . /go/src/github.com/bbarrington42/web-test

# Build the command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/bbarrington42/web-test

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/web-test

# The port the service listens on.
EXPOSE 8080
