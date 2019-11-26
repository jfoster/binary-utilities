bintools
===
A collection of cli utilities for manipulating binary files.

[Releases (precompiled binaries)](https://github.com/jfoster/bintools/releases)

## bytepad

Pads a file to a desired size with a specified padding byte

### Usage

```
./bytepad -size=<new size of file> -data=<byte data to pad file with> <path to file>
```

### Example

To pad file EVGA.GTX680.2048.120713.rom with FF's to 262144 bytes (256KiB):
```
./bytepad -size=262144 -data=0xFF ~/Downloads/EVGA.GTX680.2048.120713.rom
```

You can also omit -data in which 00's will be used for padding:
```
./bytepad -size=262144 ~/Downloads/EVGA.GTX680.2048.120713.rom
```

## cap2rom

Strips the header from an ASUS .CAP file to conver it to a .rom file for hardware flashing

### Usage

```
./cap2rom foo.CAP
```

which writes the .rom file alongside the specified .CAP file.
