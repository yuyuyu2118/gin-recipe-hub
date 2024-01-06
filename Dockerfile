# ビルドステージ
FROM golang:1.20 as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o gin-recipe-hub

# Airをインストール
RUN go install github.com/cosmtrek/air@latest

# 実行ステージ
FROM golang:1.20
WORKDIR /app
COPY --from=builder /app/gin-recipe-hub .
# $GOPATH/bin/air からバイナリをコピー
COPY --from=builder /go/bin/air /usr/local/bin/air
COPY . .
COPY templates/ templates/
# .air.tomlをコピー
COPY .air.toml ./
EXPOSE 8080

# Airを起動コマンドとして設定
CMD ["air", "-c", ".air.toml"]