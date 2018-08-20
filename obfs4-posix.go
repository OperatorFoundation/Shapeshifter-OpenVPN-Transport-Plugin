package Shapeshifter_OpenVPN_Transport_Plugin

// #include "openvpn-plugin.h"
// #include "openvpn-vsocket.h"
// #include <sys/socket.h>
// #include <netinet/in.h>
// #include <machine/types.h>
//
//struct obfs4_context {
//	struct openvpn_plugin_callbacks *global_vtab;
//};
//
//struct obfs4_socket_posix{
//struct openvpn_vsocket_handle handle;
//struct transport_skeleton_context *ctx;
//int fd;
//unsigned last_rwflags;
//};
//
//static char *get_array_string(char **array, int index) {
//	return array[index];
//}
//
//int open_helper(struct openvpn_plugin_args_open_in const *args, struct openvpn_plugin_args_open_return *out) {
//	struct obfs4_context *context;
//	context = (struct obfs4_context *) calloc(1, sizeof(struct obfs4_context));
//
//	if (!context)
//	return OPENVPN_PLUGIN_FUNC_ERROR;
//
//	context->global_vtab = args->callbacks;
//	out->type_mask = OPENVPN_PLUGIN_MASK(OPENVPN_PLUGIN_SOCKET_INTERCEPT);
//
//	out->handle = (openvpn_plugin_handle_t *) context;
//
//	return OPENVPN_PLUGIN_FUNC_SUCCESS;
//}
//
//struct openvpn_vsocket_vtab *make_vtab() {
//	return calloc(1, sizeof(struct openvpn_vsocket_vtab));
//}
//
import "C"
import (
	"github.com/OperatorFoundation/shapeshifter-transports/transports/obfs4"
	"unsafe"
)

type obfs4_socket_posix struct {

	handle C.struct_openvpn_vsocket_handle
	obfs4_context C.struct_obfs4_context
	fd int
	last_rwflags uint

}

func obfs4_posix_bind(plugin_handle unsafe.Pointer, addr C.struct_sockaddr, len C.openvpn_vsocket_socklen_t)C.struct_openvpn_vsocket_handle_t {

	//openvpn_vsocket_handle_t (*bind_function)(void *plugin_handle,
	//                                  const struct sockaddr *addr,
	//                                  openvpn_vsocket_socklen_t len);
	//

	sock := C.struct_obfs4_socket_posix{}
	sock.handle.vtab = &Obfs4_socket_vtab
	sock.ctx = (C.struct_obfs4_context)(unsafe.Pointer(plugin_handle))

}

func obfs4_posix_request_event(handle C.struct_openvpn_vsocket_handle_t, event_set C.struct_openvpn_vsocket_event_set_handle_t, rwflags uint){
	//void (*request_event_function)(openvpn_vsocket_handle_t handle,
	//                          openvpn_vsocket_event_set_handle_t event_set,
	//                          unsigned rwflags);

}

func obfs4_posix_update_event(handle C.struct_openvpn_vsocket_handle_t, arg unsafe.Pointer, rwflags uint)bool {
	//typedef bool (*update_event_function)(openvpn_vsocket_handle_t handle, void *arg,
	//                         unsigned rwflags);

	return true
}

func obfs4_posix_pump(handle C.struct_openvpn_vsocket_handle_t)uint {
	//unsigned (*pump_function)(openvpn_vsocket_handle_t handle);

	return 0
}

func obfs4_posix_recvfrom(handle C.struct_openvpn_vsocket_handle_t, buf unsafe.Pointer, len C.size_t, addr *C.struct_sockaddr, addrlen C.struct_openvpn_vsocket_socklen_t) C.ssize_t {
	//ssize_t (*recvfrom_function)(openvpn_vsocket_handle_t handle,
	//                                void *buf, size_t len,
	//                                struct sockaddr *addr,
	//                                openvpn_vsocket_socklen_t *addrlen);

	return 0
}

func obfs4_posix_sendto(handle C.struct_openvpn_vsocket_handle_t, buf unsafe.Pointer, len C.size_t, addr *C.struct_sockaddr, addrlen C.struct_openvpn_vsocket_socklen_t)C.ssize_t{
	//ssize_t (*sendto_function)(openvpn_vsocket_handle_t handle,
	//                      const void *buf, size_t len,
	//                      const struct sockaddr *addr,
	//                      openvpn_vsocket_socklen_t addrlen);

	return 0
}

func obfs4_posix_close(handle C.struct_openvpn_vsocket_handle_t){
	//void (*close_function)(openvpn_vsocket_handle_t handle);


}

func obfs4_initialize_socket_vtab() {

	Obfs4_socket_vtab = C.make_vtab()

	Obfs4_socket_vtab.bind = obfs4_posix_bind
	Obfs4_socket_vtab.request_event = obfs4_posix_request_event
	Obfs4_socket_vtab.update_event = obfs4_posix_update_event
	Obfs4_socket_vtab.pump = obfs4_posix_pump
	Obfs4_socket_vtab.recvfrom = obfs4_posix_recvfrom
	Obfs4_socket_vtab.sendto = obfs4_posix_sendto
	Obfs4_socket_vtab.close = obfs4_posix_close

}
