package tests

import (
	"strconv"
	"testing"

	. "treehole_next/models"

	"github.com/stretchr/testify/assert"
)

const testNotificationUserID = 9999

func ensureTestNotificationUser(t *testing.T) {
	t.Helper()

	user := User{
		ID: testNotificationUserID,
		Config: UserConfig{
			Notify: []string{"mention", "favorite", "report"},
		},
	}

	err := DB.Where("id = ?", testNotificationUserID).Assign(user).FirstOrCreate(&user).Error
	assert.NoError(t, err)
}

func createReplyNotificationForFloor(t *testing.T) (Hole, Floor, Message) {
	t.Helper()

	ensureTestNotificationUser(t)

	hole := Hole{
		DivisionID: 1,
		UserID:     testNotificationUserID,
	}
	err := DB.Create(&hole).Error
	assert.NoError(t, err)

	floor := Floor{
		HoleID:  hole.ID,
		Content: "cascade test floor",
		UserID:  1,
	}
	err = DB.Create(&floor).Error
	assert.NoError(t, err)

	notification := floor.SendReply(DB)
	_, err = notification.Send()
	assert.NoError(t, err)

	var msg Message
	err = DB.Where("related_floor_id = ? AND type = ?", floor.ID, MessageTypeReply).
		Order("id DESC").
		First(&msg).Error
	assert.NoError(t, err)
	assert.NotNil(t, msg.RelatedFloorID)
	assert.EqualValues(t, floor.ID, *msg.RelatedFloorID)

	return hole, floor, msg
}

func assertMessageExists(t *testing.T, messageID int, shouldExist bool) {
	t.Helper()

	var count int64
	err := DB.Model(&Message{}).Where("id = ?", messageID).Count(&count).Error
	assert.NoError(t, err)

	if shouldExist {
		assert.EqualValues(t, 1, count, "message should exist")
		return
	}

	assert.EqualValues(t, 0, count, "message should not exist")
}

func assertMessageUserExists(t *testing.T, messageID int, shouldExist bool) {
	t.Helper()

	var count int64
	err := DB.Model(&MessageUser{}).Where("message_id = ?", messageID).Count(&count).Error
	assert.NoError(t, err)

	if shouldExist {
		assert.Greater(t, count, int64(0), "message_user should exist")
		return
	}

	assert.EqualValues(t, 0, count, "message_user should not exist")
}

func TestNotificationRelatedFloorIDPersisted(t *testing.T) {
	_, floor, msg := createReplyNotificationForFloor(t)

	assert.NotZero(t, msg.ID)
	assert.NotNil(t, msg.RelatedFloorID)
	assert.EqualValues(t, floor.ID, *msg.RelatedFloorID)
	assertMessageUserExists(t, msg.ID, true)
}

func TestDeleteFloorCascadeNotification(t *testing.T) {
	_, floor, msg := createReplyNotificationForFloor(t)

	assertMessageExists(t, msg.ID, true)
	assertMessageUserExists(t, msg.ID, true)

	data := Map{"delete_reason": "cascade test"}
	testAPI(t, "delete", "/api/floors/"+strconv.Itoa(floor.ID), 200, data)

	assertMessageExists(t, msg.ID, false)
	assertMessageUserExists(t, msg.ID, false)
}

func TestHideHoleCascadeNotification(t *testing.T) {
	hole, _, msg := createReplyNotificationForFloor(t)

	assertMessageExists(t, msg.ID, true)
	assertMessageUserExists(t, msg.ID, true)

	testCommon(t, "delete", "/api/holes/"+strconv.Itoa(hole.ID), 204)

	assertMessageExists(t, msg.ID, false)
	assertMessageUserExists(t, msg.ID, false)
}

func TestForceDeleteHoleCascadeNotification(t *testing.T) {
	hole, _, msg := createReplyNotificationForFloor(t)

	assertMessageExists(t, msg.ID, true)
	assertMessageUserExists(t, msg.ID, true)

	testCommon(t, "delete", "/api/holes/"+strconv.Itoa(hole.ID)+"/_force", 204)

	assertMessageExists(t, msg.ID, false)
	assertMessageUserExists(t, msg.ID, false)
}
