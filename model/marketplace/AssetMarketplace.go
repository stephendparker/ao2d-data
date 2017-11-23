package marketplace

type AssetChidleyRoot314159 struct {
	AssetMarketplaces *AssetMarketplaces `xml:" marketplaces,omitempty" json:"marketplaces,omitempty"`
}

type AssetGuaranteedstock struct {
	AttrAmount string `xml:" amount,attr"  json:",omitempty"`
	AttrItem string `xml:" item,attr"  json:",omitempty"`
	AttrPrice string `xml:" price,attr"  json:",omitempty"`
}

type AssetMarketplace struct {
	AttrId string `xml:" id,attr"  json:",omitempty"`
	AssetGuaranteedstock *AssetGuaranteedstock `xml:" guaranteedstock,omitempty" json:"guaranteedstock,omitempty"`
}

type AssetMarketplaces struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetMarketplace []*AssetMarketplace `xml:" marketplace,omitempty" json:"marketplace,omitempty"`
}


