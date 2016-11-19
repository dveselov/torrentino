package main

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
)

var (
	TorrentinoHTTPS, _ = strconv.ParseBool(os.Getenv("TORRENTINO_HTTPS"))
	TorrentinoHostname = os.Getenv("TORRENTINO_HOSTNAME")

	TokenSigningKey    = []byte(os.Getenv("TORRENTINO_TOKEN_SIGNING_KEY"))
	TokenSigningMethod = jwt.SigningMethodHS256
	TokenTTL, _        = strconv.Atoi(os.Getenv("TORRENTINO_TOKEN_TTL"))
	TokenIssuer        = "torrentino"

	DownloadDirectory = os.Getenv("TORRENTINO_DOWNLOAD_DIRECTORY")
	MaxConnections, _ = strconv.Atoi(os.Getenv("TORRENTINO_MAX_CONNECTIONS"))
)
