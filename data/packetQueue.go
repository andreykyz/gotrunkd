package data

import (
	//	"encoding/binary"
	"time"
)

type PacketQueue struct {
	lastWrittenSeq     uint32
	lastWrittenTime    time.Time
	frameAmount        int
	missingFrameAmount int
}

func (packetQueue *PacketQueue) newPacketAdd(framePacket *FramePacket) int {

	if framePacket.seqNum <= packetQueue.lastWrittenSeq {
		if (packetQueue.lastWrittenSeq - framePacket.seqNum) >= STRANGE_SEQ_PAST {
			return STRANGE_PAST_FRAME
		}
		return DUP_FRAME
	}
	if (framePacket.seqNum - packetQueue.lastWrittenSeq) >= STRANGE_SEQ_FUTURE {
		return STRANGE_FUTURE_FRAME
	}

	return NORMAL
}
