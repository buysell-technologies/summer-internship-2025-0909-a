package handler

import (
	"testing"

	"github.com/buysell-technologies/summer-internship-2024-backend/api/domain/model"
)

// Test to verify that convertGenderForDisplay handles nil gender without panic
func TestConvertGenderForDisplay_NilGender(t *testing.T) {
	// Test with nil gender - should not panic
	result := convertGenderForDisplay(nil)
	if result != nil {
		t.Errorf("Expected nil when gender is nil, got %v", result)
	}
}

func TestConvertGenderForDisplay_ValidGenders(t *testing.T) {
	// Test with male
	male := "male"
	result := convertGenderForDisplay(&male)
	if result == nil || *result != "男性" {
		t.Errorf("Expected '男性' for male, got %v", result)
	}

	// Test with female
	female := "female"
	result = convertGenderForDisplay(&female)
	if result == nil || *result != "女性" {
		t.Errorf("Expected '女性' for female, got %v", result)
	}

	// Test with other value
	other := "other"
	result = convertGenderForDisplay(&other)
	if result == nil || *result != "other" {
		t.Errorf("Expected 'other' for other, got %v", result)
	}
}

func TestConvertUserForDisplay_NilGender(t *testing.T) {
	// Test with user that has nil gender
	user := &model.User{
		ID:     "test-id",
		Name:   "Test User",
		Email:  "test@example.com",
		Gender: nil, // This should not cause panic
	}

	result := convertUserForDisplay(user)
	if result == nil {
		t.Errorf("Expected user to be converted, got nil")
	}
	if result.Gender != nil {
		t.Errorf("Expected gender to remain nil, got %v", result.Gender)
	}
}

func TestConvertUsersForDisplay_MixedGenders(t *testing.T) {
	// Test with multiple users including nil genders
	male := "male"
	female := "female"
	users := []*model.User{
		{
			ID:     "user1",
			Name:   "User with nil gender",
			Gender: nil,
		},
		{
			ID:     "user2", 
			Name:   "Male user",
			Gender: &male,
		},
		{
			ID:     "user3",
			Name:   "Female user", 
			Gender: &female,
		},
	}

	result := convertUsersForDisplay(users)
	
	if len(result) != 3 {
		t.Errorf("Expected 3 users, got %d", len(result))
	}

	// Check nil gender user
	if result[0].Gender != nil {
		t.Errorf("Expected first user gender to be nil, got %v", result[0].Gender)
	}

	// Check male user
	if result[1].Gender == nil || *result[1].Gender != "男性" {
		t.Errorf("Expected second user gender to be '男性', got %v", result[1].Gender)
	}

	// Check female user
	if result[2].Gender == nil || *result[2].Gender != "女性" {
		t.Errorf("Expected third user gender to be '女性', got %v", result[2].Gender)
	}
}
