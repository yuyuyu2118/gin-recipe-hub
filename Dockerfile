# ビルドステージ
FROM golang:1.20 as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
# 静的にリンクされたバイナリをビルドするコマンドを追加
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o gin-recipe-hub

# 実行ステージ
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=builder /app/gin-recipe-hub .
COPY templates/ templates/
EXPOSE 8080
CMD ["./gin-recipe-hub"]