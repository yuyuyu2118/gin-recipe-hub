package main

import (
	"github.com/gin-gonic/gin"
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
	router := gin.Default()

	// レシピの一覧を取得するエンドポイント
	router.GET("/recipes", func(c *gin.Context) {
	})

	// 新しいレシピを投稿するエンドポイント
	router.POST("/recipes", func(c *gin.Context) {
	})

	// 特定のIDのレシピを取得するエンドポイント
	router.GET("/recipes/:id", func(c *gin.Context) {
	})

	// 特定のIDのレシピを削除するエンドポイント
	router.DELETE("/recipes/:id", func(c *gin.Context) {
	})

	router.Run(":8080")
}
