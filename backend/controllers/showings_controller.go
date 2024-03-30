package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"log"
	"math/rand"
	"net/http"
	"piaccho/cinema-api/configs"
	"piaccho/cinema-api/models"
	"piaccho/cinema-api/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var showingCollection = configs.GetCollection("showings")

func GetAllShowings() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var showings []models.Showing
		defer cancel()

		results, err := showingCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// reading from the db in an optimal way
		defer func(results *mongo.Cursor, ctx context.Context) {
			err := results.Close(ctx)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}
		}(results, ctx)

		for results.Next(ctx) {
			var singleShowing models.Showing
			if err = results.Decode(&singleShowing); err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			showings = append(showings, singleShowing)
		}

		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": showings}},
		)
	}
}

func CreateShowing() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var showing models.Showing
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&showing); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// use the validator library to validate required fields
		if validationErr := utils.Validator.Struct(&showing); validationErr != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newShowing := models.Showing{
			Id:              primitive.NewObjectID(),
			MovieShowingRef: showing.MovieShowingRef,
			HallId:          showing.HallId,
			StartTime:       showing.StartTime,
			EndTime:         showing.EndTime,
			AvailableSeats:  showing.AvailableSeats,
			BookedSeats:     showing.BookedSeats,
			PricePerTicket:  showing.PricePerTicket,
			AudioType:       showing.AudioType,
			VideoType:       showing.VideoType,
		}

		result, err := showingCollection.InsertOne(ctx, newShowing)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, models.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetShowingById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		showingId := c.Param("showingId")
		var showing models.Showing
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(showingId)

		err := showingCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&showing)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				c.JSON(http.StatusNotFound, models.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "No showing found with the provided ID"}})
				return
			}
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": showing}})
	}
}

func UpdateShowing() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		showingId := c.Param("showingId")
		var showing models.Showing
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(showingId)

		// validate the request body
		if err := c.BindJSON(&showing); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// use the validator library to validate required fields
		if validationErr := utils.Validator.Struct(&showing); validationErr != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"movie_id": showing.MovieShowingRef, "hall_id": showing.HallId, "start_time": showing.StartTime, "end_time": showing.EndTime, "available_seats": showing.AvailableSeats, "booked_seats": showing.BookedSeats, "price_per_ticket": showing.PricePerTicket, "audio_type": showing.AudioType, "video_type": showing.VideoType}
		result, err := showingCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// get updated showing details
		var updatedShowing models.Showing
		if result.MatchedCount == 1 {
			err := showingCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedShowing)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedShowing}})
	}
}

func DeleteShowing() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		showingId := c.Param("showingId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(showingId)

		result, err := showingCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				models.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Showing with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Showing successfully deleted!"}},
		)
	}
}

func GetShowingsByDate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse datetime from the URL parameter
		datetimeStr := c.Param("datetime")
		datetime, err := time.Parse(time.RFC3339, datetimeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid datetime format. Please use RFC3339."})
			return
		}

		// Find all documents in the showingCollection that have the given datetime
		cursor, err := showingCollection.Find(context.Background(), bson.D{{"datetime", datetime}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Decode the documents
		var showings []models.Showing
		if err = cursor.All(context.Background(), &showings); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the documents as JSON
		c.JSON(http.StatusOK, showings)
	}
}

func GetShowingsByMovieId() gin.HandlerFunc {
	return func(c *gin.Context) {
		movieId := c.Param("movieId")

		// Find all documents in the showingCollection that contain the movie with the given id
		cursor, err := showingCollection.Find(context.Background(), bson.D{{"movie._id", movieId}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Decode the documents
		var showings []models.Showing
		if err = cursor.All(context.Background(), &showings); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the documents as JSON
		c.JSON(http.StatusOK, showings)
	}
}

func GetShowingsByHallId() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		hallId := c.Param("hallId")
		defer cancel()

		// Find all documents in the showingCollection that contain the hall with the given id
		cursor, err := showingCollection.Find(ctx, bson.D{{"hall._id", hallId}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Decode the documents
		var showings []models.Showing
		if err = cursor.All(ctx, &showings); err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Return the documents as JSON
		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": showings}})
	}
}

func GenerateShowingsForNextDays() gin.HandlerFunc {
	return func(c *gin.Context) {
		const MinTimeHour = 8
		const MaxTimeHour = 22
		const HourMinutes = 60
		const QuarterMinutes = 15
		const QuarterRound = 14

		daysNum, err := strconv.Atoi(c.Param("daysNumber"))
		if err != nil || daysNum < 1 {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Invalid number of days"}})
			return
		}

		allMoviesRefs, err := getAllMoviesRefs()
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"description": "Failed to get all movies", "data": err.Error()}})
			return
		}
		allHalls, err := getAllHalls()
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"description": "Failed to get all halls", "data": err.Error()}})
			return
		}

		// Start generating showings from the next day
		//startDay := time.Now().AddDate(0, 0, 1)

		// Start generating showings from the current day
		startDay := time.Now()

		// Start a new session for the transaction
		session, err := configs.DB.StartSession()
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"description": "Failed to start session", "data": err.Error()}})
			return
		}
		defer session.EndSession(context.Background())

		// Start the transaction
		err = session.StartTransaction()
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"description": "Failed to start transaction", "data": err.Error()}})
			return
		}

		// Create a SessionContext
		ctx := mongo.NewSessionContext(context.Background(), session)

		for d := 0; d < daysNum; d++ {
			for _, hall := range allHalls {
				initTime := MinTimeHour
				for initTime < MaxTimeHour {
					randomMovie := allMoviesRefs[rand.Intn(len(allMoviesRefs))]
					startDate := time.Date(startDay.Year(), startDay.Month(), startDay.Day()+d, initTime, 0, 0, 0, time.UTC)
					endDate := startDate.Add(time.Minute * time.Duration(randomMovie.Length))

					// Create showing in hall at time
					err := createShowingInHallAtTime(ctx, randomMovie, hall, startDate, endDate)
					if err != nil {
						// Abort the transaction in case of an error
						_ = session.AbortTransaction(context.Background())
						c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"description": "Failed to create showing in hall at time", "data": err.Error()}})
						return
					}

					initTime = (initTime + randomMovie.Length + QuarterRound) / QuarterMinutes * QuarterMinutes // Round to the nearest quarter-hour

					if initTime >= HourMinutes {
						initTime = initTime % HourMinutes
					}
				}
			}
		}

		// Commit the transaction
		err = session.CommitTransaction(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"description": "Failed to commit transaction", "data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Showings successfully generated"}})
	}
}

var baseURL = "http://localhost:" + configs.EnvPort()

func getAllMoviesRefs() ([]models.MovieRef, error) {
	resp, err := http.Get(baseURL + "/api/movies/ref")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	var response models.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	data, ok := response.Data["data"]
	if !ok {
		return nil, errors.New("field 'data' not found")
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var moviesRefs []models.MovieRef
	err = json.Unmarshal(dataBytes, &moviesRefs)
	if err != nil {
		return nil, err
	}

	return moviesRefs, nil
}

func getAllHalls() ([]models.Hall, error) {
	resp, err := http.Get(baseURL + "/api/halls/")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	var response models.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	data, ok := response.Data["data"]
	if !ok {
		return nil, errors.New("field 'data' not found")
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var halls []models.Hall
	err = json.Unmarshal(dataBytes, &halls)
	if err != nil {
		return nil, err
	}

	return halls, nil
}

func createShowingInHallAtTime(ctx mongo.SessionContext, movie models.MovieRef, hall models.Hall, startDate, endDate time.Time) error {
	var audioTypes = []string{"Dubbing", "Subtitles", "VoiceOver"}
	var videoTypes = []string{"2D", "3D"}
	const InitBookedSeatsNumber = 0
	const MinTicketPrice = 5.00
	const MaxTicketPrice = 10.00

	// Random audio and video type
	audioType := audioTypes[rand.Intn(len(audioTypes))]
	videoType := videoTypes[rand.Intn(len(videoTypes))]

	// Number of available seats equals all possible in the hall
	availableSeats := hall.Rows * hall.SeatsPerRow

	// Zero booked seats

	bookedSeats := InitBookedSeatsNumber

	// Random ticket price from 5.00 to 10.00
	pricePerTicket := MinTicketPrice + rand.Float64()*(MaxTicketPrice-MinTicketPrice)

	// Create new showing
	showing := models.Showing{
		Id:              primitive.NewObjectID(),
		MovieShowingRef: movie,
		HallId:          hall.Id,
		StartTime:       startDate,
		EndTime:         endDate,
		AvailableSeats:  availableSeats,
		BookedSeats:     bookedSeats,
		PricePerTicket:  pricePerTicket,
		AudioType:       audioType,
		VideoType:       videoType,
	}

	// Save showing to the database
	_, err := showingCollection.InsertOne(ctx, showing)
	if err != nil {
		return err
	}
	return nil
}
