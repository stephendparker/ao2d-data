package mobs

type AssetAssetVfxPreset struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetAudioInfo struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetChidleyRoot314159 struct {
	AssetMobs *AssetMobs `xml:" Mobs,omitempty" json:"Mobs,omitempty"`
}

type AssetDeathSpells struct {
	AssetSpell []*AssetSpell `xml:" spell,omitempty" json:"spell,omitempty"`
}

type AssetFootStepVfxPreset struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetHarvestable struct {
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
}

type AssetHellgate struct {
	AttrType string `xml:" type,attr"  json:",omitempty"`
}

type AssetLoot struct {
	AssetHarvestable *AssetHarvestable `xml:" Harvestable,omitempty" json:"Harvestable,omitempty"`
	AssetHellgate *AssetHellgate `xml:" Hellgate,omitempty" json:"Hellgate,omitempty"`
	AssetLootListReference []*AssetLootListReference `xml:" LootListReference,omitempty" json:"LootListReference,omitempty"`
}

type AssetLootListReference struct {
	AttrChance string `xml:" chance,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetMob struct {
	AttrAbilitypower string `xml:" abilitypower,attr"  json:",omitempty"`
	AttrAggrodelayafterspawn string `xml:" aggrodelayafterspawn,attr"  json:",omitempty"`
	AttrAggroradius string `xml:" aggroradius,attr"  json:",omitempty"`
	AttrAlertradius string `xml:" alertradius,attr"  json:",omitempty"`
	AttrAlertsfactions string `xml:" alertsfactions,attr"  json:",omitempty"`
	AttrAttackcollisionradius string `xml:" attackcollisionradius,attr"  json:",omitempty"`
	AttrAttackdamage string `xml:" attackdamage,attr"  json:",omitempty"`
	AttrAttackmovespeed string `xml:" attackmovespeed,attr"  json:",omitempty"`
	AttrAttackrange string `xml:" attackrange,attr"  json:",omitempty"`
	AttrAttackspeed string `xml:" attackspeed,attr"  json:",omitempty"`
	AttrAttacktype string `xml:" attacktype,attr"  json:",omitempty"`
	AttrAvatar string `xml:" avatar,attr"  json:",omitempty"`
	AttrAvatarring string `xml:" avatarring,attr"  json:",omitempty"`
	AttrCategory string `xml:" category,attr"  json:",omitempty"`
	AttrChargesperchargeup string `xml:" chargesperchargeup,attr"  json:",omitempty"`
	AttrCollisionradius string `xml:" collisionradius,attr"  json:",omitempty"`
	AttrCrowdcontrolresistance string `xml:" crowdcontrolresistance,attr"  json:",omitempty"`
	AttrDamageaggrofactor string `xml:" damageaggrofactor,attr"  json:",omitempty"`
	AttrDangerstate string `xml:" dangerstate,attr"  json:",omitempty"`
	AttrEnergymax string `xml:" energymax,attr"  json:",omitempty"`
	AttrEnergyregeneration string `xml:" energyregeneration,attr"  json:",omitempty"`
	AttrExecutionspell string `xml:" executionspell,attr"  json:",omitempty"`
	AttrFaction string `xml:" faction,attr"  json:",omitempty"`
	AttrFame string `xml:" fame,attr"  json:",omitempty"`
	AttrFixedrotation string `xml:" fixedrotation,attr"  json:",omitempty"`
	AttrFocusfireattackersthreshold string `xml:" focusfireattackersthreshold,attr"  json:",omitempty"`
	AttrHealingaggrofactor string `xml:" healingaggrofactor,attr"  json:",omitempty"`
	AttrHitpointsmax string `xml:" hitpointsmax,attr"  json:",omitempty"`
	AttrHitpointsregeneration string `xml:" hitpointsregeneration,attr"  json:",omitempty"`
	AttrImmunetoforcedmovement string `xml:" immunetoforcedmovement,attr"  json:",omitempty"`
	AttrIsabletoexecute string `xml:" isabletoexecute,attr"  json:",omitempty"`
	AttrIsboss string `xml:" isboss,attr"  json:",omitempty"`
	AttrMagicattackdamagetime string `xml:" magicattackdamagetime,attr"  json:",omitempty"`
	AttrMagicresistance string `xml:" magicresistance,attr"  json:",omitempty"`
	AttrMaxcharges string `xml:" maxcharges,attr"  json:",omitempty"`
	AttrMeleeattackdamagetime string `xml:" meleeattackdamagetime,attr"  json:",omitempty"`
	AttrMovespeed string `xml:" movespeed,attr"  json:",omitempty"`
	AttrNpchostility string `xml:" npchostility,attr"  json:",omitempty"`
	AttrOnlysimpleroaming string `xml:" onlysimpleroaming,attr"  json:",omitempty"`
	AttrPhysicalarmor string `xml:" physicalarmor,attr"  json:",omitempty"`
	AttrPrefab string `xml:" prefab,attr"  json:",omitempty"`
	AttrPrefabscale string `xml:" prefabscale,attr"  json:",omitempty"`
	AttrPursuitradius string `xml:" pursuitradius,attr"  json:",omitempty"`
	AttrRangedattackdamagetime string `xml:" rangedattackdamagetime,attr"  json:",omitempty"`
	AttrRespawntimesecondsmax string `xml:" respawntimesecondsmax,attr"  json:",omitempty"`
	AttrRespawntimesecondsmin string `xml:" respawntimesecondsmin,attr"  json:",omitempty"`
	AttrRoamingidletimemax string `xml:" roamingidletimemax,attr"  json:",omitempty"`
	AttrRoamingidletimemin string `xml:" roamingidletimemin,attr"  json:",omitempty"`
	AttrRoamingradius string `xml:" roamingradius,attr"  json:",omitempty"`
	AttrShieldaggrofactor string `xml:" shieldaggrofactor,attr"  json:",omitempty"`
	AttrStartcharges string `xml:" startcharges,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrTimeperchargeseconds string `xml:" timeperchargeseconds,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AssetAssetVfxPreset *AssetAssetVfxPreset `xml:" AssetVfxPreset,omitempty" json:"AssetVfxPreset,omitempty"`
	AssetAudioInfo *AssetAudioInfo `xml:" AudioInfo,omitempty" json:"AudioInfo,omitempty"`
	AssetDeathSpells *AssetDeathSpells `xml:" DeathSpells,omitempty" json:"DeathSpells,omitempty"`
	AssetFootStepVfxPreset *AssetFootStepVfxPreset `xml:" FootStepVfxPreset,omitempty" json:"FootStepVfxPreset,omitempty"`
	AssetLoot *AssetLoot `xml:" Loot,omitempty" json:"Loot,omitempty"`
	AssetProjectile []*AssetProjectile `xml:" projectile,omitempty" json:"projectile,omitempty"`
	AssetSocketPreset *AssetSocketPreset `xml:" SocketPreset,omitempty" json:"SocketPreset,omitempty"`
	AssetSpells *AssetSpells `xml:" Spells,omitempty" json:"Spells,omitempty"`
}

type AssetMobs struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetMob []*AssetMob `xml:" Mob,omitempty" json:"Mob,omitempty"`
}

type AssetSocketPreset struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetSpells struct {
	AssetSpell []*AssetSpell `xml:" spell,omitempty" json:"spell,omitempty"`
}

type AssetAnd struct {
	AssetCondition []*AssetCondition `xml:" condition,omitempty" json:"condition,omitempty"`
}

type AssetCondition struct {
	AttrType string `xml:" type,attr"  json:",omitempty"`
	AttrValue string `xml:" value,attr"  json:",omitempty"`
}

type AssetImpactvfx struct {
	AttrConstraintpreset string `xml:" constraintpreset,attr"  json:",omitempty"`
	AttrIgnoresocketscale string `xml:" ignoresocketscale,attr"  json:",omitempty"`
	AttrImpactsocket string `xml:" impactsocket,attr"  json:",omitempty"`
	AttrPrefab string `xml:" prefab,attr"  json:",omitempty"`
}

type AssetNot struct {
	AssetCondition []*AssetCondition `xml:" condition,omitempty" json:"condition,omitempty"`
	AssetOr *AssetOr `xml:" or,omitempty" json:"or,omitempty"`
}

type AssetOr struct {
	AssetAnd []*AssetAnd `xml:" and,omitempty" json:"and,omitempty"`
	AssetCondition []*AssetCondition `xml:" condition,omitempty" json:"condition,omitempty"`
	AssetNot []*AssetNot `xml:" not,omitempty" json:"not,omitempty"`
}

type AssetProjectile struct {
	AttrFlyspeed string `xml:" flyspeed,attr"  json:",omitempty"`
	AttrHitsocket string `xml:" hitsocket,attr"  json:",omitempty"`
	AttrPrefab string `xml:" prefab,attr"  json:",omitempty"`
	AttrStartsocket string `xml:" startsocket,attr"  json:",omitempty"`
	AttrTimeoffset string `xml:" timeoffset,attr"  json:",omitempty"`
	AssetImpactvfx *AssetImpactvfx `xml:" impactvfx,omitempty" json:"impactvfx,omitempty"`
}

type AssetSpell struct {
	AttrGroundtargetoffset string `xml:" groundtargetoffset,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrSaytext string `xml:" saytext,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AssetCondition []*AssetCondition `xml:" condition,omitempty" json:"condition,omitempty"`
	AssetNot []*AssetNot `xml:" not,omitempty" json:"not,omitempty"`
	AssetOr *AssetOr `xml:" or,omitempty" json:"or,omitempty"`
}


