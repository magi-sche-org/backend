# build for deploy
FROM golang:1.20-bullseye as deploy-builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -trimpath -ldflags="-s -w" -o app
RUN go build -trimpath -ldflags="-s -w" -o migrate ./cmd/migrate
RUN go build -trimpath -ldflags="-s -w" -o healthcheck ./cmd/healthcheck

# for deploy
FROM gcr.io/distroless/base-debian11:latest as deploy

WORKDIR /app
COPY --from=deploy-builder /app/app /app
COPY --from=deploy-builder /app/healthcheck /app

# for migration
FROM golang:1.20-bullseye as migrate

WORKDIR /app
COPY --from=deploy-builder /app/migrate /app
COPY --from=deploy-builder /app/db/migrations /app/db/migrations
CMD ["./migrate"]


# for local development with air
FROM golang:1.20-bullseye as dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest && \
    go install github.com/volatiletech/sqlboiler/v4@latest && \
    go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest && \
    go install github.com/pressly/goose/v3/cmd/goose@latest && \
    go install github.com/golang/mock/mockgen@latest
CMD ["air"]
