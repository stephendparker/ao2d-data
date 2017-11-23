package factions

type AssetChidleyRoot314159 struct {
	AssetFactions *AssetFactions `xml:" factions,omitempty" json:"factions,omitempty"`
}

type AssetFaction struct {
	AttrDescription string `xml:" description,attr"  json:",omitempty"`
	AttrIcon string `xml:" icon,attr"  json:",omitempty"`
	AttrTitle string `xml:" title,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
}

type AssetFactions struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetFaction []*AssetFaction `xml:" faction,omitempty" json:"faction,omitempty"`
}


