package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/spf13/viper"
)

var organizationUrl string
var organizationToken string
var home string
var configPath string
var validUrl = regexp.MustCompile(`^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/|\/|\/\/)?[A-z0-9_-]*?[:]?[A-z0-9_-]*?[@]?[A-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`)

func AuthConfig() error {

	fmt.Println("Enter your organization URL: ")
	fmt.Scanln(&organizationUrl)

	//url := &organizationUrl

	isValidUrl := validUrl.MatchString(organizationUrl)
	if !isValidUrl {
		return errors.New("input provisioned is not a valid URL")
	}

	fmt.Println("Enter your access token: ")
	fmt.Scanln(&organizationToken)

	viper.Set("OrganizationUrl", organizationUrl)
	encodedOrganizationToken := EncodeStringBase64(organizationToken)
	viper.Set("OrganizationToken", encodedOrganizationToken)

	home, _ = os.UserHomeDir()
	configPath = home + "/.adoctl"
	_, err := createFile(configPath, "adoctl.yaml")
	if err != nil {
		fmt.Println(err)
	}
	viper.SetConfigName("adoctl")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	err = viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func InlineAuthConfig(organizationUrl string, organizationToken string) error {

	isValidUrl := validUrl.MatchString(organizationUrl)
	if !isValidUrl {
		return errors.New("input provisioned is not a valid URL")
	}

	viper.Set("OrganizationUrl", organizationUrl)
	encodedOrganizationToken := EncodeStringBase64(organizationToken)
	viper.Set("OrganizationToken", encodedOrganizationToken)

	home, _ = os.UserHomeDir()
	configPath = home + "/.adoctl"
	_, err := createFile(configPath, "adoctl.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	viper.SetConfigName("adoctl")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	err = viper.WriteConfig()
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
