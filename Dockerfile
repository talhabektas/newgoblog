# Go uygulamasını derlemek için bir yapı aşaması oluştur
FROM golang:1.20-alpine AS builder

# Çalışma dizini ayarla
WORKDIR /app

# go.mod ve go.sum dosyalarını kopyala ve bağımlılıkları indir
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Tüm uygulama dosyalarını kopyala
COPY . .

# Uygulamayı derle
RUN go build -o app .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/app .

COPY .env .env

EXPOSE 80
EXPOSE 8000

# Uygulamayı çalıştır
CMD ["./app"]
