package xelasticsearch

import "github.com/elastic/go-elasticsearch/v8"

type Config struct {
	Addresses []string `json:"addresses" mapsturcture:"addresses"`
	Username  string   `json:"username" mapsturcture:"username"`
	Password  string   `json:"password" mapsturcture:"password"`
}

type Client struct {
	client *elasticsearch.Client
}

func New(c *Config) (*Client, error) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: c.Addresses,
		Username:  c.Username,
		Password:  c.Password, // todo 密码使用enc加密
	})
	if err != nil {
		return nil, err
	}
	return &Client{
		client: client,
	}, nil
}

func (c *Client) Client() *elasticsearch.Client {
	return c.client
}
