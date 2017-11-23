package service

import (
	"github.com/stephendparker/ao2d-data/assets"
	"github.com/stephendparker/ao2d-data/model"
	"github.com/stephendparker/ao2d-data/model/items"
)

func GenerateGameInfo(assetData assets.AssetData) (fullGameInfo model.FullGameInfo) {
	fullGameInfo = model.NewFullGameInfo()

	mapItems(assetData, fullGameInfo)

	return
}
func mapItems(assetData assets.AssetData, fullGameInfo model.FullGameInfo) {
	for _, v := range assetData.ItemsData.AssetItems.AssetSimpleitem {

		item := model.ItemData{}
		item.Simpleitem = *v
		fullGameInfo.Items[v.AttrUniquename] = item

		fullGameInfo.MinItems = append(fullGameInfo.MinItems, MapItem(*v))
	}

	for _, v := range assetData.ItemsData.AssetItems.AssetWeapon {

		item := model.ItemData{}
		item.Weaponitem = *v
		fullGameInfo.Items[v.AttrUniquename] = item
	}
}

func MapItem(item items.AssetSimpleitem) (retVal model.ItemMinData) {

	retVal = model.ItemMinData{}

	return
}
