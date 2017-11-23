package expeditionagents

type AssetAudioInfo struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetChidleyRoot314159 struct {
	AssetExpeditionagents *AssetExpeditionagents `xml:" expeditionagents,omitempty" json:"expeditionagents,omitempty"`
}

type AssetExpedition struct {
	AttrCategory string `xml:" category,attr"  json:",omitempty"`
}

type AssetExpeditionagent struct {
	AttrNamegenerator string `xml:" namegenerator,attr"  json:",omitempty"`
	AttrUiavatar string `xml:" uiavatar,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AssetAudioInfo *AssetAudioInfo `xml:" AudioInfo,omitempty" json:"AudioInfo,omitempty"`
	AssetExpeditions *AssetExpeditions `xml:" expeditions,omitempty" json:"expeditions,omitempty"`
}

type AssetExpeditionagents struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetExpeditionagent []*AssetExpeditionagent `xml:" expeditionagent,omitempty" json:"expeditionagent,omitempty"`
}

type AssetExpeditions struct {
	AssetExpedition []*AssetExpedition `xml:" expedition,omitempty" json:"expedition,omitempty"`
}


