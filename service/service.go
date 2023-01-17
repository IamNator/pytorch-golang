package service

import (
	"encoding/json"

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
  
  if err :=  storeEncryptedModelToDB(ciphertext); err != nil {
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
 
  return runPytorchModel(input, model),  nil
}



// Load the PyTorch model
func loadModelFromLibrary() *tensor.Dense {
  return &tensor.Dense{}
}

// Store the PyTorch Model to DB
func storeEncryptedModelToDB(e string) error {
  return nil 
}


// Fetch the PyTorch Model from DB
func fetchEncryptedModelToDB(modelKey string) (string, error) {
  return "", nil
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
    ciphertext, err := encrypter.Encrypt(modelJSON)
    if err != nil {
        return "", err
    }

    return ciphertext, nil
}



//decrypt a PyTorch model
func decryptModel(key []byte, ciphertext string) (*tensor.Dense, error) {
   

    // Decrypte the model
    decrypter, err := jose.NewEncrypter(jose.A256GCM, jose.Recipient{Algorithm: jose.DIRECT, Key: key}, nil)
    if err != nil {
        return "", err
    }
  
    modelJSON, err := decrypter.Decrypt(ciphertext)
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
func runPytorchModel(input string, model *tensor.Dense) string {
    // Run the model on input data
    return model.Apply(tensor.NewDense(input))
}



