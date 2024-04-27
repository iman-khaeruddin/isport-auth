package main

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/iman-khaeruddin/isport-auth/midleware"
	"github.com/iman-khaeruddin/isport-auth/modules/signin"
	"github.com/iman-khaeruddin/isport-auth/modules/signup"
	"github.com/iman-khaeruddin/isport-auth/utils/db"
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
