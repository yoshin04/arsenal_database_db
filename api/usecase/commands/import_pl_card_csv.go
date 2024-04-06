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

type IImportPlCardCsvUsecase interface {
	Run(file io.Reader) (string, error)
}

type importPlCardCsvUsecase struct {
	plCardRepo      repository.IPLCardRepository
	linkAbilityRepo repository.ILinkAbilityRepository
}

func NewImportPlCardCsvUsecase(plCardRepo repository.IPLCardRepository, linkAbilityRepo repository.ILinkAbilityRepository) IImportPlCardCsvUsecase {
	return &importPlCardCsvUsecase{
		plCardRepo:      plCardRepo,
		linkAbilityRepo: linkAbilityRepo,
	}
}

func (uc *importPlCardCsvUsecase) Run(file io.Reader) (string, error) {
	log.Println("Running importPlCardCsvUsecase.Run")
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
			log.Printf("Error encountered: %v", err)
			return "", err
		}

		var firstLinkAbilityID, secondLinkAbilityID *string

		// If record[16] is not "-", find or create a new LinkAbility
		if record[16] != "-" {
			requiredCardCount, _ := strconv.ParseUint(record[18], 10, 8) // Default to 0 if parse fails
			firstLinkAbility, err := uc.linkAbilityRepo.FindOrCreate(record[16], domain.CsvTextConvertToLinkAbilityEffects(record[19]), uint8(requiredCardCount))
			if err != nil {
				log.Printf("Error firstLinkAbility parsing: %v", err)
				return "", err
			}
			firstLinkAbilityID = &firstLinkAbility.ID
		} else {
			firstLinkAbilityID = nil
		}

		// Same for second LinkAbility
		if record[20] != "-" {
			requiredCardCount, _ := strconv.ParseUint(record[22], 10, 8) // Default to 0 if parse fails
			secondLinkAbility, err := uc.linkAbilityRepo.FindOrCreate(record[20], domain.CsvTextConvertToLinkAbilityEffects(record[23]), uint8(requiredCardCount))
			if err != nil {
				log.Printf("Error secondLinkAbility parsing: %v", err)
				return "", err
			}
			secondLinkAbilityID = &secondLinkAbility.ID
		} else {
			// If no second LinkAbility, set the ID to nil
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
			log.Printf("Error LongRangeAttack parsing: %v", err)
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

		plCard := &models.PLCard{
			ID:               record[1] + record[2],
			IncludeCode:      record[1],
			No:               record[2],
			ImageURL:         "https://example.com/" + record[1] + record[2] + ".webp",
			Name:             record[4],
			Rarity:           record[3],
			Type:             record[6],
			Cost:             uint8(cost),
			Mobility:         uint16(mobility),
			LongRangeAttack:  uint16(longRangeAttack),
			CloseRangeAttack: uint16(closeRangeAttack),
			HP:               uint16(hp),
			TotalScore:       uint16(totalScore),
			PlSkillName:      record[13],
			PlSkillCondition: record[14],
			PlSkillDetail:    record[15],
			SeriesTitle:      record[24],

			FirstLinkAbilityID:  firstLinkAbilityID,
			SecondLinkAbilityID: secondLinkAbilityID,
		}

		log.Printf("Upserting pl card : %v", plCard)

		if err := uc.plCardRepo.Upsert(plCard); err != nil {
			log.Printf("Error upserting pl card: %v", err)
			return "", err
		}
	}

	return "Successfully imported pl cards", nil
}
