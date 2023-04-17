package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"strings"

	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
)

//CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request){
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
		
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = usuario.Preparar(); erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	respostas.JSON(w, http.StatusCreated, usuario)

	

}
//BuscarUsuarios busca todos os usuários salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request){
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.Buscar(nomeOuNick)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, usuarios)
}
//BuscarUsuario busca um usuario salvo no banco
func BuscarUsuario(w http.ResponseWriter, R *http.Request){
	w.Write([]byte("Buscando um usuário!"))
}
//AtualiarUsuario altera as informações de um usuário no banco
func AtualizarUsuario(w http.ResponseWriter, R *http.Request){
	w.Write([]byte("Atualizando suário!"))
}
//DeletarUsuario exclui as informações de um usuário no banco
func DeletarUsuario(w http.ResponseWriter, R *http.Request){
	w.Write([]byte("Deletando um usuário!"))
}
