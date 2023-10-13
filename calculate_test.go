// SPDX-FileCopyrightText: 2018 Joern Barthel <joern.barthel@kreuzwerker.de>
// SPDX-License-Identifier: Apache-2.0

package ykoath

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	assert := assert.New(t)

	for idx, k := range keys {
		var (
			testCard = new(testCard)
			touched  = false
			v        = vectors[k]
		)

		testCard.
			On(
				"Transmit",
				[]byte{
					0x00, 0xa4, 0x00, 0x01, 0x0a, 0x74, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x01,
				}).
			Return(
				[]byte{ //nolint:dupl // false-positive
					0x71, 0x2c, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x30, 0x31, 0x2d, 0x31, 0x65,
					0x35, 0x66, 0x32, 0x64, 0x62, 0x39, 0x2d, 0x34, 0x37, 0x37, 0x65, 0x2d,
					0x34, 0x31, 0x61, 0x66, 0x2d, 0x62, 0x64, 0x32, 0x65, 0x2d, 0x36, 0x30,
					0x62, 0x63, 0x35, 0x36, 0x39, 0x61, 0x65, 0x38, 0x37, 0x31, 0x76, 0x05,
					0x06, 0x00, 0x04, 0x61, 0x6a, 0x71, 0x2c, 0x74, 0x65, 0x73, 0x74, 0x2d,
					0x30, 0x32, 0x2d, 0x32, 0x61, 0x37, 0x63, 0x62, 0x63, 0x61, 0x39, 0x2d,
					0x62, 0x61, 0x65, 0x66, 0x2d, 0x34, 0x37, 0x65, 0x33, 0x2d, 0x38, 0x63,
					0x65, 0x38, 0x2d, 0x37, 0x38, 0x38, 0x62, 0x63, 0x36, 0x38, 0x35, 0x33,
					0x65, 0x31, 0x32, 0x7c, 0x01, 0x06, 0x71, 0x2c, 0x74, 0x65, 0x73, 0x74,
					0x2d, 0x30, 0x33, 0x2d, 0x62, 0x30, 0x31, 0x30, 0x31, 0x39, 0x65, 0x64,
					0x2d, 0x32, 0x61, 0x66, 0x31, 0x2d, 0x34, 0x38, 0x63, 0x63, 0x2d, 0x61,
					0x36, 0x34, 0x63, 0x2d, 0x66, 0x61, 0x39, 0x62, 0x34, 0x32, 0x34, 0x64,
					0x62, 0x39, 0x39, 0x33, 0x76, 0x05, 0x06, 0x00, 0x0a, 0x96, 0xb0, 0x71,
					0x2c, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x30, 0x34, 0x2d, 0x65, 0x36, 0x32,
					0x31, 0x37, 0x31, 0x66, 0x30, 0x2d, 0x34, 0x63, 0x66, 0x36, 0x2d, 0x34,
					0x39, 0x39, 0x65, 0x2d, 0x62, 0x39, 0x38, 0x38, 0x2d, 0x36, 0x65, 0x66,
					0x33, 0x36, 0x62, 0x32, 0x31, 0x33, 0x63, 0x63, 0x36, 0x7c, 0x01, 0x06,
					0x71, 0x2c, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x30, 0x35, 0x2d, 0x34, 0x35,
					0x38, 0x61, 0x66, 0x39, 0x65, 0x65, 0x2d, 0x63, 0x61, 0x61, 0x61, 0x2d,
					0x34, 0x37, 0x31, 0x36, 0x2d, 0x62, 0x66, 0x62, 0x38, 0x2d, 0x62, 0x64,
					0x38, 0x32, 0x38, 0x37, 0x35, 0x37, 0x39, 0x35, 0x35, 0x64, 0x76, 0x05,
					0x06, 0x00, 0x61, 0xff,
				},
				nil,
			).Once().
			On(
				"Transmit",
				[]byte{
					0x00, 0xa5, 0x00, 0x00,
				}).
			Return(
				[]byte{ //nolint:dupl // false-positive
					0x01, 0xd1, 0xce, 0x71, 0x2c, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x30, 0x36,
					0x2d, 0x32, 0x31, 0x33, 0x38, 0x61, 0x39, 0x39, 0x31, 0x2d, 0x65, 0x63,
					0x37, 0x30, 0x2d, 0x34, 0x38, 0x63, 0x62, 0x2d, 0x38, 0x33, 0x65, 0x36,
					0x2d, 0x66, 0x38, 0x30, 0x64, 0x61, 0x34, 0x37, 0x63, 0x39, 0x33, 0x65,
					0x34, 0x7c, 0x01, 0x06, 0x71, 0x2c, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x30,
					0x37, 0x2d, 0x61, 0x37, 0x30, 0x61, 0x32, 0x35, 0x32, 0x30, 0x2d, 0x37,
					0x65, 0x35, 0x31, 0x2d, 0x34, 0x35, 0x62, 0x32, 0x2d, 0x62, 0x61, 0x61,
					0x62, 0x2d, 0x30, 0x65, 0x33, 0x35, 0x32, 0x32, 0x30, 0x62, 0x30, 0x36,
					0x66, 0x65, 0x76, 0x05, 0x08, 0x05, 0x9e, 0xb4, 0xea, 0x71, 0x2c, 0x74,
					0x65, 0x73, 0x74, 0x2d, 0x30, 0x38, 0x2d, 0x38, 0x33, 0x66, 0x65, 0x33,
					0x32, 0x30, 0x38, 0x2d, 0x62, 0x31, 0x39, 0x32, 0x2d, 0x34, 0x36, 0x63,
					0x32, 0x2d, 0x39, 0x63, 0x62, 0x32, 0x2d, 0x31, 0x34, 0x65, 0x65, 0x39,
					0x31, 0x37, 0x62, 0x34, 0x64, 0x36, 0x30, 0x7c, 0x01, 0x08, 0x71, 0x2c,
					0x74, 0x65, 0x73, 0x74, 0x2d, 0x30, 0x39, 0x2d, 0x63, 0x63, 0x39, 0x64,
					0x31, 0x32, 0x32, 0x65, 0x2d, 0x39, 0x62, 0x35, 0x31, 0x2d, 0x34, 0x33,
					0x35, 0x65, 0x2d, 0x62, 0x34, 0x38, 0x65, 0x2d, 0x61, 0x62, 0x31, 0x61,
					0x31, 0x37, 0x31, 0x35, 0x37, 0x65, 0x33, 0x63, 0x76, 0x05, 0x08, 0x05,
					0x67, 0xe1, 0x30, 0x71, 0x2c, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x31, 0x30,
					0x2d, 0x39, 0x37, 0x61, 0x35, 0x38, 0x39, 0x33, 0x38, 0x2d, 0x38, 0x65,
					0x61, 0x36, 0x2d, 0x34, 0x31, 0x34, 0x33, 0x2d, 0x61, 0x65, 0x31, 0x30,
					0x2d, 0x38, 0x61, 0x64, 0x62, 0x39, 0x32, 0x62, 0x64, 0x63, 0x33, 0x33,
					0x35, 0x7c, 0x61, 0x68,
				},
				nil,
			).Once().
			On(
				"Transmit",
				[]byte{
					0x00, 0xa5, 0x00, 0x00,
				}).
			Return(
				[]byte{
					0x01, 0x08, 0x71, 0x2c, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x31, 0x31, 0x2d,
					0x38, 0x38, 0x37, 0x66, 0x64, 0x33, 0x38, 0x62, 0x2d, 0x38, 0x30, 0x62,
					0x33, 0x2d, 0x34, 0x64, 0x37, 0x61, 0x2d, 0x38, 0x36, 0x37, 0x31, 0x2d,
					0x38, 0x32, 0x62, 0x65, 0x66, 0x36, 0x33, 0x31, 0x35, 0x31, 0x61, 0x36,
					0x76, 0x05, 0x08, 0x02, 0xbf, 0xb9, 0x4e, 0x71, 0x2c, 0x74, 0x65, 0x73,
					0x74, 0x2d, 0x31, 0x32, 0x2d, 0x64, 0x61, 0x65, 0x65, 0x35, 0x30, 0x64,
					0x31, 0x2d, 0x37, 0x62, 0x62, 0x66, 0x2d, 0x34, 0x31, 0x65, 0x36, 0x2d,
					0x61, 0x36, 0x35, 0x62, 0x2d, 0x64, 0x33, 0x34, 0x30, 0x34, 0x36, 0x64,
					0x62, 0x61, 0x32, 0x38, 0x37, 0x7c, 0x01, 0x08, 0x90, 0x00,
				},
				nil,
			).Once()

		switch idx {
		case 1:

			testCard.On(
				"Transmit",
				[]byte{
					0x00, 0xa2, 0x00, 0x01, 0x38, 0x71, 0x2c, 0x74, 0x65, 0x73, 0x74, 0x2d,
					0x30, 0x32, 0x2d, 0x32, 0x61, 0x37, 0x63, 0x62, 0x63, 0x61, 0x39, 0x2d,
					0x62, 0x61, 0x65, 0x66, 0x2d, 0x34, 0x37, 0x65, 0x33, 0x2d, 0x38, 0x63,
					0x65, 0x38, 0x2d, 0x37, 0x38, 0x38, 0x62, 0x63, 0x36, 0x38, 0x35, 0x33,
					0x65, 0x31, 0x32, 0x74, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x01,
				}).
				Return(
					[]byte{
						0x76, 0x05, 0x06, 0x00, 0x01, 0xd1, 0xce, 0x90, 0x00,
					},
					nil,
				).Once()

		case 3:

			testCard.On(
				"Transmit",
				[]byte{
					0x00, 0xa2, 0x00, 0x01, 0x38, 0x71, 0x2c, 0x74, 0x65, 0x73, 0x74, 0x2d,
					0x30, 0x34, 0x2d, 0x65, 0x36, 0x32, 0x31, 0x37, 0x31, 0x66, 0x30, 0x2d,
					0x34, 0x63, 0x66, 0x36, 0x2d, 0x34, 0x39, 0x39, 0x65, 0x2d, 0x62, 0x39,
					0x38, 0x38, 0x2d, 0x36, 0x65, 0x66, 0x33, 0x36, 0x62, 0x32, 0x31, 0x33,
					0x63, 0x63, 0x36, 0x74, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x01,
				}).
				Return(
					[]byte{
						0x76, 0x05, 0x06, 0x00, 0x04, 0x61, 0x6a, 0x90, 0x00,
					},
					nil,
				).Once()

		case 5:

			testCard.On(
				"Transmit",
				[]byte{
					0x00, 0xa2, 0x00, 0x01, 0x38, 0x71, 0x2c, 0x74, 0x65, 0x73, 0x74, 0x2d,
					0x30, 0x36, 0x2d, 0x32, 0x31, 0x33, 0x38, 0x61, 0x39, 0x39, 0x31, 0x2d,
					0x65, 0x63, 0x37, 0x30, 0x2d, 0x34, 0x38, 0x63, 0x62, 0x2d, 0x38, 0x33,
					0x65, 0x36, 0x2d, 0x66, 0x38, 0x30, 0x64, 0x61, 0x34, 0x37, 0x63, 0x39,
					0x33, 0x65, 0x34, 0x74, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x01,
				}).
				Return(
					[]byte{
						0x76, 0x05, 0x06, 0x00, 0x0a, 0x96, 0xb0, 0x90, 0x00,
					},
					nil,
				).Once()

		case 7:

			testCard.On(
				"Transmit",
				[]byte{
					0x00, 0xa2, 0x00, 0x01, 0x38, 0x71, 0x2c, 0x74, 0x65, 0x73, 0x74, 0x2d,
					0x30, 0x38, 0x2d, 0x38, 0x33, 0x66, 0x65, 0x33, 0x32, 0x30, 0x38, 0x2d,
					0x62, 0x31, 0x39, 0x32, 0x2d, 0x34, 0x36, 0x63, 0x32, 0x2d, 0x39, 0x63,
					0x62, 0x32, 0x2d, 0x31, 0x34, 0x65, 0x65, 0x39, 0x31, 0x37, 0x62, 0x34,
					0x64, 0x36, 0x30, 0x74, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x01,
				}).
				Return(
					[]byte{
						0x76, 0x05, 0x08, 0x02, 0xbf, 0xb9, 0x4e, 0x90, 0x00,
					},
					nil,
				).Once()

		case 9:

			testCard.On(
				"Transmit",
				[]byte{
					0x00, 0xa2, 0x00, 0x01, 0x38, 0x71, 0x2c, 0x74, 0x65, 0x73, 0x74, 0x2d,
					0x31, 0x30, 0x2d, 0x39, 0x37, 0x61, 0x35, 0x38, 0x39, 0x33, 0x38, 0x2d,
					0x38, 0x65, 0x61, 0x36, 0x2d, 0x34, 0x31, 0x34, 0x33, 0x2d, 0x61, 0x65,
					0x31, 0x30, 0x2d, 0x38, 0x61, 0x64, 0x62, 0x39, 0x32, 0x62, 0x64, 0x63,
					0x33, 0x33, 0x35, 0x74, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x01,
				}).
				Return(
					[]byte{
						0x76, 0x05, 0x08, 0x05, 0x9e, 0xb4, 0xea, 0x90, 0x00,
					},
					nil,
				).Once()

		case 11:

			testCard.On(
				"Transmit",
				[]byte{
					0x00, 0xa2, 0x00, 0x01, 0x38, 0x71, 0x2c, 0x74, 0x65, 0x73, 0x74, 0x2d,
					0x31, 0x32, 0x2d, 0x64, 0x61, 0x65, 0x65, 0x35, 0x30, 0x64, 0x31, 0x2d,
					0x37, 0x62, 0x62, 0x66, 0x2d, 0x34, 0x31, 0x65, 0x36, 0x2d, 0x61, 0x36,
					0x35, 0x62, 0x2d, 0x64, 0x33, 0x34, 0x30, 0x34, 0x36, 0x64, 0x62, 0x61,
					0x32, 0x38, 0x37, 0x74, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x01,
				}).
				Return(
					[]byte{
						0x76, 0x05, 0x08, 0x05, 0x67, 0xe1, 0x30, 0x90, 0x00,
					},
					nil,
				).Once()
		}

		client := &OATH{
			card:     testCard,
			Timestep: DefaultTimeStep,
			Clock: func() time.Time {
				return time.Unix(v.time, 0)
			},
		}

		res, err := client.Calculate(k, func(_ string) error {
			touched = true
			return nil
		})
		assert.NoError(err)

		if v.touch {
			assert.True(touched)
		}

		assert.Equal(v.testvector, res)
		testCard.AssertExpectations(t)
	}
}