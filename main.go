package main

import (
	"github.com/linweiyuan/go-chatgpt/api"
	"github.com/linweiyuan/go-chatgpt/common"
	"github.com/linweiyuan/go-chatgpt/ui"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	modal := tview.NewModal().SetText(common.ChooseModeTitle).
		AddButtons([]string{common.ApiMode, common.ChatGPTMode}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		switch buttonLabel {
		case common.ApiMode:
			common.IsChatGPT = false
		case common.ChatGPTMode:
			common.IsChatGPT = true
		}
		setup(app)
	})

	if err := app.SetRoot(modal, true).SetFocus(modal).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func setup(app *tview.Application) {
	tui := ui.New(api.New(), app)
	tui.Setup()
	if common.IsChatGPT {
		go tui.GetConversations()
	}
}
