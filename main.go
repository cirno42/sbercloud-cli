/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/joho/godotenv"
	"sbercloud-cli/cmd"
)

func main() {
	_ = godotenv.Load(".config")
	cmd.Execute()
}
