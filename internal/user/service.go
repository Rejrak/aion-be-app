package user

import (
	"be/internal/database/db"
	"be/internal/database/models"
	"be/internal/utils"
	"context"
	"errors"

	userService "be/gen/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	Repository *Repository
	// kcClient   *KeyCloak
}

func NewService() *Service {
	var (
		aionDB = db.DB.AionDB
	)
	return &Service{
		// kcClient:   NewKCervice(),
		Repository: NewUserRepository(aionDB),
	}
}

// Create crea un nuovo utente sia in Keycloak che nel database
func (s *Service) Create(ctx context.Context, payload *userService.CreateUserPayload) (*userService.User, error) {
	// Creazione in Keycloak
	userModel := models.User{
		KCID:      uuid.MustParse(payload.KcID),
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Nickname:  *payload.Nickname,
		Admin:     payload.Admin,
	}

	// Salvataggio nel database
	savedModel, err := s.Repository.SaveUser(ctx, userModel)
	if err != nil {
		return nil, err
	}

	return &userService.User{
		ID:        savedModel.KCID.String(),
		KcID:      savedModel.KCID.String(),
		FirstName: savedModel.FirstName,
		LastName:  savedModel.LastName,
		Nickname:  &savedModel.Nickname,
		Admin:     savedModel.Admin,
	}, nil
}

// Get restituisce un utente dal database
func (s *Service) Get(ctx context.Context, payload *userService.GetPayload) (*userService.User, error) {
	user, err := s.Repository.FindByID(ctx, payload.ID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err := &userService.NotFound{Message: "Utente non trovato"}
			utils.Log.Info(ctx, err)
			return nil, err
		}
		return nil, &userService.NotFound{Message: "Utente non trovato"}
	}

	return &userService.User{
		ID:        user.KCID.String(),
		KcID:      user.KCID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Nickname:  &user.Nickname,
		Admin:     user.Admin,
	}, nil
}

// List restituisce un elenco di utenti con paginazione
func (s *Service) List(ctx context.Context, payload *userService.ListPayload) ([]*userService.User, error) {
	users, err := s.Repository.List(ctx, payload.Limit, payload.Offset)
	if err != nil {
		return nil, err
	}

	var response []*userService.User
	for _, user := range users {
		response = append(response, &userService.User{
			ID:        user.KCID.String(),
			KcID:      user.KCID.String(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Nickname:  &user.Nickname,
			Admin:     user.Admin,
		})
	}
	return response, nil
}

// Update aggiorna un utente nel database e in Keycloak
func (s *Service) Update(ctx context.Context, payload *userService.UpdatePayload) (*userService.User, error) {
	// Recupera l'utente dal database
	user, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &userService.NotFound{Message: "Utente non trovato"}
		}
		return nil, err
	}

	// Aggiorna i dati nel database
	user.FirstName = payload.FirstName
	user.LastName = payload.LastName
	if payload.Nickname != nil {
		user.Nickname = *payload.Nickname
	}
	user.Admin = payload.Admin

	_, err = s.Repository.SaveUser(ctx, *user)
	if err != nil {
		return nil, err
	}

	// Aggiorna in Keycloak
	// err = s.KcUpdate(ctx, &payload.FirstName, &payload.LastName, payload.KcID)
	// if err != nil {
	// 	return nil, err
	// }

	return &userService.User{
		ID:        user.KCID.String(),
		KcID:      user.KCID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Nickname:  &user.Nickname,
		Admin:     user.Admin,
	}, nil
}

// Delete elimina un utente sia dal database che da Keycloak
func (s *Service) Delete(ctx context.Context, payload *userService.DeletePayload) error {
	// Recupera l'utente dal database
	user, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &userService.NotFound{Message: "Utente non trovato"}
		}
		return err
	}

	// Cancella da Keycloak
	// err = s.KcDelete(ctx, user.KCID.String())
	// if err != nil {
	// 	return err
	// }

	// Cancella dal database
	err = s.Repository.DeleteUser(ctx, user.KCID.String())
	if err != nil {
		return err
	}

	return nil
}
