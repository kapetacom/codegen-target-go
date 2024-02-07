##FILENAME:Dockerfile:create-only
# syntax=docker/dockerfile:1

# Create a stage for building the application.
ARG GO_VERSION=1.21.5
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
WORKDIR /src

# Download dependencies as a separate step to take advantage of Docker's caching.
# Leverage a cache mount to /go/pkg/mod/ to speed up subsequent builds.
# Leverage bind mounts to go.sum and go.mod to avoid having to copy them into
# the container.
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

RUN --mount=type=bind,target=. \
    cp kapeta.yml /bin/

# Placing it here allows the previous steps to be cached across architectures.
ARG TARGETARCH

# Build the application.
# Leverage a cache mount to /go/pkg/mod/ to speed up subsequent builds.
# Leverage a bind mount to the current directory to avoid having to copy the
# source code into the container.
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/server .

FROM scratch AS final

# Copy the executable from the "build" stage.
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Add timezone data
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /bin/server /bin/
COPY --from=build /bin/kapeta.yml /

# Expose the port that the application listens on.
EXPOSE 80

# What the container should run when it is started.
ENTRYPOINT [ "/bin/server" ]
