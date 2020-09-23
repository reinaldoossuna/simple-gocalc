package display

import (
	"syscall"
	"unsafe"
	"fmt"
)

type winsize struct {
	Row uint16
	Col uint16
	Xpixel uint16
	Ypixel uint16
}

func getSizeTerminal() *winsize {
	ws := &winsize{}
    retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
        uintptr(syscall.Stdin),
        uintptr(syscall.TIOCGWINSZ),
        uintptr(unsafe.Pointer(ws)))

    if int(retCode) == -1 {
        panic(errno)
    }
    return ws
}

func GoToend() {
	ws := getSizeTerminal()
	MoveCursorTo(ws.Row, ws.Col)

}

func ClearLine(row, col uint16) {
	MoveCursorTo(row, col)
	fmt.Print("\033[K")
}

func ClearTerm() {
	fmt.Print("\033[H\033[2J")
}

func MoveCursorTo(row, col uint16) {
	fmt.Printf("\033[%v;%vH", row, col)
}

func GetInput() string {
	var s string
	fmt.Scanf("%s", &s)
	return s
}
