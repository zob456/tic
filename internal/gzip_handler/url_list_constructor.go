package gzip_handler

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/zob456/tic/internal/config"
	"github.com/zob456/tic/internal/utils"
)

type UrlListBuilder struct {
	FileChunksLimit      int
	FileChunkMemoryLimit int
	FileUrl              string
	HttpClient           *http.Client
}

func New(fileChunksLimit, fileChunkMemoryLimit int) *UrlListBuilder {
	return &UrlListBuilder{
		FileChunksLimit:      fileChunksLimit,
		FileChunkMemoryLimit: fileChunkMemoryLimit,
		FileUrl:              config.ENV.FileUrl,
		HttpClient:           http.DefaultClient,
	}
}

func (u *UrlListBuilder) FilterGzip() ([]string, []string, error) {
	req, err := http.NewRequest(http.MethodGet, u.FileUrl, nil)
	if err != nil {
		return nil, nil, utils.ErrorLoggerWithReturn(err)
	}

	resp, err := u.fetchGzip(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var r *gzip.Reader
	r, err = gzip.NewReader(resp.Body)
	if err != nil {
		return nil, nil, utils.ErrorLoggerWithReturn(err)
	}
	defer r.Close()

	var urls [][]string

	b := make([]byte, u.FileChunkMemoryLimit)
	for range u.FileChunksLimit {
		_, err = r.Read(b)
		if err != nil {
			return nil, nil, utils.ErrorLoggerWithReturn(err)
		}

		if strings.Contains(string(b), "New York") {
			urls = append(urls, strings.Split(string(b), "description"))
		}
		continue
	}
	urlList, failedToParseUrls := urlListConstructor(urls)
	return urlList, failedToParseUrls, nil
}

func urlListConstructor(urlStringBlocks [][]string) ([]string, []string) {
	var urls []string
	var failedToParseUrls []string
	for _, s := range urlStringBlocks {
		for _, s2 := range s {
			if strings.Contains(s2, "New York") {
				baseUrl := strings.TrimPrefix(strings.Split(s2, "location")[1], ":")
				urlStringToParse := strings.Trim(strings.SplitAfterN(baseUrl, ":", 2)[1], `"},{"`)
				parsedUrl, err := url.Parse(urlStringToParse)
				if err != nil {
					failedToParseUrls = append(failedToParseUrls, fmt.Errorf("failed to parse URL fro expected URL string. string: %v ERROR: %v", urlStringToParse, err).Error())
				}
				urls = append(urls, parsedUrl.String())
			}
		}
	}
	return urls, failedToParseUrls
}

func (u *UrlListBuilder) fetchGzip(req *http.Request) (*http.Response, error) {
	resp, err := u.HttpClient.Do(req)
	if err != nil {
		return nil, utils.ErrorLoggerWithReturn(err)
	}
	return resp, nil
}
