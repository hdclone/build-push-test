FROM golang:1.17.0 AS builder
ARG VERSION="v0.0-0-g0000000"

WORKDIR /builder
COPY ./ /builder

RUN go build -ldflags "-X 'broadcaster/internal/variables.Version=$VERSION'" -o ./dist/bin/broadcasterd ./cmd/broadcasterd/main.go
RUN go build -ldflags "-X 'broadcaster/internal/variables.Version=$VERSION'" -o ./dist/bin/broadcasterctl ./cmd/broadcasterctl/main.go
RUN go build -ldflags "-X 'broadcaster/internal/variables.Version=$VERSION'" -o ./dist/bin/migrator ./cmd/migrator/main.go
RUN ldd ./dist/bin/broadcasterd | grep libwasmvm.so | awk 'NF == 4 { system("cp " $3 " ./dist/") }'

FROM alpine:3.14.0
ARG CONFIG=dev

WORKDIR /app
RUN apk update && apk add libgcc gcompat

COPY --from=builder /builder/dist /app/
COPY --from=builder /builder/dist/libwasmvm.so /usr/lib/
CMD /app/bin/migrator up && /app/bin/broadcasterd
