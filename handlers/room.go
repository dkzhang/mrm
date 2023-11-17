package handlers

import (
	"github.com/gin-gonic/gin"
	"mrm/ent"
	"net/http"
	"strconv"
)

func (h *Handler) GetRooms(c *gin.Context) {
	rooms := h.DbClient.Room.Query().AllX(c)
	c.JSON(http.StatusOK, rooms)
}

func (h *Handler) CreateRoom(c *gin.Context) {
	room := ent.Room{}
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newRoom, err := h.DbClient.Room.Create().
		SetID(room.ID).
		SetName(room.Name).
		Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newRoom)
}

func (h *Handler) UpdateRoom(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room := ent.Room{}
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedRoom, err := h.DbClient.Room.UpdateOneID(id).
		SetName(room.Name).
		Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedRoom)
}

func (h *Handler) DeleteRoom(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DbClient.Room.DeleteOneID(id).Exec(c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
