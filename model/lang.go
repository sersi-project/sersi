package model

import "github.com/charmbracelet/bubbles/list"

func languageList() list.Model {

	items := []list.Item{
		item{title: "Typescript", desc: "https://www.typescriptlang.org/"},
		item{title: "Javascript", desc: "plain ol'javascript"},
	}

	l := list.New(items, list.NewDefaultDelegate(), defaultWidth, defaultHeight)
	l.SetFilteringEnabled(false)
	l.Title = "Language"

	return l
}
