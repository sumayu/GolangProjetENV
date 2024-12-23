package envchecker

import (
	"fmt"
	mydb "main/myDB"
)

func Envchecker() string {
	db := mydb.Database()
	if db == nil {
		return "" // Прекращаем выполнение, если не удалось подключиться к БД
	}

	var envchecker string
	err := db.QueryRow("SELECT env FROM envchecker").Scan(&envchecker)
	if err != nil {
		fmt.Errorf("get env from DB error: %v", err)
		return ""
	}
	return envchecker
}