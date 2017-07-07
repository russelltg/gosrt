package gosrt

// #cgo pkg-config: srt
// #include <srt/srt.h>
import "C"

import (
	"net"
	"time"
	"unsafe"
)

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

// Bind binds to a local IP and socket
// If this socket was created to be an ipv4 socket, then, ip must be an ipv4 address,
// and likewise for ipv6
// This will fail if port has already been bound on this machine
func (sock Socket) Bind(ip net.IP, port int) error {

	if len(ip) == 4 {
		// ipv4

		sockaddr := sockaddrFromIpPort(ip, port)

		// call SRT
		return chkSrtError(int(C.srt_bind(C.SRTSOCKET(sock.sockid), (*C.struct_sockaddr)(unsafe.Pointer(&sockaddr)), C.sizeof_struct_sockaddr_in)))

	} else {
		// ipv6
		sockaddr := sockaddrFromIpPort6(ip, port)

		// call SRT
		return chkSrtError(int(C.srt_bind(C.SRTSOCKET(sock.sockid), (*C.struct_sockaddr)(unsafe.Pointer(&sockaddr)), C.sizeof_struct_sockaddr_in6)))

	}
}

// Listen sets the listen flag in SRT
func (sock Socket) Listen() error {
	return chkSrtError(int(C.srt_listen(C.SRTSOCKET(sock.sockid), C.int(1))))
}

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

		sockaddr := sockaddrFromIpPort(ip, port)

		return chkSrtError(int(C.srt_connect(C.SRTSOCKET(sock.sockid), (*C.struct_sockaddr)(unsafe.Pointer(&sockaddr)), C.sizeof_struct_sockaddr_in)))

	} else {
		sockaddr := sockaddrFromIpPort6(ip, port)

		return chkSrtError(int(C.srt_connect(C.SRTSOCKET(sock.sockid), (*C.struct_sockaddr)(unsafe.Pointer(&sockaddr)), C.sizeof_struct_sockaddr_in6)))
	}
}

func (sock Socket) Close() error {
	return chkSrtError(int(C.srt_close(C.SRTSOCKET(sock.sockid))))
}

func (sock Socket) GetSockOpt(opt int) ([]byte, error) {
	var buffer [128]byte
	var addrlen C.int

	errInt := int(C.srt_getsockopt(C.SRTSOCKET(sock.sockid), C.int(0), C.SRT_SOCKOPT(opt), unsafe.Pointer(&buffer), &addrlen))

	return buffer[:int(addrlen)], chkSrtError(errInt)
}

func (sock Socket) SetSockOpt(opt int, data []byte) error {
	return chkSrtError(int(C.srt_setsockopt(C.SRTSOCKET(sock.sockid), C.int(0), C.SRT_SOCKOPT(opt), unsafe.Pointer(&data[0]), C.int(len(data)))))
}

// Helper function for setting int options
func (sock Socket) SetIntSockOpt(opt int, value int) error {
	cValue := C.int(value)

	return chkSrtError(int(C.srt_setsockopt(C.SRTSOCKET(sock.sockid), C.int(0), C.SRT_SOCKOPT(opt), unsafe.Pointer(&cValue), C.sizeof_int)))
}

func (sock Socket) SetBoolSockOpt(opt int, value bool) error {
	if value {
		return sock.SetIntSockOpt(opt, 1)
	} else {
		return sock.SetIntSockOpt(opt, 0)
	}
}

// Data over 1316 bytes will be discarded
func (sock Socket) SendMsg(data []byte) error {
	return chkSrtError(int(C.srt_sendmsg(C.SRTSOCKET(sock.sockid), (*C.char)(unsafe.Pointer(&data[0])), C.int(len(data)), C.int(-1), C.int(0))))
}

// Data over 1316 bytes will be discarded
func (sock Socket) SendMsgTimestamped(data []byte, timestamp time.Time) error {

	msgCtrl := C.struct_SRT_MsgCtrl_{}
	msgCtrl.srctime = C.uint64_t(timestamp.UnixNano() / 1000) // it accepts usec

	return chkSrtError(int(C.srt_sendmsg2(C.SRTSOCKET(sock.sockid), (*C.char)(unsafe.Pointer(&data)), C.int(len(data)), &msgCtrl)))
}

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
