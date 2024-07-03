package stress

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"fmt"

	"cf-backend/db"
	"cf-backend/models"
	"cf-backend/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MakeFilesForStressTest(ticket *models.Ticket, code string) error {
	fmt.Println("hello from MakeFilesForStressTest")
	err := makeInputAndOutputFilesForStressTest(ticket)
	if err != nil {
		return err
	}
	err = writeCodeForStressTest(ticket, code)
	if err != nil {
		return err
	}

	return nil
}

func writeCodeForStressTest(ticket *models.Ticket, code string) error {
	fmt.Println("hello from writeCodeForStressTest")
	contestID := ticket.Problem.ContestID
	index := ticket.Problem.Index
	problemIndex := strings.ToLower(index)
	path := `/playground/contest/` + strconv.Itoa(contestID) + `/` + problemIndex
	fileName := "sub_id-" + strconv.Itoa(ticket.Submission.ID) + "-ticket-" + strconv.Itoa(ticket.TicketID) + "-code.cpp"
	err := utils.CreateAndWriteFile(path, fileName, code)
	if err != nil {
		errorMessage := "Oops! Something Went Wrong..."
		err = UpdateTicketError(ticket, errorMessage)
		return err
	}
	return nil
}

func makeInputAndOutputFilesForStressTest(ticket *models.Ticket) error {
	fmt.Println("hello from makeInputAndOutputFilesForStressTest")
	contestID := ticket.Problem.ContestID
	index := ticket.Problem.Index
	problemIndex := strings.ToLower(index)
	path := "/playground/contest/" + strconv.Itoa(contestID) + "/" + problemIndex
	fileName := "sub_id-" + strconv.Itoa(ticket.Submission.ID) + "-ticket-" + strconv.Itoa(ticket.TicketID) + "-input.txt"
	err := utils.MakeFile(path, fileName)
	if err != nil {
		errorMessage := "Oops! Something Went Wrong..."
		err = UpdateTicketError(ticket, errorMessage)
		return err
	}
	fileName = "sub_id-" + strconv.Itoa(ticket.Submission.ID) + "-ticket-" + strconv.Itoa(ticket.TicketID) + "-output-participant.txt"
	err = utils.MakeFile(path, fileName)
	if err != nil {
		errorMessage := "Oops! Something Went Wrong..."
		err = UpdateTicketError(ticket, errorMessage)
		return err
	}
	fileName = "sub_id-" + strconv.Itoa(ticket.Submission.ID) + "-ticket-" + strconv.Itoa(ticket.TicketID) + "-output-jury.txt"
	err = utils.MakeFile(path, fileName)
	if err != nil {
		errorMessage := "Oops! Something Went Wrong..."
		err = UpdateTicketError(ticket, errorMessage)
		return err
	}
	return nil
}

func UpdateTicketError(ticket *models.Ticket, message string) error {
	client, err := db.DbConnection()
	if err != nil {
		return err
	}
	ticketsCollection := client.Database("cfstress").Collection("tickets")
	filter := bson.D{primitive.E{Key: "ticket_id", Value: ticket.TicketID}}
	update := bson.M{"$set": bson.M{
		"progress": "processed",
		"verdict":  false,
		"error":    message,
	}}
	_, err = ticketsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return errors.New(message)
}