FROM golang:1.20.4-alpine3.18

ARG DB_USER
ARG DB_PASSWORD

ENV DB_USER=${DB_USER}
ENV DB_PASSWORD=${DB_PASSWORD}

WORKDIR /app

COPY ./ ./
RUN go mod download

# バイナリファイルにビルド
RUN GOOS=linux GOARCH=amd64 go build -mod=readonly -v -o server

EXPOSE 8080

# バイナリファイルを実行
CMD ["./server"]