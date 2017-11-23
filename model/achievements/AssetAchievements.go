package achievements

type AssetChidleyRoot314159 struct {
	AssetAchievements *AssetAchievements `xml:" achievements,omitempty" json:"achievements,omitempty"`
}

type AssetAchievement struct {
	AttrId string `xml:" id,attr"  json:",omitempty"`
	AttrTrialmaxlevel string `xml:" trialmaxlevel,attr"  json:",omitempty"`
	AssetMasterylevels *AssetMasterylevels `xml:" masterylevels,omitempty" json:"masterylevels,omitempty"`
	AssetParentachievements *AssetParentachievements `xml:" parentachievements,omitempty" json:"parentachievements,omitempty"`
}

type AssetAchievements struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrVersion string `xml:" Version,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetAchievement []*AssetAchievement `xml:" achievement,omitempty" json:"achievement,omitempty"`
	AssetTemplate []*AssetTemplate `xml:" template,omitempty" json:"template,omitempty"`
	AssetTemplateachievement []*AssetTemplateachievement `xml:" templateachievement,omitempty" json:"templateachievement,omitempty"`
}

type AssetBonus struct {
	AttrAttribute string `xml:" attribute,attr"  json:",omitempty"`
	AttrBonus string `xml:" bonus,attr"  json:",omitempty"`
	AttrBufftype string `xml:" bufftype,attr"  json:",omitempty"`
	AttrMaxtier string `xml:" maxtier,attr"  json:",omitempty"`
	AttrMintier string `xml:" mintier,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
	AssetDescription *AssetDescription `xml:" description,omitempty" json:"description,omitempty"`
	AssetItempattern []*AssetItempattern `xml:" itempattern,omitempty" json:"itempattern,omitempty"`
	AssetResourcepattern *AssetResourcepattern `xml:" resourcepattern,omitempty" json:"resourcepattern,omitempty"`
}

type AssetDescription struct {
	AttrTag string `xml:" tag,attr"  json:",omitempty"`
}

type AssetEntry struct {
	AttrId string `xml:" id,attr"  json:",omitempty"`
}

type AssetEquipitem struct {
	AttrId string `xml:" id,attr"  json:",omitempty"`
}

type AssetItemlist struct {
	AssetItempattern []*AssetItempattern `xml:" itempattern,omitempty" json:"itempattern,omitempty"`
}

type AssetItempattern struct {
	AttrPattern string `xml:" pattern,attr"  json:",omitempty"`
}

type AssetMasterylevel struct {
	AttrLearningpoints string `xml:" learningpoints,attr"  json:",omitempty"`
	AttrProgresstospendlp string `xml:" progresstospendlp,attr"  json:",omitempty"`
	AssetMissions *AssetMissions `xml:" missions,omitempty" json:"missions,omitempty"`
	AssetRewardgroups *AssetRewardgroups `xml:" rewardgroups,omitempty" json:"rewardgroups,omitempty"`
}

type AssetMasterylevels struct {
	AssetMasterylevel *AssetMasterylevel `xml:" masterylevel,omitempty" json:"masterylevel,omitempty"`
}

type AssetMission struct {
	AttrCluster string `xml:" cluster,attr"  json:",omitempty"`
	AttrItemid string `xml:" itemid,attr"  json:",omitempty"`
	AttrMintier string `xml:" mintier,attr"  json:",omitempty"`
	AttrMobid string `xml:" mobid,attr"  json:",omitempty"`
	AttrResourcetype string `xml:" resourcetype,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
	AttrValue string `xml:" value,attr"  json:",omitempty"`
	AssetEquipitem []*AssetEquipitem `xml:" equipitem,omitempty" json:"equipitem,omitempty"`
	AssetValiditem []*AssetValiditem `xml:" validitem,omitempty" json:"validitem,omitempty"`
}

type AssetMissions struct {
	AssetMission *AssetMission `xml:" mission,omitempty" json:"mission,omitempty"`
}

type AssetOverridemasterylevel struct {
	AttrLevel string `xml:" level,attr"  json:",omitempty"`
	AssetRewards *AssetRewards `xml:" rewards,omitempty" json:"rewards,omitempty"`
}

type AssetParentachievements struct {
	AttrAllrequired string `xml:" allrequired,attr"  json:",omitempty"`
	AssetAchievement []*AssetAchievement `xml:" achievement,omitempty" json:"achievement,omitempty"`
}

type AssetResourcepattern struct {
	AttrPattern string `xml:" pattern,attr"  json:",omitempty"`
}

type AssetReward struct {
	AttrBonus string `xml:" bonus,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
	AssetEntry []*AssetEntry `xml:" entry,omitempty" json:"entry,omitempty"`
}

type AssetRewardgroup struct {
	AssetReward []*AssetReward `xml:" reward,omitempty" json:"reward,omitempty"`
}

type AssetRewardgroups struct {
	AssetRewardgroup []*AssetRewardgroup `xml:" rewardgroup,omitempty" json:"rewardgroup,omitempty"`
}

type AssetRewards struct {
	AssetBonus []*AssetBonus `xml:" bonus,omitempty" json:"bonus,omitempty"`
	AssetUnlock []*AssetUnlock `xml:" unlock,omitempty" json:"unlock,omitempty"`
}

type AssetTemplate struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrStructure string `xml:" structure,attr"  json:",omitempty"`
	Text string `xml:",chardata" json:",omitempty"`
}

type AssetTemplateachievement struct {
	AttrFamemultiplier string `xml:" famemultiplier,attr"  json:",omitempty"`
	AttrId string `xml:" id,attr"  json:",omitempty"`
	AttrLpmultiplier string `xml:" lpmultiplier,attr"  json:",omitempty"`
	AttrMissiontype string `xml:" missiontype,attr"  json:",omitempty"`
	AttrTrialmaxlevel string `xml:" trialmaxlevel,attr"  json:",omitempty"`
	AttrUsetemplate string `xml:" usetemplate,attr"  json:",omitempty"`
	AssetItemlist *AssetItemlist `xml:" itemlist,omitempty" json:"itemlist,omitempty"`
	AssetOverridemasterylevel *AssetOverridemasterylevel `xml:" overridemasterylevel,omitempty" json:"overridemasterylevel,omitempty"`
	AssetParentachievements *AssetParentachievements `xml:" parentachievements,omitempty" json:"parentachievements,omitempty"`
	AssetRewards *AssetRewards `xml:" rewards,omitempty" json:"rewards,omitempty"`
}

type AssetUnlock struct {
	AttrMaxtier string `xml:" maxtier,attr"  json:",omitempty"`
	AttrMintier string `xml:" mintier,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
	AssetItempattern []*AssetItempattern `xml:" itempattern,omitempty" json:"itempattern,omitempty"`
}

type AssetValiditem struct {
	AttrId string `xml:" id,attr"  json:",omitempty"`
}


