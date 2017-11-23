package randomspawnbehaviors

type AssetChidleyRoot314159 struct {
	AssetRandomspawnbehaviors *AssetRandomspawnbehaviors `xml:" randomspawnbehaviors,omitempty" json:"randomspawnbehaviors,omitempty"`
}

type AssetClustertype struct {
	AttrCharges string `xml:" charges,attr"  json:",omitempty"`
	AttrRespawntimesecondsmax string `xml:" respawntimesecondsmax,attr"  json:",omitempty"`
	AttrRespawntimesecondsmin string `xml:" respawntimesecondsmin,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AssetMob []*AssetMob `xml:" mob,omitempty" json:"mob,omitempty"`
}

type AssetDefault struct {
	AttrCharges string `xml:" charges,attr"  json:",omitempty"`
	AttrRespawntimesecondsmax string `xml:" respawntimesecondsmax,attr"  json:",omitempty"`
	AttrRespawntimesecondsmin string `xml:" respawntimesecondsmin,attr"  json:",omitempty"`
	AssetMob []*AssetMob `xml:" mob,omitempty" json:"mob,omitempty"`
}

type AssetMob struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrWeight string `xml:" weight,attr"  json:",omitempty"`
}

type AssetRandomspawnbehaviors struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetSpawnbehavior []*AssetSpawnbehavior `xml:" spawnbehavior,omitempty" json:"spawnbehavior,omitempty"`
}

type AssetSpawnbehavior struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AssetClustertype []*AssetClustertype `xml:" clustertype,omitempty" json:"clustertype,omitempty"`
	AssetDefault *AssetDefault `xml:" default,omitempty" json:"default,omitempty"`
}


