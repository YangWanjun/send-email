package oauth

import (
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/store"
	"send-email/model/repository"
)

func InitManager() oauth2.Manager {
	manager := manage.NewDefaultManager()

	// token store(ファイルにしないと、再起動したら無効になってしまう)
	manager.MustTokenStorage(store.NewFileTokenStore("token.db"))

	// client store
	clientStore := store.NewClientStore()
	applications, _ := repository.GetApplications()
	for _, app := range applications {
		clientStore.Set(app.ClientId, &models.Client{
			ID:     app.ClientId,
			Secret: app.ClientSecret,
			Domain: app.Domain,
		})
	}
	manager.MapClientStorage(clientStore)
	return manager
}