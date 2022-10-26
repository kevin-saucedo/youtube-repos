package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/xuri/excelize/v2"
)

var conn *sql.DB

type Persona struct {
	ID        int
	Nombre    string
	Apellidos string
	Direccion string
	Telefono  string
}

func init() {
	var err error
	const dsn = "user=postgres dbname=prueba_db password=maira002 sslmode=disable"
	conn, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln("DBConnection Error:", err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatalln("Ping Error:", err)
	}
	log.Println("DB Connexion Opened!!!!")
}
func main() {
	http.HandleFunc("/hello", handlerHello)
	http.HandleFunc("/file", filehandle)
	log.Println("Listen localhost:8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalln(err)
	}
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello !!!!!")
}

func filehandle(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("myfile")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No hay archivo")
		return
	}
	defer file.Close()
	f, err := excelize.OpenReader(file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se pudo abrir el archivo")
		return
	}
	rows, err := f.GetRows("Contactos")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se pudo leer la hoja Contactos del archivo .xlsx")
		return
	}
	var personas []Persona
	for index, row := range rows {
		if index == 0 {
			continue
		}
		id, err := strconv.Atoi(row[0])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "ID Error Row=%d Value=%s", index+1, row[0])
			return
		}
		personas = append(personas, Persona{
			ID:        id,
			Nombre:    row[1],
			Apellidos: row[2],
			Direccion: row[3],
			Telefono:  row[4],
		})
	}
	if err := InsertPersonas(r.Context(), personas); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Datos insertados!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
}

func InsertPersonas(ctx context.Context, personas []Persona) error {
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	const query = "insert into persona( id,nombre,apellido,direccion,telefono) values($1,$2,$3,$4,$5)"
	for _, p := range personas {
		stm, err := tx.Prepare(query)
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = stm.Exec(p.ID, p.Nombre, p.Apellidos, p.Direccion, p.Telefono)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
