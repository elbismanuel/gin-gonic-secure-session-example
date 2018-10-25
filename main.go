package main

import (

	//Gin dependency
	"net/http"

	"github.com/gin-gonic/gin"
	//Gorilla dependency
	"github.com/gorilla/securecookie"
)

//Define the key value for the session
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func main() {

	//Router
	r := gin.New()

	//http://localhost:2000/cookie
	r.GET("/cookie", func(c *gin.Context) {

		//Values for the cookie
		value := map[string]string{
			"name":    "Foo",
			"surname": "Bar",
		}

		//Encode cookie with Gorilla toolkit
		if encoded, err := cookieHandler.Encode("go_session", value); err == nil {

			cookie := &http.Cookie{
				//Name of the session
				Name: "go_session",
				//Encoded values
				Value: encoded,
			}

			//Create cookie
			http.SetCookie(c.Writer, cookie)
		}

		c.String(http.StatusOK, "Your session was created")
	})

	r.Run(":2000")
}
