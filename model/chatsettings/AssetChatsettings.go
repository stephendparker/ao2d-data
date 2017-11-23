package chatsettings

type AssetAllianceTab struct {
	AttrMaininputchannel string `xml:" maininputchannel,attr"  json:",omitempty"`
	AttrShowchannelnames string `xml:" showchannelnames,attr"  json:",omitempty"`
	AttrShowtimestamps string `xml:" showtimestamps,attr"  json:",omitempty"`
	AssetChannel []*AssetChannel `xml:" Channel,omitempty" json:"Channel,omitempty"`
}

type AssetChannel struct {
	AttrId string `xml:" id,attr"  json:",omitempty"`
}

type AssetChatSettings struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetChatUi *AssetChatUi `xml:" ChatUi,omitempty" json:"ChatUi,omitempty"`
}

type AssetChatUi struct {
	AttrMaxmuted string `xml:" maxmuted,attr"  json:",omitempty"`
	AttrMaxtabnamelength string `xml:" maxtabnamelength,attr"  json:",omitempty"`
	AttrMaxtabs string `xml:" maxtabs,attr"  json:",omitempty"`
	AttrSelectedtab string `xml:" selectedtab,attr"  json:",omitempty"`
	AssetAllianceTab *AssetAllianceTab `xml:" AllianceTab,omitempty" json:"AllianceTab,omitempty"`
	AssetDefaultTab []*AssetDefaultTab `xml:" DefaultTab,omitempty" json:"DefaultTab,omitempty"`
	AssetGuildTab *AssetGuildTab `xml:" GuildTab,omitempty" json:"GuildTab,omitempty"`
	AssetPartyTab *AssetPartyTab `xml:" PartyTab,omitempty" json:"PartyTab,omitempty"`
}

type AssetChidleyRoot314159 struct {
	AssetChatSettings *AssetChatSettings `xml:" ChatSettings,omitempty" json:"ChatSettings,omitempty"`
}

type AssetDefaultTab struct {
	AttrMaininputchannel string `xml:" maininputchannel,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrShowchannelnames string `xml:" showchannelnames,attr"  json:",omitempty"`
	AttrShowtimestamps string `xml:" showtimestamps,attr"  json:",omitempty"`
	AssetChannel []*AssetChannel `xml:" Channel,omitempty" json:"Channel,omitempty"`
}

type AssetGuildTab struct {
	AttrMaininputchannel string `xml:" maininputchannel,attr"  json:",omitempty"`
	AttrShowchannelnames string `xml:" showchannelnames,attr"  json:",omitempty"`
	AttrShowtimestamps string `xml:" showtimestamps,attr"  json:",omitempty"`
	AssetChannel []*AssetChannel `xml:" Channel,omitempty" json:"Channel,omitempty"`
}

type AssetPartyTab struct {
	AttrMaininputchannel string `xml:" maininputchannel,attr"  json:",omitempty"`
	AttrShowchannelnames string `xml:" showchannelnames,attr"  json:",omitempty"`
	AttrShowtimestamps string `xml:" showtimestamps,attr"  json:",omitempty"`
	AssetChannel []*AssetChannel `xml:" Channel,omitempty" json:"Channel,omitempty"`
}


