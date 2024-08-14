# syntax=docker/dockerfile:1

ARG GO_VERSION=1.22.5
FROM golang:${GO_VERSION} AS base
WORKDIR /src

# ==============================================
FROM base AS dev
# Install tools of system
WORKDIR /src
RUN apt update -y && \
    apt install -y && \
    apt upgrade -y && \
    apt install -y --no-install-recommends \
    protobuf-compiler \
    git

# Install tools of golang
RUN  go install github.com/spf13/cobra-cli@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && \
    go install -v github.com/go-delve/delve/cmd/dlv@latest

# Clean up
RUN rm -rf /var/lib/apt/lists/*

EXPOSE 50052
USER root

# ==============================================
FROM base AS build

RUN --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 go build  -v -o /bin/server

RUN curl -sSf https://atlasgo.sh | sh
RUN chmod +x /bin/server
RUN chmod +x /usr/local/bin/atlas
# ==============================================
FROM gcr.io/distroless/static:nonroot AS final
WORKDIR /src
ENV TZ=Asia/Tokyo
USER 65532:65532
EXPOSE 50052

# Copy the executable from the "build" stage.
COPY --from=build --chown=nonroot:nonroot  /bin/server /bin/
COPY --from=build --chown=nonroot:nonroot  /usr/local/bin/atlas /bin/
COPY --chown=nonroot:nonroot ent/migrate/migrations /src/ent/migrate/migrations
ENV PATH=/bin:$PATH

ENTRYPOINT [ "/bin/server" ]
