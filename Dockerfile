FROM golang as build-go
WORKDIR /your-work-dir
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/your-work-dir .

FROM alpine:latest
RUN addgroup -S alan-algo-api && adduser -S alan-algo-api -G alan-algo-api
USER alan-algo-api
WORKDIR /home/alan-algo-api
COPY --from=build-go /bin/alan-algo-api ./
EXPOSE 8080
ENTRYPOINT ["./alan-algo-api"]