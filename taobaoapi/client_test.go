package taobaoapi

import (
	"encoding/json"
	"github.com/saodd/alog"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

var (
	client  *Client
	secrets *Secrets
)

type Shop struct {
	Session string  `json:"session"`
	Iids    []int64 `json:"iids"`
}

type Secrets struct {
	AppKey         string `json:"app_key"`
	AppSecret      string `json:"app_secret"`
	TaobaoEndpoint string `json:"taobao_endpoint"`

	TmallShop  Shop `json:"tmall_shop"`
	TaobaoShop Shop `json:"taobao_shop"`
}

func getSecrets() {
	secrets = new(Secrets)
	resp, err := http.Get("http://localhost:26666/secrets")
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if err := resp.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, secrets); err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	getSecrets()
	client = &Client{
		AppKey:            secrets.AppKey,
		AppSecret:         []byte(secrets.AppSecret),
		HttpDo:            http.DefaultClient.Do,
		DebugFlag:         true,
		TaobaoApiEndpoint: secrets.TaobaoEndpoint,
		HandleError:       alog.CE,
		DebugPrintf:       log.Printf,
	}
	os.Exit(m.Run())
}
