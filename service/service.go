package service


type (
  AppService struct {
  }
)


func New() *AppService {
  return &AppService{}
}

func (a *AppService) Predict(input string) (string, error) {
 return "", nil
}
