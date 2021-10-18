# Longhand

### Installation
`$ go get github.com/christianraza/longhand`

### Usage
```
package main

import (
    "fmt"
    "log"

    "github.com/christianraza/longhand"
)

func main() {
    i, err := longhand.ParseLonghand("Forty-two")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Forty-two as an integer: %d\n", i)
    i64, err := longhand.ParseLonghand64("Forty-two")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Forty-two as a 64-bit integer: %d\n", i64)
}
```

### Supported Styles
##### Standard
`Five thousand two hundred and thirty-three` = `5233`
##### Years
`Nineteen thirty` = `1930`
##### Repetition
`Thirty thirty` = `3030`
##### Notes on style
What all this means is that something like `nineteen thirty thirty hundred` is valid input and will be parsed as `19303000`. `nineteen thirty thirty hundred` is not a common way of writing `19303000` but it is a possibility. So this, as well as many other variations, are supported styles.

### "Valid" Longhand
##### Examples of valid longhand
`Three` as opposed to `3`
`Thirty-three` as opposed to `33`
`Three hundred and thirty-three` as opposed to `333`
##### Examples of invalid longhand
`Thre` as opposed to `Three` 
`Thirty_three` as opposed to `Thirty-three`
`3 hundred and thirty-three` as opposed to `Three hundred and thirty-three`
##### Consequences of invalid longhand
There are dire dire consequences for invalid longhand!
:fearful: ...
Just kidding :smile: invalid longhand just returns zero and an error.
##### "Valid"?
You likely noticed the quotes around valid in the title of this section, that's because the way people interpret numerals sometimes depends on location and preference. I'm hard pressed to say that there is a valid way of arranging numerals, and some of this can be seen in the section on styles. I had to make compromises here and there, as is usual when writing tools or software, but generally these functions should do something reasonable so have fun! :smiley: