package territorytypes

type AssetChidleyRoot314159 struct {
	AssetTerritorytypes *AssetTerritorytypes `xml:" territorytypes,omitempty" json:"territorytypes,omitempty"`
}

type AssetFriendlyFX struct {
	AttrBeam string `xml:" beam,attr"  json:",omitempty"`
	AttrImpact string `xml:" impact,attr"  json:",omitempty"`
}

type AssetGuardSettings struct {
	AttrRespawntimeseconds string `xml:" respawntimeseconds,attr"  json:",omitempty"`
}

type AssetHostileFX struct {
	AttrBeam string `xml:" beam,attr"  json:",omitempty"`
	AttrImpact string `xml:" impact,attr"  json:",omitempty"`
}

type AssetMatchSettings struct {
	AttrConquercost string `xml:" conquercost,attr"  json:",omitempty"`
	AttrMatchtype string `xml:" matchtype,attr"  json:",omitempty"`
	AttrMindefenderticketsleftforbreakingplayerlock string `xml:" mindefenderticketsleftforbreakingplayerlock,attr"  json:",omitempty"`
	AttrMindefenderticketsleftforcounterattack string `xml:" mindefenderticketsleftforcounterattack,attr"  json:",omitempty"`
	AttrMintimeuntilbattleinminutes string `xml:" mintimeuntilbattleinminutes,attr"  json:",omitempty"`
	AttrTimeslotspaninminutes string `xml:" timeslotspaninminutes,attr"  json:",omitempty"`
	AttrTimeslotspansblocked string `xml:" timeslotspansblocked,attr"  json:",omitempty"`
	AttrTimeuntilretaliationmustbedeclaredinmin string `xml:" timeuntilretaliationmustbedeclaredinmin,attr"  json:",omitempty"`
	AttrWinningpayoutfactor string `xml:" winningpayoutfactor,attr"  json:",omitempty"`
	AssetTimeSlots *AssetTimeSlots `xml:" TimeSlots,omitempty" json:"TimeSlots,omitempty"`
}

type AssetTimeSlot struct {
	AttrId string `xml:" id,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
}

type AssetTimeSlots struct {
	AttrBlockedmatchtimestampend string `xml:" blockedmatchtimestampend,attr"  json:",omitempty"`
	AttrBlockedmatchtimestampstart string `xml:" blockedmatchtimestampstart,attr"  json:",omitempty"`
	AssetTimeSlot []*AssetTimeSlot `xml:" TimeSlot,omitempty" json:"TimeSlot,omitempty"`
}

type AssetTower struct {
	AttrAttackdamage string `xml:" attackdamage,attr"  json:",omitempty"`
	AttrAttackdistancedelay string `xml:" attackdistancedelay,attr"  json:",omitempty"`
	AttrAttackinterval string `xml:" attackinterval,attr"  json:",omitempty"`
	AttrDisengageattackrange string `xml:" disengageattackrange,attr"  json:",omitempty"`
	AttrEngageattackrange string `xml:" engageattackrange,attr"  json:",omitempty"`
	AssetFriendlyFX *AssetFriendlyFX `xml:" FriendlyFX,omitempty" json:"FriendlyFX,omitempty"`
	AssetHostileFX *AssetHostileFX `xml:" HostileFX,omitempty" json:"HostileFX,omitempty"`
}

type AssetTerritorytype struct {
	AttrBattlevaultsize string `xml:" battlevaultsize,attr"  json:",omitempty"`
	AttrClaimcost string `xml:" claimcost,attr"  json:",omitempty"`
	AttrClaimnutrition string `xml:" claimnutrition,attr"  json:",omitempty"`
	AttrConquernutrition string `xml:" conquernutrition,attr"  json:",omitempty"`
	AttrDefensebonuschangeafterloosingabattle string `xml:" defensebonuschangeafterloosingabattle,attr"  json:",omitempty"`
	AttrDefensebonuschangeafterprimetime string `xml:" defensebonuschangeafterprimetime,attr"  json:",omitempty"`
	AttrDefensepointschangeafterprimetime string `xml:" defensepointschangeafterprimetime,attr"  json:",omitempty"`
	AttrMaxdefensebonus string `xml:" maxdefensebonus,attr"  json:",omitempty"`
	AttrMaxdefensepoints string `xml:" maxdefensepoints,attr"  json:",omitempty"`
	AttrMindefensebonus string `xml:" mindefensebonus,attr"  json:",omitempty"`
	AttrMintimebetweenmatchesperconnectioninmin string `xml:" mintimebetweenmatchesperconnectioninmin,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrNutritionstorage string `xml:" nutritionstorage,attr"  json:",omitempty"`
	AttrSecondspernutrition string `xml:" secondspernutrition,attr"  json:",omitempty"`
	AttrUpkeepcost string `xml:" upkeepcost,attr"  json:",omitempty"`
	AttrVaultsize string `xml:" vaultsize,attr"  json:",omitempty"`
	AttrWeightlimit string `xml:" weightlimit,attr"  json:",omitempty"`
	AssetGuardSettings *AssetGuardSettings `xml:" GuardSettings,omitempty" json:"GuardSettings,omitempty"`
	AssetMatchSettings *AssetMatchSettings `xml:" MatchSettings,omitempty" json:"MatchSettings,omitempty"`
	AssetTower *AssetTower `xml:" Tower,omitempty" json:"Tower,omitempty"`
}

type AssetTerritorytypes struct {
	AttrUpkeeptimerinmin string `xml:" upkeeptimerinmin,attr"  json:",omitempty"`
	AttrUpkeepwarningchecktimespaninhours string `xml:" upkeepwarningchecktimespaninhours,attr"  json:",omitempty"`
	AttrUpkeepwarningintervalinmin string `xml:" upkeepwarningintervalinmin,attr"  json:",omitempty"`
	AssetTerritorytype []*AssetTerritorytype `xml:" territorytype,omitempty" json:"territorytype,omitempty"`
}


