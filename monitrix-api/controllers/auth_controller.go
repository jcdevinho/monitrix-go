package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"monitrix-api/utils"
	"monitrix-api/models"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c echo.Context) error {
	var user models.User

	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "JSON inválido"})
	}

	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password) // Remove espaços extras

	// Log para debug do valor da senha recebido
	log.Printf("Senha recebida para validação: '%s', tamanho: %d", user.Password, len(user.Password))

	// Validações
	if !validarNome(user.Name) {
		return respostaComErro(c, "Nome inválido. Apenas letras, sem números, máximo 45 caracteres.")
	}

	if !validarEmail(user.Email) {
		return respostaComErro(c, "Email inválido ou domínio não permitido.")
	}

	if ok, msg := validarSenhaDetalhada(user.Password); !ok {
		log.Printf("Senha inválida: %s", msg) // Log do motivo do erro
		return respostaComErro(c, msg)
	}

	// Verifica duplicidade
	var exists bool
	err := utils.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", user.Email).Scan(&exists)
	if err != nil {
		return respostaComErro(c, "Erro ao verificar email.")
	}
	if exists {
		return respostaComErro(c, "Email já registrado.")
	}

	// Criptografar senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return respostaComErro(c, "Erro ao processar senha.")
	}

	// Inserir no banco
	res, err := utils.DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
		user.Name, user.Email, string(hashedPassword))
	if err != nil {
		return respostaComErro(c, "Erro ao registrar usuário.")
	}

	id, _ := res.LastInsertId()
	gravarLogRegistro(user.Name, user.Email)

	// Criar cookie HttpOnly
	cookie := new(http.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = gerarTokenFake(user.Email) // Aqui você pode substituir por um JWT se quiser
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteLaxMode
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Usuário registrado com sucesso!"
	})
}

// ===== Funções auxiliares =====

func validarNome(nome string) bool {
	if len(nome) == 0 || len(nome) > 45 {
		return false
	}
	match, _ := regexp.MatchString(`^[A-Za-zÀ-ÿ\s]+$`, nome)
	return match && !regexp.MustCompile(`\d`).MatchString(nome)
}

func validarEmail(email string) bool {
	regex := `^[^\s@]+@[^\s@]+\.[^\s@]+$`
	match, _ := regexp.MatchString(regex, email)
	if !match || strings.Contains(email, " ") {
		return false
	}

	dominio := strings.Split(email, "@")[1]
	aceitos := []string{"gmail.com", "hotmail.com", "outlook.com", "yahoo.com", "yandex.com", "zoho.com"}
	for _, d := range aceitos {
		if d == dominio {
			return true
		}
	}
	return strings.HasSuffix(dominio, ".com.br") && len(strings.Split(dominio, ".")[0]) >= 10
}

// validação da senha
func validarSenhaDetalhada(senha string) (bool, string) {
	var erros []string

	if len(senha) < 8 {
		erros = append(erros, "pelo menos 8 caracteres")
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(senha) {
		erros = append(erros, "1 letra maiúscula")
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(senha) {
		erros = append(erros, "1 letra minúscula")
	}
	if !regexp.MustCompile(`\d`).MatchString(senha) {
		erros = append(erros, "1 número")
	}
	if !regexp.MustCompile(`[\W_]`).MatchString(senha) {
		erros = append(erros, "1 caractere especial")
	}

	if len(erros) == 0 {
		return true, ""
	}

	// Formata mensagem para ficar mais amigável
	msg := "A senha deve conter "
	if len(erros) == 1 {
		msg += erros[0] + "."
	} else if len(erros) == 2 {
		msg += erros[0] + " e " + erros[1] + "."
	} else {
		// Se mais de 2 erros, separa com vírgula e 'e' antes do último
		msg += strings.Join(erros[:len(erros)-1], ", ") + " e " + erros[len(erros)-1] + "."
	}

	return false, msg
}

func respostaComErro(c echo.Context, msg string) error {
	return c.JSON(http.StatusBadRequest, map[string]string{"message": msg})
}

func gravarLogRegistro(nome, email string) {
	logDir := "logs"
	os.MkdirAll(logDir, os.ModePerm)

	f, err := os.OpenFile(logDir+"/registro.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Erro ao gravar log:", err)
		return
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	logger.Printf("Novo registro: Nome=%s, Email=%s\n", nome, email)
}

func gerarTokenFake(email string) string {
	return "sessao_" + strings.ReplaceAll(email, "@", "_") // substitua por JWT real se desejar
}
