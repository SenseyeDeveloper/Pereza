package main

import "github.com/senseyedeveloper/pereza/bootstrap"

func main() {
	generator := bootstrap.NewGenerator(bootstrap.Settings{})

	generator.Run()
}
