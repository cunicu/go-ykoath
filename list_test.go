// SPDX-FileCopyrightText: 2018 Joern Barthel <joern.barthel@kreuzwerker.de>
// SPDX-License-Identifier: Apache-2.0

package ykoath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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