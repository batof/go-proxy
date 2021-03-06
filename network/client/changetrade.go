package client

import "proxy/network"

type ChangeTradePacket struct {
	MyOffers []bool
}

func (c *ChangeTradePacket) Read(p *network.GamePacket) {
	count := p.ReadInt16()
	if count <= 0 {
		return
	}
	c.MyOffers = make([]bool, count)
	for i := 0; i < int(count); i++ {
		c.MyOffers[i] = p.ReadBool()
	}
}

func (c ChangeTradePacket) Write(p *network.GamePacket) {
	count := len(c.MyOffers)
	p.WriteInt16(int16(count))
	if count <= 0 {
		return
	}
	for i := 0; i < count; i++ {
		p.WriteBool(c.MyOffers[i])
	}
}
