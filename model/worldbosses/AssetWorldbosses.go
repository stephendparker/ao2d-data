package worldbosses

type AssetChidleyRoot314159 struct {
	AssetWorldbosses *AssetWorldbosses `xml:" worldbosses,omitempty" json:"worldbosses,omitempty"`
}

type AssetSpawnTime struct {
	AttrUtctime string `xml:" utctime,attr"  json:",omitempty"`
}

type AssetSpawnTimes struct {
	AttrCalculateusingclustertimezone string `xml:" calculateusingclustertimezone,attr"  json:",omitempty"`
	AttrStartdate string `xml:" startdate,attr"  json:",omitempty"`
	AssetDay []*AssetDay `xml:" day,omitempty" json:"day,omitempty"`
}

type AssetBoss struct {
	AttrDescription string `xml:" description,attr"  json:",omitempty"`
	AttrMinimapiconempty string `xml:" minimapiconempty,attr"  json:",omitempty"`
	AttrMinimapiconfull string `xml:" minimapiconfull,attr"  json:",omitempty"`
	AttrMob string `xml:" mob,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrTitle string `xml:" title,attr"  json:",omitempty"`
	AttrWorldmapprefab string `xml:" worldmapprefab,attr"  json:",omitempty"`
	AssetSpawnTimes *AssetSpawnTimes `xml:" SpawnTimes,omitempty" json:"SpawnTimes,omitempty"`
}

type AssetDay struct {
	AttrNumber string `xml:" number,attr"  json:",omitempty"`
	AssetSpawnTime []*AssetSpawnTime `xml:" SpawnTime,omitempty" json:"SpawnTime,omitempty"`
}

type AssetWorldbosses struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetBoss []*AssetBoss `xml:" boss,omitempty" json:"boss,omitempty"`
}


