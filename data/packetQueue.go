package data

import (
	"container/list"
	"time"
)

type PacketQueue struct {
	lastWrittenSeq     uint32
	lastWrittenTime    time.Time
	frameAmount        int
	missingFrameAmount int
	broken_past_cnt    int
	broken_future_cnt  int
	broken_cnt         int
	list               *list.List
}

func (packetQueue *PacketQueue) init() {
	packetQueue.list = list.New()
	packetQueue.list.Init()
}

/** Check packet status
 *
 */
func (packetQueue *PacketQueue) newPacketCheck(framePacket *FramePacket) int {
	if framePacket.seqNum <= packetQueue.lastWrittenSeq {
		if (packetQueue.lastWrittenSeq - framePacket.seqNum) >= STRANGE_SEQ_PAST {
			packetQueue.broken_cnt++
			packetQueue.broken_past_cnt++
			packetQueue.broken_future_cnt = 0
			return STRANGE_PAST_FRAME
		}
		return DUP_FRAME
	}
	if (framePacket.seqNum - packetQueue.lastWrittenSeq) >= STRANGE_SEQ_FUTURE {
		packetQueue.broken_cnt++
		packetQueue.broken_future_cnt++
		packetQueue.broken_past_cnt = 0
		return STRANGE_FUTURE_FRAME
	}

	return NORMAL
}

func (packetQueue *PacketQueue) newPacketAdd(framePacket *FramePacket) int {
	retValue := NORMAL
	framePacket.packetStrip()
	switch packetQueue.newPacketCheck(framePacket) {
	case STRANGE_PAST_FRAME:
		if packetQueue.broken_past_cnt < BROKEN_CNT {
			return STRANGE_PAST_FRAME
		}
		packetQueue.lastWrittenSeq = framePacket.seqNum - 1
		retValue = STRANGE_PAST_FRAME_FIX
	case STRANGE_FUTURE_FRAME:
		if packetQueue.broken_future_cnt < BROKEN_CNT {
			return STRANGE_FUTURE_FRAME
		}
		packetQueue.lastWrittenSeq = framePacket.seqNum - 1
		retValue = STRANGE_FUTURE_FRAME_FIX

	case DUP_FRAME:
		return DUP_FRAME
	}
	var f, prevf *FramePacket
	for f = packetQueue.list.Front(); f != nil; f = f.Next() {

	}

	return retValue
}
