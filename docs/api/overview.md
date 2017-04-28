# 概要

- [スキーマ](#scheme)
- [ルートエンドポイント](#root_endpoint)
- [バージョン](#versioning)
- [エラー](#error)
- [レートリミット](#rate_limiting)

## <a name="scheme"/> スキーマ
APIアクセスはすべて`HTTPS`で行います。

すべてのAPIで使われるタイムスタンプは [ISO 8601](https://ja.wikipedia.org/wiki/ISO_8601) の拡張方式で表されます :。

    YYYY-MM-DDTHH:MM:SSZ

文字エンコードは `UTF-8` です。

## <a name="root_endpoint"/> ルートエンドポイント
URLのルートエンドポイントは以下のようになります :。

    https://api.<ドメイン未定>

## <a name="versioning"/> バージョン
現在のAPIバージョンは `v0`(仕様策定中) です.  
予告なくエンドポイントやJSONの構造が変わる可能性があります。

APIバージョンはURLに含まれます :。

    https://api.<ドメイン未定>/prepass/v0/search

## <a name="error"/> エラー
エラーは`errors`という配列で返却されます。

それぞれの要素については以下のようになります。

|キー|データ型|説明|
|:---|:-------|:---|
|code|int|HTTPステータスコードではなくAPI側のエラーコードです. くわしくは[error](api/error)のページを見てね|
|message|string|エラーの大まかな内容です|
|description|string|エラーの詳しい説明|

### 必要なパラメータが足りない or パラメータのデータ型が違う

    Status: 400 Bad Request

```cson
{
  "errors":[
    {
      "code": 40001,
      "message": "Invalid params",
      "description": "Missing search query parameter"
    }
    # 複数ある場合もあるよ
  ]
}
```

## <a name="rate_limiting"/> レートリミット
現在の仕様ではレートリミットはありません。
