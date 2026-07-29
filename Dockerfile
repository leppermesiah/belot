FROM golang:1.22-bookworm AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/belot-server ./cmd/server

FROM alpine:3.20
RUN adduser -D -H belot
WORKDIR /app
COPY --from=build /out/belot-server ./belot-server
COPY web ./web
USER belot
EXPOSE 8080
ENTRYPOINT ["./belot-server", "-addr", ":8080", "-web", "./web"]
