package utils

func IsMp3Format(data []byte) bool {
	if len(data) < 3 {
		return false
	}

	// Cek header ID3
	if string(data[:3]) == "ID3" {
		return true
	}

	// Cek frame header MPEG Audio (sync bits)
	return data[0] == 0xFF && (data[1]&0xE0) == 0xE0
}
