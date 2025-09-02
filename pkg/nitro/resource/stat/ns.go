package stat

var NsFieldNames = struct {
	AppFirewallAbortsRate                string
	AppFirewallAbortsTotal               string
	AppFirewallRedirectsRate             string
	AppFirewallRedirectsTotal            string
	AppFirewallRequestRate               string
	AppFirewallRequestTotal              string
	AppFirewallResponseRate              string
	AppFirewallResponseTotal             string
	AverageCpuUsage                      string
	AverageCpuUsagePercentage            string
	BandwidthLimitInboundExceededRate    string
	BandwidthLimitInboundExceededTotal   string
	BandwidthLimitOutboundExceededRate   string
	BandwidthLimitOutboundExceededTotal  string
	Cache64MaxMemoryKbs                  string
	CacheBandwidthSavedPercentage        string
	CacheHttpHits                        string
	CacheHttpHitsRate                    string
	CacheHttpMisses                      string
	CacheHttpMissesRate                  string
	CacheMaxMemoryActiveKbs              string
	CacheMaxMemoryKbs                    string
	CacheMemoryUsageKbs                  string
	ClearStats                           string
	CompressionRatio                     string
	CompressionRatioServerClient         string
	ConnectionTrackingLimitExceededRate  string
	ConnectionTrackingLimitExceededTotal string
	CpuNumberTotal                       string
	CpuUsage                             string
	CpuUtilizationPercentage             string
	Disk0AvailableSpace                  string
	Disk0UsagePercentage                 string
	Disk1AvailableSpace                  string
	Disk1UsagePercentage                 string
	HighAvailabilityCurrentState         string
	HighAvailabilityMasterState          string
	HttpCompressionRatio                 string
	HttpRequestRate                      string
	HttpRequestReceivedBytesRate         string
	HttpRequestReceivedBytesTotal        string
	HttpRequestTotal                     string
	HttpResponseRate                     string
	HttpResponseReceivedBytesRate        string
	HttpResponseReceivedBytesTotal       string
	HttpResponseTotal                    string
	LastMasterStateTransitionTime        string
	LinkLocalLimitExceededRate           string
	LinkLocalLimitExceededTotal          string
	ManagementCpuAverageUsagePercentage  string
	MemoryUsage                          string
	MemoryUsageMegaBytes                 string
	MemoryUsagePercentage                string
	MiscellaneousCounter0                string
	MiscellaneousCounter1                string
	PacketCpuAverageUsagePercentage      string
	PacketsPerSecondLimitExceededRate    string
	PacketsPerSecondLimitExceededTotal   string
	ReceivedMbitsRate                    string
	ReceivedMegabitsTotal                string
	SslCardsTotal                        string
	SslCardsUp                           string
	SslSessionReuseHits                  string
	SslSessionReuseRate                  string
	SslTransactionsRate                  string
	SslTransactionsTotal                 string
	StartTime                            string
	StartTimeLocalTimezone               string
	TcpClientConnectionsCurrent          string
	TcpClientConnectionsEstablished      string
	TcpServerConnectionsCurrent          string
	TcpServerConnectionsEstablished      string
	TransmittedMegabitsRate              string
	TransmittedMegabitsTotal             string
}{
	AppFirewallAbortsRate:                "appfirewallabortsrate",
	AppFirewallAbortsTotal:               "appfirewallaborts",
	AppFirewallRedirectsRate:             "appfirewallredirectsrate",
	AppFirewallRedirectsTotal:            "appfirewallredirects",
	AppFirewallRequestRate:               "appfirewallrequestsrate",
	AppFirewallRequestTotal:              "appfirewallrequests",
	AppFirewallResponseRate:              "appfirewallresponsesrate",
	AppFirewallResponseTotal:             "appfirewallresponses",
	AverageCpuUsage:                      "rescpuusage",
	AverageCpuUsagePercentage:            "rescpuusagepcnt",
	BandwidthLimitInboundExceededRate:    "enainbwlimitexceededrate",
	BandwidthLimitInboundExceededTotal:   "enainbwlimitexceeded",
	BandwidthLimitOutboundExceededRate:   "enaoutbwlimitexceededrate",
	BandwidthLimitOutboundExceededTotal:  "enaoutbwlimitexceeded",
	Cache64MaxMemoryKbs:                  "cache64Maxmemorykb",
	CacheBandwidthSavedPercentage:        "cachepercentoriginbandwidthsaved",
	CacheHttpHits:                        "cachetothits",
	CacheHttpHitsRate:                    "cachehitsrate",
	CacheHttpMisses:                      "cachetotmisses",
	CacheHttpMissesRate:                  "cachemissesrate",
	CacheMaxMemoryActiveKbs:              "cachemaxmemoryactivekb",
	CacheMaxMemoryKbs:                    "cachemaxmemorykb",
	CacheMemoryUsageKbs:                  "cacheutilizedmemorykb",
	ClearStats:                           "clearstats",
	CompressionRatio:                     "delcmpratio",
	CompressionRatioServerClient:         "compratio",
	ConnectionTrackingLimitExceededRate:  "enaconntracklimitexceededrate",
	ConnectionTrackingLimitExceededTotal: "enaconntracklimitexceeded",
	CpuNumberTotal:                       "numcpus",
	CpuUsage:                             "cpuusage",
	CpuUtilizationPercentage:             "cpuusagepcnt",
	Disk0AvailableSpace:                  "disk0Avail",
	Disk0UsagePercentage:                 "disk0Perusage",
	Disk1AvailableSpace:                  "disk1Avail",
	Disk1UsagePercentage:                 "disk1Perusage",
	HighAvailabilityCurrentState:         "hacurstate",
	HighAvailabilityMasterState:          "hacurmasterstate",
	HttpCompressionRatio:                 "comptotaldatacompressionratio",
	HttpRequestRate:                      "httprequestsrate",
	HttpRequestReceivedBytesRate:         "httprxrequestbytesrate",
	HttpRequestReceivedBytesTotal:        "httptotrxrequestbytes",
	HttpRequestTotal:                     "httptotrequests",
	HttpResponseRate:                     "httpresponsesrate",
	HttpResponseReceivedBytesRate:        "httprxresponsebytesrate",
	HttpResponseReceivedBytesTotal:       "httptotrxresponsebytes",
	HttpResponseTotal:                    "httptotresponses",
	LastMasterStateTransitionTime:        "transtime",
	LinkLocalLimitExceededRate:           "enalinklocallimitexceededrate",
	LinkLocalLimitExceededTotal:          "enalinklocallimitexceeded",
	ManagementCpuAverageUsagePercentage:  "mgmtcpuusagepcnt",
	MemoryUsage:                          "resmemusage",
	MemoryUsageMegaBytes:                 "memuseinmb",
	MemoryUsagePercentage:                "memusagepcnt",
	MiscellaneousCounter0:                "misccounter0",
	MiscellaneousCounter1:                "misccounter1",
	PacketCpuAverageUsagePercentage:      "pktcpuusagepcnt",
	PacketsPerSecondLimitExceededRate:    "enappslimitexceededrate",
	PacketsPerSecondLimitExceededTotal:   "enappslimitexceeded",
	ReceivedMbitsRate:                    "rxmbitsrate",
	ReceivedMegabitsTotal:                "totrxmbits",
	SslCardsTotal:                        "sslcards",
	SslCardsUp:                           "sslnumcardsup",
	SslSessionReuseHits:                  "ssltotsessionhits",
	SslSessionReuseRate:                  "sslsessionhitsrate",
	SslTransactionsRate:                  "ssltransactionsrate",
	SslTransactionsTotal:                 "ssltottransactions",
	StartTime:                            "starttime",
	StartTimeLocalTimezone:               "starttimelocal",
	TcpClientConnectionsCurrent:          "tcpcurclientconn",
	TcpClientConnectionsEstablished:      "tcpcurclientconnestablished",
	TcpServerConnectionsCurrent:          "tcpcurserverconn",
	TcpServerConnectionsEstablished:      "tcpcurserverconnestablished",
	TransmittedMegabitsRate:              "txmbitsrate",
	TransmittedMegabitsTotal:             "tottxmbits",
}

type Ns struct {
	AppFirewallAbortsRate                float64 `json:"appfirewallabortsrate,omitempty" nitro:"permission=readonly"`
	AppFirewallAbortsTotal               string  `json:"appfirewallaborts,omitempty" nitro:"permission=readonly"`
	AppFirewallRedirectsRate             float64 `json:"appfirewallredirectsrate,omitempty" nitro:"permission=readonly"`
	AppFirewallRedirectsTotal            string  `json:"appfirewallredirects,omitempty" nitro:"permission=readonly"`
	AppFirewallRequestRate               float64 `json:"appfirewallrequestsrate,omitempty" nitro:"permission=readonly"`
	AppFirewallRequestTotal              string  `json:"appfirewallrequests,omitempty" nitro:"permission=readonly"`
	AppFirewallResponseRate              float64 `json:"appfirewallresponsesrate,omitempty" nitro:"permission=readonly"`
	AppFirewallResponseTotal             string  `json:"appfirewallresponses,omitempty" nitro:"permission=readonly"`
	CpuAverageUsage                      string  `json:"rescpuusage,omitempty" nitro:"permission=readonly"`
	CpuAverageUsagePercentage            float64 `json:"rescpuusagepcnt,omitempty" nitro:"permission=readonly"`
	BandwidthLimitInboundExceededRate    float64 `json:"enainbwlimitexceededrate,omitempty" nitro:"permission=readonly"`
	BandwidthLimitInboundExceededTotal   string  `json:"enainbwlimitexceeded,omitempty" nitro:"permission=readonly"`
	BandwidthLimitOutboundExceededRate   float64 `json:"enaoutbwlimitexceededrate,omitempty" nitro:"permission=readonly"`
	BandwidthLimitOutboundExceededTotal  string  `json:"enaoutbwlimitexceeded,omitempty" nitro:"permission=readonly"`
	Cache64MaxMemoryKbs                  string  `json:"cache64Maxmemorykb,omitempty" nitro:"permission=readonly"`
	CacheBandwidthSavedPercentage        string  `json:"cachepercentoriginbandwidthsaved,omitempty" nitro:"permission=readonly"`
	CacheHttpHits                        string  `json:"cachetothits,omitempty" nitro:"permission=readonly"`
	CacheHttpHitsRate                    float64 `json:"cachehitsrate,omitempty" nitro:"permission=readonly"`
	CacheHttpMisses                      string  `json:"cachetotmisses,omitempty" nitro:"permission=readonly"`
	CacheHttpMissesRate                  float64 `json:"cachemissesrate,omitempty" nitro:"permission=readonly"`
	CacheMaxMemoryActiveKbs              string  `json:"cachemaxmemoryactivekb,omitempty" nitro:"permission=readonly"`
	CacheMaxMemoryKbs                    string  `json:"cachemaxmemorykb,omitempty" nitro:"permission=readonly"`
	CacheMemoryUsageKbs                  string  `json:"cacheutilizedmemorykb,omitempty" nitro:"permission=readonly"`
	ClearStats                           string  `json:"clearstats,omitempty" nitro:"permission=readwrite"`
	CompressionRatio                     float64 `json:"delcmpratio,omitempty" nitro:"permission=readonly"`
	CompressionRatioServerClient         float64 `json:"compratio,omitempty" nitro:"permission=readonly"`
	ConnectionTrackingLimitExceededRate  float64 `json:"enaconntracklimitexceededrate,omitempty" nitro:"permission=readonly"`
	ConnectionTrackingLimitExceededTotal string  `json:"enaconntracklimitexceeded,omitempty" nitro:"permission=readonly"`
	CpuNumberTotal                       string  `json:"numcpus,omitempty" nitro:"permission=readonly"`
	CpuUsage                             string  `json:"cpuusage,omitempty" nitro:"permission=readonly"`
	CpuUtilizationPercentage             float64 `json:"cpuusagepcnt,omitempty" nitro:"permission=readonly"`
	Disk0AvailableSpace                  float64 `json:"disk0Avail,omitempty" nitro:"permission=readonly"`
	Disk0UsagePercentage                 float64 `json:"disk0Perusage,omitempty" nitro:"permission=readonly"`
	Disk1AvailableSpace                  float64 `json:"disk1Avail,omitempty" nitro:"permission=readonly"`
	Disk1UsagePercentage                 float64 `json:"disk1Perusage,omitempty" nitro:"permission=readonly"`
	HighAvailabilityCurrentState         string  `json:"hacurstate,omitempty" nitro:"permission=readonly"`
	HighAvailabilityMasterState          string  `json:"hacurmasterstate,omitempty" nitro:"permission=readonly"`
	HttpCompressionRatio                 float64 `json:"comptotaldatacompressionratio,omitempty" nitro:"permission=readonly"`
	HttpRequestRate                      float64 `json:"httprequestsrate,omitempty" nitro:"permission=readonly"`
	HttpRequestReceivedBytesRate         float64 `json:"httprxrequestbytesrate,omitempty" nitro:"permission=readonly"`
	HttpRequestReceivedBytesTotal        string  `json:"httptotrxrequestbytes,omitempty" nitro:"permission=readonly"`
	HttpRequestTotal                     string  `json:"httptotrequests,omitempty" nitro:"permission=readonly"`
	HttpResponseRate                     float64 `json:"httpresponsesrate,omitempty" nitro:"permission=readonly"`
	HttpResponseReceivedBytesRate        float64 `json:"httprxresponsebytesrate,omitempty" nitro:"permission=readonly"`
	HttpResponseReceivedBytesTotal       string  `json:"httptotrxresponsebytes,omitempty" nitro:"permission=readonly"`
	HttpResponseTotal                    string  `json:"httptotresponses,omitempty" nitro:"permission=readonly"`
	LastMasterStateTransitionTime        string  `json:"transtime,omitempty" nitro:"permission=readonly"`
	LinkLocalLimitExceededRate           float64 `json:"enalinklocallimitexceededrate,omitempty" nitro:"permission=readonly"`
	LinkLocalLimitExceededTotal          string  `json:"enalinklocallimitexceeded,omitempty" nitro:"permission=readonly"`
	ManagementCpuAverageUsagePercentage  float64 `json:"mgmtcpuusagepcnt,omitempty" nitro:"permission=readonly"`
	MemoryUsage                          string  `json:"resmemusage,omitempty" nitro:"permission=readonly"`
	MemoryUsageMegaBytes                 string  `json:"memuseinmb,omitempty" nitro:"permission=readonly"`
	MemoryUsagePercentage                float64 `json:"memusagepcnt,omitempty" nitro:"permission=readonly"`
	MiscellaneousCounter0                float64 `json:"misccounter0,omitempty" nitro:"permission=readonly"`
	MiscellaneousCounter1                float64 `json:"misccounter1,omitempty" nitro:"permission=readonly"`
	PacketCpuAverageUsagePercentage      float64 `json:"pktcpuusagepcnt,omitempty" nitro:"permission=readonly"`
	PacketsPerSecondLimitExceededRate    float64 `json:"enappslimitexceededrate,omitempty" nitro:"permission=readonly"`
	PacketsPerSecondLimitExceededTotal   string  `json:"enappslimitexceeded,omitempty" nitro:"permission=readonly"`
	ReceivedMegabitsRate                 float64 `json:"rxmbitsrate,omitempty" nitro:"permission=readonly"`
	ReceivedMegabitsTotal                string  `json:"totrxmbits,omitempty" nitro:"permission=readonly"`
	SslCardsTotal                        string  `json:"sslcards,omitempty" nitro:"permission=readonly"`
	SslCardsUp                           string  `json:"sslnumcardsup,omitempty" nitro:"permission=readonly"`
	SslSessionReuseHits                  string  `json:"ssltotsessionhits,omitempty" nitro:"permission=readonly"`
	SslSessionReuseRate                  float64 `json:"sslsessionhitsrate,omitempty" nitro:"permission=readonly"`
	SslTransactionsRate                  float64 `json:"ssltransactionsrate,omitempty" nitro:"permission=readonly"`
	SslTransactionsTotal                 string  `json:"ssltottransactions,omitempty" nitro:"permission=readonly"`
	StartTime                            string  `json:"starttime,omitempty" nitro:"permission=readonly"`
	StartTimeLocalTimezone               string  `json:"starttimelocal,omitempty" nitro:"permission=readonly"`
	TcpClientConnectionsCurrent          string  `json:"tcpcurclientconn,omitempty" nitro:"permission=readonly"`
	TcpClientConnectionsEstablished      string  `json:"tcpcurclientconnestablished,omitempty" nitro:"permission=readonly"`
	TcpServerConnectionsCurrent          string  `json:"tcpcurserverconn,omitempty" nitro:"permission=readonly"`
	TcpServerConnectionsEstablished      string  `json:"tcpcurserverconnestablished,omitempty" nitro:"permission=readonly"`
	TransmittedMegabitsRate              float64 `json:"txmbitsrate,omitempty" nitro:"permission=readonly"`
	TransmittedMegabitsTotal             string  `json:"tottxmbits,omitempty" nitro:"permission=readonly"`
}

func (r Ns) GetTypeName() string {
	return "ns"
}
