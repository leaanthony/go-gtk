//go:build ignore
// +build ignore

package gtk

import (
	"unsafe"

	"github.com/leaanthony/go-gtk/gdk"
)

func (v *Window) XID() int32 {
	return gdk.WindowFromUnsafe(unsafe.Pointer(v.GWidget.window)).GetNativeWindowID()
}
