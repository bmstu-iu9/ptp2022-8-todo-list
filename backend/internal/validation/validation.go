package validation

import (
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"regexp"
)

// ValidateField performs basic validation on string field.
func ValidateField(field string, minLen, maxLen int, regex string) bool {
	if matched, _ := regexp.MatchString(regex, field); !matched {
		return false
	}
	if len(field) < minLen || len(field) > maxLen {
		return false
	}
	return true
}

// ValidateItemState performs basic validation on states filters.
func ValidateItemState(states []entity.ItemState) bool {
	for _, state := range states {
		switch state {
		case entity.Store:
		case entity.Inventoried:
		case entity.Equipped:
		default:
			return false
		}
	}
	return true
}

// ValidateItemRarity performs basic validation on rarities filters.
func ValidateItemRarity(rarities []string) bool {
	for _, rarity := range rarities {
		switch rarity {
		case "common":
		case "rare":
		case "legendary":
		case "epic":
		default:
			return false
		}
	}
	return true
}

// ValidateItemCategory performs basic validation on categories filters.
func ValidateItemCategory(categories []string) bool {
	for _, category := range categories {
		switch category {
		case "pet":
		case "skin":
		case "weapon":
		case "armor":
		default:
			return false
		}
	}
	return true
}
