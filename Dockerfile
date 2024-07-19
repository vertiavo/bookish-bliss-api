FROM alpine:latest AS migrator

WORKDIR /db/migrations

COPY db/migrations ./

FROM scratch

COPY bookish-bliss-api ./

CMD ["./bookish-bliss-api"]
