package main

import (
	"fmt"
	"formative-14/configs/database"
	"formative-14/modules/user"
)

const port = 5000

func main() {
	database.InitializeDB()

	user.UserRouter().Run(fmt.Sprintf(":%d", port))
}