package handlers

import (
	"github.com/gin-gonic/gin"
	"mrm/ent"
	"mrm/ent/room"
	"net/http"
	"strconv"
)

func (h *Handler) GetRooms(c *gin.Context) {
	rooms, err := h.DbClient.Room.Query().Order(ent.Asc(room.FieldID)).All(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rooms)
}

func (h *Handler) CreateRoom(c *gin.Context) {
	room := ent.Room{}
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if room.ID == 0 {
		var err error
		room.ID, err = h.DbClient.Room.Query().Aggregate(ent.Max("id")).Int(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		room.ID += 1
	}

	newRoom, err := h.DbClient.Room.Create().
		SetID(room.ID).
		SetName(room.Name).
		Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newRoom)
}

func (h *Handler) UpdateRoom(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	room := ent.Room{}
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	updatedRoom, err := h.DbClient.Room.UpdateOneID(id).
		SetName(room.Name).
		Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedRoom)
}

func (h *Handler) DeleteRoom(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.DbClient.Room.DeleteOneID(id).Exec(c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
