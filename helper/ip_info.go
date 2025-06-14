package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/barkin-kaplan/weather-api/helper/model"
)


func GetIPInfo(ip string) (*model.IPInfo, error) {
	resp, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s", ip))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var info model.IPInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}

	return &info, nil
}