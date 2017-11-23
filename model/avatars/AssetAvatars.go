package avatars

type AssetChidleyRoot314159 struct {
	AssetAvatars *AssetAvatars `xml:" avatars,omitempty" json:"avatars,omitempty"`
}

type AssetAvatar struct {
	AttrHidden string `xml:" hidden,attr"  json:",omitempty"`
	AttrLocked string `xml:" locked,attr"  json:",omitempty"`
	AttrMinfounderneeded string `xml:" minfounderneeded,attr"  json:",omitempty"`
	AttrMinstarterneeded string `xml:" minstarterneeded,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetAvatarring struct {
	AttrLocked string `xml:" locked,attr"  json:",omitempty"`
	AttrMinfounderneeded string `xml:" minfounderneeded,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetAvatarrings struct {
	AssetAvatarring []*AssetAvatarring `xml:" avatarring,omitempty" json:"avatarring,omitempty"`
}

type AssetAvatars struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetAvatar []*AssetAvatar `xml:" avatar,omitempty" json:"avatar,omitempty"`
	AssetAvatarrings *AssetAvatarrings `xml:" avatarrings,omitempty" json:"avatarrings,omitempty"`
	AssetAvatars *AssetAvatars `xml:" avatars,omitempty" json:"avatars,omitempty"`
}


