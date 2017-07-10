package gosrt

// #cgo pkg-config: srt
// #include <srt/srt.h>
import "C"

const (
	// OptMms : is the Maximum Transfer Unit
	OptMms = C.SRTO_MSS

	// OptSendSync : if sending is blocking
	OptSendSync = C.SRTO_SNDSYN

	// OptRecieveSync : if recieving is blocking
	OptRecieveSync = C.SRTO_RCVSYN

	// OptCc : custom congestion control algorithm
	OptCc = C.SRTO_CC

	// OptFc :  Flight flag size (window size)
	OptFc = C.SRTO_FC

	// OptSendBuffer : the maximum buffer in sending queue
	OptSendBuffer = C.SRTO_SNDBUF

	// OptRecieveBuffer : the maximum buffer in recieve queue
	OptRecieveBuffer = C.SRTO_RCVBUF

	// OptLinger : waiting for unsent data when closing
	OptLinger = C.SRTO_LINGER

	// OptUDPSendBufer : UDP sending buffer size
	OptUDPSendBufer = C.SRTO_UDP_SNDBUF

	// OptUDPRecieveBuffer : UDP receiving buffer size
	OptUDPRecieveBuffer = C.SRTO_UDP_RCVBUF

	// OptMaxMsg : maximum datagram message size
	OptMaxMsg = C.SRTO_MAXMSG

	// OptMsgTTL : time-to-live of a datagram message
	OptMsgTTL = C.SRTO_MSGTTL

	// OptRendezvous : rendezvous connection mode
	OptRendezvous = C.SRTO_RENDEZVOUS

	// OptSendTimeOut : SendMsg() timeout
	OptSendTimeOut = C.SRTO_SNDTIMEO

	// OptRecieveTimeOut : RecvMsg() timeout
	OptRecieveTimeOut = C.SRTO_RCVTIMEO

	// OptReusePort : reuse an existing port or create a new one
	OptReusePort = C.SRTO_REUSEADDR

	// OptMaxBandwidth : maximum bandwidth (bytes per second) that the connection can use
	OptMaxBandwidth = C.SRTO_MAXBW

	// OptState : current socket state, see UDTSTATUS, read only
	OptState = C.SRTO_STATE

	// OptEvent : current available events associated with the socket
	OptEvent = C.SRTO_EVENT

	// OptSendData : size of data in the sending buffer (read only)
	OptSendData = C.SRTO_SNDDATA

	// OptRecieveData : size of data available for recv (read only)
	OptRecieveData = C.SRTO_RCVDATA

	// OptSender : Sender mode (independent of conn mode), for encryption, tsbpd handshake.
	OptSender = C.SRTO_SENDER

	// OptTsbpdMode : Enable/Disable TsbPd. Enable -> Tx set origin timestamp, Rx deliver packet at origin time + delay. Should pretty much always be on, enabled by default
	OptTsbpdMode = C.SRTO_TSBPDMODE

	// OptLatency : TsbPd receiver delay (mSec) to absorb burst of missed packet retransmission
	OptLatency = C.SRTO_LATENCY

	// OptInputBandwidth : Estimated input stream rate. If not set, will by deduced by profiling data sent
	OptInputBandwidth = C.SRTO_INPUTBW

	// OptOverheadBandwidth : MaxBW ceiling based on % over input stream rate. Applies when OptMaxBandwidth=0 (auto).
	OptOverheadBandwidth = C.SRTO_OHEADBW

	// OptPassphrase : Crypto PBKDF2 Passphrase size[0,10..64] 0:disable crypto
	OptPassphrase = C.SRTO_PASSPHRASE

	// OptCryptoSize : Crypto key len in bytes {16,24,32} Default: 16 (128-bit)
	OptCryptoSize = C.SRTO_PBKEYLEN

	// OptKmState : Key Material exchange status (UDT_SRTKmState)
	OptKmState = C.SRTO_KMSTATE

	// OptIPTTL : IP Time To Live
	OptIPTTL = C.SRTO_IPTTL

	// OptIPTOS : IP Type of Service
	OptIPTOS = C.SRTO_IPTOS

	// OptRecievePacketDrop : Enable receiver pkt drop
	OptRecievePacketDrop = C.SRTO_TLPKTDROP

	// OptNakeReport : Enable receiver to send periodic NAK reports
	OptNakeReport = C.SRTO_NAKREPORT

	// OptVersion : Local SRT Version
	OptVersion = C.SRTO_VERSION

	// OptPeerVersion : Peer SRT Version (from SRT Handshake)
	OptPeerVersion = C.SRTO_PEERVERSION

	// OptConnectionTimeout : Connect timeout in msec. Ccaller default: 3000, rendezvous (x 10)
	OptConnectionTimeout = C.SRTO_CONNTIMEO

	// OptTwoWayData : Allow two way data exchange
	OptTwoWayData = C.SRTO_TWOWAYDATA

	// These aren't documented in SRT, but they should be here anyways. TODO: actually see what they do

	OptSendCryptoLength    = C.SRTO_SNDPBKEYLEN
	OptRecieveCryptoLength = C.SRTO_RCVPBKEYLEN
	OptSendPeerKMSState    = C.SRTO_SNDPEERKMSTATE
	OptRecieveKSMState     = C.SRTO_RCVKMSTATE
	OptMaxLossTTL          = C.SRTO_LOSSMAXTTL
)
