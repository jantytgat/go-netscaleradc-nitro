package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro"
	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

var (
	NITRO_ENV_NAME     string
	NITRO_ENV_ADDRESS  string
	NITRO_ENV_USERNAME string
	NITRO_ENV_PASSWORD string
)

func main() {
	var err error
	var file *os.File

	if file, err = os.Open("test.env"); err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "=")
		switch values[0] {
		case "NITRO_ENV_NAME":
			NITRO_ENV_NAME = values[1]
			fmt.Println("NITRO_ENV_NAME:", NITRO_ENV_NAME)
		case "NITRO_ENV_ADDRESS":
			NITRO_ENV_ADDRESS = values[1]
			fmt.Println("NITRO_ENV_ADDRESS:", NITRO_ENV_ADDRESS)
		case "NITRO_ENV_USERNAME":
			NITRO_ENV_USERNAME = values[1]
			fmt.Println("NITRO_ENV_USERNAME:", NITRO_ENV_USERNAME)
		case "NITRO_ENV_PASSWORD":
			NITRO_ENV_PASSWORD = values[1]
			if NITRO_ENV_PASSWORD == "" {
				panic(errors.New("NITRO_ENV_PASSWORD is empty"))
			}
		default:
			panic(values)
		}
		line++
	}
	var client *nitro.Client

	if client, err = nitro.NewClient(
		NITRO_ENV_NAME,
		NITRO_ENV_ADDRESS,
		nitro.Credentials{
			Username: NITRO_ENV_USERNAME,
			Password: NITRO_ENV_PASSWORD,
		},
		nitro.ConnectionSettings{
			UseSsl:                    true,
			Timeout:                   0,
			UserAgent:                 "dev-nitro",
			ValidateServerCertificate: false,
			LogTlsSecrets:             false,
			LogTlsSecretsDestination:  "",
			AutoLogin:                 true,
		},
		nitro.StrictSerializationMode); err != nil {
		panic(errors.New("error creating client: " + err.Error()))
	}

	defer func() {
		if client.IsLoggedIn() {
			if err = client.Logout(); err != nil {
				panic(err)
			}
		}
	}()

	fmt.Println(client.BaseUrl())
	ctx := context.Background()

	// Run commands against a live system
	// CsVserver
	var csvs []config.CsVserver
	if csvs, err = client.CsVserver.List(ctx, nil, nil); err != nil {
		panic(err)
	}
	var csvs_count float64
	if csvs_count, err = client.CsVserver.Count(ctx); err != nil {
		panic(err)
	}
	if len(csvs) != int(csvs_count) {
		panic("number of csv. expected " + fmt.Sprint(len(csvs)))
	}
	for _, csv := range csvs {
		if _, err = client.CsVserver.Stats(ctx, csv.Name, nil); err != nil {
			panic(err)
		}
	}

	// DnsAddRec
	var dnsAddRecs []config.DnsAddressRecord
	if dnsAddRecs, err = client.DnsAddressRecord.List(ctx, nil, nil); err != nil {
		panic(err)
	}
	var dnsAddRecs_count float64
	if dnsAddRecs_count, err = client.DnsAddressRecord.Count(ctx); err != nil {
		panic(err)
	}
	if len(dnsAddRecs) != int(dnsAddRecs_count) {
		panic("number of csv. expected " + fmt.Sprint(len(dnsAddRecs)))
	}
	// for _, dnsAddRec := range dnsAddRecs {
	// 	fmt.Println(dnsAddRec)
	// }

	// DnsTxtRec
	var dnsTxtRecs []config.DnsAddressRecord
	if dnsTxtRecs, err = client.DnsAddressRecord.List(ctx, nil, nil); err != nil {
		panic(err)
	}
	var dnsTxtRecs_count float64
	if dnsTxtRecs_count, err = client.DnsAddressRecord.Count(ctx); err != nil {
		panic(err)
	}
	if len(dnsTxtRecs) != int(dnsTxtRecs_count) {
		panic("number of csv. expected " + fmt.Sprint(len(dnsTxtRecs)))
	}
	// for _, dnsTxtRec := range dnsTxtRecs {
	// 	fmt.Println(dnsTxtRec)
	// }

	// HaNode
	var hanodes []config.HaNode
	if hanodes, err = client.HaNode.List(ctx, nil, nil); err != nil {
		panic(err)
	}
	var hanode_count float64
	if hanode_count, err = client.HaNode.Count(ctx); err != nil {
		panic(err)
	}
	if len(hanodes) != int(hanode_count) {
		panic("number of hanodes expected " + fmt.Sprint(len(hanodes)))
	}
	for range hanodes {
		if _, err = client.HaNode.Stats(ctx, nil); err != nil {
			panic(err)
		}
	}
	// var primary bool
	if _, err = client.IsPrimaryNode(ctx); err != nil {
		panic(err)
	}

	// lbVserver
	var lbvs []config.LbVserver

	if lbvs, err = client.LbVserver.List(ctx, nil, nil); err != nil {
		panic(err)
	}
	var lbvs_count float64
	if lbvs_count, err = client.LbVserver.Count(ctx); err != nil {
		panic(err)
	}
	if len(lbvs) != int(lbvs_count) {
		panic("number of lbvs expected " + fmt.Sprint(len(lbvs)))
	}
	for _, lbv := range lbvs {
		if _, err = client.LbVserver.Stats(ctx, lbv.Name, nil); err != nil {
			panic(err)
		}

		if _, err = client.LbVserver.GetServiceBindings(ctx, lbv.Name, nil, nil); err != nil {
			panic(err)
		}

		if _, err = client.LbVserver.GetServiceGroupBindings(ctx, lbv.Name, nil, nil); err != nil {
			panic(err)
		}
	}

	// nsconfig
	// var nsconfig config.NsConfig
	if _, err = client.NsConfig.Get(ctx, nil); err != nil {
		panic(err)
	}
	// fmt.Println(nsconfig)

	// nsfeature
	// var nsfeature config.NsFeature
	if _, err = client.NsFeature.Get(ctx); err != nil {
		panic(err)
	}
	// fmt.Println(nsfeature)

	// nsmode
	// var nsmode config.NsMode
	if _, err = client.NsMode.Get(ctx); err != nil {
		panic(err)
	}
	// fmt.Println(nsmode)

	// nsversion
	// var nsversion config.NsVersionDetail
	if _, err = client.NsVersion.Get(ctx); err != nil {
		panic(err)
	}

	// Policystringmap
	var psms []config.PolicyStringmap
	if psms, err = client.PolicyStringmap.List(ctx, nil, nil); err != nil {
		panic(err)
	}
	var psms_count float64
	if psms_count, err = client.PolicyStringmap.Count(ctx); err != nil {
		panic(err)
	}
	if len(psms) != int(psms_count) {
		panic("number of policy stringmap expected " + fmt.Sprint(len(psms)))
	}
	for _, psm := range psms {
		// var psm_bindings []config.PolicyStringmapPatternBinding
		if _, err = client.PolicyStringmap.GetBindings(ctx, psm.Name, nil, nil); err != nil {
			panic(err)
		}
		// for _, psm_binding := range psm_bindings {
		// 	fmt.Println(psm_binding)
		// }
	}

	// Responder Action
	var resacts []config.ResponderAction
	if resacts, err = client.ResponderAction.List(ctx, nil, nil); err != nil {
		panic(err)
	}
	var resact_count float64
	if resact_count, err = client.ResponderAction.Count(ctx); err != nil {
		panic(err)
	}
	if len(resacts) != int(resact_count) {
		panic("number of responderactions expected " + fmt.Sprint(len(resacts)))
	}

	// Responder Policy
	var respol []config.ResponderPolicy
	if respol, err = client.ResponderPolicy.List(ctx, nil, nil); err != nil {
		panic(err)
	}
	var respol_count float64
	if respol_count, err = client.ResponderPolicy.Count(ctx); err != nil {
		panic(err)
	}
	if len(respol) != int(respol_count) {
		panic("number of responderpolicy expected " + fmt.Sprint(len(respol)))
	}

	// Server
	var srvs []config.Server
	if srvs, err = client.Server.List(ctx, nil, nil); err != nil {
		panic(err)
	}
	var srv_count float64
	if srv_count, err = client.Server.Count(ctx); err != nil {
		panic(err)
	}
	if len(srvs) != int(srv_count) {
		panic("number of servers expected " + fmt.Sprint(len(srvs)))
	}

	var svgs []config.ServiceGroup
	if svgs, err = client.ServiceGroup.List(ctx, nil, nil); err != nil {
		panic(err)
	}
	for _, svg := range svgs {
		// var svgstat stat.ServiceGroup
		if _, err = client.ServiceGroup.Stats(ctx, svg.Name, nil); err != nil {
			panic(err)
		}

		var svg_memberbindings []config.ServiceGroupServiceGroupMemberBinding
		if svg_memberbindings, err = client.ServiceGroup.GetServiceGroupMemberBindings(ctx, svg.Name, nil, nil); err != nil {
			panic(err)
		}
		var svg_memberbindings_count float64
		if svg_memberbindings_count, err = client.ServiceGroup.CountServiceGroupMemberBindings(ctx, svg.Name); err != nil {
			panic(err)
		}
		if len(svg_memberbindings) != int(svg_memberbindings_count) {
			panic("number of svg memberbindings expected " + fmt.Sprint(len(svg_memberbindings)))
		}
	}

	// Service
	var svcs []config.Service
	if svcs, err = client.Service.List(ctx, nil, nil); err != nil {
		panic(err)
	}
	var svc_count float64
	if svc_count, err = client.Service.Count(ctx); err != nil {
		panic(err)
	}
	if len(svcs) != int(svc_count) {
		panic("number of services expected " + fmt.Sprint(len(svcs)))
	}
	for _, svc := range svcs {
		// var svcstat stat.Service
		if _, err = client.Service.Stats(ctx, svc.Name, nil); err != nil {
			panic(err)
		}
	}

	fmt.Println("Done.")
	os.Exit(0)
}
