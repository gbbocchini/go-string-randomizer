# go-string-randomizer

Go-string-randomizer is a simple but powerfull random strings generator for GO. 
You can use it to generate random ids, passwords, etc.

The package also includes a usefull collection of chars that can be used as argument LettersUniverse.

## Installation

```bash
go get github.com/gbbocchini/go-string-randomizer
```

## Usage

```go
import "github.com/gbbocchini/go-string-randomizer"

randomizer := StringRandomizer{
    LettersUniverse: LOWER_UPPER_LETTERS_NUMS,
    GeneratedMaxLen: 13,
    NoCollisions:    true,
    RandSeed:        123,
}

one := randomizer.GenerateOne()

bulk := randomizer.GenerateBulk(1000000)
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)