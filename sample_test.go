package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	// Ginエンジンのセットアップ
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// メインのルートのセットアップ
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// HTTPリクエストの作成
	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// レスポンスのレコーダ
	recorder := httptest.NewRecorder()

	// リクエストの実行
	router.ServeHTTP(recorder, req)

	// ステータスコードのチェック
	assert.Equal(t, http.StatusOK, recorder.Code)

	// レスポンスボディのチェック
	assert.Equal(t, `{"message":"pong"}`, recorder.Body.String())
}
