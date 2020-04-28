package meander

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// APIKey : Google Place API の API key.
var APIKey string

// MapsGoogleApisURL : Google API PlacesのURL
const MapsGoogleApisURL = "https://maps.googleapis.com/maps/api/place/nearbysearch/json"

type (
	// Query : Google Places APIへの問い合わせ
	Query struct {
		Lat          float64
		Lng          float64
		Journey      []string
		Radius       int
		CostRangeStr string
	}

	// Place : Google Places APIから返されるレスポンス
	Place struct {
		*GoogleGeometry `json:"geometry"`
		Name            string         `json:"name"`
		Icon            string         `json:"icon"`
		Photos          []*googlePhoto `json:"photos"`
		Vicinity        string         `json:"vicinity"`
	}

	googleResponse struct {
		Results []*Place `json:"results"`
	}

	GoogleGeometry struct {
		*GoogleLocation `json:"location"`
	}

	GoogleLocation struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}

	googlePhoto struct {
		PhotoRef string `json:"photo_reference"`
		URL      string `json:"url"`
	}
)

func (q *Query) find(types string) (*googleResponse, error) {
	vals := make(url.Values)
	vals.Set("location", fmt.Sprintf("%f,%f", q.Lat, q.Lng))
	vals.Set("radius", fmt.Sprintf("%d", q.Radius))
	vals.Set("types", types)
	vals.Set("key", APIKey)

	// 価格帯の設定
	if len(q.CostRangeStr) > 0 {
		r := ParseCostRange(q.CostRangeStr)
		vals.Set("minprice", fmt.Sprintf("%d", int(r.From)-1))
		vals.Set("maxprice", fmt.Sprintf("%d", int(r.To)-1))
	}

	requestURL := MapsGoogleApisURL + "?" + vals.Encode()
	res, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response googleResponse

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Run : おすすめの生成
func (q *Query) Run() []interface{} {
	rand.Seed(time.Now().UnixNano())
	var w sync.WaitGroup
	var l sync.Mutex
	places := make([]interface{}, len(q.Journey))

	for i, r := range q.Journey {
		w.Add(1)
		go func(types string, i int) {
			defer w.Done()
			response, err := q.find(types)
			if err != nil {
				log.Println("施設の検索に失敗しました:", err)
				return
			}

			if len(response.Results) == 0 {
				log.Println("施設が見つかりませんでした:", types)
				return
			}

			for _, result := range response.Results {
				for _, photo := range result.Photos {
					photo.URL = fmt.Sprintf("https://maps.googleapis.com/maps/api/place/photo?maxwidth=1000&photoreference=%s&key=%s", photo.PhotoRef, APIKey)
				}
			}

			randI := rand.Intn(len(response.Results))
			l.Lock()
			places[i] = response.Results[randI]
			l.Unlock()
		}(r, i)
	}

	w.Wait() // 全てのリクエストの完了を待つ
	return places
}

// Public :
func (p *Place) Public() interface{} {
	return map[string]interface{}{
		"name":     p.Name,
		"icon":     p.Icon,
		"photos":   p.Photos,
		"vicinity": p.Vicinity,
		"lat":      p.Lat,
		"lng":      p.Lng,
	}
}
