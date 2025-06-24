package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// Estrutura para resposta da API ViaCEP
type ViaCEPResponse struct {
	CEP        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
	Erro       bool   `json:"erro,omitempty"`
}

// Estrutura para dados do formulário
type FormData struct {
	CEP     string
	Rua     string
	Bairro  string
	Cidade  string
	Estado  string
	Erro    string
	Sucesso bool
}

// Template HTML integrado
const htmlTemplate = `
<!DOCTYPE html>
<html lang="pt-br">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Consulta CEP - Golang</title>
    <style>
        body {
            background: #b5a7a7;
            font-size: 20px;
            font-family: Arial, Helvetica, sans-serif;
            margin: 0;
            padding: 20px;
        }

        .container {
            margin: 0 auto;
            width: 100%;
            max-width: 440px;
            background: #FFF;
            border-radius: 8px;
            height: auto;
            padding: 20px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        }

        #idh2cep {
            text-align: center;
            color: #333;
            margin-bottom: 20px;
            transition: background-color 0.3s;
            padding: 10px;
            border-radius: 4px;
        }

        .frm-row {
            margin-bottom: 12px;
        }

        .frm-row label {
            display: block;
            margin-bottom: 3px;
            font-weight: bold;
            color: #555;
        }

        .frm-row input {
            width: calc(100% - 16px);
            padding: 8px;
            border-radius: 4px;
            border: 1px solid #CCC;
            height: 25px;
            font-size: 16px;
            transition: border-color 0.3s;
        }

        .frm-row input:focus {
            outline: none;
            border-color: #007bff;
            box-shadow: 0 0 5px rgba(0, 123, 255, 0.3);
        }

        .btn {
            background-color: #007bff;
            color: white;
            padding: 10px 20px;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            font-size: 16px;
            width: 100%;
            margin-top: 10px;
            transition: background-color 0.3s;
        }

        .btn:hover {
            background-color: #0056b3;
        }

        .error {
            color: #d32f2f;
            background-color: #ffebee;
            padding: 10px;
            border-radius: 4px;
            margin-bottom: 15px;
            border-left: 4px solid #d32f2f;
        }

        .success {
            color: #388e3c;
            background-color: #e8f5e8;
            padding: 10px;
            border-radius: 4px;
            margin-bottom: 15px;
            border-left: 4px solid #388e3c;
        }

        .loading {
            text-align: center;
            color: #666;
            margin-top: 10px;
        }

        @media (max-width: 480px) {
            .container {
                margin: 10px;
                padding: 15px;
            }
            
            body {
                padding: 10px;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <h2 id="idh2cep">Consulta CEP</h2>
        
        {{if .Erro}}
        <div class="error">{{.Erro}}</div>
        {{end}}
        
        {{if .Sucesso}}
        <div class="success">CEP encontrado com sucesso!</div>
        {{end}}

        <form method="POST" action="/">
            <div class="frm-row">
                <label for="cep">CEP</label>
                <input type="text" name="cep" id="cep" maxlength="8" 
                       pattern="[0-9]{8}" title="Digite apenas números (8 dígitos)" 
                       value="{{.CEP}}" placeholder="Ex: 01001000" required>
            </div>

            <div class="frm-row">
                <label for="rua">Endereço</label>
                <input type="text" name="rua" id="rua" value="{{.Rua}}" readonly>
            </div>

            <div class="frm-row">
                <label for="bairro">Bairro</label>
                <input type="text" name="bairro" id="bairro" value="{{.Bairro}}" readonly>
            </div>

            <div class="frm-row">
                <label for="cidade">Cidade</label>
                <input type="text" name="cidade" id="cidade" value="{{.Cidade}}" readonly>
            </div>

            <div class="frm-row">
                <label for="estado">Estado</label>
                <input type="text" name="estado" id="estado" value="{{.Estado}}" readonly>
            </div>

            <button type="submit" class="btn">Consultar CEP</button>
        </form>
    </div>

    <script>
        // Máscara para CEP
        document.getElementById('cep').addEventListener('input', function(e) {
            let value = e.target.value.replace(/\D/g, '');
            if (value.length > 8) {
                value = value.substring(0, 8);
            }
            e.target.value = value;
        });

        // Validação apenas números
        document.getElementById('cep').addEventListener('keypress', function(e) {
            const char = String.fromCharCode(e.which);
            if (!/[0-9]/.test(char)) {
                e.preventDefault();
                document.getElementById('idh2cep').style.backgroundColor = '#ffcdd2';
                setTimeout(() => {
                    document.getElementById('idh2cep').style.backgroundColor = '';
                }, 2000);
            }
        });
    </script>
</body>
</html>
`

// Função para validar CEP
func validaCEP(cep string) bool {
	// Remove espaços e caracteres especiais
	cep = strings.ReplaceAll(cep, " ", "")
	cep = strings.ReplaceAll(cep, "-", "")

	// Verifica se tem 8 dígitos numéricos
	regex := regexp.MustCompile(`^\d{8}$`)
	return regex.MatchString(cep)
}

// Função para consultar CEP na API ViaCEP
func consultaCEP(cep string) (*ViaCEPResponse, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro na resposta da API: status %d", resp.StatusCode)
	}

	var viaCEP ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&viaCEP); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %v", err)
	}

	// Verifica se a API retornou erro
	if viaCEP.Erro {
		return nil, fmt.Errorf("CEP não encontrado")
	}

	return &viaCEP, nil
}

// Handler principal
func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("cep").Parse(htmlTemplate)
	if err != nil {
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		log.Printf("Erro ao parsear template: %v", err)
		return
	}

	data := FormData{}

	if r.Method == http.MethodPost {
		cep := strings.TrimSpace(r.FormValue("cep"))

		if cep == "" {
			data.Erro = "Por favor, digite um CEP."
		} else if !validaCEP(cep) {
			data.Erro = "CEP inválido! Digite apenas números (8 dígitos). Exemplo: 01001000"
			data.CEP = cep
		} else {
			// Consulta o CEP na API
			viaCEP, err := consultaCEP(cep)
			if err != nil {
				data.Erro = fmt.Sprintf("Erro ao consultar CEP: %s", err.Error())
				data.CEP = cep
			} else {
				// Preenche os dados do formulário
				data.CEP = viaCEP.CEP
				data.Rua = viaCEP.Logradouro
				data.Bairro = viaCEP.Bairro
				data.Cidade = viaCEP.Localidade
				data.Estado = viaCEP.UF
				data.Sucesso = true
			}
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Erro ao executar template: %v", err)
	}
}

// Handler para favicon (evita erro 404)
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/favicon.ico", faviconHandler)

	port := ":8080"
	fmt.Printf("🚀 Servidor rodando em http://localhost%s\n", port)
	fmt.Println("📝 Digite um CEP para consultar o endereço")
	fmt.Println("💡 Exemplo de CEP: 01001000 (Praça da Sé, São Paulo)")

	log.Fatal(http.ListenAndServe(port, nil))
}
