package main

import (
	"fmt"
	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/storage"
	"github.com/dchest/uniuri"
	"io/ioutil"
	"os"
)

type Client struct {
	Client  *torrent.Client
	Torrent *torrent.Torrent
	Config  ClientConfig
}

type ClientConfig struct {
	URL            string
	Seed           bool
	Port           int
	TCP            bool
	DHT            bool
	MaxConnections int
	DownloadPath   string
}

func NewClientConfig(URL string) ClientConfig {
	directory, _ := ioutil.TempDir(DownloadDirectory, uniuri.New())
	return ClientConfig{
		URL:            URL,
		Seed:           false,
		TCP:            true,
		DHT:            false,
		MaxConnections: MaxConnections,
		DownloadPath:   directory,
	}
}

func NewClient(config ClientConfig) (client Client, err error) {
	var (
		t *torrent.Torrent
		c *torrent.Client
	)

	c, err = torrent.NewClient(&torrent.Config{
		DefaultStorage: storage.NewBoltDB(config.DownloadPath),
		Seed:           config.Seed,
		NoUpload:       !config.Seed,
		DisableTCP:     !config.TCP,
		NoDHT:          !config.DHT,
		ListenAddr:     fmt.Sprintf(":%d", config.Port),
	})
	if err != nil {
		return Client{}, err
	}

	t, err = c.AddMagnet(config.URL)
	if err != nil {
		return Client{}, err
	}
	t.SetMaxEstablishedConns(config.MaxConnections)
	return Client{
		Client:  c,
		Torrent: t,
		Config:  config,
	}, nil
}

func (c *Client) Close() {
	c.Torrent.Drop()
	c.Client.Close()
	os.RemoveAll(c.Config.DownloadPath)
}
