package main

func shift(originalSlice []byte, key int8) (modifiedSlice []byte) {
	for _, symbol := range originalSlice {
		c := int8(symbol)
		if c == 10 {
			modifiedSlice = append(modifiedSlice, 10)
			continue
		}
		c = c + key
		if c > 126 {
			c = c - 95
		}
		if c < 32 {
			c = c + 95
		}
		modifiedSymbol := byte(c)
		modifiedSlice = append(modifiedSlice, modifiedSymbol)
	}
	return
}

func caesar(text []byte, key int8) (ciphertext []byte) {
	ciphertext = shift(text, key)
	return
}

func decoder(ciphertext []byte, key int8) (text []byte) {
	text = shift(ciphertext, -key)
	return
}
