//go:build ffi

package service

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	surveyterm "github.com/AlecAivazis/survey/v2/terminal"
	"github.com/XMLHexagram/tdl_ffi/pb"
	"github.com/fatih/color"
	"github.com/iyear/tdl/action"
	"github.com/iyear/tdl/app/login"
	"github.com/iyear/tdl/cmd"
	"go.etcd.io/bbolt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
)

type Service struct {
	legacyStdout *os.File
	w            *os.File
	r            *os.File
	scanner      *bufio.Scanner
	cancel       context.CancelFunc
	ctx          context.Context

	actionChannel chan *pb.ActionResponse

	pb.UnimplementedTdlServer
}

var S = &Service{
	actionChannel: make(chan *pb.ActionResponse, 100),
}

func (s *Service) Init() {
	s.redirectStdout()
	color.Output = s.w
	color.NoColor = true
}

func (s *Service) Start(port int) {
	s.Init()
	//fmt.Println("Starting gRPC server")
	//go func() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println(err.Error())
	}
	test := grpc.NewServer()
	pb.RegisterTdlServer(test, s)
	if err = test.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Service) redirectStdout() {
	var err error
	s.r, s.w, err = os.Pipe()
	if err != nil {
		panic(err)
	}

	s.legacyStdout = os.Stdout
	os.Stdout = s.w

	s.scanner = bufio.NewScanner(s.r)
}

func (s *Service) GetAction(ctx context.Context, request *pb.ActionRequest) (*pb.ActionResponse, error) {
	s.scanner.Scan()
	if err := s.scanner.Err(); err != nil {
		panic(err)
	}
	raw := s.scanner.Text()

	//fmt.Fprintln(s.legacyStdout, raw)

	a := new(action.Action)
	switch {
	case json.Unmarshal([]byte(raw), a) == nil:
		s.actionChannel <- &pb.ActionResponse{
			ActionType: string(a.Type),
			Payload:    a.Payload,
			RawMessage: a.Raw,
		}
	default:
		s.actionChannel <- &pb.ActionResponse{
			ActionType: string(action.ActionTypePlaceholder),
			Payload:    "",
			RawMessage: raw,
		}
	}
	return <-s.actionChannel, nil
}

// RunScript example: "tdl dl xxx"
func (s *Service) RunScript(ctx context.Context, request *pb.RunScriptRequest) (*pb.RunScriptResponse, error) {
	script := request.Script

	s.ctx, s.cancel = signal.NotifyContext(context.Background(), os.Interrupt)
	os.Args = strings.Split(script, " ")

	defer s.cancel()

	humanizeErrors := map[error]string{
		bbolt.ErrTimeout:        "Current database is used by another process, please terminate it first",
		surveyterm.InterruptErr: "Interrupted",
	}

	if err := cmd.New().ExecuteContext(s.ctx); err != nil {
		for e, m := range humanizeErrors {
			if errors.Is(err, e) {
				color.Red("%s", m)
			}
		}

		color.Red("Error: %+v", err)
		fmt.Println("TODO1")
	}

	fmt.Println("TODO2")

	return &pb.RunScriptResponse{IsSuccessStart: true}, nil
}

func (s *Service) AuthInput(ctx context.Context, request *pb.AuthInputRequest) (*pb.AuthInputResponse, error) {
	fmt.Println(request)
	login.Input(request.Input)
	return &pb.AuthInputResponse{}, nil
}
