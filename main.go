package main

import (
	lada "github.com/kodemore/lada/pkg"
	"os"
)

func main() {
	cursor, _ := lada.NewCursor(os.Stdout)

	cursor.Print("\n asda \n")
	cursor.Print("Move to next line")

	cursor.MoveUp(1)
	cursor.MoveDown(2)
	cursor.Print("11")
	cursor.MoveDown(3)
	cursor.SetStyle(
		lada.NewCursorStyle(
			lada.Colors.YELLOW,
			lada.Colors.LIGHTRED,
			lada.Styles.BOLD,
			lada.Styles.STRIKE,
		))
	cursor.Print("restored position")
	cursor.ResetStyle()

	cursor.Close()
}
