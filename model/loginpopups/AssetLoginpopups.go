package loginpopups

type AssetChidleyRoot314159 struct {
	AssetLoginpopups *AssetLoginpopups `xml:" loginpopups,omitempty" json:"loginpopups,omitempty"`
}

type AssetLoginpopup struct {
	AttrSecondstowait string `xml:" secondstowait,attr"  json:",omitempty"`
	AttrTag string `xml:" tag,attr"  json:",omitempty"`
	AttrTitle string `xml:" title,attr"  json:",omitempty"`
	AttrTobeaccepted string `xml:" tobeaccepted,attr"  json:",omitempty"`
	AttrVersion string `xml:" version,attr"  json:",omitempty"`
	AssetText []*AssetText `xml:" text,omitempty" json:"text,omitempty"`
}

type AssetLoginpopups struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetLoginpopup []*AssetLoginpopup `xml:" loginpopup,omitempty" json:"loginpopup,omitempty"`
}

type AssetText struct {
	AttrTag string `xml:" tag,attr"  json:",omitempty"`
}


