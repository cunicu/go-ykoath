// SPDX-FileCopyrightText: 2018 Joern Barthel <joern.barthel@kreuzwerker.de>
// SPDX-License-Identifier: Apache-2.0

package ykoath

import (
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/ebfe/scard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCard struct {
	mock.Mock
}

func (t *testCard) Disconnect(d scard.Disposition) error {
	args := t.Called(d)
	return args.Error(0)
}

func (t *testCard) Transmit(b []byte) ([]byte, error) {
	args := t.Called(b)
	return args.Get(0).([]byte), args.Error(1) //nolint:forcetypeassert
}

type vector struct {
	a          Algorithm
	digits     uint8
	key        []byte
	name       string
	t          Type
	testvector string
	time       int64
	touch      bool
}

var (
	keys    sort.StringSlice
	vectors map[string]*vector
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

func TestList(t *testing.T) {
	var (
		assert   = assert.New(t)
		testCard = new(testCard)
	)

	testCard.
		On(
			"Transmit",
			[]byte{
				0x00, 0xa1, 0x00, 0x00,
			}).
		Return(
			[]byte{ //nolint:dupl // false-positive
				0x72, 0x2d, 0x21, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x30, 0x31, 0x2d, 0x31,
				0x65, 0x35, 0x66, 0x32, 0x64, 0x62, 0x39, 0x2d, 0x34, 0x37, 0x37, 0x65,
				0x2d, 0x34, 0x31, 0x61, 0x66, 0x2d, 0x62, 0x64, 0x32, 0x65, 0x2d, 0x36,
				0x30, 0x62, 0x63, 0x35, 0x36, 0x39, 0x61, 0x65, 0x38, 0x37, 0x31, 0x72,
				0x2d, 0x22, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x30, 0x32, 0x2d, 0x32, 0x61,
				0x37, 0x63, 0x62, 0x63, 0x61, 0x39, 0x2d, 0x62, 0x61, 0x65, 0x66, 0x2d,
				0x34, 0x37, 0x65, 0x33, 0x2d, 0x38, 0x63, 0x65, 0x38, 0x2d, 0x37, 0x38,
				0x38, 0x62, 0x63, 0x36, 0x38, 0x35, 0x33, 0x65, 0x31, 0x32, 0x72, 0x2d,
				0x23, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x30, 0x33, 0x2d, 0x62, 0x30, 0x31,
				0x30, 0x31, 0x39, 0x65, 0x64, 0x2d, 0x32, 0x61, 0x66, 0x31, 0x2d, 0x34,
				0x38, 0x63, 0x63, 0x2d, 0x61, 0x36, 0x34, 0x63, 0x2d, 0x66, 0x61, 0x39,
				0x62, 0x34, 0x32, 0x34, 0x64, 0x62, 0x39, 0x39, 0x33, 0x72, 0x2d, 0x21,
				0x74, 0x65, 0x73, 0x74, 0x2d, 0x30, 0x34, 0x2d, 0x65, 0x36, 0x32, 0x31,
				0x37, 0x31, 0x66, 0x30, 0x2d, 0x34, 0x63, 0x66, 0x36, 0x2d, 0x34, 0x39,
				0x39, 0x65, 0x2d, 0x62, 0x39, 0x38, 0x38, 0x2d, 0x36, 0x65, 0x66, 0x33,
				0x36, 0x62, 0x32, 0x31, 0x33, 0x63, 0x63, 0x36, 0x72, 0x2d, 0x22, 0x74,
				0x65, 0x73, 0x74, 0x2d, 0x30, 0x35, 0x2d, 0x34, 0x35, 0x38, 0x61, 0x66,
				0x39, 0x65, 0x65, 0x2d, 0x63, 0x61, 0x61, 0x61, 0x2d, 0x34, 0x37, 0x31,
				0x36, 0x2d, 0x62, 0x66, 0x62, 0x38, 0x2d, 0x62, 0x64, 0x38, 0x32, 0x38,
				0x37, 0x35, 0x37, 0x39, 0x35, 0x35, 0x64, 0x72, 0x2d, 0x23, 0x74, 0x65,
				0x73, 0x74, 0x2d, 0x30, 0x36, 0x2d, 0x32, 0x31, 0x33, 0x38, 0x61, 0x39,
				0x39, 0x31, 0x61, 0xff,
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
				0x2d, 0x65, 0x63, 0x37, 0x30, 0x2d, 0x34, 0x38, 0x63, 0x62, 0x2d, 0x38,
				0x33, 0x65, 0x36, 0x2d, 0x66, 0x38, 0x30, 0x64, 0x61, 0x34, 0x37, 0x63,
				0x39, 0x33, 0x65, 0x34, 0x72, 0x2d, 0x21, 0x74, 0x65, 0x73, 0x74, 0x2d,
				0x30, 0x37, 0x2d, 0x61, 0x37, 0x30, 0x61, 0x32, 0x35, 0x32, 0x30, 0x2d,
				0x37, 0x65, 0x35, 0x31, 0x2d, 0x34, 0x35, 0x62, 0x32, 0x2d, 0x62, 0x61,
				0x61, 0x62, 0x2d, 0x30, 0x65, 0x33, 0x35, 0x32, 0x32, 0x30, 0x62, 0x30,
				0x36, 0x66, 0x65, 0x72, 0x2d, 0x22, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x30,
				0x38, 0x2d, 0x38, 0x33, 0x66, 0x65, 0x33, 0x32, 0x30, 0x38, 0x2d, 0x62,
				0x31, 0x39, 0x32, 0x2d, 0x34, 0x36, 0x63, 0x32, 0x2d, 0x39, 0x63, 0x62,
				0x32, 0x2d, 0x31, 0x34, 0x65, 0x65, 0x39, 0x31, 0x37, 0x62, 0x34, 0x64,
				0x36, 0x30, 0x72, 0x2d, 0x23, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x30, 0x39,
				0x2d, 0x63, 0x63, 0x39, 0x64, 0x31, 0x32, 0x32, 0x65, 0x2d, 0x39, 0x62,
				0x35, 0x31, 0x2d, 0x34, 0x33, 0x35, 0x65, 0x2d, 0x62, 0x34, 0x38, 0x65,
				0x2d, 0x61, 0x62, 0x31, 0x61, 0x31, 0x37, 0x31, 0x35, 0x37, 0x65, 0x33,
				0x63, 0x72, 0x2d, 0x21, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x31, 0x30, 0x2d,
				0x39, 0x37, 0x61, 0x35, 0x38, 0x39, 0x33, 0x38, 0x2d, 0x38, 0x65, 0x61,
				0x36, 0x2d, 0x34, 0x31, 0x34, 0x33, 0x2d, 0x61, 0x65, 0x31, 0x30, 0x2d,
				0x38, 0x61, 0x64, 0x62, 0x39, 0x32, 0x62, 0x64, 0x63, 0x33, 0x33, 0x35,
				0x72, 0x2d, 0x22, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x31, 0x31, 0x2d, 0x38,
				0x38, 0x37, 0x66, 0x64, 0x33, 0x38, 0x62, 0x2d, 0x38, 0x30, 0x62, 0x33,
				0x2d, 0x34, 0x64, 0x37, 0x61, 0x2d, 0x38, 0x36, 0x37, 0x31, 0x2d, 0x38,
				0x32, 0x62, 0x61, 0x38,
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
				0x65, 0x66, 0x36, 0x33, 0x31, 0x35, 0x31, 0x61, 0x36, 0x72, 0x2d, 0x23,
				0x74, 0x65, 0x73, 0x74, 0x2d, 0x31, 0x32, 0x2d, 0x64, 0x61, 0x65, 0x65,
				0x35, 0x30, 0x64, 0x31, 0x2d, 0x37, 0x62, 0x62, 0x66, 0x2d, 0x34, 0x31,
				0x65, 0x36, 0x2d, 0x61, 0x36, 0x35, 0x62, 0x2d, 0x64, 0x33, 0x34, 0x30,
				0x34, 0x36, 0x64, 0x62, 0x61, 0x32, 0x38, 0x37, 0x90, 0x00,
			},
			nil,
		).Once()

	client := &OATH{
		card:     testCard,
		Timestep: DefaultTimeStep,
	}

	res, err := client.List()
	assert.NoError(err)
	assert.Len(res, len(vectors))

	for idx, r := range res {
		name := keys[idx]

		assert.Equal(vectors[name].a, r.Algorithm)
		assert.Equal(vectors[name].name, r.Name)
		assert.Equal(vectors[name].t, r.Type)
	}

	testCard.AssertExpectations(t)
}

func TestPutAndCalculateTestVector(t *testing.T) {
	tt := []struct {
		Name  string
		Query string
	}{
		{
			"full identifier",
			"testvector",
		},
		{
			"name only (substring)",
			"test",
		},
	}

	for _, test := range tt {
		t.Run(test.Name, func(t *testing.T) {
			var (
				assert   = assert.New(t)
				testCard = new(testCard)
			)

			testCard.
				On(
					"Transmit",
					[]byte{
						0x00, 0x01, 0x00, 0x00, 0x24, 0x71, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x76,
						0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x16, 0x21, 0x08, 0x31, 0x32, 0x33,
						0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35,
						0x36, 0x37, 0x38, 0x39, 0x30,
					}).
				Return(
					[]byte{
						0x90, 0x00,
					},
					nil,
				).Once().
				On(
					"Transmit",
					[]byte{
						0x00, 0xa4, 0x00, 0x01, 0x0a, 0x74, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00,
						0x00, 0x00, 0x01,
					}).
				Return(
					[]byte{
						0x71, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72,
						0x76, 0x05, 0x08, 0x05, 0x9e, 0xb4, 0xea, 0x90, 0x00,
					},
					nil,
				).Once()

			client := &OATH{
				card:     testCard,
				Timestep: DefaultTimeStep,
				Clock: func() time.Time {
					return time.Unix(59, 0)
				},
			}

			err := client.Put("testvector", HmacSha1, Totp, 8, []byte("12345678901234567890"), false)
			assert.NoError(err)

			res, err := client.Calculate(test.Query, nil)
			assert.NoError(err)
			assert.Equal("94287082", res)

			testCard.AssertExpectations(t)
		})
	}

	t.Run("multiple match error", func(t *testing.T) {
		var (
			assert   = assert.New(t)
			testCard = new(testCard)
		)

		testCard.
			On(
				"Transmit",
				[]byte{
					0x00, 0x01, 0x00, 0x00, 0x25, 0x71, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x76,
					0x65, 0x63, 0x74, 0x6f, 0x72, 0x31, 0x73, 0x16, 0x21, 0x08, 0x31, 0x32,
					0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x31, 0x32, 0x33, 0x34,
					0x35, 0x36, 0x37, 0x38, 0x39, 0x30,
				}).
			Return(
				[]byte{
					0x90, 0x00,
				},
				nil,
			).Once().
			On(
				"Transmit",
				[]byte{
					0x00, 0x01, 0x00, 0x00, 0x25, 0x71, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x76,
					0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x73, 0x16, 0x21, 0x08, 0x31, 0x32,
					0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x31, 0x32, 0x33, 0x34,
					0x35, 0x36, 0x37, 0x38, 0x39, 0x30,
				}).
			Return(
				[]byte{
					0x90, 0x00,
				},
				nil,
			).Once().
			On(
				"Transmit",
				[]byte{
					0x00, 0xa4, 0x00, 0x01, 0x0a, 0x74, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x01,
				}).
			Return(
				[]byte{
					0x71, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x31,
					0x76, 0x05, 0x08, 0x05, 0x9e, 0xb4, 0xea,
					0x71, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32,
					0x76, 0x05, 0x08, 0x05, 0x9e, 0xb4, 0xea,
					0x90, 0x00,
				},
				nil,
			).Once()

		client := &OATH{
			card:     testCard,
			Timestep: DefaultTimeStep,
			Clock: func() time.Time {
				return time.Unix(59, 0)
			},
		}

		err := client.Put("testvector1", HmacSha1, Totp, 8, []byte("12345678901234567890"), false)
		assert.NoError(err)

		err = client.Put("testvector2", HmacSha1, Totp, 8, []byte("12345678901234567890"), false)
		assert.NoError(err)

		_, err = client.Calculate("test", nil)
		assert.ErrorIs(err, errMultipleMatches)

		testCard.AssertExpectations(t)
	})
}

func TestSelectTOTP(t *testing.T) {
	var (
		assert   = assert.New(t)
		testCard = new(testCard)
	)

	testCard.
		On(
			"Transmit",
			[]byte{
				0x00, 0xa4, 0x04, 0x00, 0x07, 0xa0, 0x00, 0x00, 0x05, 0x27, 0x21, 0x01,
			}).
		Return(
			[]byte{
				0x79, 0x03, 0x04, 0x03, 0x03, 0x71, 0x08, 0x7c, 0x06, 0x60, 0x15, 0x20,
				0xfc, 0x3f, 0x8f, 0x90, 0x00,
			},
			nil,
		)

	client := &OATH{
		Timestep: DefaultTimeStep,
		card:     testCard,
	}

	res, err := client.Select()

	assert.Empty(res.Algorithm)
	assert.Empty(res.Challenge)
	assert.Equal(fmt.Sprintf("% x", []byte{0x7c, 0x06, 0x60, 0x15, 0x20, 0xfc, 0x3f, 0x8f}), fmt.Sprintf("% x", res.Name))
	assert.Equal(fmt.Sprintf("% x", []byte{0x04, 0x03, 0x03}), fmt.Sprintf("% x", res.Version))

	assert.NoError(err)

	testCard.AssertExpectations(t)
}

func init() { //nolint:gochecknoinits
	vectors = map[string]*vector{
		"test-01-1e5f2db9-477e-41af-bd2e-60bc569ae871": {
			a:          HmacSha1,
			t:          Totp,
			digits:     6,
			key:        []byte("12345678901234567890"),
			touch:      false,
			time:       59,
			testvector: "287082",
		},
		"test-02-2a7cbca9-baef-47e3-8ce8-788bc6853e12": {
			a:          HmacSha256,
			t:          Totp,
			digits:     6,
			key:        []byte("12345678901234567890123456789012"),
			touch:      true,
			time:       59,
			testvector: "119246",
		},
		"test-03-b01019ed-2af1-48cc-a64c-fa9b424db993": {
			a:          HmacSha512,
			t:          Totp,
			digits:     6,
			key:        []byte("1234567890123456789012345678901234567890123456789012345678901234"),
			touch:      false,
			time:       59,
			testvector: "693936",
		},
		"test-04-e62171f0-4cf6-499e-b988-6ef36b213cc6": {
			a:          HmacSha1,
			t:          Totp,
			digits:     6,
			key:        []byte("12345678901234567890"),
			touch:      true,
			time:       59,
			testvector: "287082",
		},
		"test-05-458af9ee-caaa-4716-bfb8-bd828757955d": {
			a:          HmacSha256,
			t:          Totp,
			digits:     6,
			key:        []byte("12345678901234567890123456789012"),
			touch:      false,
			time:       59,
			testvector: "119246",
		},
		"test-06-2138a991-ec70-48cb-83e6-f80da47c93e4": {
			a:          HmacSha512,
			t:          Totp,
			digits:     6,
			key:        []byte("1234567890123456789012345678901234567890123456789012345678901234"),
			touch:      true,
			time:       59,
			testvector: "693936",
		},
		"test-07-a70a2520-7e51-45b2-baab-0e35220b06fe": {
			a:          HmacSha1,
			t:          Totp,
			digits:     8,
			key:        []byte("12345678901234567890"),
			touch:      false,
			time:       59,
			testvector: "94287082",
		},
		"test-08-83fe3208-b192-46c2-9cb2-14ee917b4d60": {
			a:          HmacSha256,
			t:          Totp,
			digits:     8,
			key:        []byte("12345678901234567890123456789012"),
			touch:      true,
			time:       59,
			testvector: "46119246",
		},
		"test-09-cc9d122e-9b51-435e-b48e-ab1a17157e3c": {
			a:          HmacSha512,
			t:          Totp,
			digits:     8,
			key:        []byte("1234567890123456789012345678901234567890123456789012345678901234"),
			touch:      false,
			time:       59,
			testvector: "90693936",
		},
		"test-10-97a58938-8ea6-4143-ae10-8adb92bdc335": {
			a:          HmacSha1,
			t:          Totp,
			digits:     8,
			key:        []byte("12345678901234567890"),
			touch:      true,
			time:       59,
			testvector: "94287082",
		},
		"test-11-887fd38b-80b3-4d7a-8671-82bef63151a6": {
			a:          HmacSha256,
			t:          Totp,
			digits:     8,
			key:        []byte("12345678901234567890123456789012"),
			touch:      false,
			time:       59,
			testvector: "46119246",
		},
		"test-12-daee50d1-7bbf-41e6-a65b-d34046dba287": {
			a:          HmacSha512,
			t:          Totp,
			digits:     8,
			key:        []byte("1234567890123456789012345678901234567890123456789012345678901234"),
			touch:      true,
			time:       59,
			testvector: "90693936",
		},
	}

	for k, v := range vectors {
		keys = append(keys, k)
		v.name = k
	}

	keys.Sort()
}
