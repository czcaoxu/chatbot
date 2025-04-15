# chatbot

Chatbot is a system that enables human-computer interaction. It can be privately deployed and provides both RESTful API and web-based usage options.

Chatbot‘s key features are:

- Compatible with various models including ChatGPT, Qwen, DeepSeek, and Llama
- Maintains conversation context across interactions
- Enables integration with custom knowledge bases for enhanced responses

## start service
```
docker-compose up -d
```

## pull llama3 in Ollama
```azure
docker exec -it ollama ollama pull llama3
```

## RESTFul API

| Method | URL      | Description                                |
|--------|----------|--------------------------------------------|
| POST   | /chat    | Send chat request, support multiple models |
| GET    | /history | Retrieve available AI models.              |
| GET    | /models  | Retrieve chat history                      |
| DELETE | /session | Remove chat                                |