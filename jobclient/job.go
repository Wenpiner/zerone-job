// Code generated by goctl. DO NOT EDIT.
// Source: job.proto

package jobclient

import (
	"context"

	"github.com/suyuan32/simple-admin-job/job"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseIDResp      = job.BaseIDResp
	BaseResp        = job.BaseResp
	BaseUUIDResp    = job.BaseUUIDResp
	Empty           = job.Empty
	IDReq           = job.IDReq
	IDsReq          = job.IDsReq
	PageInfoReq     = job.PageInfoReq
	TaskInfo        = job.TaskInfo
	TaskListReq     = job.TaskListReq
	TaskListResp    = job.TaskListResp
	TaskLogInfo     = job.TaskLogInfo
	TaskLogListReq  = job.TaskLogListReq
	TaskLogListResp = job.TaskLogListResp
	UUIDReq         = job.UUIDReq
	UUIDsReq        = job.UUIDsReq

	Job interface {
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
		// Task management
		CreateTask(ctx context.Context, in *TaskInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateTask(ctx context.Context, in *TaskInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetTaskList(ctx context.Context, in *TaskListReq, opts ...grpc.CallOption) (*TaskListResp, error)
		GetTaskById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*TaskInfo, error)
		DeleteTask(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// TaskLog management
		CreateTaskLog(ctx context.Context, in *TaskLogInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateTaskLog(ctx context.Context, in *TaskLogInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetTaskLogList(ctx context.Context, in *TaskLogListReq, opts ...grpc.CallOption) (*TaskLogListResp, error)
		GetTaskLogById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*TaskLogInfo, error)
		DeleteTaskLog(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
	}

	defaultJob struct {
		cli zrpc.Client
	}
)

func NewJob(cli zrpc.Client) Job {
	return &defaultJob{
		cli: cli,
	}
}

func (m *defaultJob) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	client := job.NewJobClient(m.cli.Conn())
	return client.InitDatabase(ctx, in, opts...)
}

// Task management
func (m *defaultJob) CreateTask(ctx context.Context, in *TaskInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := job.NewJobClient(m.cli.Conn())
	return client.CreateTask(ctx, in, opts...)
}

func (m *defaultJob) UpdateTask(ctx context.Context, in *TaskInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := job.NewJobClient(m.cli.Conn())
	return client.UpdateTask(ctx, in, opts...)
}

func (m *defaultJob) GetTaskList(ctx context.Context, in *TaskListReq, opts ...grpc.CallOption) (*TaskListResp, error) {
	client := job.NewJobClient(m.cli.Conn())
	return client.GetTaskList(ctx, in, opts...)
}

func (m *defaultJob) GetTaskById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*TaskInfo, error) {
	client := job.NewJobClient(m.cli.Conn())
	return client.GetTaskById(ctx, in, opts...)
}

func (m *defaultJob) DeleteTask(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := job.NewJobClient(m.cli.Conn())
	return client.DeleteTask(ctx, in, opts...)
}

// TaskLog management
func (m *defaultJob) CreateTaskLog(ctx context.Context, in *TaskLogInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := job.NewJobClient(m.cli.Conn())
	return client.CreateTaskLog(ctx, in, opts...)
}

func (m *defaultJob) UpdateTaskLog(ctx context.Context, in *TaskLogInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := job.NewJobClient(m.cli.Conn())
	return client.UpdateTaskLog(ctx, in, opts...)
}

func (m *defaultJob) GetTaskLogList(ctx context.Context, in *TaskLogListReq, opts ...grpc.CallOption) (*TaskLogListResp, error) {
	client := job.NewJobClient(m.cli.Conn())
	return client.GetTaskLogList(ctx, in, opts...)
}

func (m *defaultJob) GetTaskLogById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*TaskLogInfo, error) {
	client := job.NewJobClient(m.cli.Conn())
	return client.GetTaskLogById(ctx, in, opts...)
}

func (m *defaultJob) DeleteTaskLog(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := job.NewJobClient(m.cli.Conn())
	return client.DeleteTaskLog(ctx, in, opts...)
}
