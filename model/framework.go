package model

import (
	"github.com/charmbracelet/bubbles/list"
)

func frameworkList() list.Model {
	items := []list.Item{
		item{title: "React", desc: "https://react.dev/"},
		item{title: "Svelte", desc: "https://svelte.dev/"},
		item{title: "Vue", desc: "https://vuejs.org/"},
		item{title: "Vanilla", desc: "https://vanilla-js.com/"},
	}

	l := list.New(items, list.NewDefaultDelegate(), defaultWidth, defaultHeight)
	l.SetFilteringEnabled(false)
	l.Title = "Framework"

	return l
}
