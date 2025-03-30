import { showFeedback } from "./script.js";

// validatePassword checks if the password is valid
function validatePassword(password) {
  if (password.length < 8) return "Password must be at least 8 characters long";

  if (!password.match(/[A-Z]/))
    return "Password must contain at least one uppercase letter";

  if (!password.match(/[a-z]/))
    return "Password must contain at least one lowercase letter";

  if (!password.match(/[0-9]/))
    return "Password must contain at least one number";

  if (!password.match(/[^A-Za-z0-9]/))
    return "Password must contain at least one special character";

  return "";
}

const passwordInput = document.getElementById("password");
const confirmedPasswordInput = document.getElementById("confirmedPassword");

// validateConfirmedPassword checks if the confirmed password is valid
function validateConfirmedPassword(password, confirmedPassword) {
  if (password !== confirmedPassword) return "Passwords do not match";
  return "";
}

// Password toggle visibility
const visibilityBtns = document.querySelectorAll(".toggle-password-visibility");

visibilityBtns.forEach((btn) => {
  btn.addEventListener("click", (event) => {
    event.preventDefault();
    const input = document.getElementById(btn.dataset.target);
    input.type = input.type === "password" ? "text" : "password";
  });
});

const signupForm = document.getElementById("signupForm");

signupForm.addEventListener("submit", async (event) => {
  event.preventDefault();

  const password = passwordInput.value;
  const confirmedPassword = confirmedPasswordInput.value;
  const passwordError = validatePassword(password);
  const confirmedPasswordError = validateConfirmedPassword(
    password,
    confirmedPassword
  );

  if (passwordError) {
    showFeedback(passwordError, false);
    return;
  }

  if (confirmedPasswordError) {
    showFeedback(confirmedPasswordError, false);
    return;
  }

  const submitBtn = document.getElementById("signupBtn");
  submitBtn.disabled = true;
  submitBtn.textContent = "Creating account...";

  const formData = new FormData(signupForm);

  try {
    const response = await fetch("/sign-up", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(Object.fromEntries(formData)),
    });

    if (response.ok) {
      const data = await response.json();
      showFeedback(data.message, data.success);

      setTimeout(() => {
        window.location.href = "/verify";
      }, 1000);
    } else {
      const error = await response.json();
      showFeedback(error.message, false);
    }
    submitBtn.disabled = false;
    submitBtn.textContent = "Create Account";
  } catch (error) {
    console.error(error);
    showFeedback("Failed to create account. Please try again.", false);
    submitBtn.disabled = false;
    submitBtn.textContent = "Create Account";
  }
});
