package main

import (
	"log"
	"os"
	"strings"

	"github.com/MattAitchison/froyo/provider"
)

// This is terrible...
func getEnv() map[string]string {
	env := map[string]string{}
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if strings.HasPrefix(pair[0], "FROYO") {
			key := pair[0]
			val := strings.Join(pair[1:], "=")
			env[key] = val
		}
	}

	return env
}
func main() {
	env := getEnv()

	sa, _ := provider.NewSceneAccessTracker()

	sa.Login(env["FROYO_SCENEACCESS_USERNAME"], env["FROYO_SCENEACCESS_PASSWORD"])
	res, _ := sa.Search(os.Args[1])
	log.Println(res)
}
