package main

import (
	"github.com/nikepan/govkbot"
	"log"
	"os"
	"strings"
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


func main(){
	var (
		env EnvMap
		result interface{}
		err error
	)

	env = getEnv("AUTO_ONLINE_")

	govkbot.SetAPI(env["TOKEN"], "", "")
	err = govkbot.API.CallMethod("account.setOnline", make(map[string]string, 0), &result)
	checkErr(err)
}
