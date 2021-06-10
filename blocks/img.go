package blocks

import (
	"fmt"
	"github.com/bouncepaw/mycomarkup/globals"
	"regexp"
	"strings"

	"github.com/bouncepaw/mycomarkup/links"
)

var imgRe = regexp.MustCompile(`^img {`)

// MatchesImg is true if the line starts with img {.
func MatchesImg(line string) bool {
	return imgRe.MatchString(line)
}

type imgState int

const (
	inRoot imgState = iota
	inName
	inDimensionsW
	inDimensionsH
	inDescription
)

// Img is an image gallery, consisting of zero or more images.
type Img struct {
	Entries   []ImgEntry
	currEntry ImgEntry
	hyphaName string
	state     imgState
}

func (img Img) isBlock() {}

// ID returns the gallery's id which is img- and a number.
func (img Img) ID(counter *IDCounter) string {
	counter.imgs++
	return fmt.Sprintf("img-%d", counter.imgs)
}

// HasOneImage returns true if img has exactly one image and that images has no description.
func (img *Img) HasOneImage() bool {
	return len(img.Entries) == 1 && img.Entries[0].desc.Len() == 0
}

func (img *Img) pushEntry() {
	if strings.TrimSpace(img.currEntry.path.String()) != "" {
		img.currEntry.Srclink = links.From(img.currEntry.path.String(), "", img.hyphaName)
		// img.currEntry.Srclink.DoubtExistence()
		img.Entries = append(img.Entries, img.currEntry)
		img.currEntry = ImgEntry{hyphaName: img.hyphaName}
		img.currEntry.path.Reset()
	}
}

// ProcessLine parses the line and tells if the gallery is finished.
func (img *Img) ProcessLine(line string) (imgFinished bool) {
	for _, r := range line {
		if shouldReturnTrue := img.ProcessRune(r); shouldReturnTrue {
			return true
		}
	}
	// We do that because \n are not part of line we receive as the argument:
	img.ProcessRune('\n')

	return false
}

// ProcessRune parses the rune.
func (img *Img) ProcessRune(r rune) (done bool) {
	// TODO: move to the parser module.
	if r == '\r' {
		return false
	}
	switch img.state {
	case inRoot:
		return img.processInRoot(r)
	case inName:
		return img.processInName(r)
	case inDimensionsW:
		return img.processInDimensionsW(r)
	case inDimensionsH:
		return img.processInDimensionsH(r)
	case inDescription:
		return img.processInDescription(r)
	}
	fmt.Println("ProcessRune: unreachable state", r)
	return true
}

func (img *Img) processInDescription(r rune) (shouldReturnTrue bool) {
	switch r {
	case '}':
		img.state = inName
	default:
		img.currEntry.desc.WriteRune(r)
	}
	return false
}

func (img *Img) processInRoot(r rune) (shouldReturnTrue bool) {
	switch r {
	case '}':
		img.pushEntry()
		return true
	case '\n':
		img.pushEntry()
	case ' ', '\t':
	default:
		img.state = inName
		img.currEntry = ImgEntry{}
		img.currEntry.path.Reset()
		img.currEntry.path.WriteRune(r)
	}
	return false
}

func (img *Img) processInName(r rune) (shouldReturnTrue bool) {
	switch r {
	case '}':
		img.pushEntry()
		return true
	case '|':
		img.state = inDimensionsW
	case '{':
		img.state = inDescription
	case '\n':
		img.pushEntry()
		img.state = inRoot
	default:
		img.currEntry.path.WriteRune(r)
	}
	return false
}

func (img *Img) processInDimensionsW(r rune) (shouldReturnTrue bool) {
	switch r {
	case '}':
		img.pushEntry()
		return true
	case '*':
		img.state = inDimensionsH
	case ' ', '\t', '\n':
	case '{':
		img.state = inDescription
	default:
		img.currEntry.sizeW.WriteRune(r)
	}
	return false
}

func (img *Img) processInDimensionsH(r rune) (imgFinished bool) {
	switch r {
	case '}':
		img.pushEntry()
		return true
	case ' ', '\t', '\n':
	case '{':
		img.state = inDescription
	default:
		img.currEntry.sizeH.WriteRune(r)
	}
	return false
}

// MakeImg parses the image gallery on the line and returns it. It also tells if the gallery is finished or not.
func MakeImg(line, hyphaName string) (img Img, imgFinished bool) {
	img = Img{
		hyphaName: hyphaName,
		Entries:   make([]ImgEntry, 0),
	}
	line = line[strings.IndexRune(line, '{')+1:]
	return img, img.ProcessLine(line)
}

// MarkExistenceOfSrcLinks effectively checks if the links in the gallery are blue or red.
func (img *Img) MarkExistenceOfSrcLinks() {
	globals.HyphaIterate(func(hn string) {
		for _, entry := range img.Entries {
			if hn == entry.Srclink.Address() {
				entry.Srclink.DestinationKnown = true
			}
		}
	})
}
