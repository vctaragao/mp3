package bits

type Frequency int

const (
	sf44100 Frequency = iota
	sf22050
	sf11025
	sf48000
	sf24000
	sf12000
	sf32000
	sf16000
	sf8000
	sfreserved
)

func (f Frequency) String() string {
	return []string{"44100", "22050", "11025", "48000", "24000", "12000", "32000", "16000", "8000", "reserved"}[f]
}

const (
	sfIdentifer1 BitStream = 0x00000000 // 0000 0000 0000 0000 0000 0000 0000 0000
	sfIdentifer2 BitStream = 0x00000400 // 0000 0000 0000 0000 0000 0100 0000 0000
	sfIdentifer3 BitStream = 0x00000800 // 0000 0000 0000 0000 0000 1000 0000 0000
	sfReserved   BitStream = 0x00000C00 // 0000 0000 0000 0000 0000 1100 0000 0000
)

func (b BitStream) ParseSampeFrequency(id ID) Frequency {
	return map[BitStream]map[ID]Frequency{
		sfIdentifer1: {MPEG1ID: sf44100, MPEG2ID: sf22050, MPEG25ID: sf11025},
		sfIdentifer2: {MPEG1ID: sf48000, MPEG2ID: sf24000, MPEG25ID: sf12000},
		sfIdentifer3: {MPEG1ID: sf32000, MPEG2ID: sf16000, MPEG25ID: sf8000},
		sfReserved:   {MPEG1ID: sfreserved, MPEG2ID: sfreserved, MPEG25ID: sfreserved},
	}[b.parseIdentifier()][id]
}

func (b BitStream) parseIdentifier() BitStream {
	if sfIdentifer1&b == sfIdentifer1 {
		return sfIdentifer1
	}

	if sfIdentifer2&b == sfIdentifer2 {
		return sfIdentifer2
	}

	if sfIdentifer3&b == sfIdentifer3 {
		return sfIdentifer3
	}

	return sfReserved
}
