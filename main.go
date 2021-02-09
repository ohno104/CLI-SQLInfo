package main

import (
	"log"

	"felix.bs.com/felix/BeStrongerInGO/01_CLI/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalln("cmd.Excute err: ", err)
	}
}
