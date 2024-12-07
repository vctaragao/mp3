package bits

type ChannelMode int

const (
	chStereo ChannelMode = iota
	chJointStereo
	chDualChannel
	chSingleChannel

	chanStereo BitStream = 0x00000000 // 0000 0000 0000 0000 0000 0000 0000 0000
	chanJoint  BitStream = 0x00000004 // 0000 0000 0000 0000 0000 0000 0000 0100
	chanDual   BitStream = 0x00000008 // 0000 0000 0000 0000 0000 0000 0000 1000
	chanSingle BitStream = 0x0000000C // 0000 0000 0000 0000 0000 0000 0000 1100
)

func (c ChannelMode) String() string {
	return []string{"Stereo", "Joint Stereo", "Dual Channel", "Single Channel"}[c]
}

func (c ChannelMode) IsJointStereo() bool {
	return c == chJointStereo
}

func (b BitStream) ParseChannelMode() ChannelMode {
	if b&chanSingle == chanSingle {
		return chSingleChannel
	}

	if b&chanDual == chanDual {
		return chDualChannel
	}

	if b&chanJoint == chanJoint {
		return chJointStereo
	}

	return chStereo
}
