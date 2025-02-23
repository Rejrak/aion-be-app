package user

import (
	"context"
	"errors"
	"os"

	"github.com/Nerzal/gocloak/v13"
)

type KeyCloak struct {
	client       *gocloak.GoCloak
	clientID     string
	clientSecret string
	realm        string
}

func NewKCervice() *KeyCloak {
	var (
		client   = gocloak.NewClient(os.Getenv("KC_URL"))
		kcClient = os.Getenv("KC_CLIENT_ID")
		kcSecret = os.Getenv("KC_CLIENT_SECRET")
		kcRealm  = os.Getenv("KC_REALM")
	)

	return &KeyCloak{
		client:       client,
		clientID:     kcClient,
		clientSecret: kcSecret,
		realm:        kcRealm,
	}
}

func (s *KeyCloak) GetToken(ctx context.Context) (*gocloak.JWT, error) {
	token, err := s.client.LoginClient(ctx, s.clientID, s.clientSecret, s.realm)
	if err != nil {
		rsp := errors.New("errore di comunicazione [KC-GT]")
		return nil, rsp
	}
	return token, nil
}

func (s *KeyCloak) KcCreate(ctx context.Context, username, firstName, lastName, password string) (uuid *string, err error) {
	token, err := s.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	kcUser := gocloak.User{
		Username:      gocloak.StringP(username),
		Enabled:       gocloak.BoolP(true),
		EmailVerified: gocloak.BoolP(true),
		FirstName:     (*string)(&firstName),
		LastName:      (*string)(&lastName),
	}

	userID, err := s.client.CreateUser(ctx, token.AccessToken, s.realm, kcUser)
	if err != nil {
		return nil, errors.New("user already exists [KC-CU]")
	}

	if err := s.client.SetPassword(ctx, token.AccessToken, userID, s.realm, password, false); err != nil {
		return nil, errors.New("comunication error [KC-SP]")
	}

	uuid = &userID

	return
}

func (s *KeyCloak) KcUpdate(ctx context.Context, firstName, lastName *string, uuid string) (err error) {
	token, err := s.GetToken(ctx)
	if err != nil {
		return err
	}

	kcUser := gocloak.User{ID: &uuid}
	if firstName != nil && *firstName != "" {
		kcUser.FirstName = firstName
	}
	if lastName != nil && *lastName != "" {
		kcUser.LastName = lastName
	}

	if err := s.client.UpdateUser(ctx, token.AccessToken, s.realm, kcUser); err != nil {
		return errors.New("errore di comunicazione [KC-UU]")
	}

	return nil
}

func (s *KeyCloak) KcDelete(ctx context.Context, uuid string) error {
	token, err := s.GetToken(ctx)
	if err != nil {
		return err
	}

	if err := s.client.DeleteUser(ctx, token.AccessToken, s.realm, uuid); err != nil {
		return errors.New("errore di comunicazione [KC-DU]")
	}
	return nil
}
