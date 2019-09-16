// Copyright © 2019 Alexandre Kovac <contact@kovacou.fr>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnexportedToBool(t *testing.T) {
	assert := assert.New(t)
	assert.True(toBool("1"))
	assert.True(toBool("true"))
	assert.True(toBool("a"))
	assert.True(toBool("t"))

	assert.False(toBool(""))
	assert.False(toBool("0"))
	assert.False(toBool("F"))
	assert.False(toBool("false"))
}

func TestUnexportedToInt64(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(int64(15), toInt64("15"))
	assert.Equal(int64(1), toInt64("1"))
	assert.Equal(int64(0), toInt64("a"))
}

func TestUnexportedToUint64(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(uint64(15), toUint64("15"))
	assert.Equal(uint64(1), toUint64("1"))
	assert.Equal(uint64(0), toUint64("a"))
}
