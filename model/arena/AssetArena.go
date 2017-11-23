package arena

type AssetChidleyRoot314159 struct {
	AssetArena *AssetArena `xml:" arena,omitempty" json:"arena,omitempty"`
}

type AssetArena struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetArenacategories *AssetArenacategories `xml:" arenacategories,omitempty" json:"arenacategories,omitempty"`
	AssetArenaregistrationtypes *AssetArenaregistrationtypes `xml:" arenaregistrationtypes,omitempty" json:"arenaregistrationtypes,omitempty"`
	AssetArenarules *AssetArenarules `xml:" arenarules,omitempty" json:"arenarules,omitempty"`
}

type AssetArenacategories struct {
	AssetArenacategory *AssetArenacategory `xml:" arenacategory,omitempty" json:"arenacategory,omitempty"`
}

type AssetArenacategory struct {
	AttrCluster string `xml:" cluster,attr"  json:",omitempty"`
	AttrDisplayname string `xml:" displayname,attr"  json:",omitempty"`
	AttrElok string `xml:" elok,attr"  json:",omitempty"`
	AttrElopersecond string `xml:" elopersecond,attr"  json:",omitempty"`
	AttrEloradiusbase string `xml:" eloradiusbase,attr"  json:",omitempty"`
	AttrGroupeloweight string `xml:" groupeloweight,attr"  json:",omitempty"`
	AttrMatchtype string `xml:" matchtype,attr"  json:",omitempty"`
	AttrMaxpotentials string `xml:" maxpotentials,attr"  json:",omitempty"`
	AttrMinitempoweraverage string `xml:" minitempoweraverage,attr"  json:",omitempty"`
	AttrMinpotentials string `xml:" minpotentials,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrRuleset string `xml:" ruleset,attr"  json:",omitempty"`
	AttrScorefirsthealer string `xml:" scorefirsthealer,attr"  json:",omitempty"`
	AttrScoregroupbalance string `xml:" scoregroupbalance,attr"  json:",omitempty"`
	AttrScoreperelodiff string `xml:" scoreperelodiff,attr"  json:",omitempty"`
	AttrScorepersecond string `xml:" scorepersecond,attr"  json:",omitempty"`
	AttrScoreunmatchedfullpremade string `xml:" scoreunmatchedfullpremade,attr"  json:",omitempty"`
	AttrStartingelo string `xml:" startingelo,attr"  json:",omitempty"`
	AttrSubeloweight string `xml:" subeloweight,attr"  json:",omitempty"`
	AssetGroupcompositions *AssetGroupcompositions `xml:" groupcompositions,omitempty" json:"groupcompositions,omitempty"`
	AssetNeededitemroles *AssetNeededitemroles `xml:" neededitemroles,omitempty" json:"neededitemroles,omitempty"`
	AssetRewards *AssetRewards `xml:" rewards,omitempty" json:"rewards,omitempty"`
}

type AssetArenaregistrationtype struct {
	AttrCategory string `xml:" category,attr"  json:",omitempty"`
	AttrElocategory string `xml:" elocategory,attr"  json:",omitempty"`
	AttrGroupsize string `xml:" groupsize,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetArenaregistrationtypes struct {
	AssetArenaregistrationtype []*AssetArenaregistrationtype `xml:" arenaregistrationtype,omitempty" json:"arenaregistrationtype,omitempty"`
}

type AssetArenarules struct {
	AssetArenaruleset *AssetArenaruleset `xml:" arenaruleset,omitempty" json:"arenaruleset,omitempty"`
}

type AssetArenaruleset struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrScorefactor string `xml:" scorefactor,attr"  json:",omitempty"`
	AttrTeamsize string `xml:" teamsize,attr"  json:",omitempty"`
}

type AssetGroupcomposition struct {
	AttrAllowedallies string `xml:" allowedallies,attr"  json:",omitempty"`
	AttrAllowedenemies string `xml:" allowedenemies,attr"  json:",omitempty"`
	AttrSize string `xml:" size,attr"  json:",omitempty"`
}

type AssetGroupcompositions struct {
	AssetGroupcomposition *AssetGroupcomposition `xml:" groupcomposition,omitempty" json:"groupcomposition,omitempty"`
}

type AssetItemreward struct {
	AttrCount string `xml:" count,attr"  json:",omitempty"`
	AttrEnchantmentlevel string `xml:" enchantmentlevel,attr"  json:",omitempty"`
	AttrItemid string `xml:" itemid,attr"  json:",omitempty"`
}

type AssetNeededitemroles struct {
	AssetRole *AssetRole `xml:" role,omitempty" json:"role,omitempty"`
}

type AssetRewards struct {
	AttrRewardsperday string `xml:" rewardsperday,attr"  json:",omitempty"`
	AssetItemreward []*AssetItemreward `xml:" itemreward,omitempty" json:"itemreward,omitempty"`
}

type AssetRole struct {
	AttrItemroleid string `xml:" itemroleid,attr"  json:",omitempty"`
	AttrMax string `xml:" max,attr"  json:",omitempty"`
	AttrMin string `xml:" min,attr"  json:",omitempty"`
}


