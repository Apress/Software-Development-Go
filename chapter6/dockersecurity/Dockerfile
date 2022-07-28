FROM golang:1.18 as build
COPY ./main.go .
COPY ./go.mod .
ENV GOPATH=""
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -trimpath -v -a -o sample -ldflags="-w -s"
RUN useradd -u 12345 normaluser

FROM scratch
COPY --from=build /go/sample /sample
COPY --from=build /etc/passwd /etc/passwd
USER normaluser
ENTRYPOINT ["/sample"]