package namegenerators

type AssetChidleyRoot314159 struct {
	AssetNamegenerators *AssetNamegenerators `xml:" namegenerators,omitempty" json:"namegenerators,omitempty"`
}

type AssetFirstname struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetNamegenerator struct {
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AssetFirstname []*AssetFirstname `xml:" firstname,omitempty" json:"firstname,omitempty"`
	AssetSurname []*AssetSurname `xml:" surname,omitempty" json:"surname,omitempty"`
}

type AssetNamegenerators struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetNamegenerator []*AssetNamegenerator `xml:" namegenerator,omitempty" json:"namegenerator,omitempty"`
}

type AssetSurname struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}


