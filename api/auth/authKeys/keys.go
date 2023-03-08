package authKeys

import (
	"log"
	"os"
)

type AkSkKeys struct {
	AccessKey string
	SecretKey string
}

func GetAkSkKeys() AkSkKeys {
	ak := os.Getenv("ACCESS_KEY")
	sk := os.Getenv("SECRET_KEY")
	if (sk == "") || (ak == "") {
		log.Println("ERROR: Can't find AK/SK in environment variables. Is .config file configured right?")
		return AkSkKeys{}
	}
	return AkSkKeys{
		AccessKey: ak,
		SecretKey: sk,
	}
}
