# Vigilant Waddle

Este projeto é uma API de autenticação e de gerenciamento de usuários.
São dois (2) tipos de usuários:

- Produtor
- Aluno

Nesta API será possível registrar cada usuário e autenticar.


# :pushpin: Tabela de conteúdo

- [Primeiros passos](#footprints-primeiros-passos)
- [Como rodar](#runner-como-rodar)
- [Documentação adicional](#book-documentação-adicional)
- [Dicas](#question-dicas)
- [Testes](#apple-testes)
- [Como contribuir](#barber-como-contribuir)
- [Issues](#bug-issues)

# :footprints: Primeiros passos
[Voltar ao topo](#vigilant-waddle)

Certifique-se de ter instalado os seguintes recursos:
- [Go lang](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Swag](#instalação-do-swag)

# :runner: Como rodar
[Voltar ao topo](#vigilant-waddle)

Após ter instalado os [recursos necessários](#footprints-primeiros-passos), você pode executar o seguinte comando:

```shell
make
```

Você terá o seguinte output:

```shell
Make usage:
        make run        - run the server
        make docs       - generate swagger documentation
        make rebuild_db - rebuild the database
        make stop       - stop the server
        make mock       - generate application mocks
        make test       - run tests
        make e2e-test   - run e2e tests
```

Estes acima, são os comandos que você pode rodar.

**Para rodar a aplicação:** `make run`

# :book: Documentação adicional
[Voltar ao topo](#vigilant-waddle)

Após conseguir rodar a aplicação com sucesso, você terá alguns recursos:
- [Swagger](http://localhost:8081/swagger/index.html)
- [Godoc](http://localhost:6061/pkg/github.com/jeanmolossi/vigilant-waddle/)
- [PgAdmin](http://localhost:1234/)

# :question: Dicas
[Voltar ao topo](#vigilant-waddle)

---

### Instalação Fácil de Go lang

Para facilitar basta fazer download do [gist](https://gist.github.com/jeanmolossi/8f2a643540aee671becf828d983952fd) e executar.
Confira a versão e OS antes de executar.

---

### Instalação do Swag

Certifique-se de ter instalado o [Go lang](https://go.dev/doc/install) antes de prosseguir. Após isso, você pode rodar:

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

[Documentação do Swag](https://github.com/swaggo/swag)

---

### Troubleshooting com o Banco de dados

Caso esteja com problemas ao fazer sua aplicação inicializar e se conectar com o banco de dados, siga os passos:

1. Execute:
	```shell
	make run
	```
2. Aguarde a aplicação iniciar.
	Para saber se a aplicação está funcionando:
	```shell
	curl 'http://localhost:8081/ping'
	```
	Se receber `{"message": "pong"}` a aplicação está rodando
3. Execute com sua senha sudo:
	```shell
	make rebuild_db
	```
	Isso irá reconstruir o container de banco de dados.
4. Execute:
	```shell
	docker logs vigilant_waddle_api_db -n 30 | grep "ready for start up."
	```
	Se este comando retornar algo, seu banco está funcionando.
5. Agora dê um _CTRL + S_ em algum arquivo `.go` para que a aplicação reinicie.

**Pronto!**

---

# :apple: Testes
[Voltar ao topo](#vigilant-waddle)

Esta aplicação possui testes unitários e testes de Integração.

Você pode executar os testes com os seguintes comandos:
- Unit: `make test`
- Integração: `make e2e-test`

# :barber: Como contribuir
[Voltar ao topo](#vigilant-waddle)

Para contribuir com este projeto crie uma branch com seu nome e sua contribuição. Exemplos:

- `johndoe/fix/acl-bug`
- `johndoe/feature/two-factor-login`

Certifique-se, também, de adicionar testes para sua contribuição e garanta que nenhuma outra implementação seja prejudicada por sua contribuição.

Abra um Pull request com sua modificação.

# :bug: Issues
[Voltar ao topo](#vigilant-waddle)

Encontrou algum bug ou problema de qualquer natureza. Sinta-se a vontade para comentar e abrir uma issue.
Entretanto, coloque o label apropriado na sua issue
