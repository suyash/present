package present

func addRoundKey(state, roundKey []int) {
	for i := 0; i < 64; i++ {
		state[i] = state[i] ^ roundKey[i]
	}
}

func sBox(n int) int {
	switch n {
	case 0:
		return 0xc
	case 1:
		return 5
	case 2:
		return 6
	case 3:
		return 0xb
	case 4:
		return 9
	case 5:
		return 0
	case 6:
		return 0xa
	case 7:
		return 0xd
	case 8:
		return 3
	case 9:
		return 0xe
	case 0xa:
		return 0xf
	case 0xb:
		return 8
	case 0xc:
		return 4
	case 0xd:
		return 7
	case 0xe:
		return 1
	case 0xf:
		return 2
	default:
		return n
	}
}

func sBoxLayer(state []int) {
	for i := 0; i < 16; i++ {
		val := 8*state[4*i] + 4*state[4*i+1] + 2*state[4*i+2] + state[4*i+3]
		sval := sBox(val)
		state[4*i+3], state[4*i+2], state[4*i+1], state[4*i] = sval&1, (sval>>1)&1, (sval>>2)&1, (sval>>3)&1
	}
}

func getPIndex(n int) int {
	row, col := n/16, n%16
	return ((col % 4) * 16) + (col / 4) + (row * 4)
}

func pLayer(state []int) {
	temp := make([]int, 64)

	for i := 0; i < 64; i++ {
		in := 63 - i
		temp[63-getPIndex(in)] = state[i]
	}

	for i := 0; i < 64; i++ {
		state[i] = temp[i]
	}
}

func getRoundKey(key []int) []int {
	ans := make([]int, 64)
	for i := 0; i < 64; i++ {
		ans[i] = key[i]
	}
	return ans
}

func rotateKey(key []int) {
	temp := make([]int, 80)

	for i := 0; i <= 18; i++ {
		temp[i] = key[i+61]
	}

	for i := 19; i < 80; i++ {
		temp[i] = key[i-19]
	}

	for i := 0; i < 80; i++ {
		key[i] = temp[i]
	}
}

func updateKey(key []int, roundCounter int) {
	rotateKey(key)

	v := 8*key[0] + 4*key[1] + 2*key[2] + key[3]
	sv := sBox(v)
	key[3], key[2], key[1], key[0] = sv&1, (sv>>1)&1, (sv>>2)&1, (sv>>3)&1

	v = key[79-15] + 2*key[79-16] + 4*key[79-17] + 8*key[79-18] + 16*key[79-19]
	sv = v ^ roundCounter
	key[79-15], key[79-16], key[79-17], key[79-18], key[79-19] = sv&1, (sv>>1)&1, (sv>>2)&1, (sv>>3)&1, (sv>>4)&1
}

// Encrypt encrypts the passed 64 bit state with the passed 80 bit key
func Encrypt(state, key []int) {
	for i := 1; i < 32; i++ {
		addRoundKey(state, getRoundKey(key))
		updateKey(key, i)
		sBoxLayer(state)
		pLayer(state)
	}
	addRoundKey(state, getRoundKey(key))
}
