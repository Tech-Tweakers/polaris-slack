package environment

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Env struct {
	ENVIRONMENT string
}

var lock = &sync.Mutex{}

type single struct {
	ENVIRONMENT      string
	LOG_LEVEL        string
	SLACK_AUTH_TOKEN string
	SLACK_CHANNEL_ID string
	SLACK_APP_TOKEN  string
	SLACK_API_URL    string
	AUTO_REPLY       string
}

func init() {

	err := godotenv.Load(".env.local")
	if err != nil {
		log.Println("Error loading environment file")
	}
	env := GetInstance()
	env.Setup()
}

func (e *single) Setup() {
	e.ENVIRONMENT = os.Getenv("ENVIRONMENT")
	e.LOG_LEVEL = os.Getenv("LOG_LEVEL")
	e.SLACK_AUTH_TOKEN = os.Getenv("SLACK_AUTH_TOKEN")
	e.SLACK_CHANNEL_ID = os.Getenv("SLACK_CHANNEL_ID")
	e.SLACK_APP_TOKEN = os.Getenv("SLACK_APP_TOKEN")
	e.SLACK_API_URL = os.Getenv("SLACK_API_URL")
	e.AUTO_REPLY = os.Getenv("AUTO_REPLY")
}

func (e *single) IsDevelopment() bool {
	return e.ENVIRONMENT == "development"
}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
			singleInstance.Setup()
		} else {
			fmt.Println("Single instance already created.")
		}
	}
	return singleInstance
}
