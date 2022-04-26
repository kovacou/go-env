// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package env

// LookupString parse a env variable as string and if the key is defined.
func LookupString(key string) (string, bool) {
	return Lookup(key)
}

// GetString parse a env variable as string.
func GetString(key string) string {
	return Parse(key)
}

// PrefixString use a prefix to parse a env variable as string.
func PrefixString(prefix, key string) string {
	return GetString(Prefix(prefix, key))
}

// FormatString use a format to parse a env variable as string.
func FormatString(format string, v ...any) string {
	return GetString(Format(format, v...))
}
