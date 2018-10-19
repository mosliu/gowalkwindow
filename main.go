package main

import "github.com/mosliu/gowalkwindow/ui"

// 使用icon？
//Tool for embedding .ico & manifest resources in Go programs for Windows.
//https://github.com/akavel/rsrc
//go:generate echo Hello, Go Generate assets & syso file
//go:generate go-bindata -o=assets.go -pkg=main assets/...
//go:generate rsrc -manifest main.manifest -o gowalkwindow.syso -ico .\assets\icons\setport.ico

func main() {
	//test()

	ui.ShowMainApp()
}
