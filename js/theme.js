let theme = window.matchMedia("(prefers-color-scheme: dark)").matches ? 'dark' : 'light';
let themeUrl = {
    'dark': 'assets/sun.svg',
    'light': 'assets/moon.svg'
}
document.getElementById('theme-switch').setAttribute('src', themeUrl[theme]);
document.getElementById('theme-switch').onclick = () => {
    theme = theme === 'dark' ? 'light' : 'dark';
    document.getElementById('theme-switch').setAttribute('src', themeUrl[theme]);
    document.getElementById('container').setAttribute('data-theme', theme);
};