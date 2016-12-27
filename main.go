package main

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetTorrentByToken(c echo.Context) error {
	tokenString := c.Param("token")

	claims, err := ParseToken(tokenString)
	if err != nil {
		return c.String(http.StatusForbidden, "Invalid or expired token")
	}

	config := NewClientConfig(claims.URL)
	client, err := NewClient(config)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid or unavaible torrent")
	}
	defer client.Close()

	select {
	case <-client.Torrent.GotInfo():
	case <-time.After(time.Second * 60):
		return c.String(http.StatusServiceUnavailable, "Failed to get info about torrent (timeout)")
	}

	files := client.Torrent.Files()
	if claims.Path != "" {
		log.Println(claims.Path)
		for i := 0; i < len(files); i++ {
			file := files[i]
			if file.Path() == claims.Path {
				fileLength := int(file.Length())
				file.PrioritizeRegion(file.Offset(), int64(fileLength/100*5))

				reader := client.Torrent.NewReader()
				reader.Seek(file.Offset(), 0)
				defer reader.Close()

				c.Response().Header().Set(echo.HeaderContentLength, strconv.Itoa(fileLength))
				return c.Stream(http.StatusOK, GetMimeTypeByPath(file.DisplayPath()), reader)
			}
		}
		return c.String(http.StatusNotFound, "Invalid file request")
	}

	json := SerializeTorrentFiles(&client)

	return c.JSON(http.StatusOK, json)
}

func main() {
	e := echo.New()
	e.GET("/:token/", GetTorrentByToken)
	e.Logger.Fatal(e.Start(":1323"))
}
