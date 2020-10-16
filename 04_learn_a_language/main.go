package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
		"net/http"
		"strings"
		"github.com/PuerkitoBio/goquery"
		)

func main() {

		holoscopes := map[string]string{
			"おひつじ座" : "aries",
			"おうし座" : "taurus",
			"ふたご座" : "gemini",
			"かに座" : "cancer",
			"しし座" : "leo",
			"おとめ座" : "virgo",
			"てんびん座" : "libra",
			"さそり座" : "scorpio",
			"いて座" : "sagittarius",
			"やぎ座" : "capricorn",
			"みずがめ座" : "aquarius",
			"うお座" : "pisces",
		}

		url := "https://voguegirl.jp/horoscope/shiitake/pisces/20201012/"
		// TODO: 星座は入力値から　日付はその週の月曜日のものをつける
		holoscopeUrl := url + holoscopes["うお座"] + "/20201012"

    // Getリクエスト
		res, err := http.Get(holoscopeUrl)
		if err != nil {
			panic(err)
		}
    defer res.Body.Close()

    // 読み取り
    buf, _ := ioutil.ReadAll(res.Body)

    // 文字コード変換
		bReader := bytes.NewReader(buf)

    // HTMLパース
    doc, _ := goquery.NewDocumentFromReader(bReader)

		selection := doc.Find(".o-shiitake-detail__body")
			// 今週のあなたを分析
			rsltYou  := strings.TrimSpace(selection.Find(".o-shiitake-detail__section > .a-text").First().Text())
			// 今週どう乗り切る？
			howYouGetThrough := strings.TrimSpace(selection.Find(".o-shiitake-detail__section > .a-text").Last().Text())
			// TODO： このカラーをプラスして、今週の運勢をコントロール！も取得したい

			// TODO: 句読点の。のあとに改行を入れたい

			//
			fmt.Println("　　　　　🐟🐟🐟魚座🐟🐟🐟　　　　　")

			fmt.Println("꙳✧˖°⌖꙳✧˖°⌖꙳✧˖°⌖꙳✧˖°⌖꙳✧˖°⌖꙳✧˖°⌖꙳✧˖°⌖꙳✧˖°")
			fmt.Println("\n🍄しいたけ.がズバリ！\n今週のあなたを分析")
			fmt.Println("------------------------------------------------")
			fmt.Println(rsltYou)
			fmt.Println("\n🍄しいたけ.がアドバイス！\n今週どう乗り切る？")
			fmt.Println("------------------------------------------------")
			fmt.Println(howYouGetThrough)
			fmt.Println("꙳✧˖°⌖꙳✧˖°⌖꙳✧˖°⌖꙳✧˖°⌖꙳✧˖°⌖꙳✧˖°⌖꙳✧˖°⌖꙳✧˖°")

}