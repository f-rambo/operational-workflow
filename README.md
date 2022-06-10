# Bworkflow
go-业务工作流

	1. 根据业务表上当前id初始化workflow
	2. 转化成fsm状态机
	3. 操作状态变迁
	4. LRU存储操作状态机？

	实体：
	functions、 step、conditions、actions

	某个动作导致步骤的变化，可以是一个或者多个动作；
	每个动作至少有一个条件，条件可以分为有结果和无结果的条件；
	每个动作之前和之后都有function；
	执行完动作之后可能会拆成两个或者合并成一个步骤；
	每个步骤都有一个指定的名字；
	一个工作流包含多个步骤。每一个步骤都有一个当前状态(例如, Queued, Underway, orFinished)。
	每一个步骤中都有一个或者多个动作可以被执行。每一个动作都可以设置执行条件(condition)，也可以设置执行函数(pre-function or post-function)。
	动作产生结果(result)，导致工作流的状态和当前步骤发生改变。
`
{
	"workflow_id": "xxx",
	"workflow_name": "sprint",
	"layout": {
	"initial_id": "open",
	"initial_name": "open",
	"states": ["create", "open", "close", "resolved", "inProgress", "reopen"],
	"transitions": [{
		"step_id": "1",
		"step_name": "open",
		"src_states": ["create"],
		"dts_state": "open",
		"mate_data": "",
		"actions": ["xxxx"],
		"functions": ["xxx"]
	}]
	}
	}
`