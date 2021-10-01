package packet

type SplitpacketHandler struct {
	sessions       map[uint16]map[uint16]*SplitPayload
	sessions_count map[uint16]uint16
}

func NewSplitPacketHandler() *SplitpacketHandler {
	return &SplitpacketHandler{
		sessions:       make(map[uint16]map[uint16]*SplitPayload),
		sessions_count: make(map[uint16]uint16),
	}
}

func (sph *SplitpacketHandler) AddPacket(sp *SplitPayload) []byte {

	parts := sph.sessions[sp.SeqNr]
	if parts == nil {
		parts = make(map[uint16]*SplitPayload)
		sph.sessions[sp.SeqNr] = parts
	}

	if parts[sp.ChunkNumber] == nil {
		parts[sp.ChunkNumber] = sp
		sph.sessions_count[sp.SeqNr]++

		if sph.sessions_count[sp.SeqNr] == sp.ChunkCount {
			// packet complete
			buf := []byte{}
			for i := 0; i < int(sp.ChunkCount); i++ {
				buf = append(buf, parts[uint16(i)].Data...)
			}

			sph.sessions[sp.SeqNr] = nil
			sph.sessions_count[sp.SeqNr] = 0
			return buf
		}
	}

	return nil
}
