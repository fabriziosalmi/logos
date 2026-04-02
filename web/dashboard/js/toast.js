// Toast notification system
let timer = null;

export function show(message) {
  const toast = document.getElementById('toast');
  const msg = document.getElementById('toast-msg');
  msg.textContent = String(message);
  toast.classList.add('active');
  clearTimeout(timer);
  timer = setTimeout(() => toast.classList.remove('active'), 3000);
}
