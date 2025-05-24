export function handleLogin() {
  const loginForm = document.getElementById('loginForm')

  if (!loginForm) return

  loginForm.onsubmit = async (e) => {
    e.preventDefault()
    const email = loginForm.email.value.trim()
    const password = loginForm.password.value

    // Regex simples para validar email básico
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/

    // Validações
    if (
      !email ||
      !emailRegex.test(email) ||
      email.includes(' ') ||
      password.includes(' ') ||
      password.length < 6
    ) {
      alert('Por favor, preencha todos os campos corretamente. O email deve ser válido, e a senha deve ter pelo menos 6 caracteres sem espaços.')
      return
    }

    try {
      const response = await fetch('/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email, password })
      })

      if (!response.ok) {
        const errorData = await response.json()
        alert(`Erro ao fazer login: ${errorData.message || response.statusText}`)
        return
      }

      const data = await response.json()
      alert(`Login realizado com sucesso!\nBem-vindo, ${data.userName || email}`)
      loginForm.reset()
    } catch (error) {
      alert(`Erro na comunicação com o servidor: ${error.message}`)
    }
  }
}
