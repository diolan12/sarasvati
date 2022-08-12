package main

// emsifa rest api struct
// provinsi
type MsifaProvince struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// kabupaten/kota
type MsifaRegency struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}

// kecamatan
type MsifaDistrict struct {
	ID        string `json:"id"`
	RegencyID string `json:"regency_id"`
	Name      string `json:"name"`
}

// kemenag rest api struct
type Pondok struct {
	ID       string        `json:"id"`
	Nama     string        `json:"nama"`
	NSPP     string        `json:"nspp"`
	Alamat   string        `json:"alamat"`
	Kyai     string        `json:"kyai"`
	District MsifaDistrict `json:"district"`
}

// dapodik resp api struct
type DapodikRegion struct {
	Name            string `json:"nama"`
	KodeWilayah     string `json:"kode_wilayah"`
	IDLevelWilayah  int    `json:"id_level_wilayah"`
	MstLevelWilayah int    `json:"mst_level_wilayah"`
	TK              int    `json:"tk"`
	TKN             int    `json:"tk_n"`
	TKS             int    `json:"tk_s"`
	KB              int    `json:"kb"`
	KBN             int    `json:"kb_n"`
	KBS             int    `json:"kb_s"`
	TPA             int    `json:"tpa"`
	TPAN            int    `json:"tpa_n"`
	TPAS            int    `json:"tpa_s"`
	SPS             int    `json:"sps"`
	SPSN            int    `json:"sps_n"`
	SPSS            int    `json:"sps_s"`
	PKBM            int    `json:"pkbm"`
	PKBMN           int    `json:"pkbm_n"`
	PKBMS           int    `json:"pkbm_s"`
	SKB             int    `json:"skb"`
	SKBN            int    `json:"skb_n"`
	SKBS            int    `json:"skb_s"`
	SD              int    `json:"sd"`
	SDN             int    `json:"sd_n"`
	SDS             int    `json:"sd_s"`
	SMP             int    `json:"smp"`
	SMPN            int    `json:"smp_n"`
	SMPS            int    `json:"smp_s"`
	SMA             int    `json:"sma"`
	SMAN            int    `json:"sma_n"`
	SMAS            int    `json:"sma_s"`
	SMK             int    `json:"smk"`
	SMKN            int    `json:"smk_n"`
	SMKS            int    `json:"smk_s"`
	SLB             int    `json:"slb"`
	SLBN            int    `json:"slb_n"`
	SLBS            int    `json:"slb_s"`
	Total           int    `json:"sekolah"`
	TotalN          int    `json:"sekolah_n"`
	TotalS          int    `json:"sekolah_s"`
}
type School struct {
	Nama      string `json:"nama"`
	ID        string `json:"sekolah_id"`
	IDEncrypt string `json:"sekolah_id_enkrip"`
	Alamat    string `json:"alamat"`
}

// app personalized struct
type Province struct {
	ID      string        `json:"id"`
	Name    string        `json:"name"`
	Dapodik DapodikRegion `json:"dapodik"`
}
type Regency struct {
	ID         string        `json:"id"`
	ProvinceID string        `json:"province_id"`
	Name       string        `json:"name"`
	Dapodik    DapodikRegion `json:"dapodik"`
}
type District struct {
	ID        string        `json:"id"`
	RegencyID string        `json:"regency_id"`
	Name      string        `json:"name"`
	Dapodik   DapodikRegion `json:"dapodik"`
}
type Village struct {
	ID         string `json:"id"`
	DistrictID string `json:"district_id"`
	Name       string `json:"name"`
}
