package main

// GraphicsApp simulation with undo and redo functionality
type GraphicsApp struct {
	undoStack *LLStack
	redoStack *LLStack
}

// NewGraphicsApp ctor for our graphics app
func NewGraphicsApp() *GraphicsApp {
	t := GraphicsApp{}
	t.undoStack = NewLLStack()
	t.redoStack = NewLLStack()
	return &t
}

func (ga *GraphicsApp) applyFilter(filter string) {
	ga.undoStack.push(filter)
}

func (ga *GraphicsApp) undo() {
	undone := ga.undoStack.pop()
	ga.redoStack.push(undone)
}

func (ga *GraphicsApp) redo() {
	redone := ga.redoStack.pop()
	ga.undoStack.push(redone)
}
