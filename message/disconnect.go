package message

import "github.com/galaco/bitbuf"

type MsgDisconnect struct {
	buf []byte
}

// Connectionless: is this message a connectionless message?
func (msg *MsgDisconnect) Connectionless() bool {
	return false
}

// Data Get packet data
func (msg *MsgDisconnect) Data() []byte {
	return msg.buf
}

// Disconnect returns new disconnect packet data
func Disconnect(msg string) *MsgDisconnect {
	buf := bitbuf.NewWriter(1024)

	buf.WriteUnsignedBitInt32(1, 6)
	buf.WriteString(msg)
	buf.WriteByte(0)

	return &MsgDisconnect{
		buf: buf.Data(),
	}
}
