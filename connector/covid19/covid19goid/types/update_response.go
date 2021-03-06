package types

// UpdateResponse represent Update Response
type UpdateResponse struct {
	Data   Data   `json:"data"`
	Update Update `json:"update"`
}

// Data represent Data (Indonesian lang.)
type Data struct {
	ID                   int64 `json:"id"`
	JumlahOdp            int64 `json:"jumlah_odp"`
	JumlahPDP            int64 `json:"jumlah_pdp"`
	TotalSpesimen        int64 `json:"total_spesimen"`
	TotalSpesimenNegatif int64 `json:"total_spesimen_negatif"`
}

// Update represent Update (new/total cases) (Indonesian lang.)
type Update struct {
	Penambahan PenambahanUpdate `json:"penambahan"`
	Harian     []Harian         `json:"harian"`
	Total      Total            `json:"total"`
}

// Harian represent daily cases (Indonesian lang.)
type Harian struct {
	KeyAsString        string `json:"key_as_string"`
	Key                int64  `json:"key"`
	DocCount           int64  `json:"doc_count"`
	JumlahMeninggal    Jumlah `json:"jumlah_meninggal"`
	JumlahSembuh       Jumlah `json:"jumlah_sembuh"`
	JumlahPositif      Jumlah `json:"jumlah_positif"`
	JumlahDirawat      Jumlah `json:"jumlah_dirawat"`
	JumlahPositifKum   Jumlah `json:"jumlah_positif_kum"`
	JumlahSembuhKum    Jumlah `json:"jumlah_sembuh_kum"`
	JumlahMeninggalKum Jumlah `json:"jumlah_meninggal_kum"`
	JumlahDirawatKum   Jumlah `json:"jumlah_dirawat_kum"`
}

// Jumlah represent count/summary (Indonesian lang.)
type Jumlah struct {
	Value int64 `json:"value"`
}

// PenambahanUpdate represent daily new cases (Indonesian lang.)
type PenambahanUpdate struct {
	JumlahPositif   int64  `json:"jumlah_positif"`
	JumlahMeninggal int64  `json:"jumlah_meninggal"`
	JumlahSembuh    int64  `json:"jumlah_sembuh"`
	JumlahDirawat   int64  `json:"jumlah_dirawat"`
	Tanggal         string `json:"tanggal"`
	Created         string `json:"created"`
}

// Total represent Total cases (Indonesian lang.)
type Total struct {
	JumlahPositif   int64 `json:"jumlah_positif"`
	JumlahDirawat   int64 `json:"jumlah_dirawat"`
	JumlahSembuh    int64 `json:"jumlah_sembuh"`
	JumlahMeninggal int64 `json:"jumlah_meninggal"`
}
