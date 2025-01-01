package utils

import (
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/gertd/go-pluralize"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/nyaruka/phonenumbers"
	pb "github.com/shooters/user/internal/gen/protos/shooters/user/v1"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"math/rand"
	"strings"
	"time"
)

func NewUserFromRequest(userRequest *pb.UserRequest) (*gocloak.User, error) {
	newUser := new(gocloak.User)
	newUser.ID = gocloak.StringP(uuid.NewString())

	newUser.FirstName = gocloak.StringP(userRequest.GetFirstName())
	newUser.LastName = gocloak.StringP(userRequest.GetLastName())

	// Parse the phone number
	parsedNumber, err := phonenumbers.Parse(userRequest.GetPhoneNumber(), "")
	if err != nil {
		log.Println("Error parsing phone number:", err)
		return nil, err
	}

	log.Println("Phone number parsed ::::: |", parsedNumber)

	// Get the country code
	countryCode := phonenumbers.GetRegionCodeForNumber(parsedNumber)
	if countryCode == "" {
		log.Println("Country code not found")
		return nil, err
	}

	log.Println("Country code fetched ::::: |", countryCode)

	intlPhoneNumberFormat := phonenumbers.Format(parsedNumber, phonenumbers.E164)
	log.Println("Intl phone number ::::: |", intlPhoneNumberFormat)

	if userRequest.GetEmail() != "" {
		newUser.Email = gocloak.StringP(userRequest.GetEmail())
	}

	newUser.Username = gocloak.StringP(userRequest.GetUsername())
	newUser.Enabled = gocloak.BoolP(true)
	userGroup := userRequest.GetType().String()
	group := GetValueFromEnum(userGroup)
	newUser.Groups = &[]string{group}

	attributes := make(map[string][]string)
	attributes["phoneNumber"] = []string{intlPhoneNumberFormat}
	attributes["countryCode"] = []string{strings.ToUpper(countryCode)}

	var password string
	credentialRepresentations := make([]gocloak.CredentialRepresentation, 0)
	switch userRequest.GetType() {
	case pb.UserType_USER_TYPE_VENDOR, pb.UserType_USER_TYPE_ADMINISTRATOR:
		// Generate a random string password
		password = GenerateRandomString(10)
		credential := gocloak.CredentialRepresentation{
			Temporary: gocloak.BoolP(true),
			Type:      gocloak.StringP("password"),
			Value:     gocloak.StringP(password),
		}

		credentialRepresentations = append(credentialRepresentations, credential)
		newUser.RequiredActions = &[]string{"VERIFY_EMAIL", "UPDATE_PASSWORD"}
	case pb.UserType_USER_TYPE_PLAYER:
		password = fmt.Sprintf("%06d", rand.Intn(900000)+100000)
		log.Println("User PIN ::::: |", password)
		credential := gocloak.CredentialRepresentation{
			Temporary: gocloak.BoolP(false),
			Type:      gocloak.StringP("password"),
			Value:     gocloak.StringP(password),
		}
		credentialRepresentations = append(credentialRepresentations, credential)
	}

	newUser.Credentials = &credentialRepresentations
	newUser.Attributes = &attributes

	return newUser, nil
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func NewProtoFromKCUser(user *gocloak.User) (*pb.User, error) {
	userPb := new(pb.User)
	if err := copier.Copy(userPb, user); err != nil {
		return nil, err
	}
	//userPb.Type = pb.UserType(user.)

	return userPb, nil
}

func GetValueFromEnum(w string) string {
	p := pluralize.NewClient()
	split := strings.Split(w, "_")
	word := p.Plural(split[2])
	c := cases.Title(language.English)
	log.Println("Pluralized word: ", c.String(word))
	return c.String(word)
}
