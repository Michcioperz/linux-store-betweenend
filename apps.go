package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

var FlathubBaseUrl, FlathubAppUrl, FlathubCategoryUrl *url.URL

func init() {
	FlathubBaseUrl, err := url.Parse("https://flathub.org/api/v1/")
	if err != nil {
		panic(err)
	}
	FlathubAppUrl, err = url.Parse("apps/")
	if err != nil {
		panic(err)
	}
	FlathubAppUrl = FlathubBaseUrl.ResolveReference(FlathubAppUrl)
	FlathubCategoryUrl, err = url.Parse("ctegory/")
	if err != nil {
		panic(err)
	}
	FlathubCategoryUrl = FlathubAppUrl.ResolveReference(FlathubCategoryUrl)
}

func GetApp(appId string) (App, error) {
	u, err := url.Parse(appId)
	if err != nil {
		return App{}, err
	}
	u = FlathubAppUrl.ResolveReference(u)
	resp, err := http.Get(u.String())
	if err != nil {
		return App{}, err
	}
	defer resp.Body.Close()
	var app App
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&app)
	return app, err
}

func GetCategory(categoryId string) ([]App, error) {
	u, err := url.Parse(categoryId)
	if err != nil {
		return []App{}, err
	}
	u = FlathubCategoryUrl.ResolveReference(u)
	resp, err := http.Get(u.String())
	if err != nil {
		return []App{}, err
	}
	defer resp.Body.Close()
	var apps []App
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&apps)
	return apps, err
}

func GetAllApps() ([]App, error) {
	resp, err := http.Get(FlathubAppUrl.String())
	if err != nil {
		return []App{}, err
	}
	defer resp.Body.Close()
	var apps []App
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&apps)
	return apps, err
}
