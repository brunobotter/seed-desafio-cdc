package response

import "github.com/brunobotter/casa-codigo/internal/data/model"

type StateResponse struct {
	ID        uint   `json:"id"`
	StateName string `json:"state_name"`
	CountryId int64  `json:"countru_id"`
}

type StateListResponse []StateItemResponse

type StateItemResponse struct {
	ID        uint   `json:"id"`
	StateName string `json:"name"`
}

func FromStateModel(m model.StateModel) StateResponse {
	return StateResponse{
		ID:        m.ID,
		StateName: m.Statename,
		CountryId: m.CountryId,
	}
}

func FromListStateModel(models []model.StateModel) StateListResponse {
	var response StateListResponse

	for _, m := range models {
		response = append(response, StateItemResponse{
			ID:        m.ID,
			StateName: m.Statename,
		})
	}

	return response
}
