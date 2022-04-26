// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package env

// LookupInt parse a env variable as int and if the key is defined.
func LookupInt(key string) (int, bool) {
	v, find := Lookup(key)
	return toInt(v), find
}

// GetInt parse a env variable as int.
func GetInt(key string) int {
	return toInt(Parse(key))
}

// PrefixInt use a prefix to parse a env variable as int.
func PrefixInt(prefix, key string) int {
	return GetInt(Prefix(prefix, key))
}

// FormatInt use a format to parse a env variable as int.
func FormatInt(format string, v ...any) int {
	return GetInt(Format(format, v...))
}
