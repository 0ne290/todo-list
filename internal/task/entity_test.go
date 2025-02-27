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
	} else if task.Id != 0 {
		t.Errorf("Id was incorrect, got: %d, want: %d.", task.Id, 0)
	} else if task.Title != expectedTitle {
		t.Errorf("Title was incorrect, got: %s, want: %s.", task.Title, expectedTitle)
	} else if task.Description != expectedDescription {
		t.Errorf("Description was incorrect, got: %s, want: %s.", dereferencePointerToString(task.Description), dereferencePointerToString(expectedDescription))
	} else if task.Status != New {
		t.Errorf("Status was incorrect, got: %s, want: %s.", task.Status, New)
	} else if task.CreatedAt != expectedCreatedAt {
		t.Errorf("CreatedAt was incorrect, got: %s, want: %s.", task.CreatedAt, expectedCreatedAt)
	} else if task.UpdatedAt != expectedCreatedAt {
		t.Errorf("UpdatedAt was incorrect, got: %s, want: %s.", task.UpdatedAt, expectedCreatedAt)
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
	} else if task.Id != 0 {
		t.Errorf("Id was incorrect, got: %d, want: %d.", task.Id, 0)
	} else if task.Title != expectedTitle {
		t.Errorf("Title was incorrect, got: %s, want: %s.", task.Title, expectedTitle)
	} else if task.Description != expectedDescription {
		t.Errorf("Description was incorrect, got: %s, want: %s.", dereferencePointerToString(task.Description), dereferencePointerToString(expectedDescription))
	} else if task.Status != New {
		t.Errorf("Status was incorrect, got: %s, want: %s.", task.Status, New)
	} else if task.CreatedAt != expectedCreatedAt {
		t.Errorf("CreatedAt was incorrect, got: %s, want: %s.", task.CreatedAt, expectedCreatedAt)
	} else if task.UpdatedAt != expectedCreatedAt {
		t.Errorf("UpdatedAt was incorrect, got: %s, want: %s.", task.UpdatedAt, expectedCreatedAt)
	}
}

func TestUpdate1(t *testing.T) {
	// Arrange
	expectedStatus := InProgress
	expectedUpdatedAt := time.Now()
	
	description := "AnyDescription123"
	task := NewTask("AnyTitle123", &description, time.Now())

	// Act
	actualError := task.Update(expectedUpdatedAt)

	// Assert
	if actualError != nil {
		t.Errorf("Error was incorrect, got: %s, want: %s.", actualError.Error(), "nil")
	} else if task.Status != expectedStatus {
		t.Errorf("Status was incorrect, got: %s, want: %s.", task.Status, expectedStatus)
	} else if task.UpdatedAt != expectedUpdatedAt {
		t.Errorf("UpdatedAt was incorrect, got: %s, want: %s.", task.UpdatedAt, expectedUpdatedAt)
	}
}

func TestUpdate2(t *testing.T) {
	// Arrange
	expectedStatus := Done
	expectedUpdatedAt := time.Now()
	
	description := "AnyDescription123"
	task := NewTask("AnyTitle123", &description, time.Now())

	task.Update(expectedUpdatedAt.Add(time.Duration(-2) * time.Hour))

	// Act
	actualError := task.Update(expectedUpdatedAt)

	// Assert
	if actualError != nil {
		t.Errorf("Error was incorrect, got: %s, want: %s.", actualError.Error(), "nil")
	} else if task.Status != expectedStatus {
		t.Errorf("Status was incorrect, got: %s, want: %s.", task.Status, expectedStatus)
	} else if task.UpdatedAt != expectedUpdatedAt {
		t.Errorf("UpdatedAt was incorrect, got: %s, want: %s.", task.UpdatedAt, expectedUpdatedAt)
	}
}

func TestUpdate3(t *testing.T) {
	// Arrange
	expectedStatus := Done
	expectedUpdatedAt := time.Now()
	
	description := "AnyDescription123"
	task := NewTask("AnyTitle123", &description, time.Now())

	task.Update(expectedUpdatedAt.Add(time.Duration(-2) * time.Hour))
	task.Update(expectedUpdatedAt)

	// Act
	actualError := task.Update(expectedUpdatedAt.Add(time.Duration(2) * time.Hour))

	// Assert
	if actualError == nil {
		t.Errorf("Error was incorrect, got: %s, want: %s.", "nil", "not nil")
	} else if task.Status != expectedStatus {
		t.Errorf("Status was incorrect, got: %s, want: %s.", task.Status, expectedStatus)
	} else if task.UpdatedAt != expectedUpdatedAt {
		t.Errorf("UpdatedAt was incorrect, got: %s, want: %s.", task.UpdatedAt, expectedUpdatedAt)
	}
}

func dereferencePointerToString(pointerToString *string) string {
	if pointerToString == nil {
		return "nil"
	}

	return *pointerToString
}