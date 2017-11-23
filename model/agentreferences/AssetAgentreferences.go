package agentreferences

type AssetChidleyRoot314159 struct {
	AssetAgentreferences *AssetAgentreferences `xml:" agentreferences,omitempty" json:"agentreferences,omitempty"`
}

type AssetAgentreference struct {
	AttrAgenttype string `xml:" agenttype,attr"  json:",omitempty"`
	AttrCluster string `xml:" cluster,attr"  json:",omitempty"`
	AttrReferencename string `xml:" referencename,attr"  json:",omitempty"`
}

type AssetAgentreferences struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetAgentreference []*AssetAgentreference `xml:" agentreference,omitempty" json:"agentreference,omitempty"`
}


