# fork of jdeng/goheif without C code - only for parsing HEIF container and extracting EXIF

## Install
- Tested
  - Mac OS X (High Sierra) 
  - Linux (Ubuntu 16.04 / GCC 5.4)
  - Windows 7 64bit with TDM-GCC 32 (GCC 5.1) and golang 1.12 windows/386

- Code Sample
```
func main() {
	flag.Parse()
	...
  
	fin := flag.Arg(0)
	fi, err := os.Open(fin)
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	exif, err := goheif.ExtractExif(fi)
	if err != nil {
		log.Printf("Warning: no EXIF from %s: %v\n", fin, err)
	}

	log.Printf("got exif %q", exif)
}
```

## What is done
- Removed C code and heic2jpg utility from (https://github.com/jdeng/goheif) golang heif parse

- Changes make to @bradfitz's (https://github.com/bradfitz) golang heif parser
  - Some minor bugfixes
  - A few new box parsers, noteably 'iref' and 'hvcC'

## License

- heif subdir is in Apache license

- goheif.go is in MIT license

## Credits
- heif parser by @bradfitz (https://github.com/go4org/go4/tree/master/media/heif)
- implementation learnt from libheif (https://github.com/strukturag/libheif)

## TODO
- Upstream the changes to heif?


