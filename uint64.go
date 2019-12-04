// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package env

// LookupUint64 parse a env variable as uint64 and if the key is defined.
func LookupUint64(key string) (uint64, bool) {
	v, find := Lookup(key)
	return toUint64(v), find
}

// GetUint64 parse a env variable as uint64.
func GetUint64(key string) uint64 {
	return toUint64(Parse(key))
}

// PrefixUint64 use a prefix to parse a env variable as uint64.
func PrefixUint64(prefix, key string) uint64 {
	return GetUint64(Prefix(prefix, key))
}

// FormatUint64 use a format to parse a env variable as uint64.
func FormatUint64(format string, v ...interface{}) uint64 {
	return GetUint64(Format(format, v...))
}
