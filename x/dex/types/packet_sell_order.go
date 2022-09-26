package types

// ValidateBasic is used for validating the packet
func (p SellOrderPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p SellOrderPacketData) GetBytes() ([]byte, error) {
	var modulePacket DexPacketData

	modulePacket.Packet = &DexPacketData_SellOrderPacket{&p}

	return modulePacket.Marshal()
}
