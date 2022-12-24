# xxd - a Go package for generating hex dumps

## Overview

Under Linux there is a handy stand alone command `xxd` that can be
used to generate HEX dumps of binary data files. This package
reproduces one of its output formats: the output of `xxd -g1`.

With `data := []bytes{0x00, 0xf0, 0x07, ... 0x5b, 0x00, 0x20}`, of 64
bytes total length, for example, `xxd.Print(0, data)` will generate
the following output:
```
00000000: 00 f0 07 20 c5 63 00 20 8f 31 00 20 79 31 00 20  ... .c. .1. y1. 
00000010: 95 31 00 20 9b 31 00 20 a1 31 00 20 00 00 00 00  .1. .1. .1. ....
00000020: 00 00 00 00 00 00 00 00 00 00 00 00 31 3a 00 20  ............1:. 
00000030: a7 31 00 20 00 00 00 00 f1 3a 00 20 11 5b 00 20  .1. .....:. .[. 
```

Each line is aligned to a multiple of 16 bytes, but `data` is output
starting at the offset indicated by the first argument. As such,
`xxd.Print(5, data)` will generate:
```
00000000:                00 f0 07 20 c5 63 00 20 8f 31 00       ... .c. .1.
00000010: 20 79 31 00 20 95 31 00 20 9b 31 00 20 a1 31 00   y1. .1. .1. .1.
00000020: 20 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00   ...............
00000030: 00 31 3a 00 20 a7 31 00 20 00 00 00 00 f1 3a 00  .1:. .1. .....:.
00000040: 20 11 5b 00 20                                    .[. 
```

Automated documentation for this Go package is available from [![Go
Reference](https://pkg.go.dev/badge/zappem.net/pub/debug/xxd.svg)](https://pkg.go.dev/zappem.net/pub/debug/xxd).

## License info

The `xxd` package is distributed with the same BSD 3-clause license
as that used by [golang](https://golang.org/LICENSE) itself.

## Reporting bugs and feature requests

The `xxd` package was developed purely out of self-interest to help
debug other programs and packages. Should you find a bug or want to
suggest a feature addition, please use the [bug
tracker](https://github.com/tinkerator/xxd/issues).
