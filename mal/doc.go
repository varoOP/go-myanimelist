/*
Package mal provides a client for accessing the MyAnimeList API:
https://myanimelist.net/apiconfig/references/api/v2.

Installation

This package can be installed using:

	go get github.com/nstratos/go-myanimelist/mal

Usage

Import the package using:

	import "github.com/nstratos/go-myanimelist/mal"

First construct a new mal client:

	c := mal.NewClient(nil)

Then use one of the client's services (User, Anime, Manga and Forum) to access
the different MyAnimeList API methods.

Authentication

When creating a new client, pass an http.Client that can handle authentication
for you. The recommended way is to use the golang.org/x/oauth2 package
(https://github.com/golang/oauth2). After performing the OAuth2 flow, you will
get an access token which can be used like this:

	ctx := context.Background()
	c := mal.NewClient(
		oauth2.NewClient(ctx, oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: "<your access token>"},
		)),
	)

Note that all calls made by the client above will include the specified access
token which is specific for an authenticated user. Therefore, authenticated
clients should almost never be shared between different users.

Performing the OAuth2 flow involves registering a MAL API application and then
asking for the user's consent to allow the application to access their data.

There is a detailed example of how to perform the Oauth2 flow and get an access
token through the terminal under example/malauth. The only thing you need to run
the example is a client ID and a client secret which you can acquire after
registering your MAL API application. Here's how:

 1. Navigate to https://myanimelist.net/apiconfig or go to your MyAnimeList
    profile, click Edit Profile and select the API tab on the far right.

 2. Click Create ID and submit the form with your application details.

After registering your application, you can run the example and pass the client
ID and client secret through flags:

	cd example/malauth
	go run main.go democlient.go --client-id=... --client-secret=...

	or

	go install github.com/nstratos/go-myanimelist/example/malauth
	malauth --client-id=... --client-secret=...

After you perform a successful authentication once, the access token will be
cached in a file under the same directory which makes it easier to run the
example multiple times.

Official MAL API OAuth2 docs:
https://myanimelist.net/apiconfig/references/authorization

List

To search and get anime and manga data:

	c := mal.NewClient()

	list, _, err := c.Anime.List(ctx, "hokuto no ken",
		mal.Fields{"rank", "popularity", "my_list_status"},
		mal.Limit(5),
	)
	// ...

	list, _, err := c.Manga.List(ctx, "hokuto no ken",
		mal.Fields{"rank", "popularity", "my_list_status"},
		mal.Limit(5),
	)
	// ...

You may get user specific data for a certain record by asking for the optional
field "my_list_status".

Official docs:

- https://myanimelist.net/apiconfig/references/api/v2#operation/anime_get
- https://myanimelist.net/apiconfig/references/api/v2#operation/manga_get

Search

To search for anime and manga:

	c := mal.NewClient(mal.Auth("<your username>", "<your password>"))

	result, _, err := c.Anime.Search("bebop")
	// ...

	result, _, err := c.Manga.Search("bebop")
	// ...

For more complex searches, you can provide the % operator which acts as a
wildcard and is escaped as %% in Go:

	result, _, err := c.Anime.Search("fate%%heaven%%flower")
	// ...
	// Will return: Fate/stay night Movie: Heaven's Feel - I. presage flower

Note: This is an undocumented feature of the MyAnimeList Search method.

Add

To add anime and manga, you provide their IDs and values through AnimeEntry and
MangaEntry:

	c := mal.NewClient(mal.Auth("<your username>", "<your password>"))

	_, err := c.Anime.Add(9989, mal.AnimeEntry{Status: mal.Current, Episode: 1})
	// ...

	_, err := c.Manga.Add(35733, mal.MangaEntry{Status: mal.Planned, Chapter: 1, Volume: 1})
	// ...

Note that when adding entries, Status is required.

Update

Similar to Add, Update also needs the ID of the entry and the values to be
updated:

	c := mal.NewClient(mal.Auth("<your username>", "<your password>"))

	_, err := c.Anime.Update(9989, mal.AnimeEntry{Status: mal.Completed, Score: 9})
	// ...

	_, err := c.Manga.Update(35733, mal.MangaEntry{Status: mal.OnHold})
	// ...

Delete

To delete anime and manga, simply provide their IDs:

	c := mal.NewClient(mal.Auth("<your username>", "<your password>"))

	_, err := c.Anime.Delete(9989)
	// ...

	_, err := c.Manga.Delete(35733)
	// ...

More Examples

See package examples:
https://godoc.org/github.com/nstratos/go-myanimelist/mal#pkg-examples

Advanced Control

If you need more control over the created requests, you can use an option to
pass a custom HTTP client to NewClient:

	c := mal.NewClient(&http.Client{})

For example this http.Client will make sure to cancel any request that takes
longer than 1 second:

	httpcl := &http.Client{
		Timeout: 1 * time.Second,
	}
	c := mal.NewClient(httpcl)
	// ...

Unit Testing

To run all unit tests:

	go test -cover

To see test coverage in your browser:

	go test -covermode=count -coverprofile=count.out && go tool cover -html count.out

Integration Testing

The integration tests will exercise the entire package against the live
MyAnimeList API. As a result, these tests take much longer to run and there is
also a much higher chance of false positives in test failures due to network
issues etc.

These tests are meant to be run using a dedicated test account that contains
empty anime and manga lists. A valid access token needs to be provided every
time.

To run the integration tests:

	go test --access-token '<your access token>'

License

MIT

*/
package mal
