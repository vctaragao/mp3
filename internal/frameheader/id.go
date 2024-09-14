package frameheader

type ID uint32

const (
	MPEG25ID       ID = 0x00000000 // 0000 0000 0000 0000 0000 0000 0000 0000
	MPEG2ID        ID = 0x00100000 // 0000 0000 0001 0000 0000 0000 0000 0000
	MPEGReservedID ID = 0x00080000 // 0000 0000 0000 1000 0000 0000 0000 0000
	MPEG3ID        ID = 0x00180000 // 0000 0000 0001 1000 0000 0000 0000 0000
)

func (id ID) String() string {
	return []string{"MPEG25ID", "MPEG2ID", "MPEGReservedID", "MPEG3ID"}[id.Int()]
}

func (id ID) Int() int {
	return int(id >> 19)
}

func fromFrameHeader(bitstream uint32) ID {
	fID := ID(bitstream)

	if fID&MPEG25ID == MPEG25ID {
		return MPEG25ID
	}

	if fID&MPEG2ID == MPEG2ID {
		return MPEG2ID
	}

	if fID&MPEG3ID == MPEG3ID {
		return MPEG3ID
	}

	return MPEGReservedID
}
