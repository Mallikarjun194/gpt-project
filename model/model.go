package model

type ProductList struct {
	ID                         string   `json: "id"` // The Product Page Instance id in CMS. Pattern: ^\\d{6}$
	GenericChoiceID            string   `json: "genericChoiceId"`
	URL                        string   `json: "url"`    // URL for this product on the site.
	Family                     string   `json: "family"` // Product family, e.g. 'Body by Victoria' (AKA 'brand').
	Name                       string   `json: "name"`   // Product name, e.g. 'Perfect Coverage Bra'. (AKA 'short description'.)
	Rating                     float64  `json: "rating"` // Average product rating, out of 5.
	TotalReviewCount           int      `json: "totalReviewCount"`
	GenericLabel               string   `json: "genericLabel,omitempty"`     // "Label for this selector, e.g. 'Color' or 'Fragrance'."
	SwatchCount                int64    `json: "swatchCount"`                // Number of available swatches for the product.
	SwatchLabel                string   `json: "swatchLabel"`                // The label for product swatches (color, scent, team, etc.) from CMS
	Price                      string   `json: "price"`                      // Original price, formatted.
	AltPrices                  []string `json: "altPrices"`                  // An array of promotional or alternate prices for an item
	SalePrice                  string   `json: "salePrice"`                  // Sale price, formatted.
	CollectionShortDescription string   `json: "collectionShortDescription"` // Collection short description
	DisplayGenericDescription  bool     `json: "displayGenericDescription"`
	ImageURL                   string   `json: "masterStyleId,omitempty"`
}

type Input struct {
	FileName string `json:"filename"`
}

type Result struct {
	Msg string `json:"msg"`
}

type Error struct {
	Message string `json:"message"`
}

// Request ...
type Request []string

type Badge struct {
	Type     string `json:"type"`
	Location string `json:"location"`
	Image    string `json:"image,omitempty"`
	Text     string `json:"text,omitempty"`
}

type BadgesResponse map[string]interface{}

type Msg struct {
	Msg string `json:"msg"`
}
