# Aplicação web em Golang - API VIACEP

Aplicação web em Go que consulta endereços através da API ViaCEP.

## ✨ Funcionalidades

- 🔍 Consulta de endereço por CEP
- ✅ Validação de CEP (apenas números, 8 dígitos)
- 📱 Layout responsivo
- 🎨 Interface limpa e intuitiva
- ⚡ Resposta rápida da API ViaCEP
- 🛡️ Tratamento de erros robusto

## 🚀 Como executar

### Pré-requisitos
- Go 1.21 ou superior instalado

### Passos para execução

1. **Clone ou baixe os arquivos**
   ```bash
   # Crie uma pasta para o projeto
   mkdir consulta-cep-go
   cd consulta-cep-go
   ```

2. **Salve os arquivos**
   - `apiviacep.go` - código principal da aplicação
   - `go.mod` - arquivo de módulo Go

3. **Execute a aplicação**
   ```bash
   go run apiviacep.go
   ```

4. **Acesse no navegador**
   ```
   http://localhost:8080
   ```

## 📋 Como usar

1. Digite um CEP válido (8 dígitos, apenas números)
2. Clique em "Consultar CEP" 
3. Os campos serão preenchidos automaticamente com os dados do endereço

### Exemplos de CEP para teste:
- `01001000` - Praça da Sé, São Paulo - SP
- `20040020` - Centro, Rio de Janeiro - RJ
- `30112000` - Centro, Belo Horizonte - MG
- `80010000` - Centro, Curitiba - PR

## 🔧 Principais melhorias em relação à versão original

### Validação e Segurança
- ✅ Validação server-side do CEP
- ✅ Sanitização de dados de entrada
- ✅ Tratamento robusto de erros da API
- ✅ Prevenção de XSS com templates seguros

### Interface do Usuário
- ✅ Mensagens de erro mais elegantes (sem palavrões)
- ✅ Feedback visual melhorado
- ✅ Loading states implícitos
- ✅ Campos readonly para dados consultados
- ✅ Responsividade aprimorada

### Arquitetura
- ✅ Código organizado em funções específicas
- ✅ Estruturas de dados tipadas
- ✅ Separação clara de responsabilidades
- ✅ Template HTML integrado (deploy simplificado)

## 🏗️ Estrutura do código

```
apiviacep.go
├── Estruturas de dados (ViaCEPResponse, FormData)
├── Validação de CEP (validaCEP)
├── Consulta à API (consultaCEP)
├── Handler HTTP principal (handler)
├── Template HTML integrado
└── Servidor HTTP (main)
```

## 🌐 API utilizada

- **ViaCEP**: https://viacep.com.br/
- Endpoint: `https://viacep.com.br/ws/{cep}/json/`
- Documentação: https://viacep.com.br/

## 📦 Deploy

Para fazer deploy da aplicação:

1. **Build da aplicação**
   ```bash
   go build -o consulta-cep apiviacep.go
   ```

2. **Executar o binário**
   ```bash
   ./consulta-cep
   ```

A aplicação é um binário único, sem dependências externas, facilitando o deploy em qualquer ambiente que suporte Go.
A aplicação é autocontida - não precisa de arquivos CSS/JS separados e gera um binário único para deploy.
Mantive a essência da minha aplicação Spring Boot original, mas com código mais robusto e profissional em Go!

## 🤝 Contribuições

Sinta-se à vontade para contribuir com melhorias:
- Adicionar cache de consultas
- Implementar histórico de pesquisas
- Adicionar mais validações
- Melhorar a interface

## 📄 Licença

Este projeto está sob licença MIT - sinta-se livre para usar e modificar.
