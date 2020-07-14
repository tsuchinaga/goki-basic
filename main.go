package main

import (
	"github.com/goki/gi/gi"
	"github.com/goki/gi/gimain"
)

func main() {
	gimain.Main(func() {
		width := 1024
		height := 768
		win := gi.NewMainWindow("gogi-basic", "Basic Test Windows", width, height)

		vp := win.WinViewport2D()
		updt := vp.UpdateStart()

		mfr := win.SetMainFrame()

		rlay := gi.AddNewLayout(mfr, "row-lay", gi.LayoutHoriz)
		rlay.SetProp("text-align", "center")
		gi.AddNewLabel(rlay, "label-1", "This is test text")

		// main menu
		appnm := gi.AppName()
		mmen := win.MainMenu
		mmen.ConfigMenus([]string{appnm, "Edit", "Window"})

		amen := win.MainMenu.ChildByName(appnm, 0).(*gi.Action)
		amen.Menu = make(gi.Menu, 0, 10)
		amen.Menu.AddAppMenu(win)

		emen := win.MainMenu.ChildByName("Edit", 1).(*gi.Action)
		emen.Menu = make(gi.Menu, 0, 10)
		emen.Menu.AddCopyCutPaste(win)

		win.SetCloseCleanFunc(func(*gi.Window) {
			go gi.Quit()
		})

		win.MainMenuUpdated()
		vp.UpdateEndNoSig(updt)
		win.StartEventLoop()
	})
}
