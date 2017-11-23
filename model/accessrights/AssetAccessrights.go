package accessrights

type AssetChidleyRoot314159 struct {
	AssetAccessrights *AssetAccessrights `xml:" accessrights,omitempty" json:"accessrights,omitempty"`
}

type AssetAccessclass struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AssetPreset []*AssetPreset `xml:" preset,omitempty" json:"preset,omitempty"`
	AssetRight []*AssetRight `xml:" right,omitempty" json:"right,omitempty"`
	AssetRole []*AssetRole `xml:" role,omitempty" json:"role,omitempty"`
}

type AssetAccessrights struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetAccessclass []*AssetAccessclass `xml:" accessclass,omitempty" json:"accessclass,omitempty"`
}

type AssetPreset struct {
	AttrDescription string `xml:" description,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AssetUserpreset []*AssetUserpreset `xml:" userpreset,omitempty" json:"userpreset,omitempty"`
}

type AssetRight struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrRoles string `xml:" roles,attr"  json:",omitempty"`
}

type AssetRole struct {
	AttrDescription string `xml:" description,attr"  json:",omitempty"`
	AttrDisplayname string `xml:" displayname,attr"  json:",omitempty"`
	AttrManagesroles string `xml:" managesroles,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetUserpreset struct {
	AttrRole string `xml:" role,attr"  json:",omitempty"`
	AttrWho string `xml:" who,attr"  json:",omitempty"`
}


