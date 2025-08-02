package figma

type Handler struct {
	client  HandlerClient
	service HandlerService
}

type HandlerClient interface {
	GetFileInfo(fileID string) error
}

type HandlerService interface {
}

func NewHandler(service HandlerService) *Handler {
	return &Handler{}
}
