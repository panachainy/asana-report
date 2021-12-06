package main

import "asana-report/cmd"

var VERSION = "development"

func main() {
	cmd.Execute(VERSION)
}
