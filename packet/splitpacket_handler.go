package packet

type SplitpacketHandler struct {
	sessions       map[uint16]map[uint16]*SplitPayload
	sessions_count map[uint16]uint16
	seq_nr         uint16
}

func NewSplitPacketHandler() *SplitpacketHandler {
	return &SplitpacketHandler{
		sessions:       make(map[uint16]map[uint16]*SplitPayload),
		sessions_count: make(map[uint16]uint16),
		seq_nr:         65500,
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

func (sph *SplitpacketHandler) nextSequenceNr() uint16 {
	if sph.seq_nr >= 65535 {
		sph.seq_nr = 0
	} else {
		sph.seq_nr++
	}

	return sph.seq_nr
}

const MaxPacketLength = 495

func (sph *SplitpacketHandler) SplitPayload(payload []byte) ([]*Packet, error) {
	packets := make([]*Packet, 0)
	seqNr = sph.nextSequenceNr()

	parts := split(payload, MaxPacketLength)
	chunk_count := len(parts)

	for i, part := range parts {
		pkg := &Packet{
			PacketType: Reliable,
			SubType:    Split,
			SplitPayload: &SplitPayload{
				SeqNr:       seqNr,
				ChunkCount:  uint16(chunk_count),
				ChunkNumber: uint16(i),
				Data:        part,
			},
		}
		packets = append(packets, pkg)
	}

	return packets, nil
}
