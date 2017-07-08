package encode

import "testing"

func TestToUnary(t *testing.T) {

	var cases = []struct {
		in  int
		out string
	}{
		{1, "1"},
		{2, "01"},
		{3, "001"},
		{4, "0001"},
		{5, "00001"},
		{10, "0000000001"},
		{100, "0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001"},
	}

	for _, c := range cases {
		s := ToUnary(c.in)
		if s != c.out {
			t.Errorf("ToUnary(%v) => %v, want %v", c.in, s, c.out)
		}
	}
}

func TestToBinary(t *testing.T) {

	var cases = []struct {
		in  int
		out string
	}{
		{1, "1"},
		{2, "10"},
		{3, "11"},
		{4, "100"},
		{5, "101"},
		{10, "1010"},
		{100, "1100100"},
		{1234567890, "1001001100101100000001011010010"},
	}

	for _, c := range cases {
		s := ToBinary(c.in)
		if s != c.out {
			t.Errorf("ToBinary(%v) => %v, want %v", c.in, s, c.out)
		}
	}
}

func TestToGamma(t *testing.T) {

	var cases = []struct {
		in  int
		out string
	}{
		{1, "1"},
		{2, "010"},
		{3, "011"},
		{4, "00100"},
		{5, "00101"},
		{10, "0001010"},
		{100, "0000001100100"},
	}

	for _, c := range cases {
		s := ToGamma(c.in)
		if s != c.out {
			t.Errorf("ToGamma(%v) => %v, want %v", c.in, s, c.out)
		}
	}
}

func TestToDelta(t *testing.T) {

	var cases = []struct {
		in  int
		out string
	}{
		{1, "1"},
		{2, "0100"},
		{3, "0101"},
		{4, "01100"},
		{5, "01101"},
		{10, "00100010"},
		{100, "00111100100"},
	}

	for _, c := range cases {
		s := ToDelta(c.in)
		if s != c.out {
			t.Errorf("ToDelta(%v) => %v, want %v", c.in, s, c.out)
		}
	}
}

func TestToGolomb(t *testing.T) {

	var cases = []struct {
		in  int
		k   int
		out string
	}{
		{1, 8, "1000"},
		{2, 8, "1001"},
		{3, 8, "1010"},
		{4, 8, "1011"},
		{5, 8, "1100"},
		{10, 8, "01001"},
		{100, 8, "0000000000001011"},
	}

	for _, c := range cases {
		s := ToGolomb(c.in, c.k)
		if s != c.out {
			t.Errorf("ToGolomb(%v) => %v, want %v", c.in, s, c.out)
		}
	}
}
