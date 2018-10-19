package handlers

import "errors"

/**
  packet Info Define
*/
// FIXME 使用interface 改写成定长和非定长可切换的

// 包定义格式，现在仅固定长度的
type PacketInfo struct {
	HasFixedHead bool
	HeadStyle    []byte
	//if BodyLength==0 then need to calc the body length
	BodyLength    int
	HasFixedTail  bool
	TailStyle     []byte
	HasCheckSum   bool
	CheckSumStyle []byte
}

//包格式
type Packet struct {
	*PacketInfo
	Content  []byte
	Checksum []byte
}

func (p Packet) Len() int {
	return len(p.Content)
}

//仅适用于固定长度的包
func (pi PacketInfo) CalcLen() int {
	var rtn = 0
	if pi.HasFixedHead {
		rtn += len(pi.HeadStyle)
	}
	rtn += pi.BodyLength

	if pi.HasCheckSum {
		rtn += len(pi.CheckSumStyle)
	}

	if pi.HasFixedTail {
		rtn += len(pi.TailStyle)
	}
	return rtn
}

// 数据包数据 从哪里开始 移除包头包尾校验的开始位置
// 如8字节 2头2 尾 1校验，则开始于位置2 （body含2）
func (p Packet) BodyStartAt() int {
	if p.HasFixedHead {
		return len(p.HeadStyle)
	} else {
		return 0
	}
}

// 数据包数据 从哪里结束 移除包头包尾校验的结束位置
// 如8字节 2头2 尾 1校验，则结束于位置5 （body不含5）
func (p Packet) BodyEndAt() int {
	//rtn := p.Len()
	rtn := p.CalcLen()

	if p.HasFixedTail {
		rtn -= len(p.TailStyle)
	}
	if p.HasCheckSum {
		rtn -= len(p.CheckSumStyle)
	}
	return rtn
}

// checksum位置 如果无 则返回-1
func (p Packet) CheckSumAt() int {
	if p.HasCheckSum {
		return p.BodyEndAt()
	} else {
		return -1
	}
}

// 尾开始位置 如果无 则返回-1
func (p Packet) TailStartAt() int {
	if p.HasFixedTail {
		return p.CalcLen() - len(p.TailStyle)
	} else {
		return -1
	}
}

//计算校验和
func (p Packet) CalcCheckSumByte() (checksum byte, err error) {
	if p.HasCheckSum {
		body := p.Content[p.BodyStartAt():p.BodyEndAt()]
		//var checksum byte
		checksum = 0
		for _, n := range body {
			checksum += n
		}
	} else {
		// do nothing
		err = errors.New("can not calc checksum if has check sum flag equals false")
	}
	return
}

//计算校验和
func (p Packet) CalcCheckSum() (checksum []byte, err error) {

	if len(p.CheckSumStyle) == 1 {
		checkbyte, err := p.CalcCheckSumByte()
		if err != nil {

		} else {
			checksum = make([]byte, 1)
			checksum[0] = checkbyte
		}
	} else {
		//FIXME 还没有实装
		err = errors.New("the author have not compelete this func")
	}
	return
}

//写校验和
func (p Packet) WriteCheckSum() (err error) {
	checksum, err := p.CalcCheckSum()
	if err != nil {
		return err
	} else {
		copy(p.Content[p.CheckSumAt():p.CheckSumAt()+len(p.CheckSumStyle)], checksum)
		//copy(checksum, checksum)
	}
	return nil
}
func (p Packet) GetBody() []byte {
	return p.Content[p.BodyStartAt():p.BodyEndAt()]
}

//--------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------------------------------------------------------------------
