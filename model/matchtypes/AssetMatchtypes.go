package matchtypes

type AssetChidleyRoot314159 struct {
	AssetMatchtypes *AssetMatchtypes `xml:" matchtypes,omitempty" json:"matchtypes,omitempty"`
}

type AssetFriendlyFX struct {
	AttrBeam string `xml:" beam,attr"  json:",omitempty"`
	AttrImpact string `xml:" impact,attr"  json:",omitempty"`
}

type AssetHostileFX struct {
	AttrBeam string `xml:" beam,attr"  json:",omitempty"`
	AttrImpact string `xml:" impact,attr"  json:",omitempty"`
}

type AssetDefense struct {
	AttrAttackdamage string `xml:" attackdamage,attr"  json:",omitempty"`
	AttrAttackdistancedelay string `xml:" attackdistancedelay,attr"  json:",omitempty"`
	AttrAttackinterval string `xml:" attackinterval,attr"  json:",omitempty"`
	AttrBuffradius string `xml:" buffradius,attr"  json:",omitempty"`
	AttrBuffspell string `xml:" buffspell,attr"  json:",omitempty"`
	AttrDisengageattackrange string `xml:" disengageattackrange,attr"  json:",omitempty"`
	AttrEngageattackrange string `xml:" engageattackrange,attr"  json:",omitempty"`
	AttrIslethal string `xml:" islethal,attr"  json:",omitempty"`
	AssetFriendlyFX *AssetFriendlyFX `xml:" FriendlyFX,omitempty" json:"FriendlyFX,omitempty"`
	AssetHostileFX *AssetHostileFX `xml:" HostileFX,omitempty" json:"HostileFX,omitempty"`
}

type AssetMatchtype struct {
	AttrAdvantageticketsdifference string `xml:" advantageticketsdifference,attr"  json:",omitempty"`
	AttrAlmostlostticketthreshold string `xml:" almostlostticketthreshold,attr"  json:",omitempty"`
	AttrMatchdurationminutes string `xml:" matchdurationminutes,attr"  json:",omitempty"`
	AttrMatchtimerespawnmultiplier string `xml:" matchtimerespawnmultiplier,attr"  json:",omitempty"`
	AttrMaxplayerperteam string `xml:" maxplayerperteam,attr"  json:",omitempty"`
	AttrRespawndelay string `xml:" respawndelay,attr"  json:",omitempty"`
	AttrRoundduration string `xml:" roundduration,attr"  json:",omitempty"`
	AttrTicketdefault string `xml:" ticketdefault,attr"  json:",omitempty"`
	AttrTicketreductionbydeath string `xml:" ticketreductionbydeath,attr"  json:",omitempty"`
	AttrTimeleftorbearlyreset string `xml:" timeleftorbearlyreset,attr"  json:",omitempty"`
	AttrUniquename string `xml:" uniquename,attr"  json:",omitempty"`
	AssetTicketreductionlookup []*AssetTicketreductionlookup `xml:" ticketreductionlookup,omitempty" json:"ticketreductionlookup,omitempty"`
	AssetWarcamp *AssetWarcamp `xml:" warcamp,omitempty" json:"warcamp,omitempty"`
}

type AssetMatchtypes struct {
	AttrXsiSpaceNoNamespaceSchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema-instance noNamespaceSchemaLocation,attr"  json:",omitempty"`
	AttrXmlnsXsi string `xml:"xmlns xsi,attr"  json:",omitempty"`
	AssetMatchtype []*AssetMatchtype `xml:" matchtype,omitempty" json:"matchtype,omitempty"`
}

type AssetTicketreductionlookup struct {
	AttrOrbdifference string `xml:" orbdifference,attr"  json:",omitempty"`
	AttrTicketreduction string `xml:" ticketreduction,attr"  json:",omitempty"`
}

type AssetWarcamp struct {
	AssetDefense *AssetDefense `xml:" defense,omitempty" json:"defense,omitempty"`
}


