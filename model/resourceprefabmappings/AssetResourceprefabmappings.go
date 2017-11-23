package resourceprefabmappings

type AssetChidleyRoot314159 struct {
	AssetMappings *AssetMappings `xml:" mappings,omitempty" json:"mappings,omitempty"`
}

type AssetFiber struct {
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AssetFillerprefabs *AssetFillerprefabs `xml:" fillerprefabs,omitempty" json:"fillerprefabs,omitempty"`
	AssetPrefabs *AssetPrefabs `xml:" prefabs,omitempty" json:"prefabs,omitempty"`
}

type AssetFillerprefabs struct {
	AssetPrefab []*AssetPrefab `xml:" prefab,omitempty" json:"prefab,omitempty"`
}

type AssetHide struct {
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AssetFillerprefabs *AssetFillerprefabs `xml:" fillerprefabs,omitempty" json:"fillerprefabs,omitempty"`
	AssetPrefabs *AssetPrefabs `xml:" prefabs,omitempty" json:"prefabs,omitempty"`
}

type AssetMapping struct {
	AttrBiome string `xml:" biome,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AssetFiber []*AssetFiber `xml:" fiber,omitempty" json:"fiber,omitempty"`
	AssetHide []*AssetHide `xml:" hide,omitempty" json:"hide,omitempty"`
	AssetMobcamp *AssetMobcamp `xml:" mobcamp,omitempty" json:"mobcamp,omitempty"`
	AssetOre []*AssetOre `xml:" ore,omitempty" json:"ore,omitempty"`
	AssetRock []*AssetRock `xml:" rock,omitempty" json:"rock,omitempty"`
	AssetWood []*AssetWood `xml:" wood,omitempty" json:"wood,omitempty"`
}

type AssetMappings struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetMapping []*AssetMapping `xml:" mapping,omitempty" json:"mapping,omitempty"`
}

type AssetMobcamp struct {
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AssetFillerprefabs *AssetFillerprefabs `xml:" fillerprefabs,omitempty" json:"fillerprefabs,omitempty"`
	AssetPrefabs *AssetPrefabs `xml:" prefabs,omitempty" json:"prefabs,omitempty"`
}

type AssetOre struct {
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AssetFillerprefabs *AssetFillerprefabs `xml:" fillerprefabs,omitempty" json:"fillerprefabs,omitempty"`
	AssetPrefabs *AssetPrefabs `xml:" prefabs,omitempty" json:"prefabs,omitempty"`
}

type AssetPrefab struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrWeight string `xml:" weight,attr"  json:",omitempty"`
}

type AssetPrefabs struct {
	AssetPrefab []*AssetPrefab `xml:" prefab,omitempty" json:"prefab,omitempty"`
}

type AssetRock struct {
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AssetFillerprefabs *AssetFillerprefabs `xml:" fillerprefabs,omitempty" json:"fillerprefabs,omitempty"`
	AssetPrefabs *AssetPrefabs `xml:" prefabs,omitempty" json:"prefabs,omitempty"`
}

type AssetWood struct {
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AssetFillerprefabs *AssetFillerprefabs `xml:" fillerprefabs,omitempty" json:"fillerprefabs,omitempty"`
	AssetPrefabs *AssetPrefabs `xml:" prefabs,omitempty" json:"prefabs,omitempty"`
}


