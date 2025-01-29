package routes

import (
	"fmt"

	"api.forex.com/models"
	"api.forex.com/storage"
	"github.com/kataras/iris/v12"
)

func GetContacts(ctx iris.Context) {
	var user models.Fx_User
	users := storage.DB.Find(&user)
	if users.Error != nil {
		fmt.Println(users.Error)
		return
	}
	ctx.JSON(iris.Map{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
	})
}
