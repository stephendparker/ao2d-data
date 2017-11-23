package assets

import (
	"encoding/xml"
	"github.com/stephendparker/ao2d-data/model/items"
	"github.com/stephendparker/ao2d-data/model/resourcedistpresets"
	"github.com/stephendparker/ao2d-data/model/worldbosses"
	"github.com/stephendparker/ao2d-data/model/loot"
	"github.com/stephendparker/ao2d-data/model/missions"
	"github.com/stephendparker/ao2d-data/model/treasures"
	"github.com/stephendparker/ao2d-data/model/expeditioncategories"
	"github.com/stephendparker/ao2d-data/model/expeditions"
	"github.com/stephendparker/ao2d-data/model/resources"
	"github.com/stephendparker/ao2d-data/model/achievements"
	"github.com/metaleap/go-util/fs"
	"github.com/pkg/errors"
	"github.com/stephendparker/ao2d-data/model/expeditionagents"
	"github.com/stephendparker/ao2d-data/model/worldsettings"
	"github.com/stephendparker/ao2d-data/model/buildings"
	"github.com/stephendparker/ao2d-data/model/agents"
	"github.com/stephendparker/ao2d-data/model/gamedata"
	"github.com/stephendparker/ao2d-data/model/spells"
	"github.com/stephendparker/ao2d-data/model/mobs"
	"github.com/stephendparker/ao2d-data/model/localization"
	"github.com/stephendparker/ao2d-data/model/territorytypes"
	"github.com/stephendparker/ao2d-data/model/world"
)

type AssetData struct {
	ItemsData            *ItemsEntryDoc
	Resourcedistpresets  *ResourcedistpresetsEntryDoc
	Worldbosses          *WorldbossesEntryDoc
	Loot                 *LootEntryDoc
	Missions             *MissionsEntryDoc
	Treasures            *TreasuresEntryDoc
	Expeditioncategories *ExpeditioncategoriesEntryDoc
	Expeditions          *ExpeditionsEntryDoc
	Resources            *ResourcesEntryDoc
	Achievements         *AchievementsEntryDoc
	Expeditionagents     *ExpeditionagentsEntryDoc
	Worldsettings        *WorldsettingsEntryDoc
	Buildings            *BuildingsEntryDoc
	Agents               *AgentsEntryDoc
	Gamedata             *GamedataEntryDoc
	Spells               *SpellsEntryDoc
	Mobs                 *MobsEntryDoc
	Localization         *LocalizationEntryDoc
	Territorytypes       *TerritorytypesEntryDoc
	World                *WorldEntryDoc
}

func NewAssetData() AssetData {
	retVal := AssetData{}

	retVal.ItemsData = &ItemsEntryDoc{}
	retVal.Resourcedistpresets = &ResourcedistpresetsEntryDoc{}
	retVal.Worldbosses = &WorldbossesEntryDoc{}
	retVal.Loot = &LootEntryDoc{}
	retVal.Missions = &MissionsEntryDoc{}
	retVal.Treasures = &TreasuresEntryDoc{}
	retVal.Expeditioncategories = &ExpeditioncategoriesEntryDoc{}
	retVal.Expeditions = &ExpeditionsEntryDoc{}
	retVal.Resources = &ResourcesEntryDoc{}
	retVal.Achievements = &AchievementsEntryDoc{}
	retVal.Expeditionagents = &ExpeditionagentsEntryDoc{}
	retVal.Worldsettings = &WorldsettingsEntryDoc{}
	retVal.Buildings = &BuildingsEntryDoc{}
	retVal.Agents = &AgentsEntryDoc{}
	retVal.Gamedata = &GamedataEntryDoc{}
	retVal.Spells = &SpellsEntryDoc{}
	retVal.Mobs = &MobsEntryDoc{}
	retVal.Localization = &LocalizationEntryDoc{}
	retVal.Territorytypes = &TerritorytypesEntryDoc{}
	retVal.World = &WorldEntryDoc{}

	return retVal
}

func LoadData(sourceFolderPath string) (data AssetData, err error) {

	data = NewAssetData()
	fileUpdates := make(map[string]interface{})

	fileUpdates["items.xml"] = data.ItemsData
	fileUpdates["resourcedistpresets.xml"] = data.Resourcedistpresets
	fileUpdates["worldbosses.xml"] = data.Worldbosses
	fileUpdates["loot.xml"] = data.Loot
	fileUpdates["missions.xml"] = data.Missions
	fileUpdates["treasures.xml"] = data.Treasures
	fileUpdates["expeditioncategories.xml"] = data.Expeditioncategories
	fileUpdates["expeditions.xml"] = data.Expeditions
	fileUpdates["resources.xml"] = data.Resources
	fileUpdates["achievements.xml"] = data.Achievements
	fileUpdates["expeditionagents.xml"] = data.Expeditionagents
	fileUpdates["worldsettings.xml"] = data.Worldsettings
	fileUpdates["buildings.xml"] = data.Buildings
	fileUpdates["agents.xml"] = data.Agents
	fileUpdates["gamedata.xml"] = data.Gamedata
	fileUpdates["spells.xml"] = data.Spells
	fileUpdates["mobs.xml"] = data.Mobs
	fileUpdates["localization.xml"] = data.Localization
	fileUpdates["territorytypes.xml"] = data.Territorytypes
	fileUpdates["cluster\\world.xml"] = data.World

	for fileName, dataObject := range fileUpdates {
		err = LoadXml(sourceFolderPath+GameDataPath+"\\"+fileName, dataObject)
		if err != nil {
			err = errors.Wrap(err, "Unable to load xml file: "+fileName)
			return
		}
	}
	return
}

func LoadXml(fileName string, objectToPopulate interface{}) error {

	dataOrig := ufs.ReadBinaryFile(fileName, true)
	err := xml.Unmarshal(dataOrig, objectToPopulate)
	return err
}

func LoadXmlFromByteArray(text []byte, objectToPopulate interface{}) error {

	err := xml.Unmarshal(text, objectToPopulate)
	return err
}

type ItemsEntryDoc struct {
	XMLName xml.Name `xml:"items"`
	items.AssetItems
}

type LocalizationEntryDoc struct {
	XMLName xml.Name `xml:"tmx"`
	localization.AssetTmx
}

type TerritorytypesEntryDoc struct {
	XMLName xml.Name `xml:"territorytypes"`
	territorytypes.AssetTerritorytypes
}

type WorldEntryDoc struct {
	XMLName xml.Name `xml:"world"`
	world.AssetWorld
}

type ResourcedistpresetsEntryDoc struct {
	XMLName xml.Name `xml:"presets"`
	resourcedistpresets.AssetPresets
}

type WorldbossesEntryDoc struct {
	XMLName xml.Name `xml:"worldbosses"`
	worldbosses.AssetWorldbosses
}

type LootEntryDoc struct {
	XMLName xml.Name `xml:"LootDefinition"`
	loot.AssetLootDefinition
}

type MissionsEntryDoc struct {
	XMLName xml.Name `xml:"missions"`
	missions.AssetMissions
}

type TreasuresEntryDoc struct {
	XMLName xml.Name `xml:"treasures"`
	treasures.AssetTreasures
}

type ExpeditioncategoriesEntryDoc struct {
	XMLName xml.Name `xml:"expeditioncategories"`
	expeditioncategories.AssetExpeditioncategories
}

type ExpeditionsEntryDoc struct {
	XMLName xml.Name `xml:"expeditions"`
	expeditions.AssetExpeditions
}

type ResourcesEntryDoc struct {
	XMLName xml.Name `xml:"AO-Resources"`
	resources.AssetAOHyphenResources
}

type AchievementsEntryDoc struct {
	XMLName xml.Name `xml:"achievements"`
	achievements.AssetAchievements
}

type ExpeditionagentsEntryDoc struct {
	XMLName xml.Name `xml:"expeditionagents"`
	expeditionagents.AssetExpeditionagents
}

type WorldsettingsEntryDoc struct {
	XMLName xml.Name `xml:"AO-WorldSettings"`
	worldsettings.AssetAOHyphenWorldSettings
}

type BuildingsEntryDoc struct {
	XMLName xml.Name `xml:"buildings"`
	buildings.AssetBuildings
}

type AgentsEntryDoc struct {
	XMLName xml.Name `xml:"agents"`
	agents.AssetAgents
}

type GamedataEntryDoc struct {
	XMLName xml.Name `xml:"AO-GameData"`
	gamedata.AssetAOHyphenGameData
}

type SpellsEntryDoc struct {
	XMLName xml.Name `xml:"spells"`
	spells.AssetSpells
}

type MobsEntryDoc struct {
	XMLName xml.Name `xml:"Mobs"`
	mobs.AssetMobs
}
