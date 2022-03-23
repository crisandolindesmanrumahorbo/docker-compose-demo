#STEP 1 build executeable library
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build -o go-gin-postgres

#STEP 2 build a small image
FROM alpine
WORKDIR /app
COPY --from=builder /app/go-gin-postgres /app
CMD ["./go-gin-postgres"]
