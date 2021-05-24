package lada

import (
	"fmt"
	"math"
)

const pbBackground = "░"
const pbBar = "█"
const pbWidth = 30

type ProgressBarUI struct {
	Label   	string
	total 		int
	progress 	int
	width 		int
	update 		chan int
}

func NewProgressBarUI(label string, total int, update chan int) *ProgressBarUI {
	pb := &ProgressBarUI{
		Label:      label,
		total: total,
		progress: 0,
		width: pbWidth,
		update: update,
	}

	return pb
}

func (pb *ProgressBarUI) drawBar(t *Terminal) {
	done := int(math.Ceil(float64(pb.progress) / float64(pb.total) * float64(pb.width)))
	for i := 0; i < pb.width; i++ {
		if done >= i {
			t.PrettyPrint(pbBar, Foreground.Green)
			continue
		}
		t.PrettyPrint(pbBackground, Format.Dim)
	}
	if pb.progress == pb.total {
		t.Printf(" | %d/%d", pb.progress, pb.total)
	} else {
		t.PrettyPrint(fmt.Sprintf(" | %d/%d ", pb.progress, pb.total), Format.Dim)
	}
	t.Print("\n")
}

func (pb *ProgressBarUI) Display(t *Terminal) error {
	t.Cursor.Hide()
	t.DisableInput()
	t.PrettyPrint(pb.Label, Format.Dim)
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

func (pb *ProgressBarUI) Remove(t *Terminal) error {
	return nil
}

func (pb *ProgressBarUI) refresh(t *Terminal) {
	t.Cursor.EraseLine()
	pb.drawBar(t)
}
