package present

import "testing"

func getBitVector(s string) []int {
	ans := make([]int, len(s)*4)
	for i := 0; i < len(s); i++ {
		var v int
		if s[i] >= '0' && s[i] <= '9' {
			v = int(s[i] - '0')
		} else {
			v = int(s[i] - 'A' + 10)
		}
		ans[4*i+3], ans[4*i+2], ans[4*i+1], ans[4*i] = v&1, (v>>1)&1, (v>>2)&1, (v>>3)&1
	}
	return ans
}

func getString(state []int) string {
	sl := len(state) / 4
	a := make([]byte, sl)
	for i := 0; i < sl; i++ {
		v := 8*state[4*i] + 4*state[4*i+1] + 2*state[4*i+2] + state[4*i+3]
		var c byte
		if v <= 9 {
			c = '0' + byte(v)
		} else {
			c = 'A' + byte(v-10)
		}
		a[i] = c
	}
	return string(a)
}

func TestEncrypt(t *testing.T) {
	cases := map[string]struct {
		message, key, cipher string
	}{
		"appendix case 1": {"0000000000000000", "00000000000000000000", "5579C1387B228445"},
		"appendix case 2": {"0000000000000000", "FFFFFFFFFFFFFFFFFFFF", "E72C46C0F5945049"},
		"appendix case 3": {"FFFFFFFFFFFFFFFF", "00000000000000000000", "A112FFC72F68417B"},
		"appendix case 4": {"FFFFFFFFFFFFFFFF", "FFFFFFFFFFFFFFFFFFFF", "3333DCD3213210D2"},
	}

	for name, c := range cases {
		state := getBitVector(c.message)
		Encrypt(state, getBitVector(c.key))
		ans := getString(state)
		if ans != c.cipher {
			t.Errorf("case %s, wanted %v, got: %v", name, c.cipher, ans)
		}
	}
}
