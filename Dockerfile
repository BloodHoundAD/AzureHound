# syntax=docker/dockerfile:1

FROM golang:1.20 as build
WORKDIR /app

ARG VERSION=v0.0.0
ENV CGO_ENABLED=1

COPY ./ ./
RUN go mod download
RUN go build -ldflags="-s -w -X github.com/bloodhoundad/azurehound/v2/constants.Version=$VERSION+docker"

FROM gcr.io/distroless/base-debian11
LABEL org.opencontainers.image.source https://github.com/BloodHoundAD/AzureHound
COPY --from=build /app/azurehound /
ENTRYPOINT ["/azurehound"]
