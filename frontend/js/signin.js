import { showFeedback } from "./script.js";

const signinForm = document.getElementById("signinForm");
const signinBtn = document.getElementById("signinBtn");

signinForm.addEventListener("submit", async (event) => {
  event.preventDefault();

  signinBtn.disabled = true;

  const formData = new FormData(signinForm);

  try {
    const response = await fetch("/sign-in", {
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
        window.location.href = "/";
      }, 1000);
    } else {
      const error = await response.json();
      showFeedback(error.message, false);
      signinBtn.disabled = false;
      return;
    }
  } catch (error) {
    console.error("Error:", error);
    showFeedback("Failed to sign in. Please try again.", false);
  }
});
