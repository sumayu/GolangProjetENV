package envupdate

import (
	"database/sql"
	"main/envchecker"
	"main/logger"
	mydb "main/myDB"

) 
func EnvUpdateProd() {
		
	db := mydb.Database()
	_, err := db.Exec (`
TRUNCATE TABLE envchecker;
INSERT INTO envchecker (env) VALUES ('prod');

	`)
	
	if err != nil {
		logger.Logger.Info("error message to db",err)
	}
	logger.Logger.Info("env = " , envchecker.Envchecker())

	}
	func EnvUpdateDebug() {
		db := mydb.Database()
		_, err := db.Exec (`
TRUNCATE TABLE envchecker;
INSERT INTO envchecker (env) VALUES ('debug');

		`)
		if err != nil {
			logger.Logger.Info("error message to db",err)

		}
		logger.Logger.Info("env = " , envchecker.Envchecker())
	}
		func Rt () (*sql.DB) {
			db := mydb.Database()
			return db
		}