package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID   int    `json: "id"`
	Name string `json: "name"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/name", name)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func name(c echo.Context) error {

	// 期待するURI
	// http://localhost:1323/name?id=2
	id := c.QueryParam("id")
	// return c.String(http.StatusOK, id)
	_, err := getName(id)

	if err != nil {
		// 500を返す
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// userをbindする
	u, err := bindUser()
	if err != nil {
		// log.Fatalf("ERROR: %v", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	log.Printf("名前： %v", u.Name)
	// u instanceからidで検索し、nameを返す。
	var result string
	result = "ほげ"
	// ここに処理を書く
	// 200を返す
	return c.String(http.StatusOK, result)

}

// func show(c echo.Context) error {
// 	// Get team and member from the query string
// 	team := c.QueryParam("team")
// 	member := c.QueryParam("member")
// 	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
// }

func getName(id string) (string, error) {

	if id == "1" {
		return "阪田　将", nil
	} else if id == "2" {
		return "栗原　協", nil
	} else if id == "3" {
		return "満永　淳一", nil
	}

	err := errors.New("エラーです。")
	return "", err
}

func bindUser() (*User, error) {
	raw, err := ioutil.ReadFile("./db.json")
	if err != nil {
		fmt.Println(err.Error())
		// os.Exit(1)
		return nil, errors.New("ファイルが見つかりませんでした")
	}
	u := new(User)
	err = json.Unmarshal(raw, &u)
	if err != nil {
		log.Fatalf("ERROR: %v", err.Error())
	}
	return u, nil
}
