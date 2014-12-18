package gophersauras

import (
	"errors"
	"log"
)

func Relieve(err *error) {
	if r := recover(); r != nil {

		switch x := r.(type) {
		case string:
			*err = errors.New(x)
		case error:
			*err = x
		default:
			*err = errors.New("Unknown panic")
		}
		log.Println("relieved panic:")
		log.Println((*err).Error())
		log.Println()
	}
}

func FatalShock() {
	var err error
	Relieve(&err)
	Check(err)
}
