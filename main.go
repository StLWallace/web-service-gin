package main

import (
	"context"
	"example/web-service-gin/pkg/firestore/models"
	"net/http"

	"os"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func getUnit(collection *firestore.CollectionRef, ctx context.Context) gin.HandlerFunc {
	// Creates a handler function to get a unit based on the unique ID
	fn := func(c *gin.Context) {
		// Takes path parameter id and gets the corresponding document
		id := c.Param("id")

		unit := collection.Doc(id)

		docsnap, _ := unit.Get(ctx)

		var unitData models.UnitDetail
		if err := docsnap.DataTo(&unitData); err != nil {
			print(err)
		}

		c.IndentedJSON(http.StatusOK, unitData)
	}
	return fn
}

func postUnit(collection *firestore.CollectionRef, ctx context.Context) gin.HandlerFunc {
	// Creates a handler function to post a new unit document
	// Args:
	// collection - a Firestore collection object that will be written to
	fn := func(c *gin.Context) {
		// Creates a UnitDetail object from the posted data and creates a new document in Firestore with the data
		// The new document is given a random ID
		var newUnit models.UnitDetail

		if err := c.BindJSON(&newUnit); err != nil {
			return
		}

		wr, _, err := collection.Add(ctx, newUnit)
		if err != nil {
			print(err)
		} else {
			print(wr)
		}

		c.IndentedJSON(http.StatusCreated, newUnit)

	}
	return fn
}

func main() {
	// Creates a service with the GET and POST endpoints defined above
	// Requires setting env vars for
	// GCP_PROJECT - Your Google Cloud project id
	// COLLECTION_NAME - the name of the Firestore collection you want to use
	router := gin.Default()
	ctx := context.Background()
	client, _ := firestore.NewClient(ctx, os.Getenv("GCP_PROJECT"))
	bankers := client.Collection(os.Getenv("COLLECTION_NAME"))

	router.GET("/units/:id", getUnit(bankers, ctx))
	router.POST(("units"), postUnit(bankers, ctx))

	router.Run("localhost:8000")
}
