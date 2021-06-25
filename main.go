package main

import "asana-report/cmd"

var VERSION = "develop"

func main() {
	cmd.Execute(VERSION)
}
