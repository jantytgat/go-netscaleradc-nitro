package stat

var LbVserverFieldNames = struct {
	ActiveServices                            string
	AverageClientTimeToLastByte               string
	ClearStats                                string
	ClientResponseTimeApdex                   string
	ClientTimeToLastByteTransactionsRate      string
	CpuUsagePerMille                          string
	CurrentBackupPersistenceSessions          string
	CurrentClientConnections                  string
	CurrentMultiPathTcpSessions               string
	CurrentMultipathTcpSubflowConnections     string
	CurrentPersistenceSessions                string
	CurrentServerConnections                  string
	DeferredRequests                          string
	DeferredRequestsRate                      string
	EstablishedConnections                    string
	FrustratingTimeToLastByteTransactions     string
	FrustratingTimeToLastByteTransactionsRate string
	HitsRate                                  string
	Http2RequestsRate                         string
	Http2ResponsesRate                        string
	HttpMaxHeaderFieldLengthExceededCount     string
	HttpMaxHeaderSizePackets                  string
	InactiveServices                          string
	InvalidRequestResponse                    string
	InvalidRequestResponseDropped             string
	LabelledConnections                       string
	Name                                      string
	PacketsReceivedRate                       string
	PacketsSentRate                           string
	PrimaryIpAddress                          string
	PrimaryPort                               string
	PushLabel                                 string
	RequestBytesRate                          string
	RequestRetryCount                         string
	RequestRetryCountExceeded                 string
	RequestsRate                              string
	ResponseBytesRate                         string
	ResponsesRate                             string
	ServerBusyErrorRate                       string
	ServiceSurgeCount                         string
	SortBy                                    string
	SortOrder                                 string
	SpilloverThreshold                        string
	State                                     string
	SurgeCount                                string
	TcpMaxOutOfOrderPackets                   string
	ToleratingTimeToLastByteTransactions      string
	ToleratingTimeToLastByteTransactionsRate  string
	TotalClientTimeToLastByteTransactions     string
	TotalConnectionReassemblyQueue75          string
	TotalConnectionReassemblyQueueFlush       string
	TotalHits                                 string
	TotalHttp2Requests                        string
	TotalHttp2Responses                       string
	TotalPacketsReceived                      string
	TotalPacketsSent                          string
	TotalRequestBytes                         string
	TotalRequests                             string
	TotalResponseBytes                        string
	TotalResponses                            string
	TotalServerBusyError                      string
	TotalSpillovers                           string
	TotalVserverDownBackupHits                string
	Type                                      string
	VserverLbHealth                           string
	VserverSurgeCount                         string
}{
	ActiveServices:                            "actsvcs",
	AverageClientTimeToLastByte:               "avgcltttlb",
	ClearStats:                                "clearstats",
	ClientResponseTimeApdex:                   "cltresponsetimeapdex",
	ClientTimeToLastByteTransactionsRate:      "cltttlbtransactionsrate",
	CpuUsagePerMille:                          "cpuusagepm",
	CurrentBackupPersistenceSessions:          "curbackuppersistencesessions",
	CurrentClientConnections:                  "curclntconnections",
	CurrentMultiPathTcpSessions:               "curmptcpsessions",
	CurrentMultipathTcpSubflowConnections:     "cursubflowconn",
	CurrentPersistenceSessions:                "curpersistencesessions",
	CurrentServerConnections:                  "cursrvrconnections",
	DeferredRequests:                          "deferredreq",
	DeferredRequestsRate:                      "deferredreqrate",
	EstablishedConnections:                    "establishedconn",
	FrustratingTimeToLastByteTransactions:     "frustratingttlbtransactions",
	FrustratingTimeToLastByteTransactionsRate: "frustratingttlbtransactionsrate",
	HitsRate:                                 "hitsrate",
	Http2RequestsRate:                        "h2requestsrate",
	Http2ResponsesRate:                       "h2responsesrate",
	HttpMaxHeaderFieldLengthExceededCount:    "httpmaxhdrfldlenpkts",
	HttpMaxHeaderSizePackets:                 "httpmaxhdrszpkts",
	InactiveServices:                         "inactsvcs",
	InvalidRequestResponse:                   "invalidrequestresponse",
	InvalidRequestResponseDropped:            "invalidrequestresponsedropped",
	LabelledConnections:                      "labelledconn",
	Name:                                     "name",
	PacketsReceivedRate:                      "pktsrecvdrate",
	PacketsSentRate:                          "pktssentrate",
	PrimaryIpAddress:                         "primaryipaddress",
	PrimaryPort:                              "primaryport",
	PushLabel:                                "pushlabel",
	RequestBytesRate:                         "requestbytesrate",
	RequestRetryCount:                        "reqretrycount",
	RequestRetryCountExceeded:                "reqretrycountexceeded",
	RequestsRate:                             "requestsrate",
	ResponseBytesRate:                        "responsebytesrate",
	ResponsesRate:                            "responsesrate",
	ServerBusyErrorRate:                      "svrbusyerrrate",
	ServiceSurgeCount:                        "svcsurgecount",
	SortBy:                                   "sortby",
	SortOrder:                                "sortorder",
	SpilloverThreshold:                       "sothreshold",
	State:                                    "state",
	SurgeCount:                               "surgecount",
	TcpMaxOutOfOrderPackets:                  "tcpmaxooopkts",
	ToleratingTimeToLastByteTransactions:     "toleratingttlbtransactions",
	ToleratingTimeToLastByteTransactionsRate: "toleratingttlbtransactionsrate",
	TotalClientTimeToLastByteTransactions:    "totcltttlbtransactions",
	TotalConnectionReassemblyQueue75:         "totalconnreassemblyqueue75",
	TotalConnectionReassemblyQueueFlush:      "totalconnreassemblyqueueflush",
	TotalHits:                                "tothits",
	TotalHttp2Requests:                       "totalh2requests",
	TotalHttp2Responses:                      "totalh2responses",
	TotalPacketsReceived:                     "totalpktsrecvd",
	TotalPacketsSent:                         "totalpktssent",
	TotalRequestBytes:                        "totalrequestbytes",
	TotalRequests:                            "totalrequests",
	TotalResponseBytes:                       "totalresponsebytes",
	TotalResponses:                           "totalresponses",
	TotalServerBusyError:                     "totalsvrbusyerr",
	TotalSpillovers:                          "totspillovers",
	TotalVserverDownBackupHits:               "totvserverdownbackuphits",
	Type:                                     "type",
	VserverLbHealth:                          "vslbhealth",
	VserverSurgeCount:                        "vsvrsurgecount",
}

type LbVserver struct {
	ActiveServices                            string  `json:"actsvcs,omitempty" nitro:"permission=readonly"`
	AverageClientTimeToLastByte               string  `json:"avgcltttlb,omitempty" nitro:"permission=readonly"`
	ClearStats                                string  `json:"clearstats,omitempty" nitro:"permission=readwrite"`
	ClientResponseTimeApdex                   float64 `json:"cltresponsetimeapdex,omitempty" nitro:"permission=readonly"`
	ClientTimeToLastByteTransactionsRate      float64 `json:"cltttlbtransactionsrate,omitempty" nitro:"permission=readonly"`
	CpuUsagePerMille                          string  `json:"cpuusagepm,omitempty" nitro:"permission=readonly"`
	CurrentBackupPersistenceSessions          string  `json:"curbackuppersistencesessions,omitempty" nitro:"permission=readonly"`
	CurrentClientConnections                  string  `json:"curclntconnections,omitempty" nitro:"permission=readonly"`
	CurrentMultiPathTcpSessions               string  `json:"curmptcpsessions,omitempty" nitro:"permission=readonly"`
	CurrentMultipathTcpSubflowConnections     string  `json:"cursubflowconn,omitempty" nitro:"permission=readonly"`
	CurrentPersistenceSessions                string  `json:"curpersistencesessions,omitempty" nitro:"permission=readonly"`
	CurrentServerConnections                  string  `json:"cursrvrconnections,omitempty" nitro:"permission=readonly"`
	DeferredRequests                          string  `json:"deferredreq,omitempty" nitro:"permission=readonly"`
	DeferredRequestsRate                      float64 `json:"deferredreqrate,omitempty" nitro:"permission=readonly"`
	EstablishedConnections                    string  `json:"establishedconn,omitempty" nitro:"permission=readonly"`
	FrustratingTimeToLastByteTransactions     string  `json:"frustratingttlbtransactions,omitempty" nitro:"permission=readonly"`
	FrustratingTimeToLastByteTransactionsRate float64 `json:"frustratingttlbtransactionsrate,omitempty" nitro:"permission=readonly"`
	HitsRate                                  float64 `json:"hitsrate,omitempty" nitro:"permission=readonly"`
	Http2RequestsRate                         float64 `json:"h2requestsrate,omitempty" nitro:"permission=readonly"`
	Http2ResponsesRate                        float64 `json:"h2responsesrate,omitempty" nitro:"permission=readonly"`
	HttpMaxHeaderFieldLengthExceededCount     string  `json:"httpmaxhdrfldlenpkts,omitempty" nitro:"permission=readonly"`
	HttpMaxHeaderSizePackets                  string  `json:"httpmaxhdrszpkts,omitempty" nitro:"permission=readonly"`
	InactiveServices                          string  `json:"inactsvcs,omitempty" nitro:"permission=readonly"`
	InvalidRequestResponse                    string  `json:"invalidrequestresponse,omitempty" nitro:"permission=readonly"`
	InvalidRequestResponseDropped             string  `json:"invalidrequestresponsedropped,omitempty" nitro:"permission=readonly"`
	LabelledConnections                       string  `json:"labelledconn,omitempty" nitro:"permission=readonly"`
	Name                                      string  `json:"name,omitempty" nitro:"permission=readwrite"`
	PacketsReceivedRate                       float64 `json:"pktsrecvdrate,omitempty" nitro:"permission=readonly"`
	PacketsSentRate                           float64 `json:"pktssentrate,omitempty" nitro:"permission=readonly"`
	PrimaryIpAddress                          string  `json:"primaryipaddress,omitempty" nitro:"permission=readonly"`
	PrimaryPort                               int     `json:"primaryport,omitempty" nitro:"permission=readonly"`
	PushLabel                                 string  `json:"pushlabel,omitempty" nitro:"permission=readonly"`
	RequestBytesRate                          float64 `json:"requestbytesrate,omitempty" nitro:"permission=readonly"`
	RequestRetryCount                         string  `json:"reqretrycount,omitempty" nitro:"permission=readonly"`
	RequestRetryCountExceeded                 string  `json:"reqretrycountexceeded,omitempty" nitro:"permission=readonly"`
	RequestsRate                              float64 `json:"requestsrate,omitempty" nitro:"permission=readonly"`
	ResponseBytesRate                         float64 `json:"responsebytesrate,omitempty" nitro:"permission=readonly"`
	ResponsesRate                             float64 `json:"responsesrate,omitempty" nitro:"permission=readonly"`
	ServerBusyErrorRate                       float64 `json:"svrbusyerrrate,omitempty" nitro:"permission=readonly"`
	ServiceSurgeCount                         string  `json:"svcsurgecount,omitempty" nitro:"permission=readonly"`
	SortBy                                    string  `json:"sortby,omitempty" nitro:"permission=readwrite"`
	SortOrder                                 string  `json:"sortorder,omitempty" nitro:"permission=readwrite"`
	SpilloverThreshold                        string  `json:"sothreshold,omitempty" nitro:"permission=readonly"`
	State                                     string  `json:"state,omitempty" nitro:"permission=readonly"`
	SurgeCount                                string  `json:"surgecount,omitempty" nitro:"permission=readonly"`
	TcpMaxOutOfOrderPackets                   string  `json:"tcpmaxooopkts,omitempty" nitro:"permission=readonly"`
	ToleratingTimeToLastByteTransactions      string  `json:"toleratingttlbtransactions,omitempty" nitro:"permission=readonly"`
	ToleratingTimeToLastByteTransactionsRate  float64 `json:"toleratingttlbtransactionsrate,omitempty" nitro:"permission=readonly"`
	TotalClientTimeToLastByteTransactions     string  `json:"totcltttlbtransactions,omitempty" nitro:"permission=readonly"`
	TotalConnectionReassemblyQueue75          string  `json:"totalconnreassemblyqueue75,omitempty" nitro:"permission=readonly"`
	TotalConnectionReassemblyQueueFlush       string  `json:"totalconnreassemblyqueueflush,omitempty" nitro:"permission=readonly"`
	TotalHits                                 string  `json:"tothits,omitempty" nitro:"permission=readonly"`
	TotalHttp2Requests                        string  `json:"totalh2requests,omitempty" nitro:"permission=readonly"`
	TotalHttp2Responses                       string  `json:"totalh2responses,omitempty" nitro:"permission=readonly"`
	TotalPacketsReceived                      string  `json:"totalpktsrecvd,omitempty" nitro:"permission=readonly"`
	TotalPacketsSent                          string  `json:"totalpktssent,omitempty" nitro:"permission=readonly"`
	TotalRequestBytes                         string  `json:"totalrequestbytes,omitempty" nitro:"permission=readonly"`
	TotalRequests                             string  `json:"totalrequests,omitempty" nitro:"permission=readonly"`
	TotalResponseBytes                        string  `json:"totalresponsebytes,omitempty" nitro:"permission=readonly"`
	TotalResponses                            string  `json:"totalresponses,omitempty" nitro:"permission=readonly"`
	TotalServerBusyError                      string  `json:"totalsvrbusyerr,omitempty" nitro:"permission=readonly"`
	TotalSpillovers                           string  `json:"totspillovers,omitempty" nitro:"permission=readonly"`
	TotalVserverDownBackupHits                string  `json:"totvserverdownbackuphits,omitempty" nitro:"permission=readonly"`
	Type                                      string  `json:"type,omitempty" nitro:"permission=readonly"`
	VserverLbHealth                           string  `json:"vslbhealth,omitempty" nitro:"permission=readonly"`
	VserverSurgeCount                         string  `json:"vsvrsurgecount,omitempty" nitro:"permission=readonly"`
}

func (r LbVserver) GetTypeName() string {
	return "lbvserver"
}
