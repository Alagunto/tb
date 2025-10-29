package telegram

import "time"

// Checklist represents a checklist message.
type Checklist struct {
	Title                    string
	TitleEntities            []MessageEntity
	Tasks                    []ChecklistTask
	OthersCanAddTasks        bool
	OthersCanMarkTasksAsDone bool
}

// ChecklistTask represents a task in a checklist.
type ChecklistTask struct {
	ID              int
	Text            string
	TextEntities    []MessageEntity
	CompletedByUser *User
	CompletionDate  *time.Time
}

// ChecklistTasksDone represents tasks that were marked as done.
type ChecklistTasksDone struct {
	ChecklistMessage       *Message
	MarkedAsDoneTaskIDs    []int
	MarkedAsNotDoneTaskIDs []int
}

// ChecklistTasksAdded represents tasks that were added to a checklist.
type ChecklistTasksAdded struct {
	ChecklistMessage *Message
	Tasks            []ChecklistTask
}
