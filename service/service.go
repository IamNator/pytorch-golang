package service

import (
	"encoding/json"
	"errors"

	"gopkg.in/square/go-jose.v2"
	"gorgonia.org/tensor"
)

type (
  AppService struct {
  }
)


func New() *AppService {
  
  
  return &AppService{}
}

func (a *AppService) Init() error {
  
  model := loadModelFromLibrary()
  
  ciphertext, err := encryptModel([]byte("encrption-key"), model)
  if err != nil {
    return err
  }

  key := findCatModelStoreKey
  
  if err :=  storeEncryptedModelToDB(key,ciphertext); err != nil {
    return err
  }

  return nil
}

func (a *AppService) Predict(input string) (string, error) {
  
  
  ciphertext, err := fetchEncryptedModelToDB("findCatModel")
  if err != nil {
    return "", err
  }
  
  model, err := decryptModel([]byte("encrption-key"), ciphertext)
  if err != nil {
    return "", err
  }
 
  return runPytorchModel(input, model)
}



// Load the PyTorch model
func loadModelFromLibrary() *tensor.Dense {
  return &tensor.Dense{}
}


type (
  ModelStoreKey string
)

const (
  findCatModelStoreKey ModelStoreKey = "findCatModel"
  findDogModelStoreKey ModelStoreKey = "findDogModel"
)

var (
  modelStore = make(map[ModelStoreKey]string)
)

// Store the PyTorch Model to DB
func storeEncryptedModelToDB(key ModelStoreKey, model string) error {
   modelStore[key] = model
   return nil
}


// Fetch the PyTorch Model from DB
func fetchEncryptedModelToDB(modelKey ModelStoreKey) (string, error) {
  if model, ok := modelStore[modelKey]; ok {
    return model, nil
  } else {
    return "", errors.New("model not found")
  }
}


//encrypt a PyTorch model
func encryptModel(key []byte, model *tensor.Dense) (string, error) {
    // Serialize the model
    modelJSON, err := json.Marshal(model)
    if err != nil {
        return "", err
    }

    // Encrypt the model
    encrypter, err := jose.NewEncrypter(jose.A256GCM, jose.Recipient{Algorithm: jose.DIRECT, Key: key}, nil)
    if err != nil {
        return "", err
    }

    obj, err := encrypter.Encrypt(modelJSON)
    if err != nil {
        return "", err
    }

    return obj.FullSerialize(), nil
}



//decrypt a PyTorch model
func decryptModel(key []byte, ciphertext string) (*tensor.Dense, error) {


   obj, err := jose.ParseEncrypted(ciphertext)
    if err!= nil {
        return nil, err
    }
  
    // Decrypte the model
    modelJSON, err := obj.Decrypt(ciphertext)
    if err != nil {
        return nil, err
    }

    var model tensor.Dense
  
    err = json.Unmarshal([]byte(modelJSON), &model)
    if err!= nil {
        return nil, err
    }

    return &model, nil
}
   


//running a PyTorch model
func runPytorchModel(input string, model *tensor.Dense) (string, error) {
    // Run the model on input data
  
    retValue, err := model.Apply(input)
    if err != nil {
      return "", err
    }

    return retValue.String(), nil
}



