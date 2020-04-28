# Meander

## APIのテスト
座標を調べる
https://www.google.co.jp/maps

例: 新宿駅
```md
35.7071856,139.6983773
```

```bash
$ cd cmd

$ go build -o meanderapi

$ ./meanderapi

# 新宿駅周辺でカフェ・バー・レストランを検索
$ curl http://localhost:8080/recommendations\?lat\=35.689611\&lng\=139.6983773\&radius\=5000\&journey\=cafe\|bar\|restaurant\&cost\=$...$$$
[
  {
    "icon": "https://maps.gstatic.com/mapfiles/place_api/icons/restaurant-71.png",
    "lat": 35.6902956,
    "lng": 139.6943916,
    "name": "Jurin",
    "photos": [
      {
        "photo_reference": "CmRaAAAACQv4s4nf8TYBuaAXqaxQe7qr-1f16YJAvDqYQKHu2Ti_CHDN-nUJywfdcq8FH7uFk8raLc4zZFgmn1scsSIc1n8OQKb1cxX3OgMhzxjfTPQpeW3fDgtCtSfWGjoRFAaoEhCBcEp0HrSVDAJ91x_FaJT6GhTE5sQ5PIxoVVW2pwaKA_uje_D95Q",
        "url": "https://maps.googleapis.com/maps/api/place/photo?maxwidth=1000&photoreference=CmRaAAAACQv4s4nf8TYBuaAXqaxQe7qr-1f16YJAvDqYQKHu2Ti_CHDN-nUJywfdcq8FH7uFk8raLc4zZFgmn1scsSIc1n8OQKb1cxX3OgMhzxjfTPQpeW3fDgtCtSfWGjoRFAaoEhCBcEp0HrSVDAJ91x_FaJT6GhTE5sQ5PIxoVVW2pwaKA_uje_D95Q&key=AIzaSyBKvqfotQB4UgtHZClfXp1SEo5UuRb3zgI"
      }
    ],
    "vicinity": "京王プラザホテル内, 2 Chome-2−１ Nishishinjuku, Shinjuku City"
  },
  {
    "icon": "https://maps.gstatic.com/mapfiles/place_api/icons/bar-71.png",
    "lat": 35.66402679999999,
    "lng": 139.7307679,
    "name": "Shamrock",
    "photos": [
      {
        "photo_reference": "CmRaAAAAB5pfFtQAaqP2TnHqzPDZ7PWABhD6VV74mWjr6X-NYaIQlohWXLNGwzLHeqPtN1mhae3dCUykCEydEcT7muzH9RVbmyNSgrVmiCDSLZ514n_YbvwsFMI1ALtrSTD3EfBjEhAXwZ1HdDaX8MQokRwjYyFPGhSxftUo9ZB3It4HIw5GvoIkGqpu7g",
        "url": "https://maps.googleapis.com/maps/api/place/photo?maxwidth=1000&photoreference=CmRaAAAAB5pfFtQAaqP2TnHqzPDZ7PWABhD6VV74mWjr6X-NYaIQlohWXLNGwzLHeqPtN1mhae3dCUykCEydEcT7muzH9RVbmyNSgrVmiCDSLZ514n_YbvwsFMI1ALtrSTD3EfBjEhAXwZ1HdDaX8MQokRwjYyFPGhSxftUo9ZB3It4HIw5GvoIkGqpu7g&key=AIzaSyBKvqfotQB4UgtHZClfXp1SEo5UuRb3zgI"
      }
    ],
    "vicinity": "7 Chome-13-6 Roppongi, 港区 Minato City"
  },
  {
    "icon": "https://maps.gstatic.com/mapfiles/place_api/icons/restaurant-71.png",
    "lat": 35.67253,
    "lng": 139.7203,
    "name": "Kihachi",
    "photos": [
      {
        "photo_reference": "CmRaAAAA6UrI_ehCURZePOHQoHlG6OLzVUfLcSIsRvWFlBOdB4cdMTWidWBdpL2vY1k5hzFqoXDWAV8AGBRKz0K-PsyFpONnClSE4ZMtJGf3E-5kfch4jN-dQnUPPjCWUjOx_m6MEhBR4gTmt4O3BqMejGaCDhheGhRKBUew2sGX7EqHMw7FmiBUfGSPag",
        "url": "https://maps.googleapis.com/maps/api/place/photo?maxwidth=1000&photoreference=CmRaAAAA6UrI_ehCURZePOHQoHlG6OLzVUfLcSIsRvWFlBOdB4cdMTWidWBdpL2vY1k5hzFqoXDWAV8AGBRKz0K-PsyFpONnClSE4ZMtJGf3E-5kfch4jN-dQnUPPjCWUjOx_m6MEhBR4gTmt4O3BqMejGaCDhheGhRKBUew2sGX7EqHMw7FmiBUfGSPag&key=AIzaSyBKvqfotQB4UgtHZClfXp1SEo5UuRb3zgI"
      }
    ],
    "vicinity": "2 Chome-1-19 Kitaaoyama, Minato City"
  }
]
```

ブラウザから
```
http://localhost:8080/recommendations?lat=35.689611&lng=139.6983773&radius=5000&journey=cafe|bar|restaurant&cost=$...$$$
```

