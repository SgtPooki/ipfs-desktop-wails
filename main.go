package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	// StartIpfs() // even commenting this out, getting errors from ipfs.go:
	// ../../../../../.asdf/installs/golang/1.21.1/packages/pkg/mod/github.com/ipfs/kubo@v0.22.0/thirdparty/verifbs/verifbs.go:17:33: not enough arguments in call to verifcid.ValidateCid
	// ../../../../../.asdf/installs/golang/1.21.1/packages/pkg/mod/github.com/libp2p/go-libp2p@v0.29.0/p2p/transport/quicreuse/quic_multiaddr.go:38:19: undefined: quic.VersionDraft29

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "ipfs-desktop",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err)
	}
}
