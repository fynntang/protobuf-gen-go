package usecase

type IAuthUsecase interface {
	CreateAuth(ctx context.Context, in *authorizationV1.CreateAuthRequest) error
	UpdateAuth(ctx context.Context, in *authorizationV1.UpdateAuthRequest) error
	DeleteAuth(ctx context.Context, in *authorizationV1.DeleteAuthRequest) error
	GetAuth(ctx context.Context, in *authorizationV1.GetAuthRequest) (*entities.Auth, error)
	ListAuth(ctx context.Context, in *authorizationV1.ListAuthRequest) ([]*entities.Auth, int64, error)
	Log(ctx context.Context) *zap.SugaredLogger
}

type AuthUseCase struct {
}

func NewAuthUseCase() IAuthUsecase {
	return &AuthUseCase{}
}

func (a *AuthUseCase) CreateAuth(ctx context.Context, in *authorizationV1.CreateAuthRequest) error {
	panic("todo")
}

func (a *AuthUseCase) UpdateAuth(ctx context.Context, in *authorizationV1.UpdateAuthRequest) error {
	panic("todo")
}

func (a *AuthUseCase) DeleteAuth(ctx context.Context, in *authorizationV1.DeleteAuthRequest) error {
	panic("todo")
}

func (a *AuthUseCase) GetAuth(ctx context.Context, in *authorizationV1.GetAuthRequest) (*entities.Auth, error) {
	panic("todo")
}

func (a *AuthUseCase) ListAuth(ctx context.Context, in *authorizationV1.ListAuthRequest) ([]*entities.Auth, int64, error) {
	panic("todo")
}

func (a AuthUseCase) Log(ctx context.Context) *zap.SugaredLogger {
	return global.Logger(ctx).Named("AuthRepo")
}
