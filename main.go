package main

import (
	middleware "auth/midleware"
	"auth/modules/signin"
	"auth/modules/signup"
	"auth/utils/db"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(
		middleware.AllowCORS(),
	)
	iSportDB := db.GormMysql(os.Getenv("ISPORT_MYSQL_DSN"))

	signup := signup.NewSignRequestHandler(iSportDB)
	signup.HandleSignup(router)

	signin := signin.NewSignRequestHandler(iSportDB)
	signin.HandleSignin(router)

	err := router.Run()

	if err != nil {
		log.Println("main router.Run:", err)
		return
	}
}
