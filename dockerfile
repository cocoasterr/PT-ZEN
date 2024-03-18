################
# BUILD BINARY #
################
# Gunakan golang:1.18.2-alpine3.16 sebagai base image
FROM golang:1.18.2-alpine3.16 as builder

# Install git + SSL ca certificates.
# Git diperlukan untuk mengambil dependensi.
# Ca-certificates diperlukan untuk melakukan panggilan HTTPS.
RUN apk update && apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /app

# Salin file go.mod dan go.sum ke dalam container dan jalankan go mod download untuk mengambil dependensi.
COPY go.mod .
COPY go.sum .
RUN go mod download

# Salin seluruh kode ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o /go/bin/ptzen app/main.go


#####################
# MAKE SMALL BINARY #
#####################
# Gunakan alpine:3.16 sebagai base image untuk membuat image yang lebih kecil
FROM alpine:3.16

# Salin binary dari builder stage
COPY --from=builder /go/bin/ptzen /go/bin/ptzen

# Atur working directory
WORKDIR /go/bin

# Jalankan aplikasi ketika container dimulai
CMD ["/go/bin/ptzen"]
