package menuinput

import (
	"fmt"
	"io"

	"strings"

	"github.com/sersi-project/core/tea/styles"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	listHeight   = 14
	defaultWidth = 20
)

type errMsg error

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(4).Foreground(lipgloss.Color("#cd24cd"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	//nolint:errcheck,gocritic
	str := fmt.Sprintf("%s", i)

	//nolint:errcheck,gocritic
	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("[✓]  " + strings.Join(s, " "))
		}
	} else {
		fn = func(s ...string) string {
			return itemStyle.Render("[ ]  " + strings.Join(s, " "))
		}
	}

	if _, err := fmt.Fprint(w, fn(str)); err != nil {
		return
	}
}

type ListModel struct {
	list     list.Model
	listType string
	Choice   string
	header   string
	quitting bool
	err      error
	styles   *styles.Styles
}

func (m *ListModel) Init() tea.Cmd {
	return nil
}

func InitialMenuInput(header string, itemsString []string, listType string) *ListModel {
	items := make([]item, len(itemsString))
	for i, it := range itemsString {
		items[i] = item(it)
	}

	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = list.Item(item)
	}

	l := list.New(listItems, itemDelegate{}, defaultWidth, listHeight)
	styles := styles.DefaultStyles()
	l.Title = styles.Header.Render(header)
	l.SetShowPagination(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return &ListModel{
		list:     l,
		listType: listType,
		Choice:   "",
		header:   header,
		quitting: false,
		styles:   styles,
	}
}

func (m *ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.Choice = string(i)
			}
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *ListModel) View() string {
	if m.Choice != "" {
		return fmt.Sprintf("> %-15s: %s\n", m.styles.SelectionTitle.Render(m.listType), m.styles.Selection.Render(m.Choice))
	}
	if m.quitting {
		return fmt.Sprintf("\n%s\n", m.styles.Cancel.Render("Operation cancelled!"))
	}
	return "\n" + m.list.View()
}
