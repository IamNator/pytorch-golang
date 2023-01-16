# pytorch-golang
loading and running a PyTorch model in Go using the gonum/gonum library, serving an API endpoint using go-chi/chi, and encrypting the model using go-jose/jose.


## About PyTorch
PyTorch is an open-source machine learning library for Python that provides a seamless integration of the research-to-production workflow. It is primarily used for natural language processing (NLP) and computer vision tasks.

A PyTorch model is a neural network architecture implemented in PyTorch that can be trained on a dataset to perform a specific task, such as image classification or text generation. The model is represented by a set of parameters (also known as weights) that are learned from the data during the training process. These parameters are stored in a file called a state_dict, which can be saved to disk and later loaded to perform inference on new data. The architecture of the model is defined using PyTorch's neural network library, torch.nn, which provides a set of pre-defined layers and functions to build the model.


## Why do I need to encrypt the model
Encrypting the model can prevent unauthorized access to the raw model data, this is important because the model may contain sensitive information such as proprietary algorithms, trade secrets, and valuable intellectual property. By encrypting the model, you can ensure that only authorized parties with the correct encryption key can access the model and use it for its intended purpose. This also will make it more difficult for others to reverse engineer the model and understand how it works. Additionally, it will prevent others from using the model for unintended purposes such as creating competing products or services.

## why Chi
Chi is a popular Go library for building HTTP servers and APIs. It is a lightweight and flexible router that helps developers to easily create HTTP routes and handlers. It is very fast and performant and has a simple and intuitive API. It also provides support for middlewares, which are functions that get executed before or after the request handler. This allows developers to add functionality such as logging, authentication, and rate limiting to the API without having to add it to every handler individually. Chi also has a good documentation, it is widely used and has a good community support which makes it easy to find solutions to any problems that may arise.

