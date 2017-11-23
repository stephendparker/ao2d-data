package blackmarket

type AssetChidleyRoot314159 struct {
	AssetBlackmarket *AssetBlackmarket `xml:" blackmarket,omitempty" json:"blackmarket,omitempty"`
}

type AssetBlackmarket struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetEquipmentstorage *AssetEquipmentstorage `xml:" equipmentstorage,omitempty" json:"equipmentstorage,omitempty"`
}

type AssetEnchantmentlevelfactor struct {
	AttrPricefactor string `xml:" pricefactor,attr"  json:",omitempty"`
	AttrStoragefactor string `xml:" storagefactor,attr"  json:",omitempty"`
	AttrValue string `xml:" value,attr"  json:",omitempty"`
	AttrWastefactor string `xml:" wastefactor,attr"  json:",omitempty"`
}

type AssetEnchantmentlevelfactors struct {
	AssetEnchantmentlevelfactor []*AssetEnchantmentlevelfactor `xml:" enchantmentlevelfactor,omitempty" json:"enchantmentlevelfactor,omitempty"`
}

type AssetEquipmentstorage struct {
	AttrBasepricefactor string `xml:" basepricefactor,attr"  json:",omitempty"`
	AttrBasestoragesize string `xml:" basestoragesize,attr"  json:",omitempty"`
	AttrBasewastechance string `xml:" basewastechance,attr"  json:",omitempty"`
	AssetSlottypefactors *AssetSlottypefactors `xml:" slottypefactors,omitempty" json:"slottypefactors,omitempty"`
	AssetTier []*AssetTier `xml:" tier,omitempty" json:"tier,omitempty"`
}

type AssetQualitylevelfactor struct {
	AttrPricefactor string `xml:" pricefactor,attr"  json:",omitempty"`
	AttrStoragefactor string `xml:" storagefactor,attr"  json:",omitempty"`
	AttrValue string `xml:" value,attr"  json:",omitempty"`
	AttrWastefactor string `xml:" wastefactor,attr"  json:",omitempty"`
}

type AssetQualitylevelfactors struct {
	AssetQualitylevelfactor []*AssetQualitylevelfactor `xml:" qualitylevelfactor,omitempty" json:"qualitylevelfactor,omitempty"`
}

type AssetSlot struct {
	AttrPricefactor string `xml:" pricefactor,attr"  json:",omitempty"`
	AttrStoragefactor string `xml:" storagefactor,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
	AttrWastefactor string `xml:" wastefactor,attr"  json:",omitempty"`
}

type AssetSlottypefactors struct {
	AssetSlot []*AssetSlot `xml:" slot,omitempty" json:"slot,omitempty"`
}

type AssetTier struct {
	AttrPricefactor string `xml:" pricefactor,attr"  json:",omitempty"`
	AttrStoragefactor string `xml:" storagefactor,attr"  json:",omitempty"`
	AttrValue string `xml:" value,attr"  json:",omitempty"`
	AttrWastefactor string `xml:" wastefactor,attr"  json:",omitempty"`
	AssetEnchantmentlevelfactors *AssetEnchantmentlevelfactors `xml:" enchantmentlevelfactors,omitempty" json:"enchantmentlevelfactors,omitempty"`
	AssetQualitylevelfactors *AssetQualitylevelfactors `xml:" qualitylevelfactors,omitempty" json:"qualitylevelfactors,omitempty"`
}


