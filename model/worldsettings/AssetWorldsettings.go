package worldsettings

type AssetAOHyphenWorldSettings struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetArchetypes *AssetArchetypes `xml:" archetypes,omitempty" json:"archetypes,omitempty"`
	AssetClustertypes *AssetClustertypes `xml:" clustertypes,omitempty" json:"clustertypes,omitempty"`
	AssetHellclusters *AssetHellclusters `xml:" hellclusters,omitempty" json:"hellclusters,omitempty"`
	AssetHellgatemobtypes *AssetHellgatemobtypes `xml:" hellgatemobtypes,omitempty" json:"hellgatemobtypes,omitempty"`
	AssetIslands *AssetIslands `xml:" islands,omitempty" json:"islands,omitempty"`
	AssetTimeregions *AssetTimeregions `xml:" timeregions,omitempty" json:"timeregions,omitempty"`
	AssetWorldmap *AssetWorldmap `xml:" worldmap,omitempty" json:"worldmap,omitempty"`
}

type AssetChidleyRoot314159 struct {
	AssetAOHyphenWorldSettings *AssetAOHyphenWorldSettings `xml:" AO-WorldSettings,omitempty" json:"AO-WorldSettings,omitempty"`
}

type AssetArchetype struct {
	AttrAllowtrade string `xml:" allowtrade,attr"  json:",omitempty"`
	AttrAnimalgrowthfactor string `xml:" animalgrowthfactor,attr"  json:",omitempty"`
	AttrBuildingcraftcapacityregenfactor string `xml:" buildingcraftcapacityregenfactor,attr"  json:",omitempty"`
	AttrClusterdangerbonustype string `xml:" clusterdangerbonustype,attr"  json:",omitempty"`
	AttrCraftefficiencybonus string `xml:" craftefficiencybonus,attr"  json:",omitempty"`
	AttrCropyieldfactor string `xml:" cropyieldfactor,attr"  json:",omitempty"`
	AttrDurabilitydecaydivider string `xml:" durabilitydecaydivider,attr"  json:",omitempty"`
	AttrItempowerscalingfactor string `xml:" itempowerscalingfactor,attr"  json:",omitempty"`
	AttrItempowersoftcap string `xml:" itempowersoftcap,attr"  json:",omitempty"`
	AttrKnockdownbuffaftergetup string `xml:" knockdownbuffaftergetup,attr"  json:",omitempty"`
	AttrLogoutseconds string `xml:" logoutseconds,attr"  json:",omitempty"`
	AttrMobsarelethal string `xml:" mobsarelethal,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrPlayersarelethal string `xml:" playersarelethal,attr"  json:",omitempty"`
	AttrPvprule string `xml:" pvprule,attr"  json:",omitempty"`
	AttrPvpwarningtype string `xml:" pvpwarningtype,attr"  json:",omitempty"`
	AttrRequiredreputationtoenter string `xml:" requiredreputationtoenter,attr"  json:",omitempty"`
	AttrShowfriendlycount string `xml:" showfriendlycount,attr"  json:",omitempty"`
	AttrShowhostilecount string `xml:" showhostilecount,attr"  json:",omitempty"`
	AttrShowpvpplayeronminimap string `xml:" showpvpplayeronminimap,attr"  json:",omitempty"`
	AttrSpellcooldownsonenteringseconds string `xml:" spellcooldownsonenteringseconds,attr"  json:",omitempty"`
	AttrUipvptype string `xml:" uipvptype,attr"  json:",omitempty"`
}

type AssetArchetypes struct {
	AssetArchetype []*AssetArchetype `xml:" archetype,omitempty" json:"archetype,omitempty"`
}

type AssetBanner struct {
	AttrResource string `xml:" resource,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
}

type AssetBanners struct {
	AssetBanner []*AssetBanner `xml:" banner,omitempty" json:"banner,omitempty"`
}

type AssetCastle struct {
	AttrGui string `xml:" gui,attr"  json:",omitempty"`
	AttrResource string `xml:" resource,attr"  json:",omitempty"`
}

type AssetCity struct {
	AttrGui string `xml:" gui,attr"  json:",omitempty"`
	AttrResource string `xml:" resource,attr"  json:",omitempty"`
}

type AssetCluster struct {
	AttrDescloca string `xml:" descloca,attr"  json:",omitempty"`
	AttrLevel string `xml:" level,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrPremiumrequirement string `xml:" premiumrequirement,attr"  json:",omitempty"`
	AttrSilverprice string `xml:" silverprice,attr"  json:",omitempty"`
}

type AssetClustertype struct {
	AttrAllowmounts string `xml:" allowmounts,attr"  json:",omitempty"`
	AttrArchetype string `xml:" archetype,attr"  json:",omitempty"`
	AttrEditorcolor string `xml:" editorcolor,attr"  json:",omitempty"`
	AttrEditorcolor2 string `xml:" editorcolor2,attr"  json:",omitempty"`
	AttrEscapedungeonspellavailable string `xml:" escapedungeonspellavailable,attr"  json:",omitempty"`
	AttrFurnitureobjectsallowed string `xml:" furnitureobjectsallowed,attr"  json:",omitempty"`
	AttrHellgateclustermessage string `xml:" hellgateclustermessage,attr"  json:",omitempty"`
	AttrHellgatemobminimapvisible string `xml:" hellgatemobminimapvisible,attr"  json:",omitempty"`
	AttrIsdungeon string `xml:" isdungeon,attr"  json:",omitempty"`
	AttrIsplayercity string `xml:" isplayercity,attr"  json:",omitempty"`
	AttrIsstartarea string `xml:" isstartarea,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrPersistsimulationtodatabase string `xml:" persistsimulationtodatabase,attr"  json:",omitempty"`
	AssetMobspawns *AssetMobspawns `xml:" mobspawns,omitempty" json:"mobspawns,omitempty"`
	AssetResourcespawns *AssetResourcespawns `xml:" resourcespawns,omitempty" json:"resourcespawns,omitempty"`
}

type AssetClustertypes struct {
	AssetClustertype []*AssetClustertype `xml:" clustertype,omitempty" json:"clustertype,omitempty"`
}

type AssetExit struct {
	AttrResource string `xml:" resource,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
}

type AssetExits struct {
	AssetExit []*AssetExit `xml:" exit,omitempty" json:"exit,omitempty"`
}

type AssetHellcluster struct {
	AttrAllowloggingout string `xml:" allowloggingout,attr"  json:",omitempty"`
	AttrBossspawndelayinminafterminibossdeath string `xml:" bossspawndelayinminafterminibossdeath,attr"  json:",omitempty"`
	AttrBossspawnicon string `xml:" bossspawnicon,attr"  json:",omitempty"`
	AttrBosswarningtext string `xml:" bosswarningtext,attr"  json:",omitempty"`
	AttrEnableenterwarningmessage string `xml:" enableenterwarningmessage,attr"  json:",omitempty"`
	AttrExitprefab string `xml:" exitprefab,attr"  json:",omitempty"`
	AttrGateentertitle string `xml:" gateentertitle,attr"  json:",omitempty"`
	AttrGateprefab string `xml:" gateprefab,attr"  json:",omitempty"`
	AttrHellbosstimer string `xml:" hellbosstimer,attr"  json:",omitempty"`
	AttrHellclusterconnectionchance string `xml:" hellclusterconnectionchance,attr"  json:",omitempty"`
	AttrHellclusterjointimelimit string `xml:" hellclusterjointimelimit,attr"  json:",omitempty"`
	AttrHellclusterkillsonclosing string `xml:" hellclusterkillsonclosing,attr"  json:",omitempty"`
	AttrHellclusterlevel string `xml:" hellclusterlevel,attr"  json:",omitempty"`
	AttrHellclustertimelimit string `xml:" hellclustertimelimit,attr"  json:",omitempty"`
	AttrHellgatedeath string `xml:" hellgatedeath,attr"  json:",omitempty"`
	AttrHellgateentrancetimer string `xml:" hellgateentrancetimer,attr"  json:",omitempty"`
	AttrHellgateexitcooldowninsec string `xml:" hellgateexitcooldowninsec,attr"  json:",omitempty"`
	AttrHellgateexittimer string `xml:" hellgateexittimer,attr"  json:",omitempty"`
	AttrHellgatejointime string `xml:" hellgatejointime,attr"  json:",omitempty"`
	AttrHellgateleavetime string `xml:" hellgateleavetime,attr"  json:",omitempty"`
	AttrHellgatelimit string `xml:" hellgatelimit,attr"  json:",omitempty"`
	AttrHellgatemindistance string `xml:" hellgatemindistance,attr"  json:",omitempty"`
	AttrHellgatemobchargepool string `xml:" hellgatemobchargepool,attr"  json:",omitempty"`
	AttrHellgatemobchargesadded string `xml:" hellgatemobchargesadded,attr"  json:",omitempty"`
	AttrHellgatemobchargetime string `xml:" hellgatemobchargetime,attr"  json:",omitempty"`
	AttrHellgatemobdespawntime string `xml:" hellgatemobdespawntime,attr"  json:",omitempty"`
	AttrHellgatemobspawnmaxamount string `xml:" hellgatemobspawnmaxamount,attr"  json:",omitempty"`
	AttrHellgatemobwavetimer string `xml:" hellgatemobwavetimer,attr"  json:",omitempty"`
	AttrHellgateplayerslots string `xml:" hellgateplayerslots,attr"  json:",omitempty"`
	AttrMinimapicon string `xml:" minimapicon,attr"  json:",omitempty"`
	AttrShowhellgatetimerwhenxminareleft string `xml:" showhellgatetimerwhenxminareleft,attr"  json:",omitempty"`
	AttrTeleportonknockdown string `xml:" teleportonknockdown,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
}

type AssetHellclusters struct {
	AssetHellcluster []*AssetHellcluster `xml:" hellcluster,omitempty" json:"hellcluster,omitempty"`
}

type AssetHellgatemob struct {
	AttrClustertype string `xml:" clustertype,attr"  json:",omitempty"`
	AttrFaction string `xml:" faction,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrWeight string `xml:" weight,attr"  json:",omitempty"`
}

type AssetHellgatemobs struct {
	AssetHellgatemob []*AssetHellgatemob `xml:" hellgatemob,omitempty" json:"hellgatemob,omitempty"`
}

type AssetHellgatemobtypes struct {
	AssetHellgatemobs *AssetHellgatemobs `xml:" hellgatemobs,omitempty" json:"hellgatemobs,omitempty"`
}

type AssetIsland struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AssetCluster []*AssetCluster `xml:" cluster,omitempty" json:"cluster,omitempty"`
}

type AssetIslands struct {
	AssetIsland []*AssetIsland `xml:" island,omitempty" json:"island,omitempty"`
}

type AssetMobspawn struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrTimefactor string `xml:" timefactor,attr"  json:",omitempty"`
}

type AssetMobspawns struct {
	AttrDefaultinitialchargemultiplicator string `xml:" defaultinitialchargemultiplicator,attr"  json:",omitempty"`
	AttrDefaultmaxchargemultiplicator string `xml:" defaultmaxchargemultiplicator,attr"  json:",omitempty"`
	AttrDefaulttimefactor string `xml:" defaulttimefactor,attr"  json:",omitempty"`
	AssetMobspawn []*AssetMobspawn `xml:" mobspawn,omitempty" json:"mobspawn,omitempty"`
}

type AssetResourcespawns struct {
	AttrDefaulttimefactor string `xml:" defaulttimefactor,attr"  json:",omitempty"`
	AssetResourcetier []*AssetResourcetier `xml:" resourcetier,omitempty" json:"resourcetier,omitempty"`
}

type AssetResourcetier struct {
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrTimefactor string `xml:" timefactor,attr"  json:",omitempty"`
}

type AssetResourcetype struct {
	AttrGui string `xml:" gui,attr"  json:",omitempty"`
	AttrResource string `xml:" resource,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
}

type AssetSettlement struct {
	AttrResource string `xml:" resource,attr"  json:",omitempty"`
}

type AssetSettlements struct {
	AssetSettlement []*AssetSettlement `xml:" settlement,omitempty" json:"settlement,omitempty"`
}

type AssetSiegecamp struct {
	AttrResource string `xml:" resource,attr"  json:",omitempty"`
}

type AssetTerritory struct {
	AttrGui string `xml:" gui,attr"  json:",omitempty"`
	AttrResource string `xml:" resource,attr"  json:",omitempty"`
	AssetResourcetype []*AssetResourcetype `xml:" resourcetype,omitempty" json:"resourcetype,omitempty"`
}

type AssetTimeregion struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrStarttime string `xml:" starttime,attr"  json:",omitempty"`
	AttrTimezone string `xml:" timezone,attr"  json:",omitempty"`
}

type AssetTimeregions struct {
	AssetTimeregion []*AssetTimeregion `xml:" timeregion,omitempty" json:"timeregion,omitempty"`
}

type AssetWorldmap struct {
	AssetBanners *AssetBanners `xml:" banners,omitempty" json:"banners,omitempty"`
	AssetCastle *AssetCastle `xml:" castle,omitempty" json:"castle,omitempty"`
	AssetCity *AssetCity `xml:" city,omitempty" json:"city,omitempty"`
	AssetExits *AssetExits `xml:" exits,omitempty" json:"exits,omitempty"`
	AssetSettlements *AssetSettlements `xml:" settlements,omitempty" json:"settlements,omitempty"`
	AssetSiegecamp *AssetSiegecamp `xml:" siegecamp,omitempty" json:"siegecamp,omitempty"`
	AssetTerritory *AssetTerritory `xml:" territory,omitempty" json:"territory,omitempty"`
}


