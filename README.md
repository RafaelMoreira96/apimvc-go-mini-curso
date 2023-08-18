# Mini-Curso: Desenvolvimento de API em Go com Gin e GORM

## Duração: 2 horas

**Objetivo:**
Capacitar os participantes a desenvolverem APIs utilizando a linguagem de programação Go, o framework Gin gonic e a biblioteca GORM para operações de banco de dados, com enfoque na implementação prática.

## Conteúdo do Curso

### Sessão Única - Desenvolvimento Prático de uma API em Go com Gin e GORM
1. O que é uma API e sua importância no desenvolvimento de software.
2. Visão geral sobre CRUD (Create, Read, Update, Delete) e sua aplicação em APIs.
3. Introdução à linguagem Go: principais características e benefícios.
4. Conceito de ORM (Object-Relational Mapping) e sua utilidade no acesso a bancos de dados.
5. Arquitetura Modelo-View-Controller (MVC) e como ela se aplica no desenvolvimento de APIs.
6. Configuração do ambiente de desenvolvimento: instalação do Go, Gin e GORM.
7. Implementação de rotas utilizando o framework Gin para diferentes requisições HTTP.
8. Utilização do GORM para interagir com o banco de dados: definição de modelos e migrações.
9. Criação das operações CRUD (Create, Read, Update, Delete) utilizando o GORM.
10. Organização do código seguindo o padrão MVC: separação de responsabilidades.
11. Implementação do projeto prático: desenvolvimento de uma API completa usando Go, Gin e GORM.
12. Adição de validações simples de campos nas requisições.

## Instruções de Instalação do Go

Aqui está um guia passo a passo para instalar o Go na sua máquina:

1. **Baixar a Versão 1.21 do Go usando `wget`**: Use o seguinte comando para baixar a versão 1.21 do Go:
    ```bash
    wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
    ```

2. **Remover a Versão Anterior do Go**: Antes de prosseguir com a instalação da nova versão, remova qualquer versão anterior do Go, independentemente da versão. Execute o seguinte comando:
    ```bash
    sudo rm -rf /usr/local/go
    ```

3. **Instalar a Nova Versão do Go**: Instale a nova versão do Go usando o seguinte comando (em modo root):
    ```bash
    sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
    ```

4. **Adicionar o Go ao seu PATH**: Adicione o diretório `bin` do Go ao seu PATH executando o seguinte comando fora do root:
    ```bash
    export PATH=$PATH:/usr/local/go/bin
    ```

5. **Verificar a Versão Instalada**: Verifique se a nova versão do Go foi instalada corretamente executando o seguinte comando:
    ```bash
    go version
    ```

6. **Criar uma Pasta para Iniciar um Projeto Go**: Agora que você tem o Go instalado, crie uma pasta para iniciar seus projetos Go. Vá para o diretório de sua escolha e execute os seguintes comandos:
    ```bash
    mkdir meu-projeto-go
    cd meu-projeto-go
    go mod init nome-do-seu-projeto
    ```

Agora você está pronto para começar a desenvolver em Go!

Lembre-se de que os comandos mencionados acima podem exigir permissões de administrador (root), dependendo das configurações do seu sistema. Certifique-se de entender o que cada comando faz antes de executá-lo.
