/*
 * Copyright 2023 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package nitro

type errorCode float64

const (
	NSERR_RNAT_INV         errorCode = 769
	NSERR_RNAT_INV_MESSAGE string    = "Reverse NAT not applicable for default route."

	NSERR_INVALID_IF         errorCode = 770
	NSERR_INVALID_IF_MESSAGE string    = "Invalid errorCodeerface name/number."

	NSERR_MGRLIMIT_REACHED         errorCode = 771
	NSERR_MGRLIMIT_REACHED_MESSAGE string    = "Maximum manager limit reached."

	NSERR_SP_INVALD_TABLE         errorCode = 772
	NSERR_SP_INVALD_TABLE_MESSAGE string    = "SP table entries should be in increasing order."

	NSERR_RNAT_NATIP_EXISTS         errorCode = 773
	NSERR_RNAT_NATIP_EXISTS_MESSAGE string    = "RNAT to the target network with specified NAT IP already exists."

	NSERR_RNAT_EXISTS         errorCode = 774
	NSERR_RNAT_EXISTS_MESSAGE string    = "ACL or NETWORK based RNAT rule already exists."

	NSERR_RNAT_NOT_EXISTS         errorCode = 775
	NSERR_RNAT_NOT_EXISTS_MESSAGE string    = "ACL or NETWORK based RNAT rule does not exist."

	NSERR_RNAT_NATIP_NOT_EXISTS         errorCode = 776
	NSERR_RNAT_NATIP_NOT_EXISTS_MESSAGE string    = "RNAT to the target network with specified NAT IP doesn't exist."

	NSERR_RNAT_INVALID_NATIP         errorCode = 777
	NSERR_RNAT_INVALID_NATIP_MESSAGE string    = "NAT IP is not valid."

	NSERR_RNAT_XACLWITHTTL         errorCode = 778
	NSERR_RNAT_XACLWITHTTL_MESSAGE string    = "Acl with ttl can not be used in NAT64/RNAT/RNAT6/ForwardingSession/LSN rule."

	NSERR_ARP_DISABLED         errorCode = 784
	NSERR_ARP_DISABLED_MESSAGE string    = "IP has arp disabled."

	NSERR_ARP_SEC_NOT_OWNEDIP         errorCode = 785
	NSERR_ARP_SEC_NOT_OWNEDIP_MESSAGE string    = "Secondary can not arp for this IP."

	NSERR_CPE_RULE_INVAL         errorCode = 786
	NSERR_CPE_RULE_INVAL_MESSAGE string    = "Invalid rule."

	NSERR_INVAL_FLOWTYPE         errorCode = 787
	NSERR_INVAL_FLOWTYPE_MESSAGE string    = "Only authorization,  audit,  VPN session and traffic policies can be bound to aaa user or group."

	NSERR_INVAL_POLICY_TYPE         errorCode = 788
	NSERR_INVAL_POLICY_TYPE_MESSAGE string    = "Response rule is invalid in an authorization policy."

	NSERR_CPE_RULE_ACTION_INVAL         errorCode = 789
	NSERR_CPE_RULE_ACTION_INVAL_MESSAGE string    = "Request action is valid only for request rule."

	NSERR_CPE_DEF_SET_INVAL         errorCode = 790
	NSERR_CPE_DEF_SET_INVAL_MESSAGE string    = "Default policy cannot be set."

	NSERR_INVAL_FORCECLEANUP         errorCode = 791
	NSERR_INVAL_FORCECLEANUP_MESSAGE string    = "Invalid forcecleanup value."

	NSERR_INVAL_AAA_GROUP         errorCode = 792
	NSERR_INVAL_AAA_GROUP_MESSAGE string    = "Invalid authorizationgroup value."

	NSERR_INVAL_PROXY         errorCode = 793
	NSERR_INVAL_PROXY_MESSAGE string    = "Invalid allprotocolproxy value."

	NSERR_INVAL_HTTTPPROXY         errorCode = 800
	NSERR_INVAL_HTTTPPROXY_MESSAGE string    = "Invalid HTTP proxy value."

	NSERR_INVAL_FTPPROXY         errorCode = 801
	NSERR_INVAL_FTPPROXY_MESSAGE string    = "Invalid FTP proxy value."

	NSERR_INVAL_SOCKPROXY         errorCode = 802
	NSERR_INVAL_SOCKPROXY_MESSAGE string    = "Invalid SOCKS proxy value."

	NSERR_INVAL_GOPHERPROXY         errorCode = 803
	NSERR_INVAL_GOPHERPROXY_MESSAGE string    = "Invalid GOPHER proxy value."

	NSERR_INVAL_SSLPROXY         errorCode = 804
	NSERR_INVAL_SSLPROXY_MESSAGE string    = "Invalid SSL proxy value."

	NSERR_INVAL_AAAGRP_MAX         errorCode = 805
	NSERR_INVAL_AAAGRP_MAX_MESSAGE string    = "Max 5 groups can be specified in authorizationgroup."

	NSERR_INVAL_MAX_PORT_NUM         errorCode = 806
	NSERR_INVAL_MAX_PORT_NUM_MESSAGE string    = "Maximum 16 ports can be specified in httpport."

	NSERR_INVAL_HTTPPORT         errorCode = 807
	NSERR_INVAL_HTTPPORT_MESSAGE string    = "Invalid port."

	NSERR_INVAL_VPNVSERER_POLTYPE         errorCode = 808
	NSERR_INVAL_VPNVSERER_POLTYPE_MESSAGE string    = "This type of policy cannot be bound to VPN Vserver."

	NSERR_INVAL_VPNGLOBAL_POLTYPE         errorCode = 809
	NSERR_INVAL_VPNGLOBAL_POLTYPE_MESSAGE string    = "This type of policy cannot be bound to VPN Global."

	NSERR_CPE_REM_INUSE         errorCode = 810
	NSERR_CPE_REM_INUSE_MESSAGE string    = "Bound policy cannot be removed."

	NSERR_PROXY_CONFLICT         errorCode = 811
	NSERR_PROXY_CONFLICT_MESSAGE string    = "Proxy server for all protocols already configured."

	NSERR_PROXY_INVAL         errorCode = 812
	NSERR_PROXY_INVAL_MESSAGE string    = "Domain names allowed only if proxy type is browser"

	NSERR_PXYEXCPT_INVAL         errorCode = 813
	NSERR_PXYEXCPT_INVAL_MESSAGE string    = "Proxy exception allowed only if proxy type is browser"

	NSERR_CPE_POLTYPE_NO_CSE         errorCode = 814
	NSERR_CPE_POLTYPE_NO_CSE_MESSAGE string    = "Policy type does not support client security expressions in rule"

	NSERR_SESSACT_CSE_INCOMPATIBLE         errorCode = 815
	NSERR_SESSACT_CSE_INCOMPATIBLE_MESSAGE string    = "Session action and rule are incompatible"

	NSERR_NOMEM_CSE         errorCode = 816
	NSERR_NOMEM_CSE_MESSAGE string    = "Not enough memory while adding client security expressions"

	NSERR_INCOMPAT_FS_RULE         errorCode = 817
	NSERR_INCOMPAT_FS_RULE_MESSAGE string    = "File system expressions supported in authorization policy only"

	NSERR_INCOMPAT_FS_MIX         errorCode = 818
	NSERR_INCOMPAT_FS_MIX_MESSAGE string    = "Incompatible expressions mixed with file system expressions in rule"

	NSERR_DR_ENABLE         errorCode = 819
	NSERR_DR_ENABLE_MESSAGE string    = "Dynamic routing can be enabled on only one IP per subnet"

	NSERR_MAX_DISTANCE         errorCode = 820
	NSERR_MAX_DISTANCE_MESSAGE string    = "Only null errorCodeerface routes can have distance equal to 255"

	NSERR_NULL_ROUTE_DISTANCE         errorCode = 821
	NSERR_NULL_ROUTE_DISTANCE_MESSAGE string    = "It is not possible to set the administrative distance/cost metric for a null errorCodeerface route"

	NSERR_BAD_ACTION_TCP_PROFILE_TYPE         errorCode = 822
	NSERR_BAD_ACTION_TCP_PROFILE_TYPE_MESSAGE string    = "TCP profile cannot be set to this service type"

	NSERR_BAD_ACTION_HTTP_PROFILE_TYPE         errorCode = 823
	NSERR_BAD_ACTION_HTTP_PROFILE_TYPE_MESSAGE string    = "HTTP profile cannot be set to this service type"

	NSERR_SP_INVALID_THRESHOLD         errorCode = 824
	NSERR_SP_INVALID_THRESHOLD_MESSAGE string    = "Invalid base threshold value"

	NSERR_VIP_ROUTE_EXISTS         errorCode = 825
	NSERR_VIP_ROUTE_EXISTS_MESSAGE string    = "VIP exists for this host route"

	NSERR_ACTION_BOUND         errorCode = 832
	NSERR_ACTION_BOUND_MESSAGE string    = "Action bound to policy can not be deleted."

	NSERR_MAX_IPLIMIT         errorCode = 833
	NSERR_MAX_IPLIMIT_MESSAGE string    = "Maximum limit for bound IP to this resource record reached,  remaining IP's destined for this resource record will be discarded."

	NSERR_BAD_PREFIX_LEN         errorCode = 826
	NSERR_BAD_PREFIX_LEN_MESSAGE string    = "Invalid rnat/lsn network prefix len it must be 0-128"

	NSERR_INV_NETADDR         errorCode = 827
	NSERR_INV_NETADDR_MESSAGE string    = "Invalid IPv6 network address"

	NSERR_MIN_DISTANCE         errorCode = 828
	NSERR_MIN_DISTANCE_MESSAGE string    = "Static route cannot have distance less than 1"

	NSERR_IPSEC_PROFILE_INUSE         errorCode = 829
	NSERR_IPSEC_PROFILE_INUSE_MESSAGE string    = "Profile is in use by an ip tunnel."

	NSERR_IPSEC_INVALID_PROFILE_NAME         errorCode = 830
	NSERR_IPSEC_INVALID_PROFILE_NAME_MESSAGE string    = "Invalid profile name."

	NSERR_SPOTTED         errorCode = 831
	NSERR_SPOTTED_MESSAGE string    = "Spotted IP cannot be used here"

	NSERR_IPSEC_PBR_ONLY_IP_SUPPORTED         errorCode = 834
	NSERR_IPSEC_PBR_ONLY_IP_SUPPORTED_MESSAGE string    = "Only IP range is supported"

	NSERR_IPSEC_IP_NOT_CONTIGUOUS         errorCode = 835
	NSERR_IPSEC_IP_NOT_CONTIGUOUS_MESSAGE string    = "Invalid IP range. The IP range does not cover all the IP's for a given subnet mask"

	NSERR_INVAL_USER_ACCOUNTING         errorCode = 836
	NSERR_INVAL_USER_ACCOUNTING_MESSAGE string    = "Invalid userAccounting value.  userAccounting value must be a radiusPolicy."

	NSERR_RDP_PROFILE_INUSE         errorCode = 837
	NSERR_RDP_PROFILE_INUSE_MESSAGE string    = "Profile is in use."

	NSERR_RDP_INVALID_PROFILE_NAME         errorCode = 838
	NSERR_RDP_INVALID_PROFILE_NAME_MESSAGE string    = "Invalid profile name."

	NSERR_RDP_INVALID_RDP_FILE_NAME         errorCode = 839
	NSERR_RDP_INVALID_RDP_FILE_NAME_MESSAGE string    = "RDP file name should have .rdp extension and some special characters for RDP file name are not supported."

	NSERR_RDP_INVALID_SERVERPROFILE_NAME         errorCode = 840
	NSERR_RDP_INVALID_SERVERPROFILE_NAME_MESSAGE string    = "Invalid server profile name."

	NSERR_RDP_SERVERPROFILE_ALREADY_IN_USE         errorCode = 841
	NSERR_RDP_SERVERPROFILE_ALREADY_IN_USE_MESSAGE string    = "Server profile is already in use."

	NSERR_AUTO_PROXY_URL_INVALID         errorCode = 842
	NSERR_AUTO_PROXY_URL_INVALID_MESSAGE string    = "Invalid Autoproxy URL"

	NSERR_AUTO_PROXY_NO_BROWSER         errorCode = 843
	NSERR_AUTO_PROXY_NO_BROWSER_MESSAGE string    = "Autoproxy URL is allowed only if proxy type is Browser."

	NSERR_AUTO_PROXY_CONFLICT         errorCode = 844
	NSERR_AUTO_PROXY_CONFLICT_MESSAGE string    = "Both Autoproxy URL and manual proxy parameters (allprotocolproxy,  http proxy,  ftp proxy,  proxy exception list,  ...) cannot be configured at same time."

	NSERR_RDP_INVALID_HOST_LEN         errorCode = 848
	NSERR_RDP_INVALID_HOST_LEN_MESSAGE string    = "Invalid Host length. Atleast 3 characters expected"

	NSERR_RDP_NEED_IP         errorCode = 849
	NSERR_RDP_NEED_IP_MESSAGE string    = "RDP IP is needed"

	NSERR_IPSEC_INV_PROF         errorCode = 850
	NSERR_IPSEC_INV_PROF_MESSAGE string    = "For IPSEC tunnels,  profile name cannot be 'none'"

	NSERR_RDP_NEED_PSK         errorCode = 854
	NSERR_RDP_NEED_PSK_MESSAGE string    = "PSK is needed"

	NSERR_IPSEC_ONLY_V1         errorCode = 858
	NSERR_IPSEC_ONLY_V1_MESSAGE string    = "Responder Only config is supported with IKEv1"

	NSERR_RDP_LISTENER_SPECIFY_PORT         errorCode = 859
	NSERR_RDP_LISTENER_SPECIFY_PORT_MESSAGE string    = "Invalid Listener. Specify Port"

	NSERR_RDP_INVALID_LISTENER_HOST         errorCode = 860
	NSERR_RDP_INVALID_LISTENER_HOST_MESSAGE string    = "Invalid Listener. Specify valid host,  atleast 3 characters expected"

	NSERR_RDP_LISTENER_INVALID_PORT         errorCode = 861
	NSERR_RDP_LISTENER_INVALID_PORT_MESSAGE string    = "Invalid Listener. Specify valid Port"

	NSERR_PROXY_NOT_ALLOWED         errorCode = 862
	NSERR_PROXY_NOT_ALLOWED_MESSAGE string    = "Proxy feature not licensed"

	NSERR_HOST_RT_GW_NOT_SET         errorCode = 3806
	NSERR_HOST_RT_GW_NOT_SET_MESSAGE string    = "Host Route Gateway not set"
)
