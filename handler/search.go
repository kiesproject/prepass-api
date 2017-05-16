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
	es, ctx := elasticsearch.Elastic()

	query := c.QueryParam("q")
	//genre := c.QueryParam("genre")
	// location
	lat := c.QueryParam("lat")
	lon := c.QueryParam("lon")
	distance := c.QueryParam("range")
	sizeStr := c.QueryParam("size")

	// init
	searchService := es.Search("prepass")
	apiErrors := errors.APIErrors{}

	// TODO 2017/05/16 このクソみたいによくわからない条件分岐をどうにかする

	// クエリが指定されておらず、かつ緯度経度のどちらかが指定されていない
	if query == "" && (lat == "" || lon == "") {
		apiErrors.Errors = append(apiErrors.Errors, errors.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    "Not enough params. Please set search query or location query.",
		})
		return c.JSONPretty(http.StatusBadRequest, apiErrors, "  ")
	}

	// 検索クエリ指定
	if query != "" {
		term := elastic.NewSimpleQueryStringQuery(query).
			Field("shop_name^3").      // 店舗名 (重み3倍)
			Field("address").          // 住所
			Field("building_address"). // 建物住所
			DefaultOperator("and")
		searchService = searchService.Query(term)
	}

	// 位置情報検索
	if lat != "" && lon != "" && distance != "" {
		la, _ := strconv.ParseFloat(lat, 64)
		lo, _ := strconv.ParseFloat(lon, 64)

		distanceQuery := elastic.NewGeoDistanceQuery("location")
		distanceQuery.
			Lat(la).
			Lon(lo).
			Distance(distance).
			DistanceType("plane")
		searchService = searchService.Query(distanceQuery)
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		// デフォルトの検索結果表示数は20
		size = 20
	}

	searchService = searchService.Size(size)

	// execute
	result, err := searchService.Do(ctx)
	if err != nil {
		echo.Logger.Error(err.Error())
		apiErrors.Errors = append(apiErrors.Errors, errors.ApiError{
			StatusCode: http.StatusInternalServerError,
			Message:    "An error occurred during the search process",
		})
		return c.JSONPretty(http.StatusInternalServerError, err.Error(), "  ")
	}

	var shops = []*json.RawMessage{}

	for _, shop := range result.Hits.Hits {
		shops = append(shops, shop.Source)
	}

	return c.JSONPretty(http.StatusOK, SearchResult{TotalCount: result.Hits.TotalHits, Shops: shops}, "  ")
}
