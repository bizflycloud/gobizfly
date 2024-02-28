package utils

import (
	"strings"

	"github.com/bizflycloud/gobizfly/constants"
	gobizflyErr "github.com/bizflycloud/gobizfly/errors"
)

func ParseRegionName(region string) (string, error) {
	lowerRegion := strings.ToLower(region)
	regionName, ok := constants.RegionMapping[lowerRegion]
	if !ok {
		return "", gobizflyErr.InvalidRegion.SetMetadata(map[string]interface{}{"Region": region})
	}
	return regionName, nil
}
