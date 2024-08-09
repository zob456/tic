package main

import (
	"fmt"

	"github.com/zob456/tic/internal/config"
	"github.com/zob456/tic/internal/gzip_handler"
	"github.com/zob456/tic/internal/utils"
)

func init() {
	config.LoadConfig()
}

func main() {
	urlListBuilder := gzip_handler.New(5, 1922320936)
	urlList, failedToParseUrls, err := urlListBuilder.FilterGzip()
	if err != nil {
		panic(err)
	}

	if len(failedToParseUrls) > 0 {
		panic(fmt.Errorf("url(s) failed to parse: %v", failedToParseUrls))
	}

	utils.InfoLogger(urlList)
}
