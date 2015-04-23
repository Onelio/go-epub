package epub

import (
    "encoding/xml"
    "log"
)

const (
// TODO
	temp = `<?xml version="1.0" encoding="UTF-8"?>
<package version="3.0" unique-identifier="pub-id" xmlns="http://www.idpf.org/2007/opf">
  <metadata xmlns:dc="http://purl.org/dc/elements/1.1/">
    <dc:identifier id="pub-id">urn:uuid:fe93046f-af57-475a-a0cb-a0d4bc99ba6d</dc:identifier>
    <dc:title>Your title here</dc:title>
    <dc:language>en</dc:language>
    <meta property="dcterms:modified">2011-01-01T12:00:00Z</meta>
  </metadata>
  <manifest>
    <item id="nav" href="nav.xhtml" media-type="application/xhtml+xml" properties="nav" />
    <item id="ncx" href="toc.ncx" media-type="application/x-dtbncx+xml" />
    <item id="section0001.xhtml" href="xhtml/section0001.xhtml" media-type="application/xhtml+xml" />
  </manifest>
  <spine toc="ncx">
    <itemref idref="section0001.xhtml" />
  </spine>
</package>
`
	packageFileTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<package version="3.0" unique-identifier="pub-id" xmlns="http://www.idpf.org/2007/opf">
  <metadata xmlns:dc="http://purl.org/dc/elements/1.1/">
    <dc:identifier id="pub-id"></dc:identifier>
    <dc:title></dc:title>
    <dc:language></dc:language>
    <meta property="dcterms:modified"></meta>
  </metadata>
  <manifest>
  </manifest>
  <spine toc="ncx">
  </spine>
</package>
`

	contentUniqueIdentifier = "pub-id"
	contentXmlnsDc = "http://purl.org/dc/elements/1.1/"
)

type Itemref struct {
    Idref string `xml:"idref,attr"`
}

type Spine struct {
    Itemref []Itemref `xml:"itemref"`
}

type Item struct {
    Href string `xml:"href,attr"`
    Id string `xml:"id,attr"`
    MediaType string `xml:"media-type,attr"`
    Properties string `xml:"properties,attr"`
}

type Manifest struct {
    Item []Item `xml:"item"`
}

type Meta struct {
	Property string `xml:"property,attr"`
	Data string `xml:",chardata"`
}

type Language struct {
	Data string `xml:",chardata"`
}

type Title struct {
	Data string `xml:",chardata"`
}

type Identifier struct {
    Id string `xml:"id,attr"`
    Data string `xml:",chardata"`
}

type Metadata struct {
    XmlnsDc string `xml:"xmlns:dc,attr"`
    Identifier Identifier `xml:"dc:identifier"`
    Title Title `xml:"dc:title"`
    Language Language `xml:"dc:language"`
    Meta Meta `xml:"meta"`
}

type Pkgdoc struct {
    XMLName xml.Name `xml:"http://www.idpf.org/2007/opf package"`
    UniqueIdentifier string `xml:"unique-identifier,attr"`
    Version string `xml:"version,attr"`
    Metadata Metadata `xml:"metadata"`
    Manifest Manifest `xml:"manifest"`
    Spine Spine `xml:"spine"`
}

func NewPkgdoc() *Pkgdoc {
	v := &Pkgdoc{}

	err := xml.Unmarshal([]byte(packageFileTemplate), &v)
	if err != nil {
		log.Fatalf("xml.Unmarshal error: %s", err)
	}
	
	v.Metadata.XmlnsDc = contentXmlnsDc
	v.Metadata.Identifier.Id = contentUniqueIdentifier
    
    return v
}

func (p *Pkgdoc) setLang(lang string) {
	p.Metadata.Language.Data = lang
}

func (p *Pkgdoc) setModified(timestamp string) {
	p.Metadata.Meta.Data = timestamp
}

func (p *Pkgdoc) setTitle(title string) {
	p.Metadata.Title.Data = title
}

func (p *Pkgdoc) setUUID(uuid string) {
	p.Metadata.Identifier.Data = uuid
}
