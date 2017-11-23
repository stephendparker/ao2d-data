package agents

type AssetAudioInfo struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetChidleyRoot314159 struct {
	AssetAgents *AssetAgents `xml:" agents,omitempty" json:"agents,omitempty"`
}

type AssetAgent struct {
	AttrFaction string `xml:" faction,attr"  json:",omitempty"`
	AttrLevel string `xml:" level,attr"  json:",omitempty"`
	AttrMaxsilverpool string `xml:" maxsilverpool,attr"  json:",omitempty"`
	AttrNamegenerator string `xml:" namegenerator,attr"  json:",omitempty"`
	AttrPrefab string `xml:" prefab,attr"  json:",omitempty"`
	AttrSecondspersilver string `xml:" secondspersilver,attr"  json:",omitempty"`
	AttrTags string `xml:" tags,attr"  json:",omitempty"`
	AttrUiavatar string `xml:" uiavatar,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AssetAudioInfo *AssetAudioInfo `xml:" AudioInfo,omitempty" json:"AudioInfo,omitempty"`
	AssetMissionslot *AssetMissionslot `xml:" missionslot,omitempty" json:"missionslot,omitempty"`
}

type AssetAgents struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetAgent []*AssetAgent `xml:" agent,omitempty" json:"agent,omitempty"`
}

type AssetMissionslot struct {
	AttrTags string `xml:" tags,attr"  json:",omitempty"`
}


