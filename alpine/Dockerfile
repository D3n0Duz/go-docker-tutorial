############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk add ca-certificates && apk update && apk add --no-cache git
RUN mkdir -p ${GOPATH}/go-rest-api;
RUN go get firebase.google.com/go; \
    go get google.golang.org/api/option; \
    go get github.com/go-chi/chi; \
    go get github.com/rs/xid; \
    go get google.golang.org/api/iterator; \
    go get cloud.google.com/go/firestore; \
    go get github.com/fatih/structs; \
    go get github.com/dgrijalva/jwt-go;

COPY . ${GOPATH}/go-rest-api
WORKDIR $GOPATH/go-rest-api

# Using go get.
# RUN go get -d -v ./...

# Build the binary.
RUN go build -o app .

############################
# STEP 2 build a small image
############################
FROM alpine:latest

# Add ca certificates
RUN apk add ca-certificates

# Copy our static executable.
COPY --from=builder /go/go-rest-api .

# Run the hello binary.
ENTRYPOINT ["/app"]