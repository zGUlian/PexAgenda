📅 Sistema de Agenda com Interface Gráfica - GoLang + Fyne
Este é um sistema de agenda desenvolvido em GoLang com interface gráfica utilizando o framework Fyne.
Permite agendar compromissos, como reuniões e visitas a clientes, com funcionalidades de:

✅ Cadastro
✅ Edição
✅ Exclusão
✅ Filtro por data

🚀 Funcionalidades
✅ Adicionar compromisso

✅ Editar compromisso por ID

✅ Excluir compromisso por ID

✅ Filtrar compromissos por data

✅ Resetar filtro

✅ Armazenamento em arquivo JSON

🛠️ Tecnologias utilizadas
GoLang (versão 1.18+ recomendada)

Fyne (framework para GUI)

📦 Instalação e execução
1. Clone o repositório
bash
Copiar
Editar
git clone [https://github.com/seuusuario/agenda-golang.git](https://github.com/zGUlian/PexAgenda.git)
cd agenda-golang
2. Inicialize o módulo Go
bash
Copiar
Editar
go mod init agenda
3. Adicione a dependência do Fyne
bash
Copiar
Editar
go get fyne.io/fyne/v2
4. Ajuste dependências
bash
Copiar
Editar
go mod tidy
🖥️ Executando o projeto
bash
Copiar
Editar
go run main.go
📄 Estrutura do projeto
lua
Copiar
Editar
/agenda
|-- main.go
|-- models/
|    |-- compromisso.go
|-- storage/
|    |-- storage.go
|-- ui/
|    |-- ui.go
|-- go.mod
|-- go.sum
✍️ Como usar
Preencha os campos:

Título

Descrição

Data (formato: YYYY-MM-DD)

Hora (formato: HH:MM)

Tipo: Reunião ou Visita

Clique em:

Adicionar → para adicionar novo compromisso.

Editar por ID → para editar um compromisso.

Excluir por ID → para deletar um compromisso.

Filtrar por Data → para visualizar compromissos de uma data específica.

Resetar Filtro → para visualizar todos novamente.

💾 Armazenamento
Todos os compromissos são armazenados automaticamente no arquivo compromissos.json.

📚 Requisitos
Go 1.18 ou superior

Sistema operacional com suporte a Fyne (Linux, Windows ou macOS)

🤝 Contribuições
Sinta-se à vontade para abrir issues ou enviar pull requests!

📝 Licença
Este projeto está sob a licença MIT.

🧑‍💻 Autor
Desenvolvido por Gabriel Ulian.
