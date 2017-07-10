package gosrt

// #cgo pkg-config: srt
// #include <srt/srt.h>
import "C"

import (
	"fmt"
	"net"
	"time"
	"unsafe"
)

// Socket is a SRT socket type.
// To create one, use NewSocket
type Socket struct {
	sockid int
}

// NewSocket creates a new socket
// newType is either INET_4 or INET_6, which are ipv4 and ipv6 repsectively
func NewSocket(netType int) (Socket, error) {
	ret := Socket{}

	ret.sockid = int(C.srt_socket(C.int(netType), C.int(C.SOCK_DGRAM), C.int(0)))

	return ret, chkSrtError(ret.sockid)
}

// Valid checks if the socket is not -1, which is the SRT invalid socket.
func (sock Socket) Valid() bool {
	return sock.sockid != -1
}

// Bind binds to a local IP and socket
// If this socket was created to be an ipv4 socket, then, ip must be an ipv4 address,
// and likewise for ipv6
// This will fail if port has already been bound on this machine
func (sock Socket) Bind(ip net.IP, port int) error {

	if len(ip) == 4 {
		// ipv4

		sockaddr := sockaddrFromIPPort(ip, port)

		// call SRT
		return chkSrtError(int(C.srt_bind(C.SRTSOCKET(sock.sockid), (*C.struct_sockaddr)(unsafe.Pointer(&sockaddr)), C.sizeof_struct_sockaddr_in)))

	}
	if len(ip) != 16 {
		panic(fmt.Sprintf("Unrecognized IP length: %d", len(ip)))
	}

	// ipv6
	sockaddr := sockaddrFromIPPort6(ip, port)

	// call SRT
	return chkSrtError(int(C.srt_bind(C.SRTSOCKET(sock.sockid), (*C.struct_sockaddr)(unsafe.Pointer(&sockaddr)), C.sizeof_struct_sockaddr_in6)))

}

// Listen sets the listen flag in SRT
func (sock Socket) Listen() error {
	return chkSrtError(int(C.srt_listen(C.SRTSOCKET(sock.sockid), C.int(1))))
}

// Accept starts accepting connections
// Listen() must be called first, and it must be bound to a socket
func (sock Socket) Accept() (net.IP, int, Socket, error) {

	// TODO: IPv6?
	sockaddr := C.struct_sockaddr_in{}
	var addrlen C.int

	ret := Socket{}

	ret.sockid = int(C.srt_accept(C.SRTSOCKET(sock.sockid), (*C.struct_sockaddr)(unsafe.Pointer(&sockaddr)), (*C.int)(unsafe.Pointer(&addrlen))))

	ip, socket := ipPortFromSockaddr(sockaddr)

	return ip, socket, ret, chkSrtError(ret.sockid)

}

// Connect connects to another SRT socket
// The other socket must be bound to the port
func (sock Socket) Connect(ip net.IP, port int) error {

	if len(ip) == 4 {

		sockaddr := sockaddrFromIPPort(ip, port)

		return chkSrtError(int(C.srt_connect(C.SRTSOCKET(sock.sockid), (*C.struct_sockaddr)(unsafe.Pointer(&sockaddr)), C.sizeof_struct_sockaddr_in)))

	}
	sockaddr := sockaddrFromIPPort6(ip, port)

	return chkSrtError(int(C.srt_connect(C.SRTSOCKET(sock.sockid), (*C.struct_sockaddr)(unsafe.Pointer(&sockaddr)), C.sizeof_struct_sockaddr_in6)))
}

// Close closes a socket, freeing the port
func (sock Socket) Close() error {
	return chkSrtError(int(C.srt_close(C.SRTSOCKET(sock.sockid))))
}

// GetSockOpt gets an option from the socket
// opt must be one of the gosrt.Opt* (defined in sockopt.go)
func (sock Socket) GetSockOpt(opt int) ([]byte, error) {
	var buffer [128]byte
	var addrlen C.int

	errInt := int(C.srt_getsockopt(C.SRTSOCKET(sock.sockid), C.int(0), C.SRT_SOCKOPT(opt), unsafe.Pointer(&buffer), &addrlen))

	return buffer[:int(addrlen)], chkSrtError(errInt)
}

// SetSockOpt sets an option for the socket
// opt must be one of gosrt.Opt* (defined in sockopt.go)
func (sock Socket) SetSockOpt(opt int, data []byte) error {
	return chkSrtError(int(C.srt_setsockopt(C.SRTSOCKET(sock.sockid), C.int(0), C.SRT_SOCKOPT(opt), unsafe.Pointer(&data[0]), C.int(len(data)))))
}

// SetIntSockOpt Helper function for setting int options
// opt must be one of gosrt.Opt* (defined in sockopt.go)
func (sock Socket) SetIntSockOpt(opt int, value int) error {
	cValue := C.int(value)

	return chkSrtError(int(C.srt_setsockopt(C.SRTSOCKET(sock.sockid), C.int(0), C.SRT_SOCKOPT(opt), unsafe.Pointer(&cValue), C.sizeof_int)))
}

// SetBoolSockOpt is a helper function for setting boolean options
// opt must be one of gosrt.Opt* (defined in sockopt.go)
func (sock Socket) SetBoolSockOpt(opt int, value bool) error {
	if value {
		return sock.SetIntSockOpt(opt, 1)
	}
	return sock.SetIntSockOpt(opt, 0)
}

// SendMsg sends a message over the SRT socket.
// Data over 1316 bytes will be discarded
func (sock Socket) SendMsg(data []byte) error {
	return chkSrtError(int(C.srt_sendmsg(C.SRTSOCKET(sock.sockid), (*C.char)(unsafe.Pointer(&data[0])), C.int(len(data)), C.int(-1), C.int(0))))
}

// SendMsgTimestamped sends a message with a timestamp other than time.Now().
// Data over 1316 bytes will be discarded
func (sock Socket) SendMsgTimestamped(data []byte, timestamp time.Time) error {

	msgCtrl := C.struct_SRT_MsgCtrl_{}
	msgCtrl.srctime = C.uint64_t(timestamp.UnixNano() / 1000) // it accepts usec

	return chkSrtError(int(C.srt_sendmsg2(C.SRTSOCKET(sock.sockid), (*C.char)(unsafe.Pointer(&data)), C.int(len(data)), &msgCtrl)))
}

// RecvMsg recieves a message from the SRT socket. The time is the timestamp of the packet
func (sock Socket) RecvMsg() ([]byte, time.Time, error) {
	var buffer [1316]byte // that is the max SRT payload size

	msgCtrl := C.struct_SRT_MsgCtrl_{}

	size := int(C.srt_recvmsg2(C.SRTSOCKET(sock.sockid), (*C.char)(unsafe.Pointer(&buffer[0])), C.int(len(buffer)), &msgCtrl))

	err := chkSrtError(size)
	if err != nil {
		return nil, time.Time{}, err
	}

	// convert back to nsec
	return buffer[:size], time.Unix(0, int64(msgCtrl.srctime)*1000), chkSrtError(size)
}
