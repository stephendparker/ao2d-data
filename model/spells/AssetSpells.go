package spells

type AssetAudioInfo struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetChidleyRoot314159 struct {
	AssetSpells *AssetSpells `xml:" spells,omitempty" json:"spells,omitempty"`
}

type AssetIfIsTargetChanneling struct {
	AttrSpell string `xml:" spell,attr"  json:",omitempty"`
}

type AssetIfOwnLoadPercentIsBelow struct {
	AttrValue string `xml:" value,attr"  json:",omitempty"`
}

type AssetIfSpellActive struct {
	AttrAtleastcharges string `xml:" atleastcharges,attr"  json:",omitempty"`
	AttrCasterchargesonly string `xml:" casterchargesonly,attr"  json:",omitempty"`
	AttrSpell string `xml:" spell,attr"  json:",omitempty"`
	AttrWho string `xml:" who,attr"  json:",omitempty"`
}

type AssetIfSpellCategoryActive struct {
	AttrCategory string `xml:" category,attr"  json:",omitempty"`
	AttrWho string `xml:" who,attr"  json:",omitempty"`
}

type AssetIfTargetEnergyAbovePercentage struct {
	AttrValue string `xml:" value,attr"  json:",omitempty"`
}

type AssetIfTargetHealthAbovePercentage struct {
	AttrValue string `xml:" value,attr"  json:",omitempty"`
}

type AssetIfTargetHealthBelowPercentage struct {
	AttrValue string `xml:" value,attr"  json:",omitempty"`
}

type AssetIfTargetMobCategory struct {
	AttrCategory string `xml:" category,attr"  json:",omitempty"`
}

type AssetIfTargetRangeAbove struct {
	AttrValue string `xml:" value,attr"  json:",omitempty"`
}

type AssetIfTargetType struct {
	AttrType string `xml:" type,attr"  json:",omitempty"`
}

type AssetIfTargetWeaponTypeEquipped struct {
	AttrType string `xml:" type,attr"  json:",omitempty"`
}

type AssetActivespell struct {
	AttrAdjustgroundtargettorange string `xml:" adjustgroundtargettorange,attr"  json:",omitempty"`
	AttrAllowedinknockdown string `xml:" allowedinknockdown,attr"  json:",omitempty"`
	AttrAnimIgnoreHitDelay string `xml:" animIgnoreHitDelay,attr"  json:",omitempty"`
	AttrAnimmode string `xml:" animmode,attr"  json:",omitempty"`
	AttrAssessment string `xml:" assessment,attr"  json:",omitempty"`
	AttrBuffpriority string `xml:" buffpriority,attr"  json:",omitempty"`
	AttrCancelautoattackaftercasting string `xml:" cancelautoattackaftercasting,attr"  json:",omitempty"`
	AttrCastablewhensilenced string `xml:" castablewhensilenced,attr"  json:",omitempty"`
	AttrCastablewhenstunned string `xml:" castablewhenstunned,attr"  json:",omitempty"`
	AttrCastanimation string `xml:" castanimation,attr"  json:",omitempty"`
	AttrCastingtime string `xml:" castingtime,attr"  json:",omitempty"`
	AttrCastrange string `xml:" castrange,attr"  json:",omitempty"`
	AttrCastrangetolerance string `xml:" castrangetolerance,attr"  json:",omitempty"`
	AttrCategory string `xml:" category,attr"  json:",omitempty"`
	AttrDamagecancelscasting string `xml:" damagecancelscasting,attr"  json:",omitempty"`
	AttrDisruptionfactor string `xml:" disruptionfactor,attr"  json:",omitempty"`
	AttrEnergyusage string `xml:" energyusage,attr"  json:",omitempty"`
	AttrHidespelleffecticon string `xml:" hidespelleffecticon,attr"  json:",omitempty"`
	AttrHitdelay string `xml:" hitdelay,attr"  json:",omitempty"`
	AttrIgnorecasttimereductionbonus string `xml:" ignorecasttimereductionbonus,attr"  json:",omitempty"`
	AttrInterruptiblebyspell string `xml:" interruptiblebyspell,attr"  json:",omitempty"`
	AttrKeepmovingaftercast string `xml:" keepmovingaftercast,attr"  json:",omitempty"`
	AttrMaxcharges string `xml:" maxcharges,attr"  json:",omitempty"`
	AttrMaxstack string `xml:" maxstack,attr"  json:",omitempty"`
	AttrNeutraleffect string `xml:" neutraleffect,attr"  json:",omitempty"`
	AttrRecastdelay string `xml:" recastdelay,attr"  json:",omitempty"`
	AttrRequireslineofsight string `xml:" requireslineofsight,attr"  json:",omitempty"`
	AttrShowgizmo string `xml:" showgizmo,attr"  json:",omitempty"`
	AttrShowrange string `xml:" showrange,attr"  json:",omitempty"`
	AttrSpelleffectanimation string `xml:" spelleffectanimation,attr"  json:",omitempty"`
	AttrStackingrule string `xml:" stackingrule,attr"  json:",omitempty"`
	AttrStandtime string `xml:" standtime,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTriggerautoattackcooldown string `xml:" triggerautoattackcooldown,attr"  json:",omitempty"`
	AttrUiatlas string `xml:" uiatlas,attr"  json:",omitempty"`
	AttrUisprite string `xml:" uisprite,attr"  json:",omitempty"`
	AttrUitype string `xml:" uitype,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUnlockedtoequip string `xml:" unlockedtoequip,attr"  json:",omitempty"`
	AttrUsableincombat string `xml:" usableincombat,attr"  json:",omitempty"`
	AssetApplyeffectoneventactive []*AssetApplyeffectoneventactive `xml:" applyeffectoneventactive,omitempty" json:"applyeffectoneventactive,omitempty"`
	AssetApplyspell []*AssetApplyspell `xml:" applyspell,omitempty" json:"applyspell,omitempty"`
	AssetAttributechangeovertime []*AssetAttributechangeovertime `xml:" attributechangeovertime,omitempty" json:"attributechangeovertime,omitempty"`
	AssetAudioInfo *AssetAudioInfo `xml:" AudioInfo,omitempty" json:"AudioInfo,omitempty"`
	AssetBeamvfx []*AssetBeamvfx `xml:" beamvfx,omitempty" json:"beamvfx,omitempty"`
	AssetBuffovertime []*AssetBuffovertime `xml:" buffovertime,omitempty" json:"buffovertime,omitempty"`
	AssetCannotattack *AssetCannotattack `xml:" cannotattack,omitempty" json:"cannotattack,omitempty"`
	AssetCannotattackplayers *AssetCannotattackplayers `xml:" cannotattackplayers,omitempty" json:"cannotattackplayers,omitempty"`
	AssetCastspellaftermoving *AssetCastspellaftermoving `xml:" castspellaftermoving,omitempty" json:"castspellaftermoving,omitempty"`
	AssetChangeaggro *AssetChangeaggro `xml:" changeaggro,omitempty" json:"changeaggro,omitempty"`
	AssetChangepvpstance *AssetChangepvpstance `xml:" changepvpstance,omitempty" json:"changepvpstance,omitempty"`
	AssetChannelingspell *AssetChannelingspell `xml:" channelingspell,omitempty" json:"channelingspell,omitempty"`
	AssetCoupdegrace *AssetCoupdegrace `xml:" coupdegrace,omitempty" json:"coupdegrace,omitempty"`
	AssetDamageshield *AssetDamageshield `xml:" damageshield,omitempty" json:"damageshield,omitempty"`
	AssetDash *AssetDash `xml:" dash,omitempty" json:"dash,omitempty"`
	AssetDirectattributechange []*AssetDirectattributechange `xml:" directattributechange,omitempty" json:"directattributechange,omitempty"`
	AssetDummy []*AssetDummy `xml:" dummy,omitempty" json:"dummy,omitempty"`
	AssetEscapedungeon *AssetEscapedungeon `xml:" escapedungeon,omitempty" json:"escapedungeon,omitempty"`
	AssetFofIndicator *AssetFofIndicator `xml:" fofIndicator,omitempty" json:"fofIndicator,omitempty"`
	AssetForcedmovement *AssetForcedmovement `xml:" forcedmovement,omitempty" json:"forcedmovement,omitempty"`
	AssetGainenergyondamage *AssetGainenergyondamage `xml:" gainenergyondamage,omitempty" json:"gainenergyondamage,omitempty"`
	AssetInterruptcast *AssetInterruptcast `xml:" interruptcast,omitempty" json:"interruptcast,omitempty"`
	AssetInvincibility []*AssetInvincibility `xml:" invincibility,omitempty" json:"invincibility,omitempty"`
	AssetInvisibility *AssetInvisibility `xml:" invisibility,omitempty" json:"invisibility,omitempty"`
	AssetKnockback []*AssetKnockback `xml:" knockback,omitempty" json:"knockback,omitempty"`
	AssetMaketargettopthreat *AssetMaketargettopthreat `xml:" maketargettopthreat,omitempty" json:"maketargettopthreat,omitempty"`
	AssetMatchhighestthreatlevel *AssetMatchhighestthreatlevel `xml:" matchhighestthreatlevel,omitempty" json:"matchhighestthreatlevel,omitempty"`
	AssetMovefreely *AssetMovefreely `xml:" movefreely,omitempty" json:"movefreely,omitempty"`
	AssetNothittable *AssetNothittable `xml:" nothittable,omitempty" json:"nothittable,omitempty"`
	AssetPortalrestriction *AssetPortalrestriction `xml:" portalrestriction,omitempty" json:"portalrestriction,omitempty"`
	AssetProjectile []*AssetProjectile `xml:" projectile,omitempty" json:"projectile,omitempty"`
	AssetPull []*AssetPull `xml:" pull,omitempty" json:"pull,omitempty"`
	AssetPulsingspell []*AssetPulsingspell `xml:" pulsingspell,omitempty" json:"pulsingspell,omitempty"`
	AssetReflectdamageactive *AssetReflectdamageactive `xml:" reflectdamageactive,omitempty" json:"reflectdamageactive,omitempty"`
	AssetRemoveactiveeffects []*AssetRemoveactiveeffects `xml:" removeactiveeffects,omitempty" json:"removeactiveeffects,omitempty"`
	AssetRemoveactivespell []*AssetRemoveactivespell `xml:" removeactivespell,omitempty" json:"removeactivespell,omitempty"`
	AssetRemovefromthreatlist *AssetRemovefromthreatlist `xml:" removefromthreatlist,omitempty" json:"removefromthreatlist,omitempty"`
	AssetRemovespelleffectareasaffectingtarget *AssetRemovespelleffectareasaffectingtarget `xml:" removespelleffectareasaffectingtarget,omitempty" json:"removespelleffectareasaffectingtarget,omitempty"`
	AssetRemovespelleffectareasoftarget []*AssetRemovespelleffectareasoftarget `xml:" removespelleffectareasoftarget,omitempty" json:"removespelleffectareasoftarget,omitempty"`
	AssetResetcooldowns *AssetResetcooldowns `xml:" resetcooldowns,omitempty" json:"resetcooldowns,omitempty"`
	AssetResourcegatheringbuffovertime []*AssetResourcegatheringbuffovertime `xml:" resourcegatheringbuffovertime,omitempty" json:"resourcegatheringbuffovertime,omitempty"`
	AssetResurrect *AssetResurrect `xml:" resurrect,omitempty" json:"resurrect,omitempty"`
	AssetRoot []*AssetRoot `xml:" root,omitempty" json:"root,omitempty"`
	AssetScalefx *AssetScalefx `xml:" scalefx,omitempty" json:"scalefx,omitempty"`
	AssetScreenshakeinfo *AssetScreenshakeinfo `xml:" screenshakeinfo,omitempty" json:"screenshakeinfo,omitempty"`
	AssetSilence []*AssetSilence `xml:" silence,omitempty" json:"silence,omitempty"`
	AssetSpawnmob []*AssetSpawnmob `xml:" spawnmob,omitempty" json:"spawnmob,omitempty"`
	AssetSpelleffectarea []*AssetSpelleffectarea `xml:" spelleffectarea,omitempty" json:"spelleffectarea,omitempty"`
	AssetSpelleffectfx []*AssetSpelleffectfx `xml:" spelleffectfx,omitempty" json:"spelleffectfx,omitempty"`
	AssetSpellimmunity []*AssetSpellimmunity `xml:" spellimmunity,omitempty" json:"spellimmunity,omitempty"`
	AssetSpellindicationarea *AssetSpellindicationarea `xml:" spellindicationarea,omitempty" json:"spellindicationarea,omitempty"`
	AssetSpellvfx []*AssetSpellvfx `xml:" spellvfx,omitempty" json:"spellvfx,omitempty"`
	AssetStun *AssetStun `xml:" stun,omitempty" json:"stun,omitempty"`
	AssetTeleport *AssetTeleport `xml:" teleport,omitempty" json:"teleport,omitempty"`
	AssetTintfx *AssetTintfx `xml:" tintfx,omitempty" json:"tintfx,omitempty"`
}

type AssetAnd struct {
	AssetIfSpellActive []*AssetIfSpellActive `xml:" IfSpellActive,omitempty" json:"IfSpellActive,omitempty"`
	AssetIfTargetHealthAbovePercentage *AssetIfTargetHealthAbovePercentage `xml:" IfTargetHealthAbovePercentage,omitempty" json:"IfTargetHealthAbovePercentage,omitempty"`
	AssetIfTargetHealthBelowPercentage *AssetIfTargetHealthBelowPercentage `xml:" IfTargetHealthBelowPercentage,omitempty" json:"IfTargetHealthBelowPercentage,omitempty"`
	AssetNot *AssetNot `xml:" not,omitempty" json:"not,omitempty"`
}

type AssetApplyeffectoneventactive struct {
	AttrApplychance string `xml:" applychance,attr"  json:",omitempty"`
	AttrApplyeffecttarget string `xml:" applyeffecttarget,attr"  json:",omitempty"`
	AttrCount string `xml:" count,attr"  json:",omitempty"`
	AttrEvent string `xml:" event,attr"  json:",omitempty"`
	AttrSpell string `xml:" spell,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
	AssetIfSpellActive []*AssetIfSpellActive `xml:" IfSpellActive,omitempty" json:"IfSpellActive,omitempty"`
}

type AssetApplyeffectoneventpassive struct {
	AttrChance string `xml:" chance,attr"  json:",omitempty"`
	AttrEvent string `xml:" event,attr"  json:",omitempty"`
	AttrSpell string `xml:" spell,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
}

type AssetApplyspell struct {
	AttrChance string `xml:" chance,attr"  json:",omitempty"`
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrMaxeffectareatargets string `xml:" maxeffectareatargets,attr"  json:",omitempty"`
	AttrShowarearadius string `xml:" showarearadius,attr"  json:",omitempty"`
	AttrSpell string `xml:" spell,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AssetAnd *AssetAnd `xml:" and,omitempty" json:"and,omitempty"`
	AssetIfSpellActive []*AssetIfSpellActive `xml:" IfSpellActive,omitempty" json:"IfSpellActive,omitempty"`
	AssetIfSpellCategoryActive *AssetIfSpellCategoryActive `xml:" IfSpellCategoryActive,omitempty" json:"IfSpellCategoryActive,omitempty"`
	AssetIfTargetEnergyAbovePercentage *AssetIfTargetEnergyAbovePercentage `xml:" IfTargetEnergyAbovePercentage,omitempty" json:"IfTargetEnergyAbovePercentage,omitempty"`
	AssetIfTargetHealthAbovePercentage *AssetIfTargetHealthAbovePercentage `xml:" IfTargetHealthAbovePercentage,omitempty" json:"IfTargetHealthAbovePercentage,omitempty"`
	AssetIfTargetHealthBelowPercentage *AssetIfTargetHealthBelowPercentage `xml:" IfTargetHealthBelowPercentage,omitempty" json:"IfTargetHealthBelowPercentage,omitempty"`
	AssetIfTargetMobCategory *AssetIfTargetMobCategory `xml:" IfTargetMobCategory,omitempty" json:"IfTargetMobCategory,omitempty"`
	AssetIfTargetType *AssetIfTargetType `xml:" IfTargetType,omitempty" json:"IfTargetType,omitempty"`
	AssetNot *AssetNot `xml:" not,omitempty" json:"not,omitempty"`
	AssetOr *AssetOr `xml:" or,omitempty" json:"or,omitempty"`
}

type AssetArc struct {
	AttrAngle string `xml:" angle,attr"  json:",omitempty"`
	AttrInnerradius string `xml:" innerradius,attr"  json:",omitempty"`
	AttrOuterradius string `xml:" outerradius,attr"  json:",omitempty"`
	AssetSpellvfx []*AssetSpellvfx `xml:" spellvfx,omitempty" json:"spellvfx,omitempty"`
}

type AssetArea struct {
	AssetFill *AssetFill `xml:" fill,omitempty" json:"fill,omitempty"`
	AssetShape *AssetShape `xml:" shape,omitempty" json:"shape,omitempty"`
}

type AssetAttributechangeovertime struct {
	AttrAttackingbreakseffect string `xml:" attackingbreakseffect,attr"  json:",omitempty"`
	AttrAttribute string `xml:" attribute,attr"  json:",omitempty"`
	AttrCastingbreakseffect string `xml:" castingbreakseffect,attr"  json:",omitempty"`
	AttrChange string `xml:" change,attr"  json:",omitempty"`
	AttrCount string `xml:" count,attr"  json:",omitempty"`
	AttrDamagebreakseffect string `xml:" damagebreakseffect,attr"  json:",omitempty"`
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrEffecttype string `xml:" effecttype,attr"  json:",omitempty"`
	AttrIgnoreabilitypowerscaling string `xml:" ignoreabilitypowerscaling,attr"  json:",omitempty"`
	AttrIgnorearmor string `xml:" ignorearmor,attr"  json:",omitempty"`
	AttrInitialinterval string `xml:" initialinterval,attr"  json:",omitempty"`
	AttrInterval string `xml:" interval,attr"  json:",omitempty"`
	AttrMaxeffectareatargets string `xml:" maxeffectareatargets,attr"  json:",omitempty"`
	AttrMountingbreakseffect string `xml:" mountingbreakseffect,attr"  json:",omitempty"`
	AttrReflectdamage string `xml:" reflectdamage,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTargetcountvaluebonusfactor string `xml:" targetcountvaluebonusfactor,attr"  json:",omitempty"`
	AttrWorldinteractionbreakseffect string `xml:" worldinteractionbreakseffect,attr"  json:",omitempty"`
	AssetIfSpellActive []*AssetIfSpellActive `xml:" IfSpellActive,omitempty" json:"IfSpellActive,omitempty"`
	AssetIfTargetMobCategory *AssetIfTargetMobCategory `xml:" IfTargetMobCategory,omitempty" json:"IfTargetMobCategory,omitempty"`
	AssetIfTargetType *AssetIfTargetType `xml:" IfTargetType,omitempty" json:"IfTargetType,omitempty"`
}

type AssetBeamvfx struct {
	AttrAlive string `xml:" alive,attr"  json:",omitempty"`
	AttrDuration string `xml:" duration,attr"  json:",omitempty"`
	AttrEvent string `xml:" event,attr"  json:",omitempty"`
	AttrPrefab string `xml:" prefab,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AssetSource *AssetSource `xml:" source,omitempty" json:"source,omitempty"`
	AssetTarget *AssetTarget `xml:" target,omitempty" json:"target,omitempty"`
}

type AssetBuff struct {
	AttrType string `xml:" type,attr"  json:",omitempty"`
	AttrValue string `xml:" value,attr"  json:",omitempty"`
}

type AssetBuffovertime struct {
	AttrAttackingbreakseffect string `xml:" attackingbreakseffect,attr"  json:",omitempty"`
	AttrCastingbreakseffect string `xml:" castingbreakseffect,attr"  json:",omitempty"`
	AttrDamagebreakseffect string `xml:" damagebreakseffect,attr"  json:",omitempty"`
	AttrDiminishingreturnfactor string `xml:" diminishingreturnfactor,attr"  json:",omitempty"`
	AttrDonttriggerdiminishingreturns string `xml:" donttriggerdiminishingreturns,attr"  json:",omitempty"`
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrExpireafterdistance string `xml:" expireafterdistance,attr"  json:",omitempty"`
	AttrIgnoreabilitypowerscaling string `xml:" ignoreabilitypowerscaling,attr"  json:",omitempty"`
	AttrIgnorecrowdcontrolresistance string `xml:" ignorecrowdcontrolresistance,attr"  json:",omitempty"`
	AttrIgnorediminishingreturns string `xml:" ignorediminishingreturns,attr"  json:",omitempty"`
	AttrMaxeffectareatargets string `xml:" maxeffectareatargets,attr"  json:",omitempty"`
	AttrMountingbreakseffect string `xml:" mountingbreakseffect,attr"  json:",omitempty"`
	AttrPersistsafterdeath string `xml:" persistsafterdeath,attr"  json:",omitempty"`
	AttrPersistsafterknockdown string `xml:" persistsafterknockdown,attr"  json:",omitempty"`
	AttrShowarearadius string `xml:" showarearadius,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTargetcountdurationbonusfactor string `xml:" targetcountdurationbonusfactor,attr"  json:",omitempty"`
	AttrTargetcountvaluebonusfactor string `xml:" targetcountvaluebonusfactor,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
	AttrValue string `xml:" value,attr"  json:",omitempty"`
	AttrWorldinteractionbreakseffect string `xml:" worldinteractionbreakseffect,attr"  json:",omitempty"`
	AssetIfSpellActive []*AssetIfSpellActive `xml:" IfSpellActive,omitempty" json:"IfSpellActive,omitempty"`
	AssetIfTargetType *AssetIfTargetType `xml:" IfTargetType,omitempty" json:"IfTargetType,omitempty"`
}

type AssetCannotattack struct {
	AttrAttackingbreakseffect string `xml:" attackingbreakseffect,attr"  json:",omitempty"`
	AttrCastingbreakseffect string `xml:" castingbreakseffect,attr"  json:",omitempty"`
	AttrDonttriggerdiminishingreturns string `xml:" donttriggerdiminishingreturns,attr"  json:",omitempty"`
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrExpireafterdistance string `xml:" expireafterdistance,attr"  json:",omitempty"`
	AttrIgnoreabilitypowerscaling string `xml:" ignoreabilitypowerscaling,attr"  json:",omitempty"`
	AttrIgnorecrowdcontrolresistance string `xml:" ignorecrowdcontrolresistance,attr"  json:",omitempty"`
	AttrIgnorediminishingreturns string `xml:" ignorediminishingreturns,attr"  json:",omitempty"`
	AttrMountingbreakseffect string `xml:" mountingbreakseffect,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
	AttrWorldinteractionbreakseffect string `xml:" worldinteractionbreakseffect,attr"  json:",omitempty"`
	AssetIfTargetType *AssetIfTargetType `xml:" IfTargetType,omitempty" json:"IfTargetType,omitempty"`
}

type AssetCannotattackplayers struct {
	AttrIgnorecrowdcontrolresistance string `xml:" ignorecrowdcontrolresistance,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
}

type AssetCastspellaftermoving struct {
	AttrCooldownafterdamage string `xml:" cooldownafterdamage,attr"  json:",omitempty"`
	AttrDamageresetstimetocast string `xml:" damageresetstimetocast,attr"  json:",omitempty"`
	AttrMinmovespeed string `xml:" minmovespeed,attr"  json:",omitempty"`
	AttrMovetimetocast string `xml:" movetimetocast,attr"  json:",omitempty"`
	AttrSpell string `xml:" spell,attr"  json:",omitempty"`
	AttrStandtimetoendspell string `xml:" standtimetoendspell,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
}

type AssetChangeaggro struct {
	AttrAggroradius string `xml:" aggroradius,attr"  json:",omitempty"`
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrValue string `xml:" value,attr"  json:",omitempty"`
	AssetIfSpellActive []*AssetIfSpellActive `xml:" IfSpellActive,omitempty" json:"IfSpellActive,omitempty"`
}

type AssetChangepvpstance struct {
}

type AssetChannelingspell struct {
	AttrAllowcancel string `xml:" allowcancel,attr"  json:",omitempty"`
	AttrAllowmovement string `xml:" allowmovement,attr"  json:",omitempty"`
	AttrChannelinganimation string `xml:" channelinganimation,attr"  json:",omitempty"`
	AttrDamagecancelscasting string `xml:" damagecancelscasting,attr"  json:",omitempty"`
	AttrDisruptionfactor string `xml:" disruptionfactor,attr"  json:",omitempty"`
	AttrEffectcount string `xml:" effectcount,attr"  json:",omitempty"`
	AttrEffectinterval string `xml:" effectinterval,attr"  json:",omitempty"`
	AttrEnergyusage string `xml:" energyusage,attr"  json:",omitempty"`
	AttrHitdelay string `xml:" hitdelay,attr"  json:",omitempty"`
	AttrInitialeffectinterval string `xml:" initialeffectinterval,attr"  json:",omitempty"`
	AttrUpperbodychannelinganim string `xml:" upperbodychannelinganim,attr"  json:",omitempty"`
	AssetApplyspell []*AssetApplyspell `xml:" applyspell,omitempty" json:"applyspell,omitempty"`
	AssetAudioInfo *AssetAudioInfo `xml:" AudioInfo,omitempty" json:"AudioInfo,omitempty"`
	AssetBeamvfx []*AssetBeamvfx `xml:" beamvfx,omitempty" json:"beamvfx,omitempty"`
	AssetBuffovertime []*AssetBuffovertime `xml:" buffovertime,omitempty" json:"buffovertime,omitempty"`
	AssetChangeaggro *AssetChangeaggro `xml:" changeaggro,omitempty" json:"changeaggro,omitempty"`
	AssetDamageshield *AssetDamageshield `xml:" damageshield,omitempty" json:"damageshield,omitempty"`
	AssetDirectattributechange []*AssetDirectattributechange `xml:" directattributechange,omitempty" json:"directattributechange,omitempty"`
	AssetInvincibility []*AssetInvincibility `xml:" invincibility,omitempty" json:"invincibility,omitempty"`
	AssetInvisibility *AssetInvisibility `xml:" invisibility,omitempty" json:"invisibility,omitempty"`
	AssetKnockback []*AssetKnockback `xml:" knockback,omitempty" json:"knockback,omitempty"`
	AssetProjectile []*AssetProjectile `xml:" projectile,omitempty" json:"projectile,omitempty"`
	AssetPulsingspell []*AssetPulsingspell `xml:" pulsingspell,omitempty" json:"pulsingspell,omitempty"`
	AssetReflectdamageactive *AssetReflectdamageactive `xml:" reflectdamageactive,omitempty" json:"reflectdamageactive,omitempty"`
	AssetRoot []*AssetRoot `xml:" root,omitempty" json:"root,omitempty"`
	AssetScreenshakeinfo *AssetScreenshakeinfo `xml:" screenshakeinfo,omitempty" json:"screenshakeinfo,omitempty"`
	AssetSpawnmob []*AssetSpawnmob `xml:" spawnmob,omitempty" json:"spawnmob,omitempty"`
	AssetSpelleffectarea []*AssetSpelleffectarea `xml:" spelleffectarea,omitempty" json:"spelleffectarea,omitempty"`
	AssetSpelleffectfx []*AssetSpelleffectfx `xml:" spelleffectfx,omitempty" json:"spelleffectfx,omitempty"`
	AssetSpellvfx []*AssetSpellvfx `xml:" spellvfx,omitempty" json:"spellvfx,omitempty"`
	AssetTintfx *AssetTintfx `xml:" tintfx,omitempty" json:"tintfx,omitempty"`
}

type AssetCircle struct {
	AttrRadius string `xml:" radius,attr"  json:",omitempty"`
	AssetFofIndicator *AssetFofIndicator `xml:" fofIndicator,omitempty" json:"fofIndicator,omitempty"`
	AssetSpellvfx []*AssetSpellvfx `xml:" spellvfx,omitempty" json:"spellvfx,omitempty"`
}

type AssetCoupdegrace struct {
}

type AssetDamageshield struct {
	AttrDamageabsorbed string `xml:" damageabsorbed,attr"  json:",omitempty"`
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrEffecttype string `xml:" effecttype,attr"  json:",omitempty"`
	AttrMaxeffectareatargets string `xml:" maxeffectareatargets,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
}

type AssetDash struct {
	AttrAnimID string `xml:" animID,attr"  json:",omitempty"`
	AttrEndeffect string `xml:" endeffect,attr"  json:",omitempty"`
	AttrMovespeed string `xml:" movespeed,attr"  json:",omitempty"`
	AttrRuleset string `xml:" ruleset,attr"  json:",omitempty"`
	AttrScaleAnim string `xml:" scaleAnim,attr"  json:",omitempty"`
	AttrStandtime string `xml:" standtime,attr"  json:",omitempty"`
	AssetIfOwnLoadPercentIsBelow *AssetIfOwnLoadPercentIsBelow `xml:" IfOwnLoadPercentIsBelow,omitempty" json:"IfOwnLoadPercentIsBelow,omitempty"`
}

type AssetDirectattributechange struct {
	AttrAttribute string `xml:" attribute,attr"  json:",omitempty"`
	AttrChange string `xml:" change,attr"  json:",omitempty"`
	AttrChangetype string `xml:" changetype,attr"  json:",omitempty"`
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrEffecttype string `xml:" effecttype,attr"  json:",omitempty"`
	AttrIgnoreabilitypowerscaling string `xml:" ignoreabilitypowerscaling,attr"  json:",omitempty"`
	AttrIgnorearmor string `xml:" ignorearmor,attr"  json:",omitempty"`
	AttrInterruptscasting string `xml:" interruptscasting,attr"  json:",omitempty"`
	AttrMaxeffectareatargets string `xml:" maxeffectareatargets,attr"  json:",omitempty"`
	AttrReflectdamage string `xml:" reflectdamage,attr"  json:",omitempty"`
	AttrShowarearadius string `xml:" showarearadius,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTargetcountvaluebonusfactor string `xml:" targetcountvaluebonusfactor,attr"  json:",omitempty"`
	AssetAnd *AssetAnd `xml:" and,omitempty" json:"and,omitempty"`
	AssetIfSpellActive []*AssetIfSpellActive `xml:" IfSpellActive,omitempty" json:"IfSpellActive,omitempty"`
	AssetIfTargetHealthAbovePercentage *AssetIfTargetHealthAbovePercentage `xml:" IfTargetHealthAbovePercentage,omitempty" json:"IfTargetHealthAbovePercentage,omitempty"`
	AssetIfTargetHealthBelowPercentage *AssetIfTargetHealthBelowPercentage `xml:" IfTargetHealthBelowPercentage,omitempty" json:"IfTargetHealthBelowPercentage,omitempty"`
	AssetIfTargetRangeAbove *AssetIfTargetRangeAbove `xml:" IfTargetRangeAbove,omitempty" json:"IfTargetRangeAbove,omitempty"`
	AssetIfTargetType *AssetIfTargetType `xml:" IfTargetType,omitempty" json:"IfTargetType,omitempty"`
	AssetNot *AssetNot `xml:" not,omitempty" json:"not,omitempty"`
}

type AssetDummy struct {
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrIgnorecrowdcontrolresistance string `xml:" ignorecrowdcontrolresistance,attr"  json:",omitempty"`
	AttrMountingbreakseffect string `xml:" mountingbreakseffect,attr"  json:",omitempty"`
	AttrPersistsafterdeath string `xml:" persistsafterdeath,attr"  json:",omitempty"`
	AttrPersistsafterknockdown string `xml:" persistsafterknockdown,attr"  json:",omitempty"`
	AttrShowarearadius string `xml:" showarearadius,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
	AssetIfSpellActive []*AssetIfSpellActive `xml:" IfSpellActive,omitempty" json:"IfSpellActive,omitempty"`
	AssetNot *AssetNot `xml:" not,omitempty" json:"not,omitempty"`
}

type AssetEscapedungeon struct {
}

type AssetFill struct {
	AttrMaterial string `xml:" material,attr"  json:",omitempty"`
	AssetArc []*AssetArc `xml:" arc,omitempty" json:"arc,omitempty"`
	AssetCircle *AssetCircle `xml:" circle,omitempty" json:"circle,omitempty"`
	AssetNode []*AssetNode `xml:" node,omitempty" json:"node,omitempty"`
	AssetRectangle *AssetRectangle `xml:" rectangle,omitempty" json:"rectangle,omitempty"`
}

type AssetFofIndicator struct {
	AttrAlive string `xml:" alive,attr"  json:",omitempty"`
	AttrConstraintpreset string `xml:" constraintpreset,attr"  json:",omitempty"`
	AttrDuration string `xml:" duration,attr"  json:",omitempty"`
	AttrEvent string `xml:" event,attr"  json:",omitempty"`
	AttrFriendlyPrefab string `xml:" friendlyPrefab,attr"  json:",omitempty"`
	AttrHostilePrefab string `xml:" hostilePrefab,attr"  json:",omitempty"`
	AttrIgnoresocketscale string `xml:" ignoresocketscale,attr"  json:",omitempty"`
	AttrSize string `xml:" size,attr"  json:",omitempty"`
	AttrSocket string `xml:" socket,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTimeoffset string `xml:" timeoffset,attr"  json:",omitempty"`
}

type AssetForcedmovement struct {
	AttrDamagebreakseffect string `xml:" damagebreakseffect,attr"  json:",omitempty"`
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrInitialdirection string `xml:" initialdirection,attr"  json:",omitempty"`
	AttrMaxdirectionchangetime string `xml:" maxdirectionchangetime,attr"  json:",omitempty"`
	AttrMaxwaittime string `xml:" maxwaittime,attr"  json:",omitempty"`
	AttrMindirectionchangetime string `xml:" mindirectionchangetime,attr"  json:",omitempty"`
	AttrMinwaittime string `xml:" minwaittime,attr"  json:",omitempty"`
	AttrMovespeedpreset string `xml:" movespeedpreset,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
}

type AssetGainenergyondamage struct {
	AttrEnergyperhit string `xml:" energyperhit,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
}

type AssetHealonhit struct {
	AttrChance string `xml:" chance,attr"  json:",omitempty"`
	AttrPercentage string `xml:" percentage,attr"  json:",omitempty"`
}

type AssetImpactvfx struct {
	AttrConstraintpreset string `xml:" constraintpreset,attr"  json:",omitempty"`
	AttrIgnoresocketscale string `xml:" ignoresocketscale,attr"  json:",omitempty"`
	AttrImpactsocket string `xml:" impactsocket,attr"  json:",omitempty"`
	AttrPrefab string `xml:" prefab,attr"  json:",omitempty"`
	AttrSize string `xml:" size,attr"  json:",omitempty"`
}

type AssetInterruptcast struct {
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AssetIfIsTargetChanneling *AssetIfIsTargetChanneling `xml:" IfIsTargetChanneling,omitempty" json:"IfIsTargetChanneling,omitempty"`
}

type AssetInvincibility struct {
	AttrAttackingbreakseffect string `xml:" attackingbreakseffect,attr"  json:",omitempty"`
	AttrCastingbreakseffect string `xml:" castingbreakseffect,attr"  json:",omitempty"`
	AttrDamagebreakseffect string `xml:" damagebreakseffect,attr"  json:",omitempty"`
	AttrExpireafterdistance string `xml:" expireafterdistance,attr"  json:",omitempty"`
	AttrIgnoreabilitypowerscaling string `xml:" ignoreabilitypowerscaling,attr"  json:",omitempty"`
	AttrIgnorecrowdcontrolresistance string `xml:" ignorecrowdcontrolresistance,attr"  json:",omitempty"`
	AttrIgnorediminishingreturns string `xml:" ignorediminishingreturns,attr"  json:",omitempty"`
	AttrMountingbreakseffect string `xml:" mountingbreakseffect,attr"  json:",omitempty"`
	AttrPersistsafterdeath string `xml:" persistsafterdeath,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
	AttrWorldinteractionbreakseffect string `xml:" worldinteractionbreakseffect,attr"  json:",omitempty"`
}

type AssetInvisibility struct {
	AttrAttackingbreakseffect string `xml:" attackingbreakseffect,attr"  json:",omitempty"`
	AttrCastingbreakseffect string `xml:" castingbreakseffect,attr"  json:",omitempty"`
	AttrDamagebreakseffect string `xml:" damagebreakseffect,attr"  json:",omitempty"`
	AttrExpireafterdistance string `xml:" expireafterdistance,attr"  json:",omitempty"`
	AttrMountingbreakseffect string `xml:" mountingbreakseffect,attr"  json:",omitempty"`
	AttrPersistsafterdeath string `xml:" persistsafterdeath,attr"  json:",omitempty"`
	AttrPersistsafterknockdown string `xml:" persistsafterknockdown,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
	AttrTo string `xml:" to,attr"  json:",omitempty"`
	AttrWorldinteractionbreakseffect string `xml:" worldinteractionbreakseffect,attr"  json:",omitempty"`
}

type AssetKnockback struct {
	AttrDiminishingreturnfactor string `xml:" diminishingreturnfactor,attr"  json:",omitempty"`
	AttrDistance string `xml:" distance,attr"  json:",omitempty"`
	AttrDonttriggerdiminishingreturns string `xml:" donttriggerdiminishingreturns,attr"  json:",omitempty"`
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrEndeffect string `xml:" endeffect,attr"  json:",omitempty"`
	AttrFromwhere string `xml:" fromwhere,attr"  json:",omitempty"`
	AttrIgnoreabilitypowerscaling string `xml:" ignoreabilitypowerscaling,attr"  json:",omitempty"`
	AttrIgnorecrowdcontrolresistance string `xml:" ignorecrowdcontrolresistance,attr"  json:",omitempty"`
	AttrIgnorediminishingreturns string `xml:" ignorediminishingreturns,attr"  json:",omitempty"`
	AttrMovespeed string `xml:" movespeed,attr"  json:",omitempty"`
	AttrRuleset string `xml:" ruleset,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AssetIfOwnLoadPercentIsBelow *AssetIfOwnLoadPercentIsBelow `xml:" IfOwnLoadPercentIsBelow,omitempty" json:"IfOwnLoadPercentIsBelow,omitempty"`
	AssetIfTargetType *AssetIfTargetType `xml:" IfTargetType,omitempty" json:"IfTargetType,omitempty"`
}

type AssetMaketargettopthreat struct {
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
}

type AssetMatchhighestthreatlevel struct {
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
}

type AssetMovefreely struct {
	AttrPersistsafterdeath string `xml:" persistsafterdeath,attr"  json:",omitempty"`
	AttrPersistsafterknockdown string `xml:" persistsafterknockdown,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
}

type AssetNode struct {
	AttrAngle string `xml:" angle,attr"  json:",omitempty"`
	AttrDistance string `xml:" distance,attr"  json:",omitempty"`
	AssetArc []*AssetArc `xml:" arc,omitempty" json:"arc,omitempty"`
	AssetCircle *AssetCircle `xml:" circle,omitempty" json:"circle,omitempty"`
	AssetNode []*AssetNode `xml:" node,omitempty" json:"node,omitempty"`
	AssetRectangle *AssetRectangle `xml:" rectangle,omitempty" json:"rectangle,omitempty"`
}

type AssetNot struct {
	AssetIfSpellActive []*AssetIfSpellActive `xml:" IfSpellActive,omitempty" json:"IfSpellActive,omitempty"`
	AssetIfTargetHealthBelowPercentage *AssetIfTargetHealthBelowPercentage `xml:" IfTargetHealthBelowPercentage,omitempty" json:"IfTargetHealthBelowPercentage,omitempty"`
	AssetIfTargetWeaponTypeEquipped *AssetIfTargetWeaponTypeEquipped `xml:" IfTargetWeaponTypeEquipped,omitempty" json:"IfTargetWeaponTypeEquipped,omitempty"`
	AssetOr *AssetOr `xml:" or,omitempty" json:"or,omitempty"`
}

type AssetNothittable struct {
	AttrPersistsafterdeath string `xml:" persistsafterdeath,attr"  json:",omitempty"`
	AttrPersistsafterknockdown string `xml:" persistsafterknockdown,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
}

type AssetOr struct {
	AssetIfSpellActive []*AssetIfSpellActive `xml:" IfSpellActive,omitempty" json:"IfSpellActive,omitempty"`
	AssetNot *AssetNot `xml:" not,omitempty" json:"not,omitempty"`
}

type AssetPassivespell struct {
	AttrUiatlas string `xml:" uiatlas,attr"  json:",omitempty"`
	AttrUisprite string `xml:" uisprite,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUnlockedtoequip string `xml:" unlockedtoequip,attr"  json:",omitempty"`
	AssetApplyeffectoneventpassive []*AssetApplyeffectoneventpassive `xml:" applyeffectoneventpassive,omitempty" json:"applyeffectoneventpassive,omitempty"`
	AssetBuff []*AssetBuff `xml:" buff,omitempty" json:"buff,omitempty"`
	AssetHealonhit *AssetHealonhit `xml:" healonhit,omitempty" json:"healonhit,omitempty"`
	AssetPulsingspellpassive *AssetPulsingspellpassive `xml:" pulsingspellpassive,omitempty" json:"pulsingspellpassive,omitempty"`
	AssetReflectdamage []*AssetReflectdamage `xml:" reflectdamage,omitempty" json:"reflectdamage,omitempty"`
	AssetWeightbonus []*AssetWeightbonus `xml:" weightbonus,omitempty" json:"weightbonus,omitempty"`
}

type AssetPortalrestriction struct {
	AttrPersistsafterdeath string `xml:" persistsafterdeath,attr"  json:",omitempty"`
	AttrPersistsafterknockdown string `xml:" persistsafterknockdown,attr"  json:",omitempty"`
	AttrPortaltype string `xml:" portaltype,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
}

type AssetProjectile struct {
	AttrFlyspeed string `xml:" flyspeed,attr"  json:",omitempty"`
	AttrHitsocket string `xml:" hitsocket,attr"  json:",omitempty"`
	AttrPrefab string `xml:" prefab,attr"  json:",omitempty"`
	AttrStartsocket string `xml:" startsocket,attr"  json:",omitempty"`
	AttrTimeoffset string `xml:" timeoffset,attr"  json:",omitempty"`
	AssetAudioInfo *AssetAudioInfo `xml:" AudioInfo,omitempty" json:"AudioInfo,omitempty"`
	AssetImpactvfx []*AssetImpactvfx `xml:" impactvfx,omitempty" json:"impactvfx,omitempty"`
}

type AssetPull struct {
	AttrAnimID string `xml:" animID,attr"  json:",omitempty"`
	AttrDonttriggerdiminishingreturns string `xml:" donttriggerdiminishingreturns,attr"  json:",omitempty"`
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrEnddistance string `xml:" enddistance,attr"  json:",omitempty"`
	AttrEndeffect string `xml:" endeffect,attr"  json:",omitempty"`
	AttrIgnoreabilitypowerscaling string `xml:" ignoreabilitypowerscaling,attr"  json:",omitempty"`
	AttrIgnorecrowdcontrolresistance string `xml:" ignorecrowdcontrolresistance,attr"  json:",omitempty"`
	AttrIgnorediminishingreturns string `xml:" ignorediminishingreturns,attr"  json:",omitempty"`
	AttrMaxeffectareatargets string `xml:" maxeffectareatargets,attr"  json:",omitempty"`
	AttrMovespeed string `xml:" movespeed,attr"  json:",omitempty"`
	AttrPullwhere string `xml:" pullwhere,attr"  json:",omitempty"`
	AttrShowarearadius string `xml:" showarearadius,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AssetIfSpellActive []*AssetIfSpellActive `xml:" IfSpellActive,omitempty" json:"IfSpellActive,omitempty"`
	AssetIfTargetType *AssetIfTargetType `xml:" IfTargetType,omitempty" json:"IfTargetType,omitempty"`
	AssetNot *AssetNot `xml:" not,omitempty" json:"not,omitempty"`
}

type AssetPulsingspell struct {
	AttrCount string `xml:" count,attr"  json:",omitempty"`
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrInitialinterval string `xml:" initialinterval,attr"  json:",omitempty"`
	AttrInterval string `xml:" interval,attr"  json:",omitempty"`
	AttrMaxeffectareatargets string `xml:" maxeffectareatargets,attr"  json:",omitempty"`
	AttrMountingbreakseffect string `xml:" mountingbreakseffect,attr"  json:",omitempty"`
	AttrPersistsafterdeath string `xml:" persistsafterdeath,attr"  json:",omitempty"`
	AttrSpell string `xml:" spell,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTargetcountvaluebonusfactor string `xml:" targetcountvaluebonusfactor,attr"  json:",omitempty"`
	AssetIfSpellActive []*AssetIfSpellActive `xml:" IfSpellActive,omitempty" json:"IfSpellActive,omitempty"`
}

type AssetPulsingspellpassive struct {
	AttrBindtoitem string `xml:" bindtoitem,attr"  json:",omitempty"`
	AttrCount string `xml:" count,attr"  json:",omitempty"`
	AttrInitialinterval string `xml:" initialinterval,attr"  json:",omitempty"`
	AttrInterval string `xml:" interval,attr"  json:",omitempty"`
	AttrSpell string `xml:" spell,attr"  json:",omitempty"`
}

type AssetRectangle struct {
	AttrHeight string `xml:" height,attr"  json:",omitempty"`
	AttrWidth string `xml:" width,attr"  json:",omitempty"`
	AssetBeamvfx []*AssetBeamvfx `xml:" beamvfx,omitempty" json:"beamvfx,omitempty"`
	AssetSpellvfx []*AssetSpellvfx `xml:" spellvfx,omitempty" json:"spellvfx,omitempty"`
}

type AssetReflectdamage struct {
	AttrAmountpercent string `xml:" amountpercent,attr"  json:",omitempty"`
	AttrAstype string `xml:" astype,attr"  json:",omitempty"`
	AttrDamagetype string `xml:" damagetype,attr"  json:",omitempty"`
}

type AssetReflectdamageactive struct {
	AttrAmountpercent string `xml:" amountpercent,attr"  json:",omitempty"`
	AttrAstype string `xml:" astype,attr"  json:",omitempty"`
	AttrDamagetype string `xml:" damagetype,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
}

type AssetRemoveactiveeffects struct {
	AttrCategory string `xml:" category,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
}

type AssetRemoveactivespell struct {
	AttrCasterchargesonly string `xml:" casterchargesonly,attr"  json:",omitempty"`
	AttrCharges string `xml:" charges,attr"  json:",omitempty"`
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrSpell string `xml:" spell,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AssetIfSpellActive []*AssetIfSpellActive `xml:" IfSpellActive,omitempty" json:"IfSpellActive,omitempty"`
	AssetIfTargetHealthBelowPercentage *AssetIfTargetHealthBelowPercentage `xml:" IfTargetHealthBelowPercentage,omitempty" json:"IfTargetHealthBelowPercentage,omitempty"`
	AssetNot *AssetNot `xml:" not,omitempty" json:"not,omitempty"`
}

type AssetRemovefromthreatlist struct {
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
}

type AssetRemovespelleffectareasaffectingtarget struct {
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrOwnedby string `xml:" ownedby,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
}

type AssetRemovespelleffectareasoftarget struct {
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AssetNot *AssetNot `xml:" not,omitempty" json:"not,omitempty"`
}

type AssetResetcooldowns struct {
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
}

type AssetResourcegatheringbuffovertime struct {
	AttrBufftype string `xml:" bufftype,attr"  json:",omitempty"`
	AttrResourcetype string `xml:" resourcetype,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
	AttrValue string `xml:" value,attr"  json:",omitempty"`
}

type AssetResurrect struct {
	AttrRestoration string `xml:" restoration,attr"  json:",omitempty"`
}

type AssetRoot struct {
	AttrAttackingbreakseffect string `xml:" attackingbreakseffect,attr"  json:",omitempty"`
	AttrDamagebreakseffect string `xml:" damagebreakseffect,attr"  json:",omitempty"`
	AttrDonttriggerdiminishingreturns string `xml:" donttriggerdiminishingreturns,attr"  json:",omitempty"`
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrIgnoreabilitypowerscaling string `xml:" ignoreabilitypowerscaling,attr"  json:",omitempty"`
	AttrIgnorecrowdcontrolresistance string `xml:" ignorecrowdcontrolresistance,attr"  json:",omitempty"`
	AttrIgnorediminishingreturns string `xml:" ignorediminishingreturns,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
	AssetIfSpellActive []*AssetIfSpellActive `xml:" IfSpellActive,omitempty" json:"IfSpellActive,omitempty"`
}

type AssetScalefx struct {
	AttrScalevalue string `xml:" scalevalue,attr"  json:",omitempty"`
}

type AssetScreenshakeinfo struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetShape struct {
	AttrMaterial string `xml:" material,attr"  json:",omitempty"`
	AssetArc []*AssetArc `xml:" arc,omitempty" json:"arc,omitempty"`
	AssetCircle *AssetCircle `xml:" circle,omitempty" json:"circle,omitempty"`
	AssetNode []*AssetNode `xml:" node,omitempty" json:"node,omitempty"`
}

type AssetSilence struct {
	AttrAttackingbreakseffect string `xml:" attackingbreakseffect,attr"  json:",omitempty"`
	AttrCastingbreakseffect string `xml:" castingbreakseffect,attr"  json:",omitempty"`
	AttrDonttriggerdiminishingreturns string `xml:" donttriggerdiminishingreturns,attr"  json:",omitempty"`
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrExpireafterdistance string `xml:" expireafterdistance,attr"  json:",omitempty"`
	AttrIgnoreabilitypowerscaling string `xml:" ignoreabilitypowerscaling,attr"  json:",omitempty"`
	AttrIgnorecrowdcontrolresistance string `xml:" ignorecrowdcontrolresistance,attr"  json:",omitempty"`
	AttrIgnorediminishingreturns string `xml:" ignorediminishingreturns,attr"  json:",omitempty"`
	AttrMountingbreakseffect string `xml:" mountingbreakseffect,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTargetcountdurationbonusfactor string `xml:" targetcountdurationbonusfactor,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
	AttrWorldinteractionbreakseffect string `xml:" worldinteractionbreakseffect,attr"  json:",omitempty"`
}

type AssetSource struct {
	AttrConstraintpreset string `xml:" constraintpreset,attr"  json:",omitempty"`
	AttrSocket string `xml:" socket,attr"  json:",omitempty"`
}

type AssetSpawnmob struct {
	AttrAnglefromcaster string `xml:" anglefromcaster,attr"  json:",omitempty"`
	AttrDistancefromtarget string `xml:" distancefromtarget,attr"  json:",omitempty"`
	AttrLimit string `xml:" limit,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTtl string `xml:" ttl,attr"  json:",omitempty"`
	AttrTtlaftercasterdeath string `xml:" ttlaftercasterdeath,attr"  json:",omitempty"`
	AttrTtlafterreset string `xml:" ttlafterreset,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
	AssetNot *AssetNot `xml:" not,omitempty" json:"not,omitempty"`
}

type AssetSpelleffectarea struct {
	AttrApplyeffectonlyoncepertarget string `xml:" applyeffectonlyoncepertarget,attr"  json:",omitempty"`
	AttrDeletewhenmaxtargetshit string `xml:" deletewhenmaxtargetshit,attr"  json:",omitempty"`
	AttrEffect string `xml:" effect,attr"  json:",omitempty"`
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrInterruptscasting string `xml:" interruptscasting,attr"  json:",omitempty"`
	AttrMaxeffectareatargets string `xml:" maxeffectareatargets,attr"  json:",omitempty"`
	AttrMaxheightdelta string `xml:" maxheightdelta,attr"  json:",omitempty"`
	AttrPosition string `xml:" position,attr"  json:",omitempty"`
	AttrRemoveoncasterdeath string `xml:" removeoncasterdeath,attr"  json:",omitempty"`
	AttrRemoveonlandscapecollision string `xml:" removeonlandscapecollision,attr"  json:",omitempty"`
	AttrRemoveonleave string `xml:" removeonleave,attr"  json:",omitempty"`
	AttrRemovewhencastergone string `xml:" removewhencastergone,attr"  json:",omitempty"`
	AttrRotation string `xml:" rotation,attr"  json:",omitempty"`
	AttrShowarearadius string `xml:" showarearadius,attr"  json:",omitempty"`
	AttrSkillshotindicatormax string `xml:" skillshotindicatormax,attr"  json:",omitempty"`
	AttrSkillshotindicatormin string `xml:" skillshotindicatormin,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
	AttrVisible string `xml:" visible,attr"  json:",omitempty"`
	AssetArea *AssetArea `xml:" area,omitempty" json:"area,omitempty"`
}

type AssetSpelleffectfx struct {
	AttrConstraintpreset string `xml:" constraintpreset,attr"  json:",omitempty"`
	AttrDuration string `xml:" duration,attr"  json:",omitempty"`
	AttrFadeouttime string `xml:" fadeouttime,attr"  json:",omitempty"`
	AttrIgnoresocketscale string `xml:" ignoresocketscale,attr"  json:",omitempty"`
	AttrPrefab string `xml:" prefab,attr"  json:",omitempty"`
	AttrSize string `xml:" size,attr"  json:",omitempty"`
	AttrSocket string `xml:" socket,attr"  json:",omitempty"`
	AttrTimeoffset string `xml:" timeoffset,attr"  json:",omitempty"`
}

type AssetSpellimmunity struct {
	AttrAttackingbreakseffect string `xml:" attackingbreakseffect,attr"  json:",omitempty"`
	AttrCastingbreakseffect string `xml:" castingbreakseffect,attr"  json:",omitempty"`
	AttrCategory string `xml:" category,attr"  json:",omitempty"`
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrExpireafterdistance string `xml:" expireafterdistance,attr"  json:",omitempty"`
	AttrIgnoreabilitypowerscaling string `xml:" ignoreabilitypowerscaling,attr"  json:",omitempty"`
	AttrIgnorecrowdcontrolresistance string `xml:" ignorecrowdcontrolresistance,attr"  json:",omitempty"`
	AttrIgnorediminishingreturns string `xml:" ignorediminishingreturns,attr"  json:",omitempty"`
	AttrMaxeffectareatargets string `xml:" maxeffectareatargets,attr"  json:",omitempty"`
	AttrMountingbreakseffect string `xml:" mountingbreakseffect,attr"  json:",omitempty"`
	AttrPersistsafterknockdown string `xml:" persistsafterknockdown,attr"  json:",omitempty"`
	AttrSpellname string `xml:" spellname,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
	AttrWorldinteractionbreakseffect string `xml:" worldinteractionbreakseffect,attr"  json:",omitempty"`
	AssetIfSpellActive []*AssetIfSpellActive `xml:" IfSpellActive,omitempty" json:"IfSpellActive,omitempty"`
	AssetIfTargetType *AssetIfTargetType `xml:" IfTargetType,omitempty" json:"IfTargetType,omitempty"`
}

type AssetSpellindicationarea struct {
	AttrPosition string `xml:" position,attr"  json:",omitempty"`
	AttrRotation string `xml:" rotation,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrVisible string `xml:" visible,attr"  json:",omitempty"`
	AssetArea *AssetArea `xml:" area,omitempty" json:"area,omitempty"`
}

type AssetSpells struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetActivespell []*AssetActivespell `xml:" activespell,omitempty" json:"activespell,omitempty"`
	AssetPassivespell []*AssetPassivespell `xml:" passivespell,omitempty" json:"passivespell,omitempty"`
}

type AssetSpellvfx struct {
	AttrAlive string `xml:" alive,attr"  json:",omitempty"`
	AttrConstraintpreset string `xml:" constraintpreset,attr"  json:",omitempty"`
	AttrDuration string `xml:" duration,attr"  json:",omitempty"`
	AttrEvent string `xml:" event,attr"  json:",omitempty"`
	AttrFadeouttime string `xml:" fadeouttime,attr"  json:",omitempty"`
	AttrIgnoresocketscale string `xml:" ignoresocketscale,attr"  json:",omitempty"`
	AttrPrefab string `xml:" prefab,attr"  json:",omitempty"`
	AttrSize string `xml:" size,attr"  json:",omitempty"`
	AttrSocket string `xml:" socket,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTimeoffset string `xml:" timeoffset,attr"  json:",omitempty"`
}

type AssetStun struct {
	AttrAttackingbreakseffect string `xml:" attackingbreakseffect,attr"  json:",omitempty"`
	AttrDamagebreakseffect string `xml:" damagebreakseffect,attr"  json:",omitempty"`
	AttrDiminishingreturnfactor string `xml:" diminishingreturnfactor,attr"  json:",omitempty"`
	AttrEffectarearadius string `xml:" effectarearadius,attr"  json:",omitempty"`
	AttrIgnoreabilitypowerscaling string `xml:" ignoreabilitypowerscaling,attr"  json:",omitempty"`
	AttrIgnorecrowdcontrolresistance string `xml:" ignorecrowdcontrolresistance,attr"  json:",omitempty"`
	AttrIgnorediminishingreturns string `xml:" ignorediminishingreturns,attr"  json:",omitempty"`
	AttrInterruptscasting string `xml:" interruptscasting,attr"  json:",omitempty"`
	AttrTarget string `xml:" target,attr"  json:",omitempty"`
	AttrTargetcountdurationbonusfactor string `xml:" targetcountdurationbonusfactor,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
}

type AssetTarget struct {
	AttrConstraintpreset string `xml:" constraintpreset,attr"  json:",omitempty"`
	AttrSocket string `xml:" socket,attr"  json:",omitempty"`
}

type AssetTeleport struct {
	AttrEndeffect string `xml:" endeffect,attr"  json:",omitempty"`
	AssetIfOwnLoadPercentIsBelow *AssetIfOwnLoadPercentIsBelow `xml:" IfOwnLoadPercentIsBelow,omitempty" json:"IfOwnLoadPercentIsBelow,omitempty"`
}

type AssetTintfx struct {
	AttrColor string `xml:" color,attr"  json:",omitempty"`
}

type AssetWeightbonus struct {
	AttrItem string `xml:" item,attr"  json:",omitempty"`
	AttrValue string `xml:" value,attr"  json:",omitempty"`
}


