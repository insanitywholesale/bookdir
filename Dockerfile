# build stage
FROM golang:1.17 as build

ENV CGO_ENABLED 0
ENV GO111MODULE on

WORKDIR /go/src/bookdir
COPY . .

RUN go get -v
RUN go vet -v
RUN go install -v

# run stage
FROM busybox as run
COPY --from=build /go/bin/bookdir /bookdir
EXPOSE 11000
EXPOSE 8080
CMD ["/bookdir"]
