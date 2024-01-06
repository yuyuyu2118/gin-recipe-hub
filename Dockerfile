# 使用するGoのバージョンを指定
FROM golang:1.20 as builder

# ワーキングディレクトリを設定
WORKDIR /app

# 依存関係をコピー
COPY go.mod ./
COPY go.sum ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o gin-recipe-hub

# 実行ステージ
FROM gcr.io/distroless/base-debian10

WORKDIR /

# ビルドステージからバイナリをコピー
COPY --from=builder /app/gin-recipe-hub .

# テンプレートと静的ファイルをコピー
COPY templates/ templates/

# ポートを公開
EXPOSE 8080

# 実行コマンド
CMD ["./gin-recipe-hub"]