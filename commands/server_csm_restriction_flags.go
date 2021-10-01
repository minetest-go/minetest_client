package commands

type ServerCSMRestrictionFlags struct{}

func (p *ServerCSMRestrictionFlags) GetCommandId() uint16 {
	return ServerCommandCSMRestrictionFlags
}

func (p *ServerCSMRestrictionFlags) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerCSMRestrictionFlags) UnmarshalPacket(payload []byte) error {
	return nil
}

func (p *ServerCSMRestrictionFlags) String() string {
	return "{ServerCSMRestrictionFlags}"
}
