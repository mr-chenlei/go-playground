package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	pb "code.lstaas.com/lightspeed/gravity/core/app/authentication/command"
	env "code.lstaas.com/lightspeed/gravity/core/app/env/command"
	v1 "code.lstaas.com/lightspeed/gravity/core/app/genesis/grpcapi/v1"
	paths "code.lstaas.com/lightspeed/gravity/core/app/paths/command/version1"

	"google.golang.org/grpc"
)

func main() {
	apiServerAddr := flag.String("a", "", "gravity server addr")
	serverList := flag.Bool("l", false, "server list")
	groupInfo := flag.Bool("g", false, "group info")
	template := flag.Bool("d", false, "division template")
	pingTarget := flag.Bool("t", false, "ping targets")
	publicIP := flag.String("p", "", "public IP")
	sub := flag.String("s", "", "path or subnet path")
	end := flag.String("e", "", "end in path")

	flag.Parse()

	if *apiServerAddr == "" {
		log.Println(`need gravity addr, please try again with "a" filed.`)
		return
	}
	fmt.Printf("get Info from:%v\n", *apiServerAddr)
	if *serverList {
		fmt.Println("------ getServerListFromGravity ------")
		err := getServerListFromGravity(*apiServerAddr)
		if err != nil {
			log.Println("getServerListFromGravity err:", err)
		}
		return
	}
	if *groupInfo {
		fmt.Println("------ getGroupFromGravity ------")
		err := getGroupFromGravity(*apiServerAddr)
		if err != nil {
			log.Println("getGroupFromGravity err:", err)
		}
		return
	}
	if *template {
		fmt.Println("------ getDivisionFromGravity ------")
		err := getDivisionFromGravity("DivisionCurrent", *apiServerAddr)
		if err != nil {
			log.Println("getDivisionFromGravity err:", err)
		}
		return
	}

	if *pingTarget {
		fmt.Println("------ getPingTargetFromGravity ------")
		err := getPingTargetFromGravity(*publicIP, *apiServerAddr)
		if err != nil {
			log.Println("getPingTargetFromGravity err:", err)
		}
		return
	}
	if *sub == "" {
		fmt.Println("------ getPathFromGravity ------")
		err := getPathFromGravity(*publicIP, *apiServerAddr)
		if err != nil {
			log.Println("getPathFromGravity err:", err)
		}
		return
	} else {
		ss, err := strconv.Atoi(*sub)
		if err != nil {
			log.Println("strconv.Atoi err:", err)
			return
		}
		service := ss / 10
		strategy := ss % 10
		fmt.Println("------ getSubnetPathFromGravity ------")
		err = getSubnetPathFromGravity(*publicIP, *end, *apiServerAddr, uint32(service), uint32(strategy))
		if err != nil {
			log.Println("getPathFromGravity err:", err)
		}
		return
	}
	//getServerListFromGravity(*apiServerAddr)
	//getGroupFromGravity(*apiServerAddr)
	//getDivisionFromGravity(*apiServerAddr)
	//getPingTargetFromGravity(*publicIP, *apiServerAddr)
	//getPathFromGravity(*publicIP, *apiServerAddr)
	//getSubnetPathFromGravity(*publicIP, *apiServerAddr, 1, 1)
}

func getPingTargetFromGravity(publicIP, addr string) error {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := v1.NewNodeManagerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()
	fmt.Println("publicIP:", publicIP)

	stream, err := client.PullPingTargets(ctx, &v1.PullTargetRequest{PublicIP: publicIP})
	if err != nil {
		return err
	}
	response, err := stream.Recv()
	if err != nil {
		return err
	}
	//fmt.Println(response)
	for i, v := range response.Target {
		fmt.Printf("[%v]:%v\n", i, v)
	}
	return nil
}

func getServerListFromGravity(addr string) error {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()
	client := pb.NewServerListClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	serverListData, err := client.DownloadServerList(ctx, &pb.ServerListRequest{Msg: []byte{0x01, 0x02}})
	if err != nil {
		return err
	}
	for i, v := range serverListData.List {
		fmt.Printf("[%v]:%v\n", i, v)
	}
	return nil
}

func getSubnetPathFromGravity(start, end, addr string, service, strategy uint32) error {
	fmt.Println("getSubnetPathFromGravity start")
	//fmt.Printf("service: %v\tstrategy: %v\n", service, strategy)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	pathConn := paths.NewPathServiceClient(conn)

	pullRequest := paths.SubnetPathsRequest{
		DBGStartIP: start,
	}

	// client
	if service < 3 {
		pullRequest.Param = []*paths.LocalData{
			{
				DimensionInfo: &paths.Dimension{
					Service:  service,
					Strategy: strategy,
				},
			},
		}
	}
	//pullRequest := paths.SubnetPathsRequest{
	//	DBGStartIP: start,
	//	Param: []*paths.LocalData{
	//		{
	//			DimensionInfo: &paths.Dimension{
	//				Service:  service,
	//				Strategy: strategy,
	//			},
	//		},
	//	},
	//}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	pullStream, err := pathConn.PullSubnetPaths(ctx, &pullRequest)
	if err != nil {
		return err
	}
	pullReply, err := pullStream.Recv()
	if err == io.EOF || err != nil {
		return err
	}
	for sub, Paths := range pullReply.RoutePath {
		service := (sub >> 8) & 0xff
		strategy := sub & 0xff
		fmt.Println("service:", service, "strategy:", strategy)
		for from, toPath := range Paths.PathFrom {
			fmt.Println("from:", from)

			if end == "" {
				for to, p := range toPath.To {
					fmt.Printf("\tto:%v, path:%v, time:%v\n", to, p.GroupN, time.Now().Format(time.RFC3339))
					fmt.Printf("\t\tquality:")
					for _, v := range p.QualityN {
						fmt.Printf("%v,", v.Estimate)
					}
					fmt.Println()
				}
			} else {
				for to, p := range toPath.To {
					if to == end {
						fmt.Printf("\tto:%v, path:%v, time:%v\n", to, p.GroupN, time.Now().Format(time.RFC3339))
						fmt.Printf("\t\tquality:")
						for _, v := range p.QualityN {
							fmt.Printf("%v,", v.Estimate)
						}
						fmt.Println()
					}
				}
			}

		}
	}

	return nil
}

func getPathFromGravity(start, addr string) error {
	fmt.Println("getPathFromGravity")
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	pathConn := paths.NewPathServiceClient(conn)
	pullRequest := paths.GroupPathRequest{DBGStartIP: start}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	pullStream, err := pathConn.PullGroupPaths(ctx, &pullRequest)
	if err != nil {
		return err
	}

	//for {
	pullReply, err := pullStream.Recv()
	if err == io.EOF || err != nil {
		return err
	}
	for from, toPath := range pullReply.PathFrom {
		fmt.Println("from:", from)
		for to, path := range toPath.To {
			fmt.Printf("\tto:%v, path:%v, time:%v\n", to, path.GroupN, time.Now().Format(time.RFC3339))
			fmt.Printf("\t\tquality:")
			for _, v := range path.QualityN {
				fmt.Printf("%v,", v.Estimate)
			}
			fmt.Println()
		}
	}
	//}
	return nil
}

func getGroupFromGravity(addr string) error {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	pathConn := paths.NewPathServiceClient(conn)
	pullRequest := paths.GroupInfoRequest{Version: 0}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	pullStream, err := pathConn.PullGroupInfo(ctx, &pullRequest)
	if err != nil {
		return err
	}

	pullReply, err := pullStream.Recv()
	if err == io.EOF || err != nil {
		return err
	}
	totalGroupNum, totalServerNum := 0, 0
	for _, group := range pullReply.Groups {
		fmt.Println("group name:", group.Name)
		totalGroupNum++
		for _, node := range group.Info.Nodes {
			fmt.Printf("\thostname:%v, addr:%v\n", node.HostName, node.Addr)
			totalServerNum++
		}
	}
	fmt.Println("total Group num:", totalGroupNum)
	fmt.Println("total server num:", totalServerNum)
	return nil
}

func getDivisionFromGravity(divName, addr string) error {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	envConn := env.NewEnvServiceClient(conn)
	request := &env.EnvConfigRequest{
		ConfigName: divName,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := envConn.PullEnvCfg(ctx, request)
	if err != nil {
		return err
	}
	if resp.Content == "" {
		log.Println("content is empty, do not update")
		return nil
	}
	log.Printf("content:\n%v", resp.Content)
	return nil
}
