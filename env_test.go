// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package env

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	os.Setenv("SLICE_STRING", "v1,v2,v3")
	os.Setenv("STRING", "value")
	os.Setenv("STRING_PTR", "value_ptr")
	os.Setenv("INT64", "256")
	os.Setenv("FLOAT64", "2.25")
	os.Setenv("FLOAT32", "2.1")
	os.Setenv("INT", "123")
	os.Setenv("UINT", "456")
	os.Setenv("UINT64", "1337")
	os.Setenv("DURATION", "60s")

	{
		s := struct {
			Embeded struct {
				StringPtr *string `env:"STRING_PTR"`
			}
			Duration    time.Duration `env:"DURATION"`
			String      string        `env:"STRING"`
			SliceString []string      `env:"SLICE_STRING"`
			Float64     float64       `env:"FLOAT64"`
			Float32     float32       `env:"FLOAT32"`
			Int64       int64         `env:"INT64"`
			Int         int           `env:"INT"`
			Uint        uint          `env:"UINT"`
			Uint64      uint64        `env:"UINT64"`
		}{}

		Unmarshal(&s)

		assert.Equal(t, "value", s.String)
		assert.Equal(t, "value_ptr", *s.Embeded.StringPtr)
		assert.Equal(t, float64(2.25), s.Float64)
		assert.Equal(t, float32(2.1), s.Float32)
		assert.Equal(t, int64(256), s.Int64)
		assert.Equal(t, int(123), s.Int)
		assert.Equal(t, uint(456), s.Uint)
		assert.Equal(t, uint64(1337), s.Uint64)
		assert.Equal(t, "1m0s", s.Duration.String())
		assert.Equal(t, s.SliceString, []string{"v1", "v2", "v3"})
	}
}
