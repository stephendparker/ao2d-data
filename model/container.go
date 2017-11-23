package model

import (
	"github.com/stephendparker/ao2d-data/model/items"
	"github.com/stephendparker/ao2d-data/model/buildings"
)

type FullGameInfo struct {
	MinItems  []ItemMinData
	Items     map[string]ItemData
	Buildings map[string]BuildingData
}

func NewFullGameInfo() (retVal FullGameInfo) {
	retVal = FullGameInfo{}

	retVal.Items = make(map[string]ItemData)

	return
}

type ItemMinData struct {
	Name         string
	Tier         string
	Enchancement string
	Category     string
	Subcategory  string
}

type ItemData struct {
	Farmableitem   items.AssetFarmableitem
	Consumableitem items.AssetConsumableitem
	Equipmentitem  items.AssetEquipmentitem
	Furnitureitem  items.AssetFurnitureitem
	Journalitem    items.AssetJournalitem
	Mountitem      items.AssetMount
	Simpleitem     items.AssetSimpleitem
	Weaponitem     items.AssetWeapon
}

type BuildingData struct {
	Arenabuilding       buildings.AssetArenabuilding
	Castlegatebuilding  buildings.AssetCastlegate
	Bankbuilding        buildings.AssetBankbuilding
	Craftbuilding       buildings.AssetCraftbuilding
	Farmbuilding        buildings.AssetFarmbuilding
	Marketplacebuilding buildings.AssetMarketplacebuilding
	Playerbuilding      buildings.AssetPlayerbuilding
	Repairbuilding      buildings.AssetRepairbuilding
	Meldbuilding        buildings.AssetMeldbuilding
	Blackmarketbuilding buildings.AssetBlackmarketbuilding
}
