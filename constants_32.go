package ed448

const (
	lmask    = 0xffffffff
	allZeros = word(0x0)
	allOnes  = word(0xffffffff)

	decafTrue  = word(0xffffffff)
	decafFalse = word(0x0)

	limbs       = 16
	scalarLimbs = 14
	radix       = 28
	radixMask   = word(0xfffffff)

	// The size of the Goldilocks field, in bits.
	fieldBits = 448
	edwardsD  = -39081

	// The size of the Goldilocks field, in bytes.
	fieldBytes = fieldBits / 8 // 56

	// The size of the Goldilocks scalars, in bits.
	scalarBits = fieldBits - 2 // 446
	// The size of the Goldilocks field, in bytes.
	scalarBytes = (scalarBits + 7) / 8 // 56

	wordBits = 32 // 32-bits
	//wordBits = 64 // 64-bits

	// The number of words in the Goldilocks field.
	// 14 for 32-bit and 7 for 64-bits
	scalarWords = (scalarBits + wordBits - 1) / wordBits

	bitSize  = scalarBits
	byteSize = fieldBytes

	symKeyBytes  = 32
	pubKeyBytes  = fieldBytes
	privKeyBytes = 2*fieldBytes + symKeyBytes // 144

	signatureBytes = 2 * fieldBytes // 112

	// Comb configuration
	combNumber  = uint(8)
	combTeeth   = uint(4)
	combSpacing = uint(14)

	// DecafComb configuration
	decafCombNumber  = uint(5)
	decafCombTeeth   = uint(5)
	decafCombSpacing = uint(18)

	uintZero = uint(0)

	montgomeryFactor = word(0xae918bc5)
)

var (
	bigZero = &bigNumber{0}
	bigOne  = &bigNumber{1}
	bigTwo  = &bigNumber{2}

	sqrtDminus1 = mustDeserialize(serialized{
		0x46, 0x9f, 0x74, 0x36, 0x18, 0xe2, 0xd2, 0x79,
		0x01, 0x4f, 0x2b, 0xb4, 0x8d, 0x88, 0x38, 0xea,
		0xde, 0xab, 0x9a, 0x18, 0x5a, 0x06, 0x4c, 0xf1,
		0xa6, 0x5c, 0xe6, 0x51, 0x70, 0x97, 0x4d, 0x42,
		0x7b, 0x9f, 0xa4, 0x56, 0xf6, 0xc5, 0x28, 0x46,
		0xac, 0xdc, 0x4a, 0x73, 0x48, 0x87, 0x3b, 0x44,
		0x49, 0x7a, 0x5b, 0xb2, 0xc0, 0xc0, 0xfe, 0x12,
	})

	//ScalarQ is the prime order of the curve.
	ScalarQ = &decafScalar{
		0xab5844f3, 0x2378c292,
		0x8dc58f55, 0x216cc272,
		0xaed63690, 0xc44edb49,
		0x7cca23e9, 0xffffffff,
		0xffffffff, 0xffffffff,
		0xffffffff, 0xffffffff,
		0xffffffff, 0x3fffffff,
	}

	scalarR2 = &decafScalar{
		0x049b9b60, 0xe3539257,
		0xc1b195d9, 0x7af32c4b,
		0x88ea1859, 0x0d66de23,
		0x5ee4d838, 0xae17cf72,
		0xa3c47c44, 0x1a9cc14b,
		0xe4d070af, 0x2052bcb7,
		0xf823b729, 0x3402a939,
	}

	//BasePoint is the basepoint of the curve.
	BasePoint = &twExtendedPoint{
		&bigNumber{
			0x0ffffffe, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x00000003, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
		},
		&bigNumber{
			0x0f752992, 0x081e6d37,
			0x01c28721, 0x03078ead,
			0x0394666c, 0x0135cfd2,
			0x00506061, 0x041149c5,
			0x0f5490b3, 0x031d30e4,
			0x090dc141, 0x09020149,
			0x04c1e328, 0x052341b0,
			0x03c10a1b, 0x01423785,
		},
		&bigNumber{
			0x0ffffffb, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0ffffffe, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
		},
		&bigNumber{
			0x00660415, 0x08f205b7,
			0x0fd3824f, 0x0881c60c,
			0x0d08500d, 0x0377a638,
			0x04672615, 0x08c66d5d,
			0x08e08e13, 0x0e52fa55,
			0x01b6983d, 0x087770ae,
			0x0a0aa7ff, 0x04388f55,
			0x05cf1a91, 0x0b4d9a78,
		},
	}
)
