package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const azureInstanceMetadataEndpoint = "http://169.254.169.254/metadata/instance"

// ComputeInstanceMetadata denotes the compute specific instance metadata from Azure Instance
// Metadata Service
type ComputeInstanceMetadata struct {
	Location             string `json:"location,omitempty"`
	Name                 string `json:"name,omitempty"`
	Offer                string `json:"offer,omitempty"`
	OsType               string `json:"osType,omitempty"`
	PlacementGroupID     string `json:"placement_group_id,omitempty"`
	PlatformFaultDomain  string `json:"platformFaultDomain,omitempty"`
	PlatformUpdateDomain string `json:"platformUpdateDomain,omitempty"`
	Publisher            string `json:"publisher,omitempty"`
	ResourceGroupName    string `json:"resourceGroupName,omitempty"`
	Sku                  string `json:"sku,omitempty"`
	SubscriptionID       string `json:"subscriptionId,omitempty"`
	Tags                 string `json:"tags,omitempty"`
	Version              string `json:"version,omitempty"`
	VMID                 string `json:"vmId,omitempty"`
	VMScaleSetName       string `json:"vmScaleSetName,omitempty"`
	VMSize               string `json:"vmSize,omitempty"`
	Zone                 string `json:"zone,omitempty"`
}

// Queries the Azure Instance Metadata Service for the instance's metadata
func retrieveComputeInstanceMetadata() (metadata ComputeInstanceMetadata, err error) {
	var m ComputeInstanceMetadata
	c := &http.Client{}
	req, _ := http.NewRequest("GET", azureInstanceMetadataEndpoint+"/compute", nil)
	req.Header.Add("Metadata", "True")
	q := req.URL.Query()
	q.Add("format", "json")
	q.Add("api-version", "2017-08-01")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Do(req)
	if err != nil {
		return m, fmt.Errorf("sending Azure Instance Metadata Service request failed: %v", err)
	}
	defer resp.Body.Close()
	rawJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return m, fmt.Errorf("reading response body failed: %v", err)
	}
	if err := json.Unmarshal(rawJSON, &m); err != nil {
		return m, fmt.Errorf("unmarshaling JSON response failed: %v", err)
	}
	return m, nil
}

func main() {
	m, err := retrieveComputeInstanceMetadata()
	if err != nil {
		panic(fmt.Errorf("failed to retrieve instance metadata: %v", err))
	}
	fmt.Printf("Instance Name: %s", m.Name)
	fmt.Printf("VM Scale Set Name: %s", m.VMScaleSetName)
	fmt.Printf("Subscription ID: %s", m.SubscriptionID)
}
