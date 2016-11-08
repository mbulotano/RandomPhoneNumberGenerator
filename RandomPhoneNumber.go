package main

import(
	"math/rand"
	"time"
	"strconv"
)

//Generates a random phone number from a random country.
func generatePhoneNumber() string {
	var num string
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(1)
	switch i {
	case 0:
		num = generateUSPhoneNumber()
	}
	return num
}

//Generates a number, split into area code, prefix, and suffix
func generatePhoneNumberSplit() (area string, prefix string,suffix string) {
	var area,pre,suf string
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(1)
	switch i {
	case 0:
		area,pre,suf = generateUSPhoneNumberSplit()
	}
	return area,pre,suf
}

//Generates a random phone number for the 50 states inside the U.S.
func generateUSPhoneNumberSplit() (string, string, string) {

	USAreaCodes := []int{201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,224,225,228,229,231,234,236,239,240,242,246,248,250,251,252,253,254,256,260,262,264,267,268,270,276,278,281,283,284,289,301,302,303,304,305,306,307,308,309,310,311,312,313,314,315,316,317,318,319,320,321,323,327,330,331,334,336,337,339,340,341,345,347,351,352,360,361,369,385,386,401,402,403,404,405,406,407,408,409,410,411,412,413,414,415,416,417,418,419,420,423,424,425,434,435,440,441,442,443,445,450,456,464,469,470,473,475,478,479,480,484,500,501,502,503,504,505,506,507,508,509,510,511,512,513,514,515,516,517,518,519,520,530,540,541,546,551,557,559,561,562,563,564,567,570,571,573,574,580,585,586,590,600,601,602,603,604,605,606,607,608,609,610,611,612,613,614,615,616,617,618,619,620,623,626,627,628,630,631,636,641,646,647,649,650,651,657,660,661,662,664,669,670,671,678,679,682,684,700,701,702,703,704,705,706,707,708,709,710,711,712,713,714,715,716,717,718,719,720,724,727,731,732,734,737,740,747,752,754,757,758,760,763,764,765,767,770,772,773,774,775,778,780,781,784,785,786,787,800,801,802,803,804,805,806,807,808,809,810,811,812,813,814,815,816,817,818,819,822,828,830,831,832,833,835,836,843,844,845,847,848,850,856,857,858,859,860,862,863,864,865,866,867,868,869,870,872,876,877,878,880,881,882,888,900,901,902,903,904,905,906,907,908,909,910,911,912,913,914,915,916,917,918,919,920,925,928,931,935,936,937,939,940,941,947,949,951,952,954,956,957,959,969,970,971,972,973,975,978,979,980,984,985,989};
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	areacode := USAreaCodes[r.Intn(len(USAreaCodes))]
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	prefix := r.Intn((999 - 200) + 1) + 200
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	suffix := r.Intn(10000)
	areacodeString := strconv.Itoa(areacode)
	prefixString := strconv.Itoa(prefix)
	suffixString := strconv.Itoa(suffix)

	if suffix > 99 && suffix < 1000 {
		suffixString = "0" + suffixString
	} else if suffix > 9 && suffix < 100 {
		suffixString = "00" + suffixString
	} else if(suffix < 10 && suffix > 0) {
		suffixString = "000" + suffixString
	}

	return areacodeString, prefixString, suffixString, areacodeString+prefixString+suffixString
}

func generateUSPhoneNumber() string {
	USAreaCodes := []int{201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,224,225,228,229,231,234,236,239,240,242,246,248,250,251,252,253,254,256,260,262,264,267,268,270,276,278,281,283,284,289,301,302,303,304,305,306,307,308,309,310,311,312,313,314,315,316,317,318,319,320,321,323,327,330,331,334,336,337,339,340,341,345,347,351,352,360,361,369,385,386,401,402,403,404,405,406,407,408,409,410,411,412,413,414,415,416,417,418,419,420,423,424,425,434,435,440,441,442,443,445,450,456,464,469,470,473,475,478,479,480,484,500,501,502,503,504,505,506,507,508,509,510,511,512,513,514,515,516,517,518,519,520,530,540,541,546,551,557,559,561,562,563,564,567,570,571,573,574,580,585,586,590,600,601,602,603,604,605,606,607,608,609,610,611,612,613,614,615,616,617,618,619,620,623,626,627,628,630,631,636,641,646,647,649,650,651,657,660,661,662,664,669,670,671,678,679,682,684,700,701,702,703,704,705,706,707,708,709,710,711,712,713,714,715,716,717,718,719,720,724,727,731,732,734,737,740,747,752,754,757,758,760,763,764,765,767,770,772,773,774,775,778,780,781,784,785,786,787,800,801,802,803,804,805,806,807,808,809,810,811,812,813,814,815,816,817,818,819,822,828,830,831,832,833,835,836,843,844,845,847,848,850,856,857,858,859,860,862,863,864,865,866,867,868,869,870,872,876,877,878,880,881,882,888,900,901,902,903,904,905,906,907,908,909,910,911,912,913,914,915,916,917,918,919,920,925,928,931,935,936,937,939,940,941,947,949,951,952,954,956,957,959,969,970,971,972,973,975,978,979,980,984,985,989};
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	areacode := USAreaCodes[r.Intn(len(USAreaCodes))]
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	prefix := r.Intn((999 - 200) + 1) + 200
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	suffix := r.Intn(10000)
	areacodeString := strconv.Itoa(areacode)
	prefixString := strconv.Itoa(prefix)
	suffixString := strconv.Itoa(suffix)

	if suffix > 99 && suffix < 1000 {
		suffixString = "0" + suffixString
	} else if suffix > 9 && suffix < 100 {
		suffixString = "00" + suffixString
	} else if(suffix < 10 && suffix > 0) {
		suffixString = "000" + suffixString
	}

	return areacodeString+prefixString+suffixString
}
