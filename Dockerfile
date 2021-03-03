FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/alexunjm/golang_persons-api
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/golang_persons-api main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/golang_persons-api /go/bin/golang_persons-api
ENTRYPOINT ["/go/bin/golang_persons-api"]
