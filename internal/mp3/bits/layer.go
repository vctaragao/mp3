package bits

type Layer uint32

const (
	LayerReserved Layer = 0x00000000 // 0000 0000 0000 0000 0000 0000 0000 0000
	Layer3        Layer = 0x00020000 // 0000 0000 0000 0010 0000 0000 0000 0000
	Layer2        Layer = 0x00040000 // 0000 0000 0000 0100 0000 0000 0000 0000
	Layer1        Layer = 0x00060000 // 0000 0000 0000 0110 0000 0000 0000 0000
)

func (l Layer) String() string {
	return map[Layer]string{
		LayerReserved: "reserved",
		Layer3:        "Layer 3",
		Layer2:        "Layer 2",
		Layer1:        "Layer 1",
	}[l]
}

func (b BitStream) ParseLayer() Layer {
	l := Layer(b)

	if l&Layer1 == Layer1 {
		return Layer1
	}

	if l&Layer2 == Layer2 {
		return Layer2
	}

	if l&Layer3 == Layer3 {
		return Layer3
	}

	return LayerReserved

}
