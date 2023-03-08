// package configs contains configuration settings that can
// be used by other packages.
package configs

const (

	// CLIENT_APOD_URI is the URL of the Nasa Picture of the Day API
	CLIENT_APOD_URI string = "https://api.nasa.gov/planetary/apod"
)

var (

	// CLIENT_APOD_PARAMS is a map which contains the parameters that
	// will be used in the APOD API's request
	CLIENT_APOD_PARAMS = map[string]string{"api_key": "DEMO_KEY"}
)
