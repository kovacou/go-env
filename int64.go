// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package env

// LookupInt64 parse a env variable as int64 and if the key is defined.
func LookupInt64(key string) (int64, bool) {
	v, find := Lookup(key)
	return toInt64(v), find
}

// GetInt64 parse a env variable as int64.
func GetInt64(key string) int64 {
	return toInt64(Parse(key))
}

// PrefixInt64 use a prefix to parse a env variable as int64.
func PrefixInt64(prefix, key string) int64 {
	return GetInt64(Prefix(prefix, key))
}

// FormatInt64 use a format to parse a env variable as int64.
func FormatInt64(format string, v ...any) int64 {
	return GetInt64(Format(format, v...))
}
