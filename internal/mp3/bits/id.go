package bits

type ID uint32

const (
	MPEG25ID       ID = 0x00000000 // 0000 0000 0000 0000 0000 0000 0000 0000
	MPEG2ID        ID = 0x00100000 // 0000 0000 0001 0000 0000 0000 0000 0000
	MPEGReservedID ID = 0x00080000 // 0000 0000 0000 1000 0000 0000 0000 0000
	MPEG1ID        ID = 0x00180000 // 0000 0000 0001 1000 0000 0000 0000 0000
)

func (id ID) String() string {
	return []string{"MPEG2.5", "MPEG2", "MPEGReserved", "MPEG1"}[id.Int()]
}

func (id ID) Int() int {
	return int(id >> 19)
}

func (b BitStream) IDFromFrameHeader() ID {
	fID := ID(b)

	if fID&MPEG1ID == MPEG1ID {
		return MPEG1ID
	}

	if fID&MPEG2ID == MPEG2ID {
		return MPEG2ID
	}

	if fID&MPEGReservedID == MPEGReservedID {
		return MPEG25ID
	}

	return MPEG25ID
}
