# Go-Faker

Go-Faker é uma biblioteca para Go que gera dados falsos para testes e desenvolvimento, semelhante ao Faker Gem no Ruby on Rails.

## Sumário

- [Instalação](#instalação)
- [Uso](#uso)
- [Funções Disponíveis](#funções-disponíveis)
- [Exemplos](#exemplos)
- [Contribuição](#contribuição)
- [Licença](#licença)

## Instalação

Para instalar a biblioteca, execute o seguinte comando:

```sh
  go get github.com/br4tech/go-faker
```

## Uso

Para usar a biblioteca, importe o pacote faker e chame as funções para gerar dados falsos.

```bash
  package main

  import (
      "fmt"
      "github.com/seuusuario/go-faker/faker"
  )

  func main() {
      fmt.Println("Random Name:", faker.Name())
      fmt.Println("Random Address:", faker.Address())
  }
```

## Funções Disponíveis

 - Name() string: Gera um nome completo aleatório.
-  Address() string: Gera um endereço aleatório.


## Contribuição

Se você quiser contribuir com a biblioteca, siga os passos abaixo:

Faça um fork do repositório.
Crie uma nova branch (git checkout -b feature/nova-funcionalidade).
Faça as mudanças desejadas e adicione commits (git commit -am 'Adiciona nova funcionalidade').
Faça push para a branch (git push origin feature/nova-funcionalidade).
Abra um Pull Request.