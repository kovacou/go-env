# github.com/kovacou/go-env

Personal project.

You can automatically load your environment variable into specific struct by using the tag `env:""` and calling the function `env.Unmarshal()`.

**Note:** The package doesn't check the validity of the data.

## ➡ install

```
go get github.com/kovacou/go-env
```
## ➡ usage

```ini
APP_TITLE=Titre de mon application
APP_VERBOSE=true
APP_KEYS=facebook:key1,google:key2,twitter:key3
APP_SLICE_KEYS=key1,key2,key3
```  

```go
package main

import (
    "github.com/kovacou/go-env"
)

type Config   struct {
    Title     string            `env:"APP_TITLE"`
    Verbose   bool              `env:"APP_VERBOSE"`
    Keys      map[string]string `env:"APP_KEYS"`
    SliceKeys []string          `env:"APP_SLICE_KEYS"`
}

func main() {
    cfg := Config{}
    env.Unmarshal(&cfg)
    // Config{
    //     Title:     "My Application",
    //     Verbose:   true,
    //     Keys:      {"facebook":"key1", "google":"key2", "twitter":"key3"},
    //     SliceKeys: {"key1", "key2", "key3"}
    // }
    println(cfg.Title) // print "My Application"
}
```
## ➡ **Supported types**

- Atomic types :
    - `string`, `*string`, 
    - `int`, `*int`, 
    - `uint`, `*uint`,
    - `int64`, `*int64`, 
    - `uint64`, `*uint64`, 
    - `float64`, `*float64`, 
    - `bool`, `*bool`
    - `interface{}`
- Special types :
    - `time.Duration`, `*time.Duration`
- Maps :
    - `map[string]interface{}`
    - `map[string]bool`
    - `map[string]int`
    - `map[string]uint`
    - `map[string]int64`
    - `map[string]uint64`
    - `map[string]float64`
    - `map[string]string`
- Slices :
    - `[]interface{}`
    - `[]bool`
    - `[]int`
    - `[]uint`
    - `[]int64`
    - `[]uint64`
    - `[]float64`
    - `[]string`
