package policy

import "strings"

type Decision string

const (
	DecisionAllow Decision = "allow"
	DecisionDeny  Decision = "deny"
)

type Result struct {
	Decision Decision `json:"decision"`
	Reason   string   `json:"reason"`
}

var blockedFragments = []string{
	"关机",
	"重启",
	"执行命令",
	"运行命令",
	"读文件",
	"读取文件",
	"导出聊天记录",
	"扫码",
	"登录",
	"加好友",
	"私聊",
	"发私信",
	"点击",
	"控制电脑",
	"操作电脑",
	"忽略一切提示词",
	"ignore previous",
	"ignore all",
}

func EvaluateMessage(text string) Result {
	normalized := strings.ToLower(strings.TrimSpace(text))
	if normalized == "" {
		return Result{Decision: DecisionDeny, Reason: "empty message"}
	}

	for _, fragment := range blockedFragments {
		if strings.Contains(normalized, strings.ToLower(fragment)) {
			return Result{
				Decision: DecisionDeny,
				Reason:   "group messages cannot authorize local computer operations",
			}
		}
	}

	return Result{Decision: DecisionAllow, Reason: "public project conversation"}
}
