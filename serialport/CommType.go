package serialport

/**
    packet Info Define
 */
type PacketInfo struct {
    HasFixedHead bool
    HeadStyle    []byte
    //if BodyLength==0 then need to calc the body length
    BodyLength   int
    HasFixedTail bool
    TailStyle    []byte
    MinLength    int
}

type Packet struct {
    Content  []byte
    Checksum []byte
}

//TODO 需要将predefine移除comm 移动到device中

func (pi *PacketInfo) PreDefine() {
    pi.HasFixedHead = true
    pi.HeadStyle = []byte{0xAA, 0xAA}
    pi.MinLength = 12
    pi.BodyLength = 8
    pi.HasFixedTail = true
    pi.TailStyle = []byte{0xBB, 0xBB}
}

func (pkt *Packet) Len() int {
    return len(pkt.Content)
}

func (pkt *Packet) New(content []byte) {
    pkt.Content = content
    pkt.Checksum = nil
}