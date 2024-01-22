package database

import (
	"log"

	"github.com/SimonMora/bikesams_users/models"
	"github.com/SimonMora/bikesams_users/util"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(user models.SignUp) error {
	log.Default().Println("Starting signup user.. saving in database")

	err = DbConnect()
	if err != nil {
		log.Default().Println(err.Error())
		return err
	}

	defer Db.Close()

	sqlSentence := " INSERT INTO users (User_UUID, User_Email, User_DateAdd) VALUES ('" + user.UserUUID + "', '" + user.UserEmail + "', '" + util.DateSqlFormat() + "')"
	_, err = Db.Exec(sqlSentence)

	if err != nil {
		log.Default().Println(err.Error())
		return err
	}

	log.Default().Println("Signup user successfully executed..")
	return nil
}
