package ui

import (
	"fmt"
	"github.com/kodemore/lada/pkg"
	"github.com/kodemore/lada/pkg/style"
	"math"
)

const pbBackground = "░"
const pbBar = "█"
const pbWidth = 30

type ProgressBar struct {
	Label   	string
	total 		int
	progress 	int
	width 		int
	update 		chan int
}

func NewProgressBar(label string, total int, update chan int) *ProgressBar {
	pb := &ProgressBar{
		Label:    label,
		total:    total,
		progress: 0,
		width:    pbWidth,
		update:   update,
	}

	return pb
}

func (pb *ProgressBar) drawBar(t *lada.Terminal) {
	done := int(math.Ceil(float64(pb.progress) / float64(pb.total) * float64(pb.width)))
	for i := 0; i < pb.width; i++ {
		if done >= i {
			t.PrettyPrint(pbBar, style.Foreground.Green)
			continue
		}
		t.PrettyPrint(pbBackground, style.Format.Dim)
	}
	if pb.progress == pb.total {
		t.Printf(" | %d/%d", pb.progress, pb.total)
	} else {
		t.PrettyPrint(fmt.Sprintf(" | %d/%d ", pb.progress, pb.total), style.Format.Dim)
	}
	t.Print("\n")
}

func (pb *ProgressBar) Display(t *lada.Terminal) error {
	t.Cursor.Hide()
	t.DisableInput()
	t.PrettyPrint(pb.Label, style.Format.Dim)
	t.Print("\n")
	pb.drawBar(t)
	go func() {
		for {
			if update, ok := <- pb.update; ok {
				pb.progress = update
			}
			if pb.progress > pb.total {
				pb.progress = pb.total
			}
			if pb.progress == pb.total {
				break
			}
			pb.refresh(t)
		}
		pb.refresh(t)
		t.RestoreDefaultMode()
		t.Cursor.Show()
	}()
	return nil
}

func (pb *ProgressBar) Remove(t *lada.Terminal) error {
	return nil
}

func (pb *ProgressBar) refresh(t *lada.Terminal) {
	t.Cursor.EraseLine()
	pb.drawBar(t)
}
