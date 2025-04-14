# Clean Architecture - Projeto do Curso GoExpert

Este projeto representa a solução para o Desafio de Clean Architecture do curso GoExpert, implementando a feature de Listagem de Pedidos no sistema de pedidos.

## Como Executar o Projeto

### Pré-requisitos
- Docker
- Docker Compose

### Passo a passo
1. Clone o repositório:
   ```bash
   git clone https://github.com/rafaelpapastamatiou/goexpert.git
   cd goexpert/20-clean-architecture
   ```

2. Inicie os containers com Docker Compose:
   ```bash
   docker-compose up -d
   ```

3. O sistema estará disponível nas seguintes portas:
   - API REST: http://localhost:8000
   - gRPC: localhost:50051
   - GraphQL: http://localhost:8080
   - RabbitMQ Management: http://localhost:15672 (usuário: guest, senha: guest)
   - MySQL: localhost:3306 (usuário: root, senha: root)

## Migrações de Banco de Dados
As migrações são executadas automaticamente ao iniciar o container através do script entrypoint.sh.