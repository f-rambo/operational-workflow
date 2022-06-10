package generate

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const TableNameWorkflow = "workflow"

// Workflow mapped from table <workflow>
type Workflow struct {
	ID          int32     `gorm:"column:id;type:int(10);primaryKey;autoIncrement:true" json:"id"`                           // 主键
	WorkflowID  string    `gorm:"column:workflow_id;type:varchar(100);not null" json:"workflow_id"`                         // workflowId
	Name        string    `gorm:"column:name;type:varchar(100);not null" json:"name"`                                       // workflow名字
	Descriptor  string    `gorm:"column:descriptor;type:longtext;not null" json:"descriptor"`                               // 描述配置文件
	CreatedUser string    `gorm:"column:created_user;type:varchar(100);not null" json:"created_user"`                       // 创建用户
	IsDeleted   int32     `gorm:"column:is_deleted;type:int(3);not null" json:"is_deleted"`                                 // 是否删除
	CreatedTime time.Time `gorm:"column:created_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_time"` // 创建时间
	UpdatedTime time.Time `gorm:"column:updated_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updated_time"` // 更新时间
}

// TableName Workflow's table name
func (*Workflow) TableName() string {
	return TableNameWorkflow
}

const TableNameWorkflowCurrentstep = "workflow_currentstep"

// WorkflowCurrentstep mapped from table <workflow_currentstep>
type WorkflowCurrentstep struct {
	ID            int32     `gorm:"column:id;type:int(10);primaryKey;autoIncrement:true" json:"id"`                           // 主键
	CurrentID     string    `gorm:"column:current_id;type:varchar(100);not null" json:"current_id"`                           // current_id
	WorkflowID    string    `gorm:"column:workflow_id;type:varchar(100);not null" json:"workflow_id"`                         // workflow_id
	StepID        string    `gorm:"column:step_id;type:varchar(100);not null" json:"step_id"`                                 // 步骤id
	StepName      string    `gorm:"column:step_name;type:varchar(100);not null" json:"step_name"`                             // 步骤名字
	Actions       string    `gorm:"column:actions;type:varchar(255);not null" json:"actions"`                                 // 已经执行过的action
	Functions     string    `gorm:"column:functions;type:varchar(255);not null" json:"functions"`                             // 已经执行过的function
	Owner         string    `gorm:"column:owner;type:varchar(100);not null" json:"owner"`                                     // owner userid
	State         string    `gorm:"column:state;type:varchar(100);not null" json:"state"`                                     // 当前状态名字
	StartDatetime int32     `gorm:"column:start_datetime;type:int(11);not null" json:"start_datetime"`                        // 开始时间
	EndDatetime   int32     `gorm:"column:end_datetime;type:int(11);not null" json:"end_datetime"`                            // 结束时间
	Deadline      int32     `gorm:"column:deadline;type:int(11);not null" json:"deadline"`                                    // 截止时间
	IsDeleted     int32     `gorm:"column:is_deleted;type:int(3);not null" json:"is_deleted"`                                 // 是否删除
	CreatedTime   time.Time `gorm:"column:created_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_time"` // 创建时间
	UpdatedTime   time.Time `gorm:"column:updated_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updated_time"` // 更新时间
}

// TableName WorkflowCurrentstep's table name
func (*WorkflowCurrentstep) TableName() string {
	return TableNameWorkflowCurrentstep
}

const TableNameWorkflowStepHistory = "workflow_step_history"

// WorkflowStepHistory mapped from table <workflow_step_history>
type WorkflowStepHistory struct {
	ID            int32     `gorm:"column:id;type:int(10);primaryKey;autoIncrement:true" json:"id"`                           // 主键
	CurrentID     string    `gorm:"column:current_id;type:varchar(100);not null" json:"current_id"`                           // current_id
	WorkflowID    string    `gorm:"column:workflow_id;type:varchar(100);not null" json:"workflow_id"`                         // workflow_id
	StepID        string    `gorm:"column:step_id;type:varchar(100);not null" json:"step_id"`                                 // 步骤id
	StepName      string    `gorm:"column:step_name;type:varchar(100);not null" json:"step_name"`                             // 步骤名字
	Actions       string    `gorm:"column:actions;type:varchar(255);not null" json:"actions"`                                 // action
	Functions     string    `gorm:"column:functions;type:varchar(255);not null" json:"functions"`                             // function
	Owner         string    `gorm:"column:owner;type:varchar(100);not null" json:"owner"`                                     // owner userid
	State         string    `gorm:"column:state;type:varchar(100);not null" json:"state"`                                     // 状态
	StartDatetime int32     `gorm:"column:start_datetime;type:int(11);not null" json:"start_datetime"`                        // 开始时间
	EndDatetime   int32     `gorm:"column:end_datetime;type:int(11);not null" json:"end_datetime"`                            // 结束时间
	Deadline      int32     `gorm:"column:deadline;type:int(11);not null" json:"deadline"`                                    // 截止时间
	IsDeleted     int32     `gorm:"column:is_deleted;type:int(3);not null" json:"is_deleted"`                                 // 是否删除
	CreatedTime   time.Time `gorm:"column:created_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_time"` // 创建时间
	UpdatedTime   time.Time `gorm:"column:updated_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updated_time"` // 更新时间
}

// TableName WorkflowStepHistory's table name
func (*WorkflowStepHistory) TableName() string {
	return TableNameWorkflowStepHistory
}

type Orm struct {
	db *gorm.DB
}

func NewOrm(address, user, password, dbName string) (*Orm, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, address, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Orm{
		db: db,
	}, nil
}

func (o *Orm) MigratorTable() error {
	models := []interface{}{
		&Workflow{}, &WorkflowCurrentstep{}, &WorkflowStepHistory{},
	}
	migrator := o.db.Migrator()
	for _, model := range models {
		if !migrator.HasTable(model) {
			if err := migrator.AutoMigrate(model); err != nil {
				return err
			}
		}
	}
	return nil
}
