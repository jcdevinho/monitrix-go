function showToast(msg, type = 'success') {
  Toastify({
    text: msg,
    duration: 4000,
    gravity: "top",
    position: "right",
    backgroundColor: type === 'success' ? "#4caf50" : "#f44336", // verde ou vermelho
    stopOnFocus: true,
    close: true
  }).showToast()
}

export function handleSignup() {
  const signupForm = document.getElementById('signupForm')
  if (!signupForm) return

  signupForm.onsubmit = async (e) => {
    e.preventDefault()

    const name = signupForm.name.value.trim()
    const email = signupForm.email.value.trim()
    const password = signupForm.password.value

    const nameRegex = /^[A-Za-zÀ-ÿ\s]{1,45}$/
    const hasNumber = /\d/
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
    const allowedDomains = ['gmail.com', 'hotmail.com', 'outlook.com', 'yahoo.com', 'yandex.com', 'zoho.com']
    const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{8,}$/

    if (!name || !nameRegex.test(name) || hasNumber.test(name)) {
      return showToast('Nome inválido. Use apenas letras e espaços.', 'error')
    }

    if (!emailRegex.test(email) || email.includes(' ')) {
      return showToast('Email inválido.', 'error')
    }

    const domain = email.split('@')[1]
    const isAllowedDomain = allowedDomains.includes(domain)
    const isCorporate = domain.endsWith('.com.br') && domain.split('.')[0].length >= 10

    if (!isAllowedDomain && !isCorporate) {
      return showToast('Use um email válido como Gmail ou domínio corporativo com 10+ letras.', 'error')
    }

    if (!passwordRegex.test(password)) {
      return showToast('Senha fraca. Use ao menos 8 caracteres, maiúscula, número e símbolo.', 'error')
    }

    try {
      const response = await fetch('/registro', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({ name, email, password })
      })

      if (!response.ok) {
        const errorData = await response.json()
        return showToast(`Erro: ${errorData.message || 'Tente novamente.'}`, 'error')
      }

      showToast(`Registrado com sucesso!`, 'success')
      signupForm.reset()
    } catch (error) {
      console.error(error)
      showToast('Erro de comunicação com o servidor.', 'error')
    }
  }
}
