package tb

import "time"

type Checklist struct {
	Title                    string
	TitleEntities            []MessageEntity
	Tasks                    []ChecklistTask
	OthersCanAddTasks        bool
	OthersCanMarkTasksAsDone bool
}

type ChecklistTask struct {
	ID              int
	Text            string
	TextEntities    []MessageEntity
	CompletedByUser *User
	CompletionDate  *time.Time
}

type ChecklistTasksDone struct {
	ChecklistMessage       *Message
	MarkedAsDoneTaskIDs    []int
	MarkedAsNotDoneTaskIDs []int
}

type ChecklistTasksAdded struct {
	ChecklistMessage *Message
	Tasks            []ChecklistTask
}
