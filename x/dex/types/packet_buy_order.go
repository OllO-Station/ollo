package types

// ValidateBasic is used for validating the packet
func (p BuyOrderPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p BuyOrderPacketData) GetBytes() ([]byte, error) {
	var modulePacket DexPacketData

	modulePacket.Packet = &DexPacketData_BuyOrderPacket{&p}

	return modulePacket.Marshal()
}
