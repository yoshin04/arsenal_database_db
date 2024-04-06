package usecase

import (
	"app/db/models"
	domain "app/domain/card"
	repository "app/repositories"
	"encoding/csv"
	"io"
	"log"
	"strconv"
)

type IImportMsCardCsvUsecase interface {
	Run(file io.Reader) (string, error)
}

type importMsCardCsvUsecase struct {
	msCardRepo      repository.IMSCardRepository
	linkAbilityRepo repository.ILinkAbilityRepository
}

func NewImportMsCardCsvUsecase(msCardRepo repository.IMSCardRepository, linkAbilityRepo repository.ILinkAbilityRepository) IImportMsCardCsvUsecase {
	return &importMsCardCsvUsecase{
		msCardRepo:      msCardRepo,
		linkAbilityRepo: linkAbilityRepo,
	}
}

func (uc *importMsCardCsvUsecase) Run(file io.Reader) (string, error) {
	log.Println("Running importMsCardCsvUsecase.Run")
	csvReader := csv.NewReader(file)

	// ヘッダ行を読み飛ばす
	if _, err := csvReader.Read(); err != nil {
		log.Printf("Error reading header row: %v", err)
		return "", err
	}

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Print("Error encountered: %v", err)
			return "", err
		}

		var firstLinkAbilityID, secondLinkAbilityID *string

		if record[30] != "-" {
			requiredCardCount, _ := strconv.ParseUint(record[32], 10, 8)
			firstLinkAbility, err := uc.linkAbilityRepo.FindOrCreate(record[30], domain.CsvTextConvertToLinkAbilityEffects(record[33]), uint8(requiredCardCount))
			if err != nil {
				log.Printf("Error firstLinkAbility parsing: %v", err)
				return "", err
			}
			firstLinkAbilityID = &firstLinkAbility.ID
		} else {
			firstLinkAbilityID = nil
		}

		if record[34] != "-" {
			requiredCardCount, _ := strconv.ParseUint(record[36], 10, 8)
			secondLinkAbility, err := uc.linkAbilityRepo.FindOrCreate(record[34], domain.CsvTextConvertToLinkAbilityEffects(record[37]), uint8(requiredCardCount))
			if err != nil {
				log.Printf("Error secondLinkAbility parsing: %v", err)
				return "", err
			}
			secondLinkAbilityID = &secondLinkAbility.ID
		} else {
			secondLinkAbilityID = nil
		}

		cost, err := strconv.ParseUint(record[7], 10, 8)
		if err != nil {
			log.Printf("Error cost parsing: %v", err)
			return "", err
		}

		mobility, err := strconv.ParseUint(record[8], 10, 16)
		if err != nil {
			log.Printf("Error mobility parsing: %v", err)
			return "", err
		}

		longRangeAttack, err := strconv.ParseUint(record[9], 10, 16)
		if err != nil {
			log.Printf("Error longRangeAttack parsing: %v", err)
			return "", err
		}

		closeRangeAttack, err := strconv.ParseUint(record[10], 10, 16)
		if err != nil {
			log.Printf("Error closeRangeAttack parsing: %v", err)
			return "", err
		}

		hp, err := strconv.ParseUint(record[11], 10, 16)
		if err != nil {
			log.Printf("Error hp parsing: %v", err)
			return "", err
		}

		totalScore, err := strconv.ParseUint(record[12], 10, 16)
		if err != nil {
			log.Printf("Error totalScore parsing: %v", err)
			return "", err
		}

		mainRange, err := strconv.ParseUint(record[13], 10, 16)
		if err != nil {
			log.Printf("Error mainRange parsing: %v", err)
			return "", err
		}

		subRange, err := strconv.ParseUint(record[14], 10, 16)
		if err != nil {
			log.Printf("Error subRange parsing: %v", err)
			subRange = 0
		}

		spCost, err := strconv.ParseUint(record[20], 10, 8)
		if err != nil {
			log.Printf("Error spCost parsing: %v", err)
			return "", err
		}

		spPower, err := strconv.ParseUint(record[21], 10, 16)
		if err != nil {
			log.Printf("Error spPower parsing: %v", err)
			return "", err
		}

		spRange, err := strconv.ParseUint(record[22], 10, 8)
		if err != nil {
			log.Printf("Error spRange parsing: %v", err)
			spRange = 0
		}

		// AbilityNameの処理
		var abilityName *string
		if record[26] != "" && record[26] != "-" {
			abilityName = &record[26]
		}

		// AbilityCostの処理
		var abilityCost *uint8
		if costStr := record[27]; costStr != "" && costStr != "-" {
			if cost, err := strconv.ParseUint(costStr, 10, 8); err == nil {
				costUint8 := uint8(cost)
				abilityCost = &costUint8
			}
		}

		// AbilityRangeの処理
		var abilityRange *uint8
		if rangeStr := record[28]; rangeStr != "" && rangeStr != "-" {
			if rangeVal, err := strconv.ParseUint(rangeStr, 10, 8); err == nil {
				rangeUint8 := uint8(rangeVal)
				abilityRange = &rangeUint8
			}
		}

		// AbilityDetailの処理
		var abilityDetail *string
		if record[29] != "" && record[29] != "-" {
			abilityDetail = &record[29]
		}

		msCard := &models.MSCard{
			ID:                    record[1] + record[2],
			IncludedCode:          record[1],
			No:                    record[2],
			ImageURL:              "https://example.com/" + record[1] + record[2] + ".webp",
			Name:                  record[4],
			Rarity:                record[3],
			Type:                  record[6],
			Cost:                  uint8(cost),
			Mobility:              uint16(mobility),
			LongRangeAttack:       uint16(longRangeAttack),
			CloseRangeAttack:      uint16(closeRangeAttack),
			HP:                    uint16(hp),
			TotalScore:            uint16(totalScore),
			MainRange:             uint16(mainRange),
			SubRange:              uint16(subRange),
			GroundSuitability:     record[15],
			SpaceSuitability:      record[16],
			DesertSuitability:     record[17],
			UnderwaterSuitability: record[18],
			SpAttackName:          record[19],
			SpCost:                uint8(spCost),
			SpPower:               uint16(spPower),
			SpRange:               uint8(spRange),
			SpTargetType:          domain.CsvTextConvertToSpTargetType(record[23]),
			SpType:                record[24],
			SpDetail:              record[25],
			AbilityName:           abilityName,
			AbilityCost:           abilityCost,
			AbilityRange:          abilityRange,
			AbilityDetail:         abilityDetail,
			SeriesTitle:           record[38],
			FirstLinkAbilityID:    firstLinkAbilityID,
			SecondLinkAbilityID:   secondLinkAbilityID,
		}

		log.Printf("Upserting: %v", msCard)

		if err := uc.msCardRepo.Upsert(msCard); err != nil {
			log.Printf("Error upserting ms card: %v", err)
			return "", err
		}
	}

	return "successfully imported tactical cards", nil
}
