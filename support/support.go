package support

import "log"

// Handle Error
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
