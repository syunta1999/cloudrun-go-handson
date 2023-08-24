package main

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		panic(err)
	}

	DB_PASSWORD := "projects/20799997600/secrets/DB_PASSWORD/versions/latest"
	DB_USER := "projects/20799997600/secrets/DB_USER/versions/latest"
	passed, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{Name: DB_PASSWORD})
	if err != nil {
		panic(err)
	}
	dbUser, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{Name: DB_USER})
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(aws.connect.psdb.cloud)/neko?tls=true", string(dbUser.Payload.Data), string(passed.Payload.Data))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Deployed",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		var users []User
		db.Unscoped().Find(&users)
		c.JSON(200, users)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
