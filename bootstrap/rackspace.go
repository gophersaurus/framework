package bootstrap

import (
	"log"

	"git.target.com/gophersaurus/gophersaurus/config"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/rackspace"
)

// Rackspace bootstraps a Rackspace connection.
func RackspaceObjectStorage(c config.Config) *gophercloud.ServiceClient {

	// Authenticate to Rackspace.
	provider, err := rackspace.AuthenticatedClient(gophercloud.AuthOptions{
		Username: c.Services.Rackspace.User,
		APIKey:   c.Services.Rackspace.Key,
	})

	// Check error.
	if err != nil {
		log.Fatal("Unable to connect to Rackspace: " + err.Error())
	}

	// Create an Open Stack object storage client.
	client, err := rackspace.NewObjectStorageV1(
		provider,
		gophercloud.EndpointOpts{
			Region: c.Services.Rackspace.Region,
		},
	)

	// Check error.
	if err != nil {
		log.Fatal("Unable to create a client: " + err.Error())
	}

	return client
}
