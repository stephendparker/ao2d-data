package resourcedistpresets

type AssetChidleyRoot314159 struct {
	AssetPresets *AssetPresets `xml:" presets,omitempty" json:"presets,omitempty"`
}

type AssetFiber struct {
	AttrAmount string `xml:" amount,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
}

type AssetHide struct {
	AttrAmount string `xml:" amount,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
}

type AssetHigh struct {
	AssetFiber []*AssetFiber `xml:" fiber,omitempty" json:"fiber,omitempty"`
	AssetHide []*AssetHide `xml:" hide,omitempty" json:"hide,omitempty"`
	AssetMobcamp *AssetMobcamp `xml:" mobcamp,omitempty" json:"mobcamp,omitempty"`
	AssetOre []*AssetOre `xml:" ore,omitempty" json:"ore,omitempty"`
	AssetRock []*AssetRock `xml:" rock,omitempty" json:"rock,omitempty"`
	AssetWood []*AssetWood `xml:" wood,omitempty" json:"wood,omitempty"`
}

type AssetLow struct {
	AssetFiber []*AssetFiber `xml:" fiber,omitempty" json:"fiber,omitempty"`
	AssetHide []*AssetHide `xml:" hide,omitempty" json:"hide,omitempty"`
	AssetOre []*AssetOre `xml:" ore,omitempty" json:"ore,omitempty"`
	AssetRock []*AssetRock `xml:" rock,omitempty" json:"rock,omitempty"`
	AssetWood []*AssetWood `xml:" wood,omitempty" json:"wood,omitempty"`
}

type AssetMedium struct {
	AssetFiber []*AssetFiber `xml:" fiber,omitempty" json:"fiber,omitempty"`
	AssetHide []*AssetHide `xml:" hide,omitempty" json:"hide,omitempty"`
	AssetMobcamp *AssetMobcamp `xml:" mobcamp,omitempty" json:"mobcamp,omitempty"`
	AssetOre []*AssetOre `xml:" ore,omitempty" json:"ore,omitempty"`
	AssetRock []*AssetRock `xml:" rock,omitempty" json:"rock,omitempty"`
	AssetWood []*AssetWood `xml:" wood,omitempty" json:"wood,omitempty"`
}

type AssetMobcamp struct {
	AttrAmount string `xml:" amount,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
}

type AssetOre struct {
	AttrAmount string `xml:" amount,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
}

type AssetPreset struct {
	AttrBiome string `xml:" biome,attr"  json:",omitempty"`
	AttrClusterquality string `xml:" clusterquality,attr"  json:",omitempty"`
	AttrContinent string `xml:" continent,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
	AssetHigh *AssetHigh `xml:" high,omitempty" json:"high,omitempty"`
	AssetLow *AssetLow `xml:" low,omitempty" json:"low,omitempty"`
	AssetMedium *AssetMedium `xml:" medium,omitempty" json:"medium,omitempty"`
}

type AssetPresets struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetPreset []*AssetPreset `xml:" preset,omitempty" json:"preset,omitempty"`
}

type AssetRock struct {
	AttrAmount string `xml:" amount,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
}

type AssetWood struct {
	AttrAmount string `xml:" amount,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
}


