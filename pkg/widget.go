package lada

type Widget struct {
}

func (w *Widget) IsDrawn() bool {
	return false
}

func (w *Widget) Redraw() error {
	return nil
}
