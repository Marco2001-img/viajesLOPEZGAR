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
# Usamos 'scratch' para la imagen final si no necesitas shell/librerías C (el tamaño es mínimo)
# Si tu aplicación *realmente* necesita CGO o librerías dinámicas, usa 'alpine:latest'
# En el contexto de tu error de zona horaria, usaremos 'alpine' y lo corregiremos.
FROM alpine:latest

WORKDIR /app

# INSTALACIÓN DE ZONA HORARIA (TZDATA):
# Esto es CRUCIAL para solucionar el error "unknown time zone America/Lima".
# Instalar tzdata permite al sistema operativo reconocer y usar zonas horarias.
RUN apk update && apk add --no-cache tzdata

# OPCIONAL PERO RECOMENDADO: Establece la zona horaria del contenedor a una de Latam (México/CDMX).
# Esto asegurará que cualquier proceso dentro del contenedor use esta TZ por defecto.
ENV TZ=America/Mexico_City

# Copiamos el binario desde el builder
COPY --from=builder /app/main .

# Variables de entorno opcionales
ENV PORT=8080

# Exponemos el puerto
EXPOSE 8080

# Comando para ejecutar la app
CMD ["./main"]