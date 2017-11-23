package expeditioncategories

type AssetChidleyRoot314159 struct {
	AssetExpeditioncategories *AssetExpeditioncategories `xml:" expeditioncategories,omitempty" json:"expeditioncategories,omitempty"`
}

type AssetBonusrewards struct {
	AssetFamepoints *AssetFamepoints `xml:" famepoints,omitempty" json:"famepoints,omitempty"`
	AssetItemreward []*AssetItemreward `xml:" itemreward,omitempty" json:"itemreward,omitempty"`
	AssetSilver *AssetSilver `xml:" silver,omitempty" json:"silver,omitempty"`
}

type AssetExpeditioncategories struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetExpeditioncategory []*AssetExpeditioncategory `xml:" expeditioncategory,omitempty" json:"expeditioncategory,omitempty"`
}

type AssetExpeditioncategory struct {
	AttrGrouptype string `xml:" grouptype,attr"  json:",omitempty"`
	AttrMinitempoweraverage string `xml:" minitempoweraverage,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUnlockedtoregister string `xml:" unlockedtoregister,attr"  json:",omitempty"`
	AssetBonusrewards *AssetBonusrewards `xml:" bonusrewards,omitempty" json:"bonusrewards,omitempty"`
	AssetNeededitemroles *AssetNeededitemroles `xml:" neededitemroles,omitempty" json:"neededitemroles,omitempty"`
	AssetRewards *AssetRewards `xml:" rewards,omitempty" json:"rewards,omitempty"`
}

type AssetFamepoints struct {
	AttrPoints string `xml:" points,attr"  json:",omitempty"`
}

type AssetItemreward struct {
	AttrCount string `xml:" count,attr"  json:",omitempty"`
	AttrEnchantmentlevel string `xml:" enchantmentlevel,attr"  json:",omitempty"`
	AttrItemid string `xml:" itemid,attr"  json:",omitempty"`
}

type AssetNeededitemroles struct {
	AssetRole []*AssetRole `xml:" role,omitempty" json:"role,omitempty"`
}

type AssetRewards struct {
	AssetFamepoints *AssetFamepoints `xml:" famepoints,omitempty" json:"famepoints,omitempty"`
	AssetItemreward []*AssetItemreward `xml:" itemreward,omitempty" json:"itemreward,omitempty"`
	AssetSilver *AssetSilver `xml:" silver,omitempty" json:"silver,omitempty"`
}

type AssetRole struct {
	AttrItemroleid string `xml:" itemroleid,attr"  json:",omitempty"`
	AttrMax string `xml:" max,attr"  json:",omitempty"`
	AttrMin string `xml:" min,attr"  json:",omitempty"`
}

type AssetSilver struct {
	AttrAbsolute string `xml:" absolute,attr"  json:",omitempty"`
}


