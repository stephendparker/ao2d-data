package emotes

type AssetChidleyRoot314159 struct {
	AssetEmoteData *AssetEmoteData `xml:" EmoteData,omitempty" json:"EmoteData,omitempty"`
}

type AssetEmote struct {
	AttrAllow_target string `xml:" allow_target,attr"  json:",omitempty"`
	AttrAllow_while_moving string `xml:" allow_while_moving,attr"  json:",omitempty"`
	AttrAnimation string `xml:" animation,attr"  json:",omitempty"`
	AttrLoop string `xml:" loop,attr"  json:",omitempty"`
	AttrMessage string `xml:" message,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrTarget_message string `xml:" target_message,attr"  json:",omitempty"`
}

type AssetEmoteData struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetEmote []*AssetEmote `xml:" Emote,omitempty" json:"Emote,omitempty"`
}


