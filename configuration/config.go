package configuration

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"reflect"
	"strconv"
)

type Configuration struct {
	RootUsername string
	RootPassword string
	MongoUri     string
	DbName       string
	Port         int
}

var ServerConfiguration Configuration

func (conf *Configuration) Load(envPath string) {
	err := godotenv.Load(envPath)
	fmt.Println(os.Getwd())
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	conf.MongoUri = os.Getenv("MONGO_URI")
	conf.DbName = os.Getenv("DB_NAME")
	conf.Port, err = strconv.Atoi(os.Getenv("PORT"))
	conf.RootUsername = os.Getenv("DB_ROOT_USERNAME")
	conf.RootPassword = os.Getenv("DB_ROOT_PASSWORD")
	if err != nil {
		log.Fatal("Error loading .env file, wrong PORT data type")
	}
}

func (conf Configuration) Print() {
	v := reflect.ValueOf(conf)
	typeOf := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %-10s \t Value: %v\n", typeOf.Field(i).Name, v.Field(i).Interface())
	}
}
