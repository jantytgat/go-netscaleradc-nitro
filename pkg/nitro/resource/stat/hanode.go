package stat

type HaNode struct {
	ClearStats                    string  `json:"clearstats,omitempty" nitro:"permission=readwrite"`
	CurrentStatus                 string  `json:"hacurstatus,omitempty" nitro:"permission=readonly"`
	CurrentState                  string  `json:"hacurstate,omitempty" nitro:"permission=readonly"`
	CurrentMasterState            string  `json:"hacurmasterstate,omitempty" nitro:"permission=readonly"`
	TransitionTime                string  `json:"transtime,omitempty" nitro:"permission=readonly"`
	HeartbeatReceivedPackets      string  `json:"hatotpktrx,omitempty" nitro:"permission=readonly"`
	HeartbeatReceivedPacketRate   float64 `json:"hapktrxrate,omitempty" nitro:"permission=readonly"`
	HeartbeatSentPackets          string  `json:"hatotpkttx,omitempty" nitro:"permission=readonly"`
	HeartbeatSentPacketRate       float64 `json:"hapkttxrate,omitempty" nitro:"permission=readonly"`
	PropagationTimeoutCounter     string  `json:"haerrproptimeout,omitempty" nitro:"permission=readonly"`
	SynchronizationFailureCounter string  `json:"haerrsyncfailure,omitempty" nitro:"permission=readonly"`
}

func (r HaNode) GetTypeName() string {
	return "hanode"
}
