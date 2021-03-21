package main

import "testing"

func TestUndo(t *testing.T) {
	ga := NewGraphicsApp()

	ga.applyFilter("blur")

	if ga.undoStack.peek() != "blur" {
		t.Error("expected undo to contain last filter")
	}

	if ga.redoStack.peek() != nil {
		t.Error("expected redo stack to contain nothing")
	}
}

func TestRedoUndo(t *testing.T) {
	ga := NewGraphicsApp()

	ga.applyFilter("blur")
	ga.applyFilter("liquify")

	ga.undo()
	ga.redo()

	if ga.undoStack.peek() != "liquify" {
		t.Error("expected filter (liquify) to be the last in undo stack")
	}

	if ga.redoStack.peek() != nil {
		t.Error("expected redo stack to be empty")
	}

	ga.undo()
	ga.undo()

	if ga.undoStack.peek() != nil {
		t.Error("expected undo stack to be empty now")
	}

	if ga.redoStack.peek() != "blur" {
		t.Error("expected last in redo to be filter (blur)")
	}
}
