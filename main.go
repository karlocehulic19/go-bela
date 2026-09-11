package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"slices"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/exp/charmtone"
)

var possibleCardValues = []string{"A", "K", "Q", "J", "X", "XI", "VIII", "VII"}
var possibleCardColors = []string{"Heart", "Spade", "Diamond", "Club"}

type card struct {
	color string
	value string
}

type player struct {
	cards []card
}

var inputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(charmtone.Cherry.Hex()))

func main() {
	p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}

type model struct {
	players [4]player
}

func initialModel() model { 
	cards := []card{}
	for _, color := range possibleCardColors {
		for _, value := range possibleCardValues {
			cards = append(cards, card{
				color: color,
				value: value,
			})
		}
	}

	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	players := [4]player{}
	for j := range 4 {
		players[j] = player{
			cards: cards[j * 8:(j + 1) * 8],
		}
	}

	return model{
		players: players,
	}
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    case tea.KeyPressMsg:
        switch msg.String() {

        case "ctrl+c", "q":
            return m, tea.Quit
						}
					}

    return m, nil
}

func get_card_string(card_type string) (string, error) {
	if !slices.Contains(possibleCardValues, card_type) {
		return "", errors.New("An invalid card type is specified for card string")
	}

	top_named_line := ""
	bottom_named_line := ""
	switch card_type {
	case "XI":
		top_named_line = "|%s    |\n"
		bottom_named_line = "|    %s|\n"
	case "VIII":
		top_named_line = "|%s  |\n"
		bottom_named_line = "|  %s|\n"
	case "VII":
		top_named_line = "|%s   |\n"
		bottom_named_line = "|   %s|\n"
	default:
		top_named_line = "| %s    |\n"
		bottom_named_line = "|    %s |\n"
	}

	s := fmt.Sprintf(
		"/------\\\n" +
		top_named_line +
		"|      |\n" +
		"|      |\n" +
		bottom_named_line +
	 "\\------/\n", card_type, card_type)

	return s, nil;
}

func (m model) View() tea.View {
		card_s, err := get_card_string("VII")
		if err != nil {
			log.Fatal(err)
		}
		card_s_2, err_2 := get_card_string("X")
		if err_2 != nil {
			log.Fatal(err)
		}

		card_s_3, err_3 := get_card_string("A")
		if err_3 != nil {
			log.Fatal(err)
		}
		// Create some layers.
		a := lipgloss.NewLayer(inputStyle.Render(card_s)).X(0)
		b := lipgloss.NewLayer(inputStyle.Render(card_s_2)).X(6)
		c := lipgloss.NewLayer(inputStyle.Render(card_s_3)).X(12)
		layers := []*lipgloss.Layer {a, b, c}

		// Composite 'em and render.
		outter_layer := lipgloss.NewLayer(inputStyle.Width(30).String())
		compositor := lipgloss.NewCompositor(outter_layer.AddLayers(layers...));
		output := compositor.Render()
    return tea.NewView(output)
}
