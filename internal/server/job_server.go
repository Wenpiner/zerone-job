// Code generated by goctl. DO NOT EDIT.
// Source: job.proto

package server

import (
	"context"

	"github.com/suyuan32/simple-admin-job/internal/logic/base"
	"github.com/suyuan32/simple-admin-job/internal/logic/task"
	"github.com/suyuan32/simple-admin-job/internal/svc"
	"github.com/suyuan32/simple-admin-job/job"
)

type JobServer struct {
	svcCtx *svc.ServiceContext
	job.UnimplementedJobServer
}

func NewJobServer(svcCtx *svc.ServiceContext) *JobServer {
	return &JobServer{
		svcCtx: svcCtx,
	}
}

func (s *JobServer) InitDatabase(ctx context.Context, in *job.Empty) (*job.BaseResp, error) {
	l := base.NewInitDatabaseLogic(ctx, s.svcCtx)
	return l.InitDatabase(in)
}

// Task management
func (s *JobServer) CreateTask(ctx context.Context, in *job.TaskInfo) (*job.BaseIDResp, error) {
	l := task.NewCreateTaskLogic(ctx, s.svcCtx)
	return l.CreateTask(in)
}

func (s *JobServer) UpdateTask(ctx context.Context, in *job.TaskInfo) (*job.BaseResp, error) {
	l := task.NewUpdateTaskLogic(ctx, s.svcCtx)
	return l.UpdateTask(in)
}

func (s *JobServer) GetTaskList(ctx context.Context, in *job.TaskListReq) (*job.TaskListResp, error) {
	l := task.NewGetTaskListLogic(ctx, s.svcCtx)
	return l.GetTaskList(in)
}

func (s *JobServer) GetTaskById(ctx context.Context, in *job.IDReq) (*job.TaskInfo, error) {
	l := task.NewGetTaskByIdLogic(ctx, s.svcCtx)
	return l.GetTaskById(in)
}

func (s *JobServer) DeleteTask(ctx context.Context, in *job.IDsReq) (*job.BaseResp, error) {
	l := task.NewDeleteTaskLogic(ctx, s.svcCtx)
	return l.DeleteTask(in)
}
