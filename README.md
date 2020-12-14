### We can:
1. Translate phrases with [Pig Latin](https://en.wikipedia.org/wiki/Pig_Latin)
2. Encode and decode input string replacing vowels with rules:
    * a -> 1
    * e -> 2
    * i -> 3
    * o -> 4
    * u -> 5
   
    Examples:
   
   ```encode("hello") == "h2ll4"```
   
   ```decode("h3 th2r2") == "hi there"```
* should work regardless of input string case

## Requirements
* make
* golang v1.15.*

## Usage:
1. build `make build`
2. run 
    * `$ ./build/app/vowels (your text)`
    * `$ ./build/app/vowels pig-latin (your text)`

## Test:
`$ make test`

## Examples:

```
$ ./build/app/vowels "hello it's me"
h2ll4 3t's m2

$ ./build/app/vowels -reverse "h2ll4 3t's m2"
hello it's me

$ ./build/app/pig-latin The modern version of Pig Latin appears in a 1919
```
