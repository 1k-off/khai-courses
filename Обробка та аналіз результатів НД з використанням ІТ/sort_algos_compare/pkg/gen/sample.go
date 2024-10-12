package gen

import (
	"sort_algos_compare/pkg/util"
)

var (
	rootPath   = "./samples"
	pathSmall  = rootPath + "/small"
	pathMedium = rootPath + "/medium"
	pathBig    = rootPath + "/big"

	smallSampleCount  = 100
	mediumSampleCount = 10000
	bigSampleCount    = 1000000

	sampleDataValueSmall  = 10
	sampleDataValueMedium = 1000000
	sampleDataValueBig    = 1000000000

	sampleDataRandomMinValueSmall  = 10
	sampleDataRandomMaxValueSmall  = 1000
	sampleDataRandomMinValueMedium = 10000
	sampleDataRandomMaxValueMedium = 10000000
	sampleDataRandomMinValueBig    = 10000000
	sampleDataRandomMaxValueBig    = 10000000000
)

func SampleData() {
	// Create directories if not exist
	_ = util.CreateDirIfNotExist(rootPath)
	_ = util.CreateDirIfNotExist(pathSmall)
	_ = util.CreateDirIfNotExist(pathMedium)
	_ = util.CreateDirIfNotExist(pathBig)

	// Generate sample data
	// Zeros
	util.ToNamedJson(Zeros(smallSampleCount), pathSmall+"/zeros")
	util.ToNamedJson(Zeros(mediumSampleCount), pathMedium+"/zeros")
	util.ToNamedJson(Zeros(bigSampleCount), pathBig+"/zeros")
	// Ones
	util.ToNamedJson(Ones(smallSampleCount), pathSmall+"/ones")
	util.ToNamedJson(Ones(mediumSampleCount), pathMedium+"/ones")
	util.ToNamedJson(Ones(bigSampleCount), pathBig+"/ones")
	// Ordered
	util.ToNamedJson(Ordered(smallSampleCount, sampleDataValueSmall, "asc"), pathSmall+"/ordered_asc_small")
	util.ToNamedJson(Ordered(smallSampleCount, sampleDataValueMedium, "asc"), pathSmall+"/ordered_asc_medium")
	util.ToNamedJson(Ordered(smallSampleCount, sampleDataValueBig, "asc"), pathSmall+"/ordered_asc_big")
	util.ToNamedJson(Ordered(mediumSampleCount, sampleDataValueSmall, "asc"), pathMedium+"/ordered_asc_small")
	util.ToNamedJson(Ordered(mediumSampleCount, sampleDataValueMedium, "asc"), pathMedium+"/ordered_asc_medium")
	util.ToNamedJson(Ordered(mediumSampleCount, sampleDataValueBig, "asc"), pathMedium+"/ordered_asc_big")
	util.ToNamedJson(Ordered(bigSampleCount, sampleDataValueSmall, "asc"), pathBig+"/ordered_asc_small")
	util.ToNamedJson(Ordered(bigSampleCount, sampleDataValueMedium, "asc"), pathBig+"/ordered_asc_medium")
	util.ToNamedJson(Ordered(bigSampleCount, sampleDataValueBig, "asc"), pathBig+"/ordered_asc_big")
	util.ToNamedJson(Ordered(smallSampleCount, sampleDataValueSmall, "desc"), pathSmall+"/ordered_desc_small")
	util.ToNamedJson(Ordered(smallSampleCount, sampleDataValueMedium, "desc"), pathSmall+"/ordered_desc_medium")
	util.ToNamedJson(Ordered(smallSampleCount, sampleDataValueBig, "desc"), pathSmall+"/ordered_desc_big")
	util.ToNamedJson(Ordered(mediumSampleCount, sampleDataValueSmall, "desc"), pathMedium+"/ordered_desc_small")
	util.ToNamedJson(Ordered(mediumSampleCount, sampleDataValueMedium, "desc"), pathMedium+"/ordered_desc_medium")
	util.ToNamedJson(Ordered(mediumSampleCount, sampleDataValueBig, "desc"), pathMedium+"/ordered_desc_big")
	util.ToNamedJson(Ordered(bigSampleCount, sampleDataValueSmall, "desc"), pathBig+"/ordered_desc_small")
	util.ToNamedJson(Ordered(bigSampleCount, sampleDataValueMedium, "desc"), pathBig+"/ordered_desc_medium")
	util.ToNamedJson(Ordered(bigSampleCount, sampleDataValueBig, "desc"), pathBig+"/ordered_desc_big")
	// Random
	util.ToNamedJson(Random(smallSampleCount, sampleDataRandomMinValueSmall, sampleDataRandomMaxValueSmall), pathSmall+"/random_small")
	util.ToNamedJson(Random(smallSampleCount, sampleDataRandomMinValueMedium, sampleDataRandomMaxValueMedium), pathSmall+"/random_medium")
	util.ToNamedJson(Random(smallSampleCount, sampleDataRandomMinValueBig, sampleDataRandomMaxValueBig), pathSmall+"/random_big")
	util.ToNamedJson(Random(mediumSampleCount, sampleDataRandomMinValueSmall, sampleDataRandomMaxValueSmall), pathMedium+"/random_small")
	util.ToNamedJson(Random(mediumSampleCount, sampleDataRandomMinValueMedium, sampleDataRandomMaxValueMedium), pathMedium+"/random_medium")
	util.ToNamedJson(Random(mediumSampleCount, sampleDataRandomMinValueBig, sampleDataRandomMaxValueBig), pathMedium+"/random_big")
	util.ToNamedJson(Random(bigSampleCount, sampleDataRandomMinValueSmall, sampleDataRandomMaxValueSmall), pathBig+"/random_small")
	util.ToNamedJson(Random(bigSampleCount, sampleDataRandomMinValueMedium, sampleDataRandomMaxValueMedium), pathBig+"/random_medium")
	util.ToNamedJson(Random(bigSampleCount, sampleDataRandomMinValueBig, sampleDataRandomMaxValueBig), pathBig+"/random_big")
}
