// Copyright © 2019 Alexandre Kovac <contact@kovacou.fr>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package env

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUnexportedToBool(t *testing.T) {
	assert.True(t, toBool("1"))
	assert.True(t, toBool("true"))
	assert.True(t, toBool("a"))
	assert.True(t, toBool("t"))

	assert.False(t, toBool(""))
	assert.False(t, toBool("0"))
	assert.False(t, toBool("F"))
	assert.False(t, toBool("false"))
}

func TestUnexportedToInt64(t *testing.T) {
	assert.Equal(t, int64(15), toInt64("15"))
	assert.Equal(t, int64(1), toInt64("1"))
	assert.Equal(t, int64(0), toInt64("a"))
}

func TestUnexportedToUint64(t *testing.T) {
	assert.Equal(t, uint64(15), toUint64("15"))
	assert.Equal(t, uint64(1), toUint64("1"))
	assert.Equal(t, uint64(0), toUint64("a"))
}

func TestUnexportedToInt(t *testing.T) {
	assert.Equal(t, int(15), toInt("15"))
	assert.Equal(t, int(-2), toInt("-2"))
	assert.Equal(t, int(0), toInt("a"))
}

func TestUnexportedToDuration(t *testing.T) {
	{
		assert.Equal(t, time.Second, toDuration("1s"))
		assert.Equal(t, time.Hour, toDuration("60m"))
	}
}
