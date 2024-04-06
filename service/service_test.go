package service

import (
	"context"
	"fmt"
	"github.com/XMLHexagram/go-must/noerror"
	"os/signal"

	//"github.com/XMLHexagram/tdl_ffi"
	"github.com/XMLHexagram/tdl_ffi/pb"
	"github.com/iyear/tdl/action"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	_ "os"
	"testing"
	"time"
)

func newGRPCClient() (pb.TdlClient, *grpc.ClientConn) {
	conn := noerror.Guard(grpc.NewClient("localhost:12210", grpc.WithTransportCredentials(insecure.NewCredentials())))
	c := pb.NewTdlClient(conn)
	return c, conn
}

func TestService_Login(t *testing.T) {
	go S.Start(12210)
	<-time.After(3 * time.Second)

	_, conn := newGRPCClient()
	c := pb.NewTdlClient(conn)
	//defer noerror.Guard0(conn.Close())
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	go func() {
		_ = noerror.Guard(c.RunScript(ctx, &pb.RunScriptRequest{Script: "tdl login -T code"}))
	}()

	for {
		res := noerror.Guard(c.GetAction(ctx, &pb.ActionRequest{}))

		noerror.Guard(fmt.Fprintln(S.legacyStdout, res))
		actionType := noerror.Guard(action.ParseActionType(res.ActionType))
		if actionType == action.ActionTypeCodeLoginPhoneInput {
			break
		}
		<-time.After(1 * time.Second)
	}

	noerror.Guard(fmt.Fprintln(S.legacyStdout, "please input phone number by grpc"))

	for {
		res := noerror.Guard(c.GetAction(ctx, &pb.ActionRequest{}))
		noerror.Guard(fmt.Fprintln(S.legacyStdout, res))
		actionType := noerror.Guard(action.ParseActionType(res.ActionType))
		if actionType == action.ActionTypeCodeLoginCodeInput {
			break
		}
		<-time.After(1 * time.Second)
	}

	noerror.Guard(fmt.Fprintln(S.legacyStdout, "please input code by grpc"))

	for {
		res := noerror.Guard(c.GetAction(ctx, &pb.ActionRequest{}))

		noerror.Guard(fmt.Fprintln(S.legacyStdout, res))
	}

	return
}

func TestService_AuthInput(t *testing.T) {
	c, _ := newGRPCClient()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_ = noerror.Guard(c.AuthInput(ctx, &pb.AuthInputRequest{Input: os.Getenv("INPUT")}))
}
