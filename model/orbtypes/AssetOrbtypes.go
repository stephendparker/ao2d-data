package orbtypes

type AssetChidleyRoot314159 struct {
	AssetOrbtypes *AssetOrbtypes `xml:" orbtypes,omitempty" json:"orbtypes,omitempty"`
}

type AssetChannelcapture struct {
	AttrOrbclaimtime string `xml:" orbclaimtime,attr"  json:",omitempty"`
}

type AssetConstantscore struct {
	AttrClaimedondefaultbyteam string `xml:" claimedondefaultbyteam,attr"  json:",omitempty"`
	AttrClaimedunclaimedstep string `xml:" claimedunclaimedstep,attr"  json:",omitempty"`
	AttrOrbtickinterval string `xml:" orbtickinterval,attr"  json:",omitempty"`
	AttrTicketreductionontick string `xml:" ticketreductionontick,attr"  json:",omitempty"`
	AttrTicketreductionorbclaim string `xml:" ticketreductionorbclaim,attr"  json:",omitempty"`
}

type AssetInstantscore struct {
	AttrHasreset string `xml:" hasreset,attr"  json:",omitempty"`
	AttrOrbcooldowntimeonstart string `xml:" orbcooldowntimeonstart,attr"  json:",omitempty"`
	AttrOrbresetsooninterval string `xml:" orbresetsooninterval,attr"  json:",omitempty"`
	AttrTicketreductionorbclaim string `xml:" ticketreductionorbclaim,attr"  json:",omitempty"`
}

type AssetOrbtype struct {
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AssetChannelcapture *AssetChannelcapture `xml:" channelcapture,omitempty" json:"channelcapture,omitempty"`
	AssetConstantscore *AssetConstantscore `xml:" constantscore,omitempty" json:"constantscore,omitempty"`
	AssetInstantscore *AssetInstantscore `xml:" instantscore,omitempty" json:"instantscore,omitempty"`
	AssetRoundbasedscore *AssetRoundbasedscore `xml:" roundbasedscore,omitempty" json:"roundbasedscore,omitempty"`
	AssetZonecapture *AssetZonecapture `xml:" zonecapture,omitempty" json:"zonecapture,omitempty"`
}

type AssetOrbtypes struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetOrbtype []*AssetOrbtype `xml:" orbtype,omitempty" json:"orbtype,omitempty"`
}

type AssetRoundbasedscore struct {
	AttrClaimedondefaultbyteam string `xml:" claimedondefaultbyteam,attr"  json:",omitempty"`
}

type AssetZonecapture struct {
	AttrOrbclaimtime string `xml:" orbclaimtime,attr"  json:",omitempty"`
	AttrOrbnoteamunclaimtime string `xml:" orbnoteamunclaimtime,attr"  json:",omitempty"`
	AttrOrbunclaimtime string `xml:" orbunclaimtime,attr"  json:",omitempty"`
	AttrOrbzoneradius string `xml:" orbzoneradius,attr"  json:",omitempty"`
}


