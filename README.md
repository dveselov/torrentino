# Torrentino

BitTorrent -> HTTP gateway

# API

Each request to `torrentino` currently done via JWT tokens with additional fields:

* URL - magnet-link of torrent
* Path - optional, if blank `torrentino` will return list of files in torrent, else return streaming response with file content 

```bash
$ curl -i http://localhost/{ jwt token with `path` field equal to '' }/
Content-Type: application/json; charset=utf-8
Date: Sat, 19 Nov 2016 11:00:00 GMT
Transfer-Encoding: chunked

[
    {
        "length": 553111,
        "path": "Motörhead (mp3)/01 - Studio albums/1977 - Motörhead/Artwork/Back.jpg",
        "mimetype": "image/jpeg",
        "url": "%full url to this file%"
    },
    {
        "length": 600077,
        "path": "Motörhead (mp3)/01 - Studio albums/1977 - Motörhead/Artwork/Front.jpg",
        "mimetype": "image/jpeg",
        "url": "%full url to this file%"
    },
]

$ curl -i http://hostname:port/{ jwt token with `path` == 'some_song.mp3' }/
Content-Length: 7771843
Content-Type: audio/mpeg
Date: Sat, 19 Nov 2016 11:01:37 GMT

... binary data ...

```

# Settings

```bash
$ export TORRENTINO_TOKEN_SIGNING_KEY=secret
$ export TORRENTINO_TOKEN_TTL=600
$ export TORRENTINO_DOWNLOAD_DIRECTORY=/tmp/ # files'll be removed as download completes
$ export TORRENTINO_MAX_CONNECTIONS=50 # don't set too much, or extend limits via ulimit
$ export TORRENTINO_HOSTNAME=localhost # used when creating links for files in torrents
$ export TORRENTINO_HTTPS=true # and this too
```
