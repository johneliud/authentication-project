import { showFeedback } from "./script.js";

const logoutBtn = document.getElementById("logoutBtn");

logoutBtn.addEventListener("click", async () => {
  try {
    const response = await fetch("/logout", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
    });

    const data = await response.json();

    if (response.ok && data.success) {
      showFeedback("Successfully logged out", true);

      setTimeout(() => {
        window.location.href = "/sign-in";
      }, 1000);
    } else {
      const error = await response.json();
      showFeedback(error.message, false);
      return;
    }
  } catch (error) {
    console.error(error);
    showFeedback("Failed to log out", false);
    return;
  }
});
