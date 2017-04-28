# 検索API

## 検索
検索APIは検索する文字を指定してそれに合致するお店を抽出できたり, ジャンルを使った絞り込みができます。 また, 指定した位置情報から半径nメートルのお店を抽出することもできます。

    GET /prepass/v0/search

#### URLクエリパラメータ

|パラメータ名|データ型|説明|
|:-----------|:-------|:---|
|q           |string  |検索のための検索クエリ|
|genre       |int     |検索したお店をジャンルでフィルターをかけます。複数指定する場合はコンマで繋いでください。例: |
|location    |string  |位置情報を指定してその周囲のお店を検索するためのクエリ。 緯度と経度をコンマで結合してください。|
|range       |int     |locationパラメータで指定した位置情報から検索に含める半径を指定します。 単位はメートルです。 `range`パラメータは`location`パラメータと一緒に指定してください, `location`パラメータまたは`range`パラメータのどちらか1つしか指定しなかった場合は`400`を返します。|

#### レスポンス

    Status: 200 OK

```cson
{
  "total_count": 5,
  "shops": [
    {
      "company_id": 123456,
      "shop_id": 1,
      "shop_name": "ほげ商店",
      "zipcode": "123-4567",
      "address": "石川県金沢市ほげ町ふが1-1",
      "building_address": "ほげビル1階",
      "latitude": 36.8895771,
      "longitude": 136.7679974,
      "tel": "123-456-7890",
      "fax": "123-456-7890",
      "url": "http://www.hogehoge.com/",
      "open_time": "平日10：00～19：00、土日祝10：00～18：00",
      "close_time": "毎週火曜日",
      "pr_message": "当店はほげほげでふがふがなおみせです",
      "image_urls": [
        "http://www.i-oyacomi.net/prepass/upimages/co123456ofid0001ofPic1_middle.jpg"
      ],
      "genres": [
        {
          "id": 2,
          "name": "食品"
        },
        {
          "id": 3,
          "name": "日用品"
        }
      ],
      "privilege": "",
      "privilege_content": "",
      "is_feed_space": false,
      "is_change_diaper_space": false,
      "is_microwave_oven": false,
      "can_buy_wet_tissues": false,
      "is_boil_water": false,
      "is_child_toilet": false,
      "is_kids_corner": false,
      "is_lent_stroller": false,
      "is_child_privilege": false,
      "is_child_menu": false,
      "is_no_smoking_room": false,
      "is_private_room": false,
      "is_zashiki": false,
      "last_update": "2017-4-12T13:23:33+09:00"
    },
    # ...
  ]
}
```
