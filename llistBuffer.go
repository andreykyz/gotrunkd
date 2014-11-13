// llistBuffer.go
package main

import (
	"gotrunkd/data"
)

type LinkedListFrame struct {
	next  int
	prev  int
	frame data.FramePacket
}

type LinkedListBuffer struct {
	head  int
	tail  int
	size  int
	frame []LinkedListFrame
}

func (llistBuffer *LinkedListBuffer) add(packet []byte, len int, trunk int) int {

}
