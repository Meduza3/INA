package hamming

import "strconv"

func GetGreaterHamming(a, b *int) (greater, lesser *int) {
	ah := GetHammingWeight(*a)
	bh := GetHammingWeight(*b)
	if ah >= bh {
		return a, b
	} else {
		return b, a
	}
}

func GetHammingDistance(from, to int) int {
	fromBitstring := GetBit(from, -1)
	toBitstring := GetBit(to, -1)
	var distance int
	for i := 0; i < len(fromBitstring) || i < len(toBitstring); i++ {
		var fromBit, toBit byte = '0', '0'
		if i < len(fromBitstring) {
			fromBit = fromBitstring[len(fromBitstring)-1-i]
		}
		if i < len(toBitstring) {
			toBit = toBitstring[len(toBitstring)-1-i]
		}
		if fromBit != toBit {
			distance++
		}
	}
	return distance
}

func GetHammingWeight(num int) int {
	bitstring := GetBit(num, -1)
	var result int
	for _, c := range bitstring {
		if c == '1' {
			result++
		}
	}
	return result
}

func GetZeroWeight(num int) int {
	bitstring := GetBit(num, -1)
	var result int
	for _, c := range bitstring {
		if c == '0' {
			result++
		}
	}
	return result
}

func GetBit(num, length int) string {
	result := strconv.FormatInt(int64(num), 2)
	if length == -1 {
		return result
	}
	for len(result) < length {
		result = "0" + result
	}
	return result
}
