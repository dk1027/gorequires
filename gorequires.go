package gorequires

import (
	"log"
	"os"
)

func Env(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("%v not defined", key)
		panic("Required environment variable not defined.")
	}
	return val
}

func File(path string) os.FileInfo {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Printf("%v does not exist", path)
		panic(err)
	}
	return fi
}
