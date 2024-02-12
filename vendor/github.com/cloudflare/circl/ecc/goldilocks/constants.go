// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package goldilocks

import fp "github.com/cloudflare/circl/math/fp448"

var (
	// genX is the x-coordinate of the generator of Goldilocks curve.
	genX = fp.Elt{
		0x5e, 0xc0, 0x0c, 0xc7, 0x2b, 0xa8, 0x26, 0x26,
		0x8e, 0x93, 0x00, 0x8b, 0xe1, 0x80, 0x3b, 0x43,
		0x11, 0x65, 0xb6, 0x2a, 0xf7, 0x1a, 0xae, 0x12,
		0x64, 0xa4, 0xd3, 0xa3, 0x24, 0xe3, 0x6d, 0xea,
		0x67, 0x17, 0x0f, 0x47, 0x70, 0x65, 0x14, 0x9e,
		0xda, 0x36, 0xbf, 0x22, 0xa6, 0x15, 0x1d, 0x22,
		0xed, 0x0d, 0xed, 0x6b, 0xc6, 0x70, 0x19, 0x4f,
	}
	// genY is the y-coordinate of the generator of Goldilocks curve.
	genY = fp.Elt{
		0x14, 0xfa, 0x30, 0xf2, 0x5b, 0x79, 0x08, 0x98,
		0xad, 0xc8, 0xd7, 0x4e, 0x2c, 0x13, 0xbd, 0xfd,
		0xc4, 0x39, 0x7c, 0xe6, 0x1c, 0xff, 0xd3, 0x3a,
		0xd7, 0xc2, 0xa0, 0x05, 0x1e, 0x9c, 0x78, 0x87,
		0x40, 0x98, 0xa3, 0x6c, 0x73, 0x73, 0xea, 0x4b,
		0x62, 0xc7, 0xc9, 0x56, 0x37, 0x20, 0x76, 0x88,
		0x24, 0xbc, 0xb6, 0x6e, 0x71, 0x46, 0x3f, 0x69,
	}
	// paramD is -39081 in Fp.
	paramD = fp.Elt{
		0x56, 0x67, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xfe, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	}
	// order is 2^446-0x8335dc163bb124b65129c96fde933d8d723a70aadc873d6d54a7bb0d,
	// which is the number of points in the prime subgroup.
	order = Scalar{
		0xf3, 0x44, 0x58, 0xab, 0x92, 0xc2, 0x78, 0x23,
		0x55, 0x8f, 0xc5, 0x8d, 0x72, 0xc2, 0x6c, 0x21,
		0x90, 0x36, 0xd6, 0xae, 0x49, 0xdb, 0x4e, 0xc4,
		0xe9, 0x23, 0xca, 0x7c, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x3f,
	}
	// residue448 is 2^448 mod order.
	residue448 = [4]uint64{
		0x721cf5b5529eec34, 0x7a4cf635c8e9c2ab, 0xeec492d944a725bf, 0x20cd77058,
	}
	// invFour is 1/4 mod order.
	invFour = Scalar{
		0x3d, 0x11, 0xd6, 0xaa, 0xa4, 0x30, 0xde, 0x48,
		0xd5, 0x63, 0x71, 0xa3, 0x9c, 0x30, 0x5b, 0x08,
		0xa4, 0x8d, 0xb5, 0x6b, 0xd2, 0xb6, 0x13, 0x71,
		0xfa, 0x88, 0x32, 0xdf, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0f,
	}
	// paramDTwist is -39082 in Fp. The D parameter of the twist curve.
	paramDTwist = fp.Elt{
		0x55, 0x67, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xfe, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	}
)
