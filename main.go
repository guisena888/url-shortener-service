package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"url-shortener-service/db"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, world")
	router := gin.Default()
	db.ConnectDatabase()
	router.GET("/:hash", getUrl)
	router.POST("/api/v1/shorten", createShortenUrl)
	router.Run("localhost:8081")
}

func getUrl(c *gin.Context) {
	hash := c.Param("hash")
	fmt.Println("Redirecting... " + hash)
	urlMap := &db.UrlMap{}
	db.Db.First(urlMap, "hash = ?", hash)
	c.Redirect(http.StatusFound, urlMap.LongUrl)
}

func createShortenUrl(c *gin.Context) {
	var shortenUrlRequest ShortenUrlRequest

	if err := c.BindJSON(&shortenUrlRequest); err != nil {
		fmt.Println("Bad request ", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	hash := sha1.New()
	hash.Write([]byte(shortenUrlRequest.LongUrl))
	shortURL := base64.URLEncoding.EncodeToString(hash.Sum(nil))[:8]
	urlMap := &db.UrlMap{
		Hash:    shortURL,
		LongUrl: shortenUrlRequest.LongUrl,
	}

	db.Db.Create(urlMap)
	c.IndentedJSON(http.StatusCreated, urlMap)
}

type ShortenUrlRequest struct {
	LongUrl string `json:"longUrl"`
}
