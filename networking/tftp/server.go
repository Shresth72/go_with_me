package main

import (
	"bytes"
	"errors"
	"log"
	"net"
	"time"

	"github.com/Shresth72/sysd/ntw/tftp/types"
)

type Server struct {
	Payload []byte // payload served for read requests
	Retries uint8 // max Tries for failed transmission
	Timeout time.Duration // duration to wait for ACK
}

func (s Server) ListenAndServe(addr string) error {
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	log.Printf("Listening on %s ...\n", conn.LocalAddr())

	return s.Serve(conn)
}

func (s *Server) Serve(conn net.PacketConn) error {
	if conn == nil {
		return errors.New("nil connection")
	}

	if s.Payload == nil {
		return errors.New("payload is required")
	}

	if s.Retries == 0 {
		s.Retries = 10
	}

	if s.Timeout == 0 {
		s.Timeout = 6 * time.Second
	}

	var rrq types.ReadReq

	for {
		buf := make([]byte, types.DatagramSize)

		_, addr, err := conn.ReadFrom(buf)
		if err != nil {
			return err
		}

		err = rrq.UnmarshalBinary(buf)
		if err != nil {
			log.Printf("[%s] bad request: %v", addr, err)
			continue
		}

		go s.handle(addr.String(), rrq)
	}
}

func (s Server) handle(clientAddr string, rrq types.ReadReq) {
	log.Printf("[%s] requested file: %s", clientAddr, rrq.Filename)

	conn, err := net.Dial("udp", clientAddr)
	if err != nil {
		log.Printf("[%s] dial: %v", clientAddr, err)
	}
	defer func() { _ = conn.Close() }()

	var (
		ackPkt  types.Ack
		errPkt  types.Err
		dataPkt = types.Data{Payload: bytes.NewReader(s.Payload)}
		buf     = make([]byte, types.DatagramSize)
	)

NEXTPACKET:
	for n := types.DatagramSize; n == types.DatagramSize; {
		data, err := dataPkt.MarshalBinary()
		if err != nil {
			log.Printf("[%s] preparing data packet: %v", clientAddr, err)
			return
		}

	RETRY:
		for i := s.Retries; i > 0; i-- {
			n, err = conn.Write(data) // send the data packet
			if err != nil {
				log.Printf("[%s] write: %v", clientAddr, err)
				return
			}

			// Wait for Client ACK packet
			_ = conn.SetReadDeadline(time.Now().Add(s.Timeout))

			_, err = conn.Read(buf)
			if err != nil {
				if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
					continue RETRY
				}

				log.Printf("[%s] waiting for ACK: %v", clientAddr, err)
				return
			}

			switch {
			case ackPkt.UnmarshalBinary(buf) == nil:
				if uint16(ackPkt) == dataPkt.Block {
					// received ACK; send next data packet
					continue NEXTPACKET
				}

			case errPkt.UnmarshalBinary(buf) == nil:
				log.Printf("[%s] received error: %v", clientAddr, errPkt.Message)
				return

			default:
				log.Printf("[%s] bad packet", clientAddr)
			}
		}

		log.Printf("[%s] exhausted retries", clientAddr)
	}

	log.Printf("[%s] sent %d blocks", clientAddr, dataPkt.Block)
}
