/*
Package metar includes abstractions to help parse METAR strings.
*/
package metar

import "time"

// Type of METAR report
type MetarType int8

const (
	// TypeMetar indicates that the report is of type METAR
	TypeMetar MetarType = iota
	// TypeSpeci specifies that the report is of type SPECI
	TypeSpeci
)

type SpeedUnit int

const (
	SpeedUnitKnots SpeedUnit = iota
	SpeedUnitMetersPerSecond
)

type VisibilityUnit int

const (
	// UnitStatuteMiles indicates distance is measured in SM (Only used in USA)
	UnitStatuteMiles VisibilityUnit = iota
	// UnitMeters indicates distance is measured in Meters
	UnitMeters
)

type VisibilityModifier int

const (
	VisibilityModifierOrLess VisibilityModifier = iota
	VisibilityModifierExactly
	VisibilityModifierOrMore
)

type VisibilityTrend int

const (
	VisibilityTrendUp VisibilityTrend = iota
	VisibilityTrendDown
	VisibilityTrendNil
	VisibilityTrendNotProvided
)

type WeatherIntensity int

const (
	WeatherIntensityLight WeatherIntensity = iota
	WeatherIntensityModerate
	WeatherIntensityHeavy
)

type WeatherDescriptor int

const (
	WeatherDescriptorNone WeatherDescriptor = iota
	WeatherDescriptorShallow
	WeatherDescriptorPartial
	WeatherDescriptorPatches
	WeatherDescriptorLowDrifting
	WeatherDescriptorBlowing
	WeatherDescriptorShowers
	WeatherDescriptorThunderstorm
	WeatherDescriptorFreezing
)

type WeatherPrecipitation int

const (
	WeatherPrecipitationNone WeatherPrecipitation = iota
	WeatherPrecipitationDrizzle
	WeatherPrecipitationRain
	WeatherPrecipitationSnow
	WeatherPrecipitationSnowGrains
	WeatherPrecipitationIceCrystals
	WeatherPrecipitationIcePellets
	WeatherPrecipitationHail
	WeatherPrecipitationSmallHailandOrSnowPellets
	WeatherPrecipitationUnknownPrecipitation
)

type WeatherObscuration int

const (
	WeatherObscurationNone WeatherObscuration = iota
	WeatherObscurationMist
	WeatherObscurationFog
	WeatherObscurationSmoke
	WeatherObscurationVolcanicAsh
	WeatherObscurationWidespreadDust
	WeatherObscurationSand
	WeatherObscurationHaze
	WeatherObscurationSpray
)

type WeatherOtherPhen int

const (
	WeatherOtherPhenNone WeatherOtherPhen = iota
	WeatherOtherPhenWellDevelopedDustSandWhirls
	WeatherOtherPhenSqualls
	WeatherOtherPhenFunnelCloudTornadoWaterspout
	WeatherOtherPhenSandstorm
	WeatherOtherPhenDuststorm
)

type CloudAmount int

const (
	CloudAmountFew CloudAmount = iota
	CloudAmountScattered
	CloudAmountBroken
	CloudAmountOvercast
	CloudAmountNilSignificant
	CloudAmountNilDetected
)

type CloudType int

const (
	CloudTypeNone CloudType = iota
	CloudTypeToweringCumulus
	CloudTypeCumulonimbusOrashowerThunderstorm
	CloudTypeAltocumulusCastellanus
)

type PressureUnit int

const (
	PressureUnitHectoPascals PressureUnit = iota
	PressureUnitInchesOfMercury
)

type Report struct {
	// Station is the ICAO location indicator that this report describes.
	Station string

	// DateTime represents the date and time (UTC) of this report.
	DateTime time.Time

	// Auto indicates if the report contains only automated observations.
	Auto bool

	// Wind information for the report
	Wind struct {

		// Variable indicates that the direction cannot be determined.
		Variable bool

		// Source of the wind in degrees from true north.
		Source int

		// VarianceFrom is the minimum observed wind direction represented
		// in degrees from true north. Only given if direction varies
		// substantially
		VarianceFrom int

		// VarianceFrom is the maximum observed wind direction represented
		// in degrees from true north. Only given if direction varies
		// substantially
		VarianceTo int

		// Speed is the mean value for speed observed in the sampling period.
		Speed struct {
			Speed int
			Unit  SpeedUnit
		}

		// Gust is the maximum speed measured in the sampling period.
		Gust int
	}

	// Cavok indicates Cloud and Visbility OK. If set to true then Visibility
	// RunwayVisualRange, Weather and Cloud sections can be ignored.
	Cavok bool

	// Visibility describes the visibility conditions of the report.
	Visibility struct {
		Distance   int
		Unit       int
		Modifier   VisibilityModifier
		ToDistance int
		ToModifier VisibilityModifier
		Trend      VisibilityTrend
	}

	RunwayVisualRange struct {
		Runway     string
		Unit       VisibilityUnit
		Visibility float32
		// Modifier is used for expressing bounds within a distance
		// measurement. Examples: 3000 meters OR MORE, 3/4 Statute miles OR
		// LESS etc.
		Modifier VisibilityModifier
	}

	Weather struct {
		Descriptor    WeatherDescriptor
		Precipitation WeatherPrecipitation
		Obscuration   WeatherObscuration
		Other         WeatherOtherPhen
		Vecinity      bool
	}

	Clouds []struct {
		Amount CloudAmount
		// Height is how high the cloud is from the airfield (in feet)
		Height int
		Type   CloudType
	}

	Temperature struct {
		Temperature int
		DewPoint    int
	}

	Pressure struct {
		Qnh  int
		Unit PressureUnit
	}

	Supplementary struct {
		RecentWeather string
		WindSheer     int
	}

	Remarks string
}
