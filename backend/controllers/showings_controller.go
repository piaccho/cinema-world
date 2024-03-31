package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"log"
	"math/rand"
	"net/http"
	"piaccho/cinema-api/configs"
	"piaccho/cinema-api/models"
	"piaccho/cinema-api/utils"
	"reflect"
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

		// pass through the validator
		err, errDescription := validateShowing(showing, ctx)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": errDescription}})
			return
		}

		newShowing := models.Showing{
			Id:              primitive.NewObjectID(),
			MovieShowingRef: showing.MovieShowingRef,
			Hall:            showing.Hall,
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

func UpdateShowing() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		showingId := c.Param("showingId")
		var showing models.Showing
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(showingId)

		// validate the request body
		if err := c.BindJSON(&showing); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to bind JSON"}})
			return
		}

		// pass through the validator
		if err, errDescription := validateShowing(showing, ctx); errDescription != "" {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": errDescription}})
			return
		}

		update := bson.M{"movie_id": showing.MovieShowingRef, "hall": showing.Hall, "start_time": showing.StartTime, "end_time": showing.EndTime, "available_seats": showing.AvailableSeats, "booked_seats": showing.BookedSeats, "price_per_ticket": showing.PricePerTicket, "audio_type": showing.AudioType, "video_type": showing.VideoType}
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

func validateShowing(showing models.Showing, ctx context.Context) (error, string) {

	// use the validator library to validate required fields
	if validationErr := utils.Validator.Struct(&showing); validationErr != nil {
		return validationErr, "error caught Validator library"
	}
	// check if hall exists and if its name is correct one with the provided ID
	hallCollection := configs.GetCollection("halls")
	var hall models.Hall
	err := hallCollection.FindOne(ctx, bson.M{"_id": showing.Hall.Id}).Decode(&hall)
	if err != nil {
		return err, "Hall with the provided ID does not exist"
	}
	if !reflect.DeepEqual(hall, showing.Hall) {
		return errors.New("not Deep Equal"), "Hall document is not identical with the provided one"
	}
	// check if movie exists and if its name is correct one with the provided ID
	movieCollection := configs.GetCollection("movies")
	var movie models.MovieRef
	err = movieCollection.FindOne(ctx, bson.M{"_id": showing.MovieShowingRef.Id}).Decode(&movie)
	if err != nil {
		return err, "Movie with the provided ID does not exist"
	}
	if !reflect.DeepEqual(movie, showing.MovieShowingRef) {
		return errors.New("not Deep Equal"), "Movie document is not identical with the provided one"
	}

	if showing.StartTime.After(showing.EndTime) ||
		showing.StartTime.Equal(showing.EndTime) ||
		showing.StartTime.Before(time.Now()) ||
		showing.EndTime.Before(time.Now()) ||
		showing.StartTime.Equal(time.Now()) ||
		showing.EndTime.Equal(time.Now()) {
		return errors.New("incorrect time range"), "Incorrect time range"
	}
	// check if showing available seats are correct
	if showing.AvailableSeats != showing.Hall.Rows*showing.Hall.SeatsPerRow ||
		showing.BookedSeats >= showing.AvailableSeats {
		return errors.New("incorrect number of seats"), "Incorrect number of seats"
	}

	return nil, ""
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
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		dateParam := c.Param("datetime")
		datetime, err := time.Parse("2006-01-02", dateParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to parse datetime"}})
			return
		}

		startOfDay := time.Date(datetime.Year(), datetime.Month(), datetime.Day(), 0, 0, 0, 0, datetime.Location())
		endOfDay := startOfDay.Add(24 * time.Hour)
		filter := bson.M{"start_time": bson.M{"$gte": startOfDay, "$lt": endOfDay}}

		cursor, err := showingCollection.Find(ctx, filter)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				c.JSON(http.StatusNotFound, models.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "No showing found with the provided date"}})
				return
			}
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to find showing with the provided date"}})
			return
		}
		defer func(cursor *mongo.Cursor, ctx context.Context) {
			err := cursor.Close(ctx)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to close cursor"}})
			}
		}(cursor, ctx)

		// Decode the found showings
		var showings []models.Showing
		if err = cursor.All(ctx, &showings); err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to decode showings"}})
			return
		}

		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"info": "Fetched " + strconv.Itoa(len(showings)) + " showings documents", "data": showings}})
	}
}

func GetShowingsByMovieId() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		movieId := c.Param("movieId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(movieId)

		cursor, err := showingCollection.Find(ctx, bson.M{"movie._id": objId})
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				c.JSON(http.StatusNotFound, models.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "No showing found with the provided movie ID"}})
				return
			}
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to find showing with the provided movie ID"}})
			return
		}
		defer func(cursor *mongo.Cursor, ctx context.Context) {
			err := cursor.Close(ctx)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to close cursor"}})
			}
		}(cursor, ctx)

		// Decode the found showings
		var showings []models.Showing
		if err = cursor.All(ctx, &showings); err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to decode showings"}})
			return
		}

		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"info": "Fetched " + strconv.Itoa(len(showings)) + " showings documents", "data": showings}})
	}
}

func GetShowingsByHallId() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		hallId := c.Param("hallId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(hallId)

		cursor, err := showingCollection.Find(ctx, bson.M{"hall._id": objId})
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				c.JSON(http.StatusNotFound, models.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "No showing found with the provided hall ID"}})
				return
			}
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to find showing with the provided hall ID"}})
			return
		}
		defer func(cursor *mongo.Cursor, ctx context.Context) {
			err := cursor.Close(ctx)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to close cursor"}})
			}
		}(cursor, ctx)

		// Decode the found showings
		var showings []models.Showing
		if err = cursor.All(ctx, &showings); err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to decode halls"}})
			return
		}

		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"info": "Fetched " + strconv.Itoa(len(showings)) + " halls documents", "data": showings}})
	}
}

func GenerateShowingsForDays() gin.HandlerFunc {
	return func(c *gin.Context) {
		const MinTimeHour = 14 // cinema opening time - 2 PM
		const MaxTimeHour = 22 // cinema closing time - 10 PM
		const HourMinutes = 60
		const QuarterMinutes = 15
		const QuarterRound = 14
		const BreakMinutes = 15
		const MovieBatchSize = 4

		var createdShowingsNumber = 0

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

		fmt.Println("\nStart generating showings for next " + strconv.Itoa(daysNum) + " days...\n")
		for d := 0; d < daysNum; d++ {
			// Movies chosen by window slide of movies array, moved each day by one movie
			var moviesBatch = allMoviesRefs[d : MovieBatchSize+d]
			fmt.Println("Day:", d, "\n")
			for _, hall := range allHalls {
				fmt.Println("\t", hall.Name, ":")
				var initTime = time.Date(startDay.Year(), startDay.Month(), startDay.Day()+d, MinTimeHour, 0, 0, 0, time.UTC)
				var finalTime = time.Date(startDay.Year(), startDay.Month(), startDay.Day()+d, MaxTimeHour, 0, 0, 0, time.UTC)
				for initTime.Before(finalTime) {
					randomMovie := moviesBatch[rand.Intn(len(moviesBatch))]
					startDate := initTime
					showingLength := (initTime.Minute() + randomMovie.Length + QuarterRound) / QuarterMinutes * QuarterMinutes // Round to the nearest quarter-hour
					endDate := startDate.Add(time.Minute * time.Duration(showingLength))

					// Create showing in hall at time
					err := createShowingInHallAtTime(c, randomMovie, hall, startDate, endDate)
					if err != nil {
						c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"description": "Failed to create showing in hall at time", "data": err.Error()}})
						return
					}
					createdShowingsNumber++

					initTime = endDate.Add(time.Minute * time.Duration(BreakMinutes))
					fmt.Println(
						"Showing from:",
						strconv.Itoa(startDate.Hour())+":"+strconv.Itoa(startDate.Minute()),
						"to:",
						strconv.Itoa(endDate.Hour())+":"+strconv.Itoa(endDate.Minute()))
				}
			}
		}

		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"info": "Created " + strconv.Itoa(createdShowingsNumber) + " showings", "data": "Showings successfully generated"}})
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

func createShowingInHallAtTime(ctx context.Context, movie models.MovieRef, hall models.Hall, startDate, endDate time.Time) error {
	var audioTypes = []string{"Dubbing", "Subtitles", "VoiceOver"}
	var videoTypes = []string{"2D", "3D"}
	var prices = []float64{7.00, 8.00, 10.00, 12.00}
	const InitBookedSeatsNumber = 0

	// Random audio and video type
	audioType := audioTypes[rand.Intn(len(audioTypes))]
	videoType := videoTypes[rand.Intn(len(videoTypes))]

	// Number of available seats equals all possible in the hall
	availableSeats := hall.Rows * hall.SeatsPerRow

	// Zero booked seats
	bookedSeats := InitBookedSeatsNumber

	// Random ticket price from 5.00 to 10.00
	pricePerTicket := prices[rand.Intn(len(prices))]

	// Create [][]Seat array of hall.Rows x hall.SeatsPerRow size with Seat struct
	seats := make([][]models.Seat, hall.Rows)
	for i := range seats {
		seats[i] = make([]models.Seat, hall.SeatsPerRow)
		for j := range seats[i] {
			seats[i][j] = models.Seat{
				RowNumber:  i + 1,
				SeatNumber: j + 1,
				IsReserved: false,
			}
		}
	}

	// Create new showing
	showing := models.Showing{
		Id:              primitive.NewObjectID(),
		MovieShowingRef: movie,
		Hall:            hall,
		StartTime:       startDate,
		EndTime:         endDate,
		Seats:           seats,
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
