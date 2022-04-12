package tabby

import "fmt"

type Tabby struct {
	Spacing     int
	ColumnCount int
	ColMax      []int
	Rows        [][]TabbyItem
}

type TabbyItem struct {
	Direction int
	Text      string
}

func Item(direction int, text string) TabbyItem {
	return TabbyItem{Direction: direction, Text: text}
}

func Left(text string) TabbyItem {
	return TabbyItem{Direction: 0, Text: text}
}

func Right(text string) TabbyItem {
	return TabbyItem{Direction: 1, Text: text}
}

func (t *Tabby) AddLine(args ...TabbyItem) {
	t.ColumnCount = len(args)
	t.ColMax = []int{}
	t.Rows = append(t.Rows, args)
	if len(t.ColMax) == 0 {
		for _, arg := range args {
			t.ColMax = append(t.ColMax, len(arg.Text))
		}
	}
	for idx, arg := range args {
		t.ColMax[idx] = max(len(arg.Text), t.ColMax[idx])
	}
}

func (t *Tabby) Print() {
	for _, row := range t.Rows {
		for c, column := range row {
			spacing := t.ColMax[c]
			gap := spacing - len(column.Text)
			if column.Direction == 0 {
				fmt.Printf(column.Text)
				printSpace(gap)
				printSpace(t.Spacing)
			} else {
				printSpace(gap)
				fmt.Printf(column.Text)
				printSpace(t.Spacing)
			}
		}
		fmt.Println()
	}
}

func printSpace(count int) {
	for i := 0; i < count; i++ {
		fmt.Printf(" ")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func New(space int) Tabby {
	return Tabby{Spacing: space}
}
