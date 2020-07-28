package common

type Job struct {
	Name       string `json:"name"`       // 任务名
	Command    string `json:"command"`    // shell命令
	Expression string `json:"expression"` // cron表达式
}
