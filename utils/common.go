package utils

import (
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
