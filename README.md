# gin-recipe-hub

### 課題: Gin を使ったレシピ共有アプリケーション

#### 目標:

Gin フレームワークを使用して、レシピを共有・閲覧できる Web アプリケーションを作成し、インターネット上で公開します。

#### 要件:

1. **レシピモデルの作成**:

   - レシピを表す構造体（`Recipe`）を作成します。各レシピには`ID`, `Title`, `Description`, `Ingredients`, `Instructions` のフィールドが必要です。

2. **エンドポイントの設定**:

   - `GET /recipes`: すべてのレシピを一覧表示。
   - `POST /recipes`: 新しいレシピを投稿。
   - `GET /recipes/:id`: 特定の ID のレシピを表示。
   - `DELETE /recipes/:id`: 特定の ID のレシピを削除。

3. **データストレージ**:

   - レシピはメモリ内に保存するか、任意のデータベースを使用して永続化します。

4. **フロントエンド**:

   - 簡単な HTML フロントエンドを提供し、レシピの一覧表示と新規投稿ができるようにします。

5. **公開**:

   - Heroku や Google Cloud Platform などのクラウドサービスを使用してアプリケーションを公開します。

6. **テスト**:
   - エンドポイントに対する単体テストを書いて、機能が正しく動作することを確認します。

#### ステップバイステップの指示:

1. Gin フレームワークをプロジェクトにインストールします。
2. `Recipe` 構造体を定義し、必要なフィールドを含めます。
3. 必要なエンドポイントを定義し、それぞれのエンドポイントに対応するハンドラ関数を作成します。
4. レシピのデータを保存するためのインメモリデータベースまたは外部データベースをセットアップします。
5. フロントエンドの HTML ページを作成し、レシピの一覧表示と新規投稿フォームを用意します。
6. ローカルでアプリケーションをテストし、すべてのエンドポイントが期待通りに機能することを確認します。
7. アプリケーションをクラウドプラットフォームにデプロイし、公開します。
8. 公開したアプリケーションにアクセスして、インターネット上で正常に動作することを確認します。

### コマンドメモ

```
docker cp gin-recipe-hub-app-1:/app/templates/index.html .
<!-- データベース削除 -->
docker volume rm postgres_data
<!-- heroku変更をコミットしてHerokuにpush -->
git add .
git commit -m "Use DATABASE_URL for Heroku PostgreSQL connection"
git push heroku main
```

### ステップ 2: Docker イメージの作成と ECR へのプッシュ

1. **Docker イメージをビルド**します。Dockerfile があるディレクトリで以下のコマンドを実行します。

   ```sh
   docker build -t your-app-name .
   ```

2. **Amazon Elastic Container Registry (ECR) リポジトリを作成**します。これは、Docker イメージを AWS に保存する場所です。
   AWS Management Console で ECR に移動し、「Create repository」をクリックして新しいリポジトリを作成します。

3. **AWS CLI を使用して ECR にログイン**します。

   ```sh
   aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 387107580630.dkr.ecr.ap-northeast-1.amazonaws.com
   ```

4. **Docker イメージにタグを付け**、ECR リポジトリの URI を使用します。

   ```sh
   docker tag your-app-name:latest your-aws-account-id.dkr.ecr.your-region.amazonaws.com/your-app-name:latest

   docker tag gin-recipe-hub-app:latest 387107580630.dkr.ecr.ap-northeast-1.amazonaws.com/gin-recipe-hub-app:latest
   ```

5. **イメージを ECR にプッシュ**します。

   ```sh
   docker push your-aws-account-id.dkr.ecr.your-region.amazonaws.com/your-app-name:latest

   docker push 387107580630.dkr.ecr.ap-northeast-1.amazonaws.com/gin-recipe-hub-app:latest
   ```
