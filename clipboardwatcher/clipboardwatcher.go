package main

import (
	"encoding/json"
	"os"
	"syscall"
	"unsafe"

	"github.com/AllenDang/w32"
)

func main() {
	className := syscall.StringToUTF16Ptr("clipboard_watcher")
	wndClassEx := w32.WNDCLASSEX{
		ClassName: className,
		WndProc: syscall.NewCallback(func(hwnd w32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
			if msg == w32.WM_CLIPBOARDUPDATE {
				if !w32.OpenClipboard(0) {
					return 0
				}
				defer w32.CloseClipboard()
				h := w32.HGLOBAL(w32.GetClipboardData(w32.CF_UNICODETEXT))
				if h == 0 {
					return 0
				}
				defer w32.GlobalUnlock(h)

				text := w32.UTF16PtrToString((*uint16)(w32.GlobalLock(h)))
				json.NewEncoder(os.Stdout).Encode(text)
				return 0
			}
			return w32.DefWindowProc(hwnd, msg, wParam, lParam)
		}),
	}
	wndClassEx.Size = uint32(unsafe.Sizeof(wndClassEx))
	w32.RegisterClassEx(&wndClassEx)

	hwnd := w32.CreateWindowEx(0, className, className, 0, 0, 0, 0, 0, w32.HWND_MESSAGE, 0, 0, nil)
	w32.AddClipboardFormatListener(hwnd)
	defer w32.RemoveClipboardFormatListener(hwnd)

	msg := w32.MSG{}
	for w32.GetMessage(&msg, 0, 0, 0) > 0 {
		w32.DispatchMessage(&msg)
	}
}
