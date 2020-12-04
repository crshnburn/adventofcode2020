package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadPassports(t *testing.T) {
	lines := []string{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in"}

	passportRecords := ReadPassports(lines)

	require.Equal(t, 4, len(passportRecords))
	require.Equal(t, []string{"ecl", "pid", "eyr", "hcl", "byr", "iyr", "cid", "hgt"}, passportRecords[0])
	require.Equal(t, []string{"iyr", "ecl", "cid", "eyr", "pid", "hcl", "byr"}, passportRecords[1])
	require.Equal(t, []string{"hcl", "iyr", "eyr", "ecl", "pid", "byr", "hgt"}, passportRecords[2])
	require.Equal(t, []string{"hcl", "eyr", "pid", "iyr", "ecl", "hgt"}, passportRecords[3])
}

func TestValidPassport(t *testing.T) {
	lines := []string{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in"}

	passportRecords := ReadPassports(lines)

	require.True(t, ValidatePassport(passportRecords[0]))
	require.False(t, ValidatePassport(passportRecords[1]))
	require.True(t, ValidatePassport(passportRecords[2]))
	require.False(t, ValidatePassport(passportRecords[3]))
}

func TestValidatePassportFields(t *testing.T) {
	var invalidPassport = map[string]string{
		"eyr": "1972", "cid": "100",
		"hcl": "#18171d", "ecl": "amb", "hgt": "170", "pid": "186cm", "iyr": "2018", "byr": "1926",
	}

	var validPassport = map[string]string{
		"pid": "087499704", "hgt": "74in", "ecl": "grn", "iyr": "2012", "eyr": "2030", "byr": "1980",
		"hcl": "#623a2f",
	}

	require.True(t, ValidatePassportFields(validPassport))
	require.False(t, ValidatePassportFields(invalidPassport))
	require.False(t, ValidatePassportFields(map[string]string{
		"iyr": "2019",
		"hcl": "#602927", "eyr": "1967", "hgt": "170cm",
		"ecl": "grn", "pid": "012533040", "byr": "1946",
	}))
	require.False(t, ValidatePassportFields(map[string]string{
		"hcl": "dab227", "iyr": "2012",
		"ecl": "brn", "hgt": "182cm", "pid": "021572410", "eyr": "2020", "byr": "1992", "cid": "277",
	}))
	require.False(t, ValidatePassportFields(map[string]string{
		"hgt": "59cm", "ecl": "zzz",
		"eyr": "2038", "hcl": "74454a", "iyr": "2023",
		"pid": "3556412378", "byr": "2007",
	}))
}
