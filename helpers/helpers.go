package helpers

import (
	"encoding/binary"
	"io"
	"net"
)

// Read reads a message from a connection
func Read(conn net.Conn) (msg []byte, err error) {
	var l uint64
	if err = binary.Read(conn, binary.BigEndian, &l); err != nil {
		return nil, err
	}

	msg = make([]byte, l)
	_, err = io.ReadFull(conn, msg)

	return msg, err
}

// Write writes a message to a connection
func Write(conn net.Conn, msg []byte) (err error) {
	if err = binary.Write(conn, binary.BigEndian, uint64(len(msg))); err != nil {
		return err
	}
	_, err = conn.Write(msg)
	return err
}
