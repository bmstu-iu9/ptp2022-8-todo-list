# syntax=docker/dockerfile:1

FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /minimal-login-page

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /minimal-login-page /minimal-login-page

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/minimal-login-page"]
