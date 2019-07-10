package main

import (
	"fmt"
	"github.com/nikepan/govkbot"
	"log"
	"os"
	"strings"
	"time"
)

type EnvMap map[string]string

func getEnv(prefix string) (env EnvMap) {
	env = EnvMap{}
	_env := os.Environ()
	for _, s := range _env {
		if strings.HasPrefix(s, prefix) {
			s = strings.TrimPrefix(s, prefix)
			separated := strings.Split(s, "=")
			env[separated[0]] = separated[1]
		}
	}
	return
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func run(){
	var (
		result interface{}
		err error
	)
	err = govkbot.API.CallMethod("account.setOnline", make(map[string]string, 0), &result)
	checkErr(err)
}

func main(){
	var (
		env EnvMap
		ticker = time.Tick(5 * time.Minute)
	)
	env = getEnv("AUTO_ONLINE_BOT_")
	govkbot.SetAPI(env["TOKEN"], "", "")

	for now := range ticker {
		fmt.Printf("%v - setting online\n", now)
		run()
	}
}
