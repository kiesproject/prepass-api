package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kiesproject/prepass-api/elasticsearch"
	"github.com/kiesproject/prepass-api/errors"
	"github.com/labstack/echo"
	"gopkg.in/olivere/elastic.v5"
)

type SearchResult struct {
	TotalCount int64              `json:"total_count"`
	Shops      []*json.RawMessage `json:"shops"`
}

func GetSearch(c echo.Context) error {
	apiErrors := errors.NewApiErrors()
	version := c.Param("version")

	if version != "v0" {
		apiErrors = apiErrors.AddError(http.StatusNotFound, "This API endpoint does not exist.")
		return c.JSONPretty(http.StatusNotFound, apiErrors, "  ")
	}

	es, ctx := elasticsearch.Elastic()

	query := c.QueryParam("q")
	//genre := c.QueryParam("genre")
	// location
	lat := c.QueryParam("lat")
	lon := c.QueryParam("lon")
	distance := c.QueryParam("range")

	// init
	searchService := es.Search("prepass")

	// TODO 2017/05/16 このクソみたいによくわからない条件分岐をどうにかする

	// クエリが指定されておらず、かつ緯度経度のどちらかが指定されていない
	if query == "" && (lat == "" || lon == "") {
		apiErrors = apiErrors.AddError(
			http.StatusBadRequest,
			"Not enough params. Please set `q` or `lat`, `lon`, `range` query.",
		)
		return c.JSONPretty(http.StatusBadRequest, apiErrors, "  ")
	}

	// 検索クエリ指定
	if query != "" {
		term := elastic.NewSimpleQueryStringQuery(query).
			Field("shop_name.ja^3"). // 店舗名 (重み3倍)
			Field("shop_name.fact^3").
			Field("shop_name.ngram^3").
			Field("genres.genre.ja^3"). // ジャンル
			Field("pr_message.ja^2"). // PRメッセージ (重み2倍)
			Field("pr_message.fact^2").
			Field("pr_message.ngram^2").
			Field("address.ja"). // 住所
			Field("address.fact").
			Field("address.ngram").
			Field("building_address.ja"). // 建物住所
			Field("building_address.fact").
			Field("building_address.ngram").
			DefaultOperator("and").
			Flags("OR|AND|NOT") // 検索時の特殊検索を使えるようにする
		searchService = searchService.Query(term)
	}

	// 位置情報検索
	if lat != "" && lon != "" && distance != "" {
		la, errLat := strconv.ParseFloat(lat, 64)
		if errLat != nil {
			apiErrors = apiErrors.AddError(
				http.StatusBadRequest,
				"Invalid latitude format: must be a positive float number",
			)
		}
		lo, errLon := strconv.ParseFloat(lon, 64)
		if errLon != nil {
			apiErrors = apiErrors.AddError(
				http.StatusBadRequest,
				"Invalid longitude format: must be a positive float number",
			)
		}

		if errLat != nil || errLon != nil {
			return c.JSONPretty(http.StatusBadRequest, apiErrors, "  ")
		}

		distanceQuery := elastic.NewGeoDistanceQuery("location")
		distanceQuery.
			Lat(la).
			Lon(lo).
			Distance(distance).
			DistanceType("plane")
		searchService = searchService.Query(distanceQuery)

		distanceSort := elastic.NewGeoDistanceSort("location")
		distanceSort.
			Point(la, lo).
			Asc().
			Unit("meters")

		searchService = searchService.SortBy(distanceSort)
	}

	// 頭悪いけど1000件以上合致する検索しなさそうだし1000件で決め打ち
	// TODO: ページング処理の追加
	searchService = searchService.Size(1000)

	// execute
	result, err := searchService.Do(ctx)
	if err != nil {
		apiErrors = apiErrors.AddError(
			http.StatusInternalServerError,
			"An error occurred during the search process",
		)
		return c.JSONPretty(http.StatusInternalServerError, err.Error(), "  ")
	}

	var shops = []*json.RawMessage{}

	for _, shop := range result.Hits.Hits {
		shops = append(shops, shop.Source)
	}

	return c.JSONPretty(http.StatusOK, SearchResult{TotalCount: result.Hits.TotalHits, Shops: shops}, "  ")
}
