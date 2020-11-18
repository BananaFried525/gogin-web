package utils

import (
	"log"
)

func HandleError(err error) {
	if err != nil {
		log.Println("ERROR : " + err.Error())
	}
	return
}
