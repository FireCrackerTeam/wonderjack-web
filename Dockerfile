FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=root

# Create a new database for our Golang app
ENV MYSQL_DATABASE=wonderjack_web

# Set the character set and collation for the database
ENV MYSQL_CHARSET=utf8mb4
ENV MYSQL_COLLATION=utf8mb4_unicode_ci

FROM golang:1.20

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go get -u github.com/go-sql-driver/mysql
RUN go build -o ./out/dist .

CMD ./out/dist
