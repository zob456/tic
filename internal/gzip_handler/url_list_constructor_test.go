package gzip_handler

import (
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zob456/tic/internal/config"
)

func init() {
	err := os.Setenv("FILE_URL", "https://antm-pt-prod-dataz-nogbd-nophi-us-east1.s3.amazonaws.com/anthem/2024-08-01_anthem_index.json.gz")
	if err != nil {
		panic(err)
	}

	config.LoadConfig()
}

const (
	nyPresentInput    = `{"description":"Highmark BS Northeastern NY : Highmark Blue Shield of Northeastern New York - PPO","location":"https://anthembcbsoh.mrf.bcbs.com/2024-08_800_72A0_in-network-rates_02_of_02.json.gz?&Expires=1726840845&Signature=uCkjd8P0JFuAfo6IvsW4ioKVeww2Y1H3rae0~7oXkIddycuC1aTN8IsLNGLK3uW4ROWCTG2twwFGLd~rEvkwG1pKGkTqDDVQtjHw8A0q~nxz14MsD7tz7XuHp~Cnvw4G2ugfrOytnqPVxq11Ai4iP60L7oSoDhKtSaUYT44YdbA1O-ZAIiWx~ypHx5VO4Fqc3V4uTzsJtVMYRQsfKffUq2p9sBJO7n7n1St~nTcQEzPNZEiTUn6glK0Z9S2ENwAxylTLWxLLh6RGM8Pic3NPMwmKp5gZ7AdNGBcO67RGglv9mEw4iXv861RfIkVEOljLHZkfvVl-z8ALAf8sHKskxw__&Key-Pair-Id=K27TQMT39R1C8A"},{"`
	nyNotPresentInput = `{"description":"BCBS Tennessee, Inc. : Network P","location":"https://anthembcbsoh.mrf.bcbs.com/2024-08_890_58B0_in-network-rates_52_of_60.json.gz?&Expires=1726840845&Signature=S~DtzrtOlcU-XvMYeb2RQxjRfih~Fn9rdi3THOfDZud1JuxWY38ikovMlJzeA95sKtBn-a~26SFYXuFjR8~qUGMKZscNCGZNMc8Xd1gpb2~O0x7-gxkKHHv8fKuKtWP2HamgQxUgeloLq5y1SzlSFei-vHhouZx6mCHchNYdte-tD3PFqI1MKT~q7hY4hGgrE2P0ZM6-r15THyLRWI33fDnS6Wzg7VcZTQee5Wr~sr-oPC~0kDTcaDEVP2ItGtFX4u5rMOX6cdQThPuG6MYJHRvltFokZU~5-KctEzmE68PPU8HRJ2H5D2K~0BYyYpzreZ2J8Bw01Cs0kT-FeLE1Xg__&Key-Pair-Id=K27TQMT39R1C8A"},{"`
	badUrlInput       = `{"description":"Highmark BS Northeastern NY : Highmark Blue Shield of Northeastern New York - PPO","location":"https:anthembcbsoh.mrf.bcbs.com/2024-08_890_58B0_in-network-rates_52_of_60.json.gz?&Expires=1726840845&Signature=S~DtzrtOlcU-XvMYeb2RQxjRfih~Fn9rdi3THOfDZud1JuxWY38ikovMlJzeA95sKtBn-a~26SFYXuFjR8~qUGMKZscNCGZNMc8Xd1gpb2~O0x7-gxkKHHv8fKuKtWP2HamgQxUgeloLq5y1SzlSFei-vHhouZx6mCHchNYdte-tD3PFqI1MKT~q7hY4hGgrE2P0ZM6-r15THyLRWI33fDnS6Wzg7VcZTQee5Wr~sr-oPC~0kDTcaDEVP2ItGtFX4u5rMOX6cdQThPuG6MYJHRvltFokZU~5-KctEzmE68PPU8HRJ2H5D2K~0BYyYpzreZ2J8Bw01Cs0kT-FeLE1Xg__&Key-Pair-Id=K27TQMT39R1C8A"},{"description":"Highmark BS Northeastern NY : Highmark Blue Shield of Northeastern New York - PPO","location":"https://anthembcbsoh.mrf.bcbs.com/2024-08_800_72A0_in-network-rates_02_of_02.json.gz?&Expires=1726840845&Signature=uCkjd8P0JFuAfo6IvsW4ioKVeww2Y1H3rae0~7oXkIddycuC1aTN8IsLNGLK3uW4ROWCTG2twwFGLd~rEvkwG1pKGkTqDDVQtjHw8A0q~nxz14MsD7tz7XuHp~Cnvw4G2ugfrOytnqPVxq11Ai4iP60L7oSoDhKtSaUYT44YdbA1O-ZAIiWx~ypHx5VO4Fqc3V4uTzsJtVMYRQsfKffUq2p9sBJO7n7n1St~nTcQEzPNZEiTUn6glK0Z9S2ENwAxylTLWxLLh6RGM8Pic3NPMwmKp5gZ7AdNGBcO67RGglv9mEw4iXv861RfIkVEOljLHZkfvVl-z8ALAf8sHKskxw__&Key-Pair-Id=K27TQMT39R1C8A"},{""`
)

var (
	nyPresentExpectedUrl = []string{"https://anthembcbsoh.mrf.bcbs.com/2024-08_800_72A0_in-network-rates_02_of_02.json.gz?&Expires=1726840845&Signature=uCkjd8P0JFuAfo6IvsW4ioKVeww2Y1H3rae0~7oXkIddycuC1aTN8IsLNGLK3uW4ROWCTG2twwFGLd~rEvkwG1pKGkTqDDVQtjHw8A0q~nxz14MsD7tz7XuHp~Cnvw4G2ugfrOytnqPVxq11Ai4iP60L7oSoDhKtSaUYT44YdbA1O-ZAIiWx~ypHx5VO4Fqc3V4uTzsJtVMYRQsfKffUq2p9sBJO7n7n1St~nTcQEzPNZEiTUn6glK0Z9S2ENwAxylTLWxLLh6RGM8Pic3NPMwmKp5gZ7AdNGBcO67RGglv9mEw4iXv861RfIkVEOljLHZkfvVl-z8ALAf8sHKskxw__&Key-Pair-Id=K27TQMT39R1C8A"}
	badUrlExpectedOutput = []string{"https:anthembcbsoh.mrf.bcbs.com/2024-08_890_58B0_in-network-rates_52_of_60.json.gz?&Expires=1726840845&Signature=S~DtzrtOlcU-XvMYeb2RQxjRfih~Fn9rdi3THOfDZud1JuxWY38ikovMlJzeA95sKtBn-a~26SFYXuFjR8~qUGMKZscNCGZNMc8Xd1gpb2~O0x7-gxkKHHv8fKuKtWP2HamgQxUgeloLq5y1SzlSFei-vHhouZx6mCHchNYdte-tD3PFqI1MKT~q7hY4hGgrE2P0ZM6-r15THyLRWI33fDnS6Wzg7VcZTQee5Wr~sr-oPC~0kDTcaDEVP2ItGtFX4u5rMOX6cdQThPuG6MYJHRvltFokZU~5-KctEzmE68PPU8HRJ2H5D2K~0BYyYpzreZ2J8Bw01Cs0kT-FeLE1Xg__&Key-Pair-Id=K27TQMT39R1C8A"}
)

func Test_FilterGzip(t *testing.T) {
	tests := map[string]struct {
		input          string
		expectedUrls   []string
		expectedErrors []string
	}{
		"NY_present_valid_list":     {input: nyPresentInput, expectedUrls: nyPresentExpectedUrl, expectedErrors: []string(nil)},
		"NY_not_present_valid_list": {input: nyNotPresentInput, expectedUrls: []string(nil), expectedErrors: []string(nil)},
		"NY_present_bad_url":        {input: badUrlInput, expectedUrls: []string(nil), expectedErrors: badUrlExpectedOutput},
	}

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(nyPresentInput))
		assert.NoError(t, err)
	}))
	defer svr.Close()
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		assert.NoError(t, err)
	}

	err = svr.Config.Serve(l)
	if err != nil {
		assert.NoError(t, err)
	}

	urlListBuilder := &UrlListBuilder{
		FileChunksLimit:      5,
		FileChunkMemoryLimit: 1922320936,
		FileUrl:              "http://localhost:8080",
		HttpClient:           svr.Client(),
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			urlList, failedToParseUrls, err := urlListBuilder.FilterGzip()
			assert.Error(t, err)
			assert.Equal(t, tc.expectedUrls, urlList)
			assert.Equal(t, tc.expectedErrors, failedToParseUrls)
		})
	}
}
