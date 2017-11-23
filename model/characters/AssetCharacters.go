package characters

type AssetAudioInfo struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetAvatarToCustomizationPresets struct {
	AssetPreset []*AssetPreset `xml:" Preset,omitempty" json:"Preset,omitempty"`
}

type AssetBeardState struct {
	AttrMesh string `xml:" mesh,attr"  json:",omitempty"`
	AttrState string `xml:" state,attr"  json:",omitempty"`
}

type AssetCharacter struct {
	AttrEquipmentsuffix string `xml:" equipmentsuffix,attr"  json:",omitempty"`
	AttrNamelocatag string `xml:" namelocatag,attr"  json:",omitempty"`
	AttrPrefab string `xml:" prefab,attr"  json:",omitempty"`
	AttrPrefabscale string `xml:" prefabscale,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AssetAudioInfo *AssetAudioInfo `xml:" AudioInfo,omitempty" json:"AudioInfo,omitempty"`
	AssetFootStepVfxPreset *AssetFootStepVfxPreset `xml:" FootStepVfxPreset,omitempty" json:"FootStepVfxPreset,omitempty"`
	AssetSocketPreset *AssetSocketPreset `xml:" SocketPreset,omitempty" json:"SocketPreset,omitempty"`
	AssetVfxPreset *AssetVfxPreset `xml:" VfxPreset,omitempty" json:"VfxPreset,omitempty"`
}

type AssetCharacterData struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetAvatarToCustomizationPresets *AssetAvatarToCustomizationPresets `xml:" AvatarToCustomizationPresets,omitempty" json:"AvatarToCustomizationPresets,omitempty"`
	AssetCharacters *AssetCharacters `xml:" Characters,omitempty" json:"Characters,omitempty"`
	AssetCustomizations *AssetCustomizations `xml:" Customizations,omitempty" json:"Customizations,omitempty"`
	AssetRaces *AssetRaces `xml:" Races,omitempty" json:"Races,omitempty"`
}

type AssetCharacters struct {
	AttrGlobalcastdelay string `xml:" globalcastdelay,attr"  json:",omitempty"`
	AssetCharacter []*AssetCharacter `xml:" Character,omitempty" json:"Character,omitempty"`
	AssetDefaultValues *AssetDefaultValues `xml:" DefaultValues,omitempty" json:"DefaultValues,omitempty"`
}

type AssetChidleyRoot314159 struct {
	AssetCharacterData *AssetCharacterData `xml:" CharacterData,omitempty" json:"CharacterData,omitempty"`
}

type AssetCustomization struct {
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AssetItem []*AssetItem `xml:" Item,omitempty" json:"Item,omitempty"`
}

type AssetCustomizationItem struct {
	AttrColor string `xml:" color,attr"  json:",omitempty"`
	AttrItem string `xml:" item,attr"  json:",omitempty"`
	AttrMesh string `xml:" mesh,attr"  json:",omitempty"`
}

type AssetCustomizations struct {
	AssetCustomization []*AssetCustomization `xml:" Customization,omitempty" json:"Customization,omitempty"`
}

type AssetDefaultValues struct {
	AttrAttackcollisionradius string `xml:" attackcollisionradius,attr"  json:",omitempty"`
	AttrBarehandattackspeed string `xml:" barehandattackspeed,attr"  json:",omitempty"`
	AttrBarehandphysicaldamage string `xml:" barehandphysicaldamage,attr"  json:",omitempty"`
	AttrBarehandrange string `xml:" barehandrange,attr"  json:",omitempty"`
	AttrCollisionradius string `xml:" collisionradius,attr"  json:",omitempty"`
	AttrCraftingfocusatstart string `xml:" craftingfocusatstart,attr"  json:",omitempty"`
	AttrCraftingfocusmax string `xml:" craftingfocusmax,attr"  json:",omitempty"`
	AttrCraftingfocusregenerationperday string `xml:" craftingfocusregenerationperday,attr"  json:",omitempty"`
	AttrEnergymax string `xml:" energymax,attr"  json:",omitempty"`
	AttrEnergyregeneration string `xml:" energyregeneration,attr"  json:",omitempty"`
	AttrHarvestrange string `xml:" harvestrange,attr"  json:",omitempty"`
	AttrHitpointsmax string `xml:" hitpointsmax,attr"  json:",omitempty"`
	AttrHitpointsregeneration string `xml:" hitpointsregeneration,attr"  json:",omitempty"`
	AttrInventorysize string `xml:" inventorysize,attr"  json:",omitempty"`
	AttrLearningpointsatstart string `xml:" learningpointsatstart,attr"  json:",omitempty"`
	AttrLearningpointsecondsperpoint string `xml:" learningpointsecondsperpoint,attr"  json:",omitempty"`
	AttrLearningpointsecondsperpointpremium string `xml:" learningpointsecondsperpointpremium,attr"  json:",omitempty"`
	AttrLearningpointsmax string `xml:" learningpointsmax,attr"  json:",omitempty"`
	AttrLearningpointsmaxpremium string `xml:" learningpointsmaxpremium,attr"  json:",omitempty"`
	AttrMagicattackdamagetime string `xml:" magicattackdamagetime,attr"  json:",omitempty"`
	AttrMagicattackspeed string `xml:" magicattackspeed,attr"  json:",omitempty"`
	AttrMaxfoodslotstacksize string `xml:" maxfoodslotstacksize,attr"  json:",omitempty"`
	AttrMaxload string `xml:" maxload,attr"  json:",omitempty"`
	AttrMaxpotionslotstacksize string `xml:" maxpotionslotstacksize,attr"  json:",omitempty"`
	AttrMeleeattackdamagetime string `xml:" meleeattackdamagetime,attr"  json:",omitempty"`
	AttrMonolithrange string `xml:" monolithrange,attr"  json:",omitempty"`
	AttrMovespeed string `xml:" movespeed,attr"  json:",omitempty"`
	AttrOrbclaimrange string `xml:" orbclaimrange,attr"  json:",omitempty"`
	AttrPhysicalattackspeed string `xml:" physicalattackspeed,attr"  json:",omitempty"`
	AttrPremiumanimalgrowthfactor string `xml:" premiumanimalgrowthfactor,attr"  json:",omitempty"`
	AttrPremiumcraftingfocusmax string `xml:" premiumcraftingfocusmax,attr"  json:",omitempty"`
	AttrPremiumcraftingfocusregenerationperday string `xml:" premiumcraftingfocusregenerationperday,attr"  json:",omitempty"`
	AttrPremiumcropyieldfactor string `xml:" premiumcropyieldfactor,attr"  json:",omitempty"`
	AttrPremiumdestinyboardprogressionfactor string `xml:" premiumdestinyboardprogressionfactor,attr"  json:",omitempty"`
	AttrPremiumgatheringyieldbonuschance string `xml:" premiumgatheringyieldbonuschance,attr"  json:",omitempty"`
	AttrPremiummarkettransactionfeefactor string `xml:" premiummarkettransactionfeefactor,attr"  json:",omitempty"`
	AttrPremiummoblootfactor string `xml:" premiummoblootfactor,attr"  json:",omitempty"`
	AttrPremiumsilverlootfactor string `xml:" premiumsilverlootfactor,attr"  json:",omitempty"`
	AttrRangedattackdamagetime string `xml:" rangedattackdamagetime,attr"  json:",omitempty"`
	AttrTreasurechestrange string `xml:" treasurechestrange,attr"  json:",omitempty"`
	AttrUsecraftbuildingrange string `xml:" usecraftbuildingrange,attr"  json:",omitempty"`
	AttrUseobjectrange string `xml:" useobjectrange,attr"  json:",omitempty"`
	AttrWarcamprange string `xml:" warcamprange,attr"  json:",omitempty"`
	AssetToolsForActionAnimation *AssetToolsForActionAnimation `xml:" ToolsForActionAnimation,omitempty" json:"ToolsForActionAnimation,omitempty"`
}

type AssetFaceState struct {
	AttrMesh string `xml:" mesh,attr"  json:",omitempty"`
	AttrState string `xml:" state,attr"  json:",omitempty"`
}

type AssetFemale struct {
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
}

type AssetFootStepVfxPreset struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetItem struct {
	AttrColorid string `xml:" colorid,attr"  json:",omitempty"`
	AttrColorlocatag string `xml:" colorlocatag,attr"  json:",omitempty"`
	AttrColors string `xml:" colors,attr"  json:",omitempty"`
	AttrReplaceslot string `xml:" replaceslot,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AssetMesh []*AssetMesh `xml:" Mesh,omitempty" json:"Mesh,omitempty"`
}

type AssetMale struct {
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
}

type AssetMesh struct {
	AttrMesh string `xml:" mesh,attr"  json:",omitempty"`
	AttrNamelocatag string `xml:" namelocatag,attr"  json:",omitempty"`
	AssetBeardState []*AssetBeardState `xml:" BeardState,omitempty" json:"BeardState,omitempty"`
	AssetFaceState []*AssetFaceState `xml:" FaceState,omitempty" json:"FaceState,omitempty"`
}

type AssetPreset struct {
	AttrAvatar string `xml:" avatar,attr"  json:",omitempty"`
	AttrCharactertype string `xml:" charactertype,attr"  json:",omitempty"`
	AssetCustomizationItem []*AssetCustomizationItem `xml:" CustomizationItem,omitempty" json:"CustomizationItem,omitempty"`
}

type AssetRace struct {
	AttrAvatar string `xml:" avatar,attr"  json:",omitempty"`
	AttrNamelocatag string `xml:" namelocatag,attr"  json:",omitempty"`
	AssetFemale *AssetFemale `xml:" Female,omitempty" json:"Female,omitempty"`
	AssetMale *AssetMale `xml:" Male,omitempty" json:"Male,omitempty"`
}

type AssetRaces struct {
	AssetRace *AssetRace `xml:" Race,omitempty" json:"Race,omitempty"`
}

type AssetSocketPreset struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetToolsForActionAnimation struct {
	AssetCareing *AssetCareing `xml:" careing,omitempty" json:"careing,omitempty"`
	AssetCrafting *AssetCrafting `xml:" crafting,omitempty" json:"crafting,omitempty"`
	AssetInstalling *AssetInstalling `xml:" installing,omitempty" json:"installing,omitempty"`
	AssetPlacing *AssetPlacing `xml:" placing,omitempty" json:"placing,omitempty"`
}

type AssetVfxPreset struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetCareing struct {
	AttrDefaultmesh string `xml:" defaultmesh,attr"  json:",omitempty"`
	AssetTarget []*AssetTarget `xml:" target,omitempty" json:"target,omitempty"`
}

type AssetCrafting struct {
	AttrDefaultmesh string `xml:" defaultmesh,attr"  json:",omitempty"`
	AssetTarget []*AssetTarget `xml:" target,omitempty" json:"target,omitempty"`
}

type AssetInstalling struct {
	AttrDefaultmesh string `xml:" defaultmesh,attr"  json:",omitempty"`
}

type AssetPlacing struct {
	AttrDefaultmesh string `xml:" defaultmesh,attr"  json:",omitempty"`
	AssetTarget []*AssetTarget `xml:" target,omitempty" json:"target,omitempty"`
}

type AssetTarget struct {
	AttrId string `xml:" id,attr"  json:",omitempty"`
	AttrMesh string `xml:" mesh,attr"  json:",omitempty"`
}


