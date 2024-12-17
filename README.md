# Desafio técnico Go Expert FullCycle

## Configuração

Copiar o arquivo `.env.exemplo` com o nome `.env`

## Rodando localmente

- Faça o build da imagem `docker build . -t weather_app`
- Rode a imagem `docker run -p 8080 weather_app`

### Exemplos de requição

```bash
curl --location 'https://go-expert-weather-gcp-991159307010.us-central1.run.app/weather?CEP=81130070'
```
