function setTokenToLocalStorage(event) {
    const response = event.detail.response;    
    if (response?.token) {
        localStorage.setItem('token', response.token);
    }
}

document.getElementById('login-form')?.addEventListener('htmx:afterRequest', setTokenToLocalStorage);
document.getElementById('signup-form')?.addEventListener('htmx:afterRequest', setTokenToLocalStorage);
document.getElementById('logout-button').addEventListener('htmx:afterRequest', (event) => {
    localStorage.removeItem('token');
});