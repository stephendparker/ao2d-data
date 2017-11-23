package buildings

type AssetChidleyRoot314159 struct {
	AssetBuildings *AssetBuildings `xml:" buildings,omitempty" json:"buildings,omitempty"`
}

type AssetArenabuilding struct {
	AttrAccessrightspreset string `xml:" accessrightspreset,attr"  json:",omitempty"`
	AttrIconAtlas string `xml:" iconAtlas,attr"  json:",omitempty"`
	AttrIconSprite string `xml:" iconSprite,attr"  json:",omitempty"`
	AttrNutritionloss string `xml:" nutritionloss,attr"  json:",omitempty"`
	AttrNutritionstorage string `xml:" nutritionstorage,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
}

type AssetBankbuilding struct {
	AttrBattlevaultsize string `xml:" battlevaultsize,attr"  json:",omitempty"`
	AttrHideonclustermap string `xml:" hideonclustermap,attr"  json:",omitempty"`
	AttrIconAtlas string `xml:" iconAtlas,attr"  json:",omitempty"`
	AttrIconSprite string `xml:" iconSprite,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrVaultsize string `xml:" vaultsize,attr"  json:",omitempty"`
	AttrWeightlimit string `xml:" weightlimit,attr"  json:",omitempty"`
}

type AssetBlackmarketbuilding struct {
	AttrHideonclustermap string `xml:" hideonclustermap,attr"  json:",omitempty"`
	AttrIconAtlas string `xml:" iconAtlas,attr"  json:",omitempty"`
	AttrIconSprite string `xml:" iconSprite,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUsesbuildingplacementrules string `xml:" usesbuildingplacementrules,attr"  json:",omitempty"`
}

type AssetBuildableground struct {
	AttrGround string `xml:" ground,attr"  json:",omitempty"`
}

type AssetBuildablegrounds struct {
	AssetBuildableground *AssetBuildableground `xml:" buildableground,omitempty" json:"buildableground,omitempty"`
}

type AssetBuildingreference struct {
	AttrTemplate string `xml:" template,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
}

type AssetBuildings struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetArenabuilding *AssetArenabuilding `xml:" arenabuilding,omitempty" json:"arenabuilding,omitempty"`
	AssetBankbuilding *AssetBankbuilding `xml:" bankbuilding,omitempty" json:"bankbuilding,omitempty"`
	AssetBlackmarketbuilding *AssetBlackmarketbuilding `xml:" blackmarketbuilding,omitempty" json:"blackmarketbuilding,omitempty"`
	AssetBuildingreference []*AssetBuildingreference `xml:" buildingreference,omitempty" json:"buildingreference,omitempty"`
	AssetCastlegate []*AssetCastlegate `xml:" castlegate,omitempty" json:"castlegate,omitempty"`
	AssetCraftbuilding []*AssetCraftbuilding `xml:" craftbuilding,omitempty" json:"craftbuilding,omitempty"`
	AssetFarmbuilding []*AssetFarmbuilding `xml:" farmbuilding,omitempty" json:"farmbuilding,omitempty"`
	AssetLabourer []*AssetLabourer `xml:" labourer,omitempty" json:"labourer,omitempty"`
	AssetMarketplacebuilding []*AssetMarketplacebuilding `xml:" marketplacebuilding,omitempty" json:"marketplacebuilding,omitempty"`
	AssetMeldbuilding *AssetMeldbuilding `xml:" meldbuilding,omitempty" json:"meldbuilding,omitempty"`
	AssetPlayerbuilding []*AssetPlayerbuilding `xml:" playerbuilding,omitempty" json:"playerbuilding,omitempty"`
	AssetRepairbuilding []*AssetRepairbuilding `xml:" repairbuilding,omitempty" json:"repairbuilding,omitempty"`
}

type AssetCastlegate struct {
	AttrDescriptionlocatag string `xml:" descriptionlocatag,attr"  json:",omitempty"`
	AttrDurability string `xml:" durability,attr"  json:",omitempty"`
	AttrNamelocatag string `xml:" namelocatag,attr"  json:",omitempty"`
	AttrRecoverytime string `xml:" recoverytime,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrUispritename string `xml:" uispritename,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
}

type AssetCombination struct {
	AttrAspect string `xml:" aspect,attr"  json:",omitempty"`
	AttrLootlist string `xml:" lootlist,attr"  json:",omitempty"`
	AttrPowerlevel string `xml:" powerlevel,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
	AssetNeededitem *AssetNeededitem `xml:" neededitem,omitempty" json:"neededitem,omitempty"`
}

type AssetCraftbuilding struct {
	AttrAccessrightspreset string `xml:" accessrightspreset,attr"  json:",omitempty"`
	AttrAllowsstudy string `xml:" allowsstudy,attr"  json:",omitempty"`
	AttrCategory string `xml:" category,attr"  json:",omitempty"`
	AttrCraftanimation string `xml:" craftanimation,attr"  json:",omitempty"`
	AttrCraftcapacity string `xml:" craftcapacity,attr"  json:",omitempty"`
	AttrCraftcapacityregeneration string `xml:" craftcapacityregeneration,attr"  json:",omitempty"`
	AttrDefaultcostfactorfriends string `xml:" defaultcostfactorfriends,attr"  json:",omitempty"`
	AttrDefaultcostfactorusers string `xml:" defaultcostfactorusers,attr"  json:",omitempty"`
	AttrDestructionreturnfactor string `xml:" destructionreturnfactor,attr"  json:",omitempty"`
	AttrDisplaygroup string `xml:" displaygroup,attr"  json:",omitempty"`
	AttrDurability string `xml:" durability,attr"  json:",omitempty"`
	AttrDurabilityloss string `xml:" durabilityloss,attr"  json:",omitempty"`
	AttrIconAtlas string `xml:" iconAtlas,attr"  json:",omitempty"`
	AttrIconSprite string `xml:" iconSprite,attr"  json:",omitempty"`
	AttrInitialnutrition string `xml:" initialnutrition,attr"  json:",omitempty"`
	AttrIsarena string `xml:" isarena,attr"  json:",omitempty"`
	AttrNutritionloss string `xml:" nutritionloss,attr"  json:",omitempty"`
	AttrNutritionstorage string `xml:" nutritionstorage,attr"  json:",omitempty"`
	AttrPlacecost string `xml:" placecost,attr"  json:",omitempty"`
	AttrRepaircostfactor string `xml:" repaircostfactor,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrTile string `xml:" tile,attr"  json:",omitempty"`
	AttrUisoundevent string `xml:" uisoundevent,attr"  json:",omitempty"`
	AttrUispritename string `xml:" uispritename,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUnlocked string `xml:" unlocked,attr"  json:",omitempty"`
	AttrUpgradeableto string `xml:" upgradeableto,attr"  json:",omitempty"`
	AttrUsesbuildingplacementrules string `xml:" usesbuildingplacementrules,attr"  json:",omitempty"`
	AttrWarningpopupstring string `xml:" warningpopupstring,attr"  json:",omitempty"`
	AssetCraftingitemlist *AssetCraftingitemlist `xml:" craftingitemlist,omitempty" json:"craftingitemlist,omitempty"`
	AssetCraftingrequirements *AssetCraftingrequirements `xml:" craftingrequirements,omitempty" json:"craftingrequirements,omitempty"`
	AssetFavoritedish *AssetFavoritedish `xml:" favoritedish,omitempty" json:"favoritedish,omitempty"`
}

type AssetCraftingitemlist struct {
	AssetCraftitem []*AssetCraftitem `xml:" craftitem,omitempty" json:"craftitem,omitempty"`
}

type AssetCraftingrequirements struct {
	AttrSilverinrvpercentage string `xml:" silverinrvpercentage,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
	AssetCraftresource []*AssetCraftresource `xml:" craftresource,omitempty" json:"craftresource,omitempty"`
}

type AssetCraftitem struct {
	AttrCategory string `xml:" category,attr"  json:",omitempty"`
	AttrDisplaygroup string `xml:" displaygroup,attr"  json:",omitempty"`
	AttrStacksize string `xml:" stacksize,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
}

type AssetCraftresource struct {
	AttrCount string `xml:" count,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUsedinrepair string `xml:" usedinrepair,attr"  json:",omitempty"`
}

type AssetDish struct {
	AttrBonus string `xml:" bonus,attr"  json:",omitempty"`
	AttrItem string `xml:" item,attr"  json:",omitempty"`
	AttrWeight string `xml:" weight,attr"  json:",omitempty"`
}

type AssetFarmableitem struct {
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
}

type AssetFarmableitems struct {
	AssetFarmableitem []*AssetFarmableitem `xml:" farmableitem,omitempty" json:"farmableitem,omitempty"`
}

type AssetFarmbuilding struct {
	AttrAccessrightspreset string `xml:" accessrightspreset,attr"  json:",omitempty"`
	AttrCategory string `xml:" category,attr"  json:",omitempty"`
	AttrDestructionreturnfactor string `xml:" destructionreturnfactor,attr"  json:",omitempty"`
	AttrDisplaygroup string `xml:" displaygroup,attr"  json:",omitempty"`
	AttrDurability string `xml:" durability,attr"  json:",omitempty"`
	AttrDurabilityloss string `xml:" durabilityloss,attr"  json:",omitempty"`
	AttrIconAtlas string `xml:" iconAtlas,attr"  json:",omitempty"`
	AttrIconSprite string `xml:" iconSprite,attr"  json:",omitempty"`
	AttrPlacecost string `xml:" placecost,attr"  json:",omitempty"`
	AttrRepaircostfactor string `xml:" repaircostfactor,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrUispritename string `xml:" uispritename,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUnlocked string `xml:" unlocked,attr"  json:",omitempty"`
	AssetBuildablegrounds *AssetBuildablegrounds `xml:" buildablegrounds,omitempty" json:"buildablegrounds,omitempty"`
	AssetCraftingrequirements *AssetCraftingrequirements `xml:" craftingrequirements,omitempty" json:"craftingrequirements,omitempty"`
	AssetFarmableitems *AssetFarmableitems `xml:" farmableitems,omitempty" json:"farmableitems,omitempty"`
}

type AssetFavoritedish struct {
	AssetDish *AssetDish `xml:" dish,omitempty" json:"dish,omitempty"`
}

type AssetJournal struct {
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
}

type AssetJournalsaccepted struct {
	AssetJournal []*AssetJournal `xml:" journal,omitempty" json:"journal,omitempty"`
}

type AssetJournalssold struct {
	AssetJournal []*AssetJournal `xml:" journal,omitempty" json:"journal,omitempty"`
}

type AssetLaborer struct {
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
}

type AssetLabourer struct {
	AttrAccessrightspreset string `xml:" accessrightspreset,attr"  json:",omitempty"`
	AttrAvatar string `xml:" avatar,attr"  json:",omitempty"`
	AttrAvatarring string `xml:" avatarring,attr"  json:",omitempty"`
	AttrDisplaygroup string `xml:" displaygroup,attr"  json:",omitempty"`
	AttrDurability string `xml:" durability,attr"  json:",omitempty"`
	AttrDurabilityloss string `xml:" durabilityloss,attr"  json:",omitempty"`
	AttrFametoprogress string `xml:" fametoprogress,attr"  json:",omitempty"`
	AttrHappinessmovementtickrate string `xml:" happinessmovementtickrate,attr"  json:",omitempty"`
	AttrHappinessunits string `xml:" happinessunits,attr"  json:",omitempty"`
	AttrHappinessunitsafterhire string `xml:" happinessunitsafterhire,attr"  json:",omitempty"`
	AttrHireprice string `xml:" hireprice,attr"  json:",omitempty"`
	AttrIconAtlas string `xml:" iconAtlas,attr"  json:",omitempty"`
	AttrIconSprite string `xml:" iconSprite,attr"  json:",omitempty"`
	AttrJoblength string `xml:" joblength,attr"  json:",omitempty"`
	AttrMaxbedcontribution string `xml:" maxbedcontribution,attr"  json:",omitempty"`
	AttrMaxtablecontribution string `xml:" maxtablecontribution,attr"  json:",omitempty"`
	AttrMaxtrophycontribution string `xml:" maxtrophycontribution,attr"  json:",omitempty"`
	AttrNamegeneratorid string `xml:" namegeneratorid,attr"  json:",omitempty"`
	AttrNutritionloss string `xml:" nutritionloss,attr"  json:",omitempty"`
	AttrNutritionstorage string `xml:" nutritionstorage,attr"  json:",omitempty"`
	AttrProfession string `xml:" profession,attr"  json:",omitempty"`
	AttrSettleintime string `xml:" settleintime,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUnlocked string `xml:" unlocked,attr"  json:",omitempty"`
	AttrUpgradeableto string `xml:" upgradeableto,attr"  json:",omitempty"`
	AttrYieldmultiplier string `xml:" yieldmultiplier,attr"  json:",omitempty"`
	AssetJournalsaccepted *AssetJournalsaccepted `xml:" journalsaccepted,omitempty" json:"journalsaccepted,omitempty"`
	AssetJournalssold *AssetJournalssold `xml:" journalssold,omitempty" json:"journalssold,omitempty"`
}

type AssetMarketplacebuilding struct {
	AttrHideonclustermap string `xml:" hideonclustermap,attr"  json:",omitempty"`
	AttrIconAtlas string `xml:" iconAtlas,attr"  json:",omitempty"`
	AttrIconSprite string `xml:" iconSprite,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUsesbuildingplacementrules string `xml:" usesbuildingplacementrules,attr"  json:",omitempty"`
}

type AssetMeldbuilding struct {
	AttrAccessrightspreset string `xml:" accessrightspreset,attr"  json:",omitempty"`
	AttrCategory string `xml:" category,attr"  json:",omitempty"`
	AttrCraftanimation string `xml:" craftanimation,attr"  json:",omitempty"`
	AttrCraftcapacity string `xml:" craftcapacity,attr"  json:",omitempty"`
	AttrCraftcapacityregeneration string `xml:" craftcapacityregeneration,attr"  json:",omitempty"`
	AttrDefaultcostfactorfriends string `xml:" defaultcostfactorfriends,attr"  json:",omitempty"`
	AttrDefaultcostfactorusers string `xml:" defaultcostfactorusers,attr"  json:",omitempty"`
	AttrDisplaygroup string `xml:" displaygroup,attr"  json:",omitempty"`
	AttrDurability string `xml:" durability,attr"  json:",omitempty"`
	AttrDurabilityloss string `xml:" durabilityloss,attr"  json:",omitempty"`
	AttrIconAtlas string `xml:" iconAtlas,attr"  json:",omitempty"`
	AttrIconSprite string `xml:" iconSprite,attr"  json:",omitempty"`
	AttrNutritionloss string `xml:" nutritionloss,attr"  json:",omitempty"`
	AttrNutritionstorage string `xml:" nutritionstorage,attr"  json:",omitempty"`
	AttrPlacecost string `xml:" placecost,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrUisoundevent string `xml:" uisoundevent,attr"  json:",omitempty"`
	AttrUispritename string `xml:" uispritename,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUnlocked string `xml:" unlocked,attr"  json:",omitempty"`
	AssetCraftingitemlist *AssetCraftingitemlist `xml:" craftingitemlist,omitempty" json:"craftingitemlist,omitempty"`
	AssetCraftingrequirements *AssetCraftingrequirements `xml:" craftingrequirements,omitempty" json:"craftingrequirements,omitempty"`
	AssetFavoritedish *AssetFavoritedish `xml:" favoritedish,omitempty" json:"favoritedish,omitempty"`
	AssetMeldingcombinations *AssetMeldingcombinations `xml:" meldingcombinations,omitempty" json:"meldingcombinations,omitempty"`
}

type AssetMeldingcombinations struct {
	AssetCombination []*AssetCombination `xml:" combination,omitempty" json:"combination,omitempty"`
}

type AssetNeededitem struct {
	AttrQuantity string `xml:" quantity,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
}

type AssetPlacelaborer struct {
	AssetLaborer []*AssetLaborer `xml:" laborer,omitempty" json:"laborer,omitempty"`
}

type AssetPlayerbuilding struct {
	AttrAccessrightspreset string `xml:" accessrightspreset,attr"  json:",omitempty"`
	AttrCategory string `xml:" category,attr"  json:",omitempty"`
	AttrDestructionreturnfactor string `xml:" destructionreturnfactor,attr"  json:",omitempty"`
	AttrDisplaygroup string `xml:" displaygroup,attr"  json:",omitempty"`
	AttrDurability string `xml:" durability,attr"  json:",omitempty"`
	AttrDurabilityloss string `xml:" durabilityloss,attr"  json:",omitempty"`
	AttrFurniturelimit string `xml:" furniturelimit,attr"  json:",omitempty"`
	AttrIconAtlas string `xml:" iconAtlas,attr"  json:",omitempty"`
	AttrIconSprite string `xml:" iconSprite,attr"  json:",omitempty"`
	AttrLaborerslimit string `xml:" laborerslimit,attr"  json:",omitempty"`
	AttrNutritionloss string `xml:" nutritionloss,attr"  json:",omitempty"`
	AttrNutritionstorage string `xml:" nutritionstorage,attr"  json:",omitempty"`
	AttrPlacecost string `xml:" placecost,attr"  json:",omitempty"`
	AttrRepaircostfactor string `xml:" repaircostfactor,attr"  json:",omitempty"`
	AttrResidentslimit string `xml:" residentslimit,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrUisoundevent string `xml:" uisoundevent,attr"  json:",omitempty"`
	AttrUispritename string `xml:" uispritename,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUnlocked string `xml:" unlocked,attr"  json:",omitempty"`
	AttrUnlockedbyfounderstate string `xml:" unlockedbyfounderstate,attr"  json:",omitempty"`
	AttrUpgradeableto string `xml:" upgradeableto,attr"  json:",omitempty"`
	AssetCraftingrequirements *AssetCraftingrequirements `xml:" craftingrequirements,omitempty" json:"craftingrequirements,omitempty"`
	AssetPlacelaborer *AssetPlacelaborer `xml:" placelaborer,omitempty" json:"placelaborer,omitempty"`
}

type AssetRepairbuilding struct {
	AttrCategory string `xml:" category,attr"  json:",omitempty"`
	AttrCraftcapacity string `xml:" craftcapacity,attr"  json:",omitempty"`
	AttrCraftcapacityregeneration string `xml:" craftcapacityregeneration,attr"  json:",omitempty"`
	AttrDefaultcostfactorfriends string `xml:" defaultcostfactorfriends,attr"  json:",omitempty"`
	AttrDefaultcostfactorusers string `xml:" defaultcostfactorusers,attr"  json:",omitempty"`
	AttrDestructionreturnfactor string `xml:" destructionreturnfactor,attr"  json:",omitempty"`
	AttrDisplaygroup string `xml:" displaygroup,attr"  json:",omitempty"`
	AttrDurability string `xml:" durability,attr"  json:",omitempty"`
	AttrDurabilityloss string `xml:" durabilityloss,attr"  json:",omitempty"`
	AttrIconAtlas string `xml:" iconAtlas,attr"  json:",omitempty"`
	AttrIconSprite string `xml:" iconSprite,attr"  json:",omitempty"`
	AttrNutritionloss string `xml:" nutritionloss,attr"  json:",omitempty"`
	AttrNutritionstorage string `xml:" nutritionstorage,attr"  json:",omitempty"`
	AttrPlacecost string `xml:" placecost,attr"  json:",omitempty"`
	AttrRepaircostfactor string `xml:" repaircostfactor,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrUisoundevent string `xml:" uisoundevent,attr"  json:",omitempty"`
	AttrUispritename string `xml:" uispritename,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUnlocked string `xml:" unlocked,attr"  json:",omitempty"`
	AttrUpgradeableto string `xml:" upgradeableto,attr"  json:",omitempty"`
	AssetCraftingrequirements *AssetCraftingrequirements `xml:" craftingrequirements,omitempty" json:"craftingrequirements,omitempty"`
	AssetFavoritedish *AssetFavoritedish `xml:" favoritedish,omitempty" json:"favoritedish,omitempty"`
}


