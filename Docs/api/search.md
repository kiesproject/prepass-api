# 検索API

## 検索
検索APIは検索する文字を指定してそれに合致するお店を抽出できたり, ジャンルを使った絞り込みができます。 また, 指定した位置情報から半径n (m or km)のお店を抽出することもできます。
`shops`の中身は最大1000件入っています。

    GET /prepass/v0/search

#### URLクエリパラメータ

|パラメータ名|データ型|説明|
|:-----------|:-------|:---|
|q           |string  |検索のためのクエリ|
|genre       |int     |検索したお店をジャンルでフィルターをかけます。複数指定する場合はコンマで繋いでください。例: |
|lat         |float   |位置情報を指定してその周囲のお店を検索を使用するときの緯度を指定するためのクエリ。|
|lon         |float   |位置情報を指定してその周囲のお店を検索を使用するときの経度を指定するためのクエリ。|
|range       |int     |`lat`,`lon`パラメータで指定した位置情報から検索に含める半径を指定します。 単位を省略するとメートルで判定します `range`パラメータは`lat`,`lon`パラメータと一緒に指定してください, `lat`,`lon`パラメータまたは`range`パラメータのどちらか片方しか指定しなかった場合は`400`を返します。|

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
      "zip_code": "123-4567",
      "address": "石川県金沢市ほげ町ふが1-1",
      "building_address": "ほげビル1階",
      "location": {
        "lat": 36.8895771, #緯度
        "lon": 136.7679974 #経度
      },
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
        "食品",
        "日用品"
      ],
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
      "antiallergic_support": "", #アレルギー対応
      "privileges": {
        "two_children": "ほげほげをほげ%OFF!!!!",   # 子供が2人いる家庭の特典
        "three_children": "ほげほげをふが%OFF!!!!", # 子供が3人いる家庭の特典
      },
      "last_update": "2017-05-19T00:30:41+0900" # データ更新日時
    },
    # ...
  ]
}
```
