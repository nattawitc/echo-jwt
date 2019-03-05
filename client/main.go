package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	b, err := ioutil.ReadFile("jwtRSA.key")
	if err != nil {
		fmt.Println("fail to read key file jwtRSA.key:", err)
		return
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		fmt.Println("fail to parse rsa key:", err)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"foo": "bar",
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println("fail to sign jwt:", err)
		return
	}
	fmt.Println("jwt token:", tokenString)
	fmt.Println()

	req, err := http.NewRequest("GET", "http://localhost:9000", nil)
	if err != nil {
		fmt.Println("cannot create http request:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+tokenString)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("cannot call to server:", err)
		return
	}
	fmt.Println("==========================")
	fmt.Println("response")
	fmt.Println("==========================")
	io.Copy(os.Stdout, resp.Body)
	fmt.Println()
}
