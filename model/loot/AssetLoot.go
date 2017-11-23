package loot

type AssetAND struct {
	AttrChance string `xml:" chance,attr"  json:",omitempty"`
	AttrWeight string `xml:" weight,attr"  json:",omitempty"`
	AssetOR []*AssetOR `xml:" OR,omitempty" json:"OR,omitempty"`
}

type AssetChidleyRoot314159 struct {
	AssetLootDefinition *AssetLootDefinition `xml:" LootDefinition,omitempty" json:"LootDefinition,omitempty"`
}

type AssetItem struct {
	AttrAmount string `xml:" amount,attr"  json:",omitempty"`
	AttrChance string `xml:" chance,attr"  json:",omitempty"`
	AttrEnchantmentlevel string `xml:" enchantmentlevel,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
	AttrUseblackmarket string `xml:" useblackmarket,attr"  json:",omitempty"`
	AttrWeight string `xml:" weight,attr"  json:",omitempty"`
}

type AssetLootDefinition struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetLootlist []*AssetLootlist `xml:" Lootlist,omitempty" json:"Lootlist,omitempty"`
}

type AssetLootListReference struct {
	AttrChance string `xml:" chance,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrWeight string `xml:" weight,attr"  json:",omitempty"`
}

type AssetLootlist struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AssetAND *AssetAND `xml:" AND,omitempty" json:"AND,omitempty"`
	AssetItem []*AssetItem `xml:" Item,omitempty" json:"Item,omitempty"`
	AssetLootListReference []*AssetLootListReference `xml:" LootListReference,omitempty" json:"LootListReference,omitempty"`
	AssetOR []*AssetOR `xml:" OR,omitempty" json:"OR,omitempty"`
}

type AssetOR struct {
	AttrChance string `xml:" chance,attr"  json:",omitempty"`
	AttrWeight string `xml:" weight,attr"  json:",omitempty"`
	AssetAND *AssetAND `xml:" AND,omitempty" json:"AND,omitempty"`
	AssetItem []*AssetItem `xml:" Item,omitempty" json:"Item,omitempty"`
	AssetLootListReference []*AssetLootListReference `xml:" LootListReference,omitempty" json:"LootListReference,omitempty"`
	AssetOR []*AssetOR `xml:" OR,omitempty" json:"OR,omitempty"`
}


