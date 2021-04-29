FROM golang:1.15.6 as builder

ENV GO111MODULE=on

# All these steps will be cached
RUN mkdir -p /go/src/klauskie.com/microservice-aurant/session-api

WORKDIR /go/src/klauskie.com/microservice-aurant/session-service
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

## OPTION 1

# Build the binary
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app
# FROM scratch 
# COPY --from=builder /go/bin/app /go/bin/app
# ENTRYPOINT ["/go/bin/app"]


## OPTION 2

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app
# RUN chmod +x /go/bin/app
# ENTRYPOINT ["/go/bin/app"]

## OPTION 3 : dev

CMD ["go","run","session-main.go"]