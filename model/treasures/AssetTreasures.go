package treasures

type AssetChidleyRoot314159 struct {
	AssetTreasures *AssetTreasures `xml:" treasures,omitempty" json:"treasures,omitempty"`
}

type AssetSpawnTime struct {
	AttrUtctime string `xml:" utctime,attr"  json:",omitempty"`
}

type AssetSpawnTimes struct {
	AttrCalculateusingclustertimezone string `xml:" calculateusingclustertimezone,attr"  json:",omitempty"`
	AttrStartdate string `xml:" startdate,attr"  json:",omitempty"`
	AssetDay []*AssetDay `xml:" day,omitempty" json:"day,omitempty"`
}

type AssetDay struct {
	AttrNumber string `xml:" number,attr"  json:",omitempty"`
	AssetSpawnTime []*AssetSpawnTime `xml:" SpawnTime,omitempty" json:"SpawnTime,omitempty"`
}

type AssetTreasure struct {
	AttrDescription string `xml:" description,attr"  json:",omitempty"`
	AttrLootlist string `xml:" lootlist,attr"  json:",omitempty"`
	AttrMaxsilverreward string `xml:" maxsilverreward,attr"  json:",omitempty"`
	AttrMinimapiconempty string `xml:" minimapiconempty,attr"  json:",omitempty"`
	AttrMinimapiconfull string `xml:" minimapiconfull,attr"  json:",omitempty"`
	AttrMinsilverreward string `xml:" minsilverreward,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrOpeningtimeinsec string `xml:" openingtimeinsec,attr"  json:",omitempty"`
	AttrRequirescastleownership string `xml:" requirescastleownership,attr"  json:",omitempty"`
	AttrTitle string `xml:" title,attr"  json:",omitempty"`
	AttrWorldmapprefab string `xml:" worldmapprefab,attr"  json:",omitempty"`
	AssetSpawnTimes *AssetSpawnTimes `xml:" SpawnTimes,omitempty" json:"SpawnTimes,omitempty"`
}

type AssetTreasures struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetTreasure []*AssetTreasure `xml:" treasure,omitempty" json:"treasure,omitempty"`
}


