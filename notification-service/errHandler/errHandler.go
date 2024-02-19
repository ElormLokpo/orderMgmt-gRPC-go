package errHandler

import "log"

func ErrHandler(err error, message string) {
	if err != nil {
		log.Fatal(err, message)
	}
}
