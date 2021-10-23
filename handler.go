package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/macaron.v1"
)

func handlerMain(ctx *macaron.Context) {
	ctx.JSON(http.StatusOK, "{'status': 1}")
}

func handleQueryNFTItem(ctx *macaron.Context) {
	coordX := ctx.QueryFloat64("x")
	coordY := ctx.QueryFloat64("y")

	if coordX == 0 {
		ctx.JSON(http.StatusBadRequest, "invalid X value")
		return
	}
	if coordY == 0 {
		ctx.JSON(http.StatusBadRequest, "invalid Y value")
		return
	}

	resp := NFTItemResponse{}
	resp.ID = "12332232232"
	resp.CoordX = fmt.Sprintf("%f", coordX)
	resp.CoordY = fmt.Sprintf("%f", coordY)

	bPlan := true
	var body []byte
	openseaResp := OpenseaData{}
	openseaAlternativeContent := `{"id":158831,"token_id":"1","num_sales":3,"background_color":null,"image_url":"https://lh3.googleusercontent.com/7bRocEaoBrWYBX3vThkHj4kAV3b3mKG-Kem85xeT-D8oHpvQ19kcoiBd9mIFeNU0GrwZGvj6Oc5NAEGBSsGlrww","image_preview_url":"https://lh3.googleusercontent.com/7bRocEaoBrWYBX3vThkHj4kAV3b3mKG-Kem85xeT-D8oHpvQ19kcoiBd9mIFeNU0GrwZGvj6Oc5NAEGBSsGlrww=s250","image_thumbnail_url":"https://lh3.googleusercontent.com/7bRocEaoBrWYBX3vThkHj4kAV3b3mKG-Kem85xeT-D8oHpvQ19kcoiBd9mIFeNU0GrwZGvj6Oc5NAEGBSsGlrww=s128","image_original_url":"https://www.larvalabs.com/cryptopunks/cryptopunk1.png","animation_url":null,"animation_original_url":null,"name":"CryptoPunk #1","description":null,"external_link":"https://www.larvalabs.com/cryptopunks/details/1","asset_contract":{"address":"0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb","asset_contract_type":"non-fungible","created_date":"2018-01-23T04:51:38.832339","name":"CryptoPunks","nft_version":"unsupported","opensea_version":null,"owner":null,"schema_name":"CRYPTOPUNKS","symbol":"PUNK","total_supply":null,"description":"CryptoPunks launched as a fixed set of 10,000 items in mid-2017 and became one of the inspirations for the ERC-721 standard. They have been featured in places like The New York Times, Christie’s of London, Art|Basel Miami, and The PBS NewsHour.","external_link":"https://www.larvalabs.com/cryptopunks","image_url":"https://lh3.googleusercontent.com/BdxvLseXcfl57BiuQcQYdJ64v-aI8din7WPk0Pgo3qQFhAUH-B6i-dCqqc_mCkRIzULmwzwecnohLhrcH8A9mpWIZqA7ygc52Sr81hE=s120","default_to_fiat":false,"dev_buyer_fee_basis_points":0,"dev_seller_fee_basis_points":0,"only_proxied_transfers":false,"opensea_buyer_fee_basis_points":0,"opensea_seller_fee_basis_points":250,"buyer_fee_basis_points":0,"seller_fee_basis_points":250,"payout_address":null},"permalink":"https://opensea.io/assets/0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb/1","collection":{"payment_tokens":[{"id":13689077,"symbol":"ETH","address":"0x0000000000000000000000000000000000000000","image_url":"https://storage.opensea.io/files/6f8e2979d428180222796ff4a33ab929.svg","name":"Ether","decimals":18,"eth_price":0,"usd_price":4023.84},{"id":12182941,"symbol":"DAI","address":"0x6b175474e89094c44da98b954eedeac495271d0f","image_url":"https://storage.opensea.io/files/8ef8fb3fe707f693e57cdbfea130c24c.svg","name":"Dai Stablecoin","decimals":18,"eth_price":0,"usd_price":1.01},{"id":4645681,"symbol":"WETH","address":"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2","image_url":"https://storage.opensea.io/files/accae6b6fb3888cbff27a013729c22dc.svg","name":"Wrapped Ether","decimals":18,"eth_price":0,"usd_price":4023.84},{"id":4403908,"symbol":"USDC","address":"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48","image_url":"https://storage.opensea.io/files/749015f009a66abcb3bbb3502ae2f1ce.svg","name":"USD Coin","decimals":6,"eth_price":0,"usd_price":1.01}],"primary_asset_contracts":[{"address":"0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb","asset_contract_type":"non-fungible","created_date":"2018-01-23T04:51:38.832339","name":"CryptoPunks","nft_version":"unsupported","opensea_version":null,"owner":null,"schema_name":"CRYPTOPUNKS","symbol":"PUNK","total_supply":null,"description":"CryptoPunks launched as a fixed set of 10,000 items in mid-2017 and became one of the inspirations for the ERC-721 standard. They have been featured in places like The New York Times, Christie’s of London, Art|Basel Miami, and The PBS NewsHour.","external_link":"https://www.larvalabs.com/cryptopunks","image_url":"https://lh3.googleusercontent.com/BdxvLseXcfl57BiuQcQYdJ64v-aI8din7WPk0Pgo3qQFhAUH-B6i-dCqqc_mCkRIzULmwzwecnohLhrcH8A9mpWIZqA7ygc52Sr81hE=s120","default_to_fiat":false,"dev_buyer_fee_basis_points":0,"dev_seller_fee_basis_points":0,"only_proxied_transfers":false,"opensea_buyer_fee_basis_points":0,"opensea_seller_fee_basis_points":250,"buyer_fee_basis_points":0,"seller_fee_basis_points":250,"payout_address":null}],"traits":{"trait_type":"","value":"","display_type":null,"max_value":null,"trait_count":0,"order":null},"stats":{"one_day_volume":1343.48,"one_day_change":-0.0315027609971309,"one_day_sales":0,"one_day_average_price":149.27555555555557,"seven_day_volume":8727.1669,"seven_day_change":0.505047235539613,"seven_day_sales":0,"seven_day_average_price":122.91784366197183,"thirty_day_volume":55904.3569,"thirty_day_change":-0.6010463913441105,"thirty_day_sales":0,"thirty_day_average_price":131.8498983490566,"total_volume":560936.648501429,"total_sales":0,"total_supply":0,"count":0,"num_owners":3168,"average_price":32.07551741202121,"num_reports":1,"market_cap":1229055.5187760564,"floor_price":0},"banner_image_url":"https://lh3.googleusercontent.com/48oVuDyfe_xhs24BC2TTVcaYCX7rrU5mpuQLyTgRDbKHj2PtzKZsQ5qC3xTH4ar34wwAXxEKH8uUDPAGffbg7boeGYqX6op5vBDcbA=s2500","chat_url":null,"created_date":"2019-04-26T22:13:09.691572","default_to_fiat":false,"description":"CryptoPunks launched as a fixed set of 10,000 items in mid-2017 and became one of the inspirations for the ERC-721 standard. They have been featured in places like The New York Times, Christie’s of London, Art|Basel Miami, and The PBS NewsHour.","dev_buyer_fee_basis_points":"0","dev_seller_fee_basis_points":"0","discord_url":"https://discord.gg/tQp4pSE","display_data":{"card_display_style":"cover"},"external_url":"https://www.larvalabs.com/cryptopunks","featured":false,"featured_image_url":"https://storage.opensea.io/0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb-featured-1556589448.png","hidden":false,"safelist_request_status":"verified","image_url":"https://lh3.googleusercontent.com/BdxvLseXcfl57BiuQcQYdJ64v-aI8din7WPk0Pgo3qQFhAUH-B6i-dCqqc_mCkRIzULmwzwecnohLhrcH8A9mpWIZqA7ygc52Sr81hE=s120","is_subject_to_whitelist":false,"large_image_url":"https://lh3.googleusercontent.com/QB2kKuQEw04X02V9EoC2BNYZV652LYuewUv9ZdR7KJfI9Jocwmd28jIfsGg0umSCr2bOMV8O9UpLAkoaqfYwvwmC","medium_username":null,"name":"CryptoPunks","only_proxied_transfers":false,"opensea_buyer_fee_basis_points":"0","opensea_seller_fee_basis_points":"250","payout_address":null,"require_email":false,"short_description":null,"slug":"cryptopunks","telegram_url":null,"twitter_username":"larvalabs","instagram_username":null,"wiki_url":null},"decimals":null,"token_metadata":"","owner":{"user":null,"profile_img_url":"https://storage.googleapis.com/opensea-static/opensea-profile/32.png","address":"0xb88f61e6fbda83fbfffabe364112137480398018","config":""},"sell_orders":null,"creator":{"user":null,"profile_img_url":"https://storage.googleapis.com/opensea-static/opensea-profile/22.png","address":"0xc352b534e8b987e036a93539fd6897f53488e56a","config":""},"traits":[{"trait_type":"type","value":"Male","display_type":null,"max_value":null,"trait_count":6039,"order":null},{"trait_type":"accessory","value":"Mohawk","display_type":null,"max_value":null,"trait_count":441,"order":null},{"trait_type":"accessory","value":"Smile","display_type":null,"max_value":null,"trait_count":238,"order":null}],"last_sale":{"asset":{"token_id":"1","decimals":null},"asset_bundle":null,"event_type":"successful","event_timestamp":"2020-11-30T18:44:26","auction_type":null,"total_price":"60000000000000000000","payment_token":{"id":1,"symbol":"ETH","address":"0x0000000000000000000000000000000000000000","image_url":"https://storage.opensea.io/files/6f8e2979d428180222796ff4a33ab929.svg","name":"Ether","decimals":18,"eth_price":"1.000000000000000","usd_price":"4022.070000000000164000"},"transaction":{"block_hash":"0xe3eacbc6f4d6bb43525b69baffe09477f452f0f5e1a5b1d5232dc23b6cb176cb","block_number":"11361817","from_account":{"user":{"username":"GoWestBTC"},"profile_img_url":"https://storage.googleapis.com/opensea-static/opensea-profile/4.png","address":"0xee3766e4f996dc0e0f8c929954eaafef3441de87","config":""},"id":63641091,"timestamp":"2020-11-30T18:44:26","to_account":{"user":null,"profile_img_url":"https://storage.googleapis.com/opensea-static/opensea-profile/23.png","address":"0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb","config":""},"transaction_hash":"0xf4af5563f3c4c3b26dee3ab027902f113bee5985b28d9ede5b81ab42b46abb30","transaction_index":"58"},"created_date":"2020-11-30T18:45:11.476313","quantity":"1"},"top_bid":null,"listing_date":null,"is_presale":false,"transfer_fee_payment_token":null,"transfer_fee":null,"related_assets":[],"orders":[],"auctions":[],"supports_wyvern":false,"top_ownerships":[{"owner":{"user":null,"profile_img_url":"https://storage.googleapis.com/opensea-static/opensea-profile/32.png","address":"0xb88f61e6fbda83fbfffabe364112137480398018","config":""},"quantity":"1"}],"ownership":null,"highest_buyer_commitment":null}`
	raribleResp := RaribleData{}
	raribleAlternativeContent := `{"id":"0x1a92f7381b9f03921564a437210bb9396471050c:7034","contract":"0x1a92f7381b9f03921564a437210bb9396471050c","tokenId":7034,"unlockable":false,"creators":[{"account":"0x8ab83d869f2bc250b781d26f6584fd5c562fdd9d","value":10000}],"supply":1,"lazySupply":0,"owners":["0xc5398e0c67610e8ed4ec95b76fa709bb6db7131a"],"royalties":[],"date":"2021-08-20T23:53:15Z","pending":[],"meta":{"name":"Cool Cat #7034","description":"Cool Cats is a collection of 9,999 randomly generated and stylistically curated NFTs that exist on the Ethereum Blockchain. Cool Cat holders can participate in exclusive events such as NFT claims, raffles, community giveaways, and more. Remember, all cats are cool, but some are cooler than others. Visit [www.coolcatsnft.com](https://www.coolcatsnft.com/) to learn more.","attributes":[{"key":"body","value":"blue cat skin"},{"key":"hats","value":"arrowhead"},{"key":"shirt","value":"tanktop pink"},{"key":"face","value":"angry"},{"key":"tier","value":"cool_1"}],"image":{"url":{"ORIGINAL":"ipfs://ipfs/QmZC7nsUnzZaU5CqNv7SQVdg4HXWS22sdtdY6zyscXgfRS","BIG":"https://lh3.googleusercontent.com/Np-KIGet3zkI_poid2z-K3l3LSMCOlW1ZgDBosS2tYD_6AjcmDCqtJv_QWywvimme-2Muuv4kqdr9VZJFGrQTKpRkB5DK3vTWulgqcw","PREVIEW":"https://lh3.googleusercontent.com/Np-KIGet3zkI_poid2z-K3l3LSMCOlW1ZgDBosS2tYD_6AjcmDCqtJv_QWywvimme-2Muuv4kqdr9VZJFGrQTKpRkB5DK3vTWulgqcw=s250"},"meta":{"PREVIEW":{"type":"image/png","width":250,"height":250}}}},"bestSellOrder":{"type":"RARIBLE_V2","maker":"0xc5398e0c67610e8ed4ec95b76fa709bb6db7131a","make":{"assetType":{"assetClass":"ERC721","contract":"0x1a92f7381b9f03921564a437210bb9396471050c","tokenId":7034},"value":1,"valueDecimal":1},"take":{"assetType":{"assetClass":"ETH"},"value":50000000000000000000,"valueDecimal":50.000000000000000000},"fill":0,"fillValue":0E-18,"makeStock":1,"makeStockValue":1,"cancelled":false,"salt":"0xa6792092bba9a1cf1ab69376af65508373bca318013494c5f13b3782c9b002e4","signature":"0x2338adc21bd53089d323b659b1dc90d20b42802aa0820189b87409ed8d197b1431de02c8c93e730e9112425f564b38737c1eb4cf8a8110a8e3e09a6b735efecf1c","createdAt":"2021-10-10T02:34:18Z","lastUpdateAt":"2021-10-10T02:34:18Z","pending":[],"hash":"0x9cd2ca661b7eef441bfc0da565116706d87efcf4aee48f86dcb6aea79c38b3ae","makeBalance":0,"makePriceUsd":193842.0505212929000000000000000000,"priceHistory":[{"date":"2021-10-10T02:34:18Z","makeValue":1,"takeValue":50.000000000000000000}],"data":{"dataType":"RARIBLE_V2_DATA_V1","payouts":[],"originFees":[{"account":"0x1cf0df2a5a20cd61d68d4489eebbf85b8d39e18a","value":250}]}},"bestBidOrder":{"type":"RARIBLE_V2","maker":"0xe4e0ac30637525afa371730b5a3106b49bdc64e1","make":{"assetType":{"assetClass":"ERC20","contract":"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"},"value":19000000000000000,"valueDecimal":0.019000000000000000},"take":{"assetType":{"assetClass":"ERC721","contract":"0x1a92f7381b9f03921564a437210bb9396471050c","tokenId":7034},"value":1,"valueDecimal":1},"fill":0,"fillValue":0,"makeStock":19000000000000000,"makeStockValue":0.019000000000000000,"cancelled":false,"salt":"0x8a75230e77918e5fc8c307fa17902393672865d3891057ae9e7d3297bb5a3544","signature":"0x188a0fb85fa415176c307e7dc4b95bdf596cfa9abc121667e8d019519c05bbe806e28c0b77f03f5db1aed98b9c8ec5ae2ffd5567d8a08619e0dd274b3e5cfb9a1b","createdAt":"2021-10-19T19:07:44.382Z","lastUpdateAt":"2021-10-19T19:07:44.382Z","pending":[],"hash":"0x91fc392b64e3700a824e321045dab51345ee90037cb4996cdcf12a5d984a0cbf","makeBalance":0,"takePriceUsd":72.573663506131547000000000000000,"priceHistory":[{"date":"2021-10-19T19:07:44.382Z","makeValue":0.019000000000000000,"takeValue":1}],"data":{"dataType":"RARIBLE_V2_DATA_V1","payouts":[],"originFees":[{"account":"0x1cf0df2a5a20cd61d68d4489eebbf85b8d39e18a","value":250}]}},"totalStock":1,"sellers":1}`

	if !bPlan {
		url := "https://api.opensea.io/api/v1/asset/0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb/1/"

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("Error req %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("Error res %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		defer res.Body.Close()

		body, err = ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Error reading body %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		// fmt.Println(res)
		if res.StatusCode != 200 {
			fmt.Printf("Error status code %s", res.Status)
			ctx.JSON(http.StatusInternalServerError, res.Status)
			return
		}

	} else {
		body = []byte(openseaAlternativeContent)
	}

	err := json.Unmarshal(body, &openseaResp)
	if err != nil {
		fmt.Printf("Error data format %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	resp.OpenseaDetails = openseaResp

	if !bPlan {
		url := "https://ethereum-api.rarible.org/v0.1/nft-order/items/0x1a92f7381b9f03921564a437210bb9396471050c:7034"

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("Rarible Error req %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("Rarible Error res %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		defer res.Body.Close()

		body, err = ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Rarible Error reading body %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		// fmt.Println(res)
		if res.StatusCode != 200 {
			fmt.Printf("Rarible Error status code %s", res.Status)
			ctx.JSON(http.StatusInternalServerError, res.Status)
			return
		}

	} else {
		body = []byte(raribleAlternativeContent)
	}

	err = json.Unmarshal(body, &raribleResp)
	if err != nil {
		fmt.Printf("Rarible error data format %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	resp.RaribleDetails = raribleResp

	ctx.JSON(http.StatusOK, resp)
}
