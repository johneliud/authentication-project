const themeToggler = document.getElementById('themeToggler');

themeToggler.addEventListener('click', () => {
  themeToggler.classList.toggle('active');
  document.body.classList.toggle('dark-mode');

  localStorage.setItem(
    'dark-mode',
    document.body.classList.contains('dark-mode')
  );
});

function applyTheme() {
  if (localStorage.getItem('dark-mode') === 'true') {
    document.body.classList.add('dark-mode');
  } else {
    document.body.classList.remove('dark-mode');
  }
}

document.addEventListener('DOMContentLoaded', applyTheme);

// showFeedback displays a pop-up message with a success or error message.
export function showFeedback(message, isSuccess) {
  const feedbackPopup = document.getElementById('feedbackPopup');

  feedbackPopup.textContent = message;
  feedbackPopup.classList.remove('success', 'error');

  feedbackPopup.classList.add('show', isSuccess ? 'success' : 'error');

  setTimeout(() => {
    feedbackPopup.classList.remove('show', 'success', 'error');
  }, 3000);
}

// Password toggle visibility
const passwordInputs = document.querySelectorAll('.toggle-password-visibility');

passwordInputs.forEach((passInput) => {
  passInput.addEventListener('click', (event) => {
    event.preventDefault();
    const input = document.getElementById(passInput.dataset.target);
    input.type = input.type === 'password' ? 'text' : 'password';
  });
});
