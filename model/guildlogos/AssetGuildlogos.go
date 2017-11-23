package guildlogos

type AssetChidleyRoot314159 struct {
	AssetGuildLogoData *AssetGuildLogoData `xml:" GuildLogoData,omitempty" json:"GuildLogoData,omitempty"`
}

type AssetGuildLogoData struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetSchemas *AssetSchemas `xml:" Schemas,omitempty" json:"Schemas,omitempty"`
	AssetSymbolColors *AssetSymbolColors `xml:" SymbolColors,omitempty" json:"SymbolColors,omitempty"`
	AssetSymbols *AssetSymbols `xml:" Symbols,omitempty" json:"Symbols,omitempty"`
}

type AssetQuadrantMarked struct {
	AttrId string `xml:" id,attr"  json:",omitempty"`
}

type AssetSchema struct {
	AttrDisplayname string `xml:" displayname,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AssetQuadrantMarked []*AssetQuadrantMarked `xml:" QuadrantMarked,omitempty" json:"QuadrantMarked,omitempty"`
}

type AssetSchemas struct {
	AttrColors string `xml:" colors,attr"  json:",omitempty"`
	AttrDefaultcolor string `xml:" defaultcolor,attr"  json:",omitempty"`
	AttrQuadrants string `xml:" quadrants,attr"  json:",omitempty"`
	AssetSchema []*AssetSchema `xml:" Schema,omitempty" json:"Schema,omitempty"`
}

type AssetSymbol struct {
	AttrDisplayname string `xml:" displayname,attr"  json:",omitempty"`
	AttrLocked string `xml:" locked,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
}

type AssetSymbolColor struct {
	AttrB string `xml:" b,attr"  json:",omitempty"`
	AttrG string `xml:" g,attr"  json:",omitempty"`
	AttrR string `xml:" r,attr"  json:",omitempty"`
}

type AssetSymbolColors struct {
	AssetSymbolColor []*AssetSymbolColor `xml:" SymbolColor,omitempty" json:"SymbolColor,omitempty"`
}

type AssetSymbols struct {
	AttrMaxoffsety string `xml:" maxoffsety,attr"  json:",omitempty"`
	AttrMaxscale string `xml:" maxscale,attr"  json:",omitempty"`
	AttrMinoffsety string `xml:" minoffsety,attr"  json:",omitempty"`
	AttrMinscale string `xml:" minscale,attr"  json:",omitempty"`
	AssetSymbol []*AssetSymbol `xml:" Symbol,omitempty" json:"Symbol,omitempty"`
}


