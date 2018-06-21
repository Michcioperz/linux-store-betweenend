package main

import (
	"time"
)

type App struct {
	FlatpakAppId string `json:"flatpakAppId"`
	Name string `json:"name"`
	Summary string `json:"summary"`
	Description string `json:"description"`
	DeveloperName string `json:"developerName"`
	ProjectLicense string `json:"projectLicense"`
	HomepageUrl string `json:"homepageUrl"`
	BugtrackerUrl string `json:"bugtrackerUrl"`
	HelpUrl string `json:"helpUrl"`
	DonationUrl string `json:"donationUrl"`
	TranslateUrl string `json:"translateUrl"`
	IconDesktopUrl string `json:"iconDesktopUrl"`
	IconMobileUrl string `json:"iconMobileUrl"`
	DownloadFlatpakRefUrl string `json:"downloadFlatpakRefUrl"`
	CurrentReleaseVersion string `json:"currentReleaseVersion"`
	CurrentReleaseDate time.Time `json:"currentReleaseDate"`
	CurrentReleaseDescription string `json:"currentReleaseDescription"`
	InStoreSinceDate time.Time `json:"inStoreSinceDate"`
	Rating float64 `json:"rating"`
	RatingVotes int64 `json:"ratingVotes"`
	Categories []Category `json:"categories"`
	Screenshots []struct{
		ThumbUrl string `json:"thumbUrl"`
		ImgMobileUrl string `json:"imgMobileUrl"`
		ImgDesktopUrl string `json:"imgDesktopUrl"`
	} `json:"screenshots"`
}
