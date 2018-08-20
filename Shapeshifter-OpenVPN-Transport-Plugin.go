// package name: Shapeshifter_OpenVPN_Transport_Plugin
package Shapeshifter_OpenVPN_Transport_Plugin

// #include <stdio.h>
// #include <stdlib.h>
// #include "openvpn-plugin.h"
// #include "openvpn-vsocket.h"
// #include "helper.h"
import "C"
import (
	"github.com/OperatorFoundation/shapeshifter-transports/transports/obfs4"
	"unsafe"
)

var Obfs4_socket_vtab *C.struct_openvpn_vsocket_vtab = C.struct_openvpn_vsocket_vtab{}

//export openvpn_plugin_open_v3
func openvpn_plugin_open_v3(version int, arguments C.struct_openvpn_plugin_args_open_in, output C.struct_openvpn_plugin_args_open_return) int{

	return C.open_helper(arguments, output)
	//// int openvpn_plugin_open_v3(int version, struct openvpn_plugin_args_open_in const *args, struct openvpn_plugin_args_open_return *out)
	//var argv = arguments.argv
	//var certString = C.GoString(C.get_array_string(argv, 1))
	//var iatModeString = C.get_array_string(argv, 2)
	//iatMode, error := strconv.Atoi(C.GoString(iatModeString))
	//if error != nil {
	//	//OPENVPN_PLUGIN_FUNC_ERROR    1
	//	return 1
	//}
	//
	//var obfs4Client  = obfs4.NewObfs4Client(certString, iatMode)
	//
	//// This is just setting up the context
	//// Currently context only does logging to OpenVPN
	//context := &C.struct_obfs4_context{global_vtab: arguments.callbacks}
	//
	//// Sets up the VTable, useful stuff
	////Here is where we should call obfs4_initialize_socket_vtab
	//
	//// Tells openVPN what events we want the plugin to handle
	////OPENVPN_PLUGIN_MASK(OPENVPN_PLUGIN_SOCKET_INTERCEPT)
	//var OPENVPN_PLUGIN_SOCKET_INTERCEPT uint = 13
	//output.type_mask = 1 << OPENVPN_PLUGIN_SOCKET_INTERCEPT
	//
	//// Gives OpenVPN the handle object to save and later give back to us in other calls
	////out->handle = (openvpn_plugin_handle_t *) context
	//
	//output.handle = unsafe.Pointer(context)
	//
	//// #define OPENVPN_PLUGIN_FUNC_SUCCESS  0
	//return 0
}

//export openvpn_plugin_close_v1
func openvpn_plugin_close_v1(handle unsafe.Pointer) {
	// void openvpn_plugin_close_v1(openvpn_plugin_handle_t handle)

}

//export openvpn_plugin_func_v3
func openvpn_plugin_func_v3(version int, arguments C.struct_openvpn_plugin_args_func_in, retrnptr C.struct_openvpn_plugin_args_func_return) int {

	// int openvpn_plugin_func_v3(int version, struct openvpn_plugin_args_func_in const *arguments, struct openvpn_plugin_args_func_return *retptr)

	return 0
}

//export openvpn_plugin_get_vtab_v1
func openvpn_plugin_get_vtab_v1(selector int, size_out C.size_t) unsafe.Pointer {

	// Sets up the VTable, useful stuff
	obfs4_initialize_socket_vtab()

	// void * openvpn_plugin_get_vtab_v1(int selector, size_t *size_out)

	// Provides OpenVPN with the VTable
	// Functions on the VTable are called when there are network events

	//FIXME *size_out = sizeof(struct openvpn_vsocket_vtab);
	return &Obfs4_socket_vtab;
}

func main() {}
