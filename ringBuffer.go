// queueRing.go
package main

/*
	Ring Buffer

	/   RingBuffer.size     \
	|                       |
	-------------------------
	\/                      /\
	# # # # # # # # # # # # #

	head    		   tail
	<--- 1 2 3 4 5 6 7 <----
*/

type RingBuffer struct {
	head  int
	tail  int
	size  int
	frame []FramePacket
}

func (ringBuffer *RingBuffer) add(packet []byte, len int, trunk int) int {
	retValue := NORMAL
	seqNum := binary.BigEndian.Uint32(packet[2:6])
	seqPointer := int(seqNum-ringBuffer.frame[ringBuffer.tail].seq_num) + ringBuffer.tail
	switch {
	case (seqNum - ringBuffer.frame[ringBuffer.tail].seq_num) > (ringBuffer.size*20)/100:
		ringBuffer.head = 0
		ringBuffer.tail = 0
		seqPointer = 0
		retValue = BIG_LOSS
	case seqPointer >= ringBuffer.size:
		seqPointer = seqPointer - ringBuffer.size
	}

	if seqPointer >= ringBuffer.head { //todo fixme detect overflow with ring
		//		retValue = RINGBUFFER_OVERFLOW
	}

	if ringBuffer.frame[seqPointer].len != 0 {
		return DUP_FRAME
	}

	copy(ringBuffer.frame[seqPointer].packet, packet[6:len])
	ringBuffer.frame[seqPointer].len = len
	ringBuffer.frame[seqPointer].trunk = trunk
	return retValue
}

func (ringBuffer *RingBuffer) getWaitData(lastWrittenSeq int) bool {
	if lastWrittenSeq+1 == ringBuffer.head {
		return true
	}
	return false
}
