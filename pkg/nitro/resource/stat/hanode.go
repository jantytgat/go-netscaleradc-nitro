package stat

var HaNodeFieldNames = struct {
	ClearStats                    string
	CurrentMasterState            string
	CurrentState                  string
	CurrentStatus                 string
	HeartbeatReceivedPacketRate   string
	HeartbeatReceivedPackets      string
	HeartbeatSentPacketRate       string
	HeartbeatSentPackets          string
	PropagationTimeoutCounter     string
	SynchronizationFailureCounter string
	TransitionTime                string
}{
	ClearStats:                    "clearstats",
	CurrentMasterState:            "hacurmasterstate",
	CurrentState:                  "hacurstate",
	CurrentStatus:                 "hacurstatus",
	HeartbeatReceivedPacketRate:   "hapktrxrate",
	HeartbeatReceivedPackets:      "hatotpktrx",
	HeartbeatSentPacketRate:       "hapkttxrate",
	HeartbeatSentPackets:          "hatotpkttx",
	PropagationTimeoutCounter:     "haerrproptimeout",
	SynchronizationFailureCounter: "haerrsyncfailure",
	TransitionTime:                "transtime",
}

type HaNode struct {
	ClearStats                    string  `json:"clearstats,omitempty" nitro:"permission=readwrite"`
	CurrentMasterState            string  `json:"hacurmasterstate,omitempty" nitro:"permission=readonly"`
	CurrentState                  string  `json:"hacurstate,omitempty" nitro:"permission=readonly"`
	CurrentStatus                 string  `json:"hacurstatus,omitempty" nitro:"permission=readonly"`
	HeartbeatReceivedPacketRate   float64 `json:"hapktrxrate,omitempty" nitro:"permission=readonly"`
	HeartbeatReceivedPackets      string  `json:"hatotpktrx,omitempty" nitro:"permission=readonly"`
	HeartbeatSentPacketRate       float64 `json:"hapkttxrate,omitempty" nitro:"permission=readonly"`
	HeartbeatSentPackets          string  `json:"hatotpkttx,omitempty" nitro:"permission=readonly"`
	PropagationTimeoutCounter     string  `json:"haerrproptimeout,omitempty" nitro:"permission=readonly"`
	SynchronizationFailureCounter string  `json:"haerrsyncfailure,omitempty" nitro:"permission=readonly"`
	TransitionTime                string  `json:"transtime,omitempty" nitro:"permission=readonly"`
}

func (r HaNode) GetTypeName() string {
	return "hanode"
}
