# GoGym

API para rastreamento de cargas de treino de musculação, permitindo ajustar progressivamente peso, séries e repetições dos exercícios ao longo do tempo.

## O que é

GoGym é uma API REST construída em Go que permite gerenciar treinos e exercícios. O foco principal é o controle de carga — você registra seus treinos, adiciona exercícios a eles e atualiza o peso, séries e repetições conforme evolui.

## Funcionalidades

- Cadastro de exercícios com grupo muscular
- Criação de treinos
- Adição de exercícios a um treino com carga definida (séries, repetições e peso)
- Atualização de carga dos exercícios

## Tecnologias

- **Go** — linguagem principal
- **Chi** — roteador HTTP
- **PostgreSQL** — banco de dados
- **sqlx** — mapeamento de queries SQL para structs

## Estrutura do projeto

```
gogym/
├── cmd/
│   └── api/
│       └── main.go         — ponto de entrada da aplicação
├── internal/
│   ├── domain/             — entidades e regras de negócio
│   ├── service/            — casos de uso
│   ├── handler/            — handlers HTTP
│   └── infra/
│       └── postgres/       — implementação dos repositórios
└── pkg/                    — utilitários reutilizáveis
```

## Como rodar

```bash
# Instalar dependências
go mod tidy

# Rodar a aplicação
go run ./cmd/api
```
