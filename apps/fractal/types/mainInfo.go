package types

type ImagesInfo struct {
	ImageRows    []ImageRow
	RowsCount    int
	ColumnsCount int
}

type ImageRow struct {
	Images []Image
}

type Image struct {
	Base64Coded string
	AltText     string
	Params      string
}
