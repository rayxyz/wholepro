package algos

// Checksum : checksum
func Checksum(b []byte) uint16 {
	csumcv := len(b) - 1 // checksum coverage
	s := uint16(0)
	for i := 0; i < csumcv; i += 2 {
		// s += uint32(b[i+1])<<8 | uint32(b[i])
		s += uint16(b[i])
		s += s & uint16(b[i+1])
	}
	// if csumcv&1 == 0 {
	// 	s += uint32(b[csumcv])
	// }
	// s = s>>16 + s&0xffff
	// s = s + s>>16

	return ^s
}
