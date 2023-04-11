# #Dockerfile asynq worker server
# FROM golang:1.19 as build

# WORKDIR /build
# COPY . .
# RUN CGO_ENABLED=0 go build -o worker worker/server.go

# FROM alpine:3.12

# WORKDIR /app
# EXPOSE 9090
# COPY --from=build /build/worker .

# CMD ["./server"]

#Dockerfile asynq client
FROM golang:1.19 as build

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -o client client/main.go

FROM alpine:3.12

WORKDIR /app
EXPOSE 9090
COPY --from=build /build/client .

CMD ["./main"]