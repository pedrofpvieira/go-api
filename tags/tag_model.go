package tags

import (
	"encoding/json"
	"github.com/mongodb/mongo-go-driver/bson"
	"github/pedroportasvieira/go-api/db"
	"github/pedroportasvieira/go-api/errors"
	"net/http"
)

type Tag struct {
	Id  int    `json:"id,omitempty"`
	Tag string `json:"tag,omitempty"`
}

// Returns a json response with the available tags
func Tags(w http.ResponseWriter, r *http.Request) {

	var tagsEntity []bson.M

	results, ctx := db.GetMongoRecords("tags")

	for results.Next(ctx) {
		var result bson.M
		err := results.Decode(&result)
		errors.HandleError(err)

		tagsEntity = append(tagsEntity, result)
	}

	buildResponse(w, tagsEntity)
}

// Responsible to save the tag into db source
func SaveTag(w http.ResponseWriter, r *http.Request) {
	var tag Tag

	err := json.NewDecoder(r.Body).Decode(&tag)
	errors.HandleError(err)

	db.InsertMongoRecord("tags", bson.M{"Id": tag.Id, "Tag": tag.Tag})

	buildResponse(w, "Saved successfully!")
}

// Builds a valid response in json with status code 200
func buildResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}
