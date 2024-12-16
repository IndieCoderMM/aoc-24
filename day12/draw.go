package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	colors  map[string]lipgloss.Style
	grid    [][]string
	visited [][]bool
	ans     int
	running bool
	stepFn  func() (int, bool)
}

func initialModel(grid [][]string, visited [][]bool, stepFn func() (int, bool)) tea.Model {
	m := Model{
		getColorMap(),
		grid,
		visited,
		0,
		false,
		stepFn,
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return m.UpdateGrid()
}

type stepMsg struct{}

func (m Model) UpdateGrid() tea.Cmd {
	return tea.Batch(tea.Tick(time.Millisecond*500, func(time.Time) tea.Msg {
		return stepMsg{}
	}))
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case stepMsg:
		if m.running {
			m.RunAlgorithm()
			return m, m.UpdateGrid()
		}
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	// only color visited cell
	defaultCell := lipgloss.NewStyle().Background(lipgloss.Color("#3a3a3a")).Foreground(lipgloss.Color("#f9f6f2"))
	titleCell := lipgloss.NewStyle().Foreground(lipgloss.Color("#f9f6f2"))

	s := ""
	// Render empty space above the grid
	for x := 0; x < len(m.grid[0]); x++ {
		s += titleCell.Render("   ")
	}
	s += "\n"

	title := ""
	if !m.running {
		title = fmt.Sprintf("Total price: $%4d", m.ans)
	} else {
		title = fmt.Sprintf("Calculating: $%4d", m.ans)
	}

	s += writeText(title, titleCell, 3*len(m.grid[0]))

	// Render empty space above the grid
	for x := 0; x < len(m.grid[0]); x++ {
		s += titleCell.Render("   ")
	}

	s += "\n"
	for y := 0; y < len(m.grid); y++ {
		for x := 0; x < len(m.grid[y]); x++ {
			txt := m.grid[y][x]
			if m.visited[y][x] {
				s += m.colors[m.grid[y][x]].Render(" ")
			} else {
				s += defaultCell.Render(" ")
			}
			if m.visited[y][x] {
				s += m.colors[m.grid[y][x]].Render(txt + " ")
			} else {

				s += defaultCell.Render(txt + " ")
			}
		}
		// s += "\n"
		// for x := 0; x < len(m.grid[y]); x++ {
		// 	if m.visited[y][x] {
		// 		s += m.colors[m.grid[y][x]].Render("   ")
		// 	} else {
		// 		s += defaultStyle.Render("   ")
		// 	}
		// }
		s += "\n"
	}

	s += "\nPress Q to quit\n"

	return s
}

func writeText(title string, titleCell lipgloss.Style, width int) string {
	s := ""

	// Calculate padding for centering the title
	padding := (width - len(title)) / 2
	extra := (width - len(title)) % 2

	// Left padding
	for x := 0; x < padding-extra; x += 3 {
		s += titleCell.Render("   ")
	}

	// Render the title
	s += titleCell.Render(title)

	// Right padding (to ensure the title is centered)
	for x := 0; x < width-len(title)-padding-extra; x += 3 {
		s += titleCell.Render("   ")
	}
	s += "\n"

	return s
}

func Run(grid [][]string, visited [][]bool, stepFn func() (int, bool)) {
	p := tea.NewProgram(initialModel(grid, visited, stepFn))

	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

func (m *Model) RunAlgorithm() {
	if m.running {
		ans, completed := m.stepFn()
		m.ans = ans
		m.running = !completed
	}
}

func getColorMap() map[string]lipgloss.Style {
	defaultStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#f9f6f2"))
	return map[string]lipgloss.Style{
		"A": defaultStyle.Background(lipgloss.Color("#f94144")), // Red
		"B": defaultStyle.Background(lipgloss.Color("#f3722c")), // Orange
		"C": defaultStyle.Background(lipgloss.Color("#f8961e")), // Amber
		"D": defaultStyle.Background(lipgloss.Color("#f9c74f")), // Yellow
		"E": defaultStyle.Background(lipgloss.Color("#90be6d")), // Green
		"F": defaultStyle.Background(lipgloss.Color("#43aa8b")), // Teal
		"G": defaultStyle.Background(lipgloss.Color("#577590")), // Navy
		"H": defaultStyle.Background(lipgloss.Color("#7d4e57")), // Wine
		"I": defaultStyle.Background(lipgloss.Color("#2a9d8f")), // Turquoise
		"J": defaultStyle.Background(lipgloss.Color("#264653")), // Dark Green
		"K": defaultStyle.Background(lipgloss.Color("#a8dadc")), // Soft Cyan
		"L": defaultStyle.Background(lipgloss.Color("#457b9d")), // Blue
		"M": defaultStyle.Background(lipgloss.Color("#1d3557")), // Deep Blue
		"N": defaultStyle.Background(lipgloss.Color("#6d597a")), // Purple
		"O": defaultStyle.Background(lipgloss.Color("#b5838d")), // Blush Pink
		"P": defaultStyle.Background(lipgloss.Color("#ffb4a2")), // Coral Pink
		"Q": defaultStyle.Background(lipgloss.Color("#e5989b")), // Rose
		"R": defaultStyle.Background(lipgloss.Color("#d672d6")), // Lavender
		"S": defaultStyle.Background(lipgloss.Color("#9d4edd")), // Violet
		"T": defaultStyle.Background(lipgloss.Color("#7209b7")), // Deep Violet
		"U": defaultStyle.Background(lipgloss.Color("#4a148c")), // Purple Night
		"V": defaultStyle.Background(lipgloss.Color("#d00000")), // Crimson Red
		"W": defaultStyle.Background(lipgloss.Color("#dc2f02")), // Burnt Orange
		"X": defaultStyle.Background(lipgloss.Color("#e85d04")), // Fire Orange
		"Y": defaultStyle.Background(lipgloss.Color("#faa307")), // Gold
		"Z": defaultStyle.Background(lipgloss.Color("#ffba08")), // Sun Yellow
	}
}
