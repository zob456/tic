package main

import (
	"github.com/zob456/tic/internal/config"
	"github.com/zob456/tic/internal/utils"
)

func init() {
	config.LoadConfig()
}

func main() {
	utils.InfoLogger("TIC started")
}
