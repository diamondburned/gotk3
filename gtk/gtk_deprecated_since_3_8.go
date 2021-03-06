// Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
//
// This file originated from: http://opensource.conformal.com/
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// This file includes wrapers for symbols deprecated beginning with GTK 3.8,
// and should only be included in a build targeted intended to target GTK
// 3.6 or earlier.  To target an earlier build build, use the build tag
// gtk_MAJOR_MINOR.  For example, to target GTK 3.6, run
// 'go build -tags gtk_3_6'.
// +build gtk_3_6

package gtk

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

// SetOpacity is a wrapper around gtk_window_set_opacity()
// Deprecated since 3.8, replaced by SetOpacity on Widget (gtk_widget_set_opacity)
func (v *Window) SetOpacity(opacity float64) {
	C.gtk_window_set_opacity(v.native(), C.gdouble(opacity))
}

// GetOpacity is a wrapper around gtk_window_get_opacity()
// Deprecated since 3.8, replaced by GetOpacity on Widget (gtk_widget_get_opacity)
func (v *Window) GetOpacity() float64 {
	c := C.gtk_window_get_opacity(v.native())
	return float64(c)
}
