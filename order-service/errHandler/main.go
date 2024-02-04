package errHandler

import "log"

func StdErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
