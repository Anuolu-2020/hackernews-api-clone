package pkg

import (
	"log"
	"os"
)

func GetEnv(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		log.Fatalf("%s env variable not found", name)
	}

	return value
}
