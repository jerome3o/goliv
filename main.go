package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type note struct {
	ID   string   `json:"id"`
	Text string   `json:"text"`
	Tags []string `json:"tags"`
	User string   `json:"user"`
}

type idCounter struct {
	m       sync.Mutex
	counter int
}

var notes = []note{
	{
		ID:   "1",
		Text: "I want to make an effort to talk to my family more",
		Tags: []string{"family", "goal"},
		User: "Jerome",
	},
	{
		ID:   "2",
		Text: "Listen to Philosophize This! Philosophy is awesome",
		Tags: []string{"philosophy", "goal", "podcasts"},
		User: "Jerome",
	},
	{
		ID:   "3",
		Text: "Do washing this weekend",
		Tags: []string{"important", "task", "soon"},
		User: "Jerome",
	},
}

var counter idCounter = idCounter{counter: len(notes) + 1}

func getNextId() string {
	counter.m.Lock()
	defer counter.m.Unlock()

	nextId := counter.counter
	counter.counter += 1

	return fmt.Sprintf("%v", nextId)
}

func main() {
	router := gin.Default()
	router.GET("/notes", getNotes)
	router.GET("/notes/:id", getNoteByID)
	router.POST("/notes", postNotes)

	router.Run("localhost:8080")
}

// getNotes resposnds with the list of all Notes as JSON.
func getNotes(c *gin.Context) {
	requiredTags, any := c.GetQueryArray("tag")
	if !any {
		c.IndentedJSON(http.StatusOK, notes)
		return
	}

	filteredNotes := []note{}

	for _, n := range notes {
		if hasAllTags(n.Tags, requiredTags) {
			filteredNotes = append(filteredNotes, n)
		}
	}

	c.IndentedJSON(http.StatusOK, filteredNotes)
}

func hasAllTags(tags, requiredTags []string) bool {
	for _, t := range requiredTags {
		if !hasTag(tags, t) {
			return false
		}
	}
	return true
}

func hasTag(tags []string, tag string) bool {
	for _, t := range tags {
		if tag == t {
			return true
		}
	}
	return false
}

// postNotes adds an note from JSON received in the request body
func postNotes(c *gin.Context) {
	var newNote note

	// Call BindJSON to bind the received JSON to newnote
	if err := c.BindJSON(&newNote); err != nil {
		return
	}

	newNote.ID = getNextId()
	notes = append(notes, newNote)
	c.IndentedJSON(http.StatusCreated, newNote)
}

// getNoteByID locates the note whose ID value matches the id
// parameter sent by the client, then returns that note as a response.
func getNoteByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range notes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "note not found"})
}
