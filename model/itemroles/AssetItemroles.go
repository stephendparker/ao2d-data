package itemroles

type AssetChidleyRoot314159 struct {
	AssetItemroles *AssetItemroles `xml:" itemroles,omitempty" json:"itemroles,omitempty"`
}

type AssetCategory struct {
	AttrRating string `xml:" rating,attr"  json:",omitempty"`
	AttrValue string `xml:" value,attr"  json:",omitempty"`
}

type AssetItemroles struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetRole []*AssetRole `xml:" role,omitempty" json:"role,omitempty"`
}

type AssetRole struct {
	AttrDisplayname string `xml:" displayname,attr"  json:",omitempty"`
	AttrUisprite string `xml:" uisprite,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AssetCategory []*AssetCategory `xml:" category,omitempty" json:"category,omitempty"`
}


