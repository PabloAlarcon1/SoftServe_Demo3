# Utilizar una imagen base de Go
FROM golang:latest

# Establecer el directorio de trabajo en /app
WORKDIR /app

# Copiar los archivos de la aplicación al directorio de trabajo
COPY . .

# Compilar la aplicación
RUN go build -o app

# Establecer las variables de entorno (opcional)
ENV PORT=8080

# Exponer el puerto en el contenedor
EXPOSE $PORT

# Establecer el comando de ejecución de la aplicación
CMD ["./app"]