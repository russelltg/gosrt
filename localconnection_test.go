package gosrt

import (
	"net"
	"sync"
	"testing"
)

func TestLocalConnectionIPV4(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(2)

	// server
	go func() {
		defer wg.Done()

		serverBindSock, err := NewSocket(Ipv4)
		if err != nil {
			t.Error("Failed to create a SRT socket", err)
		}
		defer serverBindSock.Close()

		err = serverBindSock.SetBoolSockOpt(OptSender, true)
		if err != nil {
			t.Error("Failed to set sender mode", err)
		}

		// create IP
		ips, err := net.LookupIP("localhost")
		if err != nil {
			t.Error("Failed to get localhost IP", err)
		}

		err = serverBindSock.Bind(ips[0].To4(), 9761)
		if err != nil {
			t.Error("Failed to bind sock to 9761", err)
		}

		err = serverBindSock.Listen()
		if err != nil {
			t.Error("Failed to listen on SRT socket")
		}

		_, _, sock, err := serverBindSock.Accept()
		if err != nil {
			t.Error("Failed to accept", err)
		}
		defer sock.Close()

		// send a few messages
		err = sock.SendMsg(([]byte)("Message 1"))
		if err != nil {
			t.Error("Failed to send", err)
		}

		sock.SendMsg(([]byte)("Message 2"))
		if err != nil {
			t.Error("Failed to send", err)
		}

	}()

	// client
	go func() {
		defer wg.Done()

		sock, err := NewSocket(Ipv4)
		if err != nil {
			t.Error("Failed to create a SRT socket", err)
		}
		defer sock.Close()

		// create IP
		ips, err := net.LookupIP("localhost")
		if err != nil {
			t.Error("Failed to get localhost IP", err)
		}

		err = sock.Connect(ips[0].To4(), 9761)
		if err != nil {
			t.Error("Failed to connect", err)
		}

		data, _, err := sock.RecvMsg()
		if err != nil {
			t.Error("Failed to rcv", err)
		}
		if (string)(data) != "Message 1" {
			t.Error("Mismatching packet data", data)
		}

		data, _, err = sock.RecvMsg()
		if err != nil {
			t.Error("Failed to rcv", err)
		}
		if (string)(data) != "Message 2" {
			t.Error("Mismatching packet data", data)
		}
	}()

	wg.Wait()
}

func TestLocalConnectionIPV6(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(2)

	// server
	go func() {
		defer wg.Done()

		serverBindSock, err := NewSocket(Ipv6)
		if err != nil {
			t.Error("Failed to create a SRT socket", err)
		}
		defer serverBindSock.Close()

		err = serverBindSock.SetBoolSockOpt(OptSender, true)
		if err != nil {
			t.Error("Failed to set sender mode", err)
		}

		// create IP
		ips, err := net.LookupIP("localhost")
		if err != nil {
			t.Error("Failed to get localhost IP", err)
		}

		err = serverBindSock.Bind(ips[0].To16(), 9762)
		if err != nil {
			t.Error("Failed to bind sock to 9762", err)
		}

		err = serverBindSock.Listen()
		if err != nil {
			t.Error("Failed to listen on SRT socket")
		}

		_, _, sock, err := serverBindSock.Accept()
		if err != nil {
			t.Error("Failed to accept", err)
		}
		defer sock.Close()

		// send a few messages
		err = sock.SendMsg(([]byte)("Message 1"))
		if err != nil {
			t.Error("Failed to send", err)
		}

		sock.SendMsg(([]byte)("Message 2"))
		if err != nil {
			t.Error("Failed to send", err)
		}

	}()

	// client
	go func() {
		defer wg.Done()

		sock, err := NewSocket(Ipv6)
		if err != nil {
			t.Error("Failed to create a SRT socket", err)
		}
		defer sock.Close()

		// create IP
		ips, err := net.LookupIP("localhost")
		if err != nil {
			t.Error("Failed to get localhost IP", err)
		}

		err = sock.Connect(ips[0].To16(), 9762)
		if err != nil {
			t.Error("Failed to connect", err)
		}

		data, _, err := sock.RecvMsg()
		if err != nil {
			t.Error("Failed to rcv", err)
		}
		if (string)(data) != "Message 1" {
			t.Error("Mismatching packet data", data)
		}

		data, _, err = sock.RecvMsg()
		if err != nil {
			t.Error("Failed to rcv", err)
		}
		if (string)(data) != "Message 2" {
			t.Error("Mismatching packet data", data)
		}
	}()

	wg.Wait()
}
