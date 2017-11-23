package resources

type AssetAOHyphenResources struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetHarvestables *AssetHarvestables `xml:" Harvestables,omitempty" json:"Harvestables,omitempty"`
	AssetResources *AssetResources `xml:" Resources,omitempty" json:"Resources,omitempty"`
}

type AssetCharge struct {
	AttrGfxstate string `xml:" gfxstate,attr"  json:",omitempty"`
	AttrLevel string `xml:" level,attr"  json:",omitempty"`
	AttrRespawncharges string `xml:" respawncharges,attr"  json:",omitempty"`
	AttrRespawnfactormax string `xml:" respawnfactormax,attr"  json:",omitempty"`
	AttrRespawnfactormin string `xml:" respawnfactormin,attr"  json:",omitempty"`
	AttrYield string `xml:" yield,attr"  json:",omitempty"`
}

type AssetChidleyRoot314159 struct {
	AssetAOHyphenResources *AssetAOHyphenResources `xml:" AO-Resources,omitempty" json:"AO-Resources,omitempty"`
}

type AssetHarvestable struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrResource string `xml:" resource,attr"  json:",omitempty"`
	AssetTier []*AssetTier `xml:" Tier,omitempty" json:"Tier,omitempty"`
	AssetToolModifier *AssetToolModifier `xml:" ToolModifier,omitempty" json:"ToolModifier,omitempty"`
}

type AssetHarvestables struct {
	AssetHarvestable []*AssetHarvestable `xml:" Harvestable,omitempty" json:"Harvestable,omitempty"`
}

type AssetModifier struct {
	AttrTierdifference string `xml:" tierdifference,attr"  json:",omitempty"`
	AttrTimefactor string `xml:" timefactor,attr"  json:",omitempty"`
}

type AssetRareState struct {
	AttrItem string `xml:" item,attr"  json:",omitempty"`
	AttrState string `xml:" state,attr"  json:",omitempty"`
}

type AssetResource struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AssetResourceTier []*AssetResourceTier `xml:" ResourceTier,omitempty" json:"ResourceTier,omitempty"`
}

type AssetResourceTier struct {
	AttrFamevalue string `xml:" famevalue,attr"  json:",omitempty"`
	AttrResourcevalue string `xml:" resourcevalue,attr"  json:",omitempty"`
	AttrValue string `xml:" value,attr"  json:",omitempty"`
}

type AssetResources struct {
	AssetResource []*AssetResource `xml:" Resource,omitempty" json:"Resource,omitempty"`
}

type AssetTier struct {
	AttrDecaytimeseconds string `xml:" decaytimeseconds,attr"  json:",omitempty"`
	AttrDecaytimewhenexhaustedseconds string `xml:" decaytimewhenexhaustedseconds,attr"  json:",omitempty"`
	AttrHarvesttimeseconds string `xml:" harvesttimeseconds,attr"  json:",omitempty"`
	AttrIsscaled string `xml:" isscaled,attr"  json:",omitempty"`
	AttrItem string `xml:" item,attr"  json:",omitempty"`
	AttrMaxchargesperharvest string `xml:" maxchargesperharvest,attr"  json:",omitempty"`
	AttrNotooltimefactor string `xml:" notooltimefactor,attr"  json:",omitempty"`
	AttrRequirestool string `xml:" requirestool,attr"  json:",omitempty"`
	AttrRespawntimeseconds string `xml:" respawntimeseconds,attr"  json:",omitempty"`
	AttrStartcharges string `xml:" startcharges,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrTile string `xml:" tile,attr"  json:",omitempty"`
	AttrTilepremium string `xml:" tilepremium,attr"  json:",omitempty"`
	AssetCharge []*AssetCharge `xml:" Charge,omitempty" json:"Charge,omitempty"`
	AssetRareState []*AssetRareState `xml:" RareState,omitempty" json:"RareState,omitempty"`
}

type AssetToolModifier struct {
	AssetModifier []*AssetModifier `xml:" Modifier,omitempty" json:"Modifier,omitempty"`
}


