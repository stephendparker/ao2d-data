package items

type AssetAssetVfxPreset struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetAudioInfo struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetChidleyRoot314159 struct {
	AssetItems *AssetItems `xml:" items,omitempty" json:"items,omitempty"`
}

type AssetFootStepVfxPreset struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetSocketPreset struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetAcceptedfood struct {
	AttrFoodcategory string `xml:" foodcategory,attr"  json:",omitempty"`
}

type AssetCanharvest struct {
	AttrResourcetype string `xml:" resourcetype,attr"  json:",omitempty"`
}

type AssetCheatprovider struct {
}

type AssetConsumableitem struct {
	AttrAbilitypower string `xml:" abilitypower,attr"  json:",omitempty"`
	AttrConsumespell string `xml:" consumespell,attr"  json:",omitempty"`
	AttrDummyitempower string `xml:" dummyitempower,attr"  json:",omitempty"`
	AttrMaxstacksize string `xml:" maxstacksize,attr"  json:",omitempty"`
	AttrNutrition string `xml:" nutrition,attr"  json:",omitempty"`
	AttrShopcategory string `xml:" shopcategory,attr"  json:",omitempty"`
	AttrShopsubcategory1 string `xml:" shopsubcategory1,attr"  json:",omitempty"`
	AttrSlottype string `xml:" slottype,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrUiatlas string `xml:" uiatlas,attr"  json:",omitempty"`
	AttrUicraftsoundfinish string `xml:" uicraftsoundfinish,attr"  json:",omitempty"`
	AttrUicraftsoundstart string `xml:" uicraftsoundstart,attr"  json:",omitempty"`
	AttrUisprite string `xml:" uisprite,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUnlockedtocraft string `xml:" unlockedtocraft,attr"  json:",omitempty"`
	AttrUnlockedtoequip string `xml:" unlockedtoequip,attr"  json:",omitempty"`
	AttrWeight string `xml:" weight,attr"  json:",omitempty"`
	AssetCraftingrequirements []*AssetCraftingrequirements `xml:" craftingrequirements,omitempty" json:"craftingrequirements,omitempty"`
	AssetEnchantments *AssetEnchantments `xml:" enchantments,omitempty" json:"enchantments,omitempty"`
}

type AssetConsumption struct {
	AssetFood *AssetFood `xml:" food,omitempty" json:"food,omitempty"`
}

type AssetContainer struct {
	AttrCapacity string `xml:" capacity,attr"  json:",omitempty"`
	AttrWeightlimit string `xml:" weightlimit,attr"  json:",omitempty"`
}

type AssetCraftingrequirements struct {
	AttrAmountcrafted string `xml:" amountcrafted,attr"  json:",omitempty"`
	AttrCraftingfocus string `xml:" craftingfocus,attr"  json:",omitempty"`
	AttrForcesinglecraft string `xml:" forcesinglecraft,attr"  json:",omitempty"`
	AttrGold string `xml:" gold,attr"  json:",omitempty"`
	AttrSilver string `xml:" silver,attr"  json:",omitempty"`
	AttrSwaptransaction string `xml:" swaptransaction,attr"  json:",omitempty"`
	AttrTime string `xml:" time,attr"  json:",omitempty"`
	AssetCraftresource []*AssetCraftresource `xml:" craftresource,omitempty" json:"craftresource,omitempty"`
}

type AssetCraftingspelllist struct {
	AssetCraftspell []*AssetCraftspell `xml:" craftspell,omitempty" json:"craftspell,omitempty"`
}

type AssetCraftitemfame struct {
	AttrMintier string `xml:" mintier,attr"  json:",omitempty"`
	AttrValue string `xml:" value,attr"  json:",omitempty"`
	AssetValiditem []*AssetValiditem `xml:" validitem,omitempty" json:"validitem,omitempty"`
}

type AssetCraftresource struct {
	AttrCount string `xml:" count,attr"  json:",omitempty"`
	AttrEnchantmentlevel string `xml:" enchantmentlevel,attr"  json:",omitempty"`
	AttrMaxreturnamount string `xml:" maxreturnamount,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
}

type AssetCraftspell struct {
	AttrSlots string `xml:" slots,attr"  json:",omitempty"`
	AttrTag string `xml:" tag,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
}

type AssetEnchantment struct {
	AttrAbilitypower string `xml:" abilitypower,attr"  json:",omitempty"`
	AttrAttackdamage string `xml:" attackdamage,attr"  json:",omitempty"`
	AttrAttackspeedbonus string `xml:" attackspeedbonus,attr"  json:",omitempty"`
	AttrBonusccdurationvsmobs string `xml:" bonusccdurationvsmobs,attr"  json:",omitempty"`
	AttrBonusccdurationvsplayers string `xml:" bonusccdurationvsplayers,attr"  json:",omitempty"`
	AttrBonusdefensevsmobs string `xml:" bonusdefensevsmobs,attr"  json:",omitempty"`
	AttrBonusdefensevsplayers string `xml:" bonusdefensevsplayers,attr"  json:",omitempty"`
	AttrConsumespell string `xml:" consumespell,attr"  json:",omitempty"`
	AttrCrowdcontrolresistance string `xml:" crowdcontrolresistance,attr"  json:",omitempty"`
	AttrDummyitempower string `xml:" dummyitempower,attr"  json:",omitempty"`
	AttrDurability string `xml:" durability,attr"  json:",omitempty"`
	AttrEnchantmentlevel string `xml:" enchantmentlevel,attr"  json:",omitempty"`
	AttrEnergymax string `xml:" energymax,attr"  json:",omitempty"`
	AttrEnergyregenerationbonus string `xml:" energyregenerationbonus,attr"  json:",omitempty"`
	AttrItempower string `xml:" itempower,attr"  json:",omitempty"`
	AttrMagiccasttimereduction string `xml:" magiccasttimereduction,attr"  json:",omitempty"`
	AttrMagiccooldownreduction string `xml:" magiccooldownreduction,attr"  json:",omitempty"`
	AttrMagicspelldamagebonus string `xml:" magicspelldamagebonus,attr"  json:",omitempty"`
	AttrMovespeedbonus string `xml:" movespeedbonus,attr"  json:",omitempty"`
	AttrPhysicalspelldamagebonus string `xml:" physicalspelldamagebonus,attr"  json:",omitempty"`
	AttrThreatbonus string `xml:" threatbonus,attr"  json:",omitempty"`
	AssetAssetVfxPreset *AssetAssetVfxPreset `xml:" AssetVfxPreset,omitempty" json:"AssetVfxPreset,omitempty"`
	AssetCraftingrequirements []*AssetCraftingrequirements `xml:" craftingrequirements,omitempty" json:"craftingrequirements,omitempty"`
	AssetUpgraderequirements *AssetUpgraderequirements `xml:" upgraderequirements,omitempty" json:"upgraderequirements,omitempty"`
}

type AssetEnchantments struct {
	AssetEnchantment []*AssetEnchantment `xml:" enchantment,omitempty" json:"enchantment,omitempty"`
}

type AssetEquipmentitem struct {
	AttrAbilitypower string `xml:" abilitypower,attr"  json:",omitempty"`
	AttrActivespellslots string `xml:" activespellslots,attr"  json:",omitempty"`
	AttrAttackspeedbonus string `xml:" attackspeedbonus,attr"  json:",omitempty"`
	AttrBeardstate string `xml:" beardstate,attr"  json:",omitempty"`
	AttrBonusccdurationvsmobs string `xml:" bonusccdurationvsmobs,attr"  json:",omitempty"`
	AttrBonusccdurationvsplayers string `xml:" bonusccdurationvsplayers,attr"  json:",omitempty"`
	AttrBonusdefensevsmobs string `xml:" bonusdefensevsmobs,attr"  json:",omitempty"`
	AttrBonusdefensevsplayers string `xml:" bonusdefensevsplayers,attr"  json:",omitempty"`
	AttrCraftfamegainfactor string `xml:" craftfamegainfactor,attr"  json:",omitempty"`
	AttrCrowdcontrolresistance string `xml:" crowdcontrolresistance,attr"  json:",omitempty"`
	AttrDurability string `xml:" durability,attr"  json:",omitempty"`
	AttrDurabilityloss_attack string `xml:" durabilityloss_attack,attr"  json:",omitempty"`
	AttrDurabilityloss_receivedattack string `xml:" durabilityloss_receivedattack,attr"  json:",omitempty"`
	AttrDurabilityloss_receivedspell string `xml:" durabilityloss_receivedspell,attr"  json:",omitempty"`
	AttrDurabilityloss_spelluse string `xml:" durabilityloss_spelluse,attr"  json:",omitempty"`
	AttrEnergymax string `xml:" energymax,attr"  json:",omitempty"`
	AttrEnergyregenerationbonus string `xml:" energyregenerationbonus,attr"  json:",omitempty"`
	AttrFacestate string `xml:" facestate,attr"  json:",omitempty"`
	AttrHealbonus string `xml:" healbonus,attr"  json:",omitempty"`
	AttrHitpointsmax string `xml:" hitpointsmax,attr"  json:",omitempty"`
	AttrHitpointsregenerationbonus string `xml:" hitpointsregenerationbonus,attr"  json:",omitempty"`
	AttrItempower string `xml:" itempower,attr"  json:",omitempty"`
	AttrItempowerprogressiontype string `xml:" itempowerprogressiontype,attr"  json:",omitempty"`
	AttrMagicattackdamagebonus string `xml:" magicattackdamagebonus,attr"  json:",omitempty"`
	AttrMagiccasttimereduction string `xml:" magiccasttimereduction,attr"  json:",omitempty"`
	AttrMagiccooldownreduction string `xml:" magiccooldownreduction,attr"  json:",omitempty"`
	AttrMagicresistance string `xml:" magicresistance,attr"  json:",omitempty"`
	AttrMagicspelldamagebonus string `xml:" magicspelldamagebonus,attr"  json:",omitempty"`
	AttrMaxload string `xml:" maxload,attr"  json:",omitempty"`
	AttrMaxqualitylevel string `xml:" maxqualitylevel,attr"  json:",omitempty"`
	AttrMesh string `xml:" mesh,attr"  json:",omitempty"`
	AttrMovespeed string `xml:" movespeed,attr"  json:",omitempty"`
	AttrMovespeedbonus string `xml:" movespeedbonus,attr"  json:",omitempty"`
	AttrOffhandanimationtype string `xml:" offhandanimationtype,attr"  json:",omitempty"`
	AttrPassivespellslots string `xml:" passivespellslots,attr"  json:",omitempty"`
	AttrPhysicalarmor string `xml:" physicalarmor,attr"  json:",omitempty"`
	AttrPhysicalattackdamagebonus string `xml:" physicalattackdamagebonus,attr"  json:",omitempty"`
	AttrPhysicalspelldamagebonus string `xml:" physicalspelldamagebonus,attr"  json:",omitempty"`
	AttrShopcategory string `xml:" shopcategory,attr"  json:",omitempty"`
	AttrShopsubcategory1 string `xml:" shopsubcategory1,attr"  json:",omitempty"`
	AttrShowinmarketplace string `xml:" showinmarketplace,attr"  json:",omitempty"`
	AttrSkincount string `xml:" skincount,attr"  json:",omitempty"`
	AttrSlottype string `xml:" slottype,attr"  json:",omitempty"`
	AttrThreatbonus string `xml:" threatbonus,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrUiatlas string `xml:" uiatlas,attr"  json:",omitempty"`
	AttrUicraftsoundfinish string `xml:" uicraftsoundfinish,attr"  json:",omitempty"`
	AttrUicraftsoundstart string `xml:" uicraftsoundstart,attr"  json:",omitempty"`
	AttrUisprite string `xml:" uisprite,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUnlockedtocraft string `xml:" unlockedtocraft,attr"  json:",omitempty"`
	AttrUnlockedtoequip string `xml:" unlockedtoequip,attr"  json:",omitempty"`
	AttrWeight string `xml:" weight,attr"  json:",omitempty"`
	AssetAssetVfxPreset *AssetAssetVfxPreset `xml:" AssetVfxPreset,omitempty" json:"AssetVfxPreset,omitempty"`
	AssetAudioInfo *AssetAudioInfo `xml:" AudioInfo,omitempty" json:"AudioInfo,omitempty"`
	AssetCraftingrequirements []*AssetCraftingrequirements `xml:" craftingrequirements,omitempty" json:"craftingrequirements,omitempty"`
	AssetCraftingspelllist *AssetCraftingspelllist `xml:" craftingspelllist,omitempty" json:"craftingspelllist,omitempty"`
	AssetEnchantments *AssetEnchantments `xml:" enchantments,omitempty" json:"enchantments,omitempty"`
	AssetSocketPreset *AssetSocketPreset `xml:" SocketPreset,omitempty" json:"SocketPreset,omitempty"`
}

type AssetFameearned struct {
	AttrValue string `xml:" value,attr"  json:",omitempty"`
}

type AssetFamefillingmissions struct {
	AssetCraftitemfame *AssetCraftitemfame `xml:" craftitemfame,omitempty" json:"craftitemfame,omitempty"`
	AssetFameearned *AssetFameearned `xml:" fameearned,omitempty" json:"fameearned,omitempty"`
	AssetGatherfame *AssetGatherfame `xml:" gatherfame,omitempty" json:"gatherfame,omitempty"`
	AssetKillmobfame *AssetKillmobfame `xml:" killmobfame,omitempty" json:"killmobfame,omitempty"`
}

type AssetFarmableitem struct {
	AttrActivefarmactiondurationseconds string `xml:" activefarmactiondurationseconds,attr"  json:",omitempty"`
	AttrActivefarmbonus string `xml:" activefarmbonus,attr"  json:",omitempty"`
	AttrActivefarmcyclelengthseconds string `xml:" activefarmcyclelengthseconds,attr"  json:",omitempty"`
	AttrActivefarmfocuscost string `xml:" activefarmfocuscost,attr"  json:",omitempty"`
	AttrActivefarmmaxcycles string `xml:" activefarmmaxcycles,attr"  json:",omitempty"`
	AttrAnimationid string `xml:" animationid,attr"  json:",omitempty"`
	AttrCraftfamegainfactor string `xml:" craftfamegainfactor,attr"  json:",omitempty"`
	AttrDestroyable string `xml:" destroyable,attr"  json:",omitempty"`
	AttrKind string `xml:" kind,attr"  json:",omitempty"`
	AttrMaxstacksize string `xml:" maxstacksize,attr"  json:",omitempty"`
	AttrPickupable string `xml:" pickupable,attr"  json:",omitempty"`
	AttrPlacecost string `xml:" placecost,attr"  json:",omitempty"`
	AttrPlacefame string `xml:" placefame,attr"  json:",omitempty"`
	AttrPrefabname string `xml:" prefabname,attr"  json:",omitempty"`
	AttrPrefabscale string `xml:" prefabscale,attr"  json:",omitempty"`
	AttrResourcevalue string `xml:" resourcevalue,attr"  json:",omitempty"`
	AttrShopcategory string `xml:" shopcategory,attr"  json:",omitempty"`
	AttrShopsubcategory1 string `xml:" shopsubcategory1,attr"  json:",omitempty"`
	AttrShowinmarketplace string `xml:" showinmarketplace,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrTile string `xml:" tile,attr"  json:",omitempty"`
	AttrUiatlas string `xml:" uiatlas,attr"  json:",omitempty"`
	AttrUisprite string `xml:" uisprite,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUnlockedtocraft string `xml:" unlockedtocraft,attr"  json:",omitempty"`
	AttrUnlockedtoplace string `xml:" unlockedtoplace,attr"  json:",omitempty"`
	AttrWeight string `xml:" weight,attr"  json:",omitempty"`
	AssetAudioInfo *AssetAudioInfo `xml:" AudioInfo,omitempty" json:"AudioInfo,omitempty"`
	AssetConsumption *AssetConsumption `xml:" consumption,omitempty" json:"consumption,omitempty"`
	AssetCraftingrequirements []*AssetCraftingrequirements `xml:" craftingrequirements,omitempty" json:"craftingrequirements,omitempty"`
	AssetGrownitem *AssetGrownitem `xml:" grownitem,omitempty" json:"grownitem,omitempty"`
	AssetHarvest *AssetHarvest `xml:" harvest,omitempty" json:"harvest,omitempty"`
	AssetProducts *AssetProducts `xml:" products,omitempty" json:"products,omitempty"`
}

type AssetFood struct {
	AttrNutritionmax string `xml:" nutritionmax,attr"  json:",omitempty"`
	AttrSecondspernutrition string `xml:" secondspernutrition,attr"  json:",omitempty"`
	AssetAcceptedfood *AssetAcceptedfood `xml:" acceptedfood,omitempty" json:"acceptedfood,omitempty"`
}

type AssetFurnitureitem struct {
	AttrAccessrightspreset string `xml:" accessrightspreset,attr"  json:",omitempty"`
	AttrCraftfamegainfactor string `xml:" craftfamegainfactor,attr"  json:",omitempty"`
	AttrCustomizewithguildlogo string `xml:" customizewithguildlogo,attr"  json:",omitempty"`
	AttrDescriptionlocatag string `xml:" descriptionlocatag,attr"  json:",omitempty"`
	AttrDurability string `xml:" durability,attr"  json:",omitempty"`
	AttrDurabilitylossperdayfactor string `xml:" durabilitylossperdayfactor,attr"  json:",omitempty"`
	AttrDurabilitylossperusefactor string `xml:" durabilitylossperusefactor,attr"  json:",omitempty"`
	AttrEnchantmentlevel string `xml:" enchantmentlevel,attr"  json:",omitempty"`
	AttrItemvalue string `xml:" itemvalue,attr"  json:",omitempty"`
	AttrLabourerfurnituretype string `xml:" labourerfurnituretype,attr"  json:",omitempty"`
	AttrLabourerhappiness string `xml:" labourerhappiness,attr"  json:",omitempty"`
	AttrLabourersaffected string `xml:" labourersaffected,attr"  json:",omitempty"`
	AttrLabourersperfurnitureitem string `xml:" labourersperfurnitureitem,attr"  json:",omitempty"`
	AttrPlaceableindoors string `xml:" placeableindoors,attr"  json:",omitempty"`
	AttrPlaceableindungeons string `xml:" placeableindungeons,attr"  json:",omitempty"`
	AttrPlaceableonlyonislands string `xml:" placeableonlyonislands,attr"  json:",omitempty"`
	AttrPlaceableoutdoors string `xml:" placeableoutdoors,attr"  json:",omitempty"`
	AttrResidencyslots string `xml:" residencyslots,attr"  json:",omitempty"`
	AttrShopcategory string `xml:" shopcategory,attr"  json:",omitempty"`
	AttrShopsubcategory1 string `xml:" shopsubcategory1,attr"  json:",omitempty"`
	AttrShowinmarketplace string `xml:" showinmarketplace,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrTile string `xml:" tile,attr"  json:",omitempty"`
	AttrUiatlas string `xml:" uiatlas,attr"  json:",omitempty"`
	AttrUicraftsoundfinish string `xml:" uicraftsoundfinish,attr"  json:",omitempty"`
	AttrUicraftsoundstart string `xml:" uicraftsoundstart,attr"  json:",omitempty"`
	AttrUisprite string `xml:" uisprite,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUnlockedtocraft string `xml:" unlockedtocraft,attr"  json:",omitempty"`
	AttrWeight string `xml:" weight,attr"  json:",omitempty"`
	AssetCheatprovider *AssetCheatprovider `xml:" cheatprovider,omitempty" json:"cheatprovider,omitempty"`
	AssetContainer *AssetContainer `xml:" container,omitempty" json:"container,omitempty"`
	AssetCraftingrequirements []*AssetCraftingrequirements `xml:" craftingrequirements,omitempty" json:"craftingrequirements,omitempty"`
	AssetRepairkit *AssetRepairkit `xml:" repairkit,omitempty" json:"repairkit,omitempty"`
}

type AssetGatherfame struct {
	AttrMintier string `xml:" mintier,attr"  json:",omitempty"`
	AttrValue string `xml:" value,attr"  json:",omitempty"`
	AssetValiditem []*AssetValiditem `xml:" validitem,omitempty" json:"validitem,omitempty"`
}

type AssetGrownitem struct {
	AttrFame string `xml:" fame,attr"  json:",omitempty"`
	AttrGrowtime string `xml:" growtime,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AssetOffspring *AssetOffspring `xml:" offspring,omitempty" json:"offspring,omitempty"`
}

type AssetHarvest struct {
	AttrFame string `xml:" fame,attr"  json:",omitempty"`
	AttrGrowtime string `xml:" growtime,attr"  json:",omitempty"`
	AttrLootchance string `xml:" lootchance,attr"  json:",omitempty"`
	AttrLootlist string `xml:" lootlist,attr"  json:",omitempty"`
	AssetSeed *AssetSeed `xml:" seed,omitempty" json:"seed,omitempty"`
}

type AssetImpactvfx struct {
	AttrImpactsocket string `xml:" impactsocket,attr"  json:",omitempty"`
	AttrPrefab string `xml:" prefab,attr"  json:",omitempty"`
}

type AssetItems struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetConsumableitem []*AssetConsumableitem `xml:" consumableitem,omitempty" json:"consumableitem,omitempty"`
	AssetEquipmentitem []*AssetEquipmentitem `xml:" equipmentitem,omitempty" json:"equipmentitem,omitempty"`
	AssetFarmableitem []*AssetFarmableitem `xml:" farmableitem,omitempty" json:"farmableitem,omitempty"`
	AssetFurnitureitem []*AssetFurnitureitem `xml:" furnitureitem,omitempty" json:"furnitureitem,omitempty"`
	AssetJournalitem []*AssetJournalitem `xml:" journalitem,omitempty" json:"journalitem,omitempty"`
	AssetMount []*AssetMount `xml:" mount,omitempty" json:"mount,omitempty"`
	AssetSimpleitem []*AssetSimpleitem `xml:" simpleitem,omitempty" json:"simpleitem,omitempty"`
	AssetWeapon []*AssetWeapon `xml:" weapon,omitempty" json:"weapon,omitempty"`
}

type AssetJournalitem struct {
	AttrBaselootamount string `xml:" baselootamount,attr"  json:",omitempty"`
	AttrCraftfamegainfactor string `xml:" craftfamegainfactor,attr"  json:",omitempty"`
	AttrMaxfame string `xml:" maxfame,attr"  json:",omitempty"`
	AttrShopcategory string `xml:" shopcategory,attr"  json:",omitempty"`
	AttrShopsubcategory1 string `xml:" shopsubcategory1,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrUiatlas string `xml:" uiatlas,attr"  json:",omitempty"`
	AttrUisprite string `xml:" uisprite,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUnlockedtocraft string `xml:" unlockedtocraft,attr"  json:",omitempty"`
	AttrWeight string `xml:" weight,attr"  json:",omitempty"`
	AssetCraftingrequirements []*AssetCraftingrequirements `xml:" craftingrequirements,omitempty" json:"craftingrequirements,omitempty"`
	AssetFamefillingmissions *AssetFamefillingmissions `xml:" famefillingmissions,omitempty" json:"famefillingmissions,omitempty"`
	AssetLootlist *AssetLootlist `xml:" lootlist,omitempty" json:"lootlist,omitempty"`
}

type AssetKillmobfame struct {
	AttrMintier string `xml:" mintier,attr"  json:",omitempty"`
	AttrValue string `xml:" value,attr"  json:",omitempty"`
}

type AssetLoot struct {
	AttrItemamount string `xml:" itemamount,attr"  json:",omitempty"`
	AttrItemenchantmentlevel string `xml:" itemenchantmentlevel,attr"  json:",omitempty"`
	AttrItemname string `xml:" itemname,attr"  json:",omitempty"`
	AttrLabourerfame string `xml:" labourerfame,attr"  json:",omitempty"`
	AttrSilveramount string `xml:" silveramount,attr"  json:",omitempty"`
	AttrWeight string `xml:" weight,attr"  json:",omitempty"`
}

type AssetLootlist struct {
	AssetLoot []*AssetLoot `xml:" loot,omitempty" json:"loot,omitempty"`
}

type AssetMount struct {
	AttrAbilitypower string `xml:" abilitypower,attr"  json:",omitempty"`
	AttrActivespellslots string `xml:" activespellslots,attr"  json:",omitempty"`
	AttrCraftfamegainfactor string `xml:" craftfamegainfactor,attr"  json:",omitempty"`
	AttrDespawneffect string `xml:" despawneffect,attr"  json:",omitempty"`
	AttrDespawneffectscaling string `xml:" despawneffectscaling,attr"  json:",omitempty"`
	AttrDismountbuff string `xml:" dismountbuff,attr"  json:",omitempty"`
	AttrDismounttime string `xml:" dismounttime,attr"  json:",omitempty"`
	AttrDurability string `xml:" durability,attr"  json:",omitempty"`
	AttrDurabilityloss_attack string `xml:" durabilityloss_attack,attr"  json:",omitempty"`
	AttrDurabilityloss_mounting string `xml:" durabilityloss_mounting,attr"  json:",omitempty"`
	AttrDurabilityloss_receivedattack string `xml:" durabilityloss_receivedattack,attr"  json:",omitempty"`
	AttrDurabilityloss_receivedattack_mounted string `xml:" durabilityloss_receivedattack_mounted,attr"  json:",omitempty"`
	AttrDurabilityloss_receivedspell string `xml:" durabilityloss_receivedspell,attr"  json:",omitempty"`
	AttrDurabilityloss_receivedspell_mounted string `xml:" durabilityloss_receivedspell_mounted,attr"  json:",omitempty"`
	AttrDurabilityloss_spelluse string `xml:" durabilityloss_spelluse,attr"  json:",omitempty"`
	AttrEnchantmentlevel string `xml:" enchantmentlevel,attr"  json:",omitempty"`
	AttrForceddismountcooldown string `xml:" forceddismountcooldown,attr"  json:",omitempty"`
	AttrForceddismountspellcooldown string `xml:" forceddismountspellcooldown,attr"  json:",omitempty"`
	AttrFulldismountcooldown string `xml:" fulldismountcooldown,attr"  json:",omitempty"`
	AttrHalfmountedbuff string `xml:" halfmountedbuff,attr"  json:",omitempty"`
	AttrHalfmountrange string `xml:" halfmountrange,attr"  json:",omitempty"`
	AttrItempower string `xml:" itempower,attr"  json:",omitempty"`
	AttrMaxqualitylevel string `xml:" maxqualitylevel,attr"  json:",omitempty"`
	AttrMountedbuff string `xml:" mountedbuff,attr"  json:",omitempty"`
	AttrMounthitpointsmax string `xml:" mounthitpointsmax,attr"  json:",omitempty"`
	AttrMounthitpointsregeneration string `xml:" mounthitpointsregeneration,attr"  json:",omitempty"`
	AttrMounttime string `xml:" mounttime,attr"  json:",omitempty"`
	AttrPassivespellslots string `xml:" passivespellslots,attr"  json:",omitempty"`
	AttrPrefabname string `xml:" prefabname,attr"  json:",omitempty"`
	AttrPrefabscaling string `xml:" prefabscaling,attr"  json:",omitempty"`
	AttrRemountdistance string `xml:" remountdistance,attr"  json:",omitempty"`
	AttrRemounttime string `xml:" remounttime,attr"  json:",omitempty"`
	AttrShopcategory string `xml:" shopcategory,attr"  json:",omitempty"`
	AttrShopsubcategory1 string `xml:" shopsubcategory1,attr"  json:",omitempty"`
	AttrShowinmarketplace string `xml:" showinmarketplace,attr"  json:",omitempty"`
	AttrSlottype string `xml:" slottype,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrUiatlas string `xml:" uiatlas,attr"  json:",omitempty"`
	AttrUicraftsoundfinish string `xml:" uicraftsoundfinish,attr"  json:",omitempty"`
	AttrUicraftsoundstart string `xml:" uicraftsoundstart,attr"  json:",omitempty"`
	AttrUisprite string `xml:" uisprite,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUnlockedtocraft string `xml:" unlockedtocraft,attr"  json:",omitempty"`
	AttrUnlockedtoequip string `xml:" unlockedtoequip,attr"  json:",omitempty"`
	AttrWeight string `xml:" weight,attr"  json:",omitempty"`
	AssetAssetVfxPreset *AssetAssetVfxPreset `xml:" AssetVfxPreset,omitempty" json:"AssetVfxPreset,omitempty"`
	AssetAudioInfo *AssetAudioInfo `xml:" AudioInfo,omitempty" json:"AudioInfo,omitempty"`
	AssetCraftingrequirements []*AssetCraftingrequirements `xml:" craftingrequirements,omitempty" json:"craftingrequirements,omitempty"`
	AssetFootStepVfxPreset *AssetFootStepVfxPreset `xml:" FootStepVfxPreset,omitempty" json:"FootStepVfxPreset,omitempty"`
	AssetMountspelllist *AssetMountspelllist `xml:" mountspelllist,omitempty" json:"mountspelllist,omitempty"`
	AssetSocketPreset *AssetSocketPreset `xml:" SocketPreset,omitempty" json:"SocketPreset,omitempty"`
}

type AssetMountspell struct {
	AttrSpellslot string `xml:" spellslot,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
}

type AssetMountspelllist struct {
	AssetMountspell *AssetMountspell `xml:" mountspell,omitempty" json:"mountspell,omitempty"`
}

type AssetOffspring struct {
	AttrAmount string `xml:" amount,attr"  json:",omitempty"`
	AttrChance string `xml:" chance,attr"  json:",omitempty"`
}

type AssetProduct struct {
	AttrActionname string `xml:" actionname,attr"  json:",omitempty"`
	AttrFame string `xml:" fame,attr"  json:",omitempty"`
	AttrLootchance string `xml:" lootchance,attr"  json:",omitempty"`
	AttrLootlist string `xml:" lootlist,attr"  json:",omitempty"`
	AttrProductiontime string `xml:" productiontime,attr"  json:",omitempty"`
}

type AssetProducts struct {
	AssetProduct *AssetProduct `xml:" product,omitempty" json:"product,omitempty"`
}

type AssetProjectile struct {
	AttrHitsocket string `xml:" hitsocket,attr"  json:",omitempty"`
	AttrPrefab string `xml:" prefab,attr"  json:",omitempty"`
	AttrStartsocket string `xml:" startsocket,attr"  json:",omitempty"`
	AttrTimeoffset string `xml:" timeoffset,attr"  json:",omitempty"`
	AssetImpactvfx *AssetImpactvfx `xml:" impactvfx,omitempty" json:"impactvfx,omitempty"`
}

type AssetRepairkit struct {
	AttrMaxtier string `xml:" maxtier,attr"  json:",omitempty"`
	AttrRepaircostfactor string `xml:" repaircostfactor,attr"  json:",omitempty"`
}

type AssetSeed struct {
	AttrAmount string `xml:" amount,attr"  json:",omitempty"`
	AttrChance string `xml:" chance,attr"  json:",omitempty"`
}

type AssetSimpleitem struct {
	AttrCraftfamegainfactor string `xml:" craftfamegainfactor,attr"  json:",omitempty"`
	AttrDescriptionlocatag string `xml:" descriptionlocatag,attr"  json:",omitempty"`
	AttrDescvariable0 string `xml:" descvariable0,attr"  json:",omitempty"`
	AttrDescvariable1 string `xml:" descvariable1,attr"  json:",omitempty"`
	AttrEnchantmentlevel string `xml:" enchantmentlevel,attr"  json:",omitempty"`
	AttrFoodcategory string `xml:" foodcategory,attr"  json:",omitempty"`
	AttrItemvalue string `xml:" itemvalue,attr"  json:",omitempty"`
	AttrMaxstacksize string `xml:" maxstacksize,attr"  json:",omitempty"`
	AttrNutrition string `xml:" nutrition,attr"  json:",omitempty"`
	AttrResourcetype string `xml:" resourcetype,attr"  json:",omitempty"`
	AttrSalvageable string `xml:" salvageable,attr"  json:",omitempty"`
	AttrShopcategory string `xml:" shopcategory,attr"  json:",omitempty"`
	AttrShopsubcategory1 string `xml:" shopsubcategory1,attr"  json:",omitempty"`
	AttrShowinmarketplace string `xml:" showinmarketplace,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrUiatlas string `xml:" uiatlas,attr"  json:",omitempty"`
	AttrUicraftsoundfinish string `xml:" uicraftsoundfinish,attr"  json:",omitempty"`
	AttrUicraftsoundstart string `xml:" uicraftsoundstart,attr"  json:",omitempty"`
	AttrUisprite string `xml:" uisprite,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUnlockedtocraft string `xml:" unlockedtocraft,attr"  json:",omitempty"`
	AttrWeight string `xml:" weight,attr"  json:",omitempty"`
	AssetCraftingrequirements []*AssetCraftingrequirements `xml:" craftingrequirements,omitempty" json:"craftingrequirements,omitempty"`
}

type AssetUpgraderequirements struct {
	AssetUpgraderesource *AssetUpgraderesource `xml:" upgraderesource,omitempty" json:"upgraderesource,omitempty"`
}

type AssetUpgraderesource struct {
	AttrCount string `xml:" count,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
}

type AssetValiditem struct {
	AttrId string `xml:" id,attr"  json:",omitempty"`
}

type AssetWeapon struct {
	AttrAbilitypower string `xml:" abilitypower,attr"  json:",omitempty"`
	AttrActivespellslots string `xml:" activespellslots,attr"  json:",omitempty"`
	AttrAttackbuildingdamage string `xml:" attackbuildingdamage,attr"  json:",omitempty"`
	AttrAttackdamage string `xml:" attackdamage,attr"  json:",omitempty"`
	AttrAttackdamagetimefactor string `xml:" attackdamagetimefactor,attr"  json:",omitempty"`
	AttrAttackrange string `xml:" attackrange,attr"  json:",omitempty"`
	AttrAttackspeed string `xml:" attackspeed,attr"  json:",omitempty"`
	AttrAttacktype string `xml:" attacktype,attr"  json:",omitempty"`
	AttrDurability string `xml:" durability,attr"  json:",omitempty"`
	AttrDurabilityloss_attack string `xml:" durabilityloss_attack,attr"  json:",omitempty"`
	AttrDurabilityloss_receivedattack string `xml:" durabilityloss_receivedattack,attr"  json:",omitempty"`
	AttrDurabilityloss_receivedspell string `xml:" durabilityloss_receivedspell,attr"  json:",omitempty"`
	AttrDurabilityloss_spelluse string `xml:" durabilityloss_spelluse,attr"  json:",omitempty"`
	AttrFocusfireprotectionpeneration string `xml:" focusfireprotectionpeneration,attr"  json:",omitempty"`
	AttrFxbonename string `xml:" fxbonename,attr"  json:",omitempty"`
	AttrFxboneoffset string `xml:" fxboneoffset,attr"  json:",omitempty"`
	AttrHitpointsmax string `xml:" hitpointsmax,attr"  json:",omitempty"`
	AttrHitpointsregenerationbonus string `xml:" hitpointsregenerationbonus,attr"  json:",omitempty"`
	AttrItempower string `xml:" itempower,attr"  json:",omitempty"`
	AttrItempowerprogressiontype string `xml:" itempowerprogressiontype,attr"  json:",omitempty"`
	AttrMagicspelldamagebonus string `xml:" magicspelldamagebonus,attr"  json:",omitempty"`
	AttrMainhandanimationtype string `xml:" mainhandanimationtype,attr"  json:",omitempty"`
	AttrMaxqualitylevel string `xml:" maxqualitylevel,attr"  json:",omitempty"`
	AttrMesh string `xml:" mesh,attr"  json:",omitempty"`
	AttrPassivespellslots string `xml:" passivespellslots,attr"  json:",omitempty"`
	AttrPhysicalspelldamagebonus string `xml:" physicalspelldamagebonus,attr"  json:",omitempty"`
	AttrShopcategory string `xml:" shopcategory,attr"  json:",omitempty"`
	AttrShopsubcategory1 string `xml:" shopsubcategory1,attr"  json:",omitempty"`
	AttrShowinmarketplace string `xml:" showinmarketplace,attr"  json:",omitempty"`
	AttrSlottype string `xml:" slottype,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
	AttrTwohanded string `xml:" twohanded,attr"  json:",omitempty"`
	AttrUiatlas string `xml:" uiatlas,attr"  json:",omitempty"`
	AttrUicraftsoundfinish string `xml:" uicraftsoundfinish,attr"  json:",omitempty"`
	AttrUicraftsoundstart string `xml:" uicraftsoundstart,attr"  json:",omitempty"`
	AttrUisprite string `xml:" uisprite,attr"  json:",omitempty"`
	AttrUnequipincombat string `xml:" unequipincombat,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AttrUnlockedtocraft string `xml:" unlockedtocraft,attr"  json:",omitempty"`
	AttrUnlockedtoequip string `xml:" unlockedtoequip,attr"  json:",omitempty"`
	AttrWeight string `xml:" weight,attr"  json:",omitempty"`
	AssetAudioInfo *AssetAudioInfo `xml:" AudioInfo,omitempty" json:"AudioInfo,omitempty"`
	AssetCanharvest *AssetCanharvest `xml:" canharvest,omitempty" json:"canharvest,omitempty"`
	AssetCraftingrequirements []*AssetCraftingrequirements `xml:" craftingrequirements,omitempty" json:"craftingrequirements,omitempty"`
	AssetCraftingspelllist *AssetCraftingspelllist `xml:" craftingspelllist,omitempty" json:"craftingspelllist,omitempty"`
	AssetEnchantments *AssetEnchantments `xml:" enchantments,omitempty" json:"enchantments,omitempty"`
	AssetProjectile []*AssetProjectile `xml:" projectile,omitempty" json:"projectile,omitempty"`
	AssetSocketPreset *AssetSocketPreset `xml:" SocketPreset,omitempty" json:"SocketPreset,omitempty"`
}


