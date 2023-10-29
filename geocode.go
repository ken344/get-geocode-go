package address2geocode

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
	"log"
	"os"
)

// SetDotenv .envファイルを読み込む
func SetDotenv(envPath string) {
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func init() {
	log.SetFlags(log.Lshortfile)
	// envファイルを読み込む
	SetDotenv(".env")
}

// Address2Geocode address2Geocode 住所から緯度経度を取得する
// google-mapsのGeocoding APIを使用
func Address2Geocode(address string) (*maps.LatLng, error) {
	// Geocoding APIのクライアントを作成
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_API_KEY_MAP")))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	// Geocoding APIのリクエストを作成
	// Addressに住所を指定する
	r := &maps.GeocodingRequest{
		Address: address,
	}

	// Geocoding APIのリクエストを実行
	resp, err := c.Geocode(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	fmt.Println(pretty.Sprint(resp))
	if len(resp) == 0 {
		log.Fatalf("fatal error: %s", "No results")
		return &resp[0].Geometry.Location, nil
	}
	// Geocoding APIのレスポンスから緯度経度を取得
	return &resp[0].Geometry.Location, nil
}

func example() {
	// 住所を関数に渡して、緯度経度を戻り値として取得。
	location, err := Address2Geocode("東京都新宿区西新宿２丁目８−１")
	if err != nil {
		log.Fatal(err)
	}
	//緯度経度を出力
	fmt.Println(location)
}
