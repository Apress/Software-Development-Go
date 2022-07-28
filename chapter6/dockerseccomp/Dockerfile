FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN GOOS=linux GOARCH=amd64 go build -o /dockerseccomp

# package into distroless
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /dockerseccomp /dockerseccomp

USER nonroot:nonroot

ENTRYPOINT ["/dockerseccomp"]