package godot

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewEncodedObjectAsIDFromPointer(ptr gdnative.Pointer) EncodedObjectAsID {
func newEncodedObjectAsIDFromPointer(ptr gdnative.Pointer) EncodedObjectAsID {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := EncodedObjectAsID{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type EncodedObjectAsID struct {
	Reference
	owner gdnative.Object
}

func (o *EncodedObjectAsID) BaseClass() string {
	return "EncodedObjectAsID"
}

/*

	Args: [], Returns: int
*/
func (o *EncodedObjectAsID) GetObjectId() gdnative.Int {
	//log.Println("Calling EncodedObjectAsID.GetObjectId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EncodedObjectAsID", "get_object_id")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false id int}], Returns: void
*/
func (o *EncodedObjectAsID) SetObjectId(id gdnative.Int) {
	//log.Println("Calling EncodedObjectAsID.SetObjectId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EncodedObjectAsID", "set_object_id")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// EncodedObjectAsIDImplementer is an interface that implements the methods
// of the EncodedObjectAsID class.
type EncodedObjectAsIDImplementer interface {
	ReferenceImplementer
	GetObjectId() gdnative.Int
	SetObjectId(id gdnative.Int)
}
