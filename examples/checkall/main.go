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
	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/stat"
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
		switch line {
		case 0:
			NITRO_ENV_NAME = values[1]
			fmt.Println("NITRO_ENV_NAME:", NITRO_ENV_NAME)
		case 1:
			NITRO_ENV_ADDRESS = values[1]
			fmt.Println("NITRO_ENV_ADDRESS:", NITRO_ENV_ADDRESS)
		case 2:
			NITRO_ENV_USERNAME = values[1]
			fmt.Println("NITRO_ENV_USERNAME:", NITRO_ENV_USERNAME)
		case 3:
			NITRO_ENV_PASSWORD = values[1]
			if NITRO_ENV_PASSWORD == "" {
				panic(errors.New("NITRO_ENV_PASSWORD is empty"))
			}
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
	var primary bool
	if primary, err = client.IsPrimaryNode(ctx); err != nil {
		panic(err)
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
	os.Exit(0)
	var cfg config.NsConfig
	if cfg, err = client.NsConfig.Get(ctx, nil); err != nil {
		fmt.Println(err)
		fmt.Println(errors.Unwrap(err))
	}
	fmt.Println(cfg)

	if primary && cfg.ConfigChanged {
		if err = client.SaveConfig(ctx); err != nil {
			fmt.Println(err)
			fmt.Println(errors.Unwrap(err))
		}
	}

	var version config.NsVersionDetail
	if version, err = client.NsVersion.Get(ctx); err != nil {
		fmt.Println(err)
		fmt.Println(errors.Unwrap(err))
	}
	fmt.Println(version)

	var bindings []config.PolicyStringmapPatternBinding
	if bindings, err = client.PolicyStringmap.GetBindings(ctx, "SM_CL1009_CS_CONTROL", []string{"name", "key"}, nil); err != nil {
		fmt.Println(err)
		fmt.Println(errors.Unwrap(err))
	}
	for _, binding := range bindings {
		fmt.Println(binding)
	}

	var svgs []config.ServiceGroup
	if svgs, err = client.ServiceGroup.List(ctx, nil, nil); err != nil {
		fmt.Println(err)
		fmt.Println(errors.Unwrap(err))
		return
	}
	for _, svg := range svgs {
		fmt.Println(svg)
		var svgstat stat.ServiceGroup
		if svgstat, err = client.ServiceGroup.Stats(ctx, svg.Name, nil); err != nil {
			fmt.Println(errors.Unwrap(err))
			return
		}
		fmt.Println(svgstat)

		var countsvgbindings float64
		if countsvgbindings, err = client.ServiceGroup.CountServiceGroupMemberBindings(ctx, svg.Name); err != nil {
			fmt.Println(errors.Unwrap(err))
			return
		}

		var svgbindings []config.ServiceGroupServiceGroupMemberBinding
		if svgbindings, err = client.ServiceGroup.GetServiceGroupMemberBindings(ctx, svg.Name, nil, nil); err != nil {
			fmt.Println(errors.Unwrap(err))
			return
		}
		for _, svgbinding := range svgbindings {
			fmt.Println(svgbinding)
		}
		fmt.Println("Total Bindings", countsvgbindings, len(svgbindings), "\n")
	}

	var svcs []config.Service
	if svcs, err = client.Service.List(ctx, nil, nil); err != nil {
		fmt.Println(errors.Unwrap(err))
	}
	for _, svc := range svcs {
		fmt.Println(svc, "\n")

		var svcstat stat.Service
		if svcstat, err = client.Service.Stats(ctx, svc.Name, nil); err != nil {
			fmt.Println(errors.Unwrap(err))
			return
		}
		fmt.Println(svcstat, "\n")
	}
}
