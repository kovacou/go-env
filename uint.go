// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package env

// LookupUint parse a env variable as uint and if the key is defined.
func LookupUint(key string) (uint, bool) {
	v, find := Lookup(key)
	return toUint(v), find
}

// GetUint parse a env variable as uint.
func GetUint(key string) uint {
	return toUint(Parse(key))
}

// PrefixUint use a prefix to parse a env variable as uint.
func PrefixUint(prefix, key string) uint {
	return GetUint(Prefix(prefix, key))
}

// FormatUint use a format to parse a env variable as uint.
func FormatUint(format string, v ...any) uint {
	return GetUint(Format(format, v...))
}
