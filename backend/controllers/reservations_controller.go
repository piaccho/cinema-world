package controllers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"piaccho/cinema-api/configs"
	"piaccho/cinema-api/models"
	"piaccho/cinema-api/utils"
	"reflect"
	"time"
)

var reservationCollection = configs.GetCollection("reservations")

func CreateReservation() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var reservation models.Reservation

		// validate the request body
		if err := c.BindJSON(&reservation); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Validation error"}})
			return
		}

		// use the validator library to validate required fields
		if validationErr := utils.Validator.Struct(&reservation); validationErr != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error(), "description": "Validation error"}})
			return
		}

		// start a session
		session, err := configs.DB.StartSession()
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to start session"}})
			return
		}
		defer session.EndSession(ctx)

		// start a transaction
		err = session.StartTransaction()
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to start transaction"}})
			return
		}

		// pass through the validator
		err, errDescription := validateReservation(reservation, ctx)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": errDescription}})
			return
		}

		// get the showing
		var showing models.Showing
		err = showingCollection.FindOne(ctx, bson.M{"_id": reservation.ShowingId}).Decode(&showing)
		if err != nil {
			errAbort := session.AbortTransaction(ctx)
			if errAbort != nil {
				return
			}
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get showing"}})
			return
		}

		// update the seats
		for _, reservedSeat := range reservation.ReservedSeats {
			showing.Seats[reservedSeat.RowNumber][reservedSeat.SeatNumber].IsReserved = true
		}

		// update the available and booked seats
		showing.AvailableSeats -= len(reservation.ReservedSeats)
		showing.BookedSeats += len(reservation.ReservedSeats)

		// update the showing in the database
		_, err = showingCollection.UpdateOne(ctx, bson.M{"_id": reservation.ShowingId}, bson.M{"$set": showing})
		if err != nil {
			errAbort := session.AbortTransaction(ctx)
			if errAbort != nil {
				return
			}
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to update showing"}})
			return
		}

		newReservation := models.Reservation{
			ShowingId:       reservation.ShowingId,
			UserId:          reservation.UserId,
			MovieShowingRef: reservation.MovieShowingRef,
			ReservedSeats:   reservation.ReservedSeats,
			TotalPrice:      reservation.TotalPrice,
		}

		result, err := reservationCollection.InsertOne(ctx, newReservation)
		if err != nil {
			errAbort := session.AbortTransaction(ctx)
			if errAbort != nil {
				return
			}
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to create reservation"}})
			return
		}

		// commit the transaction
		err = session.CommitTransaction(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to commit transaction"}})
			return
		}

		c.JSON(http.StatusCreated, models.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result, "description": "Reservation created successfully"}})
	}
}

func GetReservationById() gin.HandlerFunc {

	//var users []models.User
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		reservationId := c.Param("reservationId")
		objId, _ := primitive.ObjectIDFromHex(reservationId)

		var reservation models.Reservation

		err := reservationCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&reservation)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				c.JSON(http.StatusNotFound, models.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "No reservation found with the provided ID"}})
				return
			}
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get reservation"}})
			return
		}

		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": reservation, "description": "Reservation found"}})
		//cursor, err := userCollection.Find(ctx, bson.M{})
		//if err != nil {
		//	c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get reservation"}})
		//
		//}
		//if err = cursor.All(ctx, &users); err != nil {
		//	c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get reservation"}})
		//}
		//
		//for _, user := range users {
		//	for _, reservation := range user.Reservations {
		//		if reservation.Id == objId {
		//			c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": reservation, "description": "Reservation found"}})
		//		}
		//	}
		//}
		//c.JSON(http.StatusNotFound, models.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": errors.New("reservation not found"), "description": "Reservation not found"}})
	}
}

func UpdateReservation() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		reservationId := c.Param("reservationId")
		objId, _ := primitive.ObjectIDFromHex(reservationId)
		var reservation models.Reservation

		// validate the request body
		if err := c.BindJSON(&reservation); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Validation error"}})
			return
		}

		// use the validator library to validate required fields
		if validationErr := utils.Validator.Struct(&reservation); validationErr != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error(), "description": "Validation error"}})
			return
		}

		// start a session
		session, err := configs.DB.StartSession()
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to start session"}})
			return
		}
		defer session.EndSession(ctx)

		// start a transaction
		err = session.StartTransaction()
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to start transaction"}})
			return
		}

		// pass through the validator
		err, errDescription := validateReservation(reservation, ctx)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": errDescription}})
			return
		}

		// get the reservation before updating to get reserved seats
		var oldReservation models.Reservation
		err = reservationCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&oldReservation)
		if err != nil {
			errAbort := session.AbortTransaction(ctx)
			if errAbort != nil {
				return
			}
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get reservation"}})
			return
		}

		// get the showing
		var showing models.Showing
		err = showingCollection.FindOne(ctx, bson.M{"_id": reservation.ShowingId}).Decode(&showing)
		if err != nil {
			errAbort := session.AbortTransaction(ctx)
			if errAbort != nil {
				return
			}
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get showing"}})
			return
		}

		// update the seats
		for _, reservedSeat := range oldReservation.ReservedSeats {
			showing.Seats[reservedSeat.RowNumber][reservedSeat.SeatNumber].IsReserved = false
		}

		for _, reservedSeat := range reservation.ReservedSeats {
			showing.Seats[reservedSeat.RowNumber][reservedSeat.SeatNumber].IsReserved = true
		}

		// update the available and booked seats
		showing.AvailableSeats += len(oldReservation.ReservedSeats)
		showing.BookedSeats -= len(oldReservation.ReservedSeats)
		showing.AvailableSeats -= len(reservation.ReservedSeats)
		showing.BookedSeats += len(reservation.ReservedSeats)

		// update the showing in the database
		_, err = showingCollection.UpdateOne(ctx, bson.M{"_id": reservation.ShowingId}, bson.M{"$set": showing})
		if err != nil {
			errAbort := session.AbortTransaction(ctx)
			if errAbort != nil {
				return
			}
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to update showing"}})
			return
		}

		result, err := reservationCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": reservation})
		if err != nil {
			errAbort := session.AbortTransaction(ctx)
			if errAbort != nil {
				return
			}
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to update reservation"}})
			return
		}

		// commit the transaction
		err = session.CommitTransaction(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to commit transaction"}})
			return
		}

		// get updated user details
		var updatedReservation models.Reservation
		if result.MatchedCount == 1 {
			err := reservationCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedReservation)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get updated reservation"}})
				return
			}
		}

		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedReservation, "description": "Reservation updated successfully"}})
	}
}

func DeleteReservation() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		reservationId := c.Param("reservationId")
		objId, _ := primitive.ObjectIDFromHex(reservationId)

		// start a session
		session, err := configs.DB.StartSession()
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to start session"}})
			return
		}
		defer session.EndSession(ctx)

		// start a transaction
		err = session.StartTransaction()
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to start transaction"}})
			return
		}

		// get the reservation
		var reservation models.Reservation
		err = reservationCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&reservation)
		if err != nil {
			errAbort := session.AbortTransaction(ctx)
			if errAbort != nil {
				return
			}
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get reservation"}})
			return
		}

		// get the showing
		var showing models.Showing
		err = showingCollection.FindOne(ctx, bson.M{"_id": reservation.ShowingId}).Decode(&showing)
		if err != nil {
			errAbort := session.AbortTransaction(ctx)
			if errAbort != nil {
				return
			}
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get showing"}})
			return
		}

		// update the seats
		for _, reservedSeat := range reservation.ReservedSeats {
			showing.Seats[reservedSeat.RowNumber][reservedSeat.SeatNumber].IsReserved = false
		}

		// update the available and booked seats
		showing.AvailableSeats += len(reservation.ReservedSeats)
		showing.BookedSeats -= len(reservation.ReservedSeats)

		// update the showing in the database
		_, err = showingCollection.UpdateOne(ctx, bson.M{"_id": reservation.ShowingId}, bson.M{"$set": showing})
		if err != nil {
			errAbort := session.AbortTransaction(ctx)
			if errAbort != nil {
				return
			}
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to update showing"}})
			return
		}

		result, err := reservationCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			errAbort := session.AbortTransaction(ctx)
			if errAbort != nil {
				return
			}
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to delete reservation"}})
			return
		}

		// commit the transaction
		err = session.CommitTransaction(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to commit transaction"}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				models.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": nil, "description": "Reservation with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": nil, "description": "Reservation successfully deleted!"}},
		)
	}
}

func GetAllReservations() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var reservations []models.Reservation

		results, err := reservationCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get reservations"}})
			return
		}

		// reading from the db in an optimal way
		defer func(results *mongo.Cursor, ctx context.Context) {
			err := results.Close(ctx)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get reservations"}})
			}
		}(results, ctx)

		for results.Next(ctx) {
			var singleReservation models.Reservation
			if err = results.Decode(&singleReservation); err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get reservations"}})
			}

			reservations = append(reservations, singleReservation)
		}

		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": reservations, "description": "Reservations found"}},
		)
	}
}

func GetReservationsByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		userId := c.Param("userId")
		objId, _ := primitive.ObjectIDFromHex(userId)
		var reservations []models.Reservation

		results, err := reservationCollection.Find(ctx, bson.M{"user_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get reservations"}})
			return
		}

		// reading from the db in an optimal way
		defer func(results *mongo.Cursor, ctx context.Context) {
			err := results.Close(ctx)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get reservations"}})
			}
		}(results, ctx)

		for results.Next(ctx) {
			var singleReservation models.Reservation
			if err = results.Decode(&singleReservation); err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get reservations"}})
			}

			reservations = append(reservations, singleReservation)
		}

		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": reservations, "description": "Reservations found"}},
		)
	}
}

func validateReservation(reservation models.Reservation, ctx context.Context) (error, string) {

	// use the validator library to validate required fields
	if validationErr := utils.Validator.Struct(&reservation); validationErr != nil {
		return validationErr, "error caught Validator library"
	}
	// check if showing_id exists
	err := showingCollection.FindOne(ctx, bson.M{"_id": reservation.Id}).Decode(&reservation)
	if err != nil {
		return err, "Showing with the provided ID does not exist"
	}

	// check if user_id exists
	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": reservation.UserId}).Decode(&user)
	if err != nil {
		return err, "User with the provided ID does not exist"
	}

	// check if movie exists and if its name is correct one with the provided ID
	movieCollection := configs.GetCollection("movies")
	var movie models.MovieRef
	err = movieCollection.FindOne(ctx, bson.M{"_id": reservation.MovieShowingRef.Id}).Decode(&movie)
	if err != nil {
		return err, "Movie with the provided ID does not exist"
	}
	if !reflect.DeepEqual(movie, reservation.MovieShowingRef) {
		return errors.New("not Deep Equal"), "Movie document is not identical with the provided one"
	}

	// check if total price is correct according to the reserved seats and the price of the showing
	var showing models.Showing
	err = showingCollection.FindOne(ctx, bson.M{"_id": reservation.ShowingId}).Decode(&showing)
	if err != nil {
		return err, "Showing with the provided ID does not exist"
	}
	if reservation.TotalPrice != (showing.PricePerTicket * float64(len(reservation.ReservedSeats))) {
		return errors.New("total price is incorrect"), "Total price is incorrect"
	}

	return nil, ""
}
