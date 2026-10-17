package primitive

import (
	"reflect"
	"testing"
	"time"
)

func TestDatedInitializesTimestamps(t *testing.T) {
	datable := Dated()

	if datable.GetcreatedAt().IsZero() || datable.GetupdatedAt().IsZero() {
		t.Fatal("Dated returned zero timestamps")
	}
	if datable.GetcreatedAt().Location() != time.UTC || datable.GetupdatedAt().Location() != time.UTC {
		t.Fatal("Dated timestamps are not in UTC")
	}
	if !datable.GetcreatedAt().Equal(datable.GetupdatedAt()) {
		t.Fatal("Dated timestamps are not initialized to the same instant")
	}
}

func TestDatedTouchUpdatesOnlyUpdatedAt(t *testing.T) {
	datable := Dated()
	createdAt := datable.GetcreatedAt()
	updatedAt := datable.GetupdatedAt()
	datable.UpdatedAt = updatedAt.Add(-time.Second)
	beforeTouch := datable.GetupdatedAt()

	datable.Touch()

	if !datable.GetcreatedAt().Equal(createdAt) {
		t.Fatal("Touch changed CreatedAt")
	}
	if !datable.GetupdatedAt().After(beforeTouch) {
		t.Fatal("Touch did not advance UpdatedAt")
	}
}

func TestDatableHasSerializationTags(t *testing.T) {
	for tagName, want := range map[string]map[string]string{
		"bun": {
			"CreatedAt": "created_at,notnull,default:current_timestamp",
			"UpdatedAt": "updated_at,notnull,default:current_timestamp",
			"immutable": "-",
		},
		"bson": {
			"CreatedAt": "createdAt",
			"UpdatedAt": "updatedAt",
			"immutable": "-",
		},
		"yaml": {
			"CreatedAt": "createdAt",
			"UpdatedAt": "updatedAt",
			"immutable": "-",
		},
		"json": {
			"CreatedAt": "createdAt",
			"UpdatedAt": "updatedAt",
			"immutable": "-",
		},
	} {
		t.Run(tagName, func(t *testing.T) {
			assertDatableTags(t, tagName, want)
		})
	}
}

func assertDatableTags(t *testing.T, tagName string, want map[string]string) {
	t.Helper()

	typeOfDatable := reflect.TypeOf(Datable{})
	for fieldName, expectedTag := range want {
		field, ok := typeOfDatable.FieldByName(fieldName)
		if !ok {
			t.Fatalf("Datable is missing field %q", fieldName)
		}
		if got := field.Tag.Get(tagName); got != expectedTag {
			t.Errorf("Datable.%s %s tag = %q, want %q", fieldName, tagName, got, expectedTag)
		}
	}
}

func TestGetAgePreservesDatableState(t *testing.T) {
	datable := Steady()

	age := datable.GetAge()

	if !age.GetcreatedAt().Equal(datable.GetcreatedAt()) {
		t.Fatal("GetAge changed CreatedAt")
	}
	if !age.GetupdatedAt().Equal(datable.GetupdatedAt()) {
		t.Fatal("GetAge changed UpdatedAt")
	}
	if !age.immutable {
		t.Fatal("GetAge changed immutability")
	}
}

func TestSteadyTouchDoesNotUpdateTimestamp(t *testing.T) {
	datable := Steady()
	updatedAt := datable.GetupdatedAt()

	datable.Touch()

	if !datable.GetupdatedAt().Equal(updatedAt) {
		t.Fatal("Touch changed the timestamp of a steady Datable")
	}
}
