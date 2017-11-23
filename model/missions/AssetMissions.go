package missions

type AssetChidleyRoot314159 struct {
	AssetMissions *AssetMissions `xml:" missions,omitempty" json:"missions,omitempty"`
}

type AssetAcquireitem struct {
	AttrCount string `xml:" count,attr"  json:",omitempty"`
	AttrItemid string `xml:" itemid,attr"  json:",omitempty"`
	AttrTextloca string `xml:" textloca,attr"  json:",omitempty"`
}

type AssetAgenttag struct {
	AttrTag string `xml:" tag,attr"  json:",omitempty"`
}

type AssetCompletionstep struct {
	AssetGotoagent *AssetGotoagent `xml:" gotoagent,omitempty" json:"gotoagent,omitempty"`
}

type AssetCraftitem struct {
	AttrCount string `xml:" count,attr"  json:",omitempty"`
	AttrItemid string `xml:" itemid,attr"  json:",omitempty"`
	AttrTextloca string `xml:" textloca,attr"  json:",omitempty"`
}

type AssetDistancefromagent struct {
	AttrMaxdistance string `xml:" maxdistance,attr"  json:",omitempty"`
	AttrMindistance string `xml:" mindistance,attr"  json:",omitempty"`
}

type AssetDistancetocluster struct {
	AttrMaxdistance string `xml:" maxdistance,attr"  json:",omitempty"`
	AttrMindistance string `xml:" mindistance,attr"  json:",omitempty"`
}

type AssetDynamicagent struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AssetAgenttag *AssetAgenttag `xml:" agenttag,omitempty" json:"agenttag,omitempty"`
	AssetDistancefromagent *AssetDistancefromagent `xml:" distancefromagent,omitempty" json:"distancefromagent,omitempty"`
}

type AssetDynamiccluster struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AssetDistancetocluster *AssetDistancetocluster `xml:" distancetocluster,omitempty" json:"distancetocluster,omitempty"`
	AssetPresentresource *AssetPresentresource `xml:" presentresource,omitempty" json:"presentresource,omitempty"`
}

type AssetGather struct {
	AttrCount string `xml:" count,attr"  json:",omitempty"`
	AttrItemid string `xml:" itemid,attr"  json:",omitempty"`
}

type AssetGathersilver struct {
	AttrTextloca string `xml:" textloca,attr"  json:",omitempty"`
	AttrValue string `xml:" value,attr"  json:",omitempty"`
}

type AssetGotoagent struct {
	AttrAgent string `xml:" agent,attr"  json:",omitempty"`
	AssetItem *AssetItem `xml:" item,omitempty" json:"item,omitempty"`
}

type AssetItem struct {
	AttrCount string `xml:" count,attr"  json:",omitempty"`
	AttrItemid string `xml:" itemid,attr"  json:",omitempty"`
	AttrTextloca string `xml:" textloca,attr"  json:",omitempty"`
}

type AssetItemreward struct {
	AttrCount string `xml:" count,attr"  json:",omitempty"`
	AttrItemid string `xml:" itemid,attr"  json:",omitempty"`
}

type AssetMission struct {
	AttrCompletetext string `xml:" completetext,attr"  json:",omitempty"`
	AttrFaction string `xml:" faction,attr"  json:",omitempty"`
	AttrMaxlevel string `xml:" maxlevel,attr"  json:",omitempty"`
	AttrMinlevel string `xml:" minlevel,attr"  json:",omitempty"`
	AttrMissionslotcooldown string `xml:" missionslotcooldown,attr"  json:",omitempty"`
	AttrMissiontimeout string `xml:" missiontimeout,attr"  json:",omitempty"`
	AttrOffertext string `xml:" offertext,attr"  json:",omitempty"`
	AttrOffertimeout string `xml:" offertimeout,attr"  json:",omitempty"`
	AttrOnceonly string `xml:" onceonly,attr"  json:",omitempty"`
	AttrPreviousmissions string `xml:" previousmissions,attr"  json:",omitempty"`
	AttrPriority string `xml:" priority,attr"  json:",omitempty"`
	AttrProgresstext string `xml:" progresstext,attr"  json:",omitempty"`
	AttrRequiresonlyonepreviousmission string `xml:" requiresonlyonepreviousmission,attr"  json:",omitempty"`
	AttrSelectmissions string `xml:" selectmissions,attr"  json:",omitempty"`
	AttrSilverrequired string `xml:" silverrequired,attr"  json:",omitempty"`
	AttrTags string `xml:" tags,attr"  json:",omitempty"`
	AttrTitle string `xml:" title,attr"  json:",omitempty"`
	AttrUicategory string `xml:" uicategory,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AssetRewards *AssetRewards `xml:" rewards,omitempty" json:"rewards,omitempty"`
	AssetTasks *AssetTasks `xml:" tasks,omitempty" json:"tasks,omitempty"`
	AssetVariables *AssetVariables `xml:" variables,omitempty" json:"variables,omitempty"`
}

type AssetMissions struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetMission []*AssetMission `xml:" mission,omitempty" json:"mission,omitempty"`
}

type AssetMissionstep struct {
	AssetAcquireitem []*AssetAcquireitem `xml:" acquireitem,omitempty" json:"acquireitem,omitempty"`
	AssetCraftitem []*AssetCraftitem `xml:" craftitem,omitempty" json:"craftitem,omitempty"`
	AssetGather []*AssetGather `xml:" gather,omitempty" json:"gather,omitempty"`
	AssetGathersilver *AssetGathersilver `xml:" gathersilver,omitempty" json:"gathersilver,omitempty"`
	AssetUsefunction *AssetUsefunction `xml:" usefunction,omitempty" json:"usefunction,omitempty"`
}

type AssetPresentresource struct {
	AttrResourceid string `xml:" resourceid,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
}

type AssetRewards struct {
	AssetItemreward *AssetItemreward `xml:" itemreward,omitempty" json:"itemreward,omitempty"`
}

type AssetTasks struct {
	AssetCompletionstep *AssetCompletionstep `xml:" completionstep,omitempty" json:"completionstep,omitempty"`
	AssetMissionstep *AssetMissionstep `xml:" missionstep,omitempty" json:"missionstep,omitempty"`
}

type AssetUsefunction struct {
	AttrFunction string `xml:" function,attr"  json:",omitempty"`
}

type AssetVariables struct {
	AssetDynamicagent *AssetDynamicagent `xml:" dynamicagent,omitempty" json:"dynamicagent,omitempty"`
	AssetDynamiccluster []*AssetDynamiccluster `xml:" dynamiccluster,omitempty" json:"dynamiccluster,omitempty"`
}


