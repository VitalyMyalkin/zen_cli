package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

var db *sql.DB
var hostport string
var err error

type KeytoIncrement struct {
	key   string `json:"key"`
	value int    `json:"value"`
}

type TexttoHex struct {
	text string `json:"id"`
	key  string `json:"title"`
}

type User struct {
	id   int    `json:"id"`
	name string `json:"name"`
	age  int    `json:"age"`
}

func incrementKey(c *gin.Context) {
	var newKey KeytoIncrement

	// распарсить ключ и инкремент
	if err := c.BindJSON(&newKey); err != nil {
		return
	}
	// подключение к редис
	rdb := redis.NewClient(&redis.Options{
		Addr:     hostport,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//инкрементируем на значение и отправляем ответ
	answer, err := rdb.Do(context.Background(), "incrby", newKey.key, newKey.value).Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"value": answer,
	})
}

func getSign(c *gin.Context) {
	var newTexttoHex TexttoHex

	// распарсиваем текст и ключ
	if err := c.BindJSON(&newTexttoHex); err != nil {
		return
	}
	//делаем подпись
	key := []byte(newTexttoHex.key)

	signature := hmac.New(sha512.New, key)
	signature.Write([]byte(newTexttoHex.text))

	//подпись в hex-строку и отправляем ответ
	c.IndentedJSON(http.StatusOK, hex.EncodeToString(signature.Sum(nil)))
}

func addUser(c *gin.Context) {
	var newUser User

	// распарсиваем данные нового пользователя
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	//создаем таблицу пользователей если ее нет
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, age INTEGER)")
	if err != nil {
		return
	}

	//добавляем его данные в таблицу пользователей
	err := db.QueryRow("INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id", newUser.name, newUser.age).Scan(&newUser.id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	//отправляем айдишник нового пользователя
	c.JSON(http.StatusCreated, gin.H{
		"id": newUser.id,
	})
}

func main() {
	//считываем хост и порт для редиса
	var host, port string
	fmt.Println("Введите хост подключения к Redis")
	fmt.Scanf("%s\n", &host)

	fmt.Println("Введите порт подключения к Redis")
	fmt.Scanf("%s\n", &port)

	hostport = host + ":" + port

	//подключаемся к постгрес
	db, err = sql.Open("postgres", "host=localhost port=5433 user=postgres password=postgres dbname=zen sslmode=disable")
	if err != nil {
		return
	}
	defer db.Close()
	// задаем роутер и хендлеры
	router := gin.Default()
	router.POST("/redis/incr", incrementKey)
	router.POST("/sign/hmacsha512", getSign)
	router.POST("/postgres/users", addUser)

	// запускаем сервер
	router.Run("localhost:8080")
}
