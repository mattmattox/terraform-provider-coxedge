/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package apiclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

//GetOriginSettingss Get originSettingss in account
func (c *Client) GetOriginSettingss(environmentName string) ([]OriginSettings, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/originSettingss",
		nil,
	)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedOriginSettingsSet
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

//GetOriginSettings Get originSettings in account by id
func (c *Client) GetOriginSettings(environmentName string, id string) (*OriginSettings, error) {
	//Create the request
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/originSettingss/"+id,
		nil,
	)
	if err != nil {
		return nil, err
	}

	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	//Unmarshal, unwrap, and return
	var wrappedAPIStruct WrappedOriginSettings
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

//CreateOriginSettings Create the originSettings
func (c *Client) CreateOriginSettings(newOriginSettings OriginSettings) (*OriginSettings, error) {
	return nil, errors.New("cannot create OriginSettings!")
}

//UpdateOriginSettings Update a originSettings
func (c *Client) UpdateOriginSettings(originSettingsId string, newOriginSettings OriginSettings) (*OriginSettings, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(newOriginSettings)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("PUT",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+newOriginSettings.EnvironmentName+"/originSettingss/"+originSettingsId,
		bReader,
	)
	request.Header.Set("Content-Type", "application/json")
	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	//Return struct
	var wrappedAPIStruct WrappedOriginSettings
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

//DeleteOriginSettings Delete originSettings in account by id
func (c *Client) DeleteOriginSettings(environmentName string, id string) error {
	//Create the request
	request, err := http.NewRequest("DELETE",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/originSettingss/"+id,
		nil,
	)
	if err != nil {
		return err
	}

	//Execute request
	_, err = c.doRequest(request)
	if err != nil {
		return err
	}
	return nil
}
