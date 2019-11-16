// Copyright © 2019 Alexandre Kovac <contact@kovacou.fr>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package env

import (
	"strconv"
	"time"
)

func toBool(v string) bool {
	out, err := strconv.ParseBool(v)
	if err != nil && len(v) > 0 {
		out = true
	}
	return out
}

func toInt64(v string) int64 {
	out, _ := strconv.ParseInt(v, 0, 0)
	return out
}

func toInt(v string) int {
	return int(toInt64(v))
}

func toUint(v string) uint {
	return uint(toInt64(v))
}

func toUint64(v string) uint64 {
	return uint64(toInt64(v))
}

func toFloat32(v string) float32 {
	out, _ := strconv.ParseFloat(v, 32)
	return float32(out)
}

func toFloat64(v string) float64 {
	out, _ := strconv.ParseFloat(v, 64)
	return out
}

func toDuration(v string) time.Duration {
	out, _ := time.ParseDuration(v)
	return out
}
