package main

import (
	"regexp"

	"github.com/kataras/iris/v12"
)

// Follow the steps below:
// $ go get -u github.com/go-bindata/go-bindata/...
//
// $ go-bindata -nomemcopy -prefix "../http2push/" ../http2push/assets/...
// # OR if the ./assets directory was inside this example foder:
// # go-bindata -nomemcopy ./assets/...
//
// $ go run .
// Physical files are not used, you can delete the "assets" folder and run the example.

var opts = iris.DirOptions{
	IndexName: "/index.html",
	PushTargetsRegexp: map[string]*regexp.Regexp{
		"/":              iris.MatchCommonAssets,
		"/app2/app2app3": iris.MatchCommonAssets,
	},
	Compress:   false,
	ShowList:   true,
	Asset:      Asset,
	AssetInfo:  AssetInfo,
	AssetNames: AssetNames,
}

func main() {
	app := iris.New()
	app.HandleDir("/public", "./assets", opts)

	app.Run(iris.TLS(":443", "../http2push/mycert.crt", "../http2push/mykey.key"))
}
