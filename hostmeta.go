package hostmeta

const (
	HostMetaPath     = "/.well-known/host-meta"
	HostMetaJSONPath = "/.well-known/host-meta.json"
)

type (
	// Link structure is the XRD link.
	// Please read the XRD documentation because it's a too f--kin' hassle to explain.
	Link struct {
		Text     string `xml:",chardata" json:"-"`
		Rel      string `xml:"rel,attr,omitempty" json:"rel,omitempty"`
		Type     string `xml:"type,attr,omitempty" json:"type,omitempty"`
		Href     string `xml:"href,attr,omitempty" json:"href,omitempty"`
		Template string `xml:"template,attr,omitempty" json:"template,omitempty"`
	}
)

func NewLink(Rel, Type, Href, Template string) *Link {
	return &Link{
		Rel:      Rel,
		Type:     Type,
		Href:     Href,
		Template: Template,
	}
}
