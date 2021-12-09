# goliv
A little service to help Liv with her notes project.

This is just for testing / playing around, so there is no persistent storage. This means every time you restart the 
service the notes will be reset.

## Caveats / Notes

Some massive short cuts taken just to get you started

* The User field does nothing at this point (no auth, filtering, or anything)
* A "time added" or "time updated" field is needed
* An update end point is needed
* A delete endpoint is needed

## Data Model

```go
type note struct {
	ID   string   `json:"id"`
	Text string   `json:"text"`
	Tags []string `json:"tags"`
	User string   `json:"user"`
}
```

## Endpoints

All endpoints return data as JSON

* GET `/notes`
  * Returns all notes
  * The `tag` query parameter can be used to filter notes, i.e. `/notes?tag=goal&tag=family`
* GET `/notes/:id`
  * Gets an individual note by it's ID
* POST `/notes`
  * Add a note, body of request should be JSON note payload
  * The server assigns an ID, and the value (if provided) is ignored

## Setup

* Install [go](https://go.dev/doc/install)
* Open a shell in this directory
* Run `go run .`