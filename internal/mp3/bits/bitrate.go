package bits

type BitRate uint32

const (
	bitRateFree BitRate = 0x00000000 // 0000 0000 0000 0000 0000 0000 0000 0000
	bitRate32   BitRate = 0x00001000 // 0000 0000 0000 0000 0001 0000 0000 0000
	bitRate64   BitRate = 0x00002000 // 0000 0000 0000 0000 0010 0000 0000 0000
	bitRate96   BitRate = 0x00003000 // 0000 0000 0000 0000 0011 0000 0000 0000
	bitRate128  BitRate = 0x00004000 // 0000 0000 0000 0000 0100 0000 0000 0000
	bitRate160  BitRate = 0x00005000 // 0000 0000 0000 0000 0101 0000 0000 0000
	bitRate192  BitRate = 0x00006000 // 0000 0000 0000 0000 0110 0000 0000 0000
	bitRate224  BitRate = 0x00007000 // 0000 0000 0000 0000 0111 0000 0000 0000
	bitRate256  BitRate = 0x00008000 // 0000 0000 0000 0000 1000 0000 0000 0000
	bitRate288  BitRate = 0x00009000 // 0000 0000 0000 0000 1001 0000 0000 0000
	bitRate320  BitRate = 0x0000A000 // 0000 0000 0000 0000 1010 0000 0000 0000
	bitRate352  BitRate = 0x0000B000 // 0000 0000 0000 0000 1011 0000 0000 0000
	bitRate384  BitRate = 0x0000C000 // 0000 0000 0000 0000 1100 0000 0000 0000
	bitRate416  BitRate = 0x0000D000 // 0000 0000 0000 0000 1101 0000 0000 0000
	bitRate448  BitRate = 0x0000E000 // 0000 0000 0000 0000 1110 0000 0000 0000
	bitRateBad  BitRate = 0x0000F000 // 0000 0000 0000 0000 1111 0000 0000 0000
)

var bitRateMap = map[BitRate]map[ID][]int{
	bitRate32: {
		MPEG1ID:  {32, 32, 32},
		MPEG2ID:  {32, 8, 8},
		MPEG25ID: {32, 8, 8},
	},
	bitRate64: {
		MPEG1ID:  {64, 48, 40},
		MPEG2ID:  {48, 16, 16},
		MPEG25ID: {48, 16, 16},
	},
	bitRate96: {
		MPEG1ID:  {96, 56, 48},
		MPEG2ID:  {56, 24, 24},
		MPEG25ID: {56, 24, 24},
	},
	bitRate128: {
		MPEG1ID:  {128, 64, 56},
		MPEG2ID:  {64, 32, 32},
		MPEG25ID: {64, 32, 32},
	},
	bitRate160: {
		MPEG1ID:  {160, 80, 64},
		MPEG2ID:  {80, 40, 40},
		MPEG25ID: {80, 40, 40},
	},
	bitRate192: {
		MPEG1ID:  {192, 96, 80},
		MPEG2ID:  {96, 48, 48},
		MPEG25ID: {96, 48, 48},
	},
	bitRate224: {
		MPEG1ID:  {224, 112, 96},
		MPEG2ID:  {112, 56, 56},
		MPEG25ID: {112, 56, 56},
	},
	bitRate256: {
		MPEG1ID:  {256, 128, 112},
		MPEG2ID:  {128, 64, 64},
		MPEG25ID: {128, 64, 64},
	},
	bitRate288: {
		MPEG1ID:  {288, 160, 128},
		MPEG2ID:  {144, 80, 80},
		MPEG25ID: {144, 80, 80},
	},
	bitRate320: {
		MPEG1ID:  {320, 192, 160},
		MPEG2ID:  {160, 96, 96},
		MPEG25ID: {160, 96, 96},
	},
	bitRate352: {
		MPEG1ID:  {352, 224, 192},
		MPEG2ID:  {176, 112, 112},
		MPEG25ID: {176, 112, 112},
	},
	bitRate384: {
		MPEG1ID:  {384, 256, 224},
		MPEG2ID:  {192, 128, 128},
		MPEG25ID: {192, 128, 128},
	},
	bitRate416: {
		MPEG1ID:  {416, 320, 256},
		MPEG2ID:  {224, 114, 114},
		MPEG25ID: {224, 114, 114},
	},
	bitRate448: {
		MPEG1ID:  {448, 384, 320},
		MPEG2ID:  {256, 160, 160},
		MPEG25ID: {256, 160, 160},
	},
}

func (b BitStream) ParseBitRate(layer Layer, id ID) int {
	bitRate := parseBitRate(b)

	if bitRate == bitRateFree {
		return 0
	}

	if bitRate == bitRateBad {
		return -1
	}

	index := 2
	if layer&Layer1 == Layer1 {
		index = 0
	} else if layer&Layer2 == Layer2 {
		index = 1
	}

	return bitRateMap[bitRate][id][index]
}

func parseBitRate(bitStream BitStream) BitRate {
	b := BitRate(bitStream)

	if b&bitRateBad == bitRateBad {
		return bitRateBad
	}

	if b&bitRate448 == bitRate448 {
		return bitRate448
	}

	if b&bitRate416 == bitRate416 {
		return bitRate416
	}

	if b&bitRate384 == bitRate384 {
		return bitRate384
	}

	if b&bitRate352 == bitRate352 {
		return bitRate352
	}
	if b&bitRate320 == bitRate320 {
		return bitRate320
	}

	if b&bitRate288 == bitRate288 {
		return bitRate288
	}

	if b&bitRate256 == bitRate256 {
		return bitRate256
	}

	if b&bitRate224 == bitRate224 {
		return bitRate224
	}

	if b&bitRate192 == bitRate192 {
		return bitRate192
	}

	if b&bitRate160 == bitRate160 {
		return bitRate160
	}

	if b&bitRate128 == bitRate128 {
		return bitRate128
	}

	if b&bitRate96 == bitRate96 {
		return bitRate96
	}

	if b&bitRate64 == bitRate64 {
		return bitRate64
	}

	if b&bitRate32 == bitRate32 {
		return bitRate32
	}

	return bitRateFree
}
