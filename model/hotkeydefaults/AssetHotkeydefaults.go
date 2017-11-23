package hotkeydefaults

type AssetChidleyRoot314159 struct {
	AssetHotkeydefaults *AssetHotkeydefaults `xml:" hotkeydefaults,omitempty" json:"hotkeydefaults,omitempty"`
}

type AssetCategory struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AssetHotkey []*AssetHotkey `xml:" hotkey,omitempty" json:"hotkey,omitempty"`
}

type AssetHotkey struct {
	AttrAction string `xml:" action,attr"  json:",omitempty"`
	AttrKeycode string `xml:" keycode,attr"  json:",omitempty"`
	AttrModifier string `xml:" modifier,attr"  json:",omitempty"`
}

type AssetHotkeydefaults struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetCategory []*AssetCategory `xml:" category,omitempty" json:"category,omitempty"`
}


