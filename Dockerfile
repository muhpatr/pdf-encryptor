# ---------- STAGE 1: Builder ----------
FROM golang:1.24-alpine AS builder

# Install git (diperlukan untuk go mod), ca-certificates, dan alat minimal
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go.mod dan go.sum (untuk cache layer build)
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build aplikasi ke dalam satu binary statis
RUN CGO_ENABLED=0 GOOS=linux go build -o pdf-encryptor .

# ---------- STAGE 2: Final Runtime ----------
FROM alpine:latest

# Buat user non-root
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Copy hanya binary dan docs (tanpa source code!)
COPY --from=builder /app/pdf-encryptor .
COPY --from=builder /app/docs ./docs

# Ubah ownership supaya hanya user non-root yang punya akses
RUN chown -R appuser:appgroup /app

# Switch ke user non-root
USER appuser

# Expose port yang digunakan
EXPOSE 7082

# Jalankan aplikasi
CMD ["./pdf-encryptor"]    