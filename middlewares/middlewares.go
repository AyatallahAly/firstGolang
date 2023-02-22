package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const Admin = "admin"
const Doctor = "doctor"
const Patinet = "patinet"

func RequiredAuth(c *gin.Context) {
	fmt.Println("IN middleware")

	//Get the cookie off req
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	// Parse takes the token string and a function for looking up the key. The latter is especially
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")),err
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check the exp
		// if float64(time.Now().Unix()) > claims["exp"].(float64){
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// }

		//check the roles

		if claims["role"] == Admin {
			c.Set("role", Admin)
			return
		}
		if claims["role"] == Doctor {
			c.Set("role", Doctor)
			return
		}
		if claims["role"] == Patinet {
			c.Set("role", Patinet)
			return
		}
		//continue
		c.Next()
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
