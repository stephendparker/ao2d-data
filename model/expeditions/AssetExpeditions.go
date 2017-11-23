package expeditions

type AssetChidleyRoot314159 struct {
	AssetExpeditions *AssetExpeditions `xml:" expeditions,omitempty" json:"expeditions,omitempty"`
}

type AssetExpedition struct {
	AttrCategory string `xml:" category,attr"  json:",omitempty"`
	AttrCluster string `xml:" cluster,attr"  json:",omitempty"`
	AttrTimeout string `xml:" timeout,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AssetMission *AssetMission `xml:" mission,omitempty" json:"mission,omitempty"`
	AssetNarrator *AssetNarrator `xml:" narrator,omitempty" json:"narrator,omitempty"`
}

type AssetExpeditions struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetExpedition []*AssetExpedition `xml:" expedition,omitempty" json:"expedition,omitempty"`
}

type AssetKillmob struct {
	AttrCount string `xml:" count,attr"  json:",omitempty"`
	AttrMobid string `xml:" mobid,attr"  json:",omitempty"`
}

type AssetMission struct {
	AttrExtrotext string `xml:" extrotext,attr"  json:",omitempty"`
	AttrIntrotext string `xml:" introtext,attr"  json:",omitempty"`
	AttrTitle string `xml:" title,attr"  json:",omitempty"`
	AssetStep []*AssetStep `xml:" step,omitempty" json:"step,omitempty"`
}

type AssetNarrator struct {
	AttrNamegenerator string `xml:" namegenerator,attr"  json:",omitempty"`
	AttrUiavatar string `xml:" uiavatar,attr"  json:",omitempty"`
}

type AssetStep struct {
	AssetKillmob []*AssetKillmob `xml:" killmob,omitempty" json:"killmob,omitempty"`
}


