package usecase

import (
	"app/db/models"
	repository "app/repositories"
	"encoding/csv"
	"io"
	"log"
	"strconv"
)

type IImportTacticalCardCsvUsecase interface {
	Run(file io.Reader) (string, error)
}

type importTacticalCardCsvUsecase struct {
	tacticalCardRepo repository.ITacticalCardRepository
}

func NewImportTacticalCardCsvUsecase(tacticalCardRepo repository.ITacticalCardRepository) IImportTacticalCardCsvUsecase {
	return &importTacticalCardCsvUsecase{
		tacticalCardRepo: tacticalCardRepo,
	}
}

func (uc *importTacticalCardCsvUsecase) Run(file io.Reader) (string, error) {
	log.Println("Running importTacticalCardCsvUsecase.Run")
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

		cost, err := strconv.ParseUint(record[3], 10, 8)
		if err != nil {
			log.Printf("Error parsing cost: %v", err)
			return "", err
		}

		tacticalCard := &models.TacticalCard{
			ID:             record[0],
			ImageUrl:       "https://example.com/" + record[0],
			Name:           record[1],
			Cost:           uint8(cost),
			Detail:         record[4],
			UnlockCriteria: record[5],
		}

		log.Printf("Upserting tactical card: %v", tacticalCard)

		if err := uc.tacticalCardRepo.Upsert(tacticalCard); err != nil {
			log.Printf("Error upserting tactical card: %v", err)
			return "", err
		}
	}

	return "Successfully imported tactical cards", nil
}
