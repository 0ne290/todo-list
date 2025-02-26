package task

import (
	"testing"
	"time"
)

func TestNewTask1(t *testing.T) {
	// Arrange
	expectedTitle := "AnyTitle123"
	description := "AnyDescription123"
	expectedDescription := &description
	expectedCreatedAt := time.Now()

	// Act
	task := NewTask(expectedTitle, expectedDescription, expectedCreatedAt)

	// Assert
	if task == nil {
		t.Errorf("Task was incorrect, got: %s, want: %s.", "nil", "not nil")

		return
	}
	if task.Id != 0 {
		t.Errorf("Id was incorrect, got: %d, want: %d.", task.Id, 0)

		return
	}
	if task.Title != expectedTitle {
		t.Errorf("Title was incorrect, got: %s, want: %s.", task.Title, expectedTitle)

		return
	}
	if task.Description != expectedDescription {
		t.Errorf("Description was incorrect, got: %s, want: %s.", dereferencePointerToString(task.Description), dereferencePointerToString(expectedDescription))
	
		return
	}
	if task.Status != New {
		t.Errorf("Status was incorrect, got: %s, want: %s.", task.Status, New)

		return
	}
	if task.CreatedAt != expectedCreatedAt {
		t.Errorf("CreatedAt was incorrect, got: %s, want: %s.", task.CreatedAt, expectedCreatedAt)

		return
	}
	if task.UpdatedAt != expectedCreatedAt {
		t.Errorf("UpdatedAt was incorrect, got: %s, want: %s.", task.UpdatedAt, expectedCreatedAt)

		return
	}
}

func TestNewTask2(t *testing.T) {
	// Arrange
	expectedTitle := "AnyTitle123"
	var expectedDescription *string = nil
	expectedCreatedAt := time.Now()

	// Act
	task := NewTask(expectedTitle, expectedDescription, expectedCreatedAt)

	// Assert
	if task == nil {
		t.Errorf("Task was incorrect, got: %s, want: %s.", "nil", "not nil")

		return
	}
	if task.Id != 0 {
		t.Errorf("Id was incorrect, got: %d, want: %d.", task.Id, 0)

		return
	}
	if task.Title != expectedTitle {
		t.Errorf("Title was incorrect, got: %s, want: %s.", task.Title, expectedTitle)

		return
	}
	if task.Description != expectedDescription {
		t.Errorf("Description was incorrect, got: %s, want: %s.", dereferencePointerToString(task.Description), dereferencePointerToString(expectedDescription))
	
		return
	}
	if task.Status != New {
		t.Errorf("Status was incorrect, got: %s, want: %s.", task.Status, New)

		return
	}
	if task.CreatedAt != expectedCreatedAt {
		t.Errorf("CreatedAt was incorrect, got: %s, want: %s.", task.CreatedAt, expectedCreatedAt)

		return
	}
	if task.UpdatedAt != expectedCreatedAt {
		t.Errorf("UpdatedAt was incorrect, got: %s, want: %s.", task.UpdatedAt, expectedCreatedAt)

		return
	}
}

func TestUpdate1(t *testing.T) {
	// Arrange
	expectedStatus := InProgress
	
	description := "AnyDescription123"
	task := NewTask("AnyTitle123", &description, time.Now())

	// Act
	actualError := task.Update()

	// Assert
	if actualError != nil {
		t.Errorf("Error was incorrect, got: %s, want: %s.", actualError.Error(), "nil")

		return
	}
	if task.Status != expectedStatus {
		t.Errorf("Status was incorrect, got: %s, want: %s.", task.Status, expectedStatus)

		return
	}
}

func TestUpdate2(t *testing.T) {
	// Arrange
	expectedStatus := Done
	
	description := "AnyDescription123"
	task := NewTask("AnyTitle123", &description, time.Now())

	task.Update()

	// Act
	actualError := task.Update()

	// Assert
	if actualError != nil {
		t.Errorf("Error was incorrect, got: %s, want: %s.", actualError.Error(), "nil")

		return
	}
	if task.Status != expectedStatus {
		t.Errorf("Status was incorrect, got: %s, want: %s.", task.Status, expectedStatus)

		return
	}
}

func TestUpdate3(t *testing.T) {
	// Arrange
	expectedStatus := Done
	
	description := "AnyDescription123"
	task := NewTask("AnyTitle123", &description, time.Now())

	task.Update()
	task.Update()

	// Act
	actualError := task.Update()

	// Assert
	if actualError == nil {
		t.Errorf("Error was incorrect, got: %s, want: %s.", "nil", "not nil")

		return
	}
	if task.Status != expectedStatus {
		t.Errorf("Status was incorrect, got: %s, want: %s.", task.Status, expectedStatus)

		return
	}
}

func dereferencePointerToString(pointerToString *string) string {
	if pointerToString == nil {
		return "nil"
	}

	return *pointerToString
}