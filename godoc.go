/*
Package withings is a client module for working with the Withings (previously Nokia Health (previously Withings)) API. The current version (v2) of this module has been updated to work with the newer Oauth2 implementation of the API.

Authorization Overview

As with all Oauth2 APIs, you must obtain authorization to access the users data. To do so you must first register your application with the Withings api to obtain a ClientID and ClientSecret.
	http://developer.withings.com/oauth2/#tag/introduction

Once you have your client information you can now create a new client. You will need to also provide the redirectURL you provided during application registration. This URL is where the user will be redirected to and will include the code you need to generate the accessToken needed to access the users data. Under normal situations you should have an http server listening on that address and pull the code from the URL query parameters.
	client := withings.NewClient(clientID, clientSecret, clientRedirectURL)

You can now use the client to generate the authorization URL the user needs to navigate to. This URL will take the user to a page that will prompt them to allow your application to access their data. The state string returned by the method is a randomly generated BASE64 string using the crypto/rand module. It will be returned in the redirect as a query parameter and the two values verified they match. It's not required, but useful for security reasons.
	authURL, state, err := client.AuthCodeURL()

Using the code returned by the redirect in the query parameters, you can generate a new user. This user struct can be used to immediately perform actions against the users data. It also contains the token you should save somewhere for reuse. Obviously use whatever context you would like here.
	u, err := client.NewUserFromAuthCode(context.Background(), code)

Make sure the save at least the refreshToken for accessing the user data at a later date. You may also save the accessToken, but it does expire and creating a new client from saved token data only requires the refreshToken.
	refreshToken, err := i := u.Token.Token().RefreshToken

Creating User From Saved Token

You can easily create a user from a saved token using the NewUserFromRefreshToken method. A working configured client is required for the user generated from this method to work.

Requesting Data

The user struct has various methods associated with each API endpoint to perform data retrieval. The methods take a specific param struct specifying the api options to use on the request. The API is a bit "special" so the params vary a bit between each method. The client does what it can to smooth those out but there is only so much that can be done.
	startDate := time.Now().AddDate(0, 0, -2)
	endDate := time.Now()

	p := BodyMeasureQueryParams {
		StartDate: &startDate,
		EndDate: &endDate,
	}
	m, err := u.GetBodyMeasures(&p)


Context Usage

Every method has two forms, one that accepts a context and one that does now. This allows you to provide a context on each request if you would like to.
	startDate := time.Now().AddDate(0, 0, -2)
	endDate := time.Now()

	p := BodyMeasureQueryParams {
		StartDate: &startDate,
		EndDate: &endDate,
	}
	m, err := u.GetBodyMeasures(&p)

Request Timeout

By default all methods utilize a context to timeout the request to the API. The value of the timeout is stored on the Client and can be access as/set on Client.Timeout. Setting is _not_ thread safe and should only be set on client creation. If you need to change the
timeout for different requests use the methodCtx variant of the method.

Oauth2 State Randomization

By default the state generated by the AuthCodeURL utilized crypto/rand. If you would like to implement your own random method you can do so by assigning the function to Rand field of the Client struct. The function should support the Rand type. Also this is _not_ thread safe so only perform this action on client creation.

Raw Request Data

By default every returned response will be parsed and the parsed data returned. If you need access to the raw request data you can enable it by setting the SaveRawResponse field of the client struct to true. This should be done at client creation time. With it set to true the RawResponse field of the returned structs will include the raw response.

Data Helper Methods

Some data request methods include a parseResponse field on the params struct. If this is included additional parsing is performed to make the data more usable. This can be seen on GetBodyMeasures for example.

Include Path Fields In Response

You can include the path fields sent to the API by setting IncludePath to true on the client. This is primarily used for debugging but could be helpful in some situations.

Oauth2 Scopes

By default the client will request all known scopes. If you would like to pair down this you can change the scope by using the SetScope method of the client. Consts in the form of ScopeXxx are provided to aid selection.

*/
package withings
