package constants

const (
	HaNoiRegion     = "HaNoi"
	HoCHiMinhRegion = "HoChiMinh"
	VcHaNoiRegion   = "VC-HaNoi"
)

var RegionMapping = map[string]string{
	"hn":        HaNoiRegion,
	"hanoi":     HaNoiRegion,
	"hcm":       HoCHiMinhRegion,
	"hochiminh": HoCHiMinhRegion,
	"vc-hanoi":  VcHaNoiRegion,
}
