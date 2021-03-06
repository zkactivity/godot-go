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

//func NewPacketPeerUDPFromPointer(ptr gdnative.Pointer) PacketPeerUDP {
func newPacketPeerUDPFromPointer(ptr gdnative.Pointer) PacketPeerUDP {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PacketPeerUDP{}
	obj.SetBaseObject(owner)

	return obj
}

/*
UDP packet peer. Can be used to send raw UDP packets as well as [Variant]s.
*/
type PacketPeerUDP struct {
	PacketPeer
	owner gdnative.Object
}

func (o *PacketPeerUDP) BaseClass() string {
	return "PacketPeerUDP"
}

/*
        Close the UDP socket the [code]PacketPeerUDP[/code] is currently listening on.
	Args: [], Returns: void
*/
func (o *PacketPeerUDP) Close() {
	//log.Println("Calling PacketPeerUDP.Close()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PacketPeerUDP", "close")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Return the IP of the remote peer that sent the last packet(that was received with [method get_packet] or [method get_var]).
	Args: [], Returns: String
*/
func (o *PacketPeerUDP) GetPacketIp() gdnative.String {
	//log.Println("Calling PacketPeerUDP.GetPacketIp()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PacketPeerUDP", "get_packet_ip")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Return the port of the remote peer that sent the last packet(that was received with [method get_packet] or [method get_var]).
	Args: [], Returns: int
*/
func (o *PacketPeerUDP) GetPacketPort() gdnative.Int {
	//log.Println("Calling PacketPeerUDP.GetPacketPort()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PacketPeerUDP", "get_packet_port")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Return whether this [code]PacketPeerUDP[/code] is listening.
	Args: [], Returns: bool
*/
func (o *PacketPeerUDP) IsListening() gdnative.Bool {
	//log.Println("Calling PacketPeerUDP.IsListening()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PacketPeerUDP", "is_listening")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Make this [code]PacketPeerUDP[/code] listen on the "port" binding to "bind_address" with a buffer size "recv_buf_size". If "bind_address" is set as "*" (default), the peer will listen on all available addresses (both IPv4 and IPv6). If "bind_address" is set as "0.0.0.0" (for IPv4) or "::" (for IPv6), the peer will listen on all available addresses matching that IP type. If "bind_address" is set to any valid address (e.g. "192.168.1.101", "::1", etc), the peer will only listen on the interface with that addresses (or fail if no interface with the given address exists).
	Args: [{ false port int} {* true bind_address String} {65536 true recv_buf_size int}], Returns: enum.Error
*/
func (o *PacketPeerUDP) Listen(port gdnative.Int, bindAddress gdnative.String, recvBufSize gdnative.Int) gdnative.Error {
	//log.Println("Calling PacketPeerUDP.Listen()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(port)
	ptrArguments[1] = gdnative.NewPointerFromString(bindAddress)
	ptrArguments[2] = gdnative.NewPointerFromInt(recvBufSize)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PacketPeerUDP", "listen")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Set the destination address and port for sending packets and variables, a hostname will be resolved using if valid.
	Args: [{ false host String} { false port int}], Returns: enum.Error
*/
func (o *PacketPeerUDP) SetDestAddress(host gdnative.String, port gdnative.Int) gdnative.Error {
	//log.Println("Calling PacketPeerUDP.SetDestAddress()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(host)
	ptrArguments[1] = gdnative.NewPointerFromInt(port)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PacketPeerUDP", "set_dest_address")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Wait for a packet to arrive on the listening port, see [method listen].
	Args: [], Returns: enum.Error
*/
func (o *PacketPeerUDP) Wait() gdnative.Error {
	//log.Println("Calling PacketPeerUDP.Wait()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PacketPeerUDP", "wait")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

// PacketPeerUDPImplementer is an interface that implements the methods
// of the PacketPeerUDP class.
type PacketPeerUDPImplementer interface {
	PacketPeerImplementer
	Close()
	GetPacketIp() gdnative.String
	GetPacketPort() gdnative.Int
	IsListening() gdnative.Bool
}
