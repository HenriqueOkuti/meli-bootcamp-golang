package exercise02

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

/*
Requirements:

- TL;DR -> RTF
*/
func SolveExercise02() {
	const (
		validFileName1   = "file1.txt"
		validFileName2   = "file2.txt"
		invalidFileName  = "invalid.txt"
		databaseFileName = "customers.txt"
	)

	fileNames := []string{validFileName1, validFileName2, invalidFileName, databaseFileName}

	solveEVERYTHING(fileNames)

}

func solveEVERYTHING(fileNames []string) {
	for _, fileName := range fileNames {
		var fileData = readFile(fileName)
		var clientData = convertFileDataToClientData(fileData)
		if clientData.Name != "" {
			fmt.Println("Reading user data: ", clientData)
		}
	}

}

func readFile(filename string) *os.File {
	defer func() {
		if r := recover(); r != nil {
			errMsg := errors.New("error: unable to read file")
			fmt.Println(errMsg)
		}
	}()

	data, err := os.Open(filename)
	if err != nil {
		panic("error: unable to read file")
	}

	return data
}

type Client struct {
	ID        uint8
	Name      string
	CitizenID string
	Phone     string
	Address   string
}

func convertFileDataToClientData(file *os.File) Client {
	defer func() {
		if r := recover(); r != nil {
			errMsg := errors.New("error: unable to convert file data to client data")
			fmt.Println(errMsg)
		}
	}()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fileBytesString := string(fileBytes)
	fileData := strings.Split(fileBytesString, ";")

	var newClientId uint8 = uint8(rand.Uint64()) // Requirement from Task 1

	fmt.Println("Trying to read database file: customers.txt")
	var databaseFile = fetchUserDatabaseFileData() // Requirement from Task 2
	if databaseFile != nil {
		fmt.Println("Database file read successfully")
	}

	return Client{
		ID:        newClientId,
		Name:      fileData[0],
		CitizenID: fileData[1],
		Phone:     fileData[2],
		Address:   fileData[3],
	}

}

func fetchUserDatabaseFileData() *os.File {
	defer func() {
		if r := recover(); r != nil {
			errMsg := errors.New("error: unable to fetch user database file data")
			fmt.Println(errMsg)
		}
	}()

	const databaseFileName = "customers.txt"

	var fileData, err = os.Open(databaseFileName)
	if err != nil {
		panic("error: unable to read file")
	}

	return fileData

}
