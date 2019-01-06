FROM golang:1.10-alpine3.8 as builder

LABEL author_name="Alvin Rizki"
LABEL author_emails="cikupin@gmail.com, 4lvin.rizki@gmail.com"

RUN apk update
RUN apk add --no-cache git tzdata

WORKDIR /go/src/github.com/cikupin/filtering_input
COPY . /go/src/github.com/cikupin/filtering_input

# install go dep
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

# build executable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' .

# multi-stage building
FROM scratch

ENV BUILDDIR /go/src/github.com/cikupin/filtering_input

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder ${BUILDDIR}/filtering_input /go/bin/filtering_input

EXPOSE 8080
ENTRYPOINT ["/go/bin/filtering_input"]
