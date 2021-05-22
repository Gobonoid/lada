package lada

import (
	"bufio"
	"fmt"
	"golang.org/x/sys/unix"
	"io"
	"os"
	"syscall"
)

const (
	ioctlReadTermios  = syscall.TIOCGETA
	ioctlWriteTermios = syscall.TIOCSETA
	ioctlGetWindowSize = syscall.TIOCGWINSZ
)

type TerminalMode int

const (
	TerminalDefaultMode TerminalMode = iota
	TerminalDisabledInputMode
)

const osTtyPath = "/dev/tty"

type TerminalWindowSize struct {
	Rows int
	Columns int
}

type Terminal struct {
	bstdin        *bufio.Reader
	bstdout       *bufio.Writer
	stdin         io.Reader
	stdout        io.Writer
	ttyPath       string
	ttyFd         uintptr
	stashedMode   *unix.Termios
	mode          TerminalMode
	Cursor        *Cursor
}

func NewTerminal() (*Terminal, error) {
	term := &Terminal{
		ttyPath: osTtyPath,
		mode: TerminalDefaultMode,
	}
	stdin, err := os.OpenFile(term.ttyPath, os.O_RDONLY, 0)
	if err != nil {
		return &Terminal{}, err
	}
	term.stdin = stdin
	term.ttyFd = stdin.Fd()

	term.stdout, err = os.OpenFile(term.ttyPath, os.O_WRONLY, 0)
	if err != nil {
		return &Terminal{}, err
	}

	term.stashedMode, err = term.readTermios()
	if err != nil {
		return &Terminal{}, err
	}

	term.bstdin = bufio.NewReader(term.stdin)
	term.bstdout = bufio.NewWriter(term.stdout)
	term.Cursor, err = NewCursor(term.stdout)
	if err != nil {
		return &Terminal{}, err
	}

	return term, nil
}

func (t *Terminal) writeTermios(termios *unix.Termios) error {
	return unix.IoctlSetTermios(int(t.ttyFd), ioctlWriteTermios, termios)
}

func (t *Terminal) readTermios() (*unix.Termios, error) {
	return unix.IoctlGetTermios(int(t.ttyFd), ioctlReadTermios)
}

func (t *Terminal) CaptureKeys(o interface{OnKey(t *Terminal, k Key) bool}) error {
	t.Cursor.Hide()
	termios := *t.stashedMode
	termios.Lflag &^= unix.ECHO | unix.ECHONL | unix.ICANON
	err := t.writeTermios(&termios)
	if err != nil {
		return err
	}
	for {
		b := make(Key, 4)
		_, err := t.bstdin.Read(b)
		if err != nil {
			return err
		}

		if !o.OnKey(t, b) {
			break
		}
	}
	t.RestoreDefaultMode()
	t.Cursor.Show()
	return nil
}

func (t *Terminal) DisableInput() error {
	err := t.RestoreDefaultMode()
	if err != nil {
		return err
	}

	termios := *t.stashedMode
	termios.Lflag &^= unix.ECHO | unix.ICANON
	err = t.writeTermios(&termios)
	if err != nil {
		return err
	}

	t.mode = TerminalDisabledInputMode

	return nil
}

func (t *Terminal) Prompt(s string) (string, error) {
	err := t.Print(s)
	if err != nil {
		return "", err
	}

	line, _, err := t.bstdin.ReadLine()
	if err != nil {
		return "", err
	}

	return string(line), nil
}

func (t *Terminal) Secret(s string) (string, error) {
	t.Cursor.Hide()
	err := t.Print(s)
	if err != nil {
		return "", err
	}
	t.DisableInput()
	if err != nil {
		return "", err
	}
	line, _, err := t.bstdin.ReadLine()
	if err != nil {
		return "", err
	}
	t.RestoreDefaultMode()

	t.Print("\n")
	t.Cursor.Show()
	return string(line), nil
}

func (t *Terminal) Display(ui UIElement) error {
	ui.Display(t)
	if keyboardListener, ok := ui.(UIKeyboardListener); ok {
		err := t.CaptureKeys(keyboardListener)
		if err != nil {
			return err
		}
	}

	if removableUi, ok := ui.(UIRemovable); ok {
		err := removableUi.Remove(t)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Terminal) GetWindowSize() (*TerminalWindowSize, error) {
	size, err := unix.IoctlGetWinsize(int(t.ttyFd), ioctlGetWindowSize)
	if err != nil {
		return &TerminalWindowSize{}, err
	}

	return &TerminalWindowSize{
		Rows: int(size.Row),
		Columns: int(size.Col),
	}, nil
}

func (t *Terminal) Print(s string) error {
	_, err := fmt.Fprint(t.stdout, s)
	if err != nil {
		return err
	}
	return nil
}

func (t *Terminal) Println(s string) error {
	_, err := fmt.Fprintln(t.stdout, s)
	if err != nil {
		return err
	}
	return nil
}

func (t *Terminal) Printf(format string, params ...interface{}) error {
	_, err := fmt.Fprintf(t.stdout, format, params...)
	if err != nil {
		return err
	}
	return nil
}

func (t *Terminal) PrettyPrint(s string, style ...Style) error {
	err := t.Cursor.SetStyle(style...)
	if err != nil {
		return err
	}

	err = t.Print(s)
	if err != nil {
		return err
	}

	err = t.Cursor.ResetStyle()
	if err != nil {
		return err
	}

	return nil
}

func (t *Terminal) RestoreDefaultMode() error {
	err := t.writeTermios(t.stashedMode)
	if err != nil {
		return err
	}

	t.mode = TerminalDefaultMode
	return nil
}

func (t *Terminal) close() {
	t.RestoreDefaultMode()
	t.Cursor.close()
}