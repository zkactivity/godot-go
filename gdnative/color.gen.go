package gdnative

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "types.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
#include "gdnative.gen.h"
#include <gdnative/color.h>
// Include all headers for now. TODO: Look up all the required
// headers we need to import based on the method arguments and return types.
#include <gdnative/aabb.h>
#include <gdnative/array.h>
#include <gdnative/basis.h>
#include <gdnative/color.h>
#include <gdnative/dictionary.h>
#include <gdnative/gdnative.h>
#include <gdnative/node_path.h>
#include <gdnative/plane.h>
#include <gdnative/pool_arrays.h>
#include <gdnative/quat.h>
#include <gdnative/rect2.h>
#include <gdnative/rid.h>
#include <gdnative/string.h>
#include <gdnative/string_name.h>
#include <gdnative/transform.h>
#include <gdnative/transform2d.h>
#include <gdnative/variant.h>
#include <gdnative/vector2.h>
#include <gdnative/vector3.h>
#include <gdnative_api_struct.gen.h>
*/
import "C"

type Color struct {
	base *C.godot_color
}

func (gdt Color) getBase() *C.godot_color {
	return gdt.base
}

// GetR godot_color_get_r [[const godot_color * p_self]] godot_real
func (gdt *Color) GetR() Real {
	arg0 := gdt.getBase()

	ret := C.go_godot_color_get_r(GDNative.api, arg0)

	return Real(ret)
}

// SetR godot_color_set_r [[godot_color * p_self] [const godot_real r]] void
func (gdt *Color) SetR(r Real) {
	arg0 := gdt.getBase()
	arg1 := r.getBase()

	C.go_godot_color_set_r(GDNative.api, arg0, arg1)
}

// GetG godot_color_get_g [[const godot_color * p_self]] godot_real
func (gdt *Color) GetG() Real {
	arg0 := gdt.getBase()

	ret := C.go_godot_color_get_g(GDNative.api, arg0)

	return Real(ret)
}

// SetG godot_color_set_g [[godot_color * p_self] [const godot_real g]] void
func (gdt *Color) SetG(g Real) {
	arg0 := gdt.getBase()
	arg1 := g.getBase()

	C.go_godot_color_set_g(GDNative.api, arg0, arg1)
}

// GetB godot_color_get_b [[const godot_color * p_self]] godot_real
func (gdt *Color) GetB() Real {
	arg0 := gdt.getBase()

	ret := C.go_godot_color_get_b(GDNative.api, arg0)

	return Real(ret)
}

// SetB godot_color_set_b [[godot_color * p_self] [const godot_real b]] void
func (gdt *Color) SetB(b Real) {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	C.go_godot_color_set_b(GDNative.api, arg0, arg1)
}

// GetA godot_color_get_a [[const godot_color * p_self]] godot_real
func (gdt *Color) GetA() Real {
	arg0 := gdt.getBase()

	ret := C.go_godot_color_get_a(GDNative.api, arg0)

	return Real(ret)
}

// SetA godot_color_set_a [[godot_color * p_self] [const godot_real a]] void
func (gdt *Color) SetA(a Real) {
	arg0 := gdt.getBase()
	arg1 := a.getBase()

	C.go_godot_color_set_a(GDNative.api, arg0, arg1)
}

// GetH godot_color_get_h [[const godot_color * p_self]] godot_real
func (gdt *Color) GetH() Real {
	arg0 := gdt.getBase()

	ret := C.go_godot_color_get_h(GDNative.api, arg0)

	return Real(ret)
}

// GetS godot_color_get_s [[const godot_color * p_self]] godot_real
func (gdt *Color) GetS() Real {
	arg0 := gdt.getBase()

	ret := C.go_godot_color_get_s(GDNative.api, arg0)

	return Real(ret)
}

// GetV godot_color_get_v [[const godot_color * p_self]] godot_real
func (gdt *Color) GetV() Real {
	arg0 := gdt.getBase()

	ret := C.go_godot_color_get_v(GDNative.api, arg0)

	return Real(ret)
}

// AsString godot_color_as_string [[const godot_color * p_self]] godot_string
func (gdt *Color) AsString() String {
	arg0 := gdt.getBase()

	ret := C.go_godot_color_as_string(GDNative.api, arg0)

	return String{base: &ret}

}

// ToRgba32 godot_color_to_rgba32 [[const godot_color * p_self]] godot_int
func (gdt *Color) ToRgba32() Int {
	arg0 := gdt.getBase()

	ret := C.go_godot_color_to_rgba32(GDNative.api, arg0)

	return Int(ret)
}

// ToArgb32 godot_color_to_argb32 [[const godot_color * p_self]] godot_int
func (gdt *Color) ToArgb32() Int {
	arg0 := gdt.getBase()

	ret := C.go_godot_color_to_argb32(GDNative.api, arg0)

	return Int(ret)
}

// Gray godot_color_gray [[const godot_color * p_self]] godot_real
func (gdt *Color) Gray() Real {
	arg0 := gdt.getBase()

	ret := C.go_godot_color_gray(GDNative.api, arg0)

	return Real(ret)
}

// Inverted godot_color_inverted [[const godot_color * p_self]] godot_color
func (gdt *Color) Inverted() Color {
	arg0 := gdt.getBase()

	ret := C.go_godot_color_inverted(GDNative.api, arg0)

	return Color{base: &ret}

}

// Contrasted godot_color_contrasted [[const godot_color * p_self]] godot_color
func (gdt *Color) Contrasted() Color {
	arg0 := gdt.getBase()

	ret := C.go_godot_color_contrasted(GDNative.api, arg0)

	return Color{base: &ret}

}

// LinearInterpolate godot_color_linear_interpolate [[const godot_color * p_self] [const godot_color * p_b] [const godot_real p_t]] godot_color
func (gdt *Color) LinearInterpolate(b Color, t Real) Color {
	arg0 := gdt.getBase()
	arg1 := b.getBase()
	arg2 := t.getBase()

	ret := C.go_godot_color_linear_interpolate(GDNative.api, arg0, arg1, arg2)

	return Color{base: &ret}

}

// Blend godot_color_blend [[const godot_color * p_self] [const godot_color * p_over]] godot_color
func (gdt *Color) Blend(over Color) Color {
	arg0 := gdt.getBase()
	arg1 := over.getBase()

	ret := C.go_godot_color_blend(GDNative.api, arg0, arg1)

	return Color{base: &ret}

}

// ToHtml godot_color_to_html [[const godot_color * p_self] [const godot_bool p_with_alpha]] godot_string
func (gdt *Color) ToHtml(withAlpha Bool) String {
	arg0 := gdt.getBase()
	arg1 := withAlpha.getBase()

	ret := C.go_godot_color_to_html(GDNative.api, arg0, arg1)

	return String{base: &ret}

}

// OperatorEqual godot_color_operator_equal [[const godot_color * p_self] [const godot_color * p_b]] godot_bool
func (gdt *Color) OperatorEqual(b Color) Bool {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_color_operator_equal(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// OperatorLess godot_color_operator_less [[const godot_color * p_self] [const godot_color * p_b]] godot_bool
func (gdt *Color) OperatorLess(b Color) Bool {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_color_operator_less(GDNative.api, arg0, arg1)

	return Bool(ret)
}
