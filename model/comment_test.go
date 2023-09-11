package model

import (
	"testing"
)

func TestCommentDao(t *testing.T) {
	// db, err := gorm.Open("sqlite3", "test.db")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Comment{})

	// Create some comments
	comment1 := Comment{Content: "This is a great post!", Likes: 10, Reply: "", IP: "192.168.1.1", CommentCount: 0}
	comment2 := Comment{Content: "I agree with you!", Likes: 20, Reply: "", IP: "192.168.1.2", CommentCount: 0}

	// Insert comments into the database
	if err := NewCommentDaoImpl(db).InsertComment(comment1); err != nil {
		t.Fatalf("Failed to insert comment: %v", err)
	}
	if err := NewCommentDaoImpl(db).InsertComment(comment2); err != nil {
		t.Fatalf("Failed to insert comment: %v", err)
	}

	// Get all comments from the database
	comments, err := NewCommentDaoImpl(db).GetComments()
	if err != nil {
		t.Fatalf("Failed to get comments: %v", err)
	}
	if len(comments) != 2 {
		t.Fatalf("Expected 2 comments, got %d", len(comments))
	}

	// Delete a comment from the database
	if err := NewCommentDaoImpl(db).DeleteComment(comment1.ID); err != nil {
		t.Fatalf("Failed to delete comment: %v", err)
	}

	// Get all comments from the database again to verify the deletion
	comments, err = NewCommentDaoImpl(db).GetComments()
	if err != nil {
		t.Fatalf("Failed to get comments: %v", err)
	}
	if len(comments) != 1 {
		t.Fatalf("Expected 1 comment, got %d", len(comments))
	}
}
