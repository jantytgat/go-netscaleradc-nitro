package stat

var ServiceFieldNames = struct {
	ActiveTransactions                    string
	AverageServerTimeToFirstByte          string
	ClearStats                            string
	CurrentClientConnections              string
	CurrentLoad                           string
	CurrentReusePool                      string
	CurrentServerConnections              string
	Curtflags                             string
	FrustratingTimeToLastByteTransactions string
	HttpMaxHeaderSizePacketsCount         string
	MaxClients                            string
	MaxHttpHeaderFieldLengthExceededCount string
	MaxOutOfOrderPacketsCount             string
	Name                                  string
	PrimaryIpAddress                      string
	PrimaryPort                           string
	RequestBytesRate                      string
	RequestsRate                          string
	ResponseBytesRate                     string
	ResponsesRate                         string
	ServerEstablishedConnections          string
	ServiceOrder                          string
	ServiceType                           string
	State                                 string
	SurgeCount                            string
	Throughput                            string
	ThroughputRate                        string
	ToleratingTimeToLastByteTransactions  string
	TotalConnectionReassemblyQueueFlush   string
	TotalConnectionReassemblyQueueFlush75 string
	TotalRequestBytes                     string
	TotalRequests                         string
	TotalResponseBytes                    string
	TotalResponses                        string
	TotalServerTimeToLastByteTransactions string
	VserverServiceHits                    string
	VserverServiceHitsRate                string
}{
	ActiveTransactions:                    "activetransactions",
	AverageServerTimeToFirstByte:          "avgsvrttfb",
	ClearStats:                            "clearstats",
	CurrentClientConnections:              "curclntconnections",
	CurrentLoad:                           "curload",
	CurrentReusePool:                      "curreusepool",
	CurrentServerConnections:              "cursrvrconnections",
	Curtflags:                             "curtflags",
	FrustratingTimeToLastByteTransactions: "frustratingttlbtransactions",
	HttpMaxHeaderSizePacketsCount:         "httpmaxhdrszpkts",
	MaxClients:                            "maxclients",
	MaxHttpHeaderFieldLengthExceededCount: "httpmaxhdrfldlenpkts",
	MaxOutOfOrderPacketsCount:             "tcpmaxooopkts",
	Name:                                  "name",
	PrimaryIpAddress:                      "primaryipaddress",
	PrimaryPort:                           "primaryport",
	RequestBytesRate:                      "requestbytesrate",
	RequestsRate:                          "requestsrate",
	ResponseBytesRate:                     "responsebytesrate",
	ResponsesRate:                         "responsesrate",
	ServerEstablishedConnections:          "svrestablishedconn",
	ServiceOrder:                          "serviceorder",
	ServiceType:                           "servicetype",
	State:                                 "state",
	SurgeCount:                            "surgecount",
	Throughput:                            "throughput",
	ThroughputRate:                        "throughputrate",
	ToleratingTimeToLastByteTransactions:  "toleratingttlbtransactions",
	TotalConnectionReassemblyQueueFlush:   "totalconnreassemblyqueueflush",
	TotalConnectionReassemblyQueueFlush75: "totalconnreassemblyqueue75",
	TotalRequestBytes:                     "totalrequestbytes",
	TotalRequests:                         "totalrequests",
	TotalResponseBytes:                    "totalresponsebytes",
	TotalResponses:                        "totalresponses",
	TotalServerTimeToLastByteTransactions: "totsvrttlbtransactions",
	VserverServiceHits:                    "vsvrservicehits",
	VserverServiceHitsRate:                "vsvrservicehitsrate",
}

type Service struct {
	ActiveTransactions                    string  `json:"activetransactions,omitempty" nitro:"permission=readonly"`
	AverageServerTimeToFirstByte          string  `json:"avgsvrttfb,omitempty" nitro:"permission=readonly"`
	ClearStats                            string  `json:"clearstats,omitempty" nitro:"permission=readwrite"`
	CurrentClientConnections              string  `json:"curclntconnections,omitempty" nitro:"permission=readonly"`
	CurrentLoad                           string  `json:"curload,omitempty" nitro:"permission=readonly"`
	CurrentReusePool                      string  `json:"curreusepool,omitempty" nitro:"permission=readonly"`
	CurrentServerConnections              string  `json:"cursrvrconnections,omitempty" nitro:"permission=readonly"`
	Curtflags                             string  `json:"curtflags,omitempty" nitro:"permission=readonly"`
	FrustratingTimeToLastByteTransactions string  `json:"frustratingttlbtransactions,omitempty" nitro:"permission=readonly"`
	HttpMaxHeaderSizePacketsCount         string  `json:"httpmaxhdrszpkts,omitempty" nitro:"permission=readonly"`
	MaxClients                            string  `json:"maxclients,omitempty" nitro:"permission=readonly"`
	MaxHttpHeaderFieldLengthExceededCount string  `json:"httpmaxhdrfldlenpkts,omitempty" nitro:"permission=readonly"`
	MaxOutOfOrderPacketsCount             string  `json:"tcpmaxooopkts,omitempty" nitro:"permission=readonly"`
	Name                                  string  `json:"name,omitempty" nitro:"permission=readwrite"`
	PrimaryIpAddress                      string  `json:"primaryipaddress,omitempty" nitro:"permission=readonly"`
	PrimaryPort                           int     `json:"primaryport,omitempty" nitro:"permission=readonly"`
	RequestBytesRate                      float64 `json:"requestbytesrate,omitempty" nitro:"permission=readonly"`
	RequestsRate                          float64 `json:"requestsrate,omitempty" nitro:"permission=readonly"`
	ResponseBytesRate                     float64 `json:"responsebytesrate,omitempty" nitro:"permission=readonly"`
	ResponsesRate                         float64 `json:"responsesrate,omitempty" nitro:"permission=readonly"`
	ServerEstablishedConnections          string  `json:"svrestablishedconn,omitempty" nitro:"permission=readonly"`
	ServiceOrder                          string  `json:"serviceorder,omitempty" nitro:"permission=readonly"`
	ServiceType                           string  `json:"servicetype,omitempty" nitro:"permission=readonly"`
	State                                 string  `json:"state,omitempty" nitro:"permission=readonly"`
	SurgeCount                            string  `json:"surgecount,omitempty" nitro:"permission=readonly"`
	Throughput                            string  `json:"throughput,omitempty" nitro:"permission=readonly"`
	ThroughputRate                        float64 `json:"throughputrate,omitempty" nitro:"permission=readonly"`
	ToleratingTimeToLastByteTransactions  string  `json:"toleratingttlbtransactions,omitempty" nitro:"permission=readonly"`
	TotalConnectionReassemblyQueueFlush   string  `json:"totalconnreassemblyqueueflush,omitempty" nitro:"permission=readonly"`
	TotalConnectionReassemblyQueueFlush75 string  `json:"totalconnreassemblyqueue75,omitempty" nitro:"permission=readonly"`
	TotalRequestBytes                     string  `json:"totalrequestbytes,omitempty" nitro:"permission=readonly"`
	TotalRequests                         string  `json:"totalrequests,omitempty" nitro:"permission=readonly"`
	TotalResponseBytes                    string  `json:"totalresponsebytes,omitempty" nitro:"permission=readonly"`
	TotalResponses                        string  `json:"totalresponses,omitempty" nitro:"permission=readonly"`
	TotalServerTimeToLastByteTransactions string  `json:"totsvrttlbtransactions,omitempty" nitro:"permission=readonly"`
	VserverServiceHits                    string  `json:"vsvrservicehits,omitempty" nitro:"permission=readonly"`
	VserverServiceHitsRate                float64 `json:"vsvrservicehitsrate,omitempty" nitro:"permission=readonly"`
}

func (r Service) GetTypeName() string {
	return "service"
}
