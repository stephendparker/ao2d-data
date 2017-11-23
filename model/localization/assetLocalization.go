package localization

type AssetChidleyRoot314159 struct {
	AssetTmx *AssetTmx `xml:" tmx,omitempty" json:"tmx,omitempty"`
}

type AssetBody struct {
	AssetTu []*AssetTu `xml:" tu,omitempty" json:"tu,omitempty"`
}

type AssetSeg struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type AssetTmx struct {
	AttrVersion string     `xml:" version,attr"  json:",omitempty"`
	AssetBody   *AssetBody `xml:" body,omitempty" json:"body,omitempty"`
}

type AssetTu struct {
	AttrTuid string      `xml:" tuid,attr"  json:",omitempty"`
	AssetTuv []*AssetTuv `xml:" tuv,omitempty" json:"tuv,omitempty"`
}

type AssetTuv struct {
	// manually removing this - bug in Chidley
	//	AttrHttp://www.w3.org/XML/1998/namespaceLang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr"  json:",omitempty"`
	AssetSeg *AssetSeg `xml:" seg,omitempty" json:"seg,omitempty"`
}
