package commands

type ServerAddParticleSpawner struct{}

func (p *ServerAddParticleSpawner) GetCommandId() uint16 {
	return ServerCommandAddParticleSpawner
}

func (p *ServerAddParticleSpawner) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerAddParticleSpawner) UnmarshalPacket(payload []byte) error {
	return nil
}

func (p *ServerAddParticleSpawner) String() string {
	return "{ServerAddParticleSpawner}"
}
