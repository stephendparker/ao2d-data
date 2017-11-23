package world

type AssetChidleyRoot314159 struct {
	AssetWorld *AssetWorld `xml:" world,omitempty" json:"world,omitempty"`
}

type AssetHellgateExit struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
	AttrUnused string `xml:" unused,attr"  json:",omitempty"`
}

type AssetActionbuilding struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
	AttrUnused string `xml:" unused,attr"  json:",omitempty"`
}

type AssetAgent struct {
	AttrFirstname string `xml:" firstname,attr"  json:",omitempty"`
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrId string `xml:" id,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
	AttrPos string `xml:" pos,attr"  json:",omitempty"`
	AttrSurname string `xml:" surname,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
}

type AssetAgents struct {
	AssetAgent []*AssetAgent `xml:" agent,omitempty" json:"agent,omitempty"`
}

type AssetArenaagent struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
}

type AssetCluster struct {
	AttrCategoryname string `xml:" categoryname,attr"  json:",omitempty"`
	AttrDisplayname string `xml:" displayname,attr"  json:",omitempty"`
	AttrEditoroffset string `xml:" editoroffset,attr"  json:",omitempty"`
	AttrEnabled string `xml:" enabled,attr"  json:",omitempty"`
	AttrFile string `xml:" file,attr"  json:",omitempty"`
	AttrId string `xml:" id,attr"  json:",omitempty"`
	AttrMinimapBoundsMax string `xml:" minimapBoundsMax,attr"  json:",omitempty"`
	AttrMinimapBoundsMin string `xml:" minimapBoundsMin,attr"  json:",omitempty"`
	AttrOrigin string `xml:" origin,attr"  json:",omitempty"`
	AttrSize string `xml:" size,attr"  json:",omitempty"`
	AttrSpeciallocation string `xml:" speciallocation,attr"  json:",omitempty"`
	AttrTimeregion string `xml:" timeregion,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
	AttrWorldmapposition string `xml:" worldmapposition,attr"  json:",omitempty"`
	AssetActionbuilding []*AssetActionbuilding `xml:" actionbuilding,omitempty" json:"actionbuilding,omitempty"`
	AssetAgent []*AssetAgent `xml:" agent,omitempty" json:"agent,omitempty"`
	AssetAgents *AssetAgents `xml:" agents,omitempty" json:"agents,omitempty"`
	AssetArenaagent *AssetArenaagent `xml:" arenaagent,omitempty" json:"arenaagent,omitempty"`
	AssetClustermonolith *AssetClustermonolith `xml:" clustermonolith,omitempty" json:"clustermonolith,omitempty"`
	AssetDistribution *AssetDistribution `xml:" distribution,omitempty" json:"distribution,omitempty"`
	AssetExit []*AssetExit `xml:" exit,omitempty" json:"exit,omitempty"`
	AssetExits *AssetExits `xml:" exits,omitempty" json:"exits,omitempty"`
	AssetExpeditionagent []*AssetExpeditionagent `xml:" expeditionagent,omitempty" json:"expeditionagent,omitempty"`
	AssetExpeditionexit []*AssetExpeditionexit `xml:" expeditionexit,omitempty" json:"expeditionexit,omitempty"`
	AssetExpeditionnarrator []*AssetExpeditionnarrator `xml:" expeditionnarrator,omitempty" json:"expeditionnarrator,omitempty"`
	AssetExpeditionstartpoint []*AssetExpeditionstartpoint `xml:" expeditionstartpoint,omitempty" json:"expeditionstartpoint,omitempty"`
	AssetExpeditionstartpoints *AssetExpeditionstartpoints `xml:" expeditionstartpoints,omitempty" json:"expeditionstartpoints,omitempty"`
	AssetFasttravelpoint *AssetFasttravelpoint `xml:" fasttravelpoint,omitempty" json:"fasttravelpoint,omitempty"`
	AssetHellgateExit []*AssetHellgateExit `xml:" HellgateExit,omitempty" json:"HellgateExit,omitempty"`
	AssetHellgateexit []*AssetHellgateexit `xml:" hellgateexit,omitempty" json:"hellgateexit,omitempty"`
	AssetHellgateexits *AssetHellgateexits `xml:" hellgateexits,omitempty" json:"hellgateexits,omitempty"`
	AssetHellgatemobspawnpoint []*AssetHellgatemobspawnpoint `xml:" hellgatemobspawnpoint,omitempty" json:"hellgatemobspawnpoint,omitempty"`
	AssetHellgatemobspawnpoints *AssetHellgatemobspawnpoints `xml:" hellgatemobspawnpoints,omitempty" json:"hellgatemobspawnpoints,omitempty"`
	AssetIslandaccesspoint *AssetIslandaccesspoint `xml:" islandaccesspoint,omitempty" json:"islandaccesspoint,omitempty"`
	AssetIslandaccesspoints *AssetIslandaccesspoints `xml:" islandaccesspoints,omitempty" json:"islandaccesspoints,omitempty"`
	AssetMinimapmarkers *AssetMinimapmarkers `xml:" minimapmarkers,omitempty" json:"minimapmarkers,omitempty"`
	AssetMobcounts *AssetMobcounts `xml:" mobcounts,omitempty" json:"mobcounts,omitempty"`
	AssetMonolith []*AssetMonolith `xml:" monolith,omitempty" json:"monolith,omitempty"`
	AssetPortalentrance *AssetPortalentrance `xml:" portalentrance,omitempty" json:"portalentrance,omitempty"`
	AssetPortalentrances *AssetPortalentrances `xml:" portalentrances,omitempty" json:"portalentrances,omitempty"`
	AssetPortalexit []*AssetPortalexit `xml:" portalexit,omitempty" json:"portalexit,omitempty"`
	AssetPortalexits *AssetPortalexits `xml:" portalexits,omitempty" json:"portalexits,omitempty"`
	AssetRealestate []*AssetRealestate `xml:" realestate,omitempty" json:"realestate,omitempty"`
	AssetRespawnpoint []*AssetRespawnpoint `xml:" respawnpoint,omitempty" json:"respawnpoint,omitempty"`
	AssetRoads *AssetRoads `xml:" roads,omitempty" json:"roads,omitempty"`
	AssetSiegecamp *AssetSiegecamp `xml:" siegecamp,omitempty" json:"siegecamp,omitempty"`
	AssetTerritories *AssetTerritories `xml:" territories,omitempty" json:"territories,omitempty"`
	AssetTravelpoints *AssetTravelpoints `xml:" travelpoints,omitempty" json:"travelpoints,omitempty"`
	AssetTreasurechest []*AssetTreasurechest `xml:" treasurechest,omitempty" json:"treasurechest,omitempty"`
	AssetTreasurechests *AssetTreasurechests `xml:" treasurechests,omitempty" json:"treasurechests,omitempty"`
	AssetWorldbosses *AssetWorldbosses `xml:" worldbosses,omitempty" json:"worldbosses,omitempty"`
	AssetWorldmapmarkers *AssetWorldmapmarkers `xml:" worldmapmarkers,omitempty" json:"worldmapmarkers,omitempty"`
}

type AssetClustermonolith struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
}

type AssetClusters struct {
	AssetCluster []*AssetCluster `xml:" cluster,omitempty" json:"cluster,omitempty"`
}

type AssetConnection struct {
	AttrCostmodifier string `xml:" costmodifier,attr"  json:",omitempty"`
	AttrNoluggage string `xml:" noluggage,attr"  json:",omitempty"`
	AttrTo string `xml:" to,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
}

type AssetConnections struct {
	AssetConnection []*AssetConnection `xml:" connection,omitempty" json:"connection,omitempty"`
}

type AssetCriticalObjects struct {
	AssetCluster []*AssetCluster `xml:" cluster,omitempty" json:"cluster,omitempty"`
}

type AssetDistribution struct {
	AttrRealestates string `xml:" realestates,attr"  json:",omitempty"`
	AssetResource []*AssetResource `xml:" resource,omitempty" json:"resource,omitempty"`
}

type AssetExit struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrId string `xml:" id,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
	AttrPos string `xml:" pos,attr"  json:",omitempty"`
	AttrRestricted string `xml:" restricted,attr"  json:",omitempty"`
	AttrTargetid string `xml:" targetid,attr"  json:",omitempty"`
	AttrTargettype string `xml:" targettype,attr"  json:",omitempty"`
	AttrTerritory string `xml:" territory,attr"  json:",omitempty"`
	AttrUnused string `xml:" unused,attr"  json:",omitempty"`
}

type AssetExits struct {
	AssetExit []*AssetExit `xml:" exit,omitempty" json:"exit,omitempty"`
}

type AssetExpeditionagent struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
	AttrUnused string `xml:" unused,attr"  json:",omitempty"`
}

type AssetExpeditionexit struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
	AttrUnused string `xml:" unused,attr"  json:",omitempty"`
}

type AssetExpeditionnarrator struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
	AttrUnused string `xml:" unused,attr"  json:",omitempty"`
}

type AssetExpeditionstartpoint struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrId string `xml:" id,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
	AttrPos string `xml:" pos,attr"  json:",omitempty"`
	AttrUnused string `xml:" unused,attr"  json:",omitempty"`
}

type AssetExpeditionstartpoints struct {
	AssetExpeditionstartpoint []*AssetExpeditionstartpoint `xml:" expeditionstartpoint,omitempty" json:"expeditionstartpoint,omitempty"`
}

type AssetFasttravelpoint struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
	AttrUnused string `xml:" unused,attr"  json:",omitempty"`
}

type AssetHellgateexit struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrId string `xml:" id,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
	AttrPos string `xml:" pos,attr"  json:",omitempty"`
	AttrUnused string `xml:" unused,attr"  json:",omitempty"`
}

type AssetHellgateexits struct {
	AssetHellgateexit []*AssetHellgateexit `xml:" hellgateexit,omitempty" json:"hellgateexit,omitempty"`
}

type AssetHellgatemobspawnpoint struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrId string `xml:" id,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
	AttrPos string `xml:" pos,attr"  json:",omitempty"`
	AttrUnused string `xml:" unused,attr"  json:",omitempty"`
}

type AssetHellgatemobspawnpoints struct {
	AssetHellgatemobspawnpoint []*AssetHellgatemobspawnpoint `xml:" hellgatemobspawnpoint,omitempty" json:"hellgatemobspawnpoint,omitempty"`
}

type AssetIslandaccesspoint struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrId string `xml:" id,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
	AttrTravelpointid string `xml:" travelpointid,attr"  json:",omitempty"`
}

type AssetIslandaccesspoints struct {
	AssetIslandaccesspoint *AssetIslandaccesspoint `xml:" islandaccesspoint,omitempty" json:"islandaccesspoint,omitempty"`
}

type AssetMarker struct {
	AttrPos string `xml:" pos,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
}

type AssetMinimapmarkers struct {
	AssetMarker []*AssetMarker `xml:" marker,omitempty" json:"marker,omitempty"`
}

type AssetMob struct {
	AttrCount string `xml:" count,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

type AssetMobcounts struct {
	AssetMob []*AssetMob `xml:" mob,omitempty" json:"mob,omitempty"`
}

type AssetMonolith struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
	AttrUnused string `xml:" unused,attr"  json:",omitempty"`
}

type AssetPortalentrance struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrId string `xml:" id,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
	AttrPos string `xml:" pos,attr"  json:",omitempty"`
	AttrUnused string `xml:" unused,attr"  json:",omitempty"`
}

type AssetPortalentrances struct {
	AssetPortalentrance *AssetPortalentrance `xml:" portalentrance,omitempty" json:"portalentrance,omitempty"`
}

type AssetPortalexit struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrId string `xml:" id,attr"  json:",omitempty"`
	AttrKind string `xml:" kind,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
	AttrPos string `xml:" pos,attr"  json:",omitempty"`
	AttrTargetid string `xml:" targetid,attr"  json:",omitempty"`
	AttrUnused string `xml:" unused,attr"  json:",omitempty"`
}

type AssetPortalexits struct {
	AssetPortalexit []*AssetPortalexit `xml:" portalexit,omitempty" json:"portalexit,omitempty"`
}

type AssetRealestate struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
}

type AssetResource struct {
	AttrCount string `xml:" count,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrTier string `xml:" tier,attr"  json:",omitempty"`
}

type AssetRespawnpoint struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
	AttrUnused string `xml:" unused,attr"  json:",omitempty"`
}

type AssetRoads struct {
	AttrLinks string `xml:" links,attr"  json:",omitempty"`
	AttrNodes string `xml:" nodes,attr"  json:",omitempty"`
}

type AssetSiegecamp struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
	AttrUnused string `xml:" unused,attr"  json:",omitempty"`
}

type AssetTerritories struct {
	AssetTerritory []*AssetTerritory `xml:" territory,omitempty" json:"territory,omitempty"`
}

type AssetTerritory struct {
	AttrCenter string `xml:" center,attr"  json:",omitempty"`
	AttrId string `xml:" id,attr"  json:",omitempty"`
	AttrMonolith string `xml:" monolith,attr"  json:",omitempty"`
	AttrName string `xml:" name,attr"  json:",omitempty"`
	AttrResourcetype string `xml:" resourcetype,attr"  json:",omitempty"`
	AttrSize string `xml:" size,attr"  json:",omitempty"`
	AttrTerritorytype string `xml:" territorytype,attr"  json:",omitempty"`
	AssetConnections *AssetConnections `xml:" connections,omitempty" json:"connections,omitempty"`
}

type AssetTravelpoint struct {
	AttrId string `xml:" id,attr"  json:",omitempty"`
	AttrPos string `xml:" pos,attr"  json:",omitempty"`
	AssetConnections *AssetConnections `xml:" connections,omitempty" json:"connections,omitempty"`
}

type AssetTravelpoints struct {
	AssetTravelpoint *AssetTravelpoint `xml:" travelpoint,omitempty" json:"travelpoint,omitempty"`
}

type AssetTreasurechest struct {
	AttrGuid string `xml:" guid,attr"  json:",omitempty"`
	AttrId string `xml:" id,attr"  json:",omitempty"`
	AttrPath string `xml:" path,attr"  json:",omitempty"`
	AttrPos string `xml:" pos,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
	AttrUnused string `xml:" unused,attr"  json:",omitempty"`
}

type AssetTreasurechests struct {
	AssetTreasurechest []*AssetTreasurechest `xml:" treasurechest,omitempty" json:"treasurechest,omitempty"`
}

type AssetWorld struct {
	AttrBoundsmax string `xml:" boundsmax,attr"  json:",omitempty"`
	AttrBoundsmin string `xml:" boundsmin,attr"  json:",omitempty"`
	AttrScale string `xml:" scale,attr"  json:",omitempty"`
	AttrVersion string `xml:" version,attr"  json:",omitempty"`
	AttrViewportpos string `xml:" viewportpos,attr"  json:",omitempty"`
	AttrViewportzoom string `xml:" viewportzoom,attr"  json:",omitempty"`
	AssetClusters *AssetClusters `xml:" clusters,omitempty" json:"clusters,omitempty"`
	AssetCriticalObjects *AssetCriticalObjects `xml:" criticalObjects,omitempty" json:"criticalObjects,omitempty"`
}

type AssetWorldboss struct {
	AttrPos string `xml:" pos,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
}

type AssetWorldbosses struct {
	AssetWorldboss []*AssetWorldboss `xml:" worldboss,omitempty" json:"worldboss,omitempty"`
}

type AssetWorldmapmarkers struct {
	AssetMarker []*AssetMarker `xml:" marker,omitempty" json:"marker,omitempty"`
}


