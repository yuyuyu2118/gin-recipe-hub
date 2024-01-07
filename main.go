package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Recipe struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
}

var recipes []Recipe

var sampleRecipes = []Recipe{
	{
		ID:          "1",
		Title:       "チキンカレー",
		Description: "チキンと野菜のカレーです",
		Ingredients: []string{
			"玉ねぎ",
			"にんじん",
			"じゃがいも",
			"肉",
			"カレールー",
			"水",
		},
		Instructions: []string{
			"玉ねぎ、にんじん、じゃがいもを食べやすい大きさに切る",
			"鍋にサラダ油を入れて熱し、肉を炒める",
			"玉ねぎ、にんじん、じゃがいもの順に入れて炒める",
			"水を入れて具材が浸るまで煮る",
			"沸騰したらカレールーを入れて溶かす",
			"好みでカレー粉を入れて味を整える",
		},
	},
	{
		ID:          "2",
		Title:       "オムライス",
		Description: "玉ねぎと卵で作るオムライスです",
		Ingredients: []string{
			"玉ねぎ",
			"卵",
			"醤油",
			"ケチャップ",
			"サラダ油",
		},
		Instructions: []string{
			"玉ねぎを炒める",
			"卵を溶いて醤油と塩で味付けする",
			"卵と玉ねぎを混ぜ合わせる",
			"ご飯の上に卵をかける",
			"ケチャップをかけて完成",
		},
	},
}

func init() {
	recipes = sampleRecipes
}

func main() {
	// // 環境変数からデータベース接続情報を取得
	// dbHost := os.Getenv("DB_HOST")
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// dbPort := os.Getenv("DB_PORT")

	// // データベース接続文字列を構築
	// connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	// 	dbHost, dbUser, dbPassword, dbName, dbPort)

	// HerokuのDATABASE_URL環境変数からデータベース接続情報を取得
	databaseUrl := os.Getenv("DATABASE_URL")

	// データベースに接続
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// データベース接続を確認
	err = db.Ping()
	if err != nil {
		log.Fatal("データベース接続エラー: ", err)
	}

	// データベーススキーマの作成
	createTableSQL := `CREATE TABLE IF NOT EXISTS recipes (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		ingredients TEXT[] NOT NULL,
		instructions TEXT[] NOT NULL
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("テーブルの作成に失敗しました: %v", err)
	}

	fmt.Println("Successfully connected to the database!")

	router := gin.Default()

	// HTMLテンプレートをロードする
	router.LoadHTMLGlob("templates/*.html")

	// 静的ファイルのディレクトリを指定する
	router.Static("/assets", "./templates/assets")

	// ルートURLにアクセスしたときにindex.htmlをレンダリングする
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "ホームページ",
		})
	})

	// レシピの一覧を取得するエンドポイント
	router.GET("/recipes", func(c *gin.Context) {
		// データベースからレシピを取得するクエリ
		rows, err := db.Query("SELECT id, title, description, ingredients, instructions FROM recipes")
		if err != nil {
			c.JSON(500, gin.H{"error": "データベースからのレシピの取得に失敗しました"})
			return
		}
		defer rows.Close()

		// レシピを格納するスライスを初期化
		var recipes []Recipe

		// クエリ結果をループしてレシピをスライスに追加
		for rows.Next() {
			var r Recipe
			// ingredients, instructions が文字列のスライスである場合、
			// データベースのデータ型に応じて適切なスキャン方法を選択する必要があります。
			// 以下は、単純な文字列フィールドとしてスキャンする例です。
			if err := rows.Scan(&r.ID, &r.Title, &r.Description, pq.Array(&r.Ingredients), pq.Array(&r.Instructions)); err != nil {
				c.JSON(500, gin.H{"error": "レシピの読み込みに失敗しました"})
				return
			}
			recipes = append(recipes, r)
		}

		// エラーのチェック
		if err = rows.Err(); err != nil {
			c.JSON(500, gin.H{"error": "レシピの読み込み中にエラーが発生しました"})
			return
		}

		// レシピのスライスをJSONとして返す
		c.JSON(200, recipes)
	})

	// 新しいレシピを投稿するエンドポイント
	router.POST("/recipes", func(c *gin.Context) {
		var newRecipe Recipe
		if err := c.BindJSON(&newRecipe); err != nil {
			c.JSON(400, gin.H{"error": "リクエストが正しくありません"})
			return
		}

		// データベースに新しいレシピを保存するSQL文を実行
		_, err := db.Exec("INSERT INTO recipes (title, description, ingredients, instructions) VALUES ($1, $2, $3, $4)",
			newRecipe.Title, newRecipe.Description, pq.Array(newRecipe.Ingredients), pq.Array(newRecipe.Instructions))
		if err != nil {
			c.JSON(500, gin.H{"error": "データベースへの保存に失敗しました: " + err.Error()})
			return
		}

		c.JSON(201, newRecipe)
	})

	// 特定のIDのレシピを取得するエンドポイント
	router.GET("/recipes/:id", func(c *gin.Context) {
		id := c.Param("id")

		// データベースから指定されたIDのレシピを取得するSQLクエリを実行
		var recipe Recipe
		err := db.QueryRow("SELECT id, title, description, ingredients, instructions FROM recipes WHERE id = $1", id).Scan(&recipe.ID, &recipe.Title, &recipe.Description, pq.Array(&recipe.Ingredients), pq.Array(&recipe.Instructions))
		if err != nil {
			if err == sql.ErrNoRows {
				// レシピが見つからない場合は404エラーを返す
				c.JSON(404, gin.H{"error": "レシピが見つかりませんでした"})
			} else {
				// その他のエラーの場合は500エラーを返す
				c.JSON(500, gin.H{"error": "データベースのクエリ中にエラーが発生しました: " + err.Error()})
			}
			return
		}

		// レシピをJSONとして返す
		c.JSON(200, recipe)
	})

	// 特定のIDのレシピを削除するエンドポイント
	router.DELETE("/recipes/:id", func(c *gin.Context) {
		id := c.Param("id")

		// データベースから指定されたIDのレシピを削除するSQL文を実行
		result, err := db.Exec("DELETE FROM recipes WHERE id = $1", id)
		if err != nil {
			// SQL実行中にエラーが発生した場合は500エラーを返す
			c.JSON(500, gin.H{"error": "データベースからのレシピの削除に失敗しました: " + err.Error()})
			return
		}

		// 削除されたレコードの数を確認
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			c.JSON(500, gin.H{"error": "削除されたレシピの数の取得に失敗しました: " + err.Error()})
			return
		}

		if rowsAffected == 0 {
			// レシピが見つからない場合は404エラーを返す
			c.JSON(404, gin.H{"error": "レシピが見つかりませんでした"})
			return
		}

		// 削除に成功した場合は200ステータスコードを返す
		c.JSON(200, gin.H{"message": "レシピが削除されました"})
	})

	// 環境変数からポートを取得し、デフォルトのポートを設定
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // デフォルトポート
	}

	// 環境変数で指定されたポートでアプリケーションを起動
	router.Run(":" + port)
}
