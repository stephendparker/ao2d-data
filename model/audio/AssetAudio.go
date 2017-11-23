package audio

type AssetArmorType struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetAudioInfos struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetWwiseEventInfo []*AssetWwiseEventInfo `xml:" WwiseEventInfo,omitempty" json:"WwiseEventInfo,omitempty"`
	AssetWwiseGlobalEnvironmentEffect []*AssetWwiseGlobalEnvironmentEffect `xml:" WwiseGlobalEnvironmentEffect,omitempty" json:"WwiseGlobalEnvironmentEffect,omitempty"`
	AssetWwiseSwitchInfo []*AssetWwiseSwitchInfo `xml:" WwiseSwitchInfo,omitempty" json:"WwiseSwitchInfo,omitempty"`
}

type AssetCharType struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetChidleyRoot314159 struct {
	AssetAudioInfos *AssetAudioInfos `xml:" AudioInfos,omitempty" json:"AudioInfos,omitempty"`
}

type AssetClusterTypes struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetGender struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetImpactType struct {
}

type AssetMobFaction struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetMobType struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetMountType struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetOnCastFinished struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AssetImpactType *AssetImpactType `xml:" ImpactType,omitempty" json:"ImpactType,omitempty"`
}

type AssetOnCastStart struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetOnChannelingEnd struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetOnChannelingHit struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetOnChannelingInterrupted struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetOnChannelingStart struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetOnOpened struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetOnPlaced struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetOnProjectileHit struct {
	AssetImpactType *AssetImpactType `xml:" ImpactType,omitempty" json:"ImpactType,omitempty"`
}

type AssetOnProjectileInit struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetOnSpellHit struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AssetImpactType *AssetImpactType `xml:" ImpactType,omitempty" json:"ImpactType,omitempty"`
}

type AssetOnSpellRemoved struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetOnViewInitialized struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetSize struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetWeaponType struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetWwiseEventInfo struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AssetOnCastFinished *AssetOnCastFinished `xml:" OnCastFinished,omitempty" json:"OnCastFinished,omitempty"`
	AssetOnCastStart *AssetOnCastStart `xml:" OnCastStart,omitempty" json:"OnCastStart,omitempty"`
	AssetOnChannelingEnd *AssetOnChannelingEnd `xml:" OnChannelingEnd,omitempty" json:"OnChannelingEnd,omitempty"`
	AssetOnChannelingHit *AssetOnChannelingHit `xml:" OnChannelingHit,omitempty" json:"OnChannelingHit,omitempty"`
	AssetOnChannelingInterrupted *AssetOnChannelingInterrupted `xml:" OnChannelingInterrupted,omitempty" json:"OnChannelingInterrupted,omitempty"`
	AssetOnChannelingStart *AssetOnChannelingStart `xml:" OnChannelingStart,omitempty" json:"OnChannelingStart,omitempty"`
	AssetOnOpened *AssetOnOpened `xml:" OnOpened,omitempty" json:"OnOpened,omitempty"`
	AssetOnPlaced *AssetOnPlaced `xml:" OnPlaced,omitempty" json:"OnPlaced,omitempty"`
	AssetOnProjectileHit *AssetOnProjectileHit `xml:" OnProjectileHit,omitempty" json:"OnProjectileHit,omitempty"`
	AssetOnProjectileInit *AssetOnProjectileInit `xml:" OnProjectileInit,omitempty" json:"OnProjectileInit,omitempty"`
	AssetOnSpellHit *AssetOnSpellHit `xml:" OnSpellHit,omitempty" json:"OnSpellHit,omitempty"`
	AssetOnSpellRemoved *AssetOnSpellRemoved `xml:" OnSpellRemoved,omitempty" json:"OnSpellRemoved,omitempty"`
	AssetOnViewInitialized *AssetOnViewInitialized `xml:" OnViewInitialized,omitempty" json:"OnViewInitialized,omitempty"`
}

type AssetWwiseGlobalEnvironmentEffect struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AssetClusterTypes *AssetClusterTypes `xml:" ClusterTypes,omitempty" json:"ClusterTypes,omitempty"`
	AssetMobFaction *AssetMobFaction `xml:" MobFaction,omitempty" json:"MobFaction,omitempty"`
}

type AssetWwiseSwitchInfo struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AssetArmorType *AssetArmorType `xml:" ArmorType,omitempty" json:"ArmorType,omitempty"`
	AssetCharType *AssetCharType `xml:" CharType,omitempty" json:"CharType,omitempty"`
	AssetGender *AssetGender `xml:" Gender,omitempty" json:"Gender,omitempty"`
	AssetMobFaction *AssetMobFaction `xml:" MobFaction,omitempty" json:"MobFaction,omitempty"`
	AssetMobType *AssetMobType `xml:" MobType,omitempty" json:"MobType,omitempty"`
	AssetMountType *AssetMountType `xml:" MountType,omitempty" json:"MountType,omitempty"`
	AssetSize *AssetSize `xml:" Size,omitempty" json:"Size,omitempty"`
	AssetWeaponType *AssetWeaponType `xml:" WeaponType,omitempty" json:"WeaponType,omitempty"`
}


