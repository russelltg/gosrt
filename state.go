package gosrt

// #cgo pkg-config: srt
// #include <srt/srt.h>
import "C"


const (
	STATE_INIT       = C.SRTS_INIT
	STATE_OPENED     = C.SRTS_OPENED         
	STATE_LISTENING  = C.SRTS_LISTENING         
	STATE_CONNECTING = C.SRTS_CONNECTING         
	STATE_CONNECTED  = C.SRTS_CONNECTED         
	STATE_BROKEN     = C.SRTS_BROKEN         
	STATE_CLOSING    = C.SRTS_CLOSING         
	STATE_CLOSED     = C.SRTS_CLOSED         
	STATE_NONEXIST   = C.SRTS_NONEXIST         
)

func (sock Socket) State() int {
	return int(C.srt_getsockstate(C.SRTSOCKET(sock.sockid)))
}
