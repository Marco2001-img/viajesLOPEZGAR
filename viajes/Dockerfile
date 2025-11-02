# Stage 1: build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copiamos los módulos y hacemos cache
COPY go.mod go.sum ./
RUN go mod download

# Copiamos el resto del código
COPY . .

# Compilamos la aplicación.
# Agregamos la bandera -ldflags para reducir el tamaño del binario y mejorar la eficiencia.
RUN go build -ldflags='-w -s' -o main .

# ----------------------------------------------------------------------------------

# Stage 2: runtime
FROM alpine:latest

WORKDIR /app

# INSTALACIÓN DE ZONA HORARIA (TZDATA):
RUN apk update && apk add --no-cache tzdata

# OPCIONAL PERO RECOMENDADO: Establece la zona horaria del contenedor.
ENV TZ=America/Mexico_City

# Copiamos el binario desde el builder
COPY --from=builder /app/main .

# >>>>>>>>>>>>>>>>>> CORRECCIÓN CLAVE AQUÍ <<<<<<<<<<<<<<<<<<<<
# Otorga permisos de ejecución al binario 'main' para solucionar "permission denied".
RUN chmod +x /app/main

# Variables de entorno
ENV PORT=8082

# Exponemos el puerto
EXPOSE 8082

# Comando para ejecutar la app
CMD ["./main"]