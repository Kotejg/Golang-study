package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// Insere usuario no banco de dados
func CriarUsusario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisicao"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Falha ao converter o ususario para struct"))
		return
	}

	db, err := banco.Conectar()

	if err != nil {
		w.Write([]byte("Erro ao converter no banco de dados"))
		return
	}
	defer db.Close()

	// prepare statemen
	statement, err := db.Prepare("INSERT INTO usuarios (nome, email) VALUES(?,?)")
	if err != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	insercao, err := statement.Exec(usuario.Nome, usuario.Email)

	if err != nil {
		w.Write([]byte("Erro ao executar statement"))
		return
	}

	idInserido, err := insercao.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}

	//status code
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso! ID: %d", idInserido)))

}

// busca todos os usurios cadastrados no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
	}
	defer db.Close()

	linhas, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		w.Write([]byte("Erro ao buscar os usuarios no banco "))
	}
	defer linhas.Close()
	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario
		if err := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("erro ao scanear o usuario"))
			return
		}
		usuarios = append(usuarios, usuario)
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("erro ao converter os usuarios para JSON"))
		return
	}

}

// busca usuario unico cadastrado no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.Atoi(params["id"])
	if err != nil {
		w.Write([]byte("erro ao converter o parametro para inteiro"))
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	linha, err := db.Query("SELECT * FROM usuarios WHERE id=?", ID)
	if err != nil {
		w.Write([]byte("erro ao buscar usuario no banco de dados"))
		return
	}
	defer linha.Close()
	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("erro ao escanear usuario no banco "))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuario); err != nil {
		w.Write([]byte("erro ao converter os usuario para JSON"))
		return
	}

}

// altera os dado de um usuario no banco
func AlterarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.Atoi(params["id"])
	if err != nil {
		w.Write([]byte("erro ao converter o parametro para inteiro"))
	}

	corpoReq, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("erro ao obter o corpo de dados"))
		return
	}

	var usuario usuario
	if err := json.Unmarshal(corpoReq, &usuario); err != nil {
		w.Write([]byte("erro ao converter os usuario para JSON"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update usuarios set nome=?, email=? where id=?")
	if err != nil {
		w.Write([]byte("erro ao criar o statement"))
	}
	defer statement.Close()
	if _, err = statement.Exec(usuario.Nome, usuario.Email, ID); err != nil {
		w.Write([]byte("erro ao atualizar o usuario"))
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

// delete o usuario no banco de dados
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.Atoi(params["id"])
	if err != nil {
		w.Write([]byte("erro ao converter o parametro para inteiro"))
	}
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("delete from usuarios where id=?")
	if err != nil {
		w.Write([]byte("erro ao criar o statement"))
		return
	}
	defer statement.Close()
	if _, err = statement.Exec(ID); err != nil {
		w.Write([]byte("erro ao deletar o usuario"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
	
}
