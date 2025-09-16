# Build the Go Binary.
FROM golang:1.25 AS build_api
ENV CGO_ENABLED=0
ARG BUILD_REF

# Copy the source code into the container
COPY . /service

# Build the api binary
WORKDIR /service/cmd/api
RUN go build -ldflags "-X main.build=${BUILD_REF}"

FROM alpine:3.22
ARG BUILD_REF
ARG BUILD_DATE
RUN addgroup -g 1000 -S api && \
    adduser -u 1000 -h /service -G api -S api
COPY --from=build_api --chown=api:api /service/cmd/api/api /service/api
WORKDIR /service
USER api
# TODO: Add healthcheck to api binary
CMD ["./api"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
    org.opencontainers.image.title="api" \
    org.opencontainers.image.authors="Stephen Morgan <hello@vexaralabs.com>" \
    org.opencontainers.image.source="https://github.com/rebelopsio/hexagonal-project/tree/main/cmd/api" \
    org.opencontainers.image.revision="${BUILD_REF}" \
    org.opencontainers.image.vendor="Vexara Labs"
