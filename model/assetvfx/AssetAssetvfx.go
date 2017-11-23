package assetvfx

type AssetAssetVfx struct {
	AttrSocketname string `xml:" socketname,attr"  json:",omitempty"`
	AssetVfx []*AssetVfx `xml:" Vfx,omitempty" json:"Vfx,omitempty"`
}

type AssetAssetVfxInfo struct {
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AssetAssetVfx []*AssetAssetVfx `xml:" AssetVfx,omitempty" json:"AssetVfx,omitempty"`
	AssetRareResourceFx *AssetRareResourceFx `xml:" RareResourceFx,omitempty" json:"RareResourceFx,omitempty"`
}

type AssetAssetVfxInfos struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetAssetVfxInfo []*AssetAssetVfxInfo `xml:" AssetVfxInfo,omitempty" json:"AssetVfxInfo,omitempty"`
}

type AssetChidleyRoot314159 struct {
	AssetAssetVfxInfos *AssetAssetVfxInfos `xml:" AssetVfxInfos,omitempty" json:"AssetVfxInfos,omitempty"`
}

type AssetRareResourceFx struct {
	AssetRareState []*AssetRareState `xml:" RareState,omitempty" json:"RareState,omitempty"`
}

type AssetRareState struct {
	AttrStateid string `xml:" stateid,attr"  json:",omitempty"`
	AssetVfx []*AssetVfx `xml:" Vfx,omitempty" json:"Vfx,omitempty"`
}

type AssetVfx struct {
	AttrConstraintpreset string `xml:" constraintpreset,attr"  json:",omitempty"`
	AttrIgnoreSocketScale string `xml:" ignoreSocketScale,attr"  json:",omitempty"`
	AttrPrefabName string `xml:" prefabName,attr"  json:",omitempty"`
	AttrScale string `xml:" scale,attr"  json:",omitempty"`
	AttrSocket string `xml:" socket,attr"  json:",omitempty"`
}


