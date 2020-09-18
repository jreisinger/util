FROM golang:1.13-alpine AS build

ARG GOOS=darwin
ARG GOARCH=amd64

# Set the current working directory inside container.
WORKDIR /go/src/util

# Install tools required for building the app.
RUN apk add git

# Download all dependencies.
COPY go.* ./
RUN go mod download

# Build the app.
COPY . ./
RUN go build -o /bin/util

# Create a single layer image.
#FROM scratch # -> this doesn't work
FROM alpine:latest
WORKDIR /app/util
COPY --from=build /bin/util /app/util/util
COPY --from=build /go/src/util/template /app/util/template
COPY --from=build /go/src/util/static /app/util/static
RUN apk update
RUN apk add git

EXPOSE 5002
ENTRYPOINT ["/app/util/util"]