# pytorch-golang
loading and running a PyTorch model in Go using the gonum/gonum library, serving an API endpoint using go-chi/chi, and encrypting the model using go-jose/jose.


## About PyTorch
PyTorch is an open-source machine learning library for Python that provides a seamless integration of the research-to-production workflow. It is primarily used for natural language processing (NLP) and computer vision tasks.

A PyTorch model is a neural network architecture implemented in PyTorch that can be trained on a dataset to perform a specific task, such as image classification or text generation. The model is represented by a set of parameters (also known as weights) that are learned from the data during the training process. These parameters are stored in a file called a state_dict, which can be saved to disk and later loaded to perform inference on new data. The architecture of the model is defined using PyTorch's neural network library, torch.nn, which provides a set of pre-defined layers and functions to build the model.

## More Information
Sure! PyTorch is a library for creating machine learning models, which are programs that can learn from data. Think of it like a set of building blocks that you can use to create different types of models, like a lego set.

To create a model with PyTorch, you first need to define its architecture, which is like the blueprint of a building. This is done by specifying the types and number of layers that the model should have. Layers are like different rooms in a building, each with a specific function. For example, one layer might be used for detecting edges in an image, while another layer might be used for recognizing shapes.

Once the architecture is defined, you need to train the model using a dataset. Training is like giving the model a set of instructions, so it can learn from the data. This is done by showing the model a lot of examples, and adjusting the parameters (like the weights of the layers) until the model can accurately predict the output.

After the model is trained, you can use it to make predictions on new data. This is like asking the model to perform a specific task, like recognizing a face in an image, or translating a sentence.

So in short, PyTorch is a library that helps you create machine learning models, train them on data, and use them to make predictions.


## Why do I need to encrypt the model?
Encrypting the model can prevent unauthorized access to the raw model data, this is important because the model may contain sensitive information such as proprietary algorithms, trade secrets, and valuable intellectual property. By encrypting the model, you can ensure that only authorized parties with the correct encryption key can access the model and use it for its intended purpose. This also will make it more difficult for others to reverse engineer the model and understand how it works. Additionally, it will prevent others from using the model for unintended purposes such as creating competing products or services.

## why Chi?
Chi is a popular Go library for building HTTP servers and APIs. It is a lightweight and flexible router that helps developers to easily create HTTP routes and handlers. It is very fast and performant and has a simple and intuitive API. It also provides support for middlewares, which are functions that get executed before or after the request handler. This allows developers to add functionality such as logging, authentication, and rate limiting to the API without having to add it to every handler individually. Chi also has a good documentation, it is widely used and has a good community support which makes it easy to find solutions to any problems that may arise.

## How does the project work?
The whole project works together by combining the capabilities of different libraries and tools. Here's a high-level overview of how the project would work:

1. The PyTorch model is fine-tuned and ready to be used.
2. The model is encrypted using the go-jose/jose library to prevent unauthorized access to the raw model data.
3. The encrypted model is then stored in a MongoDB database.
4. When the model is needed, it is retrieved from the MongoDB database, and then decrypted using the key that was stored in the database.
5. The decrypted model is loaded into memory using the gonum/gonum library, so it can be used to make predictions on the input data from the API endpoint.
6. The Go backend serves an API endpoint using the go-chi/chi library.
7. The API endpoint listens for incoming requests and checks for an active API key.
8. Once a request is received, the API endpoint parses the input as a JSON, it uses the loaded model to make predictions and returns the output in JSON format.   
    
By storing the encrypted models in a MongoDB database, it allows for better management, organization and accessibility of the models, as well as making it easier to rollback to previous models or versions.

It also allows for easier scaling, if multiple instances of the Go app need to access the same models,

The Go app serves as a bridge between the client and the PyTorch model, handling the encryption and loading of the model as well as serving the API endpoint, handling the incoming requests and returning the output.



