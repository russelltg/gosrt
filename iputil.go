package gosrt

// #cgo pkg-config: srt
// #include <srt/srt.h>
import "C"

import "net"
import "unsafe"
import "errors"
import "encoding/binary"

const (
	INET_4 = C.AF_INET
	INET_6 = C.AF_INET6
)

// Check for SRT errors
func chkSrtError(errorCode int) error {
	if errorCode == -1 {
		return errors.New(C.GoString(C.srt_getlasterror_str()))
	}
	return nil
}

func sockaddrFromIpPort(ip net.IP,  port int) C.struct_sockaddr_in {
    
    if len(ip) != 4 {
        panic("Internal SRT error: cannot get sockaddr ipv4 from an ipv6 IP")
    }
    
	// create sockaddr
	sockaddr := C.struct_sockaddr_in{}

	// zero
	C.memset(unsafe.Pointer(&sockaddr), C.int(0), C.sizeof_struct_sockaddr_in)

	var noPortBytes [2]byte
	binary.LittleEndian.PutUint16(noPortBytes[:], uint16(port))
	
	var noPort C.uint16_t;
	
	C.memcpy(unsafe.Pointer(&noPort), unsafe.Pointer(&noPortBytes[0]), 2)
	
	sockaddr.sin_family = C.sa_family_t(INET_4)
	sockaddr.sin_port = C.in_port_t(noPort)

	// set it
	C.memcpy(unsafe.Pointer(&sockaddr.sin_addr), unsafe.Pointer(&ip[0]), C.size_t(4))

	return sockaddr

}

func sockaddrFromIpPort6(ip net.IP,  port int) C.struct_sockaddr_in6 {

    
    if len(ip) != 4 {
        panic("Internal SRT error: cannot get sockaddr ipv4 from an ipv6 IP")
    }
    
	// create sockaddr
	sockaddr := C.struct_sockaddr_in6{}
	
	// zero
	C.memset(unsafe.Pointer(&sockaddr), C.int(0), C.sizeof_struct_sockaddr_in6)

	var noPortBytes [2]byte
	binary.LittleEndian.PutUint16(noPortBytes[:], uint16(port))
	
	var noPort C.uint16_t;
	
	C.memcpy(unsafe.Pointer(&noPort), unsafe.Pointer(&noPortBytes[0]), 2)
	
	
	sockaddr.sin6_family = C.sa_family_t(INET_6)
	sockaddr.sin6_port = C.in_port_t(noPort)

	// set it
	C.memcpy(unsafe.Pointer(&sockaddr.sin6_addr), unsafe.Pointer(&ip[0]), C.size_t(16))

	return sockaddr
}

func ipPortFromSockaddr(sockaddr C.struct_sockaddr_in) (net.IP, int) {

	// copy the IP out
	var ipBuff [4]byte
	C.memcpy(unsafe.Pointer(&ipBuff), unsafe.Pointer(&sockaddr.sin_addr), 4)

	// get the socket
	socket := int(sockaddr.sin_port)

	return net.IP(ipBuff[:]), socket

}

func ipPortFromSockaddr6(sockaddr C.struct_sockaddr_in6) (net.IP, int) {

	// copy the IP out
	var ipBuff [16]byte
	C.memcpy(unsafe.Pointer(&ipBuff), unsafe.Pointer(&sockaddr.sin6_addr), 16)

	// get the socket
	socket := int(sockaddr.sin6_port)

	return net.IP(ipBuff[:]), socket

}
