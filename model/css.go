package model

import "github.com/charmbracelet/bubbles/list"

func cssList() list.Model {
	items := []list.Item{
		item{title: "Tailwind", desc: "https://tailwindcss.com/"},
		item{title: "Bootstrap", desc: "https://getbootstrap.com/"},
		item{title: "Traditional", desc: "plain ol csss"},
	}

	l := list.New(items, list.NewDefaultDelegate(), defaultWidth, defaultHeight)
	l.SetFilteringEnabled(false)
	l.Title = "CSS"

	return l
}
