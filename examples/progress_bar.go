package main

import (
	"github.com/kodemore/lada/pkg"
	"github.com/kodemore/lada/pkg/ui"
	"time"
)

func main() {
	app, _ := lada.NewApplication("Progress Bar Application", "1.0.0")
	app.Description = "This application showcases how progress bar works"
	app.AddCommand("progressbar", func(t *lada.Terminal, _ lada.Arguments) error {

		total := 20
		updateProgressBar := make(chan int)
		pb := ui.NewProgressBar("My Progress", total, updateProgressBar)
		t.Display(pb)

		for i := 0; i <= total; i++ {
			updateProgressBar <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(updateProgressBar)

		t.Println("Success!")
		return nil
	})

	app.Run()
}

