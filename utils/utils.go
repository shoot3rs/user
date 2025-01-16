package utils

import (
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/gertd/go-pluralize"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/nyaruka/phonenumbers"
	pb "github.com/shoot3rs/user/internal/gen/protos/shooters/user/v1"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"google.golang.org/genproto/googleapis/type/phone_number"
	"log"
	"math/rand"
	"strconv"
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
	group := GetValueFromEnum(userGroup, 2, "_")
	newUser.Groups = &[]string{group}

	attributes := make(map[string][]string)
	attributes["phoneNumber"] = []string{intlPhoneNumberFormat}
	attributes["phoneNumberVerified"] = []string{"false"}
	attributes["countryCode"] = []string{strings.ToUpper(countryCode)}

	var password string
	credentialRepresentations := make([]gocloak.CredentialRepresentation, 0)
	switch userRequest.GetType() {
	case pb.UserRole_USER_ROLE_VENDOR, pb.UserRole_USER_ROLE_ADMINISTRATOR:
		// Generate a random string password
		attributes["approved"] = []string{"false"}
		password = GenerateRandomString(10)
		credential := gocloak.CredentialRepresentation{
			Temporary: gocloak.BoolP(true),
			Type:      gocloak.StringP("password"),
			Value:     gocloak.StringP(password),
		}

		credentialRepresentations = append(credentialRepresentations, credential)
		newUser.RequiredActions = &[]string{"VERIFY_EMAIL", "UPDATE_PASSWORD"}
	case pb.UserRole_USER_ROLE_PLAYER:
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
	userPb.EmailVerified = gocloak.PBool(user.EmailVerified)
	log.Println("Roles :::: |", gocloak.PStringSlice(user.RealmRoles))
	userPb.Role = getRoleFromRealmRoles(user.RealmRoles)

	attributes := user.Attributes
	if attributes != nil {
		phoneNumberAttribute := getFieldFromUser("phoneNumber", attributes)
		phoneNumber, err := phonenumbers.Parse(phoneNumberAttribute, "")
		if err != nil {
			return nil, err
		}

		regionCode := phonenumbers.GetRegionCodeForNumber(phoneNumber)
		userPb.CountryCode = regionCode

		e164PhoneNumber := phonenumbers.Format(phoneNumber, phonenumbers.E164)
		phoneVerified := getFieldFromUser("phoneNumberVerified", attributes)

		userPb.PhoneNumber = &phone_number.PhoneNumber{
			Kind: &phone_number.PhoneNumber_E164Number{
				E164Number: e164PhoneNumber,
			},
			Extension: phoneNumber.GetExtension(),
		}

		approved := getFieldFromUser("approved", attributes)

		if isApproved, err := strconv.ParseBool(approved); err == nil {
			userPb.IsApproved = isApproved
		}

		if isPhoneVerified, err := strconv.ParseBool(phoneVerified); err == nil {
			userPb.PhoneNumberVerified = isPhoneVerified
		}
	}

	return userPb, nil
}

func getRoleFromRealmRoles(roles *[]string) pb.UserRole {
	for _, realmRole := range gocloak.PStringSlice(roles) {
		log.Println("Realm role :::: |", realmRole)
		switch realmRole {
		case "vendor":
			return pb.UserRole_USER_ROLE_VENDOR
		case "administrator":
			return pb.UserRole_USER_ROLE_ADMINISTRATOR
		case "staker":
			return pb.UserRole_USER_ROLE_PLAYER
		}
	}

	return pb.UserRole_USER_ROLE_UNSPECIFIED
}

func GetValueFromEnum(w string, splitter int, sep string) string {
	p := pluralize.NewClient()
	split := strings.Split(w, sep)
	word := p.Plural(split[splitter])
	c := cases.Title(language.English)
	return c.String(word)
}

func getFieldFromUser(field string, attr *map[string][]string) string {
	if val, ok := (*attr)[field]; ok {
		return val[0]
	}

	return ""
}
