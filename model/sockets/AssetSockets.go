package sockets

type AssetChidleyRoot314159 struct {
	AssetSocketInfos *AssetSocketInfos `xml:" SocketInfos,omitempty" json:"SocketInfos,omitempty"`
}

type AssetSocket struct {
	AttrBonename string `xml:" bonename,attr"  json:",omitempty"`
	AttrPositionOffset string `xml:" positionOffset,attr"  json:",omitempty"`
	AttrRotationOffset string `xml:" rotationOffset,attr"  json:",omitempty"`
	AttrScaling string `xml:" scaling,attr"  json:",omitempty"`
	AttrSocketname string `xml:" socketname,attr"  json:",omitempty"`
}

type AssetSocketInfo struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AssetSocket []*AssetSocket `xml:" Socket,omitempty" json:"Socket,omitempty"`
}

type AssetSocketInfos struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetSocketInfo []*AssetSocketInfo `xml:" SocketInfo,omitempty" json:"SocketInfo,omitempty"`
}


