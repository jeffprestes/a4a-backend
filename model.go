package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"
)

// NFTItemResponse is the response to NFT Item query
type NFTItemResponse struct {
	ID             string
	CoordX         string
	CoordY         string
	OpenseaDetails OpenseaData
	RaribleDetails RaribleData
}

type RaribleData struct {
	ID            string        `json:"id"`
	Contract      string        `json:"contract"`
	TokenID       int           `json:"tokenId"`
	Unlockable    bool          `json:"unlockable"`
	Creators      []Creators    `json:"creators"`
	Supply        int           `json:"supply"`
	LazySupply    int           `json:"lazySupply"`
	Owners        []string      `json:"owners"`
	Royalties     []interface{} `json:"royalties"`
	Date          time.Time     `json:"date"`
	Pending       []interface{} `json:"pending"`
	Meta          Meta          `json:"meta"`
	BestSellOrder BestSellOrder `json:"bestSellOrder"`
	BestBidOrder  BestBidOrder  `json:"bestBidOrder"`
	TotalStock    int           `json:"totalStock"`
	Sellers       int           `json:"sellers"`
}

type Creators struct {
	Account string `json:"account"`
	Value   int    `json:"value"`
}

type Attributes struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type URL struct {
	ORIGINAL string `json:"ORIGINAL"`
	BIG      string `json:"BIG"`
	PREVIEW  string `json:"PREVIEW"`
}

type PREVIEW struct {
	Type   string `json:"type"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type MetaPreview struct {
	PREVIEW PREVIEW `json:"PREVIEW"`
}

type Image struct {
	URL  URL         `json:"url"`
	Meta MetaPreview `json:"meta"`
}

type Meta struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Attributes  []Attributes `json:"attributes"`
	Image       Image        `json:"image"`
}

type AssetType struct {
	AssetClass string `json:"assetClass"`
	Contract   string `json:"contract"`
	TokenID    int    `json:"tokenId"`
}

type Make struct {
	AssetType    AssetType `json:"assetType"`
	Value        int       `json:"value"`
	ValueDecimal float64   `json:"valueDecimal"`
}

type Take struct {
	AssetType    AssetType `json:"assetType"`
	Value        BigInt    `json:"value"`
	ValueDecimal float64   `json:"valueDecimal"`
}

type PriceHistory struct {
	Date      time.Time `json:"date"`
	MakeValue BigInt    `json:"makeValue"`
	TakeValue BigInt    `json:"takeValue"`
}

type OriginFees struct {
	Account string `json:"account"`
	Value   int    `json:"value"`
}

type Data struct {
	DataType   string        `json:"dataType"`
	Payouts    []interface{} `json:"payouts"`
	OriginFees []OriginFees  `json:"originFees"`
}

type BestSellOrder struct {
	Type           string         `json:"type"`
	Maker          string         `json:"maker"`
	Make           Make           `json:"make"`
	Take           Take           `json:"take"`
	Fill           int            `json:"fill"`
	FillValue      BigInt         `json:"fillValue"`
	MakeStock      int            `json:"makeStock"`
	MakeStockValue int            `json:"makeStockValue"`
	Cancelled      bool           `json:"cancelled"`
	Salt           string         `json:"salt"`
	Signature      string         `json:"signature"`
	CreatedAt      time.Time      `json:"createdAt"`
	LastUpdateAt   time.Time      `json:"lastUpdateAt"`
	Pending        []interface{}  `json:"pending"`
	Hash           string         `json:"hash"`
	MakeBalance    int            `json:"makeBalance"`
	MakePriceUsd   float64        `json:"makePriceUsd"`
	PriceHistory   []PriceHistory `json:"priceHistory"`
	Data           Data           `json:"data"`
}

type BestBidOrder struct {
	Type           string         `json:"type"`
	Maker          string         `json:"maker"`
	Make           Make           `json:"make"`
	Take           Take           `json:"take"`
	Fill           int            `json:"fill"`
	FillValue      int            `json:"fillValue"`
	MakeStock      int64          `json:"makeStock"`
	MakeStockValue float64        `json:"makeStockValue"`
	Cancelled      bool           `json:"cancelled"`
	Salt           string         `json:"salt"`
	Signature      string         `json:"signature"`
	CreatedAt      time.Time      `json:"createdAt"`
	LastUpdateAt   time.Time      `json:"lastUpdateAt"`
	Pending        []interface{}  `json:"pending"`
	Hash           string         `json:"hash"`
	MakeBalance    int            `json:"makeBalance"`
	TakePriceUsd   float64        `json:"takePriceUsd"`
	PriceHistory   []PriceHistory `json:"priceHistory"`
	Data           Data           `json:"data"`
}

type OpenseaData struct {
	ID                      int             `json:"id"`
	TokenID                 string          `json:"token_id"`
	NumSales                int             `json:"num_sales"`
	BackgroundColor         interface{}     `json:"background_color"`
	ImageURL                string          `json:"image_url"`
	ImagePreviewURL         string          `json:"image_preview_url"`
	ImageThumbnailURL       string          `json:"image_thumbnail_url"`
	ImageOriginalURL        string          `json:"image_original_url"`
	AnimationURL            interface{}     `json:"animation_url"`
	AnimationOriginalURL    interface{}     `json:"animation_original_url"`
	Name                    string          `json:"name"`
	Description             interface{}     `json:"description"`
	ExternalLink            string          `json:"external_link"`
	AssetContract           AssetContract   `json:"asset_contract"`
	Permalink               string          `json:"permalink"`
	Collection              Collection      `json:"collection"`
	Decimals                interface{}     `json:"decimals"`
	TokenMetadata           string          `json:"token_metadata"`
	Owner                   Owner           `json:"owner"`
	SellOrders              interface{}     `json:"sell_orders"`
	Creator                 Creator         `json:"creator"`
	Traits                  []Traits        `json:"traits"`
	LastSale                LastSale        `json:"last_sale"`
	TopBid                  interface{}     `json:"top_bid"`
	ListingDate             interface{}     `json:"listing_date"`
	IsPresale               bool            `json:"is_presale"`
	TransferFeePaymentToken interface{}     `json:"transfer_fee_payment_token"`
	TransferFee             interface{}     `json:"transfer_fee"`
	RelatedAssets           []interface{}   `json:"related_assets"`
	Orders                  []interface{}   `json:"orders"`
	Auctions                []interface{}   `json:"auctions"`
	SupportsWyvern          bool            `json:"supports_wyvern"`
	TopOwnerships           []TopOwnerships `json:"top_ownerships"`
	Ownership               interface{}     `json:"ownership"`
	HighestBuyerCommitment  interface{}     `json:"highest_buyer_commitment"`
}

type AssetContract struct {
	Address                     string      `json:"address"`
	AssetContractType           string      `json:"asset_contract_type"`
	CreatedDate                 string      `json:"created_date"`
	Name                        string      `json:"name"`
	NftVersion                  string      `json:"nft_version"`
	OpenseaVersion              interface{} `json:"opensea_version"`
	Owner                       interface{} `json:"owner"`
	SchemaName                  string      `json:"schema_name"`
	Symbol                      string      `json:"symbol"`
	TotalSupply                 interface{} `json:"total_supply"`
	Description                 string      `json:"description"`
	ExternalLink                string      `json:"external_link"`
	ImageURL                    string      `json:"image_url"`
	DefaultToFiat               bool        `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int         `json:"seller_fee_basis_points"`
	PayoutAddress               interface{} `json:"payout_address"`
}

type PaymentTokens struct {
	ID       int     `json:"id"`
	Symbol   string  `json:"symbol"`
	Address  string  `json:"address"`
	ImageURL string  `json:"image_url"`
	Name     string  `json:"name"`
	Decimals int     `json:"decimals"`
	EthPrice int     `json:"eth_price"`
	UsdPrice float64 `json:"usd_price"`
}

type PrimaryAssetContracts struct {
	Address                     string      `json:"address"`
	AssetContractType           string      `json:"asset_contract_type"`
	CreatedDate                 string      `json:"created_date"`
	Name                        string      `json:"name"`
	NftVersion                  string      `json:"nft_version"`
	OpenseaVersion              interface{} `json:"opensea_version"`
	Owner                       interface{} `json:"owner"`
	SchemaName                  string      `json:"schema_name"`
	Symbol                      string      `json:"symbol"`
	TotalSupply                 interface{} `json:"total_supply"`
	Description                 string      `json:"description"`
	ExternalLink                string      `json:"external_link"`
	ImageURL                    string      `json:"image_url"`
	DefaultToFiat               bool        `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int         `json:"seller_fee_basis_points"`
	PayoutAddress               interface{} `json:"payout_address"`
}

type Stats struct {
	OneDayVolume          float64 `json:"one_day_volume"`
	OneDayChange          float64 `json:"one_day_change"`
	OneDaySales           int     `json:"one_day_sales"`
	OneDayAveragePrice    float64 `json:"one_day_average_price"`
	SevenDayVolume        float64 `json:"seven_day_volume"`
	SevenDayChange        float64 `json:"seven_day_change"`
	SevenDaySales         int     `json:"seven_day_sales"`
	SevenDayAveragePrice  float64 `json:"seven_day_average_price"`
	ThirtyDayVolume       float64 `json:"thirty_day_volume"`
	ThirtyDayChange       float64 `json:"thirty_day_change"`
	ThirtyDaySales        int     `json:"thirty_day_sales"`
	ThirtyDayAveragePrice float64 `json:"thirty_day_average_price"`
	TotalVolume           float64 `json:"total_volume"`
	TotalSales            int     `json:"total_sales"`
	TotalSupply           int     `json:"total_supply"`
	Count                 int     `json:"count"`
	NumOwners             int     `json:"num_owners"`
	AveragePrice          float64 `json:"average_price"`
	NumReports            int     `json:"num_reports"`
	MarketCap             float64 `json:"market_cap"`
	FloorPrice            int     `json:"floor_price"`
}

type DisplayData struct {
	CardDisplayStyle string `json:"card_display_style"`
}

type Collection struct {
	PaymentTokens               []PaymentTokens         `json:"payment_tokens"`
	PrimaryAssetContracts       []PrimaryAssetContracts `json:"primary_asset_contracts"`
	Traits                      Traits                  `json:"traits"`
	Stats                       Stats                   `json:"stats"`
	BannerImageURL              string                  `json:"banner_image_url"`
	ChatURL                     interface{}             `json:"chat_url"`
	CreatedDate                 string                  `json:"created_date"`
	DefaultToFiat               bool                    `json:"default_to_fiat"`
	Description                 string                  `json:"description"`
	DevBuyerFeeBasisPoints      string                  `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     string                  `json:"dev_seller_fee_basis_points"`
	DiscordURL                  string                  `json:"discord_url"`
	DisplayData                 DisplayData             `json:"display_data"`
	ExternalURL                 string                  `json:"external_url"`
	Featured                    bool                    `json:"featured"`
	FeaturedImageURL            string                  `json:"featured_image_url"`
	Hidden                      bool                    `json:"hidden"`
	SafelistRequestStatus       string                  `json:"safelist_request_status"`
	ImageURL                    string                  `json:"image_url"`
	IsSubjectToWhitelist        bool                    `json:"is_subject_to_whitelist"`
	LargeImageURL               string                  `json:"large_image_url"`
	MediumUsername              interface{}             `json:"medium_username"`
	Name                        string                  `json:"name"`
	OnlyProxiedTransfers        bool                    `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string                  `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string                  `json:"opensea_seller_fee_basis_points"`
	PayoutAddress               interface{}             `json:"payout_address"`
	RequireEmail                bool                    `json:"require_email"`
	ShortDescription            interface{}             `json:"short_description"`
	Slug                        string                  `json:"slug"`
	TelegramURL                 interface{}             `json:"telegram_url"`
	TwitterUsername             string                  `json:"twitter_username"`
	InstagramUsername           interface{}             `json:"instagram_username"`
	WikiURL                     interface{}             `json:"wiki_url"`
}

type Owner struct {
	User          interface{} `json:"user"`
	ProfileImgURL string      `json:"profile_img_url"`
	Address       string      `json:"address"`
	Config        string      `json:"config"`
}

type Creator struct {
	User          interface{} `json:"user"`
	ProfileImgURL string      `json:"profile_img_url"`
	Address       string      `json:"address"`
	Config        string      `json:"config"`
}

type Traits struct {
	TraitType   string      `json:"trait_type"`
	Value       string      `json:"value"`
	DisplayType interface{} `json:"display_type"`
	MaxValue    interface{} `json:"max_value"`
	TraitCount  int         `json:"trait_count"`
	Order       interface{} `json:"order"`
}

type Asset struct {
	TokenID  string      `json:"token_id"`
	Decimals interface{} `json:"decimals"`
}

type PaymentToken struct {
	ID       int    `json:"id"`
	Symbol   string `json:"symbol"`
	Address  string `json:"address"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
	Decimals int    `json:"decimals"`
	EthPrice string `json:"eth_price"`
	UsdPrice string `json:"usd_price"`
}

type User struct {
	Username string `json:"username"`
}

type FromAccount struct {
	User          User   `json:"user"`
	ProfileImgURL string `json:"profile_img_url"`
	Address       string `json:"address"`
	Config        string `json:"config"`
}

type ToAccount struct {
	User          interface{} `json:"user"`
	ProfileImgURL string      `json:"profile_img_url"`
	Address       string      `json:"address"`
	Config        string      `json:"config"`
}

type Transaction struct {
	BlockHash        string      `json:"block_hash"`
	BlockNumber      string      `json:"block_number"`
	FromAccount      FromAccount `json:"from_account"`
	ID               int         `json:"id"`
	Timestamp        string      `json:"timestamp"`
	ToAccount        ToAccount   `json:"to_account"`
	TransactionHash  string      `json:"transaction_hash"`
	TransactionIndex string      `json:"transaction_index"`
}

type LastSale struct {
	Asset          Asset        `json:"asset"`
	AssetBundle    interface{}  `json:"asset_bundle"`
	EventType      string       `json:"event_type"`
	EventTimestamp string       `json:"event_timestamp"`
	AuctionType    interface{}  `json:"auction_type"`
	TotalPrice     string       `json:"total_price"`
	PaymentToken   PaymentToken `json:"payment_token"`
	Transaction    Transaction  `json:"transaction"`
	CreatedDate    string       `json:"created_date"`
	Quantity       string       `json:"quantity"`
}

type TopOwnerships struct {
	Owner    Owner  `json:"owner"`
	Quantity string `json:"quantity"`
}

type BigInt struct {
	big.Int
}

func (i *BigInt) UnmarshalJSON(b []byte) error {
	var val interface{}
	err := json.Unmarshal(b, &val)
	if err != nil {
		return err
	}
	var2 := val.(float64)
	var3 := fmt.Sprintf("%f", var2)
	i.SetString(var3, 10)

	return nil
}
