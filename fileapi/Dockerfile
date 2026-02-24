FROM golang:1.25

WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o fileapi .

FROM scratch

COPY --from=0 /src/fileapi /fileapi
EXPOSE 8080
ENTRYPOINT ["/fileapi"]