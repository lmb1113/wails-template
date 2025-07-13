package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

const lockUniqueId = "C04B4B88-2B22-3ADE-9DFE-551EBA431119"

func main() {
	// Create an instance of the app structure
	app := NewApp()
	// Create application with options
	err := wails.Run(&options.App{
		Title:     "wails-template",
		Width:     1280,
		Height:    800,
		MinWidth:  1280,
		MinHeight: 800,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 100},
		OnStartup:        app.startup,
		WindowStartState: options.Normal,
		LogLevel:         logger.DEBUG,
		Bind: []interface{}{
			app,
		},
		DisableResize: true,
		Frameless:     true,
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId:               lockUniqueId,               // 单实例锁
			OnSecondInstanceLaunch: app.onSecondInstanceLaunch, // 如果重复打开就再次显示窗口
		},
		OnShutdown: nil,
		Windows: &windows.Options{
			CustomTheme: &windows.ThemeSettings{
				DarkModeTitleBar:           0xffffff,
				DarkModeTitleBarInactive:   0xffffff,
				DarkModeTitleText:          0xffffff,
				DarkModeTitleTextInactive:  0xffffff,
				DarkModeBorder:             0xffffff,
				DarkModeBorderInactive:     0xffffff,
				LightModeTitleBar:          0xffffff,
				LightModeTitleBarInactive:  0xffffff,
				LightModeTitleText:         0xffffff,
				LightModeTitleTextInactive: 0xffffff,
				LightModeBorder:            0xffffff,
				LightModeBorderInactive:    0xffffff,
			},
			DisableFramelessWindowDecorations: true,
			Theme:                             windows.Light,
		},
		Linux: &linux.Options{
			Icon: icon,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  true,
				HideTitleBar:               true,
				FullSizeContent:            true,
				UseToolbar:                 false,
				HideToolbarSeparator:       false,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "wails-template",
				Message: "wails-template",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

// onSecondInstanceLaunch 唯一进程
func (a *App) onSecondInstanceLaunch(secondInstanceData options.SecondInstanceData) {
	runtime.WindowUnminimise(a.ctx)
	runtime.Show(a.ctx)
}
