package spellanimmappings

type AssetChidleyRoot314159 struct {
	AssetSpellAnimationMappingInfos *AssetSpellAnimationMappingInfos `xml:" SpellAnimationMappingInfos,omitempty" json:"SpellAnimationMappingInfos,omitempty"`
}

type AssetDefault struct {
	AttrCastanimation string `xml:" castanimation,attr"  json:",omitempty"`
}

type AssetMappingInfo struct {
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AssetDefault *AssetDefault `xml:" Default,omitempty" json:"Default,omitempty"`
	AssetWeapon *AssetWeapon `xml:" Weapon,omitempty" json:"Weapon,omitempty"`
}

type AssetSpellAnimationMappingInfos struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetMappingInfo *AssetMappingInfo `xml:" MappingInfo,omitempty" json:"MappingInfo,omitempty"`
}

type AssetWeapon struct {
	AttrCastanimation string `xml:" castanimation,attr"  json:",omitempty"`
	AttrWeaponanimationtype string `xml:" weaponanimationtype,attr"  json:",omitempty"`
}


