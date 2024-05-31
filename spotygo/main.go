package main

// Create http request
import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sort"
)

type Artists struct {
	Artists struct {
		Items []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Followers struct {
				Href  interface{} `json:"href"`
				Total int         `json:"total"`
			} `json:"followers"`
			Genres []string `json:"genres"`
			Href   string   `json:"href"`
			ID     string   `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name       string `json:"name"`
			Popularity int    `json:"popularity"`
			Type       string `json:"type"`
			Uri        string `json:"uri"`
		} `json:"items"`
		Next    interface{} `json:"next"`
		Total   int         `json:"total"`
		Cursors struct {
			After interface{} `json:"after"`
		} `json:"cursors"`
		Limit int    `json:"limit"`
		Href  string `json:"href"`
	} `json:"artists"`
}

type ArtistQueryResponse struct {
	Tracks struct {
		Href     string `json:"href"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
		Items    []struct {
			Album struct {
				AlbumType    string `json:"album_type"`
				TotalTracks  int    `json:"total_tracks"`
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href   string `json:"href"`
				ID     string `json:"id"`
				Images []struct {
					Url    string `json:"url"`
					Height int    `json:"height"`
					Width  int    `json:"width"`
				} `json:"images"`
				Name                 string `json:"name"`
				ReleaseDate          string `json:"release_date"`
				ReleaseDatePrecision string `json:"release_date_precision"`
				Type                 string `json:"type"`
				Uri                  string `json:"uri"`
				Artists              []struct {
					ExternalUrls struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href string `json:"href"`
					ID   string `json:"id"`
					Name string `json:"name"`
					Type string `json:"type"`
					Uri  string `json:"uri"`
				} `json:"artists"`
				IsPlayable bool `json:"is_playable"`
			} `json:"album"`
			Artists []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				Uri  string `json:"uri"`
			} `json:"artists"`
			DiscNumber  int  `json:"disc_number"`
			DurationMs  int  `json:"duration_ms"`
			Explicit    bool `json:"explicit"`
			ExternalIds struct {
				Isrc string `json:"isrc"`
			} `json:"external_ids"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href        string `json:"href"`
			ID          string `json:"id"`
			IsPlayable  bool   `json:"is_playable"`
			Name        string `json:"name"`
			Popularity  int    `json:"popularity"`
			PreviewUrl  string `json:"preview_url"`
			TrackNumber int    `json:"track_number"`
			Type        string `json:"type"`
			Uri         string `json:"uri"`
			IsLocal     bool   `json:"is_local"`
		} `json:"items"`
	} `json:"tracks"`
	Albums struct {
		Href     string `json:"href"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
		Items    []struct {
			AlbumType    string `json:"album_type"`
			TotalTracks  int    `json:"total_tracks"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				Url    string `json:"url"`
				Height int    `json:"height"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name                 string `json:"name"`
			ReleaseDate          string `json:"release_date"`
			ReleaseDatePrecision string `json:"release_date_precision"`
			Type                 string `json:"type"`
			Uri                  string `json:"uri"`
			Artists              []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				Uri  string `json:"uri"`
			} `json:"artists"`
			IsPlayable bool `json:"is_playable"`
		} `json:"items"`
	} `json:"albums"`
}

type FeedItem struct {
	ArtistName   string
	TrackName    string
	ReleaseDate  string
	Link         string
	TrackOrAlbum string
}

func getFollowedArtists() Artists {
	artists := Artists{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/following?type=artist&limit=50", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return artists
	}
	req.Header.Set("Authorization", "Bearer BQDQFbp1CSX71YgjXiLL8XMFL-pX7-f7hG2fddHCNY_OOgJmB9FnsZwvT4ylNb3lOW43GSupolAOXL9EVOxSoQ9RvLluilWlvvgBbYIBXwgFgIWW8zgEdtM9xLvU4nTSpNtF3NVJzfwo5lLAV4j9hLWIGAOFPEPcRUWM2I-SaOyGKFIT8BgY")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return artists
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return artists
	}
	for _, artist := range artists.Artists.Items {
		fmt.Println("Name:", artist.Name)
	}
	return artists
}

func getTracksForArtist(artist_name string) []FeedItem {
	feedItems := []FeedItem{}
	artistQueryResponse := ArtistQueryResponse{}

	artistEncoded := url.QueryEscape(artist_name)
	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=artist:%s+year:2020-2024&type=track,album&market=NL&limit=2&offset=0", artistEncoded)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return feedItems
	}
	req.Header.Set("Authorization", "Bearer BQDQFbp1CSX71YgjXiLL8XMFL-pX7-f7hG2fddHCNY_OOgJmB9FnsZwvT4ylNb3lOW43GSupolAOXL9EVOxSoQ9RvLluilWlvvgBbYIBXwgFgIWW8zgEdtM9xLvU4nTSpNtF3NVJzfwo5lLAV4j9hLWIGAOFPEPcRUWM2I-SaOyGKFIT8BgY")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return feedItems
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
	err = json.NewDecoder(resp.Body).Decode(&artistQueryResponse)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return feedItems
	}
	for _, album := range artistQueryResponse.Tracks.Items {
		exact_artist_match_found := false
		artists := ""
		for _, artist := range album.Artists {
			if artist.Name == artist_name {
				exact_artist_match_found = true
			}
			artists += artist.Name + ", "
		}
		if !exact_artist_match_found {
			continue
		}
		fmt.Println("Name:", album.Name)
		artists = artists[:len(artists)-2]
		feedItems = append(feedItems, FeedItem{ArtistName: artists, TrackName: album.Name, ReleaseDate: album.Album.ReleaseDate, Link: album.ExternalUrls.Spotify, TrackOrAlbum: "track"})
	}

	for _, album := range artistQueryResponse.Albums.Items {
		exact_artist_match_found := false
		artists := ""
		for _, artist := range album.Artists {
			if artist.Name == artist_name {
				exact_artist_match_found = true
			}
			artists += artist.Name + ", "
		}
		if !exact_artist_match_found {
			continue
		}
		fmt.Println("Name:", album.Name)
		artists = artists[:len(artists)-2]
		feedItems = append(feedItems, FeedItem{ArtistName: artists, TrackName: album.Name, ReleaseDate: album.ReleaseDate, Link: album.ExternalUrls.Spotify, TrackOrAlbum: "album"})
	}
	return feedItems
}

func main() {
	artists := getFollowedArtists()
	feedItems := []FeedItem{}
	for _, artist := range artists.Artists.Items {
		fmt.Println("Getting tracks for artist:", artist.Name)
		newFeedItems := getTracksForArtist(artist.Name)
		feedItems = append(feedItems, newFeedItems...)
	}
	for _, feedItem := range feedItems {
		fmt.Println("Artist:", feedItem.ArtistName)
		fmt.Println("Track:", feedItem.TrackName)
		fmt.Println("Release Date:", feedItem.ReleaseDate)
		fmt.Println("Link:", feedItem.Link)
	}

	sort.Slice(feedItems, func(i, j int) bool {
		return feedItems[i].ReleaseDate > feedItems[j].ReleaseDate
	})

	// Write feedItems to a JSON file
	file, err := os.Create("../sveltekit-following-spotify/public/songs.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(feedItems)
	if err != nil {
		fmt.Println("Error encoding feedItems:", err)
		return
	}

	fmt.Println("FeedItems successfully written to feedItems.json")
}

// AQAwVnHFhNZncLal7oZrBjDaPJueDvoie4KH6aP-l1TJXtybM3fgD3yscpYKmo7cU-RsZROKiHvlUm7OBIaf9PltntnFlj-0DCIy4MXV6f3bCAFx4nXsAQlYBRsvkICUK9O8JQqmvvmlKeq6Q9O1mEq8F4hIXnyVHEyzFn_8VOf-5Q
