package main

import (
	"bytes"
	"context"
	"crypto/md5"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "code.lstaas.com/lightspeed/gravity/core/app/env/command"
	pb2 "code.lstaas.com/lightspeed/gravity/core/app/paths/command/version1"
)

var (
	gravity = flag.String("gravity", "gravity.lstaas.com:8001", "Gravity address.")
	period  = flag.String("period", "1m", "Period of pull from gravity.")
)

var (
	logger *log.Logger
)

func pullOriginPath(pathConn pb2.PathServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()
	request := &pb2.GroupPathRequest{Version: 0}
	stream, err := pathConn.PullGroupPaths(ctx, request)
	if err != nil {
		return err
	}
	if _, err := stream.Recv(); err != nil {
		return err
	}
	return nil
}

func pullOriginGroup(pathConn pb2.PathServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()
	request := &pb2.GroupInfoRequest{Version: 0}
	stream, err := pathConn.PullGroupInfo(ctx, request)
	if err != nil {
		return err
	}
	if _, err := stream.Recv(); err != nil {
		return err
	}
	return nil
}

func pullSubnetPath(pathConn pb2.PathServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()

	request := &pb2.SubnetPathsRequest{Param: []*pb2.LocalData{
		&pb2.LocalData{DimensionInfo: &pb2.Dimension{
			Service:  0,
			Strategy: 0,
		}},
	}}
	stream, err := pathConn.PullSubnetPaths(ctx, request)
	if err != nil {
		return err
	}
	if _, err := stream.Recv(); err != nil {
		return err
	}
	return nil
}

func pullSubnetGroup(pathConn pb2.PathServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()

	request := &pb2.SubnetInfosRequest{Param: []*pb2.LocalData{
		&pb2.LocalData{DimensionInfo: &pb2.Dimension{
			Service:  0,
			Strategy: 0,
		}},
	}}
	stream, err := pathConn.PullSubnetInfo(ctx, request)
	if err != nil {
		return err
	}
	if _, err := stream.Recv(); err != nil {
		return err
	}
	return nil
}

func pullConfig(addr string) error {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	envConn := pb.NewEnvServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()

	request := &pb.EnvConfigRequest{
		ConfigName: "DivisionCurrent",
		Hash:       []byte{},
	}
	resp, err := envConn.PullEnvCfg(ctx, request)
	if err != nil {
		return fmt.Errorf("pull cfg template from gravity failed")
	}
	if resp.Content == "" {
		return nil
	}
	tempHash := md5.Sum([]byte(resp.Content))
	if bytes.Compare(tempHash[:], resp.Hash) != 0 {
		return fmt.Errorf("cfg template hash do not match")
	}
	return nil
}

func pullPath(addr string) error {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	grpcConn := pb2.NewPathServiceClient(conn)
	if err := pullSubnetGroup(grpcConn); err != nil {
		return err
	}
	if err := pullSubnetPath(grpcConn); err != nil {
		return err
	}
	if err := pullOriginPath(grpcConn); err != nil {
		return err
	}
	if err := pullOriginGroup(grpcConn); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()

	logFile, err := os.OpenFile("./voyager-test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		panic(err)
	}
	logger = log.New(logFile, "", log.LUTC|log.Ldate|log.Ltime|log.Lmicroseconds)

	addr := *gravity
	duration, err := time.ParseDuration(*period)
	if err != nil {
		logger.Panic(err)
	}

	logger.Println("voyager-test start working...")
	for range time.Tick(duration) {
		t1 := time.Now()
		logger.Println("1. prepare to pull config from", addr)
		if err := pullConfig(addr); err != nil {
			logger.Println("1. pull config failed, cost:", time.Since(t1), "error:", err)
			continue
		}
		logger.Println("1. pull config success, cost:", time.Since(t1))

		t2 := time.Now()
		logger.Println("2. prepare to pull path from", addr)
		if err := pullPath(addr); err != nil {
			logger.Println("2. pull path failed, cost:", time.Since(t2), "error:", err)
			continue
		}
		logger.Println("2. pull path success, cost:", time.Since(t2))
	}
}
