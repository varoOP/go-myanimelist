package mal

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAnimeService_Delete(t *testing.T) {
	setup()
	defer teardown()

	client.SetCredentials("TestUser", "TestPass")
	client.SetUserAgent("TestAgent")

	mux.HandleFunc("/api/animelist/delete/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testID(t, r, "55")
		testBasicAuth(t, r, "TestUser", "TestPass")
		testUserAgent(t, r, "TestAgent")
		fmt.Fprintf(w, "Deleted")
	})

	_, err := client.Anime.Delete(55)
	if err != nil {
		t.Errorf("Anime.Delete returned error %v", err)
	}
}

func TestAnimeService_Add(t *testing.T) {
	setup()
	defer teardown()

	client.SetCredentials("TestUser", "TestPass")
	client.SetUserAgent("TestAgent")

	mux.HandleFunc("/api/animelist/add/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testID(t, r, "55")
		testBasicAuth(t, r, "TestUser", "TestPass")
		testUserAgent(t, r, "TestAgent")
		testContentType(t, r, "application/x-www-form-urlencoded")
		testFormValue(t, r, "data", "<entry><status>watching</status></entry>")
		fmt.Fprintf(w, "Created")
	})

	_, err := client.Anime.Add(55, AnimeEntry{Status: "watching"})
	if err != nil {
		t.Errorf("Anime.Add returned error %v", err)
	}
}

func TestAnimeService_Update(t *testing.T) {
	setup()
	defer teardown()

	client.SetCredentials("TestUser", "TestPass")
	client.SetUserAgent("TestAgent")

	mux.HandleFunc("/api/animelist/update/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testID(t, r, "55")
		testBasicAuth(t, r, "TestUser", "TestPass")
		testUserAgent(t, r, "TestAgent")
		testContentType(t, r, "application/x-www-form-urlencoded")
		testFormValue(t, r, "data", "<entry><status>onhold</status></entry>")
		fmt.Fprintf(w, "Updated")
	})

	_, err := client.Anime.Update(55, AnimeEntry{Status: "onhold"})
	if err != nil {
		t.Errorf("Anime.Update returned error %v", err)
	}
}

func TestAnimeService_Search(t *testing.T) {
	setup()
	defer teardown()

	client.SetCredentials("TestUser", "TestPass")
	client.SetUserAgent("TestAgent")

	mux.HandleFunc("/api/anime/search.xml", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testBasicAuth(t, r, "TestUser", "TestPass")
		testUserAgent(t, r, "TestAgent")
		testURLValues(t, r, urlValues{"q": "query"})
		fmt.Fprintf(w, `
			<anime>
				<entry>
					<title>title1</title>
					<id>55</id>
				</entry>
				<entry>
					<title>title2</title>
					<id>56</id>
				</entry>
			</anime>`)
	})

	result, _, err := client.Anime.Search("query")
	if err != nil {
		t.Errorf("Anime.Search returned error %v", err)
	}
	want := &AnimeResult{
		[]AnimeRow{
			AnimeRow{Row: Row{ID: 55, Title: "title1"}},
			AnimeRow{Row: Row{ID: 56, Title: "title2"}},
		},
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Anime.Search returned %+v, want %+v", result, want)
	}
}
