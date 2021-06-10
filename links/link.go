// Package links provides a data type for links.
package links

import (
	"fmt"
	"path"
	"strings"

	"github.com/bouncepaw/mycomarkup/util"
)

// LinkType tells what type the given link is.
type LinkType int

const (
	// LinkLocalRoot is a link like "/list", "/user-list", etc.
	LinkLocalRoot LinkType = iota
	// LinkLocalHypha is a link like "test", "../test", etc.
	LinkLocalHypha
	// LinkExternal is an external link with specified protocol.
	LinkExternal
	// LinkInterwiki is currently left unused. In the future it will be used for interwiki links.
	LinkInterwiki
)

// Link is an abstraction for universal representation of links, be they links in mycomarkup links or whatever.
type Link struct {
	// Known stuff
	srcAddress string
	srcDisplay string
	srcHypha   string
	// Parsed stuff
	anchor   string
	address  string
	display  string
	kind     LinkType
	protocol string
	// Settable stuff

	DestinationKnown bool
}

// From makes a link from the given source address and display text on the given hypha.
func From(srcAddress, srcDisplay, srcHypha string) *Link {
	link := Link{
		srcAddress:       strings.TrimSpace(srcAddress),
		srcDisplay:       strings.TrimSpace(srcDisplay),
		srcHypha:         strings.TrimSpace(srcHypha),
		DestinationKnown: false,
	}
	link.address = link.srcAddress

	// If there is a hash sign in the address, move everything starting from the sign to the end of the address to the anchor field and truncate the address.
	if pos := strings.IndexRune(link.srcAddress, '#'); pos != -1 {
		link.anchor = link.srcAddress[pos:]
		link.address = link.address[:pos]
	}

	// NOTE: This part will need some extending with introduction of interwiki.

	switch {
	// If is an external link
	case strings.ContainsRune(link.address, ':'):
		pos := strings.IndexRune(link.address, ':')
		link.kind = LinkExternal
		link.protocol = link.address[:pos+1]
		link.address = link.address[pos+1:]
		if strings.HasPrefix(link.address, "//") && len(link.address) > 2 {
			link.protocol += "//"
			link.address = link.address[2:]
		}
	case strings.HasPrefix(link.address, "/"):
		link.kind = LinkLocalRoot
	case strings.HasPrefix(link.address, "./"):
		link.kind = LinkLocalHypha
		link.address = util.CanonicalName(path.Join(link.srcHypha, link.address[2:]))
	case link.address == "..":
		link.address = util.CanonicalName(path.Dir(link.srcHypha))
	case strings.HasPrefix(link.address, "../"):
		link.kind = LinkLocalHypha
		link.address = util.CanonicalName(path.Join(path.Dir(link.srcHypha), link.address[3:]))
	case strings.HasPrefix(link.address, "#"):
		link.kind = LinkLocalHypha
		link.anchor = link.address
		link.address = util.CanonicalName(link.srcHypha)
	default:
		link.kind = LinkLocalHypha
		link.address = util.CanonicalName(link.address)
	}

	// If no display text is given, copy the address there.
	if link.srcDisplay == "" {
		// Drop the protocol if there is any.
		link.display = strings.TrimPrefix(link.srcAddress, link.protocol)
	} else {
		link.display = link.srcDisplay
	}

	return &link
}

// Exists is true if the link should be blue, not red. Red links are links to hyphae that do not exist, all other links are blue.
func (link *Link) Exists() bool {
	return (link.OfKind(LinkExternal)) || (link.OfKind(LinkLocalRoot)) || (link.OfKind(LinkLocalHypha) && !link.DestinationKnown)
}

// ItExists notes that the destination makes sense, exists.
func (link *Link) ItExists() *Link {
	link.DestinationKnown = true
	return link
}

// Classes returns CSS class string for given link. It is not wrapped in any quotes, wrap yourself.
func (link *Link) Classes() (classes string) {
	classes = "wikilink"
	switch link.kind {
	case LinkLocalRoot, LinkLocalHypha:
		classes += " wikilink_internal"
		if link.DestinationKnown {
			classes += " wikilink_new"
		}
	case LinkInterwiki:
		classes += " wikilink_interwiki"
	case LinkExternal:
		classes += fmt.Sprintf(" wikilink_external wikilink_%s", strings.TrimSuffix(strings.TrimSuffix(link.protocol, "://"), ":"))
	}
	return classes
}

// Href returns content for the href attrubute for hyperlink. You should always use it.
func (link *Link) Href() string {
	switch link.kind {
	case LinkExternal, LinkLocalRoot:
		return link.protocol + link.address + link.anchor
	default:
		return "/hypha/" + link.address + link.anchor
	}
}

// ImgSrc returns content for src attribute of img tag. Used with `img{}`.
func (link *Link) ImgSrc() string {
	switch link.kind {
	case LinkExternal, LinkLocalRoot:
		return link.protocol + link.address + link.anchor
	default:
		return "/binary/" + link.address
	}
}

// String returns a debugging string representation of the given link.
func (link *Link) String() string {
	return fmt.Sprintf(`Link("%s", "%s", "%s")`, link.Href(), link.Display(), link.srcHypha)
}

// Display returns the display text of the given link.
func (link *Link) Display() string {
	return link.display
}

// Address returns the address of the given link. Why would you need that?
func (link *Link) Address() string {
	return link.address
}

// OfKind returns if the given link is of the given kind.
func (link *Link) OfKind(kind LinkType) bool {
	return link.kind == kind
}
